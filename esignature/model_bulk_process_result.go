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

// checks if the BulkProcessResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkProcessResult{}

// BulkProcessResult 
type BulkProcessResult struct {
	// 
	Errors []BulkSendBatchError `json:"errors,omitempty"`
	// The GUID of the bulk send list.
	ListId *string `json:"listId,omitempty"`
	// 
	Success *string `json:"success,omitempty"`
}

// NewBulkProcessResult instantiates a new BulkProcessResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkProcessResult() *BulkProcessResult {
	this := BulkProcessResult{}
	return &this
}

// NewBulkProcessResultWithDefaults instantiates a new BulkProcessResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkProcessResultWithDefaults() *BulkProcessResult {
	this := BulkProcessResult{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *BulkProcessResult) GetErrors() []BulkSendBatchError {
	if o == nil || IsNil(o.Errors) {
		var ret []BulkSendBatchError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkProcessResult) GetErrorsOk() ([]BulkSendBatchError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *BulkProcessResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []BulkSendBatchError and assigns it to the Errors field.
func (o *BulkProcessResult) SetErrors(v []BulkSendBatchError) {
	o.Errors = v
}

// GetListId returns the ListId field value if set, zero value otherwise.
func (o *BulkProcessResult) GetListId() string {
	if o == nil || IsNil(o.ListId) {
		var ret string
		return ret
	}
	return *o.ListId
}

// GetListIdOk returns a tuple with the ListId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkProcessResult) GetListIdOk() (*string, bool) {
	if o == nil || IsNil(o.ListId) {
		return nil, false
	}
	return o.ListId, true
}

// HasListId returns a boolean if a field has been set.
func (o *BulkProcessResult) HasListId() bool {
	if o != nil && !IsNil(o.ListId) {
		return true
	}

	return false
}

// SetListId gets a reference to the given string and assigns it to the ListId field.
func (o *BulkProcessResult) SetListId(v string) {
	o.ListId = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *BulkProcessResult) GetSuccess() string {
	if o == nil || IsNil(o.Success) {
		var ret string
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkProcessResult) GetSuccessOk() (*string, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *BulkProcessResult) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given string and assigns it to the Success field.
func (o *BulkProcessResult) SetSuccess(v string) {
	o.Success = &v
}

func (o BulkProcessResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkProcessResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.ListId) {
		toSerialize["listId"] = o.ListId
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	return toSerialize, nil
}

type NullableBulkProcessResult struct {
	value *BulkProcessResult
	isSet bool
}

func (v NullableBulkProcessResult) Get() *BulkProcessResult {
	return v.value
}

func (v *NullableBulkProcessResult) Set(val *BulkProcessResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkProcessResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkProcessResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkProcessResult(val *BulkProcessResult) *NullableBulkProcessResult {
	return &NullableBulkProcessResult{value: val, isSet: true}
}

func (v NullableBulkProcessResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkProcessResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


