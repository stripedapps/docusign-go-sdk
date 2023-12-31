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

// checks if the NewAccountDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewAccountDefinition{}

// NewAccountDefinition 
type NewAccountDefinition struct {
	// The account name for the new account.
	AccountName *string `json:"accountName,omitempty"`
	AccountSettings *AccountSettingsInformation `json:"accountSettings,omitempty"`
	AddressInformation *AccountAddress `json:"addressInformation,omitempty"`
	CreditCardInformation *CreditCardInformation `json:"creditCardInformation,omitempty"`
	DirectDebitProcessorInformation *DirectDebitProcessorInformation `json:"directDebitProcessorInformation,omitempty"`
	// The Distributor Code that you received from DocuSign.
	DistributorCode *string `json:"distributorCode,omitempty"`
	// The password for the `distributorCode`.
	DistributorPassword *string `json:"distributorPassword,omitempty"`
	// 
	EnablePreAuth *string `json:"enablePreAuth,omitempty"`
	// Reserved for DocuSign.
	EnvelopePartitionId *string `json:"envelopePartitionId,omitempty"`
	InitialUser *UserInformation `json:"initialUser,omitempty"`
	// The payment method used for the billing plan. Valid values are:  - `NotSupported` - `CreditCard` - `PurchaseOrder` - `Premium` - `Freemium` - `FreeTrial` - `AppStore` - `DigitalExternal` - `DirectDebit`
	PaymentMethod *string `json:"paymentMethod,omitempty"`
	// 
	PaymentProcessor *string `json:"paymentProcessor,omitempty"`
	PaymentProcessorInformation *PaymentProcessorInformation `json:"paymentProcessorInformation,omitempty"`
	PlanInformation *PlanInformation `json:"planInformation,omitempty"`
	// 
	ProcessPayment *string `json:"processPayment,omitempty"`
	ReferralInformation *ReferralInformation `json:"referralInformation,omitempty"`
	SocialAccountInformation *SocialAccountInformation `json:"socialAccountInformation,omitempty"`
	// 
	TaxExemptId *string `json:"taxExemptId,omitempty"`
}

// NewNewAccountDefinition instantiates a new NewAccountDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewAccountDefinition() *NewAccountDefinition {
	this := NewAccountDefinition{}
	return &this
}

// NewNewAccountDefinitionWithDefaults instantiates a new NewAccountDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewAccountDefinitionWithDefaults() *NewAccountDefinition {
	this := NewAccountDefinition{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *NewAccountDefinition) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountSettings returns the AccountSettings field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetAccountSettings() AccountSettingsInformation {
	if o == nil || IsNil(o.AccountSettings) {
		var ret AccountSettingsInformation
		return ret
	}
	return *o.AccountSettings
}

// GetAccountSettingsOk returns a tuple with the AccountSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetAccountSettingsOk() (*AccountSettingsInformation, bool) {
	if o == nil || IsNil(o.AccountSettings) {
		return nil, false
	}
	return o.AccountSettings, true
}

// HasAccountSettings returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasAccountSettings() bool {
	if o != nil && !IsNil(o.AccountSettings) {
		return true
	}

	return false
}

// SetAccountSettings gets a reference to the given AccountSettingsInformation and assigns it to the AccountSettings field.
func (o *NewAccountDefinition) SetAccountSettings(v AccountSettingsInformation) {
	o.AccountSettings = &v
}

// GetAddressInformation returns the AddressInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetAddressInformation() AccountAddress {
	if o == nil || IsNil(o.AddressInformation) {
		var ret AccountAddress
		return ret
	}
	return *o.AddressInformation
}

// GetAddressInformationOk returns a tuple with the AddressInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetAddressInformationOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.AddressInformation) {
		return nil, false
	}
	return o.AddressInformation, true
}

// HasAddressInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasAddressInformation() bool {
	if o != nil && !IsNil(o.AddressInformation) {
		return true
	}

	return false
}

