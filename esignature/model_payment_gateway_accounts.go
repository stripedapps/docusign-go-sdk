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

// checks if the PaymentGatewayAccounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentGatewayAccounts{}

// PaymentGatewayAccounts Information about a connected payment gateway account.
type PaymentGatewayAccounts struct {
	// When **true,** the sender can pass custom metadata about the payment to the payment gateway. You pass in this metadata on an EnvelopeRecipientTab, in the `customMetadata` property under `paymentDetails`.   For example, this property is set to **true** for the Authorize.net gateway by default. As a result, the extra metadata that you send displays for the Authorize.net transaction in the merchant gateway portal under **Description.**  **Note:** This property is read-only and cannot be changed.
	AllowCustomMetadata *bool `json:"allowCustomMetadata,omitempty"`
	Config *PaymentGatewayAccountSetting `json:"config,omitempty"`
	// A user-defined name for a connected gateway account.  This name is used in the Admin panel in the list of connected accounts and in Tagger in the payment gateway selector.  The human-readable version of `paymentGatewayAccountId`.
	DisplayName *string `json:"displayName,omitempty"`
	// When **true,** the payment gateway account is enabled.
	IsEnabled *string `json:"isEnabled,omitempty"`
	// Reserved for DocuSign.
	IsLegacy *string `json:"isLegacy,omitempty"`
	// The UTC DateTime that the payment gateway account was last updated.
	LastModified *string `json:"lastModified,omitempty"`
	// Payment gateway used by the connected gateway account. This is the name used by the API. For a human-readable version use `paymentGatewayDisplayName`.  Possible values are:  * `Stripe` * `Braintree` * `AuthorizeDotNet` * `CyberSource` * `Zuora` * `Elavon`
	PaymentGateway *string `json:"paymentGateway,omitempty"`
	// A GUID that identifies the payment gateway account. For a human-readable version use `displayName`.
	PaymentGatewayAccountId *string `json:"paymentGatewayAccountId,omitempty"`
	// The display name of the payment gateway that the connected gateway account uses. This is the human-readable version of `paymentGateway`.  Possible values are:  * Stripe * Braintree * Authorize.Net * CyberSource * Zuora * Elavon
	PaymentGatewayDisplayName *string `json:"paymentGatewayDisplayName,omitempty"`
	PayPalLegacySettings *PayPalLegacySettings `json:"payPalLegacySettings,omitempty"`
	// A list of ISO 4217 currency codes for the currencies that the payment gateway account supports.  Examples:   - `USD` - `CAD` - `EUR` - `HKD`
	SupportedCurrencies []string `json:"supportedCurrencies,omitempty"`
	// An array of paymentMethodWithOptions objects that specify the payment methods that are available for the gateway.
	SupportedPaymentMethods []string `json:"supportedPaymentMethods,omitempty"`
	// An array of `paymentMethodWithOptions` objects that specify the payment methods that are available for the gateway, as well as the payment options that are compatible with each payment method.
	SupportedPaymentMethodsWithOptions []PaymentMethodWithOptions `json:"supportedPaymentMethodsWithOptions,omitempty"`
	// 
	ZeroDecimalCurrencies []string `json:"zeroDecimalCurrencies,omitempty"`
}

// NewPaymentGatewayAccounts instantiates a new PaymentGatewayAccounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentGatewayAccounts() *PaymentGatewayAccounts {
	this := PaymentGatewayAccounts{}
	return &this
}

// NewPaymentGatewayAccountsWithDefaults instantiates a new PaymentGatewayAccounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentGatewayAccountsWithDefaults() *PaymentGatewayAccounts {
	this := PaymentGatewayAccounts{}
	return &this
}

// GetAllowCustomMetadata returns the AllowCustomMetadata field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetAllowCustomMetadata() bool {
	if o == nil || IsNil(o.AllowCustomMetadata) {
		var ret bool
		return ret
	}
	return *o.AllowCustomMetadata
}

// GetAllowCustomMetadataOk returns a tuple with the AllowCustomMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetAllowCustomMetadataOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCustomMetadata) {
		return nil, false
	}
	return o.AllowCustomMetadata, true
}

// HasAllowCustomMetadata returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasAllowCustomMetadata() bool {
	if o != nil && !IsNil(o.AllowCustomMetadata) {
		return true
	}

	return false
}

// SetAllowCustomMetadata gets a reference to the given bool and assigns it to the AllowCustomMetadata field.
func (o *PaymentGatewayAccounts) SetAllowCustomMetadata(v bool) {
	o.AllowCustomMetadata = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetConfig() PaymentGatewayAccountSetting {
	if o == nil || IsNil(o.Config) {
		var ret PaymentGatewayAccountSetting
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetConfigOk() (*PaymentGatewayAccountSetting, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given PaymentGatewayAccountSetting and assigns it to the Config field.
func (o *PaymentGatewayAccounts) SetConfig(v PaymentGatewayAccountSetting) {
	o.Config = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *PaymentGatewayAccounts) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetIsEnabled() string {
	if o == nil || IsNil(o.IsEnabled) {
		var ret string
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetIsEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given string and assigns it to the IsEnabled field.
func (o *PaymentGatewayAccounts) SetIsEnabled(v string) {
	o.IsEnabled = &v
}

// GetIsLegacy returns the IsLegacy field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetIsLegacy() string {
	if o == nil || IsNil(o.IsLegacy) {
		var ret string
		return ret
	}
	return *o.IsLegacy
}

// GetIsLegacyOk returns a tuple with the IsLegacy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetIsLegacyOk() (*string, bool) {
	if o == nil || IsNil(o.IsLegacy) {
		return nil, false
	}
	return o.IsLegacy, true
}

// HasIsLegacy returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasIsLegacy() bool {
	if o != nil && !IsNil(o.IsLegacy) {
		return true
	}

	return false
}

// SetIsLegacy gets a reference to the given string and assigns it to the IsLegacy field.
func (o *PaymentGatewayAccounts) SetIsLegacy(v string) {
	o.IsLegacy = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetLastModified() string {
	if o == nil || IsNil(o.LastModified) {
		var ret string
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetLastModifiedOk() (*string, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given string and assigns it to the LastModified field.
func (o *PaymentGatewayAccounts) SetLastModified(v string) {
	o.LastModified = &v
}

// GetPaymentGateway returns the PaymentGateway field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetPaymentGateway() string {
	if o == nil || IsNil(o.PaymentGateway) {
		var ret string
		return ret
	}
	return *o.PaymentGateway
}

// GetPaymentGatewayOk returns a tuple with the PaymentGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetPaymentGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentGateway) {
		return nil, false
	}
	return o.PaymentGateway, true
}

// HasPaymentGateway returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasPaymentGateway() bool {
	if o != nil && !IsNil(o.PaymentGateway) {
		return true
	}

	return false
}

// SetPaymentGateway gets a reference to the given string and assigns it to the PaymentGateway field.
func (o *PaymentGatewayAccounts) SetPaymentGateway(v string) {
	o.PaymentGateway = &v
}

// GetPaymentGatewayAccountId returns the PaymentGatewayAccountId field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetPaymentGatewayAccountId() string {
	if o == nil || IsNil(o.PaymentGatewayAccountId) {
		var ret string
		return ret
	}
	return *o.PaymentGatewayAccountId
}

// GetPaymentGatewayAccountIdOk returns a tuple with the PaymentGatewayAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetPaymentGatewayAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentGatewayAccountId) {
		return nil, false
	}
	return o.PaymentGatewayAccountId, true
}

// HasPaymentGatewayAccountId returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasPaymentGatewayAccountId() bool {
	if o != nil && !IsNil(o.PaymentGatewayAccountId) {
		return true
	}

	return false
}

// SetPaymentGatewayAccountId gets a reference to the given string and assigns it to the PaymentGatewayAccountId field.
func (o *PaymentGatewayAccounts) SetPaymentGatewayAccountId(v string) {
	o.PaymentGatewayAccountId = &v
}

// GetPaymentGatewayDisplayName returns the PaymentGatewayDisplayName field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetPaymentGatewayDisplayName() string {
	if o == nil || IsNil(o.PaymentGatewayDisplayName) {
		var ret string
		return ret
	}
	return *o.PaymentGatewayDisplayName
}

