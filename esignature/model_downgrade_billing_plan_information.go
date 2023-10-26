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

// checks if the DowngradeBillingPlanInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DowngradeBillingPlanInformation{}

// DowngradeBillingPlanInformation 
type DowngradeBillingPlanInformation struct {
	// 
	DowngradeEventType *string `json:"downgradeEventType,omitempty"`
	PlanInformation *PlanInformation `json:"planInformation,omitempty"`
	// 
	PromoCode *string `json:"promoCode,omitempty"`
	// 
	SaleDiscount *string `json:"saleDiscount,omitempty"`
	// Reserved for DocuSign.
	SaleDiscountPeriods *string `json:"saleDiscountPeriods,omitempty"`
	// 
	SaleDiscountType *string `json:"saleDiscountType,omitempty"`
}

// NewDowngradeBillingPlanInformation instantiates a new DowngradeBillingPlanInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDowngradeBillingPlanInformation() *DowngradeBillingPlanInformation {
	this := DowngradeBillingPlanInformation{}
	return &this
}

// NewDowngradeBillingPlanInformationWithDefaults instantiates a new DowngradeBillingPlanInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDowngradeBillingPlanInformationWithDefaults() *DowngradeBillingPlanInformation {
	this := DowngradeBillingPlanInformation{}
	return &this
}

// GetDowngradeEventType returns the DowngradeEventType field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetDowngradeEventType() string {
	if o == nil || IsNil(o.DowngradeEventType) {
		var ret string
		return ret
	}
	return *o.DowngradeEventType
}

// GetDowngradeEventTypeOk returns a tuple with the DowngradeEventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetDowngradeEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DowngradeEventType) {
		return nil, false
	}
	return o.DowngradeEventType, true
}

// HasDowngradeEventType returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasDowngradeEventType() bool {
	if o != nil && !IsNil(o.DowngradeEventType) {
		return true
	}

	return false
}

// SetDowngradeEventType gets a reference to the given string and assigns it to the DowngradeEventType field.
func (o *DowngradeBillingPlanInformation) SetDowngradeEventType(v string) {
	o.DowngradeEventType = &v
}

// GetPlanInformation returns the PlanInformation field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetPlanInformation() PlanInformation {
	if o == nil || IsNil(o.PlanInformation) {
		var ret PlanInformation
		return ret
	}
	return *o.PlanInformation
}

// GetPlanInformationOk returns a tuple with the PlanInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetPlanInformationOk() (*PlanInformation, bool) {
	if o == nil || IsNil(o.PlanInformation) {
		return nil, false
	}
	return o.PlanInformation, true
}

// HasPlanInformation returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasPlanInformation() bool {
	if o != nil && !IsNil(o.PlanInformation) {
		return true
	}

	return false
}

// SetPlanInformation gets a reference to the given PlanInformation and assigns it to the PlanInformation field.
func (o *DowngradeBillingPlanInformation) SetPlanInformation(v PlanInformation) {
	o.PlanInformation = &v
}

// GetPromoCode returns the PromoCode field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetPromoCode() string {
	if o == nil || IsNil(o.PromoCode) {
		var ret string
		return ret
	}
	return *o.PromoCode
}

// GetPromoCodeOk returns a tuple with the PromoCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetPromoCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PromoCode) {
		return nil, false
	}
	return o.PromoCode, true
}

// HasPromoCode returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasPromoCode() bool {
	if o != nil && !IsNil(o.PromoCode) {
		return true
	}

	return false
}

// SetPromoCode gets a reference to the given string and assigns it to the PromoCode field.
func (o *DowngradeBillingPlanInformation) SetPromoCode(v string) {
	o.PromoCode = &v
}

// GetSaleDiscount returns the SaleDiscount field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetSaleDiscount() string {
	if o == nil || IsNil(o.SaleDiscount) {
		var ret string
		return ret
	}
	return *o.SaleDiscount
}

// GetSaleDiscountOk returns a tuple with the SaleDiscount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetSaleDiscountOk() (*string, bool) {
	if o == nil || IsNil(o.SaleDiscount) {
		return nil, false
	}
	return o.SaleDiscount, true
}

