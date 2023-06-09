# InlineResponse404

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**HttpStatus** | Pointer to **int32** |  | [optional] [default to 404]|
|**Message** | Pointer to [**InlineResponse404Message**](InlineResponse404Message.md) |  | [optional] |

## Methods

### NewInlineResponse404

`func NewInlineResponse404() *InlineResponse404`

NewInlineResponse404 instantiates a new InlineResponse404 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse404WithDefaults

`func NewInlineResponse404WithDefaults() *InlineResponse404`

NewInlineResponse404WithDefaults instantiates a new InlineResponse404 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpStatus

`func (o *InlineResponse404) GetHttpStatus() int32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *InlineResponse404) GetHttpStatusOk() (*int32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *InlineResponse404) SetHttpStatus(v int32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *InlineResponse404) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### GetMessage

`func (o *InlineResponse404) GetMessage() InlineResponse404Message`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InlineResponse404) GetMessageOk() (*InlineResponse404Message, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InlineResponse404) SetMessage(v InlineResponse404Message)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InlineResponse404) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


