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

// checks if the ConnectUserObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectUserObject{}

// ConnectUserObject 
type ConnectUserObject struct {
	// The type of custom Connect configuration being accessed.
	Configurationtype *string `json:"configurationtype,omitempty"`
	// The ID of the custom Connect configuration being accessed.
	ConnectId *string `json:"connectId,omitempty"`
	// Boolean value that indicates whether the custom Connect configuration is enabled or not.
	Enabled *string `json:"enabled,omitempty"`
	// 
	HasAccess *string `json:"hasAccess,omitempty"`
	// 
	SenderSearchableItems []string `json:"senderSearchableItems,omitempty"`
}

// NewConnectUserObject instantiates a new ConnectUserObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectUserObject() *ConnectUserObject {
	this := ConnectUserObject{}
	return &this
}

// NewConnectUserObjectWithDefaults instantiates a new ConnectUserObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectUserObjectWithDefaults() *ConnectUserObject {
	this := ConnectUserObject{}
	return &this
}

// GetConfigurationtype returns the Configurationtype field value if set, zero value otherwise.
func (o *ConnectUserObject) GetConfigurationtype() string {
	if o == nil || IsNil(o.Configurationtype) {
		var ret string
		return ret
	}
	return *o.Configurationtype
}

// GetConfigurationtypeOk returns a tuple with the Configurationtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectUserObject) GetConfigurationtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Configurationtype) {
		return nil, false
	}
	return o.Configurationtype, true
}

// HasConfigurationtype returns a boolean if a field has been set.
func (o *ConnectUserObject) HasConfigurationtype() bool {
	if o != nil && !IsNil(o.Configurationtype) {
		return true
	}

	return false
}

// SetConfigurationtype gets a reference to the given string and assigns it to the Configurationtype field.
func (o *ConnectUserObject) SetConfigurationtype(v string) {
	o.Configurationtype = &v
}

// GetConnectId returns the ConnectId field value if set, zero value otherwise.
func (o *ConnectUserObject) GetConnectId() string {
	if o == nil || IsNil(o.ConnectId) {
		var ret string
		return ret
	}
	return *o.ConnectId
}

// GetConnectIdOk returns a tuple with the ConnectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectUserObject) GetConnectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectId) {
		return nil, false
	}
	return o.ConnectId, true
}

// HasConnectId returns a boolean if a field has been set.
func (o *ConnectUserObject) HasConnectId() bool {
	if o != nil && !IsNil(o.ConnectId) {
		return true
	}

	return false
}

// SetConnectId gets a reference to the given string and assigns it to the ConnectId field.
func (o *ConnectUserObject) SetConnectId(v string) {
	o.ConnectId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectUserObject) GetEnabled() string {
	if o == nil || IsNil(o.Enabled) {
		var ret string
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectUserObject) GetEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectUserObject) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given string and assigns it to the Enabled field.
func (o *ConnectUserObject) SetEnabled(v string) {
	o.Enabled = &v
}

// GetHasAccess returns the HasAccess field value if set, zero value otherwise.
func (o *ConnectUserObject) GetHasAccess() string {
	if o == nil || IsNil(o.HasAccess) {
		var ret string
		return ret
	}
	return *o.HasAccess
}

// GetHasAccessOk returns a tuple with the HasAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectUserObject) GetHasAccessOk() (*string, bool) {
	if o == nil || IsNil(o.HasAccess) {
		return nil, false
	}
	return o.HasAccess, true
}

// HasHasAccess returns a boolean if a field has been set.
func (o *ConnectUserObject) HasHasAccess() bool {
	if o != nil && !IsNil(o.HasAccess) {
		return true
	}

	return false
}

// SetHasAccess gets a reference to the given string and assigns it to the HasAccess field.
func (o *ConnectUserObject) SetHasAccess(v string) {
	o.HasAccess = &v
}

// GetSenderSearchableItems returns the SenderSearchableItems field value if set, zero value otherwise.
func (o *ConnectUserObject) GetSenderSearchableItems() []string {
	if o == nil || IsNil(o.SenderSearchableItems) {
		var ret []string
		return ret
	}
	return o.SenderSearchableItems
}

// GetSenderSearchableItemsOk returns a tuple with the SenderSearchableItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectUserObject) GetSenderSearchableItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.SenderSearchableItems) {
		return nil, false
	}
	return o.SenderSearchableItems, true
}

// HasSenderSearchableItems returns a boolean if a field has been set.
func (o *ConnectUserObject) HasSenderSearchableItems() bool {
	if o != nil && !IsNil(o.SenderSearchableItems) {
		return true
	}

	return false
}

// SetSenderSearchableItems gets a reference to the given []string and assigns it to the SenderSearchableItems field.
func (o *ConnectUserObject) SetSenderSearchableItems(v []string) {
	o.SenderSearchableItems = v
}

func (o ConnectUserObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectUserObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurationtype) {
		toSerialize["configurationtype"] = o.Configurationtype
	}
	if !IsNil(o.ConnectId) {
		toSerialize["connectId"] = o.ConnectId
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.HasAccess) {
		toSerialize["hasAccess"] = o.HasAccess
	}
	if !IsNil(o.SenderSearchableItems) {
		toSerialize["senderSearchableItems"] = o.SenderSearchableItems
	}
	return toSerialize, nil
}

type NullableConnectUserObject struct {
	value *ConnectUserObject
	isSet bool
}

func (v NullableConnectUserObject) Get() *ConnectUserObject {
	return v.value
}

func (v *NullableConnectUserObject) Set(val *ConnectUserObject) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectUserObject) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectUserObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectUserObject(val *ConnectUserObject) *NullableConnectUserObject {
	return &NullableConnectUserObject{value: val, isSet: true}
}

func (v NullableConnectUserObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectUserObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


