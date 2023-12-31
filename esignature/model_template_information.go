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

// checks if the TemplateInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateInformation{}

// TemplateInformation 
type TemplateInformation struct {
	// An array of `templateSummary` objects that contain information about templates.
	Templates []TemplateSummary `json:"templates,omitempty"`
}

// NewTemplateInformation instantiates a new TemplateInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateInformation() *TemplateInformation {
	this := TemplateInformation{}
	return &this
}

// NewTemplateInformationWithDefaults instantiates a new TemplateInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateInformationWithDefaults() *TemplateInformation {
	this := TemplateInformation{}
	return &this
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *TemplateInformation) GetTemplates() []TemplateSummary {
	if o == nil || IsNil(o.Templates) {
		var ret []TemplateSummary
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateInformation) GetTemplatesOk() ([]TemplateSummary, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *TemplateInformation) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []TemplateSummary and assigns it to the Templates field.
func (o *TemplateInformation) SetTemplates(v []TemplateSummary) {
	o.Templates = v
}

func (o TemplateInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	return toSerialize, nil
}

type NullableTemplateInformation struct {
	value *TemplateInformation
	isSet bool
}

func (v NullableTemplateInformation) Get() *TemplateInformation {
	return v.value
}

func (v *NullableTemplateInformation) Set(val *TemplateInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateInformation(val *TemplateInformation) *NullableTemplateInformation {
	return &NullableTemplateInformation{value: val, isSet: true}
}

func (v NullableTemplateInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


