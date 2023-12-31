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

// checks if the DocGenFormFieldResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocGenFormFieldResponse{}

// DocGenFormFieldResponse An object for document generation responses.
type DocGenFormFieldResponse struct {
	// A list of `docGenFormFields` objects.
	DocGenFormFields []DocGenFormFields `json:"docGenFormFields,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
}

// NewDocGenFormFieldResponse instantiates a new DocGenFormFieldResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocGenFormFieldResponse() *DocGenFormFieldResponse {
	this := DocGenFormFieldResponse{}
	return &this
}

// NewDocGenFormFieldResponseWithDefaults instantiates a new DocGenFormFieldResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocGenFormFieldResponseWithDefaults() *DocGenFormFieldResponse {
	this := DocGenFormFieldResponse{}
	return &this
}

// GetDocGenFormFields returns the DocGenFormFields field value if set, zero value otherwise.
func (o *DocGenFormFieldResponse) GetDocGenFormFields() []DocGenFormFields {
	if o == nil || IsNil(o.DocGenFormFields) {
		var ret []DocGenFormFields
		return ret
	}
	return o.DocGenFormFields
}

// GetDocGenFormFieldsOk returns a tuple with the DocGenFormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocGenFormFieldResponse) GetDocGenFormFieldsOk() ([]DocGenFormFields, bool) {
	if o == nil || IsNil(o.DocGenFormFields) {
		return nil, false
	}
	return o.DocGenFormFields, true
}

// HasDocGenFormFields returns a boolean if a field has been set.
func (o *DocGenFormFieldResponse) HasDocGenFormFields() bool {
	if o != nil && !IsNil(o.DocGenFormFields) {
		return true
	}

	return false
}

// SetDocGenFormFields gets a reference to the given []DocGenFormFields and assigns it to the DocGenFormFields field.
func (o *DocGenFormFieldResponse) SetDocGenFormFields(v []DocGenFormFields) {
	o.DocGenFormFields = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *DocGenFormFieldResponse) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocGenFormFieldResponse) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *DocGenFormFieldResponse) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *DocGenFormFieldResponse) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

func (o DocGenFormFieldResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocGenFormFieldResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocGenFormFields) {
		toSerialize["docGenFormFields"] = o.DocGenFormFields
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	return toSerialize, nil
}

type NullableDocGenFormFieldResponse struct {
	value *DocGenFormFieldResponse
	isSet bool
}

func (v NullableDocGenFormFieldResponse) Get() *DocGenFormFieldResponse {
	return v.value
}

func (v *NullableDocGenFormFieldResponse) Set(val *DocGenFormFieldResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocGenFormFieldResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocGenFormFieldResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocGenFormFieldResponse(val *DocGenFormFieldResponse) *NullableDocGenFormFieldResponse {
	return &NullableDocGenFormFieldResponse{value: val, isSet: true}
}

func (v NullableDocGenFormFieldResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocGenFormFieldResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


