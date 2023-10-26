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

// checks if the MobileNotifierConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileNotifierConfiguration{}

// MobileNotifierConfiguration 
type MobileNotifierConfiguration struct {
	// 
	DeviceId *string `json:"deviceId,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The Platform of the client application
	Platform *string `json:"platform,omitempty"`
}

// NewMobileNotifierConfiguration instantiates a new MobileNotifierConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileNotifierConfiguration() *MobileNotifierConfiguration {
	this := MobileNotifierConfiguration{}
	return &this
}

// NewMobileNotifierConfigurationWithDefaults instantiates a new MobileNotifierConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileNotifierConfigurationWithDefaults() *MobileNotifierConfiguration {
	this := MobileNotifierConfiguration{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *MobileNotifierConfiguration) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileNotifierConfiguration) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *MobileNotifierConfiguration) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *MobileNotifierConfiguration) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *MobileNotifierConfiguration) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileNotifierConfiguration) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *MobileNotifierConfiguration) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *MobileNotifierConfiguration) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *MobileNotifierConfiguration) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileNotifierConfiguration) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *MobileNotifierConfiguration) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *MobileNotifierConfiguration) SetPlatform(v string) {
	o.Platform = &v
}

func (o MobileNotifierConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileNotifierConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["deviceId"] = o.DeviceId
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	return toSerialize, nil
}

type NullableMobileNotifierConfiguration struct {
	value *MobileNotifierConfiguration
	isSet bool
}

func (v NullableMobileNotifierConfiguration) Get() *MobileNotifierConfiguration {
	return v.value
}

func (v *NullableMobileNotifierConfiguration) Set(val *MobileNotifierConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileNotifierConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileNotifierConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileNotifierConfiguration(val *MobileNotifierConfiguration) *NullableMobileNotifierConfiguration {
	return &NullableMobileNotifierConfiguration{value: val, isSet: true}
}

func (v NullableMobileNotifierConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileNotifierConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

