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

// checks if the PowerFormData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFormData{}

// PowerFormData Data that recipients have entered in PowerForm fields.
type PowerFormData struct {
	// The envelope ID of the envelope status that failed to post.
	EnvelopeId *string `json:"envelopeId,omitempty"`
	// An array of powerform recipients.
	Recipients []PowerFormFormDataRecipient `json:"recipients,omitempty"`
}

// NewPowerFormData instantiates a new PowerFormData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFormData() *PowerFormData {
	this := PowerFormData{}
	return &this
}

// NewPowerFormDataWithDefaults instantiates a new PowerFormData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFormDataWithDefaults() *PowerFormData {
	this := PowerFormData{}
	return &this
}

// GetEnvelopeId returns the EnvelopeId field value if set, zero value otherwise.
func (o *PowerFormData) GetEnvelopeId() string {
	if o == nil || IsNil(o.EnvelopeId) {
		var ret string
		return ret
	}
	return *o.EnvelopeId
}

// GetEnvelopeIdOk returns a tuple with the EnvelopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormData) GetEnvelopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeId) {
		return nil, false
	}
	return o.EnvelopeId, true
}

// HasEnvelopeId returns a boolean if a field has been set.
func (o *PowerFormData) HasEnvelopeId() bool {
	if o != nil && !IsNil(o.EnvelopeId) {
		return true
	}

	return false
}

// SetEnvelopeId gets a reference to the given string and assigns it to the EnvelopeId field.
func (o *PowerFormData) SetEnvelopeId(v string) {
	o.EnvelopeId = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *PowerFormData) GetRecipients() []PowerFormFormDataRecipient {
	if o == nil || IsNil(o.Recipients) {
		var ret []PowerFormFormDataRecipient
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormData) GetRecipientsOk() ([]PowerFormFormDataRecipient, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *PowerFormData) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []PowerFormFormDataRecipient and assigns it to the Recipients field.
func (o *PowerFormData) SetRecipients(v []PowerFormFormDataRecipient) {
	o.Recipients = v
}

func (o PowerFormData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFormData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvelopeId) {
		toSerialize["envelopeId"] = o.EnvelopeId
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	return toSerialize, nil
}

type NullablePowerFormData struct {
	value *PowerFormData
	isSet bool
}

func (v NullablePowerFormData) Get() *PowerFormData {
	return v.value
}

func (v *NullablePowerFormData) Set(val *PowerFormData) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFormData) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFormData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFormData(val *PowerFormData) *NullablePowerFormData {
	return &NullablePowerFormData{value: val, isSet: true}
}

func (v NullablePowerFormData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFormData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


