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

// checks if the PaymentGatewayAccountSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentGatewayAccountSetting{}

// PaymentGatewayAccountSetting 
type PaymentGatewayAccountSetting struct {
	// 
	ApiFields *string `json:"apiFields,omitempty"`
	// 
	AuthorizationCode *string `json:"authorizationCode,omitempty"`
	// 
	CredentialStatus *string `json:"credentialStatus,omitempty"`
	// 
	MerchantId *string `json:"merchantId,omitempty"`
}

// NewPaymentGatewayAccountSetting instantiates a new PaymentGatewayAccountSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentGatewayAccountSetting() *PaymentGatewayAccountSetting {
	this := PaymentGatewayAccountSetting{}
	return &this
}

// NewPaymentGatewayAccountSettingWithDefaults instantiates a new PaymentGatewayAccountSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentGatewayAccountSettingWithDefaults() *PaymentGatewayAccountSetting {
	this := PaymentGatewayAccountSetting{}
	return &this
}

// GetApiFields returns the ApiFields field value if set, zero value otherwise.
func (o *PaymentGatewayAccountSetting) GetApiFields() string {
	if o == nil || IsNil(o.ApiFields) {
		var ret string
		return ret
	}
	return *o.ApiFields
}

// GetApiFieldsOk returns a tuple with the ApiFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccountSetting) GetApiFieldsOk() (*string, bool) {
	if o == nil || IsNil(o.ApiFields) {
		return nil, false
	}
	return o.ApiFields, true
}

// HasApiFields returns a boolean if a field has been set.
func (o *PaymentGatewayAccountSetting) HasApiFields() bool {
	if o != nil && !IsNil(o.ApiFields) {
		return true
	}

	return false
}

// SetApiFields gets a reference to the given string and assigns it to the ApiFields field.
func (o *PaymentGatewayAccountSetting) SetApiFields(v string) {
	o.ApiFields = &v
}

// GetAuthorizationCode returns the AuthorizationCode field value if set, zero value otherwise.
func (o *PaymentGatewayAccountSetting) GetAuthorizationCode() string {
	if o == nil || IsNil(o.AuthorizationCode) {
		var ret string
		return ret
	}
	return *o.AuthorizationCode
}

// GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccountSetting) GetAuthorizationCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationCode) {
		return nil, false
	}
	return o.AuthorizationCode, true
}

// HasAuthorizationCode returns a boolean if a field has been set.
func (o *PaymentGatewayAccountSetting) HasAuthorizationCode() bool {
	if o != nil && !IsNil(o.AuthorizationCode) {
		return true
	}

	return false
}

// SetAuthorizationCode gets a reference to the given string and assigns it to the AuthorizationCode field.
func (o *PaymentGatewayAccountSetting) SetAuthorizationCode(v string) {
	o.AuthorizationCode = &v
}

// GetCredentialStatus returns the CredentialStatus field value if set, zero value otherwise.
func (o *PaymentGatewayAccountSetting) GetCredentialStatus() string {
	if o == nil || IsNil(o.CredentialStatus) {
		var ret string
		return ret
	}
	return *o.CredentialStatus
}

// GetCredentialStatusOk returns a tuple with the CredentialStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccountSetting) GetCredentialStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CredentialStatus) {
		return nil, false
	}
	return o.CredentialStatus, true
}

// HasCredentialStatus returns a boolean if a field has been set.
func (o *PaymentGatewayAccountSetting) HasCredentialStatus() bool {
	if o != nil && !IsNil(o.CredentialStatus) {
		return true
	}

	return false
}

// SetCredentialStatus gets a reference to the given string and assigns it to the CredentialStatus field.
func (o *PaymentGatewayAccountSetting) SetCredentialStatus(v string) {
	o.CredentialStatus = &v
}

// GetMerchantId returns the MerchantId field value if set, zero value otherwise.
func (o *PaymentGatewayAccountSetting) GetMerchantId() string {
	if o == nil || IsNil(o.MerchantId) {
		var ret string
		return ret
	}
	return *o.MerchantId
}

// GetMerchantIdOk returns a tuple with the MerchantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentGatewayAccountSetting) GetMerchantIdOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantId) {
		return nil, false
	}
	return o.MerchantId, true
}

// HasMerchantId returns a boolean if a field has been set.
func (o *PaymentGatewayAccountSetting) HasMerchantId() bool {
	if o != nil && !IsNil(o.MerchantId) {
		return true
	}

	return false
}

// SetMerchantId gets a reference to the given string and assigns it to the MerchantId field.
func (o *PaymentGatewayAccountSetting) SetMerchantId(v string) {
	o.MerchantId = &v
}

func (o PaymentGatewayAccountSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentGatewayAccountSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiFields) {
		toSerialize["apiFields"] = o.ApiFields
	}
	if !IsNil(o.AuthorizationCode) {
		toSerialize["authorizationCode"] = o.AuthorizationCode
	}
	if !IsNil(o.CredentialStatus) {
		toSerialize["credentialStatus"] = o.CredentialStatus
	}
	if !IsNil(o.MerchantId) {
		toSerialize["merchantId"] = o.MerchantId
	}
	return toSerialize, nil
}

type NullablePaymentGatewayAccountSetting struct {
	value *PaymentGatewayAccountSetting
	isSet bool
}

func (v NullablePaymentGatewayAccountSetting) Get() *PaymentGatewayAccountSetting {
	return v.value
}

func (v *NullablePaymentGatewayAccountSetting) Set(val *PaymentGatewayAccountSetting) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentGatewayAccountSetting) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentGatewayAccountSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentGatewayAccountSetting(val *PaymentGatewayAccountSetting) *NullablePaymentGatewayAccountSetting {
	return &NullablePaymentGatewayAccountSetting{value: val, isSet: true}
}

func (v NullablePaymentGatewayAccountSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentGatewayAccountSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


