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

// checks if the BillingInvoicesSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInvoicesSummary{}

// BillingInvoicesSummary 
type BillingInvoicesSummary struct {
	// 
	AccountBalance *string `json:"accountBalance,omitempty"`
	// Reserved for DocuSign.
	BillingInvoices []BillingInvoice `json:"billingInvoices,omitempty"`
	// The [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) currency code. 
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// 
	PastDueBalance *string `json:"pastDueBalance,omitempty"`
	// 
	PaymentAllowed *string `json:"paymentAllowed,omitempty"`
}

// NewBillingInvoicesSummary instantiates a new BillingInvoicesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInvoicesSummary() *BillingInvoicesSummary {
	this := BillingInvoicesSummary{}
	return &this
}

// NewBillingInvoicesSummaryWithDefaults instantiates a new BillingInvoicesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInvoicesSummaryWithDefaults() *BillingInvoicesSummary {
	this := BillingInvoicesSummary{}
	return &this
}

// GetAccountBalance returns the AccountBalance field value if set, zero value otherwise.
func (o *BillingInvoicesSummary) GetAccountBalance() string {
	if o == nil || IsNil(o.AccountBalance) {
		var ret string
		return ret
	}
	return *o.AccountBalance
}

// GetAccountBalanceOk returns a tuple with the AccountBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoicesSummary) GetAccountBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.AccountBalance) {
		return nil, false
	}
	return o.AccountBalance, true
}

// HasAccountBalance returns a boolean if a field has been set.
func (o *BillingInvoicesSummary) HasAccountBalance() bool {
	if o != nil && !IsNil(o.AccountBalance) {
		return true
	}

	return false
}

// SetAccountBalance gets a reference to the given string and assigns it to the AccountBalance field.
func (o *BillingInvoicesSummary) SetAccountBalance(v string) {
	o.AccountBalance = &v
}

// GetBillingInvoices returns the BillingInvoices field value if set, zero value otherwise.
func (o *BillingInvoicesSummary) GetBillingInvoices() []BillingInvoice {
	if o == nil || IsNil(o.BillingInvoices) {
		var ret []BillingInvoice
		return ret
	}
	return o.BillingInvoices
}

// GetBillingInvoicesOk returns a tuple with the BillingInvoices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoicesSummary) GetBillingInvoicesOk() ([]BillingInvoice, bool) {
	if o == nil || IsNil(o.BillingInvoices) {
		return nil, false
	}
	return o.BillingInvoices, true
}

// HasBillingInvoices returns a boolean if a field has been set.
func (o *BillingInvoicesSummary) HasBillingInvoices() bool {
	if o != nil && !IsNil(o.BillingInvoices) {
		return true
	}

	return false
}

// SetBillingInvoices gets a reference to the given []BillingInvoice and assigns it to the BillingInvoices field.
func (o *BillingInvoicesSummary) SetBillingInvoices(v []BillingInvoice) {
	o.BillingInvoices = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *BillingInvoicesSummary) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoicesSummary) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *BillingInvoicesSummary) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *BillingInvoicesSummary) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetPastDueBalance returns the PastDueBalance field value if set, zero value otherwise.
func (o *BillingInvoicesSummary) GetPastDueBalance() string {
	if o == nil || IsNil(o.PastDueBalance) {
		var ret string
		return ret
	}
	return *o.PastDueBalance
}

// GetPastDueBalanceOk returns a tuple with the PastDueBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoicesSummary) GetPastDueBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.PastDueBalance) {
		return nil, false
	}
	return o.PastDueBalance, true
}

// HasPastDueBalance returns a boolean if a field has been set.
func (o *BillingInvoicesSummary) HasPastDueBalance() bool {
	if o != nil && !IsNil(o.PastDueBalance) {
		return true
	}

	return false
}

// SetPastDueBalance gets a reference to the given string and assigns it to the PastDueBalance field.
func (o *BillingInvoicesSummary) SetPastDueBalance(v string) {
	o.PastDueBalance = &v
}

// GetPaymentAllowed returns the PaymentAllowed field value if set, zero value otherwise.
func (o *BillingInvoicesSummary) GetPaymentAllowed() string {
	if o == nil || IsNil(o.PaymentAllowed) {
		var ret string
		return ret
	}
	return *o.PaymentAllowed
}

// GetPaymentAllowedOk returns a tuple with the PaymentAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoicesSummary) GetPaymentAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentAllowed) {
		return nil, false
	}
	return o.PaymentAllowed, true
}

// HasPaymentAllowed returns a boolean if a field has been set.
func (o *BillingInvoicesSummary) HasPaymentAllowed() bool {
	if o != nil && !IsNil(o.PaymentAllowed) {
		return true
	}

	return false
}

// SetPaymentAllowed gets a reference to the given string and assigns it to the PaymentAllowed field.
func (o *BillingInvoicesSummary) SetPaymentAllowed(v string) {
	o.PaymentAllowed = &v
}

func (o BillingInvoicesSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInvoicesSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountBalance) {
		toSerialize["accountBalance"] = o.AccountBalance
	}
	if !IsNil(o.BillingInvoices) {
		toSerialize["billingInvoices"] = o.BillingInvoices
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.PastDueBalance) {
		toSerialize["pastDueBalance"] = o.PastDueBalance
	}
	if !IsNil(o.PaymentAllowed) {
		toSerialize["paymentAllowed"] = o.PaymentAllowed
	}
	return toSerialize, nil
}

type NullableBillingInvoicesSummary struct {
	value *BillingInvoicesSummary
	isSet bool
}

func (v NullableBillingInvoicesSummary) Get() *BillingInvoicesSummary {
	return v.value
}

func (v *NullableBillingInvoicesSummary) Set(val *BillingInvoicesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInvoicesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInvoicesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInvoicesSummary(val *BillingInvoicesSummary) *NullableBillingInvoicesSummary {
	return &NullableBillingInvoicesSummary{value: val, isSet: true}
}

func (v NullableBillingInvoicesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInvoicesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