// GetPaymentGatewayDisplayNameOk returns a tuple with the PaymentGatewayDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetPaymentGatewayDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentGatewayDisplayName) {
		return nil, false
	}
	return o.PaymentGatewayDisplayName, true
}

// HasPaymentGatewayDisplayName returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasPaymentGatewayDisplayName() bool {
	if o != nil && !IsNil(o.PaymentGatewayDisplayName) {
		return true
	}

	return false
}

// SetPaymentGatewayDisplayName gets a reference to the given string and assigns it to the PaymentGatewayDisplayName field.
func (o *PaymentGatewayAccounts) SetPaymentGatewayDisplayName(v string) {
	o.PaymentGatewayDisplayName = &v
}

// GetPayPalLegacySettings returns the PayPalLegacySettings field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetPayPalLegacySettings() PayPalLegacySettings {
	if o == nil || IsNil(o.PayPalLegacySettings) {
		var ret PayPalLegacySettings
		return ret
	}
	return *o.PayPalLegacySettings
}

// GetPayPalLegacySettingsOk returns a tuple with the PayPalLegacySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetPayPalLegacySettingsOk() (*PayPalLegacySettings, bool) {
	if o == nil || IsNil(o.PayPalLegacySettings) {
		return nil, false
	}
	return o.PayPalLegacySettings, true
}

// HasPayPalLegacySettings returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasPayPalLegacySettings() bool {
	if o != nil && !IsNil(o.PayPalLegacySettings) {
		return true
	}

	return false
}

// SetPayPalLegacySettings gets a reference to the given PayPalLegacySettings and assigns it to the PayPalLegacySettings field.
func (o *PaymentGatewayAccounts) SetPayPalLegacySettings(v PayPalLegacySettings) {
	o.PayPalLegacySettings = &v
}

// GetSupportedCurrencies returns the SupportedCurrencies field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetSupportedCurrencies() []string {
	if o == nil || IsNil(o.SupportedCurrencies) {
		var ret []string
		return ret
	}
	return o.SupportedCurrencies
}

// GetSupportedCurrenciesOk returns a tuple with the SupportedCurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetSupportedCurrenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedCurrencies) {
		return nil, false
	}
	return o.SupportedCurrencies, true
}

// HasSupportedCurrencies returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasSupportedCurrencies() bool {
	if o != nil && !IsNil(o.SupportedCurrencies) {
		return true
	}

	return false
}

// SetSupportedCurrencies gets a reference to the given []string and assigns it to the SupportedCurrencies field.
func (o *PaymentGatewayAccounts) SetSupportedCurrencies(v []string) {
	o.SupportedCurrencies = v
}

// GetSupportedPaymentMethods returns the SupportedPaymentMethods field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetSupportedPaymentMethods() []string {
	if o == nil || IsNil(o.SupportedPaymentMethods) {
		var ret []string
		return ret
	}
	return o.SupportedPaymentMethods
}

// GetSupportedPaymentMethodsOk returns a tuple with the SupportedPaymentMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetSupportedPaymentMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedPaymentMethods) {
		return nil, false
	}
	return o.SupportedPaymentMethods, true
}

// HasSupportedPaymentMethods returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasSupportedPaymentMethods() bool {
	if o != nil && !IsNil(o.SupportedPaymentMethods) {
		return true
	}

	return false
}

// SetSupportedPaymentMethods gets a reference to the given []string and assigns it to the SupportedPaymentMethods field.
func (o *PaymentGatewayAccounts) SetSupportedPaymentMethods(v []string) {
	o.SupportedPaymentMethods = v
}

// GetSupportedPaymentMethodsWithOptions returns the SupportedPaymentMethodsWithOptions field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetSupportedPaymentMethodsWithOptions() []PaymentMethodWithOptions {
	if o == nil || IsNil(o.SupportedPaymentMethodsWithOptions) {
		var ret []PaymentMethodWithOptions
		return ret
	}
	return o.SupportedPaymentMethodsWithOptions
}

// GetSupportedPaymentMethodsWithOptionsOk returns a tuple with the SupportedPaymentMethodsWithOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetSupportedPaymentMethodsWithOptionsOk() ([]PaymentMethodWithOptions, bool) {
	if o == nil || IsNil(o.SupportedPaymentMethodsWithOptions) {
		return nil, false
	}
	return o.SupportedPaymentMethodsWithOptions, true
}

