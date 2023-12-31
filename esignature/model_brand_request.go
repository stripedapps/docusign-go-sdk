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

// checks if the BrandRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandRequest{}

// BrandRequest This request object contains information about a specific brand.
type BrandRequest struct {
	// The ID of the brand used in API calls
	BrandId *string `json:"brandId,omitempty"`
}

// NewBrandRequest instantiates a new BrandRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandRequest() *BrandRequest {
	this := BrandRequest{}
	return &this
}

// NewBrandRequestWithDefaults instantiates a new BrandRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandRequestWithDefaults() *BrandRequest {
	this := BrandRequest{}
	return &this
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *BrandRequest) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandRequest) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *BrandRequest) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *BrandRequest) SetBrandId(v string) {
	o.BrandId = &v
}

func (o BrandRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandId) {
		toSerialize["brandId"] = o.BrandId
	}
	return toSerialize, nil
}

type NullableBrandRequest struct {
	value *BrandRequest
	isSet bool
}

func (v NullableBrandRequest) Get() *BrandRequest {
	return v.value
}

func (v *NullableBrandRequest) Set(val *BrandRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandRequest(val *BrandRequest) *NullableBrandRequest {
	return &NullableBrandRequest{value: val, isSet: true}
}

func (v NullableBrandRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


