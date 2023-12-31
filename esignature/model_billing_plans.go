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

// checks if the BillingPlans type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingPlans{}

// BillingPlans Billing plans
type BillingPlans struct {
	BillingAddress *AccountAddress `json:"billingAddress,omitempty"`
	// When **true,** the credit card address information is the same as that returned as the billing address. If false, then the billing address is considered a billing contact address, and the credit card address can be different.
	BillingAddressIsCreditCardAddress *string `json:"billingAddressIsCreditCardAddress,omitempty"`
	BillingPlan *AccountBillingPlan `json:"billingPlan,omitempty"`
	CreditCardInformation *CreditCardInformation `json:"creditCardInformation,omitempty"`
	DirectDebitProcessorInformation *DirectDebitProcessorInformation `json:"directDebitProcessorInformation,omitempty"`
	DowngradePlanInformation *DowngradePlanUpdateResponse `json:"downgradePlanInformation,omitempty"`
	DowngradeRequestInformation *DowngradeRequestInformation `json:"downgradeRequestInformation,omitempty"`
	EntityInformation *BillingEntityInformationResponse `json:"entityInformation,omitempty"`
	// The payment method used for the billing plan. Valid values are:  - `NotSupported` - `CreditCard` - `PurchaseOrder` - `Premium` - `Freemium` - `FreeTrial` - `AppStore` - `DigitalExternal` - `DirectDebit`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	PaymentProcessorInformation *PaymentProcessorInformation `json:"paymentProcessorInformation,omitempty"`
	ReferralInformation *ReferralInformation `json:"referralInformation,omitempty"`
	// A list of billing plans that the current billing plan can be rolled into.
	SuccessorPlans []BillingPlan `json:"successorPlans,omitempty"`
	// 
	TaxExemptId *string `json:"taxExemptId,omitempty"`
}

// NewBillingPlans instantiates a new BillingPlans object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingPlans() *BillingPlans {
	this := BillingPlans{}
	return &this
}

// NewBillingPlansWithDefaults instantiates a new BillingPlans object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingPlansWithDefaults() *BillingPlans {
	this := BillingPlans{}
	return &this
}

// GetBillingAddress returns the BillingAddress field value if set, zero value otherwise.
func (o *BillingPlans) GetBillingAddress() AccountAddress {
	if o == nil || IsNil(o.BillingAddress) {
		var ret AccountAddress
		return ret
	}
	return *o.BillingAddress
}

// GetBillingAddressOk returns a tuple with the BillingAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetBillingAddressOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.BillingAddress) {
		return nil, false
	}
	return o.BillingAddress, true
}

// HasBillingAddress returns a boolean if a field has been set.
func (o *BillingPlans) HasBillingAddress() bool {
	if o != nil && !IsNil(o.BillingAddress) {
		return true
	}

	return false
}

// SetBillingAddress gets a reference to the given AccountAddress and assigns it to the BillingAddress field.
func (o *BillingPlans) SetBillingAddress(v AccountAddress) {
	o.BillingAddress = &v
}

// GetBillingAddressIsCreditCardAddress returns the BillingAddressIsCreditCardAddress field value if set, zero value otherwise.
func (o *BillingPlans) GetBillingAddressIsCreditCardAddress() string {
	if o == nil || IsNil(o.BillingAddressIsCreditCardAddress) {
		var ret string
		return ret
	}
	return *o.BillingAddressIsCreditCardAddress
}

// GetBillingAddressIsCreditCardAddressOk returns a tuple with the BillingAddressIsCreditCardAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetBillingAddressIsCreditCardAddressOk() (*string, bool) {
	if o == nil || IsNil(o.BillingAddressIsCreditCardAddress) {
		return nil, false
	}
	return o.BillingAddressIsCreditCardAddress, true
}

// HasBillingAddressIsCreditCardAddress returns a boolean if a field has been set.
func (o *BillingPlans) HasBillingAddressIsCreditCardAddress() bool {
	if o != nil && !IsNil(o.BillingAddressIsCreditCardAddress) {
		return true
	}

	return false
}