// HasSupportedPaymentMethodsWithOptions returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasSupportedPaymentMethodsWithOptions() bool {
	if o != nil && !IsNil(o.SupportedPaymentMethodsWithOptions) {
		return true
	}

	return false
}

// SetSupportedPaymentMethodsWithOptions gets a reference to the given []PaymentMethodWithOptions and assigns it to the SupportedPaymentMethodsWithOptions field.
func (o *PaymentGatewayAccounts) SetSupportedPaymentMethodsWithOptions(v []PaymentMethodWithOptions) {
	o.SupportedPaymentMethodsWithOptions = v
}

// GetZeroDecimalCurrencies returns the ZeroDecimalCurrencies field value if set, zero value otherwise.
func (o *PaymentGatewayAccounts) GetZeroDecimalCurrencies() []string {
	if o == nil || IsNil(o.ZeroDecimalCurrencies) {
		var ret []string
		return ret
	}
	return o.ZeroDecimalCurrencies
}

// GetZeroDecimalCurrenciesOk returns a tuple with the ZeroDecimalCurrencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccounts) GetZeroDecimalCurrenciesOk() ([]string, bool) {
	if o == nil || IsNil(o.ZeroDecimalCurrencies) {
		return nil, false
	}
	return o.ZeroDecimalCurrencies, true
}

// HasZeroDecimalCurrencies returns a boolean if a field has been set.
func (o *PaymentGatewayAccounts) HasZeroDecimalCurrencies() bool {
	if o != nil && !IsNil(o.ZeroDecimalCurrencies) {
		return true
	}

	return false
}

// SetZeroDecimalCurrencies gets a reference to the given []string and assigns it to the ZeroDecimalCurrencies field.
func (o *PaymentGatewayAccounts) SetZeroDecimalCurrencies(v []string) {
	o.ZeroDecimalCurrencies = v
}

func (o PaymentGatewayAccounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentGatewayAccounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowCustomMetadata) {
		toSerialize["allowCustomMetadata"] = o.AllowCustomMetadata
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !IsNil(o.IsLegacy) {
		toSerialize["isLegacy"] = o.IsLegacy
	}
	if !IsNil(o.LastModified) {
		toSerialize["lastModified"] = o.LastModified
	}
	if !IsNil(o.PaymentGateway) {
		toSerialize["paymentGateway"] = o.PaymentGateway
	}
	if !IsNil(o.PaymentGatewayAccountId) {
		toSerialize["paymentGatewayAccountId"] = o.PaymentGatewayAccountId
	}
	if !IsNil(o.PaymentGatewayDisplayName) {
		toSerialize["paymentGatewayDisplayName"] = o.PaymentGatewayDisplayName
	}
	if !IsNil(o.PayPalLegacySettings) {
		toSerialize["payPalLegacySettings"] = o.PayPalLegacySettings
	}
	if !IsNil(o.SupportedCurrencies) {
		toSerialize["supportedCurrencies"] = o.SupportedCurrencies
	}
	if !IsNil(o.SupportedPaymentMethods) {
		toSerialize["supportedPaymentMethods"] = o.SupportedPaymentMethods
	}
	if !IsNil(o.SupportedPaymentMethodsWithOptions) {
		toSerialize["supportedPaymentMethodsWithOptions"] = o.SupportedPaymentMethodsWithOptions
	}
	if !IsNil(o.ZeroDecimalCurrencies) {
		toSerialize["zeroDecimalCurrencies"] = o.ZeroDecimalCurrencies
	}
	return toSerialize, nil
}

type NullablePaymentGatewayAccounts struct {
	value *PaymentGatewayAccounts
	isSet bool
}

func (v NullablePaymentGatewayAccounts) Get() *PaymentGatewayAccounts {
	return v.value
}

func (v *NullablePaymentGatewayAccounts) Set(val *PaymentGatewayAccounts) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentGatewayAccounts) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentGatewayAccounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentGatewayAccounts(val *PaymentGatewayAccounts) *NullablePaymentGatewayAccounts {
	return &NullablePaymentGatewayAccounts{value: val, isSet: true}
}

func (v NullablePaymentGatewayAccounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentGatewayAccounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


