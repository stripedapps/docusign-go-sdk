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

// checks if the TemplateViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateViews{}

// TemplateViews A TemplateView contains a URL that you can embed in your application to generate a template view that uses the DocuSign user interface (UI).
type TemplateViews struct {
	// The URL that you navigate to in order to start the view.
	Url *string `json:"url,omitempty"`
}

// NewTemplateViews instantiates a new TemplateViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateViews() *TemplateViews {
	this := TemplateViews{}
	return &this
}

// NewTemplateViewsWithDefaults instantiates a new TemplateViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateViewsWithDefaults() *TemplateViews {
	this := TemplateViews{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *TemplateViews) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateViews) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *TemplateViews) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *TemplateViews) SetUrl(v string) {
	o.Url = &v
}

func (o TemplateViews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableTemplateViews struct {
	value *TemplateViews
	isSet bool
}

func (v NullableTemplateViews) Get() *TemplateViews {
	return v.value
}

func (v *NullableTemplateViews) Set(val *TemplateViews) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateViews) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateViews(val *TemplateViews) *NullableTemplateViews {
	return &NullableTemplateViews{value: val, isSet: true}
}

func (v NullableTemplateViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

