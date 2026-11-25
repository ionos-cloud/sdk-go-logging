package ionoscloud

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const testPipelineID = "7d3cfb2e-d823-11f1-8ce7-cad43111a286"

func strPtr(s string) *string { return &s }

// newLogs returns one log stream, as the operator creates it.
func newLogs() PipelineNoAddrLogs {
	return *NewPipelineNoAddrLogs("generic", "logwrite", "http", []PipelineNoAddrLogsDestinations{
		{Type: strPtr("loki"), RetentionInDays: int32Ptr(7)},
	})
}

func int32Ptr(i int32) *int32 { return &i }

// marshalMap marshals v and decodes the result into a generic map.
func marshalMap(t *testing.T, v interface{}) map[string]interface{} {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal %s: %v", b, err)
	}
	return m
}

func TestPipelineNoAddrLogsLabelsAccessors(t *testing.T) {
	logs := newLogs()

	if logs.HasLabels() {
		t.Fatal("HasLabels() = true on a new log stream, want false")
	}
	if got := logs.GetLabels(); got != nil {
		t.Fatalf("GetLabels() = %v on a new log stream, want nil", *got)
	}

	logs.SetLabels([]string{"host", "src"})

	if !logs.HasLabels() {
		t.Fatal("HasLabels() = false after SetLabels, want true")
	}
	want := []string{"host", "src"}
	if got := logs.GetLabels(); got == nil || !reflect.DeepEqual(*got, want) {
		t.Fatalf("GetLabels() = %v, want %v", got, want)
	}
	got, ok := logs.GetLabelsOk()
	if !ok || got == nil || !reflect.DeepEqual(*got, want) {
		t.Fatalf("GetLabelsOk() = %v, %v, want %v, true", got, ok, want)
	}
}

func TestPipelineNoAddrLogsLabelsNilReceiver(t *testing.T) {
	var logs *PipelineNoAddrLogs

	if logs.HasLabels() {
		t.Error("HasLabels() on nil receiver = true, want false")
	}
	if got := logs.GetLabels(); got != nil {
		t.Errorf("GetLabels() on nil receiver = %v, want nil", got)
	}
	if got, ok := logs.GetLabelsOk(); got != nil || ok {
		t.Errorf("GetLabelsOk() on nil receiver = %v, %v, want nil, false", got, ok)
	}
}

func TestPipelineNoAddrLogsLabelsMarshal(t *testing.T) {
	tests := []struct {
		name      string
		labels    *[]string
		wantKey   bool
		wantValue []interface{}
	}{
		{name: "unset is omitted", labels: nil, wantKey: false},
		{name: "host and src", labels: &[]string{"host", "src"}, wantKey: true, wantValue: []interface{}{"host", "src"}},
		{name: "three labels, the API maximum", labels: &[]string{"a", "b", "c"}, wantKey: true, wantValue: []interface{}{"a", "b", "c"}},
		{name: "empty list is kept", labels: &[]string{}, wantKey: true, wantValue: []interface{}{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logs := newLogs()
			logs.Labels = tt.labels

			m := marshalMap(t, logs)

			got, present := m["labels"]
			if present != tt.wantKey {
				t.Fatalf("labels key present = %v, want %v (body %v)", present, tt.wantKey, m)
			}
			if tt.wantKey && !reflect.DeepEqual(got, tt.wantValue) {
				t.Fatalf("labels = %v, want %v", got, tt.wantValue)
			}
			// The required fields must be unaffected by the new field.
			for key, want := range map[string]interface{}{"source": "generic", "tag": "logwrite", "protocol": "http"} {
				if m[key] != want {
					t.Errorf("%s = %v, want %v", key, m[key], want)
				}
			}
			if _, ok := m["destinations"]; !ok {
				t.Error("destinations missing from body")
			}
		})
	}
}

