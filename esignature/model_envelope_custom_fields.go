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

// checks if the EnvelopeCustomFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeCustomFields{}

// EnvelopeCustomFields An envelope custom field enables you to collect custom data about envelopes on a per-envelope basis. You can then use the custom data for sorting, organizing, searching, and other downstream processes. For example, you can use custom fields to copy envelopes or data to multiple areas in Salesforce. eOriginal customers can eVault their documents from the web app on a per-envelope basis by setting an envelope custom field with a name like \"eVault with eOriginal?\" to \"Yes\" or \"No\".  When a user creates an envelope, the envelope custom fields display in the **Envelope Settings** section of the DocuSign console. Envelope recipients do not see the envelope custom fields. For more information, see [Envelope Custom Fields](https://support.docusign.com/s/document-item?bundleId=pik1583277475390&topicId=qor1583277385137.html).
type EnvelopeCustomFields struct {
	// An array of list custom fields.
	ListCustomFields []ListCustomField `json:"listCustomFields,omitempty"`
	// An array of text custom fields.
	TextCustomFields []TextCustomField `json:"textCustomFields,omitempty"`
}

// NewEnvelopeCustomFields instantiates a new EnvelopeCustomFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeCustomFields() *EnvelopeCustomFields {
	this := EnvelopeCustomFields{}
	return &this
}

// NewEnvelopeCustomFieldsWithDefaults instantiates a new EnvelopeCustomFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeCustomFieldsWithDefaults() *EnvelopeCustomFields {
	this := EnvelopeCustomFields{}
	return &this
}

// GetListCustomFields returns the ListCustomFields field value if set, zero value otherwise.
func (o *EnvelopeCustomFields) GetListCustomFields() []ListCustomField {
	if o == nil || IsNil(o.ListCustomFields) {
		var ret []ListCustomField
		return ret
	}
	return o.ListCustomFields
}

// GetListCustomFieldsOk returns a tuple with the ListCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeCustomFields) GetListCustomFieldsOk() ([]ListCustomField, bool) {
	if o == nil || IsNil(o.ListCustomFields) {
		return nil, false
	}
	return o.ListCustomFields, true
}

// HasListCustomFields returns a boolean if a field has been set.
func (o *EnvelopeCustomFields) HasListCustomFields() bool {
	if o != nil && !IsNil(o.ListCustomFields) {
		return true
	}

	return false
}

// SetListCustomFields gets a reference to the given []ListCustomField and assigns it to the ListCustomFields field.
func (o *EnvelopeCustomFields) SetListCustomFields(v []ListCustomField) {
	o.ListCustomFields = v
}

// GetTextCustomFields returns the TextCustomFields field value if set, zero value otherwise.
func (o *EnvelopeCustomFields) GetTextCustomFields() []TextCustomField {
	if o == nil || IsNil(o.TextCustomFields) {
		var ret []TextCustomField
		return ret
	}
	return o.TextCustomFields
}

// GetTextCustomFieldsOk returns a tuple with the TextCustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeCustomFields) GetTextCustomFieldsOk() ([]TextCustomField, bool) {
	if o == nil || IsNil(o.TextCustomFields) {
		return nil, false
	}
	return o.TextCustomFields, true
}

// HasTextCustomFields returns a boolean if a field has been set.
func (o *EnvelopeCustomFields) HasTextCustomFields() bool {
	if o != nil && !IsNil(o.TextCustomFields) {
		return true
	}

	return false
}

// SetTextCustomFields gets a reference to the given []TextCustomField and assigns it to the TextCustomFields field.
func (o *EnvelopeCustomFields) SetTextCustomFields(v []TextCustomField) {
	o.TextCustomFields = v
}

func (o EnvelopeCustomFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeCustomFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ListCustomFields) {
		toSerialize["listCustomFields"] = o.ListCustomFields
	}
	if !IsNil(o.TextCustomFields) {
		toSerialize["textCustomFields"] = o.TextCustomFields
	}
	return toSerialize, nil
}

type NullableEnvelopeCustomFields struct {
	value *EnvelopeCustomFields
	isSet bool
}

func (v NullableEnvelopeCustomFields) Get() *EnvelopeCustomFields {
	return v.value
}

func (v *NullableEnvelopeCustomFields) Set(val *EnvelopeCustomFields) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeCustomFields) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeCustomFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeCustomFields(val *EnvelopeCustomFields) *NullableEnvelopeCustomFields {
	return &NullableEnvelopeCustomFields{value: val, isSet: true}
}

func (v NullableEnvelopeCustomFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeCustomFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

