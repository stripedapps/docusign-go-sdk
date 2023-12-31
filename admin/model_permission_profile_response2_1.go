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

// checks if the PermissionProfileResponse21 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionProfileResponse21{}

// PermissionProfileResponse21 
type PermissionProfileResponse21 struct {
	// The ID of the permission profile.
	PermissionProfileId *string `json:"permission_profile_id,omitempty"`
	// The human-readable name of the permission profile.
	PermissionProfileName *string `json:"permission_profile_name,omitempty"`
}

// NewPermissionProfileResponse21 instantiates a new PermissionProfileResponse21 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionProfileResponse21() *PermissionProfileResponse21 {
	this := PermissionProfileResponse21{}
	return &this
}

// NewPermissionProfileResponse21WithDefaults instantiates a new PermissionProfileResponse21 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionProfileResponse21WithDefaults() *PermissionProfileResponse21 {
	this := PermissionProfileResponse21{}
	return &this
}

// GetPermissionProfileId returns the PermissionProfileId field value if set, zero value otherwise.
func (o *PermissionProfileResponse21) GetPermissionProfileId() string {
	if o == nil || IsNil(o.PermissionProfileId) {
		var ret string
		return ret
	}
	return *o.PermissionProfileId
}

// GetPermissionProfileIdOk returns a tuple with the PermissionProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionProfileResponse21) GetPermissionProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionProfileId) {
		return nil, false
	}
	return o.PermissionProfileId, true
}

// HasPermissionProfileId returns a boolean if a field has been set.
func (o *PermissionProfileResponse21) HasPermissionProfileId() bool {
	if o != nil && !IsNil(o.PermissionProfileId) {
		return true
	}

	return false
}

// SetPermissionProfileId gets a reference to the given string and assigns it to the PermissionProfileId field.
func (o *PermissionProfileResponse21) SetPermissionProfileId(v string) {
	o.PermissionProfileId = &v
}

// GetPermissionProfileName returns the PermissionProfileName field value if set, zero value otherwise.
func (o *PermissionProfileResponse21) GetPermissionProfileName() string {
	if o == nil || IsNil(o.PermissionProfileName) {
		var ret string
		return ret
	}
	return *o.PermissionProfileName
}

// GetPermissionProfileNameOk returns a tuple with the PermissionProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionProfileResponse21) GetPermissionProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionProfileName) {
		return nil, false
	}
	return o.PermissionProfileName, true
}

// HasPermissionProfileName returns a boolean if a field has been set.
func (o *PermissionProfileResponse21) HasPermissionProfileName() bool {
	if o != nil && !IsNil(o.PermissionProfileName) {
		return true
	}

	return false
}

// SetPermissionProfileName gets a reference to the given string and assigns it to the PermissionProfileName field.
func (o *PermissionProfileResponse21) SetPermissionProfileName(v string) {
	o.PermissionProfileName = &v
}

func (o PermissionProfileResponse21) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionProfileResponse21) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionProfileId) {
		toSerialize["permission_profile_id"] = o.PermissionProfileId
	}
	if !IsNil(o.PermissionProfileName) {
		toSerialize["permission_profile_name"] = o.PermissionProfileName
	}
	return toSerialize, nil
}

type NullablePermissionProfileResponse21 struct {
	value *PermissionProfileResponse21
	isSet bool
}

func (v NullablePermissionProfileResponse21) Get() *PermissionProfileResponse21 {
	return v.value
}

func (v *NullablePermissionProfileResponse21) Set(val *PermissionProfileResponse21) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionProfileResponse21) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionProfileResponse21) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionProfileResponse21(val *PermissionProfileResponse21) *NullablePermissionProfileResponse21 {
	return &NullablePermissionProfileResponse21{value: val, isSet: true}
}

func (v NullablePermissionProfileResponse21) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionProfileResponse21) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


