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

// checks if the ProductPermissionProfilesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductPermissionProfilesRequest{}

// ProductPermissionProfilesRequest 
type ProductPermissionProfilesRequest struct {
	// A list of one or more products and their respective permissions.
	ProductPermissionProfiles []ProductPermissionProfileRequest `json:"product_permission_profiles"`
}

// NewProductPermissionProfilesRequest instantiates a new ProductPermissionProfilesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductPermissionProfilesRequest(productPermissionProfiles []ProductPermissionProfileRequest) *ProductPermissionProfilesRequest {
	this := ProductPermissionProfilesRequest{}
	this.ProductPermissionProfiles = productPermissionProfiles
	return &this
}

// NewProductPermissionProfilesRequestWithDefaults instantiates a new ProductPermissionProfilesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPermissionProfilesRequestWithDefaults() *ProductPermissionProfilesRequest {
	this := ProductPermissionProfilesRequest{}
	return &this
}

// GetProductPermissionProfiles returns the ProductPermissionProfiles field value
func (o *ProductPermissionProfilesRequest) GetProductPermissionProfiles() []ProductPermissionProfileRequest {
	if o == nil {
		var ret []ProductPermissionProfileRequest
		return ret
	}

	return o.ProductPermissionProfiles
}

// GetProductPermissionProfilesOk returns a tuple with the ProductPermissionProfiles field value
// and a boolean to check if the value has been set.
func (o *ProductPermissionProfilesRequest) GetProductPermissionProfilesOk() ([]ProductPermissionProfileRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductPermissionProfiles, true
}

// SetProductPermissionProfiles sets field value
func (o *ProductPermissionProfilesRequest) SetProductPermissionProfiles(v []ProductPermissionProfileRequest) {
	o.ProductPermissionProfiles = v
}

func (o ProductPermissionProfilesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductPermissionProfilesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["product_permission_profiles"] = o.ProductPermissionProfiles
	return toSerialize, nil
}

type NullableProductPermissionProfilesRequest struct {
	value *ProductPermissionProfilesRequest
	isSet bool
}

func (v NullableProductPermissionProfilesRequest) Get() *ProductPermissionProfilesRequest {
	return v.value
}

func (v *NullableProductPermissionProfilesRequest) Set(val *ProductPermissionProfilesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProductPermissionProfilesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProductPermissionProfilesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductPermissionProfilesRequest(val *ProductPermissionProfilesRequest) *NullableProductPermissionProfilesRequest {
	return &NullableProductPermissionProfilesRequest{value: val, isSet: true}
}

func (v NullableProductPermissionProfilesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductPermissionProfilesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


