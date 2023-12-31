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

// checks if the BillingInvoiceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInvoiceItem{}

// BillingInvoiceItem Contains information about an item on a billing invoice.
type BillingInvoiceItem struct {
	// Reserved for DocuSign.
	ChargeAmount *string `json:"chargeAmount,omitempty"`
	// Reserved for DocuSign.
	ChargeName *string `json:"chargeName,omitempty"`
	// Reserved for DocuSign.
	InvoiceItemId *string `json:"invoiceItemId,omitempty"`
	// The quantity of envelopes to add to the account.
	Quantity *string `json:"quantity,omitempty"`
	// 
	TaxAmount *string `json:"taxAmount,omitempty"`
	// 
	TaxExemptAmount *string `json:"taxExemptAmount,omitempty"`
	// Reserved for DocuSign.
	UnitPrice *string `json:"unitPrice,omitempty"`
}

// NewBillingInvoiceItem instantiates a new BillingInvoiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInvoiceItem() *BillingInvoiceItem {
	this := BillingInvoiceItem{}
	return &this
}

// NewBillingInvoiceItemWithDefaults instantiates a new BillingInvoiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInvoiceItemWithDefaults() *BillingInvoiceItem {
	this := BillingInvoiceItem{}
	return &this
}

// GetChargeAmount returns the ChargeAmount field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetChargeAmount() string {
	if o == nil || IsNil(o.ChargeAmount) {
		var ret string
		return ret
	}
	return *o.ChargeAmount
}

// GetChargeAmountOk returns a tuple with the ChargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetChargeAmountOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeAmount) {
		return nil, false
	}
	return o.ChargeAmount, true
}

// HasChargeAmount returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasChargeAmount() bool {
	if o != nil && !IsNil(o.ChargeAmount) {
		return true
	}

	return false
}

// SetChargeAmount gets a reference to the given string and assigns it to the ChargeAmount field.
func (o *BillingInvoiceItem) SetChargeAmount(v string) {
	o.ChargeAmount = &v
}

// GetChargeName returns the ChargeName field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetChargeName() string {
	if o == nil || IsNil(o.ChargeName) {
		var ret string
		return ret
	}
	return *o.ChargeName
}

// GetChargeNameOk returns a tuple with the ChargeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetChargeNameOk() (*string, bool) {
	if o == nil || IsNil(o.ChargeName) {
		return nil, false
	}
	return o.ChargeName, true
}

// HasChargeName returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasChargeName() bool {
	if o != nil && !IsNil(o.ChargeName) {
		return true
	}

	return false
}

// SetChargeName gets a reference to the given string and assigns it to the ChargeName field.
func (o *BillingInvoiceItem) SetChargeName(v string) {
	o.ChargeName = &v
}

// GetInvoiceItemId returns the InvoiceItemId field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetInvoiceItemId() string {
	if o == nil || IsNil(o.InvoiceItemId) {
		var ret string
		return ret
	}
	return *o.InvoiceItemId
}

// GetInvoiceItemIdOk returns a tuple with the InvoiceItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetInvoiceItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.InvoiceItemId) {
		return nil, false
	}
	return o.InvoiceItemId, true
}

// HasInvoiceItemId returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasInvoiceItemId() bool {
	if o != nil && !IsNil(o.InvoiceItemId) {
		return true
	}

	return false
}

// SetInvoiceItemId gets a reference to the given string and assigns it to the InvoiceItemId field.
func (o *BillingInvoiceItem) SetInvoiceItemId(v string) {
	o.InvoiceItemId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetQuantity() string {
	if o == nil || IsNil(o.Quantity) {
		var ret string
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetQuantityOk() (*string, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given string and assigns it to the Quantity field.
func (o *BillingInvoiceItem) SetQuantity(v string) {
	o.Quantity = &v
}

// GetTaxAmount returns the TaxAmount field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetTaxAmount() string {
	if o == nil || IsNil(o.TaxAmount) {
		var ret string
		return ret
	}
	return *o.TaxAmount
}

// GetTaxAmountOk returns a tuple with the TaxAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetTaxAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TaxAmount) {
		return nil, false
	}
	return o.TaxAmount, true
}

// HasTaxAmount returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasTaxAmount() bool {
	if o != nil && !IsNil(o.TaxAmount) {
		return true
	}

	return false
}

// SetTaxAmount gets a reference to the given string and assigns it to the TaxAmount field.
func (o *BillingInvoiceItem) SetTaxAmount(v string) {
	o.TaxAmount = &v
}

// GetTaxExemptAmount returns the TaxExemptAmount field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetTaxExemptAmount() string {
	if o == nil || IsNil(o.TaxExemptAmount) {
		var ret string
		return ret
	}
	return *o.TaxExemptAmount
}

// GetTaxExemptAmountOk returns a tuple with the TaxExemptAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetTaxExemptAmountOk() (*string, bool) {
	if o == nil || IsNil(o.TaxExemptAmount) {
		return nil, false
	}
	return o.TaxExemptAmount, true
}

// HasTaxExemptAmount returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasTaxExemptAmount() bool {
	if o != nil && !IsNil(o.TaxExemptAmount) {
		return true
	}

	return false
}

// SetTaxExemptAmount gets a reference to the given string and assigns it to the TaxExemptAmount field.
func (o *BillingInvoiceItem) SetTaxExemptAmount(v string) {
	o.TaxExemptAmount = &v
}

// GetUnitPrice returns the UnitPrice field value if set, zero value otherwise.
func (o *BillingInvoiceItem) GetUnitPrice() string {
	if o == nil || IsNil(o.UnitPrice) {
		var ret string
		return ret
	}
	return *o.UnitPrice
}

// GetUnitPriceOk returns a tuple with the UnitPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInvoiceItem) GetUnitPriceOk() (*string, bool) {
	if o == nil || IsNil(o.UnitPrice) {
		return nil, false
	}
	return o.UnitPrice, true
}

// HasUnitPrice returns a boolean if a field has been set.
func (o *BillingInvoiceItem) HasUnitPrice() bool {
	if o != nil && !IsNil(o.UnitPrice) {
		return true
	}

	return false
}

// SetUnitPrice gets a reference to the given string and assigns it to the UnitPrice field.
func (o *BillingInvoiceItem) SetUnitPrice(v string) {
	o.UnitPrice = &v
}

func (o BillingInvoiceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInvoiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChargeAmount) {
		toSerialize["chargeAmount"] = o.ChargeAmount
	}
	if !IsNil(o.ChargeName) {
		toSerialize["chargeName"] = o.ChargeName
	}
	if !IsNil(o.InvoiceItemId) {
		toSerialize["invoiceItemId"] = o.InvoiceItemId
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.TaxAmount) {
		toSerialize["taxAmount"] = o.TaxAmount
	}
	if !IsNil(o.TaxExemptAmount) {
		toSerialize["taxExemptAmount"] = o.TaxExemptAmount
	}
	if !IsNil(o.UnitPrice) {
		toSerialize["unitPrice"] = o.UnitPrice
	}
	return toSerialize, nil
}

type NullableBillingInvoiceItem struct {
	value *BillingInvoiceItem
	isSet bool
}

func (v NullableBillingInvoiceItem) Get() *BillingInvoiceItem {
	return v.value
}

func (v *NullableBillingInvoiceItem) Set(val *BillingInvoiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInvoiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInvoiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInvoiceItem(val *BillingInvoiceItem) *NullableBillingInvoiceItem {
	return &NullableBillingInvoiceItem{value: val, isSet: true}
}

func (v NullableBillingInvoiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInvoiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


