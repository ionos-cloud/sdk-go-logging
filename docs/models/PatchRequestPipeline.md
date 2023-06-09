# PatchRequestPipeline

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Source** | Pointer to **string** | The source parser to be used | [optional] |
|**Tag** | Pointer to **string** | Tag is to distinguish different pipelines. must be unique amongst the pipeline&#39;s array items. | [optional] |
|**Protocol** | Pointer to **string** | Protocol to use as intake | [optional] |

## Methods

### NewPatchRequestPipeline

`func NewPatchRequestPipeline() *PatchRequestPipeline`

NewPatchRequestPipeline instantiates a new PatchRequestPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestPipelineWithDefaults

`func NewPatchRequestPipelineWithDefaults() *PatchRequestPipeline`

NewPatchRequestPipelineWithDefaults instantiates a new PatchRequestPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PatchRequestPipeline) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatchRequestPipeline) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatchRequestPipeline) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatchRequestPipeline) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *PatchRequestPipeline) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PatchRequestPipeline) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PatchRequestPipeline) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PatchRequestPipeline) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProtocol

`func (o *PatchRequestPipeline) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PatchRequestPipeline) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PatchRequestPipeline) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PatchRequestPipeline) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


