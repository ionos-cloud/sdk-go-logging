# ResponsePipelineAllOf2

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Status** | Pointer to **string** | To display wether the logging stream is Ready/Not Ready | [optional] |

## Methods

### NewResponsePipelineAllOf2

`func NewResponsePipelineAllOf2() *ResponsePipelineAllOf2`

NewResponsePipelineAllOf2 instantiates a new ResponsePipelineAllOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePipelineAllOf2WithDefaults

`func NewResponsePipelineAllOf2WithDefaults() *ResponsePipelineAllOf2`

NewResponsePipelineAllOf2WithDefaults instantiates a new ResponsePipelineAllOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResponsePipelineAllOf2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResponsePipelineAllOf2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResponsePipelineAllOf2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResponsePipelineAllOf2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


