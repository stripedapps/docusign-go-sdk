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

// checks if the AccountSignatureProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSignatureProvider{}

// AccountSignatureProvider Contains information about the signature provider associated with the Identity Verification workflow. If empty, then this specific workflow is not intended for signers. 
type AccountSignatureProvider struct {
	// Reserved for DocuSign.
	IsRequired *string `json:"isRequired,omitempty"`
	// Reserved for DocuSign.
	Priority *string `json:"priority,omitempty"`
	// Reserved for DocuSign.
	SignatureProviderDisplayName *string `json:"signatureProviderDisplayName,omitempty"`
	// Reserved for DocuSign.
	SignatureProviderId *string `json:"signatureProviderId,omitempty"`
	// The name of an Electronic or Standards Based Signature (digital signature) provider for the signer to use. For details, see [the current provider list](/docs/esign-rest-api/esign101/concepts/standards-based-signatures/). You can also retrieve the list by using the [AccountSignatureProviders: List](/docs/esign-rest-api/reference/accounts/accountsignatureproviders/list/) method.  Example: `universalsignaturepen_default`  
	SignatureProviderName *string `json:"signatureProviderName,omitempty"`
	// Reserved for DocuSign.
	SignatureProviderOptionsMetadata []AccountSignatureProviderOption `json:"signatureProviderOptionsMetadata,omitempty"`
	// Reserved for DocuSign.
	SignatureProviderRequiredOptions []SignatureProviderRequiredOption `json:"signatureProviderRequiredOptions,omitempty"`
}

// NewAccountSignatureProvider instantiates a new AccountSignatureProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSignatureProvider() *AccountSignatureProvider {
	this := AccountSignatureProvider{}
	return &this
}

// NewAccountSignatureProviderWithDefaults instantiates a new AccountSignatureProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSignatureProviderWithDefaults() *AccountSignatureProvider {
	this := AccountSignatureProvider{}
	return &this
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetIsRequired() string {
	if o == nil || IsNil(o.IsRequired) {
		var ret string
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetIsRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given string and assigns it to the IsRequired field.
func (o *AccountSignatureProvider) SetIsRequired(v string) {
	o.IsRequired = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetPriority() string {
	if o == nil || IsNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetPriorityOk() (*string, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *AccountSignatureProvider) SetPriority(v string) {
	o.Priority = &v
}

// GetSignatureProviderDisplayName returns the SignatureProviderDisplayName field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetSignatureProviderDisplayName() string {
	if o == nil || IsNil(o.SignatureProviderDisplayName) {
		var ret string
		return ret
	}
	return *o.SignatureProviderDisplayName
}

// GetSignatureProviderDisplayNameOk returns a tuple with the SignatureProviderDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetSignatureProviderDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureProviderDisplayName) {
		return nil, false
	}
	return o.SignatureProviderDisplayName, true
}

// HasSignatureProviderDisplayName returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasSignatureProviderDisplayName() bool {
	if o != nil && !IsNil(o.SignatureProviderDisplayName) {
		return true
	}

	return false
}

// SetSignatureProviderDisplayName gets a reference to the given string and assigns it to the SignatureProviderDisplayName field.
func (o *AccountSignatureProvider) SetSignatureProviderDisplayName(v string) {
	o.SignatureProviderDisplayName = &v
}

// GetSignatureProviderId returns the SignatureProviderId field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetSignatureProviderId() string {
	if o == nil || IsNil(o.SignatureProviderId) {
		var ret string
		return ret
	}
	return *o.SignatureProviderId
}

// GetSignatureProviderIdOk returns a tuple with the SignatureProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetSignatureProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureProviderId) {
		return nil, false
	}
	return o.SignatureProviderId, true
}

// HasSignatureProviderId returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasSignatureProviderId() bool {
	if o != nil && !IsNil(o.SignatureProviderId) {
		return true
	}

	return false
}

// SetSignatureProviderId gets a reference to the given string and assigns it to the SignatureProviderId field.
func (o *AccountSignatureProvider) SetSignatureProviderId(v string) {
	o.SignatureProviderId = &v
}

// GetSignatureProviderName returns the SignatureProviderName field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetSignatureProviderName() string {
	if o == nil || IsNil(o.SignatureProviderName) {
		var ret string
		return ret
	}
	return *o.SignatureProviderName
}

// GetSignatureProviderNameOk returns a tuple with the SignatureProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetSignatureProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureProviderName) {
		return nil, false
	}
	return o.SignatureProviderName, true
}

