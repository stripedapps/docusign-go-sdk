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

// checks if the OrganizationImportResponseWarningRollup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationImportResponseWarningRollup{}

// OrganizationImportResponseWarningRollup 
type OrganizationImportResponseWarningRollup struct {
	// The type of warning.
	WarningType *string `json:"warning_type,omitempty"`
	// The number of warnings of this type.
	Count *int32 `json:"count,omitempty"`
}

// NewOrganizationImportResponseWarningRollup instantiates a new OrganizationImportResponseWarningRollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationImportResponseWarningRollup() *OrganizationImportResponseWarningRollup {
	this := OrganizationImportResponseWarningRollup{}
	return &this
}

// NewOrganizationImportResponseWarningRollupWithDefaults instantiates a new OrganizationImportResponseWarningRollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationImportResponseWarningRollupWithDefaults() *OrganizationImportResponseWarningRollup {
	this := OrganizationImportResponseWarningRollup{}
	return &this
}

// GetWarningType returns the WarningType field value if set, zero value otherwise.
func (o *OrganizationImportResponseWarningRollup) GetWarningType() string {
	if o == nil || IsNil(o.WarningType) {
		var ret string
		return ret
	}
	return *o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponseWarningRollup) GetWarningTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WarningType) {
		return nil, false
	}
	return o.WarningType, true
}

// HasWarningType returns a boolean if a field has been set.
func (o *OrganizationImportResponseWarningRollup) HasWarningType() bool {
	if o != nil && !IsNil(o.WarningType) {
		return true
	}

	return false
}

// SetWarningType gets a reference to the given string and assigns it to the WarningType field.
func (o *OrganizationImportResponseWarningRollup) SetWarningType(v string) {
	o.WarningType = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *OrganizationImportResponseWarningRollup) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponseWarningRollup) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *OrganizationImportResponseWarningRollup) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *OrganizationImportResponseWarningRollup) SetCount(v int32) {
	o.Count = &v
}

func (o OrganizationImportResponseWarningRollup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationImportResponseWarningRollup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WarningType) {
		toSerialize["warning_type"] = o.WarningType
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableOrganizationImportResponseWarningRollup struct {
	value *OrganizationImportResponseWarningRollup
	isSet bool
}

func (v NullableOrganizationImportResponseWarningRollup) Get() *OrganizationImportResponseWarningRollup {
	return v.value
}

func (v *NullableOrganizationImportResponseWarningRollup) Set(val *OrganizationImportResponseWarningRollup) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationImportResponseWarningRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationImportResponseWarningRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationImportResponseWarningRollup(val *OrganizationImportResponseWarningRollup) *NullableOrganizationImportResponseWarningRollup {
	return &NullableOrganizationImportResponseWarningRollup{value: val, isSet: true}
}

func (v NullableOrganizationImportResponseWarningRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationImportResponseWarningRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


