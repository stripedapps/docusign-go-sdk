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

// checks if the EnvelopeViews type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeViews{}

// EnvelopeViews Provides a URL that you can embed in your application to provide access to the DocuSign UI.  ### Related topics  - [Embedded signing and sending](/docs/esign-rest-api/esign101/concepts/embedding/) - [Send an envelope via your app](/docs/esign-rest-api/how-to/embedded-sending/) - [Introducing customizable embedded sending](https://www.docusign.com/blog/developers/introducing-customizable-embedded-sending)  
type EnvelopeViews struct {
	// The view URL to be navigated to.
	Url *string `json:"url,omitempty"`
}

// NewEnvelopeViews instantiates a new EnvelopeViews object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeViews() *EnvelopeViews {
	this := EnvelopeViews{}
	return &this
}

// NewEnvelopeViewsWithDefaults instantiates a new EnvelopeViews object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeViewsWithDefaults() *EnvelopeViews {
	this := EnvelopeViews{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *EnvelopeViews) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeViews) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *EnvelopeViews) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *EnvelopeViews) SetUrl(v string) {
	o.Url = &v
}

func (o EnvelopeViews) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeViews) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableEnvelopeViews struct {
	value *EnvelopeViews
	isSet bool
}

func (v NullableEnvelopeViews) Get() *EnvelopeViews {
	return v.value
}

func (v *NullableEnvelopeViews) Set(val *EnvelopeViews) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeViews) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeViews) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeViews(val *EnvelopeViews) *NullableEnvelopeViews {
	return &NullableEnvelopeViews{value: val, isSet: true}
}

func (v NullableEnvelopeViews) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeViews) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


