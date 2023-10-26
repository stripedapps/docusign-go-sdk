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

// checks if the AddDSGroupAndUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDSGroupAndUsersResponse{}

// AddDSGroupAndUsersResponse 
type AddDSGroupAndUsersResponse struct {
	Group *DSGroupResponse `json:"group,omitempty"`
	GroupUsers *AddDSGroupUsersResponse `json:"group_users,omitempty"`
}

// NewAddDSGroupAndUsersResponse instantiates a new AddDSGroupAndUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDSGroupAndUsersResponse() *AddDSGroupAndUsersResponse {
	this := AddDSGroupAndUsersResponse{}
	return &this
}

// NewAddDSGroupAndUsersResponseWithDefaults instantiates a new AddDSGroupAndUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDSGroupAndUsersResponseWithDefaults() *AddDSGroupAndUsersResponse {
	this := AddDSGroupAndUsersResponse{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *AddDSGroupAndUsersResponse) GetGroup() DSGroupResponse {
	if o == nil || IsNil(o.Group) {
		var ret DSGroupResponse
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDSGroupAndUsersResponse) GetGroupOk() (*DSGroupResponse, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *AddDSGroupAndUsersResponse) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given DSGroupResponse and assigns it to the Group field.
func (o *AddDSGroupAndUsersResponse) SetGroup(v DSGroupResponse) {
	o.Group = &v
}

// GetGroupUsers returns the GroupUsers field value if set, zero value otherwise.
func (o *AddDSGroupAndUsersResponse) GetGroupUsers() AddDSGroupUsersResponse {
	if o == nil || IsNil(o.GroupUsers) {
		var ret AddDSGroupUsersResponse
		return ret
	}
	return *o.GroupUsers
}

// GetGroupUsersOk returns a tuple with the GroupUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDSGroupAndUsersResponse) GetGroupUsersOk() (*AddDSGroupUsersResponse, bool) {
	if o == nil || IsNil(o.GroupUsers) {
		return nil, false
	}
	return o.GroupUsers, true
}

// HasGroupUsers returns a boolean if a field has been set.
func (o *AddDSGroupAndUsersResponse) HasGroupUsers() bool {
	if o != nil && !IsNil(o.GroupUsers) {
		return true
	}

	return false
}

// SetGroupUsers gets a reference to the given AddDSGroupUsersResponse and assigns it to the GroupUsers field.
func (o *AddDSGroupAndUsersResponse) SetGroupUsers(v AddDSGroupUsersResponse) {
	o.GroupUsers = &v
}

func (o AddDSGroupAndUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDSGroupAndUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.GroupUsers) {
		toSerialize["group_users"] = o.GroupUsers
	}
	return toSerialize, nil
}

type NullableAddDSGroupAndUsersResponse struct {
	value *AddDSGroupAndUsersResponse
	isSet bool
}

func (v NullableAddDSGroupAndUsersResponse) Get() *AddDSGroupAndUsersResponse {
	return v.value
}

func (v *NullableAddDSGroupAndUsersResponse) Set(val *AddDSGroupAndUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDSGroupAndUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDSGroupAndUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDSGroupAndUsersResponse(val *AddDSGroupAndUsersResponse) *NullableAddDSGroupAndUsersResponse {
	return &NullableAddDSGroupAndUsersResponse{value: val, isSet: true}
}

func (v NullableAddDSGroupAndUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDSGroupAndUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

