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

// checks if the DeleteResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteResponse{}

// DeleteResponse Results of deleting identities.
type DeleteResponse struct {
	// When **true,** the request succeeded.
	Success *bool `json:"success,omitempty"`
	// A list of identities to delete.
	Identities []UserIdentityResponse `json:"identities,omitempty"`
}

// NewDeleteResponse instantiates a new DeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResponse() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResponseWithDefaults() *DeleteResponse {
	this := DeleteResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *DeleteResponse) GetIdentities() []UserIdentityResponse {
	if o == nil || IsNil(o.Identities) {
		var ret []UserIdentityResponse
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteResponse) GetIdentitiesOk() ([]UserIdentityResponse, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *DeleteResponse) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []UserIdentityResponse and assigns it to the Identities field.
func (o *DeleteResponse) SetIdentities(v []UserIdentityResponse) {
	o.Identities = v
}

func (o DeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}
	return toSerialize, nil
}

type NullableDeleteResponse struct {
	value *DeleteResponse
	isSet bool
}

func (v NullableDeleteResponse) Get() *DeleteResponse {
	return v.value
}

func (v *NullableDeleteResponse) Set(val *DeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResponse(val *DeleteResponse) *NullableDeleteResponse {
	return &NullableDeleteResponse{value: val, isSet: true}
}

func (v NullableDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

