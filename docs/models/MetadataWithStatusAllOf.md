# MetadataWithStatusAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | **string** | The state of the object. The state can be: * &#x60;PROVISIONING&#x60; - resource is being created. * &#x60;AVAILABLE&#x60; - resource exists and is healthy. * &#x60;UPDATING&#x60; - resource is being updated. * &#x60;DESTROYING&#x60; - delete command was issued, the resource is being deleted. * &#x60;FAILED&#x60; - resource failed, details in &#x60;statusMessage&#x60;.  | [readonly] |
|**StatusMessage** | Pointer to **NullableString** | The message of the failure if state is &#x60;FAILED&#x60;.  | [optional] [readonly] |

## Methods

### NewMetadataWithStatusAllOf

`func NewMetadataWithStatusAllOf(state string, ) *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOf instantiates a new MetadataWithStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStatusAllOfWithDefaults

`func NewMetadataWithStatusAllOfWithDefaults() *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOfWithDefaults instantiates a new MetadataWithStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MetadataWithStatusAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStatusAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStatusAllOf) SetState(v string)`

SetState sets State field to given value.


### GetStatusMessage

`func (o *MetadataWithStatusAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MetadataWithStatusAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MetadataWithStatusAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MetadataWithStatusAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *MetadataWithStatusAllOf) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *MetadataWithStatusAllOf) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil

