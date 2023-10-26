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

// checks if the AccountNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountNotification{}

// AccountNotification A complex element that specifies notifications (expirations and reminders) for the envelope.
type AccountNotification struct {
	Expirations *Expirations `json:"expirations,omitempty"`
	Reminders *Reminders `json:"reminders,omitempty"`
	// When **true,** the user can override envelope expirations.
	UserOverrideEnabled *string `json:"userOverrideEnabled,omitempty"`
}

// NewAccountNotification instantiates a new AccountNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNotification() *AccountNotification {
	this := AccountNotification{}
	return &this
}

// NewAccountNotificationWithDefaults instantiates a new AccountNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNotificationWithDefaults() *AccountNotification {
	this := AccountNotification{}
	return &this
}

// GetExpirations returns the Expirations field value if set, zero value otherwise.
func (o *AccountNotification) GetExpirations() Expirations {
	if o == nil || IsNil(o.Expirations) {
		var ret Expirations
		return ret
	}
	return *o.Expirations
}

// GetExpirationsOk returns a tuple with the Expirations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNotification) GetExpirationsOk() (*Expirations, bool) {
	if o == nil || IsNil(o.Expirations) {
		return nil, false
	}
	return o.Expirations, true
}

// HasExpirations returns a boolean if a field has been set.
func (o *AccountNotification) HasExpirations() bool {
	if o != nil && !IsNil(o.Expirations) {
		return true
	}

	return false
}

// SetExpirations gets a reference to the given Expirations and assigns it to the Expirations field.
func (o *AccountNotification) SetExpirations(v Expirations) {
	o.Expirations = &v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *AccountNotification) GetReminders() Reminders {
	if o == nil || IsNil(o.Reminders) {
		var ret Reminders
		return ret
	}
	return *o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNotification) GetRemindersOk() (*Reminders, bool) {
	if o == nil || IsNil(o.Reminders) {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *AccountNotification) HasReminders() bool {
	if o != nil && !IsNil(o.Reminders) {
		return true
	}

	return false
}

// SetReminders gets a reference to the given Reminders and assigns it to the Reminders field.
func (o *AccountNotification) SetReminders(v Reminders) {
	o.Reminders = &v
}

// GetUserOverrideEnabled returns the UserOverrideEnabled field value if set, zero value otherwise.
func (o *AccountNotification) GetUserOverrideEnabled() string {
	if o == nil || IsNil(o.UserOverrideEnabled) {
		var ret string
		return ret
	}
	return *o.UserOverrideEnabled
}

// GetUserOverrideEnabledOk returns a tuple with the UserOverrideEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNotification) GetUserOverrideEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.UserOverrideEnabled) {
		return nil, false
	}
	return o.UserOverrideEnabled, true
}

// HasUserOverrideEnabled returns a boolean if a field has been set.
func (o *AccountNotification) HasUserOverrideEnabled() bool {
	if o != nil && !IsNil(o.UserOverrideEnabled) {
		return true
	}

	return false
}

// SetUserOverrideEnabled gets a reference to the given string and assigns it to the UserOverrideEnabled field.
func (o *AccountNotification) SetUserOverrideEnabled(v string) {
	o.UserOverrideEnabled = &v
}

func (o AccountNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expirations) {
		toSerialize["expirations"] = o.Expirations
	}
	if !IsNil(o.Reminders) {
		toSerialize["reminders"] = o.Reminders
	}
	if !IsNil(o.UserOverrideEnabled) {
		toSerialize["userOverrideEnabled"] = o.UserOverrideEnabled
	}
	return toSerialize, nil
}

type NullableAccountNotification struct {
	value *AccountNotification
	isSet bool
}

func (v NullableAccountNotification) Get() *AccountNotification {
	return v.value
}

func (v *NullableAccountNotification) Set(val *AccountNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNotification(val *AccountNotification) *NullableAccountNotification {
	return &NullableAccountNotification{value: val, isSet: true}
}

func (v NullableAccountNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


