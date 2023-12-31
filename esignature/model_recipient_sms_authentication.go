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

// checks if the RecipientSMSAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientSMSAuthentication{}

// RecipientSMSAuthentication Contains the element senderProvidedNumbers which is an Array  of phone numbers the recipient can use for SMS text authentication.
type RecipientSMSAuthentication struct {
	// An array containing a list of phone numbers that the recipient can use for SMS text authentication. 
	SenderProvidedNumbers []string `json:"senderProvidedNumbers,omitempty"`
	SenderProvidedNumbersMetadata *PropertyMetadata `json:"senderProvidedNumbersMetadata,omitempty"`
}

// NewRecipientSMSAuthentication instantiates a new RecipientSMSAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientSMSAuthentication() *RecipientSMSAuthentication {
	this := RecipientSMSAuthentication{}
	return &this
}

// NewRecipientSMSAuthenticationWithDefaults instantiates a new RecipientSMSAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientSMSAuthenticationWithDefaults() *RecipientSMSAuthentication {
	this := RecipientSMSAuthentication{}
	return &this
}

// GetSenderProvidedNumbers returns the SenderProvidedNumbers field value if set, zero value otherwise.
func (o *RecipientSMSAuthentication) GetSenderProvidedNumbers() []string {
	if o == nil || IsNil(o.SenderProvidedNumbers) {
		var ret []string
		return ret
	}
	return o.SenderProvidedNumbers
}

// GetSenderProvidedNumbersOk returns a tuple with the SenderProvidedNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSMSAuthentication) GetSenderProvidedNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.SenderProvidedNumbers) {
		return nil, false
	}
	return o.SenderProvidedNumbers, true
}

// HasSenderProvidedNumbers returns a boolean if a field has been set.
func (o *RecipientSMSAuthentication) HasSenderProvidedNumbers() bool {
	if o != nil && !IsNil(o.SenderProvidedNumbers) {
		return true
	}

	return false
}

// SetSenderProvidedNumbers gets a reference to the given []string and assigns it to the SenderProvidedNumbers field.
func (o *RecipientSMSAuthentication) SetSenderProvidedNumbers(v []string) {
	o.SenderProvidedNumbers = v
}

// GetSenderProvidedNumbersMetadata returns the SenderProvidedNumbersMetadata field value if set, zero value otherwise.
func (o *RecipientSMSAuthentication) GetSenderProvidedNumbersMetadata() PropertyMetadata {
	if o == nil || IsNil(o.SenderProvidedNumbersMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.SenderProvidedNumbersMetadata
}

// GetSenderProvidedNumbersMetadataOk returns a tuple with the SenderProvidedNumbersMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSMSAuthentication) GetSenderProvidedNumbersMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.SenderProvidedNumbersMetadata) {
		return nil, false
	}
	return o.SenderProvidedNumbersMetadata, true
}

// HasSenderProvidedNumbersMetadata returns a boolean if a field has been set.
func (o *RecipientSMSAuthentication) HasSenderProvidedNumbersMetadata() bool {
	if o != nil && !IsNil(o.SenderProvidedNumbersMetadata) {
		return true
	}

	return false
}

// SetSenderProvidedNumbersMetadata gets a reference to the given PropertyMetadata and assigns it to the SenderProvidedNumbersMetadata field.
func (o *RecipientSMSAuthentication) SetSenderProvidedNumbersMetadata(v PropertyMetadata) {
	o.SenderProvidedNumbersMetadata = &v
}

func (o RecipientSMSAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientSMSAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SenderProvidedNumbers) {
		toSerialize["senderProvidedNumbers"] = o.SenderProvidedNumbers
	}
	if !IsNil(o.SenderProvidedNumbersMetadata) {
		toSerialize["senderProvidedNumbersMetadata"] = o.SenderProvidedNumbersMetadata
	}
	return toSerialize, nil
}

type NullableRecipientSMSAuthentication struct {
	value *RecipientSMSAuthentication
	isSet bool
}

func (v NullableRecipientSMSAuthentication) Get() *RecipientSMSAuthentication {
	return v.value
}

func (v *NullableRecipientSMSAuthentication) Set(val *RecipientSMSAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientSMSAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientSMSAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientSMSAuthentication(val *RecipientSMSAuthentication) *NullableRecipientSMSAuthentication {
	return &NullableRecipientSMSAuthentication{value: val, isSet: true}
}

func (v NullableRecipientSMSAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientSMSAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


