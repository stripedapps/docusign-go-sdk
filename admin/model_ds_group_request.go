/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the DSGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSGroupRequest{}

// DSGroupRequest 
type DSGroupRequest struct {
	// 
	DsGroupId string `json:"ds_group_id"`
}

// NewDSGroupRequest instantiates a new DSGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSGroupRequest(dsGroupId string) *DSGroupRequest {
	this := DSGroupRequest{}
	this.DsGroupId = dsGroupId
	return &this
}

// NewDSGroupRequestWithDefaults instantiates a new DSGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSGroupRequestWithDefaults() *DSGroupRequest {
	this := DSGroupRequest{}
	return &this
}

// GetDsGroupId returns the DsGroupId field value
func (o *DSGroupRequest) GetDsGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DsGroupId
}

// GetDsGroupIdOk returns a tuple with the DsGroupId field value
// and a boolean to check if the value has been set.
func (o *DSGroupRequest) GetDsGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsGroupId, true
}

// SetDsGroupId sets field value
func (o *DSGroupRequest) SetDsGroupId(v string) {
	o.DsGroupId = v
}

func (o DSGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ds_group_id"] = o.DsGroupId
	return toSerialize, nil
}

type NullableDSGroupRequest struct {
	value *DSGroupRequest
	isSet bool
}

func (v NullableDSGroupRequest) Get() *DSGroupRequest {
	return v.value
}

func (v *NullableDSGroupRequest) Set(val *DSGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDSGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDSGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSGroupRequest(val *DSGroupRequest) *NullableDSGroupRequest {
	return &NullableDSGroupRequest{value: val, isSet: true}
}

func (v NullableDSGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


