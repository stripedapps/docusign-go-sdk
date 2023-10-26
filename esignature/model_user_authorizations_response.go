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

// checks if the UserAuthorizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserAuthorizationsResponse{}

// UserAuthorizationsResponse 
type UserAuthorizationsResponse struct {
	// 
	Results []UserAuthorizationWithStatus `json:"results,omitempty"`
}

// NewUserAuthorizationsResponse instantiates a new UserAuthorizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserAuthorizationsResponse() *UserAuthorizationsResponse {
	this := UserAuthorizationsResponse{}
	return &this
}

// NewUserAuthorizationsResponseWithDefaults instantiates a new UserAuthorizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserAuthorizationsResponseWithDefaults() *UserAuthorizationsResponse {
	this := UserAuthorizationsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *UserAuthorizationsResponse) GetResults() []UserAuthorizationWithStatus {
	if o == nil || IsNil(o.Results) {
		var ret []UserAuthorizationWithStatus
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserAuthorizationsResponse) GetResultsOk() ([]UserAuthorizationWithStatus, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *UserAuthorizationsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []UserAuthorizationWithStatus and assigns it to the Results field.
func (o *UserAuthorizationsResponse) SetResults(v []UserAuthorizationWithStatus) {
	o.Results = v
}

func (o UserAuthorizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserAuthorizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableUserAuthorizationsResponse struct {
	value *UserAuthorizationsResponse
	isSet bool
}

func (v NullableUserAuthorizationsResponse) Get() *UserAuthorizationsResponse {
	return v.value
}

func (v *NullableUserAuthorizationsResponse) Set(val *UserAuthorizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserAuthorizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserAuthorizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserAuthorizationsResponse(val *UserAuthorizationsResponse) *NullableUserAuthorizationsResponse {
	return &NullableUserAuthorizationsResponse{value: val, isSet: true}
}

func (v NullableUserAuthorizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserAuthorizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

