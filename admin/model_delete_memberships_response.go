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

// checks if the DeleteMembershipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteMembershipsResponse{}

// DeleteMembershipsResponse The results of closing a user's account.
type DeleteMembershipsResponse struct {
	// When **true,** the request to close the accounts succeeded.
	Success *bool `json:"success,omitempty"`
	// A list of accounts that were closed.
	Accounts []DeleteMembershipResponse `json:"accounts,omitempty"`
}

// NewDeleteMembershipsResponse instantiates a new DeleteMembershipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteMembershipsResponse() *DeleteMembershipsResponse {
	this := DeleteMembershipsResponse{}
	return &this
}

// NewDeleteMembershipsResponseWithDefaults instantiates a new DeleteMembershipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteMembershipsResponseWithDefaults() *DeleteMembershipsResponse {
	this := DeleteMembershipsResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DeleteMembershipsResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMembershipsResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DeleteMembershipsResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DeleteMembershipsResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *DeleteMembershipsResponse) GetAccounts() []DeleteMembershipResponse {
	if o == nil || IsNil(o.Accounts) {
		var ret []DeleteMembershipResponse
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteMembershipsResponse) GetAccountsOk() ([]DeleteMembershipResponse, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *DeleteMembershipsResponse) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []DeleteMembershipResponse and assigns it to the Accounts field.
func (o *DeleteMembershipsResponse) SetAccounts(v []DeleteMembershipResponse) {
	o.Accounts = v
}

func (o DeleteMembershipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteMembershipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableDeleteMembershipsResponse struct {
	value *DeleteMembershipsResponse
	isSet bool
}

func (v NullableDeleteMembershipsResponse) Get() *DeleteMembershipsResponse {
	return v.value
}

func (v *NullableDeleteMembershipsResponse) Set(val *DeleteMembershipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteMembershipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteMembershipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteMembershipsResponse(val *DeleteMembershipsResponse) *NullableDeleteMembershipsResponse {
	return &NullableDeleteMembershipsResponse{value: val, isSet: true}
}

func (v NullableDeleteMembershipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteMembershipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


