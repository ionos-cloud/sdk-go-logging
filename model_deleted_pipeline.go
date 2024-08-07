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

// DeletedPipeline pipeline response
type DeletedPipeline struct {
	// The unique ID of the resource.
	Id         *string             `json:"id,omitempty"`
	Metadata   *DeletedMetadata    `json:"metadata,omitempty"`
	Properties *PipelineProperties `json:"properties,omitempty"`
}

// NewDeletedPipeline instantiates a new DeletedPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeletedPipeline() *DeletedPipeline {
	this := DeletedPipeline{}

	return &this
}

// NewDeletedPipelineWithDefaults instantiates a new DeletedPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeletedPipelineWithDefaults() *DeletedPipeline {
	this := DeletedPipeline{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *DeletedPipeline) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedPipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *DeletedPipeline) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *DeletedPipeline) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for DeletedMetadata will be returned
func (o *DeletedPipeline) GetMetadata() *DeletedMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedPipeline) GetMetadataOk() (*DeletedMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *DeletedPipeline) SetMetadata(v DeletedMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *DeletedPipeline) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for PipelineProperties will be returned
func (o *DeletedPipeline) GetProperties() *PipelineProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeletedPipeline) GetPropertiesOk() (*PipelineProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *DeletedPipeline) SetProperties(v PipelineProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *DeletedPipeline) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o DeletedPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	return json.Marshal(toSerialize)
}

type NullableDeletedPipeline struct {
	value *DeletedPipeline
	isSet bool
}

func (v NullableDeletedPipeline) Get() *DeletedPipeline {
	return v.value
}

func (v *NullableDeletedPipeline) Set(val *DeletedPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableDeletedPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableDeletedPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeletedPipeline(val *DeletedPipeline) *NullableDeletedPipeline {
	return &NullableDeletedPipeline{value: val, isSet: true}
}

func (v NullableDeletedPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeletedPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
