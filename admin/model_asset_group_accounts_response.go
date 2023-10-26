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

// checks if the AssetGroupAccountsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetGroupAccountsResponse{}

// AssetGroupAccountsResponse A list of asset group accounts.
type AssetGroupAccountsResponse struct {
	// The list of asset group accounts.
	AssetGroupAccounts []AssetGroupAccountResponse `json:"assetGroupAccounts,omitempty"`
}

// NewAssetGroupAccountsResponse instantiates a new AssetGroupAccountsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetGroupAccountsResponse() *AssetGroupAccountsResponse {
	this := AssetGroupAccountsResponse{}
	return &this
}

// NewAssetGroupAccountsResponseWithDefaults instantiates a new AssetGroupAccountsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetGroupAccountsResponseWithDefaults() *AssetGroupAccountsResponse {
	this := AssetGroupAccountsResponse{}
	return &this
}

// GetAssetGroupAccounts returns the AssetGroupAccounts field value if set, zero value otherwise.
func (o *AssetGroupAccountsResponse) GetAssetGroupAccounts() []AssetGroupAccountResponse {
	if o == nil || IsNil(o.AssetGroupAccounts) {
		var ret []AssetGroupAccountResponse
		return ret
	}
	return o.AssetGroupAccounts
}

// GetAssetGroupAccountsOk returns a tuple with the AssetGroupAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGroupAccountsResponse) GetAssetGroupAccountsOk() ([]AssetGroupAccountResponse, bool) {
	if o == nil || IsNil(o.AssetGroupAccounts) {
		return nil, false
	}
	return o.AssetGroupAccounts, true
}

// HasAssetGroupAccounts returns a boolean if a field has been set.
func (o *AssetGroupAccountsResponse) HasAssetGroupAccounts() bool {
	if o != nil && !IsNil(o.AssetGroupAccounts) {
		return true
	}

	return false
}

// SetAssetGroupAccounts gets a reference to the given []AssetGroupAccountResponse and assigns it to the AssetGroupAccounts field.
func (o *AssetGroupAccountsResponse) SetAssetGroupAccounts(v []AssetGroupAccountResponse) {
	o.AssetGroupAccounts = v
}

func (o AssetGroupAccountsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetGroupAccountsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetGroupAccounts) {
		toSerialize["assetGroupAccounts"] = o.AssetGroupAccounts
	}
	return toSerialize, nil
}

type NullableAssetGroupAccountsResponse struct {
	value *AssetGroupAccountsResponse
	isSet bool
}

func (v NullableAssetGroupAccountsResponse) Get() *AssetGroupAccountsResponse {
	return v.value
}

func (v *NullableAssetGroupAccountsResponse) Set(val *AssetGroupAccountsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetGroupAccountsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetGroupAccountsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetGroupAccountsResponse(val *AssetGroupAccountsResponse) *NullableAssetGroupAccountsResponse {
	return &NullableAssetGroupAccountsResponse{value: val, isSet: true}
}

func (v NullableAssetGroupAccountsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetGroupAccountsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