// SetBillingAddressIsCreditCardAddress gets a reference to the given string and assigns it to the BillingAddressIsCreditCardAddress field.
func (o *BillingPlans) SetBillingAddressIsCreditCardAddress(v string) {
	o.BillingAddressIsCreditCardAddress = &v
}

// GetBillingPlan returns the BillingPlan field value if set, zero value otherwise.
func (o *BillingPlans) GetBillingPlan() AccountBillingPlan {
	if o == nil || IsNil(o.BillingPlan) {
		var ret AccountBillingPlan
		return ret
	}
	return *o.BillingPlan
}

// GetBillingPlanOk returns a tuple with the BillingPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetBillingPlanOk() (*AccountBillingPlan, bool) {
	if o == nil || IsNil(o.BillingPlan) {
		return nil, false
	}
	return o.BillingPlan, true
}

// HasBillingPlan returns a boolean if a field has been set.
func (o *BillingPlans) HasBillingPlan() bool {
	if o != nil && !IsNil(o.BillingPlan) {
		return true
	}

	return false
}

// SetBillingPlan gets a reference to the given AccountBillingPlan and assigns it to the BillingPlan field.
func (o *BillingPlans) SetBillingPlan(v AccountBillingPlan) {
	o.BillingPlan = &v
}

// GetCreditCardInformation returns the CreditCardInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetCreditCardInformation() CreditCardInformation {
	if o == nil || IsNil(o.CreditCardInformation) {
		var ret CreditCardInformation
		return ret
	}
	return *o.CreditCardInformation
}

// GetCreditCardInformationOk returns a tuple with the CreditCardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetCreditCardInformationOk() (*CreditCardInformation, bool) {
	if o == nil || IsNil(o.CreditCardInformation) {
		return nil, false
	}
	return o.CreditCardInformation, true
}

// HasCreditCardInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasCreditCardInformation() bool {
	if o != nil && !IsNil(o.CreditCardInformation) {
		return true
	}

	return false
}

// SetCreditCardInformation gets a reference to the given CreditCardInformation and assigns it to the CreditCardInformation field.
func (o *BillingPlans) SetCreditCardInformation(v CreditCardInformation) {
	o.CreditCardInformation = &v
}

// GetDirectDebitProcessorInformation returns the DirectDebitProcessorInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetDirectDebitProcessorInformation() DirectDebitProcessorInformation {
	if o == nil || IsNil(o.DirectDebitProcessorInformation) {
		var ret DirectDebitProcessorInformation
		return ret
	}
	return *o.DirectDebitProcessorInformation
}

// GetDirectDebitProcessorInformationOk returns a tuple with the DirectDebitProcessorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetDirectDebitProcessorInformationOk() (*DirectDebitProcessorInformation, bool) {
	if o == nil || IsNil(o.DirectDebitProcessorInformation) {
		return nil, false
	}
	return o.DirectDebitProcessorInformation, true
}

// HasDirectDebitProcessorInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasDirectDebitProcessorInformation() bool {
	if o != nil && !IsNil(o.DirectDebitProcessorInformation) {
		return true
	}

	return false
}

// SetDirectDebitProcessorInformation gets a reference to the given DirectDebitProcessorInformation and assigns it to the DirectDebitProcessorInformation field.
func (o *BillingPlans) SetDirectDebitProcessorInformation(v DirectDebitProcessorInformation) {
	o.DirectDebitProcessorInformation = &v
}

// GetDowngradePlanInformation returns the DowngradePlanInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetDowngradePlanInformation() DowngradePlanUpdateResponse {
	if o == nil || IsNil(o.DowngradePlanInformation) {
		var ret DowngradePlanUpdateResponse
		return ret
	}
	return *o.DowngradePlanInformation
}

// GetDowngradePlanInformationOk returns a tuple with the DowngradePlanInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetDowngradePlanInformationOk() (*DowngradePlanUpdateResponse, bool) {
	if o == nil || IsNil(o.DowngradePlanInformation) {
		return nil, false
	}
	return o.DowngradePlanInformation, true
}

// HasDowngradePlanInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasDowngradePlanInformation() bool {
	if o != nil && !IsNil(o.DowngradePlanInformation) {
		return true
	}

	return false
}

