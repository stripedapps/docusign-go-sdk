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

// checks if the ConnectFailureFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectFailureFilter{}

// ConnectFailureFilter A list of failed envelope IDs to retry.
type ConnectFailureFilter struct {
	// An array of envelope GUIDs.  Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
	EnvelopeIds []string `json:"envelopeIds,omitempty"`
	// Must be **false.** Setting this property to any other value will result in errors.
	Synchronous *string `json:"synchronous,omitempty"`
}

// NewConnectFailureFilter instantiates a new ConnectFailureFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectFailureFilter() *ConnectFailureFilter {
	this := ConnectFailureFilter{}
	return &this
}

// NewConnectFailureFilterWithDefaults instantiates a new ConnectFailureFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectFailureFilterWithDefaults() *ConnectFailureFilter {
	this := ConnectFailureFilter{}
	return &this
}

// GetEnvelopeIds returns the EnvelopeIds field value if set, zero value otherwise.
func (o *ConnectFailureFilter) GetEnvelopeIds() []string {
	if o == nil || IsNil(o.EnvelopeIds) {
		var ret []string
		return ret
	}
	return o.EnvelopeIds
}

// GetEnvelopeIdsOk returns a tuple with the EnvelopeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectFailureFilter) GetEnvelopeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.EnvelopeIds) {
		return nil, false
	}
	return o.EnvelopeIds, true
}

// HasEnvelopeIds returns a boolean if a field has been set.
func (o *ConnectFailureFilter) HasEnvelopeIds() bool {
	if o != nil && !IsNil(o.EnvelopeIds) {
		return true
	}

	return false
}

// SetEnvelopeIds gets a reference to the given []string and assigns it to the EnvelopeIds field.
func (o *ConnectFailureFilter) SetEnvelopeIds(v []string) {
	o.EnvelopeIds = v
}

// GetSynchronous returns the Synchronous field value if set, zero value otherwise.
func (o *ConnectFailureFilter) GetSynchronous() string {
	if o == nil || IsNil(o.Synchronous) {
		var ret string
		return ret
	}
	return *o.Synchronous
}

// GetSynchronousOk returns a tuple with the Synchronous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectFailureFilter) GetSynchronousOk() (*string, bool) {
	if o == nil || IsNil(o.Synchronous) {
		return nil, false
	}
	return o.Synchronous, true
}

// HasSynchronous returns a boolean if a field has been set.
func (o *ConnectFailureFilter) HasSynchronous() bool {
	if o != nil && !IsNil(o.Synchronous) {
		return true
	}

	return false
}

// SetSynchronous gets a reference to the given string and assigns it to the Synchronous field.
func (o *ConnectFailureFilter) SetSynchronous(v string) {
	o.Synchronous = &v
}

func (o ConnectFailureFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectFailureFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnvelopeIds) {
		toSerialize["envelopeIds"] = o.EnvelopeIds
	}
	if !IsNil(o.Synchronous) {
		toSerialize["synchronous"] = o.Synchronous
	}
	return toSerialize, nil
}

type NullableConnectFailureFilter struct {
	value *ConnectFailureFilter
	isSet bool
}

func (v NullableConnectFailureFilter) Get() *ConnectFailureFilter {
	return v.value
}

func (v *NullableConnectFailureFilter) Set(val *ConnectFailureFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectFailureFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectFailureFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectFailureFilter(val *ConnectFailureFilter) *NullableConnectFailureFilter {
	return &NullableConnectFailureFilter{value: val, isSet: true}
}

func (v NullableConnectFailureFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectFailureFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


