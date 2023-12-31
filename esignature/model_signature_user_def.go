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

// checks if the SignatureUserDef type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureUserDef{}

// SignatureUserDef 
type SignatureUserDef struct {
	// Boolean that specifies whether the signature is the default signature for the user.
	IsDefault *string `json:"isDefault,omitempty"`
	// Indicates whether the property is editable. Valid values are:  - `editable` - `read_only`
	Rights *string `json:"rights,omitempty"`
	// The ID of the user to access.  **Note:** Users can only access their own information. A user, even one with Admin rights, cannot access another user's settings.
	UserId *string `json:"userId,omitempty"`
}

// NewSignatureUserDef instantiates a new SignatureUserDef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureUserDef() *SignatureUserDef {
	this := SignatureUserDef{}
	return &this
}

// NewSignatureUserDefWithDefaults instantiates a new SignatureUserDef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureUserDefWithDefaults() *SignatureUserDef {
	this := SignatureUserDef{}
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *SignatureUserDef) GetIsDefault() string {
	if o == nil || IsNil(o.IsDefault) {
		var ret string
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureUserDef) GetIsDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *SignatureUserDef) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given string and assigns it to the IsDefault field.
func (o *SignatureUserDef) SetIsDefault(v string) {
	o.IsDefault = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *SignatureUserDef) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureUserDef) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *SignatureUserDef) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *SignatureUserDef) SetRights(v string) {
	o.Rights = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *SignatureUserDef) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureUserDef) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *SignatureUserDef) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *SignatureUserDef) SetUserId(v string) {
	o.UserId = &v
}

func (o SignatureUserDef) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureUserDef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableSignatureUserDef struct {
	value *SignatureUserDef
	isSet bool
}

func (v NullableSignatureUserDef) Get() *SignatureUserDef {
	return v.value
}

func (v *NullableSignatureUserDef) Set(val *SignatureUserDef) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureUserDef) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureUserDef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureUserDef(val *SignatureUserDef) *NullableSignatureUserDef {
	return &NullableSignatureUserDef{value: val, isSet: true}
}

func (v NullableSignatureUserDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureUserDef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


