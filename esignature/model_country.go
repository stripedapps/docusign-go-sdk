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

// checks if the Country type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Country{}

// Country 
type Country struct {
	// 
	IsoCode *string `json:"isoCode,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// 
	Provinces []Province `json:"provinces,omitempty"`
	// 
	ProvinceValidated *string `json:"provinceValidated,omitempty"`
}

// NewCountry instantiates a new Country object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountry() *Country {
	this := Country{}
	return &this
}

// NewCountryWithDefaults instantiates a new Country object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryWithDefaults() *Country {
	this := Country{}
	return &this
}

// GetIsoCode returns the IsoCode field value if set, zero value otherwise.
func (o *Country) GetIsoCode() string {
	if o == nil || IsNil(o.IsoCode) {
		var ret string
		return ret
	}
	return *o.IsoCode
}

// GetIsoCodeOk returns a tuple with the IsoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetIsoCodeOk() (*string, bool) {
	if o == nil || IsNil(o.IsoCode) {
		return nil, false
	}
	return o.IsoCode, true
}

// HasIsoCode returns a boolean if a field has been set.
func (o *Country) HasIsoCode() bool {
	if o != nil && !IsNil(o.IsoCode) {
		return true
	}

	return false
}

// SetIsoCode gets a reference to the given string and assigns it to the IsoCode field.
func (o *Country) SetIsoCode(v string) {
	o.IsoCode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Country) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Country) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Country) SetName(v string) {
	o.Name = &v
}

// GetProvinces returns the Provinces field value if set, zero value otherwise.
func (o *Country) GetProvinces() []Province {
	if o == nil || IsNil(o.Provinces) {
		var ret []Province
		return ret
	}
	return o.Provinces
}

// GetProvincesOk returns a tuple with the Provinces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetProvincesOk() ([]Province, bool) {
	if o == nil || IsNil(o.Provinces) {
		return nil, false
	}
	return o.Provinces, true
}

// HasProvinces returns a boolean if a field has been set.
func (o *Country) HasProvinces() bool {
	if o != nil && !IsNil(o.Provinces) {
		return true
	}

	return false
}

// SetProvinces gets a reference to the given []Province and assigns it to the Provinces field.
func (o *Country) SetProvinces(v []Province) {
	o.Provinces = v
}

// GetProvinceValidated returns the ProvinceValidated field value if set, zero value otherwise.
func (o *Country) GetProvinceValidated() string {
	if o == nil || IsNil(o.ProvinceValidated) {
		var ret string
		return ret
	}
	return *o.ProvinceValidated
}

// GetProvinceValidatedOk returns a tuple with the ProvinceValidated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Country) GetProvinceValidatedOk() (*string, bool) {
	if o == nil || IsNil(o.ProvinceValidated) {
		return nil, false
	}
	return o.ProvinceValidated, true
}

// HasProvinceValidated returns a boolean if a field has been set.
func (o *Country) HasProvinceValidated() bool {
	if o != nil && !IsNil(o.ProvinceValidated) {
		return true
	}

	return false
}

// SetProvinceValidated gets a reference to the given string and assigns it to the ProvinceValidated field.
func (o *Country) SetProvinceValidated(v string) {
	o.ProvinceValidated = &v
}

func (o Country) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Country) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsoCode) {
		toSerialize["isoCode"] = o.IsoCode
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provinces) {
		toSerialize["provinces"] = o.Provinces
	}
	if !IsNil(o.ProvinceValidated) {
		toSerialize["provinceValidated"] = o.ProvinceValidated
	}
	return toSerialize, nil
}

type NullableCountry struct {
	value *Country
	isSet bool
}

func (v NullableCountry) Get() *Country {
	return v.value
}

func (v *NullableCountry) Set(val *Country) {
	v.value = val
	v.isSet = true
}

func (v NullableCountry) IsSet() bool {
	return v.isSet
}

func (v *NullableCountry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountry(val *Country) *NullableCountry {
	return &NullableCountry{value: val, isSet: true}
}

func (v NullableCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


