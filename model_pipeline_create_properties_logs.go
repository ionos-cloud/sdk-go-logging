/*
 * IONOS Logging REST API
 *
 * The logging service offers a centralized platform to collect and store logs from various systems and applications. It includes tools to search, filter, visualize, and create alerts based on your log data.  This API provides programmatic control over logging pipelines, enabling you to create new pipelines or modify existing ones. It mirrors the functionality of the DCD visual tool, ensuring a consistent experience regardless of your chosen interface.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// PipelineCreatePropertiesLogs struct for PipelineCreatePropertiesLogs
type PipelineCreatePropertiesLogs struct {
	// The source parser to be used
	Source *string `json:"source,omitempty"`
	// Tag is to distinguish different pipelines. must be unique amongst the pipeline's array items.
	Tag *string `json:"tag,omitempty"`
	// Protocol to use as intake
	Protocol *string `json:"protocol,omitempty"`
	// Optional custom labels to filter and report logs
	Labels *[]string `json:"labels,omitempty"`
	// The configuration of the logs datastore
	Destinations *[]Destination `json:"destinations,omitempty"`
}

// NewPipelineCreatePropertiesLogs instantiates a new PipelineCreatePropertiesLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineCreatePropertiesLogs() *PipelineCreatePropertiesLogs {
	this := PipelineCreatePropertiesLogs{}

	return &this
}

// NewPipelineCreatePropertiesLogsWithDefaults instantiates a new PipelineCreatePropertiesLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineCreatePropertiesLogsWithDefaults() *PipelineCreatePropertiesLogs {
	this := PipelineCreatePropertiesLogs{}
	return &this
}

// GetSource returns the Source field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineCreatePropertiesLogs) GetSource() *string {
	if o == nil {
		return nil
	}

	return o.Source

}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreatePropertiesLogs) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Source, true
}

// SetSource sets field value
func (o *PipelineCreatePropertiesLogs) SetSource(v string) {

	o.Source = &v

}

// HasSource returns a boolean if a field has been set.
func (o *PipelineCreatePropertiesLogs) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// GetTag returns the Tag field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineCreatePropertiesLogs) GetTag() *string {
	if o == nil {
		return nil
	}

	return o.Tag

}

// GetTagOk returns a tuple with the Tag field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreatePropertiesLogs) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Tag, true
}

// SetTag sets field value
func (o *PipelineCreatePropertiesLogs) SetTag(v string) {

	o.Tag = &v

}

// HasTag returns a boolean if a field has been set.
func (o *PipelineCreatePropertiesLogs) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// GetProtocol returns the Protocol field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PipelineCreatePropertiesLogs) GetProtocol() *string {
	if o == nil {
		return nil
	}

	return o.Protocol

}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreatePropertiesLogs) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Protocol, true
}

// SetProtocol sets field value
func (o *PipelineCreatePropertiesLogs) SetProtocol(v string) {

	o.Protocol = &v

}

// HasProtocol returns a boolean if a field has been set.
func (o *PipelineCreatePropertiesLogs) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// GetLabels returns the Labels field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *PipelineCreatePropertiesLogs) GetLabels() *[]string {
	if o == nil {
		return nil
	}

	return o.Labels

}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreatePropertiesLogs) GetLabelsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Labels, true
}

// SetLabels sets field value
func (o *PipelineCreatePropertiesLogs) SetLabels(v []string) {

	o.Labels = &v

}

// HasLabels returns a boolean if a field has been set.
func (o *PipelineCreatePropertiesLogs) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// GetDestinations returns the Destinations field value
// If the value is explicit nil, the zero value for []Destination will be returned
func (o *PipelineCreatePropertiesLogs) GetDestinations() *[]Destination {
	if o == nil {
		return nil
	}

	return o.Destinations

}

// GetDestinationsOk returns a tuple with the Destinations field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelineCreatePropertiesLogs) GetDestinationsOk() (*[]Destination, bool) {
	if o == nil {
		return nil, false
	}

	return o.Destinations, true
}

// SetDestinations sets field value
func (o *PipelineCreatePropertiesLogs) SetDestinations(v []Destination) {

	o.Destinations = &v

}

// HasDestinations returns a boolean if a field has been set.
func (o *PipelineCreatePropertiesLogs) HasDestinations() bool {
	if o != nil && o.Destinations != nil {
		return true
	}

	return false
}

func (o PipelineCreatePropertiesLogs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}

	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}

	if o.Destinations != nil {
		toSerialize["destinations"] = o.Destinations
	}

	return json.Marshal(toSerialize)
}

type NullablePipelineCreatePropertiesLogs struct {
	value *PipelineCreatePropertiesLogs
	isSet bool
}

func (v NullablePipelineCreatePropertiesLogs) Get() *PipelineCreatePropertiesLogs {
	return v.value
}

func (v *NullablePipelineCreatePropertiesLogs) Set(val *PipelineCreatePropertiesLogs) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCreatePropertiesLogs) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCreatePropertiesLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCreatePropertiesLogs(val *PipelineCreatePropertiesLogs) *NullablePipelineCreatePropertiesLogs {
	return &NullablePipelineCreatePropertiesLogs{value: val, isSet: true}
}

func (v NullablePipelineCreatePropertiesLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCreatePropertiesLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
