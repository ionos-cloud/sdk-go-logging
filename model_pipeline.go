/*
 * IONOS Logging REST API
 *
 * Logging as a Service (LaaS) is a service that provides a centralized logging system where users are able to push and aggregate their system or application logs. This service also provides a visualization platform where users are able to observe, search and filter the logs and also create dashboards and alerts for their data points. This service can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an API. The API allows you to create logging pipelines or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// Pipeline pipeline response
type Pipeline struct {
	// The unique ID of the resource.
	Id         *string             `json:"id,omitempty"`
	Metadata   *Metadata           `json:"metadata,omitempty"`
	Properties *PipelineProperties `json:"properties,omitempty"`
}

// NewPipeline instantiates a new Pipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeline() *Pipeline {
	this := Pipeline{}

	return &this
}

// NewPipelineWithDefaults instantiates a new Pipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineWithDefaults() *Pipeline {
	this := Pipeline{}
	return &this
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Pipeline) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *Pipeline) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *Pipeline) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for Metadata will be returned
func (o *Pipeline) GetMetadata() *Metadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pipeline) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *Pipeline) SetMetadata(v Metadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *Pipeline) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, the zero value for PipelineProperties will be returned
func (o *Pipeline) GetProperties() *PipelineProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Pipeline) GetPropertiesOk() (*PipelineProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *Pipeline) SetProperties(v PipelineProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *Pipeline) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

func (o Pipeline) MarshalJSON() ([]byte, error) {
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

type NullablePipeline struct {
	value *Pipeline
	isSet bool
}

func (v NullablePipeline) Get() *Pipeline {
	return v.value
}

func (v *NullablePipeline) Set(val *Pipeline) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeline(val *Pipeline) *NullablePipeline {
	return &NullablePipeline{value: val, isSet: true}
}

func (v NullablePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
