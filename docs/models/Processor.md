# Processor

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Source** | Pointer to **string** | The source parser to be used | [optional] |
|**Tag** | Pointer to **string** | Tag is to distinguish different pipelines. must be unique amongst the pipeline&#39;s array items. | [optional] |
|**Protocol** | Pointer to **string** | Protocol to use as intake | [optional] |

## Methods

### NewProcessor

`func NewProcessor() *Processor`

NewProcessor instantiates a new Processor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorWithDefaults

`func NewProcessorWithDefaults() *Processor`

NewProcessorWithDefaults instantiates a new Processor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *Processor) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Processor) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Processor) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Processor) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *Processor) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Processor) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Processor) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Processor) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProtocol

`func (o *Processor) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Processor) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Processor) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Processor) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


