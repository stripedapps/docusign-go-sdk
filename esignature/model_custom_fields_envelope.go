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

// checks if the CustomFieldsEnvelope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomFieldsEnvelope{}

// CustomFieldsEnvelope 
type CustomFieldsEnvelope struct {
	// An array of list custom fields.
	ListCustomFields []ListCustomField `json:"listCustomFields,omitempty"`
	// An array of text custom fields.
	TextCustomFields []TextCustomField `json:"textCustomFields,omitempty"`
}

// NewCustomFieldsEnvelope instantiates a new CustomFieldsEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFieldsEnvelope() *CustomFieldsEnvelope {
	this := CustomFieldsEnvelope{}
	return &this
}

// NewCustomFieldsEnvelopeWithDefaults instantiates a new CustomFieldsEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldsEnvelopeWithDefaults() *CustomFieldsEnvelope {
	this := CustomFieldsEnvelope{}
	return &this
}

// GetListCustomFields returns the ListCustomFields field value if set, zero value otherwise.
func (o *CustomFieldsEnvelope) GetListCustomFields() []ListCustomField {
	if o == nil || IsNil(o.ListCustomFields) {
		var ret []ListCustomField
		return ret
	}
	return o.ListCustomFields
}

// GetListCustomFieldsOk returns a tuple with the ListCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsEnvelope) GetListCustomFieldsOk() ([]ListCustomField, bool) {
	if o == nil || IsNil(o.ListCustomFields) {
		return nil, false
	}
	return o.ListCustomFields, true
}

// HasListCustomFields returns a boolean if a field has been set.
func (o *CustomFieldsEnvelope) HasListCustomFields() bool {
	if o != nil && !IsNil(o.ListCustomFields) {
		return true
	}

	return false
}

// SetListCustomFields gets a reference to the given []ListCustomField and assigns it to the ListCustomFields field.
func (o *CustomFieldsEnvelope) SetListCustomFields(v []ListCustomField) {
	o.ListCustomFields = v
}

// GetTextCustomFields returns the TextCustomFields field value if set, zero value otherwise.
func (o *CustomFieldsEnvelope) GetTextCustomFields() []TextCustomField {
	if o == nil || IsNil(o.TextCustomFields) {
		var ret []TextCustomField
		return ret
	}
	return o.TextCustomFields
}

// GetTextCustomFieldsOk returns a tuple with the TextCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFieldsEnvelope) GetTextCustomFieldsOk() ([]TextCustomField, bool) {
	if o == nil || IsNil(o.TextCustomFields) {
		return nil, false
	}
	return o.TextCustomFields, true
}

// HasTextCustomFields returns a boolean if a field has been set.
func (o *CustomFieldsEnvelope) HasTextCustomFields() bool {
	if o != nil && !IsNil(o.TextCustomFields) {
		return true
	}

	return false
}

// SetTextCustomFields gets a reference to the given []TextCustomField and assigns it to the TextCustomFields field.
func (o *CustomFieldsEnvelope) SetTextCustomFields(v []TextCustomField) {
	o.TextCustomFields = v
}

func (o CustomFieldsEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomFieldsEnvelope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListCustomFields) {
		toSerialize["listCustomFields"] = o.ListCustomFields
	}
	if !IsNil(o.TextCustomFields) {
		toSerialize["textCustomFields"] = o.TextCustomFields
	}
	return toSerialize, nil
}

type NullableCustomFieldsEnvelope struct {
	value *CustomFieldsEnvelope
	isSet bool
}

func (v NullableCustomFieldsEnvelope) Get() *CustomFieldsEnvelope {
	return v.value
}

func (v *NullableCustomFieldsEnvelope) Set(val *CustomFieldsEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFieldsEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFieldsEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFieldsEnvelope(val *CustomFieldsEnvelope) *NullableCustomFieldsEnvelope {
	return &NullableCustomFieldsEnvelope{value: val, isSet: true}
}

func (v NullableCustomFieldsEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFieldsEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


