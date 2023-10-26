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

// checks if the PermissionProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionProfileRequest{}

// PermissionProfileRequest A permission profile.
type PermissionProfileRequest struct {
	// The ID of the permission profile.
	Id int64 `json:"id"`
	// The name of the permission profile.
	Name *string `json:"name,omitempty"`
}

// NewPermissionProfileRequest instantiates a new PermissionProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionProfileRequest(id int64) *PermissionProfileRequest {
	this := PermissionProfileRequest{}
	this.Id = id
	return &this
}

// NewPermissionProfileRequestWithDefaults instantiates a new PermissionProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionProfileRequestWithDefaults() *PermissionProfileRequest {
	this := PermissionProfileRequest{}
	return &this
}

// GetId returns the Id field value
func (o *PermissionProfileRequest) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionProfileRequest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionProfileRequest) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PermissionProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PermissionProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PermissionProfileRequest) SetName(v string) {
	o.Name = &v
}

func (o PermissionProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullablePermissionProfileRequest struct {
	value *PermissionProfileRequest
	isSet bool
}

func (v NullablePermissionProfileRequest) Get() *PermissionProfileRequest {
	return v.value
}

func (v *NullablePermissionProfileRequest) Set(val *PermissionProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionProfileRequest(val *PermissionProfileRequest) *NullablePermissionProfileRequest {
	return &NullablePermissionProfileRequest{value: val, isSet: true}
}

func (v NullablePermissionProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