// SetAddressInformation gets a reference to the given AccountAddress and assigns it to the AddressInformation field.
func (o *NewAccountDefinition) SetAddressInformation(v AccountAddress) {
	o.AddressInformation = &v
}

// GetCreditCardInformation returns the CreditCardInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetCreditCardInformation() CreditCardInformation {
	if o == nil || IsNil(o.CreditCardInformation) {
		var ret CreditCardInformation
		return ret
	}
	return *o.CreditCardInformation
}

// GetCreditCardInformationOk returns a tuple with the CreditCardInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetCreditCardInformationOk() (*CreditCardInformation, bool) {
	if o == nil || IsNil(o.CreditCardInformation) {
		return nil, false
	}
	return o.CreditCardInformation, true
}

// HasCreditCardInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasCreditCardInformation() bool {
	if o != nil && !IsNil(o.CreditCardInformation) {
		return true
	}

	return false
}

// SetCreditCardInformation gets a reference to the given CreditCardInformation and assigns it to the CreditCardInformation field.
func (o *NewAccountDefinition) SetCreditCardInformation(v CreditCardInformation) {
	o.CreditCardInformation = &v
}

// GetDirectDebitProcessorInformation returns the DirectDebitProcessorInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetDirectDebitProcessorInformation() DirectDebitProcessorInformation {
	if o == nil || IsNil(o.DirectDebitProcessorInformation) {
		var ret DirectDebitProcessorInformation
		return ret
	}
	return *o.DirectDebitProcessorInformation
}

// GetDirectDebitProcessorInformationOk returns a tuple with the DirectDebitProcessorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetDirectDebitProcessorInformationOk() (*DirectDebitProcessorInformation, bool) {
	if o == nil || IsNil(o.DirectDebitProcessorInformation) {
		return nil, false
	}
	return o.DirectDebitProcessorInformation, true
}

// HasDirectDebitProcessorInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasDirectDebitProcessorInformation() bool {
	if o != nil && !IsNil(o.DirectDebitProcessorInformation) {
		return true
	}

	return false
}

// SetDirectDebitProcessorInformation gets a reference to the given DirectDebitProcessorInformation and assigns it to the DirectDebitProcessorInformation field.
func (o *NewAccountDefinition) SetDirectDebitProcessorInformation(v DirectDebitProcessorInformation) {
	o.DirectDebitProcessorInformation = &v
}

// GetDistributorCode returns the DistributorCode field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetDistributorCode() string {
	if o == nil || IsNil(o.DistributorCode) {
		var ret string
		return ret
	}
	return *o.DistributorCode
}

// GetDistributorCodeOk returns a tuple with the DistributorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetDistributorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistributorCode) {
		return nil, false
	}
	return o.DistributorCode, true
}

// HasDistributorCode returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasDistributorCode() bool {
	if o != nil && !IsNil(o.DistributorCode) {
		return true
	}

	return false
}

// SetDistributorCode gets a reference to the given string and assigns it to the DistributorCode field.
func (o *NewAccountDefinition) SetDistributorCode(v string) {
	o.DistributorCode = &v
}

// GetDistributorPassword returns the DistributorPassword field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetDistributorPassword() string {
	if o == nil || IsNil(o.DistributorPassword) {
		var ret string
		return ret
	}
	return *o.DistributorPassword
}

// GetDistributorPasswordOk returns a tuple with the DistributorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetDistributorPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DistributorPassword) {
		return nil, false
	}
	return o.DistributorPassword, true
}

// HasDistributorPassword returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasDistributorPassword() bool {
	if o != nil && !IsNil(o.DistributorPassword) {
		return true
	}

	return false
}

// SetDistributorPassword gets a reference to the given string and assigns it to the DistributorPassword field.
func (o *NewAccountDefinition) SetDistributorPassword(v string) {
	o.DistributorPassword = &v
}

// GetEnablePreAuth returns the EnablePreAuth field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetEnablePreAuth() string {
	if o == nil || IsNil(o.EnablePreAuth) {
		var ret string
		return ret
	}
	return *o.EnablePreAuth
}

