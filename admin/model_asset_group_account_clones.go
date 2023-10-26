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

// checks if the AssetGroupAccountClones type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetGroupAccountClones{}

// AssetGroupAccountClones 
type AssetGroupAccountClones struct {
	// The list of asset group accounts.
	AssetGroupWorks []AssetGroupAccountClone `json:"assetGroupWorks,omitempty"`
}

// NewAssetGroupAccountClones instantiates a new AssetGroupAccountClones object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetGroupAccountClones() *AssetGroupAccountClones {
	this := AssetGroupAccountClones{}
	return &this
}

// NewAssetGroupAccountClonesWithDefaults instantiates a new AssetGroupAccountClones object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetGroupAccountClonesWithDefaults() *AssetGroupAccountClones {
	this := AssetGroupAccountClones{}
	return &this
}

// GetAssetGroupWorks returns the AssetGroupWorks field value if set, zero value otherwise.
func (o *AssetGroupAccountClones) GetAssetGroupWorks() []AssetGroupAccountClone {
	if o == nil || IsNil(o.AssetGroupWorks) {
		var ret []AssetGroupAccountClone
		return ret
	}
	return o.AssetGroupWorks
}

// GetAssetGroupWorksOk returns a tuple with the AssetGroupWorks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetGroupAccountClones) GetAssetGroupWorksOk() ([]AssetGroupAccountClone, bool) {
	if o == nil || IsNil(o.AssetGroupWorks) {
		return nil, false
	}
	return o.AssetGroupWorks, true
}

// HasAssetGroupWorks returns a boolean if a field has been set.
func (o *AssetGroupAccountClones) HasAssetGroupWorks() bool {
	if o != nil && !IsNil(o.AssetGroupWorks) {
		return true
	}

	return false
}

// SetAssetGroupWorks gets a reference to the given []AssetGroupAccountClone and assigns it to the AssetGroupWorks field.
func (o *AssetGroupAccountClones) SetAssetGroupWorks(v []AssetGroupAccountClone) {
	o.AssetGroupWorks = v
}

func (o AssetGroupAccountClones) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetGroupAccountClones) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssetGroupWorks) {
		toSerialize["assetGroupWorks"] = o.AssetGroupWorks
	}
	return toSerialize, nil
}

type NullableAssetGroupAccountClones struct {
	value *AssetGroupAccountClones
	isSet bool
}

func (v NullableAssetGroupAccountClones) Get() *AssetGroupAccountClones {
	return v.value
}

func (v *NullableAssetGroupAccountClones) Set(val *AssetGroupAccountClones) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetGroupAccountClones) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetGroupAccountClones) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetGroupAccountClones(val *AssetGroupAccountClones) *NullableAssetGroupAccountClones {
	return &NullableAssetGroupAccountClones{value: val, isSet: true}
}

func (v NullableAssetGroupAccountClones) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetGroupAccountClones) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


