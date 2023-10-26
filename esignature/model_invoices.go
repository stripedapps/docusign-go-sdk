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

// checks if the Invoices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invoices{}

// Invoices Invoices
type Invoices struct {
	// Reserved for DocuSign.
	Amount *string `json:"amount,omitempty"`
	// Reserved for DocuSign.
	Balance *string `json:"balance,omitempty"`
	// Reserved for DocuSign.
	DueDate *string `json:"dueDate,omitempty"`
	// Reserved for DocuSign.
	InvoiceId *string `json:"invoiceId,omitempty"`
	// Reserved for DocuSign.
	InvoiceItems []BillingInvoiceItem `json:"invoiceItems,omitempty"`
	// Reserved for DocuSign.
	InvoiceNumber *string `json:"invoiceNumber,omitempty"`
	// Contains a URI for an endpoint that you can use to retrieve invoice information.
	InvoiceUri *string `json:"invoiceUri,omitempty"`
	// 
	NonTaxableAmount *string `json:"nonTaxableAmount,omitempty"`
	// 
	PdfAvailable *string `json:"pdfAvailable,omitempty"`
	// 
	TaxableAmount *string `json:"taxableAmount,omitempty"`
}

// NewInvoices instantiates a new Invoices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvoices() *Invoices {
	this := Invoices{}
	return &this
}

// NewInvoicesWithDefaults instantiates a new Invoices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvoicesWithDefaults() *Invoices {
	this := Invoices{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *Invoices) GetAmount() string {
	if o == nil || IsNil(o.Amount) {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetAmountOk() (*string, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *Invoices) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *Invoices) SetAmount(v string) {
	o.Amount = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *Invoices) GetBalance() string {
	if o == nil || IsNil(o.Balance) {
		var ret string
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetBalanceOk() (*string, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *Invoices) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given string and assigns it to the Balance field.
func (o *Invoices) SetBalance(v string) {
	o.Balance = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *Invoices) GetDueDate() string {
	if o == nil || IsNil(o.DueDate) {
		var ret string
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetDueDateOk() (*string, bool) {
	if o == nil || IsNil(o.DueDate) {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *Invoices) HasDueDate() bool {
	if o != nil && !IsNil(o.DueDate) {
		return true
	}

	return false
}

// SetDueDate gets a reference to the given string and assigns it to the DueDate field.
func (o *Invoices) SetDueDate(v string) {
	o.DueDate = &v
}

// GetInvoiceId returns the InvoiceId field value if set, zero value otherwise.
func (o *Invoices) GetInvoiceId() string {
	if o == nil || IsNil(o.InvoiceId) {
		var ret string
		return ret
	}
	return *o.InvoiceId
}

// GetInvoiceIdOk returns a tuple with the InvoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetInvoiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceId) {
		return nil, false
	}
	return o.InvoiceId, true
}

// HasInvoiceId returns a boolean if a field has been set.
func (o *Invoices) HasInvoiceId() bool {
	if o != nil && !IsNil(o.InvoiceId) {
		return true
	}

	return false
}

// SetInvoiceId gets a reference to the given string and assigns it to the InvoiceId field.
func (o *Invoices) SetInvoiceId(v string) {
	o.InvoiceId = &v
}

// GetInvoiceItems returns the InvoiceItems field value if set, zero value otherwise.
func (o *Invoices) GetInvoiceItems() []BillingInvoiceItem {
	if o == nil || IsNil(o.InvoiceItems) {
		var ret []BillingInvoiceItem
		return ret
	}
	return o.InvoiceItems
}

// GetInvoiceItemsOk returns a tuple with the InvoiceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetInvoiceItemsOk() ([]BillingInvoiceItem, bool) {
	if o == nil || IsNil(o.InvoiceItems) {
		return nil, false
	}
	return o.InvoiceItems, true
}

// HasInvoiceItems returns a boolean if a field has been set.
func (o *Invoices) HasInvoiceItems() bool {
	if o != nil && !IsNil(o.InvoiceItems) {
		return true
	}

	return false
}

// SetInvoiceItems gets a reference to the given []BillingInvoiceItem and assigns it to the InvoiceItems field.
func (o *Invoices) SetInvoiceItems(v []BillingInvoiceItem) {
	o.InvoiceItems = v
}

// GetInvoiceNumber returns the InvoiceNumber field value if set, zero value otherwise.
func (o *Invoices) GetInvoiceNumber() string {
	if o == nil || IsNil(o.InvoiceNumber) {
		var ret string
		return ret
	}
	return *o.InvoiceNumber
}

// GetInvoiceNumberOk returns a tuple with the InvoiceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetInvoiceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceNumber) {
		return nil, false
	}
	return o.InvoiceNumber, true
}

// HasInvoiceNumber returns a boolean if a field has been set.
func (o *Invoices) HasInvoiceNumber() bool {
	if o != nil && !IsNil(o.InvoiceNumber) {
		return true
	}

	return false
}

// SetInvoiceNumber gets a reference to the given string and assigns it to the InvoiceNumber field.
func (o *Invoices) SetInvoiceNumber(v string) {
	o.InvoiceNumber = &v
}

// GetInvoiceUri returns the InvoiceUri field value if set, zero value otherwise.
func (o *Invoices) GetInvoiceUri() string {
	if o == nil || IsNil(o.InvoiceUri) {
		var ret string
		return ret
	}
	return *o.InvoiceUri
}

// GetInvoiceUriOk returns a tuple with the InvoiceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetInvoiceUriOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceUri) {
		return nil, false
	}
	return o.InvoiceUri, true
}