// GetEnablePreAuthOk returns a tuple with the EnablePreAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetEnablePreAuthOk() (*string, bool) {
	if o == nil || IsNil(o.EnablePreAuth) {
		return nil, false
	}
	return o.EnablePreAuth, true
}

// HasEnablePreAuth returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasEnablePreAuth() bool {
	if o != nil && !IsNil(o.EnablePreAuth) {
		return true
	}

	return false
}

// SetEnablePreAuth gets a reference to the given string and assigns it to the EnablePreAuth field.
func (o *NewAccountDefinition) SetEnablePreAuth(v string) {
	o.EnablePreAuth = &v
}

// GetEnvelopePartitionId returns the EnvelopePartitionId field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetEnvelopePartitionId() string {
	if o == nil || IsNil(o.EnvelopePartitionId) {
		var ret string
		return ret
	}
	return *o.EnvelopePartitionId
}

// GetEnvelopePartitionIdOk returns a tuple with the EnvelopePartitionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetEnvelopePartitionIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopePartitionId) {
		return nil, false
	}
	return o.EnvelopePartitionId, true
}

// HasEnvelopePartitionId returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasEnvelopePartitionId() bool {
	if o != nil && !IsNil(o.EnvelopePartitionId) {
		return true
	}

	return false
}

// SetEnvelopePartitionId gets a reference to the given string and assigns it to the EnvelopePartitionId field.
func (o *NewAccountDefinition) SetEnvelopePartitionId(v string) {
	o.EnvelopePartitionId = &v
}

// GetInitialUser returns the InitialUser field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetInitialUser() UserInformation {
	if o == nil || IsNil(o.InitialUser) {
		var ret UserInformation
		return ret
	}
	return *o.InitialUser
}

// GetInitialUserOk returns a tuple with the InitialUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetInitialUserOk() (*UserInformation, bool) {
	if o == nil || IsNil(o.InitialUser) {
		return nil, false
	}
	return o.InitialUser, true
}

// HasInitialUser returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasInitialUser() bool {
	if o != nil && !IsNil(o.InitialUser) {
		return true
	}

	return false
}

// SetInitialUser gets a reference to the given UserInformation and assigns it to the InitialUser field.
func (o *NewAccountDefinition) SetInitialUser(v UserInformation) {
	o.InitialUser = &v
}

// GetPaymentMethod returns the PaymentMethod field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetPaymentMethod() string {
	if o == nil || IsNil(o.PaymentMethod) {
		var ret string
		return ret
	}
	return *o.PaymentMethod
}

// GetPaymentMethodOk returns a tuple with the PaymentMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetPaymentMethodOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentMethod) {
		return nil, false
	}
	return o.PaymentMethod, true
}

// HasPaymentMethod returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasPaymentMethod() bool {
	if o != nil && !IsNil(o.PaymentMethod) {
		return true
	}

	return false
}

// SetPaymentMethod gets a reference to the given string and assigns it to the PaymentMethod field.
func (o *NewAccountDefinition) SetPaymentMethod(v string) {
	o.PaymentMethod = &v
}

// GetPaymentProcessor returns the PaymentProcessor field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetPaymentProcessor() string {
	if o == nil || IsNil(o.PaymentProcessor) {
		var ret string
		return ret
	}
	return *o.PaymentProcessor
}

// GetPaymentProcessorOk returns a tuple with the PaymentProcessor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetPaymentProcessorOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentProcessor) {
		return nil, false
	}
	return o.PaymentProcessor, true
}

// HasPaymentProcessor returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasPaymentProcessor() bool {
	if o != nil && !IsNil(o.PaymentProcessor) {
		return true
	}

	return false
}

// SetPaymentProcessor gets a reference to the given string and assigns it to the PaymentProcessor field.
func (o *NewAccountDefinition) SetPaymentProcessor(v string) {
	o.PaymentProcessor = &v
}

// GetPaymentProcessorInformation returns the PaymentProcessorInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetPaymentProcessorInformation() PaymentProcessorInformation {
	if o == nil || IsNil(o.PaymentProcessorInformation) {
		var ret PaymentProcessorInformation
		return ret
	}
	return *o.PaymentProcessorInformation
}

