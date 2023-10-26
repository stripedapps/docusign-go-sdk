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

// checks if the PermissionProfileInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionProfileInformation{}

// PermissionProfileInformation Contains details about the permission profiles associated with an account.
type PermissionProfileInformation struct {
	// A complex type containing a collection of permission profiles.
	PermissionProfiles []PermissionProfile `json:"permissionProfiles,omitempty"`
}

// NewPermissionProfileInformation instantiates a new PermissionProfileInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionProfileInformation() *PermissionProfileInformation {
	this := PermissionProfileInformation{}
	return &this
}

// NewPermissionProfileInformationWithDefaults instantiates a new PermissionProfileInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionProfileInformationWithDefaults() *PermissionProfileInformation {
	this := PermissionProfileInformation{}
	return &this
}

// GetPermissionProfiles returns the PermissionProfiles field value if set, zero value otherwise.
func (o *PermissionProfileInformation) GetPermissionProfiles() []PermissionProfile {
	if o == nil || IsNil(o.PermissionProfiles) {
		var ret []PermissionProfile
		return ret
	}
	return o.PermissionProfiles
}

// GetPermissionProfilesOk returns a tuple with the PermissionProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PermissionProfileInformation) GetPermissionProfilesOk() ([]PermissionProfile, bool) {
	if o == nil || IsNil(o.PermissionProfiles) {
		return nil, false
	}
	return o.PermissionProfiles, true
}

// HasPermissionProfiles returns a boolean if a field has been set.
func (o *PermissionProfileInformation) HasPermissionProfiles() bool {
	if o != nil && !IsNil(o.PermissionProfiles) {
		return true
	}

	return false
}

// SetPermissionProfiles gets a reference to the given []PermissionProfile and assigns it to the PermissionProfiles field.
func (o *PermissionProfileInformation) SetPermissionProfiles(v []PermissionProfile) {
	o.PermissionProfiles = v
}

func (o PermissionProfileInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionProfileInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PermissionProfiles) {
		toSerialize["permissionProfiles"] = o.PermissionProfiles
	}
	return toSerialize, nil
}

type NullablePermissionProfileInformation struct {
	value *PermissionProfileInformation
	isSet bool
}

func (v NullablePermissionProfileInformation) Get() *PermissionProfileInformation {
	return v.value
}

func (v *NullablePermissionProfileInformation) Set(val *PermissionProfileInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionProfileInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionProfileInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionProfileInformation(val *PermissionProfileInformation) *NullablePermissionProfileInformation {
	return &NullablePermissionProfileInformation{value: val, isSet: true}
}

func (v NullablePermissionProfileInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionProfileInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


