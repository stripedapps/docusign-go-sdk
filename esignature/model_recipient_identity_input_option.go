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

// checks if the RecipientIdentityInputOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientIdentityInputOption{}

// RecipientIdentityInputOption 
type RecipientIdentityInputOption struct {
	// 
	Name *string `json:"name,omitempty"`
	// 
	PhoneNumberList []RecipientIdentityPhoneNumber `json:"phoneNumberList,omitempty"`
	// 
	ValueType *string `json:"valueType,omitempty"`
}

// NewRecipientIdentityInputOption instantiates a new RecipientIdentityInputOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientIdentityInputOption() *RecipientIdentityInputOption {
	this := RecipientIdentityInputOption{}
	return &this
}

// NewRecipientIdentityInputOptionWithDefaults instantiates a new RecipientIdentityInputOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientIdentityInputOptionWithDefaults() *RecipientIdentityInputOption {
	this := RecipientIdentityInputOption{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecipientIdentityInputOption) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientIdentityInputOption) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecipientIdentityInputOption) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecipientIdentityInputOption) SetName(v string) {
	o.Name = &v
}

// GetPhoneNumberList returns the PhoneNumberList field value if set, zero value otherwise.
func (o *RecipientIdentityInputOption) GetPhoneNumberList() []RecipientIdentityPhoneNumber {
	if o == nil || IsNil(o.PhoneNumberList) {
		var ret []RecipientIdentityPhoneNumber
		return ret
	}
	return o.PhoneNumberList
}

// GetPhoneNumberListOk returns a tuple with the PhoneNumberList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientIdentityInputOption) GetPhoneNumberListOk() ([]RecipientIdentityPhoneNumber, bool) {
	if o == nil || IsNil(o.PhoneNumberList) {
		return nil, false
	}
	return o.PhoneNumberList, true
}

// HasPhoneNumberList returns a boolean if a field has been set.
func (o *RecipientIdentityInputOption) HasPhoneNumberList() bool {
	if o != nil && !IsNil(o.PhoneNumberList) {
		return true
	}

	return false
}

// SetPhoneNumberList gets a reference to the given []RecipientIdentityPhoneNumber and assigns it to the PhoneNumberList field.
func (o *RecipientIdentityInputOption) SetPhoneNumberList(v []RecipientIdentityPhoneNumber) {
	o.PhoneNumberList = v
}

// GetValueType returns the ValueType field value if set, zero value otherwise.
func (o *RecipientIdentityInputOption) GetValueType() string {
	if o == nil || IsNil(o.ValueType) {
		var ret string
		return ret
	}
	return *o.ValueType
}

// GetValueTypeOk returns a tuple with the ValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientIdentityInputOption) GetValueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ValueType) {
		return nil, false
	}
	return o.ValueType, true
}

// HasValueType returns a boolean if a field has been set.
func (o *RecipientIdentityInputOption) HasValueType() bool {
	if o != nil && !IsNil(o.ValueType) {
		return true
	}

	return false
}

// SetValueType gets a reference to the given string and assigns it to the ValueType field.
func (o *RecipientIdentityInputOption) SetValueType(v string) {
	o.ValueType = &v
}

func (o RecipientIdentityInputOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientIdentityInputOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PhoneNumberList) {
		toSerialize["phoneNumberList"] = o.PhoneNumberList
	}
	if !IsNil(o.ValueType) {
		toSerialize["valueType"] = o.ValueType
	}
	return toSerialize, nil
}

type NullableRecipientIdentityInputOption struct {
	value *RecipientIdentityInputOption
	isSet bool
}

func (v NullableRecipientIdentityInputOption) Get() *RecipientIdentityInputOption {
	return v.value
}

func (v *NullableRecipientIdentityInputOption) Set(val *RecipientIdentityInputOption) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientIdentityInputOption) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientIdentityInputOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientIdentityInputOption(val *RecipientIdentityInputOption) *NullableRecipientIdentityInputOption {
	return &NullableRecipientIdentityInputOption{value: val, isSet: true}
}

func (v NullableRecipientIdentityInputOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientIdentityInputOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


