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

// checks if the BillingDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingDiscount{}

// BillingDiscount 
type BillingDiscount struct {
	// Reserved for DocuSign.
	BeginQuantity *string `json:"beginQuantity,omitempty"`
	// 
	Discount *string `json:"discount,omitempty"`
	// 
	EndQuantity *string `json:"endQuantity,omitempty"`
}

// NewBillingDiscount instantiates a new BillingDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingDiscount() *BillingDiscount {
	this := BillingDiscount{}
	return &this
}

// NewBillingDiscountWithDefaults instantiates a new BillingDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingDiscountWithDefaults() *BillingDiscount {
	this := BillingDiscount{}
	return &this
}

// GetBeginQuantity returns the BeginQuantity field value if set, zero value otherwise.
func (o *BillingDiscount) GetBeginQuantity() string {
	if o == nil || IsNil(o.BeginQuantity) {
		var ret string
		return ret
	}
	return *o.BeginQuantity
}

// GetBeginQuantityOk returns a tuple with the BeginQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDiscount) GetBeginQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.BeginQuantity) {
		return nil, false
	}
	return o.BeginQuantity, true
}

// HasBeginQuantity returns a boolean if a field has been set.
func (o *BillingDiscount) HasBeginQuantity() bool {
	if o != nil && !IsNil(o.BeginQuantity) {
		return true
	}

	return false
}

// SetBeginQuantity gets a reference to the given string and assigns it to the BeginQuantity field.
func (o *BillingDiscount) SetBeginQuantity(v string) {
	o.BeginQuantity = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *BillingDiscount) GetDiscount() string {
	if o == nil || IsNil(o.Discount) {
		var ret string
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDiscount) GetDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *BillingDiscount) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given string and assigns it to the Discount field.
func (o *BillingDiscount) SetDiscount(v string) {
	o.Discount = &v
}

// GetEndQuantity returns the EndQuantity field value if set, zero value otherwise.
func (o *BillingDiscount) GetEndQuantity() string {
	if o == nil || IsNil(o.EndQuantity) {
		var ret string
		return ret
	}
	return *o.EndQuantity
}

// GetEndQuantityOk returns a tuple with the EndQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDiscount) GetEndQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.EndQuantity) {
		return nil, false
	}
	return o.EndQuantity, true
}

// HasEndQuantity returns a boolean if a field has been set.
func (o *BillingDiscount) HasEndQuantity() bool {
	if o != nil && !IsNil(o.EndQuantity) {
		return true
	}

	return false
}

// SetEndQuantity gets a reference to the given string and assigns it to the EndQuantity field.
func (o *BillingDiscount) SetEndQuantity(v string) {
	o.EndQuantity = &v
}

func (o BillingDiscount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginQuantity) {
		toSerialize["beginQuantity"] = o.BeginQuantity
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.EndQuantity) {
		toSerialize["endQuantity"] = o.EndQuantity
	}
	return toSerialize, nil
}

type NullableBillingDiscount struct {
	value *BillingDiscount
	isSet bool
}

func (v NullableBillingDiscount) Get() *BillingDiscount {
	return v.value
}

func (v *NullableBillingDiscount) Set(val *BillingDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingDiscount(val *BillingDiscount) *NullableBillingDiscount {
	return &NullableBillingDiscount{value: val, isSet: true}
}

func (v NullableBillingDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

