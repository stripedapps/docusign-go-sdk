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

// checks if the ProvisioningInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningInformation{}

// ProvisioningInformation 
type ProvisioningInformation struct {
	// 
	DefaultConnectionId *string `json:"defaultConnectionId,omitempty"`
	// 
	DefaultPlanId *string `json:"defaultPlanId,omitempty"`
	// The code that identifies the billing plan groups and plans for the new account.
	DistributorCode *string `json:"distributorCode,omitempty"`
	// The password for the `distributorCode`.
	DistributorPassword *string `json:"distributorPassword,omitempty"`
	// 
	PasswordRuleText *string `json:"passwordRuleText,omitempty"`
	// 
	PlanPromotionText *string `json:"planPromotionText,omitempty"`
	// 
	PurchaseOrderOrPromAllowed *string `json:"purchaseOrderOrPromAllowed,omitempty"`
}

// NewProvisioningInformation instantiates a new ProvisioningInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningInformation() *ProvisioningInformation {
	this := ProvisioningInformation{}
	return &this
}

// NewProvisioningInformationWithDefaults instantiates a new ProvisioningInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningInformationWithDefaults() *ProvisioningInformation {
	this := ProvisioningInformation{}
	return &this
}

// GetDefaultConnectionId returns the DefaultConnectionId field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetDefaultConnectionId() string {
	if o == nil || IsNil(o.DefaultConnectionId) {
		var ret string
		return ret
	}
	return *o.DefaultConnectionId
}

// GetDefaultConnectionIdOk returns a tuple with the DefaultConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetDefaultConnectionIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultConnectionId) {
		return nil, false
	}
	return o.DefaultConnectionId, true
}

// HasDefaultConnectionId returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasDefaultConnectionId() bool {
	if o != nil && !IsNil(o.DefaultConnectionId) {
		return true
	}

	return false
}

// SetDefaultConnectionId gets a reference to the given string and assigns it to the DefaultConnectionId field.
func (o *ProvisioningInformation) SetDefaultConnectionId(v string) {
	o.DefaultConnectionId = &v
}

// GetDefaultPlanId returns the DefaultPlanId field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetDefaultPlanId() string {
	if o == nil || IsNil(o.DefaultPlanId) {
		var ret string
		return ret
	}
	return *o.DefaultPlanId
}

// GetDefaultPlanIdOk returns a tuple with the DefaultPlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetDefaultPlanIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultPlanId) {
		return nil, false
	}
	return o.DefaultPlanId, true
}

// HasDefaultPlanId returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasDefaultPlanId() bool {
	if o != nil && !IsNil(o.DefaultPlanId) {
		return true
	}

	return false
}

// SetDefaultPlanId gets a reference to the given string and assigns it to the DefaultPlanId field.
func (o *ProvisioningInformation) SetDefaultPlanId(v string) {
	o.DefaultPlanId = &v
}

// GetDistributorCode returns the DistributorCode field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetDistributorCode() string {
	if o == nil || IsNil(o.DistributorCode) {
		var ret string
		return ret
	}
	return *o.DistributorCode
}

// GetDistributorCodeOk returns a tuple with the DistributorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetDistributorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.DistributorCode) {
		return nil, false
	}
	return o.DistributorCode, true
}

// HasDistributorCode returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasDistributorCode() bool {
	if o != nil && !IsNil(o.DistributorCode) {
		return true
	}

	return false
}

// SetDistributorCode gets a reference to the given string and assigns it to the DistributorCode field.
func (o *ProvisioningInformation) SetDistributorCode(v string) {
	o.DistributorCode = &v
}

// GetDistributorPassword returns the DistributorPassword field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetDistributorPassword() string {
	if o == nil || IsNil(o.DistributorPassword) {
		var ret string
		return ret
	}
	return *o.DistributorPassword
}

// GetDistributorPasswordOk returns a tuple with the DistributorPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetDistributorPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DistributorPassword) {
		return nil, false
	}
	return o.DistributorPassword, true
}

// HasDistributorPassword returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasDistributorPassword() bool {
	if o != nil && !IsNil(o.DistributorPassword) {
		return true
	}

	return false
}

// SetDistributorPassword gets a reference to the given string and assigns it to the DistributorPassword field.
func (o *ProvisioningInformation) SetDistributorPassword(v string) {
	o.DistributorPassword = &v
}

