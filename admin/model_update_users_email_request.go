/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateUsersEmailRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUsersEmailRequest{}

// UpdateUsersEmailRequest A change email request.
type UpdateUsersEmailRequest struct {
	// A list of users whose email address to change.
	Users []UpdateUserEmailRequest `json:"users,omitempty"`
}

// NewUpdateUsersEmailRequest instantiates a new UpdateUsersEmailRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUsersEmailRequest() *UpdateUsersEmailRequest {
	this := UpdateUsersEmailRequest{}
	return &this
}

// NewUpdateUsersEmailRequestWithDefaults instantiates a new UpdateUsersEmailRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUsersEmailRequestWithDefaults() *UpdateUsersEmailRequest {
	this := UpdateUsersEmailRequest{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UpdateUsersEmailRequest) GetUsers() []UpdateUserEmailRequest {
	if o == nil || IsNil(o.Users) {
		var ret []UpdateUserEmailRequest
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUsersEmailRequest) GetUsersOk() ([]UpdateUserEmailRequest, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UpdateUsersEmailRequest) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UpdateUserEmailRequest and assigns it to the Users field.
func (o *UpdateUsersEmailRequest) SetUsers(v []UpdateUserEmailRequest) {
	o.Users = v
}

func (o UpdateUsersEmailRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUsersEmailRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUpdateUsersEmailRequest struct {
	value *UpdateUsersEmailRequest
	isSet bool
}

func (v NullableUpdateUsersEmailRequest) Get() *UpdateUsersEmailRequest {
	return v.value
}

func (v *NullableUpdateUsersEmailRequest) Set(val *UpdateUsersEmailRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUsersEmailRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUsersEmailRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUsersEmailRequest(val *UpdateUsersEmailRequest) *NullableUpdateUsersEmailRequest {
	return &NullableUpdateUsersEmailRequest{value: val, isSet: true}
}

func (v NullableUpdateUsersEmailRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUsersEmailRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