// SetDowngradePlanInformation gets a reference to the given DowngradePlanUpdateResponse and assigns it to the DowngradePlanInformation field.
func (o *BillingPlans) SetDowngradePlanInformation(v DowngradePlanUpdateResponse) {
	o.DowngradePlanInformation = &v
}

// GetDowngradeRequestInformation returns the DowngradeRequestInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetDowngradeRequestInformation() DowngradeRequestInformation {
	if o == nil || IsNil(o.DowngradeRequestInformation) {
		var ret DowngradeRequestInformation
		return ret
	}
	return *o.DowngradeRequestInformation
}

// GetDowngradeRequestInformationOk returns a tuple with the DowngradeRequestInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetDowngradeRequestInformationOk() (*DowngradeRequestInformation, bool) {
	if o == nil || IsNil(o.DowngradeRequestInformation) {
		return nil, false
	}
	return o.DowngradeRequestInformation, true
}

// HasDowngradeRequestInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasDowngradeRequestInformation() bool {
	if o != nil && !IsNil(o.DowngradeRequestInformation) {
		return true
	}

	return false
}

// SetDowngradeRequestInformation gets a reference to the given DowngradeRequestInformation and assigns it to the DowngradeRequestInformation field.
func (o *BillingPlans) SetDowngradeRequestInformation(v DowngradeRequestInformation) {
	o.DowngradeRequestInformation = &v
}

// GetEntityInformation returns the EntityInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetEntityInformation() BillingEntityInformationResponse {
	if o == nil || IsNil(o.EntityInformation) {
		var ret BillingEntityInformationResponse
		return ret
	}
	return *o.EntityInformation
}

// GetEntityInformationOk returns a tuple with the EntityInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetEntityInformationOk() (*BillingEntityInformationResponse, bool) {
	if o == nil || IsNil(o.EntityInformation) {
		return nil, false
	}
	return o.EntityInformation, true
}

// HasEntityInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasEntityInformation() bool {
	if o != nil && !IsNil(o.EntityInformation) {
		return true
	}

	return false
}

// SetEntityInformation gets a reference to the given BillingEntityInformationResponse and assigns it to the EntityInformation field.
func (o *BillingPlans) SetEntityInformation(v BillingEntityInformationResponse) {
	o.EntityInformation = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *BillingPlans) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *BillingPlans) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *BillingPlans) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentProcessorInformation returns the PaymentProcessorInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetPaymentProcessorInformation() PaymentProcessorInformation {
	if o == nil || IsNil(o.PaymentProcessorInformation) {
		var ret PaymentProcessorInformation
		return ret
	}
	return *o.PaymentProcessorInformation
}

// GetPaymentProcessorInformationOk returns a tuple with the PaymentProcessorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetPaymentProcessorInformationOk() (*PaymentProcessorInformation, bool) {
	if o == nil || IsNil(o.PaymentProcessorInformation) {
		return nil, false
	}
	return o.PaymentProcessorInformation, true
}

// HasPaymentProcessorInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasPaymentProcessorInformation() bool {
	if o != nil && !IsNil(o.PaymentProcessorInformation) {
		return true
	}

	return false
}

// SetPaymentProcessorInformation gets a reference to the given PaymentProcessorInformation and assigns it to the PaymentProcessorInformation field.
func (o *BillingPlans) SetPaymentProcessorInformation(v PaymentProcessorInformation) {
	o.PaymentProcessorInformation = &v
}

// GetReferralInformation returns the ReferralInformation field value if set, zero value otherwise.
func (o *BillingPlans) GetReferralInformation() ReferralInformation {
	if o == nil || IsNil(o.ReferralInformation) {
		var ret ReferralInformation
		return ret
	}
	return *o.ReferralInformation
}

// GetReferralInformationOk returns a tuple with the ReferralInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetReferralInformationOk() (*ReferralInformation, bool) {
	if o == nil || IsNil(o.ReferralInformation) {
		return nil, false
	}
	return o.ReferralInformation, true
}

// HasReferralInformation returns a boolean if a field has been set.
func (o *BillingPlans) HasReferralInformation() bool {
	if o != nil && !IsNil(o.ReferralInformation) {
		return true
	}

	return false
}

