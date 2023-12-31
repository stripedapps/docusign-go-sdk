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

// checks if the RecipientSignatureProviderOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientSignatureProviderOptions{}

// RecipientSignatureProviderOptions Option settings for the signature provider. Different providers require or use different options. [The current provider list and the options they require.](/docs/esign-rest-api/esign101/concepts/standards-based-signatures/)
type RecipientSignatureProviderOptions struct {
	// Reserved for DocuSign.
	CpfNumber *string `json:"cpfNumber,omitempty"`
	CpfNumberMetadata *PropertyMetadata `json:"cpfNumberMetadata,omitempty"`
	// A pre-shared secret that the signer must enter to complete the signing process. Eg last six digits of the signer's government ID or Social Security number. Or a newly created pre-shared secret for the transaction. Note: some signature providers may require an exact (case-sensitive) match if alphabetic characters are included in the field.
	OneTimePassword *string `json:"oneTimePassword,omitempty"`
	OneTimePasswordMetadata *PropertyMetadata `json:"oneTimePasswordMetadata,omitempty"`
	// The role or capacity of the signing recipient. Examples: Manager, Approver, etc.
	SignerRole *string `json:"signerRole,omitempty"`
	SignerRoleMetadata *PropertyMetadata `json:"signerRoleMetadata,omitempty"`
	// The mobile phone number used to send the recipient an access code for the signing ceremony. Format: a string starting with +, then the country code followed by the full mobile phone number without any spaces or special characters. Omit leading zeroes before a city code. Examples: +14155551234, +97235551234, +33505551234.
	Sms *string `json:"sms,omitempty"`
	SmsMetadata *PropertyMetadata `json:"smsMetadata,omitempty"`
}

// NewRecipientSignatureProviderOptions instantiates a new RecipientSignatureProviderOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientSignatureProviderOptions() *RecipientSignatureProviderOptions {
	this := RecipientSignatureProviderOptions{}
	return &this
}

// NewRecipientSignatureProviderOptionsWithDefaults instantiates a new RecipientSignatureProviderOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientSignatureProviderOptionsWithDefaults() *RecipientSignatureProviderOptions {
	this := RecipientSignatureProviderOptions{}
	return &this
}

// GetCpfNumber returns the CpfNumber field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetCpfNumber() string {
	if o == nil || IsNil(o.CpfNumber) {
		var ret string
		return ret
	}
	return *o.CpfNumber
}

// GetCpfNumberOk returns a tuple with the CpfNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetCpfNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CpfNumber) {
		return nil, false
	}
	return o.CpfNumber, true
}

// HasCpfNumber returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasCpfNumber() bool {
	if o != nil && !IsNil(o.CpfNumber) {
		return true
	}

	return false
}

// SetCpfNumber gets a reference to the given string and assigns it to the CpfNumber field.
func (o *RecipientSignatureProviderOptions) SetCpfNumber(v string) {
	o.CpfNumber = &v
}

// GetCpfNumberMetadata returns the CpfNumberMetadata field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetCpfNumberMetadata() PropertyMetadata {
	if o == nil || IsNil(o.CpfNumberMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.CpfNumberMetadata
}

// GetCpfNumberMetadataOk returns a tuple with the CpfNumberMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetCpfNumberMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.CpfNumberMetadata) {
		return nil, false
	}
	return o.CpfNumberMetadata, true
}

// HasCpfNumberMetadata returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasCpfNumberMetadata() bool {
	if o != nil && !IsNil(o.CpfNumberMetadata) {
		return true
	}

	return false
}

// SetCpfNumberMetadata gets a reference to the given PropertyMetadata and assigns it to the CpfNumberMetadata field.
func (o *RecipientSignatureProviderOptions) SetCpfNumberMetadata(v PropertyMetadata) {
	o.CpfNumberMetadata = &v
}

// GetOneTimePassword returns the OneTimePassword field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetOneTimePassword() string {
	if o == nil || IsNil(o.OneTimePassword) {
		var ret string
		return ret
	}
	return *o.OneTimePassword
}

// GetOneTimePasswordOk returns a tuple with the OneTimePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetOneTimePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OneTimePassword) {
		return nil, false
	}
	return o.OneTimePassword, true
}

// HasOneTimePassword returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasOneTimePassword() bool {
	if o != nil && !IsNil(o.OneTimePassword) {
		return true
	}

	return false
}

// SetOneTimePassword gets a reference to the given string and assigns it to the OneTimePassword field.
func (o *RecipientSignatureProviderOptions) SetOneTimePassword(v string) {
	o.OneTimePassword = &v
}

// GetOneTimePasswordMetadata returns the OneTimePasswordMetadata field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetOneTimePasswordMetadata() PropertyMetadata {
	if o == nil || IsNil(o.OneTimePasswordMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.OneTimePasswordMetadata
}

// GetOneTimePasswordMetadataOk returns a tuple with the OneTimePasswordMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetOneTimePasswordMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.OneTimePasswordMetadata) {
		return nil, false
	}
	return o.OneTimePasswordMetadata, true
}

// HasOneTimePasswordMetadata returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasOneTimePasswordMetadata() bool {
	if o != nil && !IsNil(o.OneTimePasswordMetadata) {
		return true
	}

	return false
}

