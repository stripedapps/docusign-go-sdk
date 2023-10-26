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

// checks if the PropertyMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyMetadata{}

// PropertyMetadata Metadata about a property.
type PropertyMetadata struct {
	// An array of option strings supported by this setting.
	Options []string `json:"options,omitempty"`
	// Indicates whether the property is editable. Valid values are:  - `editable` - `read_only`
	Rights *string `json:"rights,omitempty"`
}

// NewPropertyMetadata instantiates a new PropertyMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyMetadata() *PropertyMetadata {
	this := PropertyMetadata{}
	return &this
}

// NewPropertyMetadataWithDefaults instantiates a new PropertyMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyMetadataWithDefaults() *PropertyMetadata {
	this := PropertyMetadata{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PropertyMetadata) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyMetadata) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PropertyMetadata) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *PropertyMetadata) SetOptions(v []string) {
	o.Options = v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *PropertyMetadata) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyMetadata) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *PropertyMetadata) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *PropertyMetadata) SetRights(v string) {
	o.Rights = &v
}

func (o PropertyMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	return toSerialize, nil
}

type NullablePropertyMetadata struct {
	value *PropertyMetadata
	isSet bool
}

func (v NullablePropertyMetadata) Get() *PropertyMetadata {
	return v.value
}

func (v *NullablePropertyMetadata) Set(val *PropertyMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyMetadata(val *PropertyMetadata) *NullablePropertyMetadata {
	return &NullablePropertyMetadata{value: val, isSet: true}
}

func (v NullablePropertyMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

