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

// checks if the EnvelopePublishTransactionErrorRollup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopePublishTransactionErrorRollup{}

// EnvelopePublishTransactionErrorRollup 
type EnvelopePublishTransactionErrorRollup struct {
	// The maximum number of results to return.
	Count *string `json:"count,omitempty"`
	// 
	ErrorType *string `json:"errorType,omitempty"`
}

// NewEnvelopePublishTransactionErrorRollup instantiates a new EnvelopePublishTransactionErrorRollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopePublishTransactionErrorRollup() *EnvelopePublishTransactionErrorRollup {
	this := EnvelopePublishTransactionErrorRollup{}
	return &this
}

// NewEnvelopePublishTransactionErrorRollupWithDefaults instantiates a new EnvelopePublishTransactionErrorRollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopePublishTransactionErrorRollupWithDefaults() *EnvelopePublishTransactionErrorRollup {
	this := EnvelopePublishTransactionErrorRollup{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EnvelopePublishTransactionErrorRollup) GetCount() string {
	if o == nil || IsNil(o.Count) {
		var ret string
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransactionErrorRollup) GetCountOk() (*string, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransactionErrorRollup) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given string and assigns it to the Count field.
func (o *EnvelopePublishTransactionErrorRollup) SetCount(v string) {
	o.Count = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *EnvelopePublishTransactionErrorRollup) GetErrorType() string {
	if o == nil || IsNil(o.ErrorType) {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransactionErrorRollup) GetErrorTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *EnvelopePublishTransactionErrorRollup) HasErrorType() bool {
	if o != nil && !IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *EnvelopePublishTransactionErrorRollup) SetErrorType(v string) {
	o.ErrorType = &v
}

func (o EnvelopePublishTransactionErrorRollup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopePublishTransactionErrorRollup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	return toSerialize, nil
}

type NullableEnvelopePublishTransactionErrorRollup struct {
	value *EnvelopePublishTransactionErrorRollup
	isSet bool
}

func (v NullableEnvelopePublishTransactionErrorRollup) Get() *EnvelopePublishTransactionErrorRollup {
	return v.value
}

func (v *NullableEnvelopePublishTransactionErrorRollup) Set(val *EnvelopePublishTransactionErrorRollup) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopePublishTransactionErrorRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopePublishTransactionErrorRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopePublishTransactionErrorRollup(val *EnvelopePublishTransactionErrorRollup) *NullableEnvelopePublishTransactionErrorRollup {
	return &NullableEnvelopePublishTransactionErrorRollup{value: val, isSet: true}
}

func (v NullableEnvelopePublishTransactionErrorRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopePublishTransactionErrorRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


