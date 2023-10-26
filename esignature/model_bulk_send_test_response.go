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

// checks if the BulkSendTestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkSendTestResponse{}

// BulkSendTestResponse This object contains the results of a bulk send test.
type BulkSendTestResponse struct {
	// When **true,** the envelope or template is compatible with the bulk send list and can be sent by using the [BulkSend: createBulkSendRequest][BulkSendRequest] method.  **Note:** This property is only returned in responses and ignored in requests.  [BulkSendRequest]:  /docs/esign-rest-api/reference/bulkenvelopes/bulksend/createbulksendrequest/ 
	CanBeSent *bool `json:"canBeSent,omitempty"`
	// Human-readable details about any validation errors that occurred.
	ValidationErrorDetails []string `json:"validationErrorDetails,omitempty"`
	// A list of validation errors that were encountered during the bulk send test.  **Note:** This information is intended to be parsed by machine.
	ValidationErrors []string `json:"validationErrors,omitempty"`
}

// NewBulkSendTestResponse instantiates a new BulkSendTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSendTestResponse() *BulkSendTestResponse {
	this := BulkSendTestResponse{}
	return &this
}

// NewBulkSendTestResponseWithDefaults instantiates a new BulkSendTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSendTestResponseWithDefaults() *BulkSendTestResponse {
	this := BulkSendTestResponse{}
	return &this
}

// GetCanBeSent returns the CanBeSent field value if set, zero value otherwise.
func (o *BulkSendTestResponse) GetCanBeSent() bool {
	if o == nil || IsNil(o.CanBeSent) {
		var ret bool
		return ret
	}
	return *o.CanBeSent
}

// GetCanBeSentOk returns a tuple with the CanBeSent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendTestResponse) GetCanBeSentOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeSent) {
		return nil, false
	}
	return o.CanBeSent, true
}

// HasCanBeSent returns a boolean if a field has been set.
func (o *BulkSendTestResponse) HasCanBeSent() bool {
	if o != nil && !IsNil(o.CanBeSent) {
		return true
	}

	return false
}

// SetCanBeSent gets a reference to the given bool and assigns it to the CanBeSent field.
func (o *BulkSendTestResponse) SetCanBeSent(v bool) {
	o.CanBeSent = &v
}

// GetValidationErrorDetails returns the ValidationErrorDetails field value if set, zero value otherwise.
func (o *BulkSendTestResponse) GetValidationErrorDetails() []string {
	if o == nil || IsNil(o.ValidationErrorDetails) {
		var ret []string
		return ret
	}
	return o.ValidationErrorDetails
}

// GetValidationErrorDetailsOk returns a tuple with the ValidationErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendTestResponse) GetValidationErrorDetailsOk() ([]string, bool) {
	if o == nil || IsNil(o.ValidationErrorDetails) {
		return nil, false
	}
	return o.ValidationErrorDetails, true
}

// HasValidationErrorDetails returns a boolean if a field has been set.
func (o *BulkSendTestResponse) HasValidationErrorDetails() bool {
	if o != nil && !IsNil(o.ValidationErrorDetails) {
		return true
	}

	return false
}

// SetValidationErrorDetails gets a reference to the given []string and assigns it to the ValidationErrorDetails field.
func (o *BulkSendTestResponse) SetValidationErrorDetails(v []string) {
	o.ValidationErrorDetails = v
}

// GetValidationErrors returns the ValidationErrors field value if set, zero value otherwise.
func (o *BulkSendTestResponse) GetValidationErrors() []string {
	if o == nil || IsNil(o.ValidationErrors) {
		var ret []string
		return ret
	}
	return o.ValidationErrors
}

// GetValidationErrorsOk returns a tuple with the ValidationErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendTestResponse) GetValidationErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.ValidationErrors) {
		return nil, false
	}
	return o.ValidationErrors, true
}

// HasValidationErrors returns a boolean if a field has been set.
func (o *BulkSendTestResponse) HasValidationErrors() bool {
	if o != nil && !IsNil(o.ValidationErrors) {
		return true
	}

	return false
}

// SetValidationErrors gets a reference to the given []string and assigns it to the ValidationErrors field.
func (o *BulkSendTestResponse) SetValidationErrors(v []string) {
	o.ValidationErrors = v
}

func (o BulkSendTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkSendTestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanBeSent) {
		toSerialize["canBeSent"] = o.CanBeSent
	}
	if !IsNil(o.ValidationErrorDetails) {
		toSerialize["validationErrorDetails"] = o.ValidationErrorDetails
	}
	if !IsNil(o.ValidationErrors) {
		toSerialize["validationErrors"] = o.ValidationErrors
	}
	return toSerialize, nil
}

type NullableBulkSendTestResponse struct {
	value *BulkSendTestResponse
	isSet bool
}

func (v NullableBulkSendTestResponse) Get() *BulkSendTestResponse {
	return v.value
}

func (v *NullableBulkSendTestResponse) Set(val *BulkSendTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSendTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSendTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSendTestResponse(val *BulkSendTestResponse) *NullableBulkSendTestResponse {
	return &NullableBulkSendTestResponse{value: val, isSet: true}
}

func (v NullableBulkSendTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSendTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

