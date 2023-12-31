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

// checks if the UserSignaturesInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSignaturesInformation{}

// UserSignaturesInformation 
type UserSignaturesInformation struct {
	// An array of  `userSignature` objects.
	UserSignatures []UserSignature `json:"userSignatures,omitempty"`
}

// NewUserSignaturesInformation instantiates a new UserSignaturesInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSignaturesInformation() *UserSignaturesInformation {
	this := UserSignaturesInformation{}
	return &this
}

// NewUserSignaturesInformationWithDefaults instantiates a new UserSignaturesInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSignaturesInformationWithDefaults() *UserSignaturesInformation {
	this := UserSignaturesInformation{}
	return &this
}

// GetUserSignatures returns the UserSignatures field value if set, zero value otherwise.
func (o *UserSignaturesInformation) GetUserSignatures() []UserSignature {
	if o == nil || IsNil(o.UserSignatures) {
		var ret []UserSignature
		return ret
	}
	return o.UserSignatures
}

// GetUserSignaturesOk returns a tuple with the UserSignatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignaturesInformation) GetUserSignaturesOk() ([]UserSignature, bool) {
	if o == nil || IsNil(o.UserSignatures) {
		return nil, false
	}
	return o.UserSignatures, true
}

// HasUserSignatures returns a boolean if a field has been set.
func (o *UserSignaturesInformation) HasUserSignatures() bool {
	if o != nil && !IsNil(o.UserSignatures) {
		return true
	}

	return false
}

// SetUserSignatures gets a reference to the given []UserSignature and assigns it to the UserSignatures field.
func (o *UserSignaturesInformation) SetUserSignatures(v []UserSignature) {
	o.UserSignatures = v
}

func (o UserSignaturesInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSignaturesInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserSignatures) {
		toSerialize["userSignatures"] = o.UserSignatures
	}
	return toSerialize, nil
}

type NullableUserSignaturesInformation struct {
	value *UserSignaturesInformation
	isSet bool
}

func (v NullableUserSignaturesInformation) Get() *UserSignaturesInformation {
	return v.value
}

func (v *NullableUserSignaturesInformation) Set(val *UserSignaturesInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSignaturesInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSignaturesInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSignaturesInformation(val *UserSignaturesInformation) *NullableUserSignaturesInformation {
	return &NullableUserSignaturesInformation{value: val, isSet: true}
}

func (v NullableUserSignaturesInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSignaturesInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