// GetPaymentProcessorInformationOk returns a tuple with the PaymentProcessorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetPaymentProcessorInformationOk() (*PaymentProcessorInformation, bool) {
	if o == nil || IsNil(o.PaymentProcessorInformation) {
		return nil, false
	}
	return o.PaymentProcessorInformation, true
}

// HasPaymentProcessorInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasPaymentProcessorInformation() bool {
	if o != nil && !IsNil(o.PaymentProcessorInformation) {
		return true
	}

	return false
}

// SetPaymentProcessorInformation gets a reference to the given PaymentProcessorInformation and assigns it to the PaymentProcessorInformation field.
func (o *NewAccountDefinition) SetPaymentProcessorInformation(v PaymentProcessorInformation) {
	o.PaymentProcessorInformation = &v
}

// GetPlanInformation returns the PlanInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetPlanInformation() PlanInformation {
	if o == nil || IsNil(o.PlanInformation) {
		var ret PlanInformation
		return ret
	}
	return *o.PlanInformation
}

// GetPlanInformationOk returns a tuple with the PlanInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetPlanInformationOk() (*PlanInformation, bool) {
	if o == nil || IsNil(o.PlanInformation) {
		return nil, false
	}
	return o.PlanInformation, true
}

// HasPlanInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasPlanInformation() bool {
	if o != nil && !IsNil(o.PlanInformation) {
		return true
	}

	return false
}

// SetPlanInformation gets a reference to the given PlanInformation and assigns it to the PlanInformation field.
func (o *NewAccountDefinition) SetPlanInformation(v PlanInformation) {
	o.PlanInformation = &v
}

// GetProcessPayment returns the ProcessPayment field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetProcessPayment() string {
	if o == nil || IsNil(o.ProcessPayment) {
		var ret string
		return ret
	}
	return *o.ProcessPayment
}

// GetProcessPaymentOk returns a tuple with the ProcessPayment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetProcessPaymentOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessPayment) {
		return nil, false
	}
	return o.ProcessPayment, true
}

// HasProcessPayment returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasProcessPayment() bool {
	if o != nil && !IsNil(o.ProcessPayment) {
		return true
	}

	return false
}

// SetProcessPayment gets a reference to the given string and assigns it to the ProcessPayment field.
func (o *NewAccountDefinition) SetProcessPayment(v string) {
	o.ProcessPayment = &v
}

// GetReferralInformation returns the ReferralInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetReferralInformation() ReferralInformation {
	if o == nil || IsNil(o.ReferralInformation) {
		var ret ReferralInformation
		return ret
	}
	return *o.ReferralInformation
}

// GetReferralInformationOk returns a tuple with the ReferralInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetReferralInformationOk() (*ReferralInformation, bool) {
	if o == nil || IsNil(o.ReferralInformation) {
		return nil, false
	}
	return o.ReferralInformation, true
}

// HasReferralInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasReferralInformation() bool {
	if o != nil && !IsNil(o.ReferralInformation) {
		return true
	}

	return false
}

// SetReferralInformation gets a reference to the given ReferralInformation and assigns it to the ReferralInformation field.
func (o *NewAccountDefinition) SetReferralInformation(v ReferralInformation) {
	o.ReferralInformation = &v
}

// GetSocialAccountInformation returns the SocialAccountInformation field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetSocialAccountInformation() SocialAccountInformation {
	if o == nil || IsNil(o.SocialAccountInformation) {
		var ret SocialAccountInformation
		return ret
	}
	return *o.SocialAccountInformation
}

// GetSocialAccountInformationOk returns a tuple with the SocialAccountInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetSocialAccountInformationOk() (*SocialAccountInformation, bool) {
	if o == nil || IsNil(o.SocialAccountInformation) {
		return nil, false
	}
	return o.SocialAccountInformation, true
}

// HasSocialAccountInformation returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasSocialAccountInformation() bool {
	if o != nil && !IsNil(o.SocialAccountInformation) {
		return true
	}

	return false
}

