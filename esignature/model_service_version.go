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

// checks if the ServiceVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVersion{}

// ServiceVersion 
type ServiceVersion struct {
	// The version of the rest API.
	Version *string `json:"version,omitempty"`
	// 
	VersionUrl *string `json:"versionUrl,omitempty"`
}

// NewServiceVersion instantiates a new ServiceVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVersion() *ServiceVersion {
	this := ServiceVersion{}
	return &this
}

// NewServiceVersionWithDefaults instantiates a new ServiceVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVersionWithDefaults() *ServiceVersion {
	this := ServiceVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceVersion) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ServiceVersion) SetVersion(v string) {
	o.Version = &v
}

// GetVersionUrl returns the VersionUrl field value if set, zero value otherwise.
func (o *ServiceVersion) GetVersionUrl() string {
	if o == nil || IsNil(o.VersionUrl) {
		var ret string
		return ret
	}
	return *o.VersionUrl
}

// GetVersionUrlOk returns a tuple with the VersionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVersion) GetVersionUrlOk() (*string, bool) {
	if o == nil || IsNil(o.VersionUrl) {
		return nil, false
	}
	return o.VersionUrl, true
}

// HasVersionUrl returns a boolean if a field has been set.
func (o *ServiceVersion) HasVersionUrl() bool {
	if o != nil && !IsNil(o.VersionUrl) {
		return true
	}

	return false
}

// SetVersionUrl gets a reference to the given string and assigns it to the VersionUrl field.
func (o *ServiceVersion) SetVersionUrl(v string) {
	o.VersionUrl = &v
}

func (o ServiceVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionUrl) {
		toSerialize["versionUrl"] = o.VersionUrl
	}
	return toSerialize, nil
}

type NullableServiceVersion struct {
	value *ServiceVersion
	isSet bool
}

func (v NullableServiceVersion) Get() *ServiceVersion {
	return v.value
}

func (v *NullableServiceVersion) Set(val *ServiceVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVersion(val *ServiceVersion) *NullableServiceVersion {
	return &NullableServiceVersion{value: val, isSet: true}
}

func (v NullableServiceVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


