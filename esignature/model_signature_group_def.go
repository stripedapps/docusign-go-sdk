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

// checks if the SignatureGroupDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureGroupDef{}

// SignatureGroupDef 
type SignatureGroupDef struct {
	// The ID of the group being accessed.
	GroupId *string `json:"groupId,omitempty"`
	// Indicates whether the property is editable. Valid values are:  - `editable` - `read_only`
	Rights *string `json:"rights,omitempty"`
}

// NewSignatureGroupDef instantiates a new SignatureGroupDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureGroupDef() *SignatureGroupDef {
	this := SignatureGroupDef{}
	return &this
}

// NewSignatureGroupDefWithDefaults instantiates a new SignatureGroupDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureGroupDefWithDefaults() *SignatureGroupDef {
	this := SignatureGroupDef{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SignatureGroupDef) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGroupDef) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SignatureGroupDef) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SignatureGroupDef) SetGroupId(v string) {
	o.GroupId = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *SignatureGroupDef) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGroupDef) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *SignatureGroupDef) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *SignatureGroupDef) SetRights(v string) {
	o.Rights = &v
}

func (o SignatureGroupDef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureGroupDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	return toSerialize, nil
}

type NullableSignatureGroupDef struct {
	value *SignatureGroupDef
	isSet bool
}

func (v NullableSignatureGroupDef) Get() *SignatureGroupDef {
	return v.value
}

func (v *NullableSignatureGroupDef) Set(val *SignatureGroupDef) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureGroupDef) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureGroupDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureGroupDef(val *SignatureGroupDef) *NullableSignatureGroupDef {
	return &NullableSignatureGroupDef{value: val, isSet: true}
}

func (v NullableSignatureGroupDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureGroupDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

