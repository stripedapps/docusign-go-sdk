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

// checks if the EnvelopeTransactionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeTransactionStatus{}

// EnvelopeTransactionStatus 
type EnvelopeTransactionStatus struct {
	// The envelope ID of the envelope status that failed to post.
	EnvelopeId *string `json:"envelopeId,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// Indicates the envelope status. Valid values are:  * `completed`: The recipients have finished working with the envelope: the documents are signed and all required tabs are filled in. * `created`: The envelope is created as a draft. It can be modified and sent later. * `declined`: The envelope has been declined by the recipients. * `delivered`: The envelope has been delivered to the recipients. * `sent`: The envelope will be sent to the recipients after the envelope is created. * `signed`: The envelope has been signed by the recipients. * `voided`: The envelope is no longer valid and recipients cannot access or sign the envelope. 
	Status *string `json:"status,omitempty"`
	//  Used to identify an envelope. The ID is a sender-generated value and is valid in the DocuSign system for 7 days. It is recommended that a transaction ID is used for offline signing to ensure that an envelope is not sent multiple times. The `transactionId` property can be used determine an envelope's status (i.e. was it created or not) in cases where the internet connection was lost before the envelope status was returned.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewEnvelopeTransactionStatus instantiates a new EnvelopeTransactionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeTransactionStatus() *EnvelopeTransactionStatus {
	this := EnvelopeTransactionStatus{}
	return &this
}

// NewEnvelopeTransactionStatusWithDefaults instantiates a new EnvelopeTransactionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeTransactionStatusWithDefaults() *EnvelopeTransactionStatus {
	this := EnvelopeTransactionStatus{}
	return &this
}

// GetEnvelopeId returns the EnvelopeId field value if set, zero value otherwise.
func (o *EnvelopeTransactionStatus) GetEnvelopeId() string {
	if o == nil || IsNil(o.EnvelopeId) {
		var ret string
		return ret
	}
	return *o.EnvelopeId
}

// GetEnvelopeIdOk returns a tuple with the EnvelopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeTransactionStatus) GetEnvelopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeId) {
		return nil, false
	}
	return o.EnvelopeId, true
}

// HasEnvelopeId returns a boolean if a field has been set.
func (o *EnvelopeTransactionStatus) HasEnvelopeId() bool {
	if o != nil && !IsNil(o.EnvelopeId) {
		return true
	}

	return false
}

// SetEnvelopeId gets a reference to the given string and assigns it to the EnvelopeId field.
func (o *EnvelopeTransactionStatus) SetEnvelopeId(v string) {
	o.EnvelopeId = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *EnvelopeTransactionStatus) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeTransactionStatus) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *EnvelopeTransactionStatus) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *EnvelopeTransactionStatus) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EnvelopeTransactionStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeTransactionStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EnvelopeTransactionStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EnvelopeTransactionStatus) SetStatus(v string) {
	o.Status = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *EnvelopeTransactionStatus) GetTransactionId() string {
	if o == nil || IsNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeTransactionStatus) GetTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *EnvelopeTransactionStatus) HasTransactionId() bool {
	if o != nil && !IsNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *EnvelopeTransactionStatus) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o EnvelopeTransactionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeTransactionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvelopeId) {
		toSerialize["envelopeId"] = o.EnvelopeId
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	return toSerialize, nil
}

type NullableEnvelopeTransactionStatus struct {
	value *EnvelopeTransactionStatus
	isSet bool
}

func (v NullableEnvelopeTransactionStatus) Get() *EnvelopeTransactionStatus {
	return v.value
}

func (v *NullableEnvelopeTransactionStatus) Set(val *EnvelopeTransactionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeTransactionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeTransactionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeTransactionStatus(val *EnvelopeTransactionStatus) *NullableEnvelopeTransactionStatus {
	return &NullableEnvelopeTransactionStatus{value: val, isSet: true}
}

func (v NullableEnvelopeTransactionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeTransactionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