// SetSocialAccountInformation gets a reference to the given SocialAccountInformation and assigns it to the SocialAccountInformation field.
func (o *NewAccountDefinition) SetSocialAccountInformation(v SocialAccountInformation) {
	o.SocialAccountInformation = &v
}

// GetTaxExemptId returns the TaxExemptId field value if set, zero value otherwise.
func (o *NewAccountDefinition) GetTaxExemptId() string {
	if o == nil || IsNil(o.TaxExemptId) {
		var ret string
		return ret
	}
	return *o.TaxExemptId
}

// GetTaxExemptIdOk returns a tuple with the TaxExemptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewAccountDefinition) GetTaxExemptIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxExemptId) {
		return nil, false
	}
	return o.TaxExemptId, true
}

// HasTaxExemptId returns a boolean if a field has been set.
func (o *NewAccountDefinition) HasTaxExemptId() bool {
	if o != nil && !IsNil(o.TaxExemptId) {
		return true
	}

	return false
}

// SetTaxExemptId gets a reference to the given string and assigns it to the TaxExemptId field.
func (o *NewAccountDefinition) SetTaxExemptId(v string) {
	o.TaxExemptId = &v
}

func (o NewAccountDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewAccountDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.AccountSettings) {
		toSerialize["accountSettings"] = o.AccountSettings
	}
	if !IsNil(o.AddressInformation) {
		toSerialize["addressInformation"] = o.AddressInformation
	}
	if !IsNil(o.CreditCardInformation) {
		toSerialize["creditCardInformation"] = o.CreditCardInformation
	}
	if !IsNil(o.DirectDebitProcessorInformation) {
		toSerialize["directDebitProcessorInformation"] = o.DirectDebitProcessorInformation
	}
	if !IsNil(o.DistributorCode) {
		toSerialize["distributorCode"] = o.DistributorCode
	}
	if !IsNil(o.DistributorPassword) {
		toSerialize["distributorPassword"] = o.DistributorPassword
	}
	if !IsNil(o.EnablePreAuth) {
		toSerialize["enablePreAuth"] = o.EnablePreAuth
	}
	if !IsNil(o.EnvelopePartitionId) {
		toSerialize["envelopePartitionId"] = o.EnvelopePartitionId
	}
	if !IsNil(o.InitialUser) {
		toSerialize["initialUser"] = o.InitialUser
	}
	if !IsNil(o.PaymentMethod) {
		toSerialize["paymentMethod"] = o.PaymentMethod
	}
	if !IsNil(o.PaymentProcessor) {
		toSerialize["paymentProcessor"] = o.PaymentProcessor
	}
	if !IsNil(o.PaymentProcessorInformation) {
		toSerialize["paymentProcessorInformation"] = o.PaymentProcessorInformation
	}
	if !IsNil(o.PlanInformation) {
		toSerialize["planInformation"] = o.PlanInformation
	}
	if !IsNil(o.ProcessPayment) {
		toSerialize["processPayment"] = o.ProcessPayment
	}
	if !IsNil(o.ReferralInformation) {
		toSerialize["referralInformation"] = o.ReferralInformation
	}
	if !IsNil(o.SocialAccountInformation) {
		toSerialize["socialAccountInformation"] = o.SocialAccountInformation
	}
	if !IsNil(o.TaxExemptId) {
		toSerialize["taxExemptId"] = o.TaxExemptId
	}
	return toSerialize, nil
}

type NullableNewAccountDefinition struct {
	value *NewAccountDefinition
	isSet bool
}

func (v NullableNewAccountDefinition) Get() *NewAccountDefinition {
	return v.value
}

func (v *NullableNewAccountDefinition) Set(val *NewAccountDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableNewAccountDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableNewAccountDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewAccountDefinition(val *NewAccountDefinition) *NullableNewAccountDefinition {
	return &NullableNewAccountDefinition{value: val, isSet: true}
}

func (v NullableNewAccountDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewAccountDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


