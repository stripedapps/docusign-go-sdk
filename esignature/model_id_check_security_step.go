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

// checks if the IdCheckSecurityStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdCheckSecurityStep{}

// IdCheckSecurityStep 
type IdCheckSecurityStep struct {
	// Type of authorization used for the security check.
	AuthType *string `json:"authType,omitempty"`
}

// NewIdCheckSecurityStep instantiates a new IdCheckSecurityStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdCheckSecurityStep() *IdCheckSecurityStep {
	this := IdCheckSecurityStep{}
	return &this
}

// NewIdCheckSecurityStepWithDefaults instantiates a new IdCheckSecurityStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdCheckSecurityStepWithDefaults() *IdCheckSecurityStep {
	this := IdCheckSecurityStep{}
	return &this
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *IdCheckSecurityStep) GetAuthType() string {
	if o == nil || IsNil(o.AuthType) {
		var ret string
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdCheckSecurityStep) GetAuthTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *IdCheckSecurityStep) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given string and assigns it to the AuthType field.
func (o *IdCheckSecurityStep) SetAuthType(v string) {
	o.AuthType = &v
}

func (o IdCheckSecurityStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdCheckSecurityStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthType) {
		toSerialize["authType"] = o.AuthType
	}
	return toSerialize, nil
}

type NullableIdCheckSecurityStep struct {
	value *IdCheckSecurityStep
	isSet bool
}

func (v NullableIdCheckSecurityStep) Get() *IdCheckSecurityStep {
	return v.value
}

func (v *NullableIdCheckSecurityStep) Set(val *IdCheckSecurityStep) {
	v.value = val
	v.isSet = true
}

func (v NullableIdCheckSecurityStep) IsSet() bool {
	return v.isSet
}

func (v *NullableIdCheckSecurityStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdCheckSecurityStep(val *IdCheckSecurityStep) *NullableIdCheckSecurityStep {
	return &NullableIdCheckSecurityStep{value: val, isSet: true}
}

func (v NullableIdCheckSecurityStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdCheckSecurityStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

