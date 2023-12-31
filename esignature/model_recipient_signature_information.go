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

// checks if the RecipientSignatureInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientSignatureInformation{}

// RecipientSignatureInformation Allows the sender to pre-specify the signature name, signature initials and signature font used in the signature stamp for the recipient.  Used only with recipient types In Person Signers and Signers.
type RecipientSignatureInformation struct {
	// The font type to use for the signature if the signature is not drawn. The following font styles  are supported. The quotes are to indicate that these values are strings, not `enums`.  - `\"1_DocuSign\"` - `\"2_DocuSign\"` - `\"3_DocuSign\"` - `\"4_DocuSign\"` - `\"5_DocuSign\"` - `\"6_DocuSign\"` - `\"7_DocuSign\"` - `\"8_DocuSign\"` - `\"Mistral\"` - `\"Rage Italic\"` 
	FontStyle *string `json:"fontStyle,omitempty"`
	// Specifies the user's signature in initials format.
	SignatureInitials *string `json:"signatureInitials,omitempty"`
	// Specifies the user's signature name.
	SignatureName *string `json:"signatureName,omitempty"`
}

// NewRecipientSignatureInformation instantiates a new RecipientSignatureInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientSignatureInformation() *RecipientSignatureInformation {
	this := RecipientSignatureInformation{}
	return &this
}

// NewRecipientSignatureInformationWithDefaults instantiates a new RecipientSignatureInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientSignatureInformationWithDefaults() *RecipientSignatureInformation {
	this := RecipientSignatureInformation{}
	return &this
}

// GetFontStyle returns the FontStyle field value if set, zero value otherwise.
func (o *RecipientSignatureInformation) GetFontStyle() string {
	if o == nil || IsNil(o.FontStyle) {
		var ret string
		return ret
	}
	return *o.FontStyle
}

// GetFontStyleOk returns a tuple with the FontStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureInformation) GetFontStyleOk() (*string, bool) {
	if o == nil || IsNil(o.FontStyle) {
		return nil, false
	}
	return o.FontStyle, true
}

// HasFontStyle returns a boolean if a field has been set.
func (o *RecipientSignatureInformation) HasFontStyle() bool {
	if o != nil && !IsNil(o.FontStyle) {
		return true
	}

	return false
}

// SetFontStyle gets a reference to the given string and assigns it to the FontStyle field.
func (o *RecipientSignatureInformation) SetFontStyle(v string) {
	o.FontStyle = &v
}

// GetSignatureInitials returns the SignatureInitials field value if set, zero value otherwise.
func (o *RecipientSignatureInformation) GetSignatureInitials() string {
	if o == nil || IsNil(o.SignatureInitials) {
		var ret string
		return ret
	}
	return *o.SignatureInitials
}

// GetSignatureInitialsOk returns a tuple with the SignatureInitials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureInformation) GetSignatureInitialsOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureInitials) {
		return nil, false
	}
	return o.SignatureInitials, true
}

// HasSignatureInitials returns a boolean if a field has been set.
func (o *RecipientSignatureInformation) HasSignatureInitials() bool {
	if o != nil && !IsNil(o.SignatureInitials) {
		return true
	}

	return false
}

// SetSignatureInitials gets a reference to the given string and assigns it to the SignatureInitials field.
func (o *RecipientSignatureInformation) SetSignatureInitials(v string) {
	o.SignatureInitials = &v
}

// GetSignatureName returns the SignatureName field value if set, zero value otherwise.
func (o *RecipientSignatureInformation) GetSignatureName() string {
	if o == nil || IsNil(o.SignatureName) {
		var ret string
		return ret
	}
	return *o.SignatureName
}

// GetSignatureNameOk returns a tuple with the SignatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureInformation) GetSignatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureName) {
		return nil, false
	}
	return o.SignatureName, true
}

// HasSignatureName returns a boolean if a field has been set.
func (o *RecipientSignatureInformation) HasSignatureName() bool {
	if o != nil && !IsNil(o.SignatureName) {
		return true
	}

	return false
}

// SetSignatureName gets a reference to the given string and assigns it to the SignatureName field.
func (o *RecipientSignatureInformation) SetSignatureName(v string) {
	o.SignatureName = &v
}

func (o RecipientSignatureInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientSignatureInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FontStyle) {
		toSerialize["fontStyle"] = o.FontStyle
	}
	if !IsNil(o.SignatureInitials) {
		toSerialize["signatureInitials"] = o.SignatureInitials
	}
	if !IsNil(o.SignatureName) {
		toSerialize["signatureName"] = o.SignatureName
	}
	return toSerialize, nil
}

type NullableRecipientSignatureInformation struct {
	value *RecipientSignatureInformation
	isSet bool
}

func (v NullableRecipientSignatureInformation) Get() *RecipientSignatureInformation {
	return v.value
}

func (v *NullableRecipientSignatureInformation) Set(val *RecipientSignatureInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientSignatureInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientSignatureInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientSignatureInformation(val *RecipientSignatureInformation) *NullableRecipientSignatureInformation {
	return &NullableRecipientSignatureInformation{value: val, isSet: true}
}

func (v NullableRecipientSignatureInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientSignatureInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


