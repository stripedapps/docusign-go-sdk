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

// checks if the BrandResourceUrls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandResourceUrls{}

// BrandResourceUrls Brands use resource files to style the following experiences:   - Email - Sending - Signing - Captive (embedded) signing   You can modify these resource files to customize these experiences.
type BrandResourceUrls struct {
	// The URI for the email resource file that the brand uses.
	Email *string `json:"email,omitempty"`
	// The URI for the sending resource file that the brand uses.
	Sending *string `json:"sending,omitempty"`
	// The URI for the signing resource file that the brand uses.
	Signing *string `json:"signing,omitempty"`
	// The URI for the captive (embedded) signing resource file that the brand uses.
	SigningCaptive *string `json:"signingCaptive,omitempty"`
}

// NewBrandResourceUrls instantiates a new BrandResourceUrls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandResourceUrls() *BrandResourceUrls {
	this := BrandResourceUrls{}
	return &this
}

// NewBrandResourceUrlsWithDefaults instantiates a new BrandResourceUrls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandResourceUrlsWithDefaults() *BrandResourceUrls {
	this := BrandResourceUrls{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BrandResourceUrls) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandResourceUrls) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BrandResourceUrls) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BrandResourceUrls) SetEmail(v string) {
	o.Email = &v
}

// GetSending returns the Sending field value if set, zero value otherwise.
func (o *BrandResourceUrls) GetSending() string {
	if o == nil || IsNil(o.Sending) {
		var ret string
		return ret
	}
	return *o.Sending
}

// GetSendingOk returns a tuple with the Sending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandResourceUrls) GetSendingOk() (*string, bool) {
	if o == nil || IsNil(o.Sending) {
		return nil, false
	}
	return o.Sending, true
}

// HasSending returns a boolean if a field has been set.
func (o *BrandResourceUrls) HasSending() bool {
	if o != nil && !IsNil(o.Sending) {
		return true
	}

	return false
}

// SetSending gets a reference to the given string and assigns it to the Sending field.
func (o *BrandResourceUrls) SetSending(v string) {
	o.Sending = &v
}

// GetSigning returns the Signing field value if set, zero value otherwise.
func (o *BrandResourceUrls) GetSigning() string {
	if o == nil || IsNil(o.Signing) {
		var ret string
		return ret
	}
	return *o.Signing
}

// GetSigningOk returns a tuple with the Signing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandResourceUrls) GetSigningOk() (*string, bool) {
	if o == nil || IsNil(o.Signing) {
		return nil, false
	}
	return o.Signing, true
}

// HasSigning returns a boolean if a field has been set.
func (o *BrandResourceUrls) HasSigning() bool {
	if o != nil && !IsNil(o.Signing) {
		return true
	}

	return false
}

// SetSigning gets a reference to the given string and assigns it to the Signing field.
func (o *BrandResourceUrls) SetSigning(v string) {
	o.Signing = &v
}

// GetSigningCaptive returns the SigningCaptive field value if set, zero value otherwise.
func (o *BrandResourceUrls) GetSigningCaptive() string {
	if o == nil || IsNil(o.SigningCaptive) {
		var ret string
		return ret
	}
	return *o.SigningCaptive
}

// GetSigningCaptiveOk returns a tuple with the SigningCaptive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandResourceUrls) GetSigningCaptiveOk() (*string, bool) {
	if o == nil || IsNil(o.SigningCaptive) {
		return nil, false
	}
	return o.SigningCaptive, true
}

// HasSigningCaptive returns a boolean if a field has been set.
func (o *BrandResourceUrls) HasSigningCaptive() bool {
	if o != nil && !IsNil(o.SigningCaptive) {
		return true
	}

	return false
}

// SetSigningCaptive gets a reference to the given string and assigns it to the SigningCaptive field.
func (o *BrandResourceUrls) SetSigningCaptive(v string) {
	o.SigningCaptive = &v
}

func (o BrandResourceUrls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandResourceUrls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Sending) {
		toSerialize["sending"] = o.Sending
	}
	if !IsNil(o.Signing) {
		toSerialize["signing"] = o.Signing
	}
	if !IsNil(o.SigningCaptive) {
		toSerialize["signingCaptive"] = o.SigningCaptive
	}
	return toSerialize, nil
}

type NullableBrandResourceUrls struct {
	value *BrandResourceUrls
	isSet bool
}

func (v NullableBrandResourceUrls) Get() *BrandResourceUrls {
	return v.value
}

func (v *NullableBrandResourceUrls) Set(val *BrandResourceUrls) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandResourceUrls) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandResourceUrls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandResourceUrls(val *BrandResourceUrls) *NullableBrandResourceUrls {
	return &NullableBrandResourceUrls{value: val, isSet: true}
}

func (v NullableBrandResourceUrls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandResourceUrls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

