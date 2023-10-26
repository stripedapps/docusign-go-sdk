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

// checks if the AppStoreProduct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreProduct{}

// AppStoreProduct Contains information about an APP store product.
type AppStoreProduct struct {
	// 
	MarketPlace *string `json:"marketPlace,omitempty"`
	// The Product ID from the AppStore.
	ProductId *string `json:"productId,omitempty"`
}

// NewAppStoreProduct instantiates a new AppStoreProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreProduct() *AppStoreProduct {
	this := AppStoreProduct{}
	return &this
}

// NewAppStoreProductWithDefaults instantiates a new AppStoreProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreProductWithDefaults() *AppStoreProduct {
	this := AppStoreProduct{}
	return &this
}

// GetMarketPlace returns the MarketPlace field value if set, zero value otherwise.
func (o *AppStoreProduct) GetMarketPlace() string {
	if o == nil || IsNil(o.MarketPlace) {
		var ret string
		return ret
	}
	return *o.MarketPlace
}

// GetMarketPlaceOk returns a tuple with the MarketPlace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreProduct) GetMarketPlaceOk() (*string, bool) {
	if o == nil || IsNil(o.MarketPlace) {
		return nil, false
	}
	return o.MarketPlace, true
}

// HasMarketPlace returns a boolean if a field has been set.
func (o *AppStoreProduct) HasMarketPlace() bool {
	if o != nil && !IsNil(o.MarketPlace) {
		return true
	}

	return false
}

// SetMarketPlace gets a reference to the given string and assigns it to the MarketPlace field.
func (o *AppStoreProduct) SetMarketPlace(v string) {
	o.MarketPlace = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AppStoreProduct) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreProduct) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AppStoreProduct) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AppStoreProduct) SetProductId(v string) {
	o.ProductId = &v
}

func (o AppStoreProduct) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreProduct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketPlace) {
		toSerialize["marketPlace"] = o.MarketPlace
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	return toSerialize, nil
}

type NullableAppStoreProduct struct {
	value *AppStoreProduct
	isSet bool
}

func (v NullableAppStoreProduct) Get() *AppStoreProduct {
	return v.value
}

func (v *NullableAppStoreProduct) Set(val *AppStoreProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreProduct(val *AppStoreProduct) *NullableAppStoreProduct {
	return &NullableAppStoreProduct{value: val, isSet: true}
}

func (v NullableAppStoreProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


