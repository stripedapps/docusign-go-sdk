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

// checks if the EnvelopeDocumentFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeDocumentFields{}

// EnvelopeDocumentFields Envelope document fields
type EnvelopeDocumentFields struct {
	// The array of name/value custom data strings to be added to a document. Custom document field information is returned in the status, but otherwise is not used by DocuSign. The array contains the elements:   * name - A string that can be a maximum of 50 characters.  * value - A string that can be a maximum of 200 characters.  *IMPORTANT*: If you are using xml, the name/value pair is contained in a nameValue element.  
	DocumentFields []NameValue `json:"documentFields,omitempty"`
}

// NewEnvelopeDocumentFields instantiates a new EnvelopeDocumentFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeDocumentFields() *EnvelopeDocumentFields {
	this := EnvelopeDocumentFields{}
	return &this
}

// NewEnvelopeDocumentFieldsWithDefaults instantiates a new EnvelopeDocumentFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeDocumentFieldsWithDefaults() *EnvelopeDocumentFields {
	this := EnvelopeDocumentFields{}
	return &this
}

// GetDocumentFields returns the DocumentFields field value if set, zero value otherwise.
func (o *EnvelopeDocumentFields) GetDocumentFields() []NameValue {
	if o == nil || IsNil(o.DocumentFields) {
		var ret []NameValue
		return ret
	}
	return o.DocumentFields
}

// GetDocumentFieldsOk returns a tuple with the DocumentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeDocumentFields) GetDocumentFieldsOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.DocumentFields) {
		return nil, false
	}
	return o.DocumentFields, true
}

// HasDocumentFields returns a boolean if a field has been set.
func (o *EnvelopeDocumentFields) HasDocumentFields() bool {
	if o != nil && !IsNil(o.DocumentFields) {
		return true
	}

	return false
}

// SetDocumentFields gets a reference to the given []NameValue and assigns it to the DocumentFields field.
func (o *EnvelopeDocumentFields) SetDocumentFields(v []NameValue) {
	o.DocumentFields = v
}

func (o EnvelopeDocumentFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeDocumentFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentFields) {
		toSerialize["documentFields"] = o.DocumentFields
	}
	return toSerialize, nil
}

type NullableEnvelopeDocumentFields struct {
	value *EnvelopeDocumentFields
	isSet bool
}

func (v NullableEnvelopeDocumentFields) Get() *EnvelopeDocumentFields {
	return v.value
}

func (v *NullableEnvelopeDocumentFields) Set(val *EnvelopeDocumentFields) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeDocumentFields) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeDocumentFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeDocumentFields(val *EnvelopeDocumentFields) *NullableEnvelopeDocumentFields {
	return &NullableEnvelopeDocumentFields{value: val, isSet: true}
}

func (v NullableEnvelopeDocumentFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeDocumentFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


