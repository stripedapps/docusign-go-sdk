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

// checks if the AccountBrands type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountBrands{}

// AccountBrands The AccountBrands resource enables you to use account-level brands to customize the styles and text that recipients see.
type AccountBrands struct {
	// A list of brands.
	Brands []Brand `json:"brands,omitempty"`
	// The brand that envelope recipients see when a brand is not explicitly set.
	RecipientBrandIdDefault *string `json:"recipientBrandIdDefault,omitempty"`
	// The brand that envelope senders see when a brand is not explicitly set.
	SenderBrandIdDefault *string `json:"senderBrandIdDefault,omitempty"`
}

// NewAccountBrands instantiates a new AccountBrands object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBrands() *AccountBrands {
	this := AccountBrands{}
	return &this
}

// NewAccountBrandsWithDefaults instantiates a new AccountBrands object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBrandsWithDefaults() *AccountBrands {
	this := AccountBrands{}
	return &this
}

// GetBrands returns the Brands field value if set, zero value otherwise.
func (o *AccountBrands) GetBrands() []Brand {
	if o == nil || IsNil(o.Brands) {
		var ret []Brand
		return ret
	}
	return o.Brands
}

// GetBrandsOk returns a tuple with the Brands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBrands) GetBrandsOk() ([]Brand, bool) {
	if o == nil || IsNil(o.Brands) {
		return nil, false
	}
	return o.Brands, true
}

// HasBrands returns a boolean if a field has been set.
func (o *AccountBrands) HasBrands() bool {
	if o != nil && !IsNil(o.Brands) {
		return true
	}

	return false
}

// SetBrands gets a reference to the given []Brand and assigns it to the Brands field.
func (o *AccountBrands) SetBrands(v []Brand) {
	o.Brands = v
}

// GetRecipientBrandIdDefault returns the RecipientBrandIdDefault field value if set, zero value otherwise.
func (o *AccountBrands) GetRecipientBrandIdDefault() string {
	if o == nil || IsNil(o.RecipientBrandIdDefault) {
		var ret string
		return ret
	}
	return *o.RecipientBrandIdDefault
}

// GetRecipientBrandIdDefaultOk returns a tuple with the RecipientBrandIdDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBrands) GetRecipientBrandIdDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientBrandIdDefault) {
		return nil, false
	}
	return o.RecipientBrandIdDefault, true
}

// HasRecipientBrandIdDefault returns a boolean if a field has been set.
func (o *AccountBrands) HasRecipientBrandIdDefault() bool {
	if o != nil && !IsNil(o.RecipientBrandIdDefault) {
		return true
	}

	return false
}

// SetRecipientBrandIdDefault gets a reference to the given string and assigns it to the RecipientBrandIdDefault field.
func (o *AccountBrands) SetRecipientBrandIdDefault(v string) {
	o.RecipientBrandIdDefault = &v
}

// GetSenderBrandIdDefault returns the SenderBrandIdDefault field value if set, zero value otherwise.
func (o *AccountBrands) GetSenderBrandIdDefault() string {
	if o == nil || IsNil(o.SenderBrandIdDefault) {
		var ret string
		return ret
	}
	return *o.SenderBrandIdDefault
}

// GetSenderBrandIdDefaultOk returns a tuple with the SenderBrandIdDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBrands) GetSenderBrandIdDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.SenderBrandIdDefault) {
		return nil, false
	}
	return o.SenderBrandIdDefault, true
}

// HasSenderBrandIdDefault returns a boolean if a field has been set.
func (o *AccountBrands) HasSenderBrandIdDefault() bool {
	if o != nil && !IsNil(o.SenderBrandIdDefault) {
		return true
	}

	return false
}

// SetSenderBrandIdDefault gets a reference to the given string and assigns it to the SenderBrandIdDefault field.
func (o *AccountBrands) SetSenderBrandIdDefault(v string) {
	o.SenderBrandIdDefault = &v
}

func (o AccountBrands) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBrands) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brands) {
		toSerialize["brands"] = o.Brands
	}
	if !IsNil(o.RecipientBrandIdDefault) {
		toSerialize["recipientBrandIdDefault"] = o.RecipientBrandIdDefault
	}
	if !IsNil(o.SenderBrandIdDefault) {
		toSerialize["senderBrandIdDefault"] = o.SenderBrandIdDefault
	}
	return toSerialize, nil
}

type NullableAccountBrands struct {
	value *AccountBrands
	isSet bool
}

func (v NullableAccountBrands) Get() *AccountBrands {
	return v.value
}

func (v *NullableAccountBrands) Set(val *AccountBrands) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBrands) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBrands) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBrands(val *AccountBrands) *NullableAccountBrands {
	return &NullableAccountBrands{value: val, isSet: true}
}

func (v NullableAccountBrands) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBrands) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


