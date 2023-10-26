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

// checks if the OrganizationUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationUsersResponse{}

// OrganizationUsersResponse A response containing information about users.
type OrganizationUsersResponse struct {
	// A list of users.
	Users []OrganizationUserResponse `json:"users,omitempty"`
	Paging *PagingResponseProperties `json:"paging,omitempty"`
}

// NewOrganizationUsersResponse instantiates a new OrganizationUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationUsersResponse() *OrganizationUsersResponse {
	this := OrganizationUsersResponse{}
	return &this
}

// NewOrganizationUsersResponseWithDefaults instantiates a new OrganizationUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationUsersResponseWithDefaults() *OrganizationUsersResponse {
	this := OrganizationUsersResponse{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *OrganizationUsersResponse) GetUsers() []OrganizationUserResponse {
	if o == nil || IsNil(o.Users) {
		var ret []OrganizationUserResponse
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUsersResponse) GetUsersOk() ([]OrganizationUserResponse, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *OrganizationUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []OrganizationUserResponse and assigns it to the Users field.
func (o *OrganizationUsersResponse) SetUsers(v []OrganizationUserResponse) {
	o.Users = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *OrganizationUsersResponse) GetPaging() PagingResponseProperties {
	if o == nil || IsNil(o.Paging) {
		var ret PagingResponseProperties
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationUsersResponse) GetPagingOk() (*PagingResponseProperties, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *OrganizationUsersResponse) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given PagingResponseProperties and assigns it to the Paging field.
func (o *OrganizationUsersResponse) SetPaging(v PagingResponseProperties) {
	o.Paging = &v
}

func (o OrganizationUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableOrganizationUsersResponse struct {
	value *OrganizationUsersResponse
	isSet bool
}

func (v NullableOrganizationUsersResponse) Get() *OrganizationUsersResponse {
	return v.value
}

func (v *NullableOrganizationUsersResponse) Set(val *OrganizationUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationUsersResponse(val *OrganizationUsersResponse) *NullableOrganizationUsersResponse {
	return &NullableOrganizationUsersResponse{value: val, isSet: true}
}

func (v NullableOrganizationUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

