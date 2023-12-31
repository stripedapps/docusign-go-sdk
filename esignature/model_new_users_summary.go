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

// checks if the NewUsersSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewUsersSummary{}

// NewUsersSummary Object representing a summary of data for new users.
type NewUsersSummary struct {
	// A list of one or more new users.
	NewUsers []NewUser `json:"newUsers,omitempty"`
}

// NewNewUsersSummary instantiates a new NewUsersSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUsersSummary() *NewUsersSummary {
	this := NewUsersSummary{}
	return &this
}

// NewNewUsersSummaryWithDefaults instantiates a new NewUsersSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUsersSummaryWithDefaults() *NewUsersSummary {
	this := NewUsersSummary{}
	return &this
}

// GetNewUsers returns the NewUsers field value if set, zero value otherwise.
func (o *NewUsersSummary) GetNewUsers() []NewUser {
	if o == nil || IsNil(o.NewUsers) {
		var ret []NewUser
		return ret
	}
	return o.NewUsers
}

// GetNewUsersOk returns a tuple with the NewUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUsersSummary) GetNewUsersOk() ([]NewUser, bool) {
	if o == nil || IsNil(o.NewUsers) {
		return nil, false
	}
	return o.NewUsers, true
}

// HasNewUsers returns a boolean if a field has been set.
func (o *NewUsersSummary) HasNewUsers() bool {
	if o != nil && !IsNil(o.NewUsers) {
		return true
	}

	return false
}

// SetNewUsers gets a reference to the given []NewUser and assigns it to the NewUsers field.
func (o *NewUsersSummary) SetNewUsers(v []NewUser) {
	o.NewUsers = v
}

func (o NewUsersSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewUsersSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewUsers) {
		toSerialize["newUsers"] = o.NewUsers
	}
	return toSerialize, nil
}

type NullableNewUsersSummary struct {
	value *NewUsersSummary
	isSet bool
}

func (v NullableNewUsersSummary) Get() *NewUsersSummary {
	return v.value
}

func (v *NullableNewUsersSummary) Set(val *NewUsersSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUsersSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUsersSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUsersSummary(val *NewUsersSummary) *NullableNewUsersSummary {
	return &NullableNewUsersSummary{value: val, isSet: true}
}

func (v NullableNewUsersSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUsersSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


