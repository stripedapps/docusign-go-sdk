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

// checks if the ProductPermissionProfilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductPermissionProfilesResponse{}

// ProductPermissionProfilesResponse 
type ProductPermissionProfilesResponse struct {
	// A list of one or more products and their respective permissions.
	ProductPermissionProfiles []ProductPermissionProfileResponse `json:"product_permission_profiles,omitempty"`
}

// NewProductPermissionProfilesResponse instantiates a new ProductPermissionProfilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductPermissionProfilesResponse() *ProductPermissionProfilesResponse {
	this := ProductPermissionProfilesResponse{}
	return &this
}

// NewProductPermissionProfilesResponseWithDefaults instantiates a new ProductPermissionProfilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductPermissionProfilesResponseWithDefaults() *ProductPermissionProfilesResponse {
	this := ProductPermissionProfilesResponse{}
	return &this
}

// GetProductPermissionProfiles returns the ProductPermissionProfiles field value if set, zero value otherwise.
func (o *ProductPermissionProfilesResponse) GetProductPermissionProfiles() []ProductPermissionProfileResponse {
	if o == nil || IsNil(o.ProductPermissionProfiles) {
		var ret []ProductPermissionProfileResponse
		return ret
	}
	return o.ProductPermissionProfiles
}

// GetProductPermissionProfilesOk returns a tuple with the ProductPermissionProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductPermissionProfilesResponse) GetProductPermissionProfilesOk() ([]ProductPermissionProfileResponse, bool) {
	if o == nil || IsNil(o.ProductPermissionProfiles) {
		return nil, false
	}
	return o.ProductPermissionProfiles, true
}

// HasProductPermissionProfiles returns a boolean if a field has been set.
func (o *ProductPermissionProfilesResponse) HasProductPermissionProfiles() bool {
	if o != nil && !IsNil(o.ProductPermissionProfiles) {
		return true
	}

	return false
}

// SetProductPermissionProfiles gets a reference to the given []ProductPermissionProfileResponse and assigns it to the ProductPermissionProfiles field.
func (o *ProductPermissionProfilesResponse) SetProductPermissionProfiles(v []ProductPermissionProfileResponse) {
	o.ProductPermissionProfiles = v
}

func (o ProductPermissionProfilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductPermissionProfilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductPermissionProfiles) {
		toSerialize["product_permission_profiles"] = o.ProductPermissionProfiles
	}
	return toSerialize, nil
}

type NullableProductPermissionProfilesResponse struct {
	value *ProductPermissionProfilesResponse
	isSet bool
}

func (v NullableProductPermissionProfilesResponse) Get() *ProductPermissionProfilesResponse {
	return v.value
}

func (v *NullableProductPermissionProfilesResponse) Set(val *ProductPermissionProfilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductPermissionProfilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductPermissionProfilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductPermissionProfilesResponse(val *ProductPermissionProfilesResponse) *NullableProductPermissionProfilesResponse {
	return &NullableProductPermissionProfilesResponse{value: val, isSet: true}
}

func (v NullableProductPermissionProfilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductPermissionProfilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


