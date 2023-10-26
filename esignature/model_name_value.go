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

// checks if the NameValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NameValue{}

// NameValue A name-value pair that describes an item and provides a value for the item.
type NameValue struct {
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The name of the item.
	Name *string `json:"name,omitempty"`
	// The initial value of the item.
	OriginalValue *string `json:"originalValue,omitempty"`
	// The current value of the item.
	Value *string `json:"value,omitempty"`
}

// NewNameValue instantiates a new NameValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameValue() *NameValue {
	this := NameValue{}
	return &this
}

// NewNameValueWithDefaults instantiates a new NameValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameValueWithDefaults() *NameValue {
	this := NameValue{}
	return &this
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *NameValue) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValue) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *NameValue) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *NameValue) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NameValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NameValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NameValue) SetName(v string) {
	o.Name = &v
}

// GetOriginalValue returns the OriginalValue field value if set, zero value otherwise.
func (o *NameValue) GetOriginalValue() string {
	if o == nil || IsNil(o.OriginalValue) {
		var ret string
		return ret
	}
	return *o.OriginalValue
}

// GetOriginalValueOk returns a tuple with the OriginalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValue) GetOriginalValueOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalValue) {
		return nil, false
	}
	return o.OriginalValue, true
}

// HasOriginalValue returns a boolean if a field has been set.
func (o *NameValue) HasOriginalValue() bool {
	if o != nil && !IsNil(o.OriginalValue) {
		return true
	}

	return false
}

// SetOriginalValue gets a reference to the given string and assigns it to the OriginalValue field.
func (o *NameValue) SetOriginalValue(v string) {
	o.OriginalValue = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *NameValue) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameValue) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *NameValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *NameValue) SetValue(v string) {
	o.Value = &v
}

func (o NameValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NameValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OriginalValue) {
		toSerialize["originalValue"] = o.OriginalValue
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableNameValue struct {
	value *NameValue
	isSet bool
}

func (v NullableNameValue) Get() *NameValue {
	return v.value
}

func (v *NullableNameValue) Set(val *NameValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNameValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNameValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameValue(val *NameValue) *NullableNameValue {
	return &NullableNameValue{value: val, isSet: true}
}

func (v NullableNameValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

