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

// checks if the NotificationDefaults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationDefaults{}

// NotificationDefaults The `NotificationDefaults` resource provides methods that enable you to manage the default notifications for envelopes.
type NotificationDefaults struct {
	ApiEmailNotifications *NotificationDefaultSettings `json:"apiEmailNotifications,omitempty"`
	EmailNotifications *NotificationDefaultSettings `json:"emailNotifications,omitempty"`
}

// NewNotificationDefaults instantiates a new NotificationDefaults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationDefaults() *NotificationDefaults {
	this := NotificationDefaults{}
	return &this
}

// NewNotificationDefaultsWithDefaults instantiates a new NotificationDefaults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationDefaultsWithDefaults() *NotificationDefaults {
	this := NotificationDefaults{}
	return &this
}

// GetApiEmailNotifications returns the ApiEmailNotifications field value if set, zero value otherwise.
func (o *NotificationDefaults) GetApiEmailNotifications() NotificationDefaultSettings {
	if o == nil || IsNil(o.ApiEmailNotifications) {
		var ret NotificationDefaultSettings
		return ret
	}
	return *o.ApiEmailNotifications
}

// GetApiEmailNotificationsOk returns a tuple with the ApiEmailNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDefaults) GetApiEmailNotificationsOk() (*NotificationDefaultSettings, bool) {
	if o == nil || IsNil(o.ApiEmailNotifications) {
		return nil, false
	}
	return o.ApiEmailNotifications, true
}

// HasApiEmailNotifications returns a boolean if a field has been set.
func (o *NotificationDefaults) HasApiEmailNotifications() bool {
	if o != nil && !IsNil(o.ApiEmailNotifications) {
		return true
	}

	return false
}

// SetApiEmailNotifications gets a reference to the given NotificationDefaultSettings and assigns it to the ApiEmailNotifications field.
func (o *NotificationDefaults) SetApiEmailNotifications(v NotificationDefaultSettings) {
	o.ApiEmailNotifications = &v
}

// GetEmailNotifications returns the EmailNotifications field value if set, zero value otherwise.
func (o *NotificationDefaults) GetEmailNotifications() NotificationDefaultSettings {
	if o == nil || IsNil(o.EmailNotifications) {
		var ret NotificationDefaultSettings
		return ret
	}
	return *o.EmailNotifications
}

// GetEmailNotificationsOk returns a tuple with the EmailNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationDefaults) GetEmailNotificationsOk() (*NotificationDefaultSettings, bool) {
	if o == nil || IsNil(o.EmailNotifications) {
		return nil, false
	}
	return o.EmailNotifications, true
}

// HasEmailNotifications returns a boolean if a field has been set.
func (o *NotificationDefaults) HasEmailNotifications() bool {
	if o != nil && !IsNil(o.EmailNotifications) {
		return true
	}

	return false
}

// SetEmailNotifications gets a reference to the given NotificationDefaultSettings and assigns it to the EmailNotifications field.
func (o *NotificationDefaults) SetEmailNotifications(v NotificationDefaultSettings) {
	o.EmailNotifications = &v
}

func (o NotificationDefaults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationDefaults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiEmailNotifications) {
		toSerialize["apiEmailNotifications"] = o.ApiEmailNotifications
	}
	if !IsNil(o.EmailNotifications) {
		toSerialize["emailNotifications"] = o.EmailNotifications
	}
	return toSerialize, nil
}

type NullableNotificationDefaults struct {
	value *NotificationDefaults
	isSet bool
}

func (v NullableNotificationDefaults) Get() *NotificationDefaults {
	return v.value
}

func (v *NullableNotificationDefaults) Set(val *NotificationDefaults) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationDefaults) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationDefaults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationDefaults(val *NotificationDefaults) *NullableNotificationDefaults {
	return &NullableNotificationDefaults{value: val, isSet: true}
}

func (v NullableNotificationDefaults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationDefaults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

