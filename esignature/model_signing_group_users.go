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

// checks if the SigningGroupUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SigningGroupUsers{}

// SigningGroupUsers Signing groups' users
type SigningGroupUsers struct {
	// User management information.
	Users []SigningGroupUser `json:"users,omitempty"`
}

// NewSigningGroupUsers instantiates a new SigningGroupUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningGroupUsers() *SigningGroupUsers {
	this := SigningGroupUsers{}
	return &this
}

// NewSigningGroupUsersWithDefaults instantiates a new SigningGroupUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningGroupUsersWithDefaults() *SigningGroupUsers {
	this := SigningGroupUsers{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *SigningGroupUsers) GetUsers() []SigningGroupUser {
	if o == nil || IsNil(o.Users) {
		var ret []SigningGroupUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningGroupUsers) GetUsersOk() ([]SigningGroupUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *SigningGroupUsers) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []SigningGroupUser and assigns it to the Users field.
func (o *SigningGroupUsers) SetUsers(v []SigningGroupUser) {
	o.Users = v
}

func (o SigningGroupUsers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SigningGroupUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableSigningGroupUsers struct {
	value *SigningGroupUsers
	isSet bool
}

func (v NullableSigningGroupUsers) Get() *SigningGroupUsers {
	return v.value
}

func (v *NullableSigningGroupUsers) Set(val *SigningGroupUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableSigningGroupUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableSigningGroupUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSigningGroupUsers(val *SigningGroupUsers) *NullableSigningGroupUsers {
	return &NullableSigningGroupUsers{value: val, isSet: true}
}

func (v NullableSigningGroupUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSigningGroupUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


