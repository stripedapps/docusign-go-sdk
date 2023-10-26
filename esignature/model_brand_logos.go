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

// checks if the BrandLogos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandLogos{}

// BrandLogos The URIs for retrieving the logos that are associated with the brand.  These are read-only properties that provide a URI to logos in use. To update a logo use [AccountBrands: updateLogo](/docs/esign-rest-api/reference/accounts/accountbrands/updatelogo/). 
type BrandLogos struct {
	// The URI for the brand's secondary logo.  This is a read-only property that provides a URI to the logo in use. To update a logo use [AccountBrands: updateLogo](/docs/esign-rest-api/reference/accounts/accountbrands/updatelogo/). 
	Email *string `json:"email,omitempty"`
	// The URI for the brand's secondary logo.  This is a read-only property that provides a URI to the logo in use. To update a logo use [AccountBrands: updateLogo](/docs/esign-rest-api/reference/accounts/accountbrands/updatelogo/). 
	Primary *string `json:"primary,omitempty"`
	// The URI for the brand's secondary logo.  This is a read-only property that provides a URI to the logo in use. To update a logo use [AccountBrands: updateLogo](/docs/esign-rest-api/reference/accounts/accountbrands/updatelogo/). 
	Secondary *string `json:"secondary,omitempty"`
}

// NewBrandLogos instantiates a new BrandLogos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandLogos() *BrandLogos {
	this := BrandLogos{}
	return &this
}

// NewBrandLogosWithDefaults instantiates a new BrandLogos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandLogosWithDefaults() *BrandLogos {
	this := BrandLogos{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BrandLogos) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogos) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BrandLogos) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BrandLogos) SetEmail(v string) {
	o.Email = &v
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *BrandLogos) GetPrimary() string {
	if o == nil || IsNil(o.Primary) {
		var ret string
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogos) GetPrimaryOk() (*string, bool) {
	if o == nil || IsNil(o.Primary) {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *BrandLogos) HasPrimary() bool {
	if o != nil && !IsNil(o.Primary) {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given string and assigns it to the Primary field.
func (o *BrandLogos) SetPrimary(v string) {
	o.Primary = &v
}

// GetSecondary returns the Secondary field value if set, zero value otherwise.
func (o *BrandLogos) GetSecondary() string {
	if o == nil || IsNil(o.Secondary) {
		var ret string
		return ret
	}
	return *o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandLogos) GetSecondaryOk() (*string, bool) {
	if o == nil || IsNil(o.Secondary) {
		return nil, false
	}
	return o.Secondary, true
}

// HasSecondary returns a boolean if a field has been set.
func (o *BrandLogos) HasSecondary() bool {
	if o != nil && !IsNil(o.Secondary) {
		return true
	}

	return false
}

// SetSecondary gets a reference to the given string and assigns it to the Secondary field.
func (o *BrandLogos) SetSecondary(v string) {
	o.Secondary = &v
}

func (o BrandLogos) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandLogos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Primary) {
		toSerialize["primary"] = o.Primary
	}
	if !IsNil(o.Secondary) {
		toSerialize["secondary"] = o.Secondary
	}
	return toSerialize, nil
}

type NullableBrandLogos struct {
	value *BrandLogos
	isSet bool
}

func (v NullableBrandLogos) Get() *BrandLogos {
	return v.value
}

func (v *NullableBrandLogos) Set(val *BrandLogos) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandLogos) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandLogos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandLogos(val *BrandLogos) *NullableBrandLogos {
	return &NullableBrandLogos{value: val, isSet: true}
}

func (v NullableBrandLogos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandLogos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