// HasSignatureProviderName returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasSignatureProviderName() bool {
	if o != nil && !IsNil(o.SignatureProviderName) {
		return true
	}

	return false
}

// SetSignatureProviderName gets a reference to the given string and assigns it to the SignatureProviderName field.
func (o *AccountSignatureProvider) SetSignatureProviderName(v string) {
	o.SignatureProviderName = &v
}

// GetSignatureProviderOptionsMetadata returns the SignatureProviderOptionsMetadata field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetSignatureProviderOptionsMetadata() []AccountSignatureProviderOption {
	if o == nil || IsNil(o.SignatureProviderOptionsMetadata) {
		var ret []AccountSignatureProviderOption
		return ret
	}
	return o.SignatureProviderOptionsMetadata
}

// GetSignatureProviderOptionsMetadataOk returns a tuple with the SignatureProviderOptionsMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetSignatureProviderOptionsMetadataOk() ([]AccountSignatureProviderOption, bool) {
	if o == nil || IsNil(o.SignatureProviderOptionsMetadata) {
		return nil, false
	}
	return o.SignatureProviderOptionsMetadata, true
}

// HasSignatureProviderOptionsMetadata returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasSignatureProviderOptionsMetadata() bool {
	if o != nil && !IsNil(o.SignatureProviderOptionsMetadata) {
		return true
	}

	return false
}

// SetSignatureProviderOptionsMetadata gets a reference to the given []AccountSignatureProviderOption and assigns it to the SignatureProviderOptionsMetadata field.
func (o *AccountSignatureProvider) SetSignatureProviderOptionsMetadata(v []AccountSignatureProviderOption) {
	o.SignatureProviderOptionsMetadata = v
}

// GetSignatureProviderRequiredOptions returns the SignatureProviderRequiredOptions field value if set, zero value otherwise.
func (o *AccountSignatureProvider) GetSignatureProviderRequiredOptions() []SignatureProviderRequiredOption {
	if o == nil || IsNil(o.SignatureProviderRequiredOptions) {
		var ret []SignatureProviderRequiredOption
		return ret
	}
	return o.SignatureProviderRequiredOptions
}

// GetSignatureProviderRequiredOptionsOk returns a tuple with the SignatureProviderRequiredOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureProvider) GetSignatureProviderRequiredOptionsOk() ([]SignatureProviderRequiredOption, bool) {
	if o == nil || IsNil(o.SignatureProviderRequiredOptions) {
		return nil, false
	}
	return o.SignatureProviderRequiredOptions, true
}

// HasSignatureProviderRequiredOptions returns a boolean if a field has been set.
func (o *AccountSignatureProvider) HasSignatureProviderRequiredOptions() bool {
	if o != nil && !IsNil(o.SignatureProviderRequiredOptions) {
		return true
	}

	return false
}

// SetSignatureProviderRequiredOptions gets a reference to the given []SignatureProviderRequiredOption and assigns it to the SignatureProviderRequiredOptions field.
func (o *AccountSignatureProvider) SetSignatureProviderRequiredOptions(v []SignatureProviderRequiredOption) {
	o.SignatureProviderRequiredOptions = v
}

func (o AccountSignatureProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSignatureProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsRequired) {
		toSerialize["isRequired"] = o.IsRequired
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.SignatureProviderDisplayName) {
		toSerialize["signatureProviderDisplayName"] = o.SignatureProviderDisplayName
	}
	if !IsNil(o.SignatureProviderId) {
		toSerialize["signatureProviderId"] = o.SignatureProviderId
	}
	if !IsNil(o.SignatureProviderName) {
		toSerialize["signatureProviderName"] = o.SignatureProviderName
	}
	if !IsNil(o.SignatureProviderOptionsMetadata) {
		toSerialize["signatureProviderOptionsMetadata"] = o.SignatureProviderOptionsMetadata
	}
	if !IsNil(o.SignatureProviderRequiredOptions) {
		toSerialize["signatureProviderRequiredOptions"] = o.SignatureProviderRequiredOptions
	}
	return toSerialize, nil
}

type NullableAccountSignatureProvider struct {
	value *AccountSignatureProvider
	isSet bool
}

func (v NullableAccountSignatureProvider) Get() *AccountSignatureProvider {
	return v.value
}

func (v *NullableAccountSignatureProvider) Set(val *AccountSignatureProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSignatureProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSignatureProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSignatureProvider(val *AccountSignatureProvider) *NullableAccountSignatureProvider {
	return &NullableAccountSignatureProvider{value: val, isSet: true}
}

func (v NullableAccountSignatureProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSignatureProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

