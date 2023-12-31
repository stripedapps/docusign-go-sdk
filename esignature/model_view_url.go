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

// checks if the ViewUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ViewUrl{}

// ViewUrl 
type ViewUrl struct {
	// URL to the Review ID page.
	Url *string `json:"url,omitempty"`
}

// NewViewUrl instantiates a new ViewUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewUrl() *ViewUrl {
	this := ViewUrl{}
	return &this
}

// NewViewUrlWithDefaults instantiates a new ViewUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewUrlWithDefaults() *ViewUrl {
	this := ViewUrl{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ViewUrl) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewUrl) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ViewUrl) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ViewUrl) SetUrl(v string) {
	o.Url = &v
}

func (o ViewUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ViewUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableViewUrl struct {
	value *ViewUrl
	isSet bool
}

func (v NullableViewUrl) Get() *ViewUrl {
	return v.value
}

func (v *NullableViewUrl) Set(val *ViewUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableViewUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableViewUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewUrl(val *ViewUrl) *NullableViewUrl {
	return &NullableViewUrl{value: val, isSet: true}
}

func (v NullableViewUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


