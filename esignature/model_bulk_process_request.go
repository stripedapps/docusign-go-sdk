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

// checks if the BulkProcessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkProcessRequest{}

// BulkProcessRequest 
type BulkProcessRequest struct {
	// 
	BatchName *string `json:"batchName,omitempty"`
	// The GUID of the envelope or template.
	EnvelopeOrTemplateId *string `json:"envelopeOrTemplateId,omitempty"`
}

// NewBulkProcessRequest instantiates a new BulkProcessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkProcessRequest() *BulkProcessRequest {
	this := BulkProcessRequest{}
	return &this
}

// NewBulkProcessRequestWithDefaults instantiates a new BulkProcessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkProcessRequestWithDefaults() *BulkProcessRequest {
	this := BulkProcessRequest{}
	return &this
}

// GetBatchName returns the BatchName field value if set, zero value otherwise.
func (o *BulkProcessRequest) GetBatchName() string {
	if o == nil || IsNil(o.BatchName) {
		var ret string
		return ret
	}
	return *o.BatchName
}

// GetBatchNameOk returns a tuple with the BatchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkProcessRequest) GetBatchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BatchName) {
		return nil, false
	}
	return o.BatchName, true
}

// HasBatchName returns a boolean if a field has been set.
func (o *BulkProcessRequest) HasBatchName() bool {
	if o != nil && !IsNil(o.BatchName) {
		return true
	}

	return false
}

// SetBatchName gets a reference to the given string and assigns it to the BatchName field.
func (o *BulkProcessRequest) SetBatchName(v string) {
	o.BatchName = &v
}

// GetEnvelopeOrTemplateId returns the EnvelopeOrTemplateId field value if set, zero value otherwise.
func (o *BulkProcessRequest) GetEnvelopeOrTemplateId() string {
	if o == nil || IsNil(o.EnvelopeOrTemplateId) {
		var ret string
		return ret
	}
	return *o.EnvelopeOrTemplateId
}

// GetEnvelopeOrTemplateIdOk returns a tuple with the EnvelopeOrTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkProcessRequest) GetEnvelopeOrTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeOrTemplateId) {
		return nil, false
	}
	return o.EnvelopeOrTemplateId, true
}

// HasEnvelopeOrTemplateId returns a boolean if a field has been set.
func (o *BulkProcessRequest) HasEnvelopeOrTemplateId() bool {
	if o != nil && !IsNil(o.EnvelopeOrTemplateId) {
		return true
	}

	return false
}

// SetEnvelopeOrTemplateId gets a reference to the given string and assigns it to the EnvelopeOrTemplateId field.
func (o *BulkProcessRequest) SetEnvelopeOrTemplateId(v string) {
	o.EnvelopeOrTemplateId = &v
}

func (o BulkProcessRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkProcessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchName) {
		toSerialize["batchName"] = o.BatchName
	}
	if !IsNil(o.EnvelopeOrTemplateId) {
		toSerialize["envelopeOrTemplateId"] = o.EnvelopeOrTemplateId
	}
	return toSerialize, nil
}

type NullableBulkProcessRequest struct {
	value *BulkProcessRequest
	isSet bool
}

func (v NullableBulkProcessRequest) Get() *BulkProcessRequest {
	return v.value
}

func (v *NullableBulkProcessRequest) Set(val *BulkProcessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkProcessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkProcessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkProcessRequest(val *BulkProcessRequest) *NullableBulkProcessRequest {
	return &NullableBulkProcessRequest{value: val, isSet: true}
}

func (v NullableBulkProcessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkProcessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