// HasInvoiceUri returns a boolean if a field has been set.
func (o *Invoices) HasInvoiceUri() bool {
	if o != nil && !IsNil(o.InvoiceUri) {
		return true
	}

	return false
}

// SetInvoiceUri gets a reference to the given string and assigns it to the InvoiceUri field.
func (o *Invoices) SetInvoiceUri(v string) {
	o.InvoiceUri = &v
}

// GetNonTaxableAmount returns the NonTaxableAmount field value if set, zero value otherwise.
func (o *Invoices) GetNonTaxableAmount() string {
	if o == nil || IsNil(o.NonTaxableAmount) {
		var ret string
		return ret
	}
	return *o.NonTaxableAmount
}

// GetNonTaxableAmountOk returns a tuple with the NonTaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetNonTaxableAmountOk() (*string, bool) {
	if o == nil || IsNil(o.NonTaxableAmount) {
		return nil, false
	}
	return o.NonTaxableAmount, true
}

// HasNonTaxableAmount returns a boolean if a field has been set.
func (o *Invoices) HasNonTaxableAmount() bool {
	if o != nil && !IsNil(o.NonTaxableAmount) {
		return true
	}

	return false
}

// SetNonTaxableAmount gets a reference to the given string and assigns it to the NonTaxableAmount field.
func (o *Invoices) SetNonTaxableAmount(v string) {
	o.NonTaxableAmount = &v
}

// GetPdfAvailable returns the PdfAvailable field value if set, zero value otherwise.
func (o *Invoices) GetPdfAvailable() string {
	if o == nil || IsNil(o.PdfAvailable) {
		var ret string
		return ret
	}
	return *o.PdfAvailable
}

// GetPdfAvailableOk returns a tuple with the PdfAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetPdfAvailableOk() (*string, bool) {
	if o == nil || IsNil(o.PdfAvailable) {
		return nil, false
	}
	return o.PdfAvailable, true
}

// HasPdfAvailable returns a boolean if a field has been set.
func (o *Invoices) HasPdfAvailable() bool {
	if o != nil && !IsNil(o.PdfAvailable) {
		return true
	}

	return false
}

// SetPdfAvailable gets a reference to the given string and assigns it to the PdfAvailable field.
func (o *Invoices) SetPdfAvailable(v string) {
	o.PdfAvailable = &v
}

// GetTaxableAmount returns the TaxableAmount field value if set, zero value otherwise.
func (o *Invoices) GetTaxableAmount() string {
	if o == nil || IsNil(o.TaxableAmount) {
		var ret string
		return ret
	}
	return *o.TaxableAmount
}

// GetTaxableAmountOk returns a tuple with the TaxableAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invoices) GetTaxableAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TaxableAmount) {
		return nil, false
	}
	return o.TaxableAmount, true
}

// HasTaxableAmount returns a boolean if a field has been set.
func (o *Invoices) HasTaxableAmount() bool {
	if o != nil && !IsNil(o.TaxableAmount) {
		return true
	}

	return false
}

// SetTaxableAmount gets a reference to the given string and assigns it to the TaxableAmount field.
func (o *Invoices) SetTaxableAmount(v string) {
	o.TaxableAmount = &v
}

func (o Invoices) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invoices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.DueDate) {
		toSerialize["dueDate"] = o.DueDate
	}
	if !IsNil(o.InvoiceId) {
		toSerialize["invoiceId"] = o.InvoiceId
	}
	if !IsNil(o.InvoiceItems) {
		toSerialize["invoiceItems"] = o.InvoiceItems
	}
	if !IsNil(o.InvoiceNumber) {
		toSerialize["invoiceNumber"] = o.InvoiceNumber
	}
	if !IsNil(o.InvoiceUri) {
		toSerialize["invoiceUri"] = o.InvoiceUri
	}
	if !IsNil(o.NonTaxableAmount) {
		toSerialize["nonTaxableAmount"] = o.NonTaxableAmount
	}
	if !IsNil(o.PdfAvailable) {
		toSerialize["pdfAvailable"] = o.PdfAvailable
	}
	if !IsNil(o.TaxableAmount) {
		toSerialize["taxableAmount"] = o.TaxableAmount
	}
	return toSerialize, nil
}

type NullableInvoices struct {
	value *Invoices
	isSet bool
}

func (v NullableInvoices) Get() *Invoices {
	return v.value
}

func (v *NullableInvoices) Set(val *Invoices) {
	v.value = val
	v.isSet = true
}

func (v NullableInvoices) IsSet() bool {
	return v.isSet
}

func (v *NullableInvoices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvoices(val *Invoices) *NullableInvoices {
	return &NullableInvoices{value: val, isSet: true}
}

func (v NullableInvoices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvoices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


