/*
DocuSign REST API

The DocuSign REST API provides you with a powerful, convenient, and simple Web services API for interacting with DocuSign.

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the EnvelopeCustomMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeCustomMetadata{}

// EnvelopeCustomMetadata 
type EnvelopeCustomMetadata struct {
	// 
	EnvelopeCustomMetadataDetails []NameValue `json:"envelopeCustomMetadataDetails,omitempty"`
}

// NewEnvelopeCustomMetadata instantiates a new EnvelopeCustomMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeCustomMetadata() *EnvelopeCustomMetadata {
	this := EnvelopeCustomMetadata{}
	return &this
}

// NewEnvelopeCustomMetadataWithDefaults instantiates a new EnvelopeCustomMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeCustomMetadataWithDefaults() *EnvelopeCustomMetadata {
	this := EnvelopeCustomMetadata{}
	return &this
}

// GetEnvelopeCustomMetadataDetails returns the EnvelopeCustomMetadataDetails field value if set, zero value otherwise.
func (o *EnvelopeCustomMetadata) GetEnvelopeCustomMetadataDetails() []NameValue {
	if o == nil || IsNil(o.EnvelopeCustomMetadataDetails) {
		var ret []NameValue
		return ret
	}
	return o.EnvelopeCustomMetadataDetails
}

// GetEnvelopeCustomMetadataDetailsOk returns a tuple with the EnvelopeCustomMetadataDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeCustomMetadata) GetEnvelopeCustomMetadataDetailsOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.EnvelopeCustomMetadataDetails) {
		return nil, false
	}
	return o.EnvelopeCustomMetadataDetails, true
}

// HasEnvelopeCustomMetadataDetails returns a boolean if a field has been set.
func (o *EnvelopeCustomMetadata) HasEnvelopeCustomMetadataDetails() bool {
	if o != nil && !IsNil(o.EnvelopeCustomMetadataDetails) {
		return true
	}

	return false
}

// SetEnvelopeCustomMetadataDetails gets a reference to the given []NameValue and assigns it to the EnvelopeCustomMetadataDetails field.
func (o *EnvelopeCustomMetadata) SetEnvelopeCustomMetadataDetails(v []NameValue) {
	o.EnvelopeCustomMetadataDetails = v
}

func (o EnvelopeCustomMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeCustomMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvelopeCustomMetadataDetails) {
		toSerialize["envelopeCustomMetadataDetails"] = o.EnvelopeCustomMetadataDetails
	}
	return toSerialize, nil
}

type NullableEnvelopeCustomMetadata struct {
	value *EnvelopeCustomMetadata
	isSet bool
}

func (v NullableEnvelopeCustomMetadata) Get() *EnvelopeCustomMetadata {
	return v.value
}

func (v *NullableEnvelopeCustomMetadata) Set(val *EnvelopeCustomMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeCustomMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeCustomMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeCustomMetadata(val *EnvelopeCustomMetadata) *NullableEnvelopeCustomMetadata {
	return &NullableEnvelopeCustomMetadata{value: val, isSet: true}
}

func (v NullableEnvelopeCustomMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeCustomMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

