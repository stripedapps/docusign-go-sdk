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

// checks if the AppStoreReceipt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReceipt{}

// AppStoreReceipt Contains information about an APP store receipt.
type AppStoreReceipt struct {
	// 
	DowngradeProductId *string `json:"downgradeProductId,omitempty"`
	// 
	IsDowngradeCancellation *string `json:"isDowngradeCancellation,omitempty"`
	// The Product ID from the AppStore.
	ProductId *string `json:"productId,omitempty"`
	// Reserved for DocuSign.
	ReceiptData *string `json:"receiptData,omitempty"`
}

// NewAppStoreReceipt instantiates a new AppStoreReceipt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReceipt() *AppStoreReceipt {
	this := AppStoreReceipt{}
	return &this
}

// NewAppStoreReceiptWithDefaults instantiates a new AppStoreReceipt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReceiptWithDefaults() *AppStoreReceipt {
	this := AppStoreReceipt{}
	return &this
}

// GetDowngradeProductId returns the DowngradeProductId field value if set, zero value otherwise.
func (o *AppStoreReceipt) GetDowngradeProductId() string {
	if o == nil || IsNil(o.DowngradeProductId) {
		var ret string
		return ret
	}
	return *o.DowngradeProductId
}

// GetDowngradeProductIdOk returns a tuple with the DowngradeProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReceipt) GetDowngradeProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.DowngradeProductId) {
		return nil, false
	}
	return o.DowngradeProductId, true
}

// HasDowngradeProductId returns a boolean if a field has been set.
func (o *AppStoreReceipt) HasDowngradeProductId() bool {
	if o != nil && !IsNil(o.DowngradeProductId) {
		return true
	}

	return false
}

// SetDowngradeProductId gets a reference to the given string and assigns it to the DowngradeProductId field.
func (o *AppStoreReceipt) SetDowngradeProductId(v string) {
	o.DowngradeProductId = &v
}

// GetIsDowngradeCancellation returns the IsDowngradeCancellation field value if set, zero value otherwise.
func (o *AppStoreReceipt) GetIsDowngradeCancellation() string {
	if o == nil || IsNil(o.IsDowngradeCancellation) {
		var ret string
		return ret
	}
	return *o.IsDowngradeCancellation
}

// GetIsDowngradeCancellationOk returns a tuple with the IsDowngradeCancellation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReceipt) GetIsDowngradeCancellationOk() (*string, bool) {
	if o == nil || IsNil(o.IsDowngradeCancellation) {
		return nil, false
	}
	return o.IsDowngradeCancellation, true
}

// HasIsDowngradeCancellation returns a boolean if a field has been set.
func (o *AppStoreReceipt) HasIsDowngradeCancellation() bool {
	if o != nil && !IsNil(o.IsDowngradeCancellation) {
		return true
	}

	return false
}

// SetIsDowngradeCancellation gets a reference to the given string and assigns it to the IsDowngradeCancellation field.
func (o *AppStoreReceipt) SetIsDowngradeCancellation(v string) {
	o.IsDowngradeCancellation = &v
}

// GetProductId returns the ProductId field value if set, zero value otherwise.
func (o *AppStoreReceipt) GetProductId() string {
	if o == nil || IsNil(o.ProductId) {
		var ret string
		return ret
	}
	return *o.ProductId
}

// GetProductIdOk returns a tuple with the ProductId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReceipt) GetProductIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProductId) {
		return nil, false
	}
	return o.ProductId, true
}

// HasProductId returns a boolean if a field has been set.
func (o *AppStoreReceipt) HasProductId() bool {
	if o != nil && !IsNil(o.ProductId) {
		return true
	}

	return false
}

// SetProductId gets a reference to the given string and assigns it to the ProductId field.
func (o *AppStoreReceipt) SetProductId(v string) {
	o.ProductId = &v
}

// GetReceiptData returns the ReceiptData field value if set, zero value otherwise.
func (o *AppStoreReceipt) GetReceiptData() string {
	if o == nil || IsNil(o.ReceiptData) {
		var ret string
		return ret
	}
	return *o.ReceiptData
}

// GetReceiptDataOk returns a tuple with the ReceiptData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreReceipt) GetReceiptDataOk() (*string, bool) {
	if o == nil || IsNil(o.ReceiptData) {
		return nil, false
	}
	return o.ReceiptData, true
}

// HasReceiptData returns a boolean if a field has been set.
func (o *AppStoreReceipt) HasReceiptData() bool {
	if o != nil && !IsNil(o.ReceiptData) {
		return true
	}

	return false
}

// SetReceiptData gets a reference to the given string and assigns it to the ReceiptData field.
func (o *AppStoreReceipt) SetReceiptData(v string) {
	o.ReceiptData = &v
}

func (o AppStoreReceipt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReceipt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DowngradeProductId) {
		toSerialize["downgradeProductId"] = o.DowngradeProductId
	}
	if !IsNil(o.IsDowngradeCancellation) {
		toSerialize["isDowngradeCancellation"] = o.IsDowngradeCancellation
	}
	if !IsNil(o.ProductId) {
		toSerialize["productId"] = o.ProductId
	}
	if !IsNil(o.ReceiptData) {
		toSerialize["receiptData"] = o.ReceiptData
	}
	return toSerialize, nil
}

type NullableAppStoreReceipt struct {
	value *AppStoreReceipt
	isSet bool
}

func (v NullableAppStoreReceipt) Get() *AppStoreReceipt {
	return v.value
}

func (v *NullableAppStoreReceipt) Set(val *AppStoreReceipt) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReceipt) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReceipt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReceipt(val *AppStoreReceipt) *NullableAppStoreReceipt {
	return &NullableAppStoreReceipt{value: val, isSet: true}
}

func (v NullableAppStoreReceipt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReceipt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


