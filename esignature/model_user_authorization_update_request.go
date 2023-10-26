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

// checks if the UserAuthorizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAuthorizationUpdateRequest{}

// UserAuthorizationUpdateRequest The request object to update a user authorization.
type UserAuthorizationUpdateRequest struct {
	// The end date for the user authorization. The default value is the max UTC value: `9999-12-31T23:59:59.0000000+00:00`.
	EndDate *string `json:"endDate,omitempty"`
	// The start date for the user authorization. The default value is the current date and time.
	StartDate *string `json:"startDate,omitempty"`
}

// NewUserAuthorizationUpdateRequest instantiates a new UserAuthorizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthorizationUpdateRequest() *UserAuthorizationUpdateRequest {
	this := UserAuthorizationUpdateRequest{}
	return &this
}

// NewUserAuthorizationUpdateRequestWithDefaults instantiates a new UserAuthorizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthorizationUpdateRequestWithDefaults() *UserAuthorizationUpdateRequest {
	this := UserAuthorizationUpdateRequest{}
	return &this
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *UserAuthorizationUpdateRequest) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthorizationUpdateRequest) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *UserAuthorizationUpdateRequest) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *UserAuthorizationUpdateRequest) SetEndDate(v string) {
	o.EndDate = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UserAuthorizationUpdateRequest) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthorizationUpdateRequest) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UserAuthorizationUpdateRequest) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *UserAuthorizationUpdateRequest) SetStartDate(v string) {
	o.StartDate = &v
}

func (o UserAuthorizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAuthorizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableUserAuthorizationUpdateRequest struct {
	value *UserAuthorizationUpdateRequest
	isSet bool
}

func (v NullableUserAuthorizationUpdateRequest) Get() *UserAuthorizationUpdateRequest {
	return v.value
}

func (v *NullableUserAuthorizationUpdateRequest) Set(val *UserAuthorizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthorizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthorizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthorizationUpdateRequest(val *UserAuthorizationUpdateRequest) *NullableUserAuthorizationUpdateRequest {
	return &NullableUserAuthorizationUpdateRequest{value: val, isSet: true}
}

func (v NullableUserAuthorizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthorizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


