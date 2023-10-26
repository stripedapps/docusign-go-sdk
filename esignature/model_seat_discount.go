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

// checks if the SeatDiscount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeatDiscount{}

// SeatDiscount This object contains information about a seat discount.
type SeatDiscount struct {
	// Reserved for DocuSign.
	BeginSeatCount *string `json:"beginSeatCount,omitempty"`
	// The percent of the discount.   Example: `\"0.00\"`
	DiscountPercent *string `json:"discountPercent,omitempty"`
	// Reserved for DocuSign.
	EndSeatCount *string `json:"endSeatCount,omitempty"`
}

// NewSeatDiscount instantiates a new SeatDiscount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeatDiscount() *SeatDiscount {
	this := SeatDiscount{}
	return &this
}

// NewSeatDiscountWithDefaults instantiates a new SeatDiscount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeatDiscountWithDefaults() *SeatDiscount {
	this := SeatDiscount{}
	return &this
}

// GetBeginSeatCount returns the BeginSeatCount field value if set, zero value otherwise.
func (o *SeatDiscount) GetBeginSeatCount() string {
	if o == nil || IsNil(o.BeginSeatCount) {
		var ret string
		return ret
	}
	return *o.BeginSeatCount
}

// GetBeginSeatCountOk returns a tuple with the BeginSeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatDiscount) GetBeginSeatCountOk() (*string, bool) {
	if o == nil || IsNil(o.BeginSeatCount) {
		return nil, false
	}
	return o.BeginSeatCount, true
}

// HasBeginSeatCount returns a boolean if a field has been set.
func (o *SeatDiscount) HasBeginSeatCount() bool {
	if o != nil && !IsNil(o.BeginSeatCount) {
		return true
	}

	return false
}

// SetBeginSeatCount gets a reference to the given string and assigns it to the BeginSeatCount field.
func (o *SeatDiscount) SetBeginSeatCount(v string) {
	o.BeginSeatCount = &v
}

// GetDiscountPercent returns the DiscountPercent field value if set, zero value otherwise.
func (o *SeatDiscount) GetDiscountPercent() string {
	if o == nil || IsNil(o.DiscountPercent) {
		var ret string
		return ret
	}
	return *o.DiscountPercent
}

// GetDiscountPercentOk returns a tuple with the DiscountPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatDiscount) GetDiscountPercentOk() (*string, bool) {
	if o == nil || IsNil(o.DiscountPercent) {
		return nil, false
	}
	return o.DiscountPercent, true
}

// HasDiscountPercent returns a boolean if a field has been set.
func (o *SeatDiscount) HasDiscountPercent() bool {
	if o != nil && !IsNil(o.DiscountPercent) {
		return true
	}

	return false
}

// SetDiscountPercent gets a reference to the given string and assigns it to the DiscountPercent field.
func (o *SeatDiscount) SetDiscountPercent(v string) {
	o.DiscountPercent = &v
}

// GetEndSeatCount returns the EndSeatCount field value if set, zero value otherwise.
func (o *SeatDiscount) GetEndSeatCount() string {
	if o == nil || IsNil(o.EndSeatCount) {
		var ret string
		return ret
	}
	return *o.EndSeatCount
}

// GetEndSeatCountOk returns a tuple with the EndSeatCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeatDiscount) GetEndSeatCountOk() (*string, bool) {
	if o == nil || IsNil(o.EndSeatCount) {
		return nil, false
	}
	return o.EndSeatCount, true
}

// HasEndSeatCount returns a boolean if a field has been set.
func (o *SeatDiscount) HasEndSeatCount() bool {
	if o != nil && !IsNil(o.EndSeatCount) {
		return true
	}

	return false
}

// SetEndSeatCount gets a reference to the given string and assigns it to the EndSeatCount field.
func (o *SeatDiscount) SetEndSeatCount(v string) {
	o.EndSeatCount = &v
}

func (o SeatDiscount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeatDiscount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginSeatCount) {
		toSerialize["beginSeatCount"] = o.BeginSeatCount
	}
	if !IsNil(o.DiscountPercent) {
		toSerialize["discountPercent"] = o.DiscountPercent
	}
	if !IsNil(o.EndSeatCount) {
		toSerialize["endSeatCount"] = o.EndSeatCount
	}
	return toSerialize, nil
}

type NullableSeatDiscount struct {
	value *SeatDiscount
	isSet bool
}

func (v NullableSeatDiscount) Get() *SeatDiscount {
	return v.value
}

func (v *NullableSeatDiscount) Set(val *SeatDiscount) {
	v.value = val
	v.isSet = true
}

func (v NullableSeatDiscount) IsSet() bool {
	return v.isSet
}

func (v *NullableSeatDiscount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeatDiscount(val *SeatDiscount) *NullableSeatDiscount {
	return &NullableSeatDiscount{value: val, isSet: true}
}

func (v NullableSeatDiscount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeatDiscount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


