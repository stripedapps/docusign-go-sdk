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

// checks if the Ssn4InformationInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ssn4InformationInput{}

// Ssn4InformationInput 
type Ssn4InformationInput struct {
	// Specifies the display level for the recipient. Valid values are: * `ReadOnly` * `Editable` * `DoNotDisplay`
	DisplayLevelCode *string `json:"displayLevelCode,omitempty"`
	// A Boolean value that specifies whether the information must be returned in the response.
	ReceiveInResponse *string `json:"receiveInResponse,omitempty"`
	// The last four digits of the recipient's Social Security Number (SSN).
	Ssn4 *string `json:"ssn4,omitempty"`
}

// NewSsn4InformationInput instantiates a new Ssn4InformationInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsn4InformationInput() *Ssn4InformationInput {
	this := Ssn4InformationInput{}
	return &this
}

// NewSsn4InformationInputWithDefaults instantiates a new Ssn4InformationInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSsn4InformationInputWithDefaults() *Ssn4InformationInput {
	this := Ssn4InformationInput{}
	return &this
}

// GetDisplayLevelCode returns the DisplayLevelCode field value if set, zero value otherwise.
func (o *Ssn4InformationInput) GetDisplayLevelCode() string {
	if o == nil || IsNil(o.DisplayLevelCode) {
		var ret string
		return ret
	}
	return *o.DisplayLevelCode
}

// GetDisplayLevelCodeOk returns a tuple with the DisplayLevelCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssn4InformationInput) GetDisplayLevelCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayLevelCode) {
		return nil, false
	}
	return o.DisplayLevelCode, true
}

// HasDisplayLevelCode returns a boolean if a field has been set.
func (o *Ssn4InformationInput) HasDisplayLevelCode() bool {
	if o != nil && !IsNil(o.DisplayLevelCode) {
		return true
	}

	return false
}

// SetDisplayLevelCode gets a reference to the given string and assigns it to the DisplayLevelCode field.
func (o *Ssn4InformationInput) SetDisplayLevelCode(v string) {
	o.DisplayLevelCode = &v
}

// GetReceiveInResponse returns the ReceiveInResponse field value if set, zero value otherwise.
func (o *Ssn4InformationInput) GetReceiveInResponse() string {
	if o == nil || IsNil(o.ReceiveInResponse) {
		var ret string
		return ret
	}
	return *o.ReceiveInResponse
}

// GetReceiveInResponseOk returns a tuple with the ReceiveInResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssn4InformationInput) GetReceiveInResponseOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiveInResponse) {
		return nil, false
	}
	return o.ReceiveInResponse, true
}

// HasReceiveInResponse returns a boolean if a field has been set.
func (o *Ssn4InformationInput) HasReceiveInResponse() bool {
	if o != nil && !IsNil(o.ReceiveInResponse) {
		return true
	}

	return false
}

// SetReceiveInResponse gets a reference to the given string and assigns it to the ReceiveInResponse field.
func (o *Ssn4InformationInput) SetReceiveInResponse(v string) {
	o.ReceiveInResponse = &v
}

// GetSsn4 returns the Ssn4 field value if set, zero value otherwise.
func (o *Ssn4InformationInput) GetSsn4() string {
	if o == nil || IsNil(o.Ssn4) {
		var ret string
		return ret
	}
	return *o.Ssn4
}

// GetSsn4Ok returns a tuple with the Ssn4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssn4InformationInput) GetSsn4Ok() (*string, bool) {
	if o == nil || IsNil(o.Ssn4) {
		return nil, false
	}
	return o.Ssn4, true
}

// HasSsn4 returns a boolean if a field has been set.
func (o *Ssn4InformationInput) HasSsn4() bool {
	if o != nil && !IsNil(o.Ssn4) {
		return true
	}

	return false
}

// SetSsn4 gets a reference to the given string and assigns it to the Ssn4 field.
func (o *Ssn4InformationInput) SetSsn4(v string) {
	o.Ssn4 = &v
}

func (o Ssn4InformationInput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ssn4InformationInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayLevelCode) {
		toSerialize["displayLevelCode"] = o.DisplayLevelCode
	}
	if !IsNil(o.ReceiveInResponse) {
		toSerialize["receiveInResponse"] = o.ReceiveInResponse
	}
	if !IsNil(o.Ssn4) {
		toSerialize["ssn4"] = o.Ssn4
	}
	return toSerialize, nil
}

type NullableSsn4InformationInput struct {
	value *Ssn4InformationInput
	isSet bool
}

func (v NullableSsn4InformationInput) Get() *Ssn4InformationInput {
	return v.value
}

func (v *NullableSsn4InformationInput) Set(val *Ssn4InformationInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSsn4InformationInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSsn4InformationInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsn4InformationInput(val *Ssn4InformationInput) *NullableSsn4InformationInput {
	return &NullableSsn4InformationInput{value: val, isSet: true}
}

func (v NullableSsn4InformationInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsn4InformationInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


