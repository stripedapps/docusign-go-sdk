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

// checks if the GroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupRequest{}

// GroupRequest A group for a user to belong to.
type GroupRequest struct {
	// The ID of the group.
	Id int64 `json:"id"`
	// The name of the group.
	Name *string `json:"name,omitempty"`
	// The type of group. One of:  - `invalid` - `admin_group` - `everyone_group` - `custom_group`
	Type *string `json:"type,omitempty"`
}

// NewGroupRequest instantiates a new GroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupRequest(id int64) *GroupRequest {
	this := GroupRequest{}
	this.Id = id
	return &this
}

// NewGroupRequestWithDefaults instantiates a new GroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupRequestWithDefaults() *GroupRequest {
	this := GroupRequest{}
	return &this
}

// GetId returns the Id field value
func (o *GroupRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupRequest) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupRequest) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GroupRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GroupRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GroupRequest) SetType(v string) {
	o.Type = &v
}

func (o GroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableGroupRequest struct {
	value *GroupRequest
	isSet bool
}

func (v NullableGroupRequest) Get() *GroupRequest {
	return v.value
}

func (v *NullableGroupRequest) Set(val *GroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupRequest(val *GroupRequest) *NullableGroupRequest {
	return &NullableGroupRequest{value: val, isSet: true}
}

func (v NullableGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

