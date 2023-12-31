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

// checks if the EnvelopeHtmlDefinitions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeHtmlDefinitions{}

// EnvelopeHtmlDefinitions 
type EnvelopeHtmlDefinitions struct {
	// Holds the properties that define how to generate the responsive-formatted HTML for the document.
	HtmlDefinitions []DocumentHtmlDefinitionOriginal `json:"htmlDefinitions,omitempty"`
}

// NewEnvelopeHtmlDefinitions instantiates a new EnvelopeHtmlDefinitions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeHtmlDefinitions() *EnvelopeHtmlDefinitions {
	this := EnvelopeHtmlDefinitions{}
	return &this
}

// NewEnvelopeHtmlDefinitionsWithDefaults instantiates a new EnvelopeHtmlDefinitions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeHtmlDefinitionsWithDefaults() *EnvelopeHtmlDefinitions {
	this := EnvelopeHtmlDefinitions{}
	return &this
}

// GetHtmlDefinitions returns the HtmlDefinitions field value if set, zero value otherwise.
func (o *EnvelopeHtmlDefinitions) GetHtmlDefinitions() []DocumentHtmlDefinitionOriginal {
	if o == nil || IsNil(o.HtmlDefinitions) {
		var ret []DocumentHtmlDefinitionOriginal
		return ret
	}
	return o.HtmlDefinitions
}

// GetHtmlDefinitionsOk returns a tuple with the HtmlDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeHtmlDefinitions) GetHtmlDefinitionsOk() ([]DocumentHtmlDefinitionOriginal, bool) {
	if o == nil || IsNil(o.HtmlDefinitions) {
		return nil, false
	}
	return o.HtmlDefinitions, true
}

// HasHtmlDefinitions returns a boolean if a field has been set.
func (o *EnvelopeHtmlDefinitions) HasHtmlDefinitions() bool {
	if o != nil && !IsNil(o.HtmlDefinitions) {
		return true
	}

	return false
}

// SetHtmlDefinitions gets a reference to the given []DocumentHtmlDefinitionOriginal and assigns it to the HtmlDefinitions field.
func (o *EnvelopeHtmlDefinitions) SetHtmlDefinitions(v []DocumentHtmlDefinitionOriginal) {
	o.HtmlDefinitions = v
}

func (o EnvelopeHtmlDefinitions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeHtmlDefinitions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HtmlDefinitions) {
		toSerialize["htmlDefinitions"] = o.HtmlDefinitions
	}
	return toSerialize, nil
}

type NullableEnvelopeHtmlDefinitions struct {
	value *EnvelopeHtmlDefinitions
	isSet bool
}

func (v NullableEnvelopeHtmlDefinitions) Get() *EnvelopeHtmlDefinitions {
	return v.value
}

func (v *NullableEnvelopeHtmlDefinitions) Set(val *EnvelopeHtmlDefinitions) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeHtmlDefinitions) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeHtmlDefinitions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeHtmlDefinitions(val *EnvelopeHtmlDefinitions) *NullableEnvelopeHtmlDefinitions {
	return &NullableEnvelopeHtmlDefinitions{value: val, isSet: true}
}

func (v NullableEnvelopeHtmlDefinitions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeHtmlDefinitions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


