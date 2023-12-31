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

// checks if the SocialAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SocialAuthentication{}

// SocialAuthentication 
type SocialAuthentication struct {
	// Reserved for DocuSign.
	Authentication *string `json:"authentication,omitempty"`
}

// NewSocialAuthentication instantiates a new SocialAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialAuthentication() *SocialAuthentication {
	this := SocialAuthentication{}
	return &this
}

// NewSocialAuthenticationWithDefaults instantiates a new SocialAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialAuthenticationWithDefaults() *SocialAuthentication {
	this := SocialAuthentication{}
	return &this
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *SocialAuthentication) GetAuthentication() string {
	if o == nil || IsNil(o.Authentication) {
		var ret string
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAuthentication) GetAuthenticationOk() (*string, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *SocialAuthentication) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given string and assigns it to the Authentication field.
func (o *SocialAuthentication) SetAuthentication(v string) {
	o.Authentication = &v
}

func (o SocialAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SocialAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	return toSerialize, nil
}

type NullableSocialAuthentication struct {
	value *SocialAuthentication
	isSet bool
}

func (v NullableSocialAuthentication) Get() *SocialAuthentication {
	return v.value
}

func (v *NullableSocialAuthentication) Set(val *SocialAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialAuthentication(val *SocialAuthentication) *NullableSocialAuthentication {
	return &NullableSocialAuthentication{value: val, isSet: true}
}

func (v NullableSocialAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


