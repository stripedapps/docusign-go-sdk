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

// checks if the CurrencyFeatureSetPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CurrencyFeatureSetPrice{}

// CurrencyFeatureSetPrice Information about the price and currency associated with the feature set. Reserved for internal DocuSign use only.
type CurrencyFeatureSetPrice struct {
	// Specifies the alternate ISO currency code for the account. 
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// Reserved for DocuSign.
	CurrencySymbol *string `json:"currencySymbol,omitempty"`
	// Reserved for DocuSign.
	EnvelopeFee *string `json:"envelopeFee,omitempty"`
	// Reserved for DocuSign.
	FixedFee *string `json:"fixedFee,omitempty"`
	// Reserved for DocuSign.
	SeatFee *string `json:"seatFee,omitempty"`
}

// NewCurrencyFeatureSetPrice instantiates a new CurrencyFeatureSetPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCurrencyFeatureSetPrice() *CurrencyFeatureSetPrice {
	this := CurrencyFeatureSetPrice{}
	return &this
}

// NewCurrencyFeatureSetPriceWithDefaults instantiates a new CurrencyFeatureSetPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCurrencyFeatureSetPriceWithDefaults() *CurrencyFeatureSetPrice {
	this := CurrencyFeatureSetPrice{}
	return &this
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *CurrencyFeatureSetPrice) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyFeatureSetPrice) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *CurrencyFeatureSetPrice) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *CurrencyFeatureSetPrice) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetCurrencySymbol returns the CurrencySymbol field value if set, zero value otherwise.
func (o *CurrencyFeatureSetPrice) GetCurrencySymbol() string {
	if o == nil || IsNil(o.CurrencySymbol) {
		var ret string
		return ret
	}
	return *o.CurrencySymbol
}

// GetCurrencySymbolOk returns a tuple with the CurrencySymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyFeatureSetPrice) GetCurrencySymbolOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencySymbol) {
		return nil, false
	}
	return o.CurrencySymbol, true
}

// HasCurrencySymbol returns a boolean if a field has been set.
func (o *CurrencyFeatureSetPrice) HasCurrencySymbol() bool {
	if o != nil && !IsNil(o.CurrencySymbol) {
		return true
	}

	return false
}

// SetCurrencySymbol gets a reference to the given string and assigns it to the CurrencySymbol field.
func (o *CurrencyFeatureSetPrice) SetCurrencySymbol(v string) {
	o.CurrencySymbol = &v
}

// GetEnvelopeFee returns the EnvelopeFee field value if set, zero value otherwise.
func (o *CurrencyFeatureSetPrice) GetEnvelopeFee() string {
	if o == nil || IsNil(o.EnvelopeFee) {
		var ret string
		return ret
	}
	return *o.EnvelopeFee
}

// GetEnvelopeFeeOk returns a tuple with the EnvelopeFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyFeatureSetPrice) GetEnvelopeFeeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeFee) {
		return nil, false
	}
	return o.EnvelopeFee, true
}

// HasEnvelopeFee returns a boolean if a field has been set.
func (o *CurrencyFeatureSetPrice) HasEnvelopeFee() bool {
	if o != nil && !IsNil(o.EnvelopeFee) {
		return true
	}

	return false
}

// SetEnvelopeFee gets a reference to the given string and assigns it to the EnvelopeFee field.
func (o *CurrencyFeatureSetPrice) SetEnvelopeFee(v string) {
	o.EnvelopeFee = &v
}

// GetFixedFee returns the FixedFee field value if set, zero value otherwise.
func (o *CurrencyFeatureSetPrice) GetFixedFee() string {
	if o == nil || IsNil(o.FixedFee) {
		var ret string
		return ret
	}
	return *o.FixedFee
}

// GetFixedFeeOk returns a tuple with the FixedFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyFeatureSetPrice) GetFixedFeeOk() (*string, bool) {
	if o == nil || IsNil(o.FixedFee) {
		return nil, false
	}
	return o.FixedFee, true
}

// HasFixedFee returns a boolean if a field has been set.
func (o *CurrencyFeatureSetPrice) HasFixedFee() bool {
	if o != nil && !IsNil(o.FixedFee) {
		return true
	}

	return false
}

// SetFixedFee gets a reference to the given string and assigns it to the FixedFee field.
func (o *CurrencyFeatureSetPrice) SetFixedFee(v string) {
	o.FixedFee = &v
}

// GetSeatFee returns the SeatFee field value if set, zero value otherwise.
func (o *CurrencyFeatureSetPrice) GetSeatFee() string {
	if o == nil || IsNil(o.SeatFee) {
		var ret string
		return ret
	}
	return *o.SeatFee
}

// GetSeatFeeOk returns a tuple with the SeatFee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CurrencyFeatureSetPrice) GetSeatFeeOk() (*string, bool) {
	if o == nil || IsNil(o.SeatFee) {
		return nil, false
	}
	return o.SeatFee, true
}

// HasSeatFee returns a boolean if a field has been set.
func (o *CurrencyFeatureSetPrice) HasSeatFee() bool {
	if o != nil && !IsNil(o.SeatFee) {
		return true
	}

	return false
}

// SetSeatFee gets a reference to the given string and assigns it to the SeatFee field.
func (o *CurrencyFeatureSetPrice) SetSeatFee(v string) {
	o.SeatFee = &v
}

func (o CurrencyFeatureSetPrice) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CurrencyFeatureSetPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.CurrencySymbol) {
		toSerialize["currencySymbol"] = o.CurrencySymbol
	}
	if !IsNil(o.EnvelopeFee) {
		toSerialize["envelopeFee"] = o.EnvelopeFee
	}
	if !IsNil(o.FixedFee) {
		toSerialize["fixedFee"] = o.FixedFee
	}
	if !IsNil(o.SeatFee) {
		toSerialize["seatFee"] = o.SeatFee
	}
	return toSerialize, nil
}

type NullableCurrencyFeatureSetPrice struct {
	value *CurrencyFeatureSetPrice
	isSet bool
}

func (v NullableCurrencyFeatureSetPrice) Get() *CurrencyFeatureSetPrice {
	return v.value
}

func (v *NullableCurrencyFeatureSetPrice) Set(val *CurrencyFeatureSetPrice) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrencyFeatureSetPrice) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrencyFeatureSetPrice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrencyFeatureSetPrice(val *CurrencyFeatureSetPrice) *NullableCurrencyFeatureSetPrice {
	return &NullableCurrencyFeatureSetPrice{value: val, isSet: true}
}

func (v NullableCurrencyFeatureSetPrice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrencyFeatureSetPrice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