// SetReferralInformation gets a reference to the given ReferralInformation and assigns it to the ReferralInformation field.
func (o *BillingPlans) SetReferralInformation(v ReferralInformation) {
	o.ReferralInformation = &v
}

// GetSuccessorPlans returns the SuccessorPlans field value if set, zero value otherwise.
func (o *BillingPlans) GetSuccessorPlans() []BillingPlan {
	if o == nil || IsNil(o.SuccessorPlans) {
		var ret []BillingPlan
		return ret
	}
	return o.SuccessorPlans
}

// GetSuccessorPlansOk returns a tuple with the SuccessorPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetSuccessorPlansOk() ([]BillingPlan, bool) {
	if o == nil || IsNil(o.SuccessorPlans) {
		return nil, false
	}
	return o.SuccessorPlans, true
}

// HasSuccessorPlans returns a boolean if a field has been set.
func (o *BillingPlans) HasSuccessorPlans() bool {
	if o != nil && !IsNil(o.SuccessorPlans) {
		return true
	}

	return false
}

// SetSuccessorPlans gets a reference to the given []BillingPlan and assigns it to the SuccessorPlans field.
func (o *BillingPlans) SetSuccessorPlans(v []BillingPlan) {
	o.SuccessorPlans = v
}

// GetTaxExemptId returns the TaxExemptId field value if set, zero value otherwise.
func (o *BillingPlans) GetTaxExemptId() string {
	if o == nil || IsNil(o.TaxExemptId) {
		var ret string
		return ret
	}
	return *o.TaxExemptId
}

// GetTaxExemptIdOk returns a tuple with the TaxExemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingPlans) GetTaxExemptIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxExemptId) {
		return nil, false
	}
	return o.TaxExemptId, true
}

// HasTaxExemptId returns a boolean if a field has been set.
func (o *BillingPlans) HasTaxExemptId() bool {
	if o != nil && !IsNil(o.TaxExemptId) {
		return true
	}

	return false
}

// SetTaxExemptId gets a reference to the given string and assigns it to the TaxExemptId field.
func (o *BillingPlans) SetTaxExemptId(v string) {
	o.TaxExemptId = &v
}

func (o BillingPlans) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingPlans) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingAddress) {
		toSerialize["billingAddress"] = o.BillingAddress
	}
	if !IsNil(o.BillingAddressIsCreditCardAddress) {
		toSerialize["billingAddressIsCreditCardAddress"] = o.BillingAddressIsCreditCardAddress
	}
	if !IsNil(o.BillingPlan) {
		toSerialize["billingPlan"] = o.BillingPlan
	}
	if !IsNil(o.CreditCardInformation) {
		toSerialize["creditCardInformation"] = o.CreditCardInformation
	}
	if !IsNil(o.DirectDebitProcessorInformation) {
		toSerialize["directDebitProcessorInformation"] = o.DirectDebitProcessorInformation
	}
	if !IsNil(o.DowngradePlanInformation) {
		toSerialize["downgradePlanInformation"] = o.DowngradePlanInformation
	}
	if !IsNil(o.DowngradeRequestInformation) {
		toSerialize["downgradeRequestInformation"] = o.DowngradeRequestInformation
	}
	if !IsNil(o.EntityInformation) {
		toSerialize["entityInformation"] = o.EntityInformation
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentProcessorInformation) {
		toSerialize["paymentProcessorInformation"] = o.PaymentProcessorInformation
	}
	if !IsNil(o.ReferralInformation) {
		toSerialize["referralInformation"] = o.ReferralInformation
	}
	if !IsNil(o.SuccessorPlans) {
		toSerialize["successorPlans"] = o.SuccessorPlans
	}
	if !IsNil(o.TaxExemptId) {
		toSerialize["taxExemptId"] = o.TaxExemptId
	}
	return toSerialize, nil
}

type NullableBillingPlans struct {
	value *BillingPlans
	isSet bool
}

func (v NullableBillingPlans) Get() *BillingPlans {
	return v.value
}

func (v *NullableBillingPlans) Set(val *BillingPlans) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingPlans) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingPlans) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingPlans(val *BillingPlans) *NullableBillingPlans {
	return &NullableBillingPlans{value: val, isSet: true}
}

func (v NullableBillingPlans) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingPlans) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