// HasSaleDiscount returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasSaleDiscount() bool {
	if o != nil && !IsNil(o.SaleDiscount) {
		return true
	}

	return false
}

// SetSaleDiscount gets a reference to the given string and assigns it to the SaleDiscount field.
func (o *DowngradeBillingPlanInformation) SetSaleDiscount(v string) {
	o.SaleDiscount = &v
}

// GetSaleDiscountPeriods returns the SaleDiscountPeriods field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetSaleDiscountPeriods() string {
	if o == nil || IsNil(o.SaleDiscountPeriods) {
		var ret string
		return ret
	}
	return *o.SaleDiscountPeriods
}

// GetSaleDiscountPeriodsOk returns a tuple with the SaleDiscountPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetSaleDiscountPeriodsOk() (*string, bool) {
	if o == nil || IsNil(o.SaleDiscountPeriods) {
		return nil, false
	}
	return o.SaleDiscountPeriods, true
}

// HasSaleDiscountPeriods returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasSaleDiscountPeriods() bool {
	if o != nil && !IsNil(o.SaleDiscountPeriods) {
		return true
	}

	return false
}

// SetSaleDiscountPeriods gets a reference to the given string and assigns it to the SaleDiscountPeriods field.
func (o *DowngradeBillingPlanInformation) SetSaleDiscountPeriods(v string) {
	o.SaleDiscountPeriods = &v
}

// GetSaleDiscountType returns the SaleDiscountType field value if set, zero value otherwise.
func (o *DowngradeBillingPlanInformation) GetSaleDiscountType() string {
	if o == nil || IsNil(o.SaleDiscountType) {
		var ret string
		return ret
	}
	return *o.SaleDiscountType
}

// GetSaleDiscountTypeOk returns a tuple with the SaleDiscountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DowngradeBillingPlanInformation) GetSaleDiscountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SaleDiscountType) {
		return nil, false
	}
	return o.SaleDiscountType, true
}

// HasSaleDiscountType returns a boolean if a field has been set.
func (o *DowngradeBillingPlanInformation) HasSaleDiscountType() bool {
	if o != nil && !IsNil(o.SaleDiscountType) {
		return true
	}

	return false
}

// SetSaleDiscountType gets a reference to the given string and assigns it to the SaleDiscountType field.
func (o *DowngradeBillingPlanInformation) SetSaleDiscountType(v string) {
	o.SaleDiscountType = &v
}

func (o DowngradeBillingPlanInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DowngradeBillingPlanInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DowngradeEventType) {
		toSerialize["downgradeEventType"] = o.DowngradeEventType
	}
	if !IsNil(o.PlanInformation) {
		toSerialize["planInformation"] = o.PlanInformation
	}
	if !IsNil(o.PromoCode) {
		toSerialize["promoCode"] = o.PromoCode
	}
	if !IsNil(o.SaleDiscount) {
		toSerialize["saleDiscount"] = o.SaleDiscount
	}
	if !IsNil(o.SaleDiscountPeriods) {
		toSerialize["saleDiscountPeriods"] = o.SaleDiscountPeriods
	}
	if !IsNil(o.SaleDiscountType) {
		toSerialize["saleDiscountType"] = o.SaleDiscountType
	}
	return toSerialize, nil
}

type NullableDowngradeBillingPlanInformation struct {
	value *DowngradeBillingPlanInformation
	isSet bool
}

func (v NullableDowngradeBillingPlanInformation) Get() *DowngradeBillingPlanInformation {
	return v.value
}

func (v *NullableDowngradeBillingPlanInformation) Set(val *DowngradeBillingPlanInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDowngradeBillingPlanInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDowngradeBillingPlanInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDowngradeBillingPlanInformation(val *DowngradeBillingPlanInformation) *NullableDowngradeBillingPlanInformation {
	return &NullableDowngradeBillingPlanInformation{value: val, isSet: true}
}

func (v NullableDowngradeBillingPlanInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDowngradeBillingPlanInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


