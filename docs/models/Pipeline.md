# Pipeline

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the pipeline. Must be not more that 20 characters long.  | |
|**Logs** | Pointer to [**[]PipelineNoAddrLogs**](PipelineNoAddrLogs.md) |  | [optional] |
|**TcpAddress** | Pointer to **string** | The TCP address of the pipeline. This is the address to which logs are sent using the TCP protocol.  | [optional] |
|**HttpAddress** | Pointer to **string** | The HTTP address of the pipeline. This is the address to which logs are sent using the HTTP protocol.  | [optional] |
|**GrafanaAddress** | Pointer to **string** | The Grafana address is where user can access their logs, create dashboards, and set up alerts.  | [optional] |
|**ResourceTier** | Pointer to **string** |  | [optional] |
|**Key** | Pointer to **string** | The key is shared once and is used to authenticate the logs sent to the pipeline.  | [optional] |

## Methods

### NewPipeline

`func NewPipeline(name string, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pipeline) SetName(v string)`

SetName sets Name field to given value.


### GetLogs

`func (o *Pipeline) GetLogs() []PipelineNoAddrLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *Pipeline) GetLogsOk() (*[]PipelineNoAddrLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *Pipeline) SetLogs(v []PipelineNoAddrLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *Pipeline) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetTcpAddress

`func (o *Pipeline) GetTcpAddress() string`

GetTcpAddress returns the TcpAddress field if non-nil, zero value otherwise.

### GetTcpAddressOk

`func (o *Pipeline) GetTcpAddressOk() (*string, bool)`

GetTcpAddressOk returns a tuple with the TcpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpAddress

`func (o *Pipeline) SetTcpAddress(v string)`

SetTcpAddress sets TcpAddress field to given value.

### HasTcpAddress

`func (o *Pipeline) HasTcpAddress() bool`

HasTcpAddress returns a boolean if a field has been set.

### GetHttpAddress

`func (o *Pipeline) GetHttpAddress() string`

GetHttpAddress returns the HttpAddress field if non-nil, zero value otherwise.

### GetHttpAddressOk

`func (o *Pipeline) GetHttpAddressOk() (*string, bool)`

GetHttpAddressOk returns a tuple with the HttpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAddress

`func (o *Pipeline) SetHttpAddress(v string)`

SetHttpAddress sets HttpAddress field to given value.

### HasHttpAddress

`func (o *Pipeline) HasHttpAddress() bool`

HasHttpAddress returns a boolean if a field has been set.

### GetGrafanaAddress

`func (o *Pipeline) GetGrafanaAddress() string`

GetGrafanaAddress returns the GrafanaAddress field if non-nil, zero value otherwise.

### GetGrafanaAddressOk

`func (o *Pipeline) GetGrafanaAddressOk() (*string, bool)`

GetGrafanaAddressOk returns a tuple with the GrafanaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaAddress

`func (o *Pipeline) SetGrafanaAddress(v string)`

SetGrafanaAddress sets GrafanaAddress field to given value.

### HasGrafanaAddress

`func (o *Pipeline) HasGrafanaAddress() bool`

HasGrafanaAddress returns a boolean if a field has been set.

### GetResourceTier

`func (o *Pipeline) GetResourceTier() string`

GetResourceTier returns the ResourceTier field if non-nil, zero value otherwise.

### GetResourceTierOk

`func (o *Pipeline) GetResourceTierOk() (*string, bool)`

GetResourceTierOk returns a tuple with the ResourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTier

`func (o *Pipeline) SetResourceTier(v string)`

SetResourceTier sets ResourceTier field to given value.

### HasResourceTier

`func (o *Pipeline) HasResourceTier() bool`

HasResourceTier returns a boolean if a field has been set.

### GetKey

`func (o *Pipeline) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Pipeline) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Pipeline) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Pipeline) HasKey() bool`

HasKey returns a boolean if a field has been set.