// SetOneTimePasswordMetadata gets a reference to the given PropertyMetadata and assigns it to the OneTimePasswordMetadata field.
func (o *RecipientSignatureProviderOptions) SetOneTimePasswordMetadata(v PropertyMetadata) {
	o.OneTimePasswordMetadata = &v
}

// GetSignerRole returns the SignerRole field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetSignerRole() string {
	if o == nil || IsNil(o.SignerRole) {
		var ret string
		return ret
	}
	return *o.SignerRole
}

// GetSignerRoleOk returns a tuple with the SignerRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetSignerRoleOk() (*string, bool) {
	if o == nil || IsNil(o.SignerRole) {
		return nil, false
	}
	return o.SignerRole, true
}

// HasSignerRole returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasSignerRole() bool {
	if o != nil && !IsNil(o.SignerRole) {
		return true
	}

	return false
}

// SetSignerRole gets a reference to the given string and assigns it to the SignerRole field.
func (o *RecipientSignatureProviderOptions) SetSignerRole(v string) {
	o.SignerRole = &v
}

// GetSignerRoleMetadata returns the SignerRoleMetadata field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetSignerRoleMetadata() PropertyMetadata {
	if o == nil || IsNil(o.SignerRoleMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.SignerRoleMetadata
}

// GetSignerRoleMetadataOk returns a tuple with the SignerRoleMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetSignerRoleMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.SignerRoleMetadata) {
		return nil, false
	}
	return o.SignerRoleMetadata, true
}

// HasSignerRoleMetadata returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasSignerRoleMetadata() bool {
	if o != nil && !IsNil(o.SignerRoleMetadata) {
		return true
	}

	return false
}

// SetSignerRoleMetadata gets a reference to the given PropertyMetadata and assigns it to the SignerRoleMetadata field.
func (o *RecipientSignatureProviderOptions) SetSignerRoleMetadata(v PropertyMetadata) {
	o.SignerRoleMetadata = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetSms() string {
	if o == nil || IsNil(o.Sms) {
		var ret string
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetSmsOk() (*string, bool) {
	if o == nil || IsNil(o.Sms) {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasSms() bool {
	if o != nil && !IsNil(o.Sms) {
		return true
	}

	return false
}

// SetSms gets a reference to the given string and assigns it to the Sms field.
func (o *RecipientSignatureProviderOptions) SetSms(v string) {
	o.Sms = &v
}

// GetSmsMetadata returns the SmsMetadata field value if set, zero value otherwise.
func (o *RecipientSignatureProviderOptions) GetSmsMetadata() PropertyMetadata {
	if o == nil || IsNil(o.SmsMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.SmsMetadata
}

// GetSmsMetadataOk returns a tuple with the SmsMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProviderOptions) GetSmsMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.SmsMetadata) {
		return nil, false
	}
	return o.SmsMetadata, true
}

// HasSmsMetadata returns a boolean if a field has been set.
func (o *RecipientSignatureProviderOptions) HasSmsMetadata() bool {
	if o != nil && !IsNil(o.SmsMetadata) {
		return true
	}

	return false
}

// SetSmsMetadata gets a reference to the given PropertyMetadata and assigns it to the SmsMetadata field.
func (o *RecipientSignatureProviderOptions) SetSmsMetadata(v PropertyMetadata) {
	o.SmsMetadata = &v
}

func (o RecipientSignatureProviderOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientSignatureProviderOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CpfNumber) {
		toSerialize["cpfNumber"] = o.CpfNumber
	}
	if !IsNil(o.CpfNumberMetadata) {
		toSerialize["cpfNumberMetadata"] = o.CpfNumberMetadata
	}
	if !IsNil(o.OneTimePassword) {
		toSerialize["oneTimePassword"] = o.OneTimePassword
	}
	if !IsNil(o.OneTimePasswordMetadata) {
		toSerialize["oneTimePasswordMetadata"] = o.OneTimePasswordMetadata
	}
	if !IsNil(o.SignerRole) {
		toSerialize["signerRole"] = o.SignerRole
	}
	if !IsNil(o.SignerRoleMetadata) {
		toSerialize["signerRoleMetadata"] = o.SignerRoleMetadata
	}
	if !IsNil(o.Sms) {
		toSerialize["sms"] = o.Sms
	}
	if !IsNil(o.SmsMetadata) {
		toSerialize["smsMetadata"] = o.SmsMetadata
	}
	return toSerialize, nil
}

type NullableRecipientSignatureProviderOptions struct {
	value *RecipientSignatureProviderOptions
	isSet bool
}

func (v NullableRecipientSignatureProviderOptions) Get() *RecipientSignatureProviderOptions {
	return v.value
}

func (v *NullableRecipientSignatureProviderOptions) Set(val *RecipientSignatureProviderOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientSignatureProviderOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientSignatureProviderOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientSignatureProviderOptions(val *RecipientSignatureProviderOptions) *NullableRecipientSignatureProviderOptions {
	return &NullableRecipientSignatureProviderOptions{value: val, isSet: true}
}

func (v NullableRecipientSignatureProviderOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientSignatureProviderOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