func TestPipelineNoAddrLogsLabelsUnmarshal(t *testing.T) {
	tests := []struct {
		name string
		body string
		want *[]string
	}{
		{name: "with labels", body: `{"source":"generic","tag":"logwrite","protocol":"http","labels":["host","src"],"destinations":[]}`, want: &[]string{"host", "src"}},
		{name: "without labels", body: `{"source":"generic","tag":"logwrite","protocol":"http","destinations":[]}`, want: nil},
		{name: "empty list", body: `{"source":"generic","tag":"logwrite","protocol":"http","labels":[],"destinations":[]}`, want: &[]string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var logs PipelineNoAddrLogs
			if err := json.Unmarshal([]byte(tt.body), &logs); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if !reflect.DeepEqual(logs.Labels, tt.want) {
				t.Fatalf("Labels = %v, want %v", logs.Labels, tt.want)
			}
		})
	}
}

func TestPipelineNoAddrLogsLabelsRoundTrip(t *testing.T) {
	in := newLogs()
	in.SetLabels([]string{"host", "src"})

	b, err := json.Marshal(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var out PipelineNoAddrLogs
	if err := json.Unmarshal(b, &out); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if !reflect.DeepEqual(in, out) {
		t.Fatalf("round trip changed the value:\n in  %+v\n out %+v", in, out)
	}
}

// seen is what the test server received.
type seen struct {
	method string
	path   string
	header http.Header
	body   map[string]interface{}
}

// newTestAPI returns a client for a server that records the request and
// answers with status and response.
func newTestAPI(t *testing.T, status int, response string) (*APIClient, *seen) {
	t.Helper()
	got := &seen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method = r.Method
		got.path = r.URL.Path
		got.header = r.Header.Clone()
		if b, err := io.ReadAll(r.Body); err == nil && len(b) > 0 {
			_ = json.Unmarshal(b, &got.body)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, _ = io.WriteString(w, response)
	}))
	t.Cleanup(srv.Close)

	cfg := NewConfiguration("", "", "test-token", srv.URL)
	cfg.MaxRetries = 0
	return NewAPIClient(cfg), got
}

// pipelineReadJSON is a pipeline as the API returns it. labels is the raw JSON
// of the labels value, an empty string leaves the key out.
func pipelineReadJSON(labels string) string {
	l := ""
	if labels != "" {
		l = `"labels":` + labels + `,`
	}
	return `{"id":"` + testPipelineID + `","type":"pipeline","href":"/pipelines/` + testPipelineID + `",` +
		`"metadata":{"state":"AVAILABLE"},` +
		`"properties":{"name":"nas-b6afd1d1-9939-44","httpAddress":"x-logs.example.test",` +
		`"logs":[{"source":"generic","tag":"logwrite","protocol":"http",` + l +
		`"destinations":[{"type":"loki","retentionInDays":7}]}]}}`
}

// requestBody builds the properties object of a create or patch request.
func requestBody(labels *[]string) PipelineNoAddr {
	logs := newLogs()
	logs.Labels = labels
	props := NewPipelineNoAddr("nas-b6afd1d1-9939-44")
	props.SetLogs([]PipelineNoAddrLogs{logs})
	return *props
}

// sentLogs returns the first log stream of the properties object the client sent.
func sentLogs(t *testing.T, body map[string]interface{}) map[string]interface{} {
	t.Helper()
	props, ok := body["properties"].(map[string]interface{})
	if !ok {
		t.Fatalf("request has no properties object: %v", body)
	}
	logs, ok := props["logs"].([]interface{})
	if !ok || len(logs) != 1 {
		t.Fatalf("properties.logs is not a list of one stream: %v", props["logs"])
	}
	log, ok := logs[0].(map[string]interface{})
	if !ok {
		t.Fatalf("log stream is not an object: %v", logs[0])
	}
	return log
}

func readLabels(t *testing.T, p PipelineRead) *[]string {
	t.Helper()
	if p.Properties == nil || p.Properties.Logs == nil || len(*p.Properties.Logs) != 1 {
		t.Fatalf("response has no single log stream: %+v", p.Properties)
	}
	return (*p.Properties.Logs)[0].Labels
}

