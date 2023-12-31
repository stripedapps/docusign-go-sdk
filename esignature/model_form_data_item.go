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

// checks if the FormDataItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDataItem{}

// FormDataItem 
type FormDataItem struct {
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The selected value in a list.
	ListSelectedValue *string `json:"listSelectedValue,omitempty"`
	// The name of the form field.
	Name *string `json:"name,omitempty"`
	// 
	NumericalValue *string `json:"numericalValue,omitempty"`
	// 
	OriginalNumericalValue *string `json:"originalNumericalValue,omitempty"`
	// The initial value associated with the form field.
	OriginalValue *string `json:"originalValue,omitempty"`
	// The current value associated with the form field.
	Value *string `json:"value,omitempty"`
}

// NewFormDataItem instantiates a new FormDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDataItem() *FormDataItem {
	this := FormDataItem{}
	return &this
}

// NewFormDataItemWithDefaults instantiates a new FormDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDataItemWithDefaults() *FormDataItem {
	this := FormDataItem{}
	return &this
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *FormDataItem) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *FormDataItem) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *FormDataItem) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetListSelectedValue returns the ListSelectedValue field value if set, zero value otherwise.
func (o *FormDataItem) GetListSelectedValue() string {
	if o == nil || IsNil(o.ListSelectedValue) {
		var ret string
		return ret
	}
	return *o.ListSelectedValue
}

// GetListSelectedValueOk returns a tuple with the ListSelectedValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetListSelectedValueOk() (*string, bool) {
	if o == nil || IsNil(o.ListSelectedValue) {
		return nil, false
	}
	return o.ListSelectedValue, true
}

// HasListSelectedValue returns a boolean if a field has been set.
func (o *FormDataItem) HasListSelectedValue() bool {
	if o != nil && !IsNil(o.ListSelectedValue) {
		return true
	}

	return false
}

// SetListSelectedValue gets a reference to the given string and assigns it to the ListSelectedValue field.
func (o *FormDataItem) SetListSelectedValue(v string) {
	o.ListSelectedValue = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormDataItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormDataItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormDataItem) SetName(v string) {
	o.Name = &v
}

// GetNumericalValue returns the NumericalValue field value if set, zero value otherwise.
func (o *FormDataItem) GetNumericalValue() string {
	if o == nil || IsNil(o.NumericalValue) {
		var ret string
		return ret
	}
	return *o.NumericalValue
}

// GetNumericalValueOk returns a tuple with the NumericalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetNumericalValueOk() (*string, bool) {
	if o == nil || IsNil(o.NumericalValue) {
		return nil, false
	}
	return o.NumericalValue, true
}

// HasNumericalValue returns a boolean if a field has been set.
func (o *FormDataItem) HasNumericalValue() bool {
	if o != nil && !IsNil(o.NumericalValue) {
		return true
	}

	return false
}

// SetNumericalValue gets a reference to the given string and assigns it to the NumericalValue field.
func (o *FormDataItem) SetNumericalValue(v string) {
	o.NumericalValue = &v
}

// GetOriginalNumericalValue returns the OriginalNumericalValue field value if set, zero value otherwise.
func (o *FormDataItem) GetOriginalNumericalValue() string {
	if o == nil || IsNil(o.OriginalNumericalValue) {
		var ret string
		return ret
	}
	return *o.OriginalNumericalValue
}

// GetOriginalNumericalValueOk returns a tuple with the OriginalNumericalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetOriginalNumericalValueOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalNumericalValue) {
		return nil, false
	}
	return o.OriginalNumericalValue, true
}

// HasOriginalNumericalValue returns a boolean if a field has been set.
func (o *FormDataItem) HasOriginalNumericalValue() bool {
	if o != nil && !IsNil(o.OriginalNumericalValue) {
		return true
	}

	return false
}

// SetOriginalNumericalValue gets a reference to the given string and assigns it to the OriginalNumericalValue field.
func (o *FormDataItem) SetOriginalNumericalValue(v string) {
	o.OriginalNumericalValue = &v
}

// GetOriginalValue returns the OriginalValue field value if set, zero value otherwise.
func (o *FormDataItem) GetOriginalValue() string {
	if o == nil || IsNil(o.OriginalValue) {
		var ret string
		return ret
	}
	return *o.OriginalValue
}

// GetOriginalValueOk returns a tuple with the OriginalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetOriginalValueOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalValue) {
		return nil, false
	}
	return o.OriginalValue, true
}

// HasOriginalValue returns a boolean if a field has been set.
func (o *FormDataItem) HasOriginalValue() bool {
	if o != nil && !IsNil(o.OriginalValue) {
		return true
	}

	return false
}

// SetOriginalValue gets a reference to the given string and assigns it to the OriginalValue field.
func (o *FormDataItem) SetOriginalValue(v string) {
	o.OriginalValue = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FormDataItem) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataItem) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FormDataItem) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FormDataItem) SetValue(v string) {
	o.Value = &v
}

func (o FormDataItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDataItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.ListSelectedValue) {
		toSerialize["listSelectedValue"] = o.ListSelectedValue
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NumericalValue) {
		toSerialize["numericalValue"] = o.NumericalValue
	}
	if !IsNil(o.OriginalNumericalValue) {
		toSerialize["originalNumericalValue"] = o.OriginalNumericalValue
	}
	if !IsNil(o.OriginalValue) {
		toSerialize["originalValue"] = o.OriginalValue
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFormDataItem struct {
	value *FormDataItem
	isSet bool
}

func (v NullableFormDataItem) Get() *FormDataItem {
	return v.value
}

func (v *NullableFormDataItem) Set(val *FormDataItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDataItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDataItem(val *FormDataItem) *NullableFormDataItem {
	return &NullableFormDataItem{value: val, isSet: true}
}

func (v NullableFormDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


