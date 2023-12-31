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

// checks if the AddressInformationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressInformationInput{}

// AddressInformationInput Contains address input information.
type AddressInformationInput struct {
	AddressInformation *AddressInformation `json:"addressInformation,omitempty"`
	// Specifies the display level for the recipient. Valid values are: * `ReadOnly` * `Editable` * `DoNotDisplay`
	DisplayLevelCode *string `json:"displayLevelCode,omitempty"`
	// A Boolean value that specifies whether the information must be returned in the response.
	ReceiveInResponse *string `json:"receiveInResponse,omitempty"`
}

// NewAddressInformationInput instantiates a new AddressInformationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressInformationInput() *AddressInformationInput {
	this := AddressInformationInput{}
	return &this
}

// NewAddressInformationInputWithDefaults instantiates a new AddressInformationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressInformationInputWithDefaults() *AddressInformationInput {
	this := AddressInformationInput{}
	return &this
}

// GetAddressInformation returns the AddressInformation field value if set, zero value otherwise.
func (o *AddressInformationInput) GetAddressInformation() AddressInformation {
	if o == nil || IsNil(o.AddressInformation) {
		var ret AddressInformation
		return ret
	}
	return *o.AddressInformation
}

// GetAddressInformationOk returns a tuple with the AddressInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInformationInput) GetAddressInformationOk() (*AddressInformation, bool) {
	if o == nil || IsNil(o.AddressInformation) {
		return nil, false
	}
	return o.AddressInformation, true
}

// HasAddressInformation returns a boolean if a field has been set.
func (o *AddressInformationInput) HasAddressInformation() bool {
	if o != nil && !IsNil(o.AddressInformation) {
		return true
	}

	return false
}

// SetAddressInformation gets a reference to the given AddressInformation and assigns it to the AddressInformation field.
func (o *AddressInformationInput) SetAddressInformation(v AddressInformation) {
	o.AddressInformation = &v
}

// GetDisplayLevelCode returns the DisplayLevelCode field value if set, zero value otherwise.
func (o *AddressInformationInput) GetDisplayLevelCode() string {
	if o == nil || IsNil(o.DisplayLevelCode) {
		var ret string
		return ret
	}
	return *o.DisplayLevelCode
}

// GetDisplayLevelCodeOk returns a tuple with the DisplayLevelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInformationInput) GetDisplayLevelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLevelCode) {
		return nil, false
	}
	return o.DisplayLevelCode, true
}

// HasDisplayLevelCode returns a boolean if a field has been set.
func (o *AddressInformationInput) HasDisplayLevelCode() bool {
	if o != nil && !IsNil(o.DisplayLevelCode) {
		return true
	}

	return false
}

// SetDisplayLevelCode gets a reference to the given string and assigns it to the DisplayLevelCode field.
func (o *AddressInformationInput) SetDisplayLevelCode(v string) {
	o.DisplayLevelCode = &v
}

// GetReceiveInResponse returns the ReceiveInResponse field value if set, zero value otherwise.
func (o *AddressInformationInput) GetReceiveInResponse() string {
	if o == nil || IsNil(o.ReceiveInResponse) {
		var ret string
		return ret
	}
	return *o.ReceiveInResponse
}

// GetReceiveInResponseOk returns a tuple with the ReceiveInResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressInformationInput) GetReceiveInResponseOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveInResponse) {
		return nil, false
	}
	return o.ReceiveInResponse, true
}

// HasReceiveInResponse returns a boolean if a field has been set.
func (o *AddressInformationInput) HasReceiveInResponse() bool {
	if o != nil && !IsNil(o.ReceiveInResponse) {
		return true
	}

	return false
}

// SetReceiveInResponse gets a reference to the given string and assigns it to the ReceiveInResponse field.
func (o *AddressInformationInput) SetReceiveInResponse(v string) {
	o.ReceiveInResponse = &v
}

func (o AddressInformationInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressInformationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressInformation) {
		toSerialize["addressInformation"] = o.AddressInformation
	}
	if !IsNil(o.DisplayLevelCode) {
		toSerialize["displayLevelCode"] = o.DisplayLevelCode
	}
	if !IsNil(o.ReceiveInResponse) {
		toSerialize["receiveInResponse"] = o.ReceiveInResponse
	}
	return toSerialize, nil
}

type NullableAddressInformationInput struct {
	value *AddressInformationInput
	isSet bool
}

func (v NullableAddressInformationInput) Get() *AddressInformationInput {
	return v.value
}

func (v *NullableAddressInformationInput) Set(val *AddressInformationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressInformationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressInformationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressInformationInput(val *AddressInformationInput) *NullableAddressInformationInput {
	return &NullableAddressInformationInput{value: val, isSet: true}
}

func (v NullableAddressInformationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressInformationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


