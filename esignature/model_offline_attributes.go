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

// checks if the OfflineAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OfflineAttributes{}

// OfflineAttributes Reserved for DocuSign.
type OfflineAttributes struct {
	// Reserved for DocuSign.
	AccountEsignId *string `json:"accountEsignId,omitempty"`
	// Reserved for DocuSign.
	DeviceModel *string `json:"deviceModel,omitempty"`
	// Reserved for DocuSign.
	DeviceName *string `json:"deviceName,omitempty"`
	// Reserved for DocuSign.
	GpsLatitude *string `json:"gpsLatitude,omitempty"`
	// Reserved for DocuSign.
	GpsLongitude *string `json:"gpsLongitude,omitempty"`
	// Reserved for DocuSign.
	OfflineSigningHash *string `json:"offlineSigningHash,omitempty"`
}

// NewOfflineAttributes instantiates a new OfflineAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfflineAttributes() *OfflineAttributes {
	this := OfflineAttributes{}
	return &this
}

// NewOfflineAttributesWithDefaults instantiates a new OfflineAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfflineAttributesWithDefaults() *OfflineAttributes {
	this := OfflineAttributes{}
	return &this
}

// GetAccountEsignId returns the AccountEsignId field value if set, zero value otherwise.
func (o *OfflineAttributes) GetAccountEsignId() string {
	if o == nil || IsNil(o.AccountEsignId) {
		var ret string
		return ret
	}
	return *o.AccountEsignId
}

// GetAccountEsignIdOk returns a tuple with the AccountEsignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetAccountEsignIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountEsignId) {
		return nil, false
	}
	return o.AccountEsignId, true
}

// HasAccountEsignId returns a boolean if a field has been set.
func (o *OfflineAttributes) HasAccountEsignId() bool {
	if o != nil && !IsNil(o.AccountEsignId) {
		return true
	}

	return false
}

// SetAccountEsignId gets a reference to the given string and assigns it to the AccountEsignId field.
func (o *OfflineAttributes) SetAccountEsignId(v string) {
	o.AccountEsignId = &v
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise.
func (o *OfflineAttributes) GetDeviceModel() string {
	if o == nil || IsNil(o.DeviceModel) {
		var ret string
		return ret
	}
	return *o.DeviceModel
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetDeviceModelOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceModel) {
		return nil, false
	}
	return o.DeviceModel, true
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *OfflineAttributes) HasDeviceModel() bool {
	if o != nil && !IsNil(o.DeviceModel) {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given string and assigns it to the DeviceModel field.
func (o *OfflineAttributes) SetDeviceModel(v string) {
	o.DeviceModel = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *OfflineAttributes) GetDeviceName() string {
	if o == nil || IsNil(o.DeviceName) {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetDeviceNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceName) {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *OfflineAttributes) HasDeviceName() bool {
	if o != nil && !IsNil(o.DeviceName) {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *OfflineAttributes) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetGpsLatitude returns the GpsLatitude field value if set, zero value otherwise.
func (o *OfflineAttributes) GetGpsLatitude() string {
	if o == nil || IsNil(o.GpsLatitude) {
		var ret string
		return ret
	}
	return *o.GpsLatitude
}

// GetGpsLatitudeOk returns a tuple with the GpsLatitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetGpsLatitudeOk() (*string, bool) {
	if o == nil || IsNil(o.GpsLatitude) {
		return nil, false
	}
	return o.GpsLatitude, true
}

// HasGpsLatitude returns a boolean if a field has been set.
func (o *OfflineAttributes) HasGpsLatitude() bool {
	if o != nil && !IsNil(o.GpsLatitude) {
		return true
	}

	return false
}

// SetGpsLatitude gets a reference to the given string and assigns it to the GpsLatitude field.
func (o *OfflineAttributes) SetGpsLatitude(v string) {
	o.GpsLatitude = &v
}

// GetGpsLongitude returns the GpsLongitude field value if set, zero value otherwise.
func (o *OfflineAttributes) GetGpsLongitude() string {
	if o == nil || IsNil(o.GpsLongitude) {
		var ret string
		return ret
	}
	return *o.GpsLongitude
}

// GetGpsLongitudeOk returns a tuple with the GpsLongitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetGpsLongitudeOk() (*string, bool) {
	if o == nil || IsNil(o.GpsLongitude) {
		return nil, false
	}
	return o.GpsLongitude, true
}

// HasGpsLongitude returns a boolean if a field has been set.
func (o *OfflineAttributes) HasGpsLongitude() bool {
	if o != nil && !IsNil(o.GpsLongitude) {
		return true
	}

	return false
}

// SetGpsLongitude gets a reference to the given string and assigns it to the GpsLongitude field.
func (o *OfflineAttributes) SetGpsLongitude(v string) {
	o.GpsLongitude = &v
}

// GetOfflineSigningHash returns the OfflineSigningHash field value if set, zero value otherwise.
func (o *OfflineAttributes) GetOfflineSigningHash() string {
	if o == nil || IsNil(o.OfflineSigningHash) {
		var ret string
		return ret
	}
	return *o.OfflineSigningHash
}

// GetOfflineSigningHashOk returns a tuple with the OfflineSigningHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfflineAttributes) GetOfflineSigningHashOk() (*string, bool) {
	if o == nil || IsNil(o.OfflineSigningHash) {
		return nil, false
	}
	return o.OfflineSigningHash, true
}

// HasOfflineSigningHash returns a boolean if a field has been set.
func (o *OfflineAttributes) HasOfflineSigningHash() bool {
	if o != nil && !IsNil(o.OfflineSigningHash) {
		return true
	}

	return false
}

// SetOfflineSigningHash gets a reference to the given string and assigns it to the OfflineSigningHash field.
func (o *OfflineAttributes) SetOfflineSigningHash(v string) {
	o.OfflineSigningHash = &v
}

func (o OfflineAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OfflineAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountEsignId) {
		toSerialize["accountEsignId"] = o.AccountEsignId
	}
	if !IsNil(o.DeviceModel) {
		toSerialize["deviceModel"] = o.DeviceModel
	}
	if !IsNil(o.DeviceName) {
		toSerialize["deviceName"] = o.DeviceName
	}
	if !IsNil(o.GpsLatitude) {
		toSerialize["gpsLatitude"] = o.GpsLatitude
	}
	if !IsNil(o.GpsLongitude) {
		toSerialize["gpsLongitude"] = o.GpsLongitude
	}
	if !IsNil(o.OfflineSigningHash) {
		toSerialize["offlineSigningHash"] = o.OfflineSigningHash
	}
	return toSerialize, nil
}

type NullableOfflineAttributes struct {
	value *OfflineAttributes
	isSet bool
}

func (v NullableOfflineAttributes) Get() *OfflineAttributes {
	return v.value
}

func (v *NullableOfflineAttributes) Set(val *OfflineAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableOfflineAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableOfflineAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfflineAttributes(val *OfflineAttributes) *NullableOfflineAttributes {
	return &NullableOfflineAttributes{value: val, isSet: true}
}

func (v NullableOfflineAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfflineAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


