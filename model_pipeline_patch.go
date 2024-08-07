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

// PipelinePatch Request payload with any data that is possible to patch a logging pipeline
type PipelinePatch struct {
	Properties *PipelinePatchProperties `json:"properties"`
}

// NewPipelinePatch instantiates a new PipelinePatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelinePatch(properties PipelinePatchProperties) *PipelinePatch {
	this := PipelinePatch{}

	this.Properties = &properties

	return &this
}

// NewPipelinePatchWithDefaults instantiates a new PipelinePatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelinePatchWithDefaults() *PipelinePatch {
	this := PipelinePatch{}
	return &this
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for PipelinePatchProperties will be returned
func (o *PipelinePatch) GetProperties() *PipelinePatchProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PipelinePatch) GetPropertiesOk() (*PipelinePatchProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *PipelinePatch) SetProperties(v PipelinePatchProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *PipelinePatch) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o PipelinePatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullablePipelinePatch struct {
	value *PipelinePatch
	isSet bool
}

func (v NullablePipelinePatch) Get() *PipelinePatch {
	return v.value
}

func (v *NullablePipelinePatch) Set(val *PipelinePatch) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelinePatch) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelinePatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelinePatch(val *PipelinePatch) *NullablePipelinePatch {
	return &NullablePipelinePatch{value: val, isSet: true}
}

func (v NullablePipelinePatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelinePatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
