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
	"time"
)

// Metadata Metadata of the resource
type Metadata struct {
	// The ISO 8601 creation timestamp.
	CreatedDate       *IonosTime `json:"createdDate,omitempty"`
	CreatedBy         *string    `json:"createdBy,omitempty"`
	CreatedByUserId   *string    `json:"createdByUserId,omitempty"`
	CreatedByUserUuid *string    `json:"createdByUserUuid,omitempty"`
	// The ISO 8601 modified timestamp.
	LastModifiedDate       *IonosTime `json:"lastModifiedDate,omitempty"`
	LastModifiedBy         *string    `json:"lastModifiedBy,omitempty"`
	LastModifiedByUserId   *string    `json:"lastModifiedByUserId,omitempty"`
	LastModifiedByUserUuid *string    `json:"lastModifiedByUserUuid,omitempty"`
	// The current state reported back by the pipeline.
	State *string `json:"state,omitempty"`
}

// NewMetadata instantiates a new Metadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadata() *Metadata {
	this := Metadata{}

	return &this
}

// NewMetadataWithDefaults instantiates a new Metadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataWithDefaults() *Metadata {
	this := Metadata{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Metadata) GetCreatedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.CreatedDate == nil {
		return nil
	}
	return &o.CreatedDate.Time

}

// GetCreatedDateOk returns a tuple with the CreatedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.CreatedDate == nil {
		return nil, false
	}
	return &o.CreatedDate.Time, true

}

// SetCreatedDate sets field value
func (o *Metadata) SetCreatedDate(v time.Time) {

	o.CreatedDate = &IonosTime{v}

}

// HasCreatedDate returns a boolean if a field has been set.
func (o *Metadata) HasCreatedDate() bool {
	if o != nil && o.CreatedDate != nil {
		return true
	}

	return false
}

// GetCreatedBy returns the CreatedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetCreatedBy() *string {
	if o == nil {
		return nil
	}

	return o.CreatedBy

}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Metadata) SetCreatedBy(v string) {

	o.CreatedBy = &v

}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Metadata) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// GetCreatedByUserId returns the CreatedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetCreatedByUserId() *string {
	if o == nil {
		return nil
	}

	return o.CreatedByUserId

}

// GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedByUserId, true
}

// SetCreatedByUserId sets field value
func (o *Metadata) SetCreatedByUserId(v string) {

	o.CreatedByUserId = &v

}

// HasCreatedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasCreatedByUserId() bool {
	if o != nil && o.CreatedByUserId != nil {
		return true
	}

	return false
}

// GetCreatedByUserUuid returns the CreatedByUserUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetCreatedByUserUuid() *string {
	if o == nil {
		return nil
	}

	return o.CreatedByUserUuid

}

// GetCreatedByUserUuidOk returns a tuple with the CreatedByUserUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetCreatedByUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.CreatedByUserUuid, true
}

// SetCreatedByUserUuid sets field value
func (o *Metadata) SetCreatedByUserUuid(v string) {

	o.CreatedByUserUuid = &v

}

// HasCreatedByUserUuid returns a boolean if a field has been set.
func (o *Metadata) HasCreatedByUserUuid() bool {
	if o != nil && o.CreatedByUserUuid != nil {
		return true
	}

	return false
}

// GetLastModifiedDate returns the LastModifiedDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Metadata) GetLastModifiedDate() *time.Time {
	if o == nil {
		return nil
	}

	if o.LastModifiedDate == nil {
		return nil
	}
	return &o.LastModifiedDate.Time

}

// GetLastModifiedDateOk returns a tuple with the LastModifiedDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}

	if o.LastModifiedDate == nil {
		return nil, false
	}
	return &o.LastModifiedDate.Time, true

}

// SetLastModifiedDate sets field value
func (o *Metadata) SetLastModifiedDate(v time.Time) {

	o.LastModifiedDate = &IonosTime{v}

}

// HasLastModifiedDate returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedDate() bool {
	if o != nil && o.LastModifiedDate != nil {
		return true
	}

	return false
}

// GetLastModifiedBy returns the LastModifiedBy field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetLastModifiedBy() *string {
	if o == nil {
		return nil
	}

	return o.LastModifiedBy

}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LastModifiedBy, true
}

// SetLastModifiedBy sets field value
func (o *Metadata) SetLastModifiedBy(v string) {

	o.LastModifiedBy = &v

}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// GetLastModifiedByUserId returns the LastModifiedByUserId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetLastModifiedByUserId() *string {
	if o == nil {
		return nil
	}

	return o.LastModifiedByUserId

}

// GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedByUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LastModifiedByUserId, true
}

// SetLastModifiedByUserId sets field value
func (o *Metadata) SetLastModifiedByUserId(v string) {

	o.LastModifiedByUserId = &v

}

// HasLastModifiedByUserId returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedByUserId() bool {
	if o != nil && o.LastModifiedByUserId != nil {
		return true
	}

	return false
}

// GetLastModifiedByUserUuid returns the LastModifiedByUserUuid field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetLastModifiedByUserUuid() *string {
	if o == nil {
		return nil
	}

	return o.LastModifiedByUserUuid

}

// GetLastModifiedByUserUuidOk returns a tuple with the LastModifiedByUserUuid field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetLastModifiedByUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.LastModifiedByUserUuid, true
}

// SetLastModifiedByUserUuid sets field value
func (o *Metadata) SetLastModifiedByUserUuid(v string) {

	o.LastModifiedByUserUuid = &v

}

// HasLastModifiedByUserUuid returns a boolean if a field has been set.
func (o *Metadata) HasLastModifiedByUserUuid() bool {
	if o != nil && o.LastModifiedByUserUuid != nil {
		return true
	}

	return false
}

// GetState returns the State field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Metadata) GetState() *string {
	if o == nil {
		return nil
	}

	return o.State

}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Metadata) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.State, true
}

// SetState sets field value
func (o *Metadata) SetState(v string) {

	o.State = &v

}

// HasState returns a boolean if a field has been set.
func (o *Metadata) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

func (o Metadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDate != nil {
		toSerialize["createdDate"] = o.CreatedDate
	}

	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}

	if o.CreatedByUserId != nil {
		toSerialize["createdByUserId"] = o.CreatedByUserId
	}

	if o.CreatedByUserUuid != nil {
		toSerialize["createdByUserUuid"] = o.CreatedByUserUuid
	}

	if o.LastModifiedDate != nil {
		toSerialize["lastModifiedDate"] = o.LastModifiedDate
	}

	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}

	if o.LastModifiedByUserId != nil {
		toSerialize["lastModifiedByUserId"] = o.LastModifiedByUserId
	}

	if o.LastModifiedByUserUuid != nil {
		toSerialize["lastModifiedByUserUuid"] = o.LastModifiedByUserUuid
	}

	if o.State != nil {
		toSerialize["state"] = o.State
	}

	return json.Marshal(toSerialize)
}

type NullableMetadata struct {
	value *Metadata
	isSet bool
}

func (v NullableMetadata) Get() *Metadata {
	return v.value
}

func (v *NullableMetadata) Set(val *Metadata) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadata(val *Metadata) *NullableMetadata {
	return &NullableMetadata{value: val, isSet: true}
}

func (v NullableMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
