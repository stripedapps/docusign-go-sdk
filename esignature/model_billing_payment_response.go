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

// checks if the BillingPaymentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPaymentResponse{}

// BillingPaymentResponse Defines an billing payment response object.
type BillingPaymentResponse struct {
	// Reserved for DocuSign.
	BillingPayments []BillingPayment `json:"billingPayments,omitempty"`
}

// NewBillingPaymentResponse instantiates a new BillingPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPaymentResponse() *BillingPaymentResponse {
	this := BillingPaymentResponse{}
	return &this
}

// NewBillingPaymentResponseWithDefaults instantiates a new BillingPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPaymentResponseWithDefaults() *BillingPaymentResponse {
	this := BillingPaymentResponse{}
	return &this
}

// GetBillingPayments returns the BillingPayments field value if set, zero value otherwise.
func (o *BillingPaymentResponse) GetBillingPayments() []BillingPayment {
	if o == nil || IsNil(o.BillingPayments) {
		var ret []BillingPayment
		return ret
	}
	return o.BillingPayments
}

// GetBillingPaymentsOk returns a tuple with the BillingPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPaymentResponse) GetBillingPaymentsOk() ([]BillingPayment, bool) {
	if o == nil || IsNil(o.BillingPayments) {
		return nil, false
	}
	return o.BillingPayments, true
}

// HasBillingPayments returns a boolean if a field has been set.
func (o *BillingPaymentResponse) HasBillingPayments() bool {
	if o != nil && !IsNil(o.BillingPayments) {
		return true
	}

	return false
}

// SetBillingPayments gets a reference to the given []BillingPayment and assigns it to the BillingPayments field.
func (o *BillingPaymentResponse) SetBillingPayments(v []BillingPayment) {
	o.BillingPayments = v
}

func (o BillingPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPaymentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingPayments) {
		toSerialize["billingPayments"] = o.BillingPayments
	}
	return toSerialize, nil
}

type NullableBillingPaymentResponse struct {
	value *BillingPaymentResponse
	isSet bool
}

func (v NullableBillingPaymentResponse) Get() *BillingPaymentResponse {
	return v.value
}

func (v *NullableBillingPaymentResponse) Set(val *BillingPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPaymentResponse(val *BillingPaymentResponse) *NullableBillingPaymentResponse {
	return &NullableBillingPaymentResponse{value: val, isSet: true}
}

func (v NullableBillingPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


