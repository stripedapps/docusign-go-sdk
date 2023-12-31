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

// checks if the TemplateNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateNotificationRequest{}

// TemplateNotificationRequest 
type TemplateNotificationRequest struct {
	Expirations *Expirations `json:"expirations,omitempty"`
	// The user's encrypted password hash.
	Password *string `json:"password,omitempty"`
	Reminders *Reminders `json:"reminders,omitempty"`
	// When **true,** the account default notification settings are used for the envelope, overriding the reminders and expirations settings. When **false,** the reminders and expirations settings specified in this request are used. The default value is **false.**
	UseAccountDefaults *string `json:"useAccountDefaults,omitempty"`
}

// NewTemplateNotificationRequest instantiates a new TemplateNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateNotificationRequest() *TemplateNotificationRequest {
	this := TemplateNotificationRequest{}
	return &this
}

// NewTemplateNotificationRequestWithDefaults instantiates a new TemplateNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateNotificationRequestWithDefaults() *TemplateNotificationRequest {
	this := TemplateNotificationRequest{}
	return &this
}

// GetExpirations returns the Expirations field value if set, zero value otherwise.
func (o *TemplateNotificationRequest) GetExpirations() Expirations {
	if o == nil || IsNil(o.Expirations) {
		var ret Expirations
		return ret
	}
	return *o.Expirations
}

// GetExpirationsOk returns a tuple with the Expirations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateNotificationRequest) GetExpirationsOk() (*Expirations, bool) {
	if o == nil || IsNil(o.Expirations) {
		return nil, false
	}
	return o.Expirations, true
}

// HasExpirations returns a boolean if a field has been set.
func (o *TemplateNotificationRequest) HasExpirations() bool {
	if o != nil && !IsNil(o.Expirations) {
		return true
	}

	return false
}

// SetExpirations gets a reference to the given Expirations and assigns it to the Expirations field.
func (o *TemplateNotificationRequest) SetExpirations(v Expirations) {
	o.Expirations = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *TemplateNotificationRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateNotificationRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *TemplateNotificationRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *TemplateNotificationRequest) SetPassword(v string) {
	o.Password = &v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *TemplateNotificationRequest) GetReminders() Reminders {
	if o == nil || IsNil(o.Reminders) {
		var ret Reminders
		return ret
	}
	return *o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateNotificationRequest) GetRemindersOk() (*Reminders, bool) {
	if o == nil || IsNil(o.Reminders) {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *TemplateNotificationRequest) HasReminders() bool {
	if o != nil && !IsNil(o.Reminders) {
		return true
	}

	return false
}

// SetReminders gets a reference to the given Reminders and assigns it to the Reminders field.
func (o *TemplateNotificationRequest) SetReminders(v Reminders) {
	o.Reminders = &v
}

// GetUseAccountDefaults returns the UseAccountDefaults field value if set, zero value otherwise.
func (o *TemplateNotificationRequest) GetUseAccountDefaults() string {
	if o == nil || IsNil(o.UseAccountDefaults) {
		var ret string
		return ret
	}
	return *o.UseAccountDefaults
}

// GetUseAccountDefaultsOk returns a tuple with the UseAccountDefaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateNotificationRequest) GetUseAccountDefaultsOk() (*string, bool) {
	if o == nil || IsNil(o.UseAccountDefaults) {
		return nil, false
	}
	return o.UseAccountDefaults, true
}

// HasUseAccountDefaults returns a boolean if a field has been set.
func (o *TemplateNotificationRequest) HasUseAccountDefaults() bool {
	if o != nil && !IsNil(o.UseAccountDefaults) {
		return true
	}

	return false
}

// SetUseAccountDefaults gets a reference to the given string and assigns it to the UseAccountDefaults field.
func (o *TemplateNotificationRequest) SetUseAccountDefaults(v string) {
	o.UseAccountDefaults = &v
}

func (o TemplateNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expirations) {
		toSerialize["expirations"] = o.Expirations
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Reminders) {
		toSerialize["reminders"] = o.Reminders
	}
	if !IsNil(o.UseAccountDefaults) {
		toSerialize["useAccountDefaults"] = o.UseAccountDefaults
	}
	return toSerialize, nil
}

type NullableTemplateNotificationRequest struct {
	value *TemplateNotificationRequest
	isSet bool
}

func (v NullableTemplateNotificationRequest) Get() *TemplateNotificationRequest {
	return v.value
}

func (v *NullableTemplateNotificationRequest) Set(val *TemplateNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateNotificationRequest(val *TemplateNotificationRequest) *NullableTemplateNotificationRequest {
	return &NullableTemplateNotificationRequest{value: val, isSet: true}
}

func (v NullableTemplateNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


