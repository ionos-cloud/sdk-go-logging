# InlineResponse404Message

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ErrorCode** | Pointer to **string** |  | [optional] [default to "NotFound"]|
|**Message** | Pointer to **string** |  | [optional] [default to "No contract"]|

## Methods

### NewInlineResponse404Message

`func NewInlineResponse404Message() *InlineResponse404Message`

NewInlineResponse404Message instantiates a new InlineResponse404Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse404MessageWithDefaults

`func NewInlineResponse404MessageWithDefaults() *InlineResponse404Message`

NewInlineResponse404MessageWithDefaults instantiates a new InlineResponse404Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *InlineResponse404Message) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *InlineResponse404Message) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *InlineResponse404Message) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *InlineResponse404Message) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMessage

`func (o *InlineResponse404Message) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse404Message) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse404Message) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse404Message) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