// GetPasswordRuleText returns the PasswordRuleText field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetPasswordRuleText() string {
	if o == nil || IsNil(o.PasswordRuleText) {
		var ret string
		return ret
	}
	return *o.PasswordRuleText
}

// GetPasswordRuleTextOk returns a tuple with the PasswordRuleText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetPasswordRuleTextOk() (*string, bool) {
	if o == nil || IsNil(o.PasswordRuleText) {
		return nil, false
	}
	return o.PasswordRuleText, true
}

// HasPasswordRuleText returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasPasswordRuleText() bool {
	if o != nil && !IsNil(o.PasswordRuleText) {
		return true
	}

	return false
}

// SetPasswordRuleText gets a reference to the given string and assigns it to the PasswordRuleText field.
func (o *ProvisioningInformation) SetPasswordRuleText(v string) {
	o.PasswordRuleText = &v
}

// GetPlanPromotionText returns the PlanPromotionText field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetPlanPromotionText() string {
	if o == nil || IsNil(o.PlanPromotionText) {
		var ret string
		return ret
	}
	return *o.PlanPromotionText
}

// GetPlanPromotionTextOk returns a tuple with the PlanPromotionText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetPlanPromotionTextOk() (*string, bool) {
	if o == nil || IsNil(o.PlanPromotionText) {
		return nil, false
	}
	return o.PlanPromotionText, true
}

// HasPlanPromotionText returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasPlanPromotionText() bool {
	if o != nil && !IsNil(o.PlanPromotionText) {
		return true
	}

	return false
}

// SetPlanPromotionText gets a reference to the given string and assigns it to the PlanPromotionText field.
func (o *ProvisioningInformation) SetPlanPromotionText(v string) {
	o.PlanPromotionText = &v
}

// GetPurchaseOrderOrPromAllowed returns the PurchaseOrderOrPromAllowed field value if set, zero value otherwise.
func (o *ProvisioningInformation) GetPurchaseOrderOrPromAllowed() string {
	if o == nil || IsNil(o.PurchaseOrderOrPromAllowed) {
		var ret string
		return ret
	}
	return *o.PurchaseOrderOrPromAllowed
}

// GetPurchaseOrderOrPromAllowedOk returns a tuple with the PurchaseOrderOrPromAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningInformation) GetPurchaseOrderOrPromAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.PurchaseOrderOrPromAllowed) {
		return nil, false
	}
	return o.PurchaseOrderOrPromAllowed, true
}

// HasPurchaseOrderOrPromAllowed returns a boolean if a field has been set.
func (o *ProvisioningInformation) HasPurchaseOrderOrPromAllowed() bool {
	if o != nil && !IsNil(o.PurchaseOrderOrPromAllowed) {
		return true
	}

	return false
}

// SetPurchaseOrderOrPromAllowed gets a reference to the given string and assigns it to the PurchaseOrderOrPromAllowed field.
func (o *ProvisioningInformation) SetPurchaseOrderOrPromAllowed(v string) {
	o.PurchaseOrderOrPromAllowed = &v
}

func (o ProvisioningInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultConnectionId) {
		toSerialize["defaultConnectionId"] = o.DefaultConnectionId
	}
	if !IsNil(o.DefaultPlanId) {
		toSerialize["defaultPlanId"] = o.DefaultPlanId
	}
	if !IsNil(o.DistributorCode) {
		toSerialize["distributorCode"] = o.DistributorCode
	}
	if !IsNil(o.DistributorPassword) {
		toSerialize["distributorPassword"] = o.DistributorPassword
	}
	if !IsNil(o.PasswordRuleText) {
		toSerialize["passwordRuleText"] = o.PasswordRuleText
	}
	if !IsNil(o.PlanPromotionText) {
		toSerialize["planPromotionText"] = o.PlanPromotionText
	}
	if !IsNil(o.PurchaseOrderOrPromAllowed) {
		toSerialize["purchaseOrderOrPromAllowed"] = o.PurchaseOrderOrPromAllowed
	}
	return toSerialize, nil
}

type NullableProvisioningInformation struct {
	value *ProvisioningInformation
	isSet bool
}

func (v NullableProvisioningInformation) Get() *ProvisioningInformation {
	return v.value
}

func (v *NullableProvisioningInformation) Set(val *ProvisioningInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningInformation(val *ProvisioningInformation) *NullableProvisioningInformation {
	return &NullableProvisioningInformation{value: val, isSet: true}
}

func (v NullableProvisioningInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

