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

// checks if the AccountPasswordMinimumPasswordAgeDays type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountPasswordMinimumPasswordAgeDays{}

// AccountPasswordMinimumPasswordAgeDays 
type AccountPasswordMinimumPasswordAgeDays struct {
	// 
	MaximumAge *string `json:"maximumAge,omitempty"`
	// 
	MinimumAge *string `json:"minimumAge,omitempty"`
}

// NewAccountPasswordMinimumPasswordAgeDays instantiates a new AccountPasswordMinimumPasswordAgeDays object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountPasswordMinimumPasswordAgeDays() *AccountPasswordMinimumPasswordAgeDays {
	this := AccountPasswordMinimumPasswordAgeDays{}
	return &this
}

// NewAccountPasswordMinimumPasswordAgeDaysWithDefaults instantiates a new AccountPasswordMinimumPasswordAgeDays object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountPasswordMinimumPasswordAgeDaysWithDefaults() *AccountPasswordMinimumPasswordAgeDays {
	this := AccountPasswordMinimumPasswordAgeDays{}
	return &this
}

// GetMaximumAge returns the MaximumAge field value if set, zero value otherwise.
func (o *AccountPasswordMinimumPasswordAgeDays) GetMaximumAge() string {
	if o == nil || IsNil(o.MaximumAge) {
		var ret string
		return ret
	}
	return *o.MaximumAge
}

// GetMaximumAgeOk returns a tuple with the MaximumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPasswordMinimumPasswordAgeDays) GetMaximumAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MaximumAge) {
		return nil, false
	}
	return o.MaximumAge, true
}

// HasMaximumAge returns a boolean if a field has been set.
func (o *AccountPasswordMinimumPasswordAgeDays) HasMaximumAge() bool {
	if o != nil && !IsNil(o.MaximumAge) {
		return true
	}

	return false
}

// SetMaximumAge gets a reference to the given string and assigns it to the MaximumAge field.
func (o *AccountPasswordMinimumPasswordAgeDays) SetMaximumAge(v string) {
	o.MaximumAge = &v
}

// GetMinimumAge returns the MinimumAge field value if set, zero value otherwise.
func (o *AccountPasswordMinimumPasswordAgeDays) GetMinimumAge() string {
	if o == nil || IsNil(o.MinimumAge) {
		var ret string
		return ret
	}
	return *o.MinimumAge
}

// GetMinimumAgeOk returns a tuple with the MinimumAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountPasswordMinimumPasswordAgeDays) GetMinimumAgeOk() (*string, bool) {
	if o == nil || IsNil(o.MinimumAge) {
		return nil, false
	}
	return o.MinimumAge, true
}

// HasMinimumAge returns a boolean if a field has been set.
func (o *AccountPasswordMinimumPasswordAgeDays) HasMinimumAge() bool {
	if o != nil && !IsNil(o.MinimumAge) {
		return true
	}

	return false
}

// SetMinimumAge gets a reference to the given string and assigns it to the MinimumAge field.
func (o *AccountPasswordMinimumPasswordAgeDays) SetMinimumAge(v string) {
	o.MinimumAge = &v
}

func (o AccountPasswordMinimumPasswordAgeDays) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountPasswordMinimumPasswordAgeDays) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaximumAge) {
		toSerialize["maximumAge"] = o.MaximumAge
	}
	if !IsNil(o.MinimumAge) {
		toSerialize["minimumAge"] = o.MinimumAge
	}
	return toSerialize, nil
}

type NullableAccountPasswordMinimumPasswordAgeDays struct {
	value *AccountPasswordMinimumPasswordAgeDays
	isSet bool
}

func (v NullableAccountPasswordMinimumPasswordAgeDays) Get() *AccountPasswordMinimumPasswordAgeDays {
	return v.value
}

func (v *NullableAccountPasswordMinimumPasswordAgeDays) Set(val *AccountPasswordMinimumPasswordAgeDays) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountPasswordMinimumPasswordAgeDays) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountPasswordMinimumPasswordAgeDays) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountPasswordMinimumPasswordAgeDays(val *AccountPasswordMinimumPasswordAgeDays) *NullableAccountPasswordMinimumPasswordAgeDays {
	return &NullableAccountPasswordMinimumPasswordAgeDays{value: val, isSet: true}
}

func (v NullableAccountPasswordMinimumPasswordAgeDays) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountPasswordMinimumPasswordAgeDays) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