func TestPipelinesPostSendsLabels(t *testing.T) {
	api, got := newTestAPI(t, http.StatusAccepted, pipelineReadJSON(`["host","src"]`))

	created, resp, err := api.PipelinesApi.PipelinesPost(context.Background()).
		PipelineCreate(*NewPipelineCreate(requestBody(&[]string{"host", "src"}))).Execute()
	if err != nil {
		t.Fatalf("PipelinesPost: %v", err)
	}
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusAccepted)
	}

	if got.method != http.MethodPost || got.path != "/pipelines" {
		t.Errorf("request = %s %s, want POST /pipelines", got.method, got.path)
	}
	if auth := got.header.Get("Authorization"); auth != "Bearer test-token" {
		t.Errorf("Authorization = %q, want a bearer token", auth)
	}
	if labels := sentLogs(t, got.body)["labels"]; !reflect.DeepEqual(labels, []interface{}{"host", "src"}) {
		t.Errorf("sent labels = %v, want [host src]", labels)
	}

	if labels := readLabels(t, created); labels == nil || !reflect.DeepEqual(*labels, []string{"host", "src"}) {
		t.Errorf("returned labels = %v, want [host src]", labels)
	}
}

func TestPipelinesPostWithoutLabelsSendsNoLabelsKey(t *testing.T) {
	api, got := newTestAPI(t, http.StatusAccepted, pipelineReadJSON(``))

	created, _, err := api.PipelinesApi.PipelinesPost(context.Background()).
		PipelineCreate(*NewPipelineCreate(requestBody(nil))).Execute()
	if err != nil {
		t.Fatalf("PipelinesPost: %v", err)
	}

	if _, present := sentLogs(t, got.body)["labels"]; present {
		t.Error("labels key was sent although no labels were set")
	}
	if labels := readLabels(t, created); labels != nil {
		t.Errorf("returned labels = %v, want nil", *labels)
	}
}

func TestPipelinesPatchSendsFullLogsWithLabels(t *testing.T) {
	api, got := newTestAPI(t, http.StatusOK, pipelineReadJSON(`["host","src"]`))

	patched, _, err := api.PipelinesApi.PipelinesPatch(context.Background(), testPipelineID).
		PipelinePatch(*NewPipelinePatch(requestBody(&[]string{"host", "src"}))).Execute()
	if err != nil {
		t.Fatalf("PipelinesPatch: %v", err)
	}

	if got.method != http.MethodPatch || got.path != "/pipelines/"+testPipelineID {
		t.Errorf("request = %s %s, want PATCH /pipelines/%s", got.method, got.path, testPipelineID)
	}

	// The API replaces the values it is given, so the request has to carry the
	// name and the whole log stream, not only the labels.
	props := got.body["properties"].(map[string]interface{})
	if props["name"] != "nas-b6afd1d1-9939-44" {
		t.Errorf("name = %v, want nas-b6afd1d1-9939-44", props["name"])
	}
	log := sentLogs(t, got.body)
	for key, want := range map[string]interface{}{
		"source":       "generic",
		"tag":          "logwrite",
		"protocol":     "http",
		"labels":       []interface{}{"host", "src"},
		"destinations": []interface{}{map[string]interface{}{"type": "loki", "retentionInDays": float64(7)}},
	} {
		if !reflect.DeepEqual(log[key], want) {
			t.Errorf("%s = %v, want %v", key, log[key], want)
		}
	}

	if labels := readLabels(t, patched); labels == nil || !reflect.DeepEqual(*labels, []string{"host", "src"}) {
		t.Errorf("returned labels = %v, want [host src]", labels)
	}
}

func TestPipelinesFindByIdDecodesLabels(t *testing.T) {
	tests := []struct {
		name   string
		labels string
		want   *[]string
	}{
		{name: "with labels", labels: `["host","src"]`, want: &[]string{"host", "src"}},
		{name: "without labels", labels: ``, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, got := newTestAPI(t, http.StatusOK, pipelineReadJSON(tt.labels))

			p, _, err := api.PipelinesApi.PipelinesFindById(context.Background(), testPipelineID).Execute()
			if err != nil {
				t.Fatalf("PipelinesFindById: %v", err)
			}

			if got.method != http.MethodGet || got.path != "/pipelines/"+testPipelineID {
				t.Errorf("request = %s %s, want GET /pipelines/%s", got.method, got.path, testPipelineID)
			}
			if labels := readLabels(t, p); !reflect.DeepEqual(labels, tt.want) {
				t.Errorf("labels = %v, want %v", labels, tt.want)
			}
		})
	}
}
