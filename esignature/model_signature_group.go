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

// checks if the SignatureGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureGroup{}

// SignatureGroup 
type SignatureGroup struct {
	// The ID of the group being accessed.
	GroupId *string `json:"groupId,omitempty"`
	// The name of the group. The search_text provided in the call automatically performs a wild card search on group_name.
	GroupName *string `json:"groupName,omitempty"`
	// Indicates whether the property is editable. Valid values are:  - `editable` - `read_only`
	Rights *string `json:"rights,omitempty"`
}

// NewSignatureGroup instantiates a new SignatureGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureGroup() *SignatureGroup {
	this := SignatureGroup{}
	return &this
}

// NewSignatureGroupWithDefaults instantiates a new SignatureGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureGroupWithDefaults() *SignatureGroup {
	this := SignatureGroup{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SignatureGroup) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SignatureGroup) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SignatureGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *SignatureGroup) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGroup) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *SignatureGroup) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *SignatureGroup) SetGroupName(v string) {
	o.GroupName = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *SignatureGroup) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureGroup) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *SignatureGroup) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *SignatureGroup) SetRights(v string) {
	o.Rights = &v
}

func (o SignatureGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	return toSerialize, nil
}

type NullableSignatureGroup struct {
	value *SignatureGroup
	isSet bool
}

func (v NullableSignatureGroup) Get() *SignatureGroup {
	return v.value
}

func (v *NullableSignatureGroup) Set(val *SignatureGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureGroup(val *SignatureGroup) *NullableSignatureGroup {
	return &NullableSignatureGroup{value: val, isSet: true}
}

func (v NullableSignatureGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


