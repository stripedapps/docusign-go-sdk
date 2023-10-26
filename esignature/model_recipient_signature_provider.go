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

// checks if the RecipientSignatureProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientSignatureProvider{}

// RecipientSignatureProvider An Electronic or Standards Based Signature (digital signature) provider for the signer to use. [More information](/docs/esign-rest-api/esign101/concepts/standards-based-signatures/). 
type RecipientSignatureProvider struct {
	// By default, electronic seals apply on all documents in an envelope. If any of the documents has a `signHere` tab, then a visual representation of the electronic seal will show up in the final document. If not, the electronic seal will be visible in the metadata but not in the content of the document.  To apply electronic seals on specific documents only, you must enable the  `sealDocumentsWithTabsOnly` parameter. In this case, Electronic Seal applies only on documents that have `signHere` tabs set for the Electronic Seal recipient. Other documents won't be sealed. 
	SealDocumentsWithTabsOnly *string `json:"sealDocumentsWithTabsOnly,omitempty"`
	// Indicates the name of the electronic seal to apply on documents. 
	SealName *string `json:"sealName,omitempty"`
	// The name of an Electronic or Standards Based Signature (digital signature) provider for the signer to use. For details, see [the current provider list](/docs/esign-rest-api/esign101/concepts/standards-based-signatures/). You can also retrieve the list by using the [AccountSignatureProviders: List](/docs/esign-rest-api/reference/accounts/accountsignatureproviders/list/) method.  Example: `universalsignaturepen_default`  
	SignatureProviderName *string `json:"signatureProviderName,omitempty"`
	SignatureProviderNameMetadata *PropertyMetadata `json:"signatureProviderNameMetadata,omitempty"`
	SignatureProviderOptions *RecipientSignatureProviderOptions `json:"signatureProviderOptions,omitempty"`
}

// NewRecipientSignatureProvider instantiates a new RecipientSignatureProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientSignatureProvider() *RecipientSignatureProvider {
	this := RecipientSignatureProvider{}
	return &this
}

// NewRecipientSignatureProviderWithDefaults instantiates a new RecipientSignatureProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientSignatureProviderWithDefaults() *RecipientSignatureProvider {
	this := RecipientSignatureProvider{}
	return &this
}

// GetSealDocumentsWithTabsOnly returns the SealDocumentsWithTabsOnly field value if set, zero value otherwise.
func (o *RecipientSignatureProvider) GetSealDocumentsWithTabsOnly() string {
	if o == nil || IsNil(o.SealDocumentsWithTabsOnly) {
		var ret string
		return ret
	}
	return *o.SealDocumentsWithTabsOnly
}

// GetSealDocumentsWithTabsOnlyOk returns a tuple with the SealDocumentsWithTabsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProvider) GetSealDocumentsWithTabsOnlyOk() (*string, bool) {
	if o == nil || IsNil(o.SealDocumentsWithTabsOnly) {
		return nil, false
	}
	return o.SealDocumentsWithTabsOnly, true
}

// HasSealDocumentsWithTabsOnly returns a boolean if a field has been set.
func (o *RecipientSignatureProvider) HasSealDocumentsWithTabsOnly() bool {
	if o != nil && !IsNil(o.SealDocumentsWithTabsOnly) {
		return true
	}

	return false
}

// SetSealDocumentsWithTabsOnly gets a reference to the given string and assigns it to the SealDocumentsWithTabsOnly field.
func (o *RecipientSignatureProvider) SetSealDocumentsWithTabsOnly(v string) {
	o.SealDocumentsWithTabsOnly = &v
}

// GetSealName returns the SealName field value if set, zero value otherwise.
func (o *RecipientSignatureProvider) GetSealName() string {
	if o == nil || IsNil(o.SealName) {
		var ret string
		return ret
	}
	return *o.SealName
}

// GetSealNameOk returns a tuple with the SealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProvider) GetSealNameOk() (*string, bool) {
	if o == nil || IsNil(o.SealName) {
		return nil, false
	}
	return o.SealName, true
}

// HasSealName returns a boolean if a field has been set.
func (o *RecipientSignatureProvider) HasSealName() bool {
	if o != nil && !IsNil(o.SealName) {
		return true
	}

	return false
}

// SetSealName gets a reference to the given string and assigns it to the SealName field.
func (o *RecipientSignatureProvider) SetSealName(v string) {
	o.SealName = &v
}

// GetSignatureProviderName returns the SignatureProviderName field value if set, zero value otherwise.
func (o *RecipientSignatureProvider) GetSignatureProviderName() string {
	if o == nil || IsNil(o.SignatureProviderName) {
		var ret string
		return ret
	}
	return *o.SignatureProviderName
}

// GetSignatureProviderNameOk returns a tuple with the SignatureProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProvider) GetSignatureProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureProviderName) {
		return nil, false
	}
	return o.SignatureProviderName, true
}

// HasSignatureProviderName returns a boolean if a field has been set.
func (o *RecipientSignatureProvider) HasSignatureProviderName() bool {
	if o != nil && !IsNil(o.SignatureProviderName) {
		return true
	}

	return false
}

// SetSignatureProviderName gets a reference to the given string and assigns it to the SignatureProviderName field.
func (o *RecipientSignatureProvider) SetSignatureProviderName(v string) {
	o.SignatureProviderName = &v
}

// GetSignatureProviderNameMetadata returns the SignatureProviderNameMetadata field value if set, zero value otherwise.
func (o *RecipientSignatureProvider) GetSignatureProviderNameMetadata() PropertyMetadata {
	if o == nil || IsNil(o.SignatureProviderNameMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.SignatureProviderNameMetadata
}

// GetSignatureProviderNameMetadataOk returns a tuple with the SignatureProviderNameMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProvider) GetSignatureProviderNameMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.SignatureProviderNameMetadata) {
		return nil, false
	}
	return o.SignatureProviderNameMetadata, true
}

// HasSignatureProviderNameMetadata returns a boolean if a field has been set.
func (o *RecipientSignatureProvider) HasSignatureProviderNameMetadata() bool {
	if o != nil && !IsNil(o.SignatureProviderNameMetadata) {
		return true
	}

	return false
}

// SetSignatureProviderNameMetadata gets a reference to the given PropertyMetadata and assigns it to the SignatureProviderNameMetadata field.
func (o *RecipientSignatureProvider) SetSignatureProviderNameMetadata(v PropertyMetadata) {
	o.SignatureProviderNameMetadata = &v
}

// GetSignatureProviderOptions returns the SignatureProviderOptions field value if set, zero value otherwise.
func (o *RecipientSignatureProvider) GetSignatureProviderOptions() RecipientSignatureProviderOptions {
	if o == nil || IsNil(o.SignatureProviderOptions) {
		var ret RecipientSignatureProviderOptions
		return ret
	}
	return *o.SignatureProviderOptions
}

// GetSignatureProviderOptionsOk returns a tuple with the SignatureProviderOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientSignatureProvider) GetSignatureProviderOptionsOk() (*RecipientSignatureProviderOptions, bool) {
	if o == nil || IsNil(o.SignatureProviderOptions) {
		return nil, false
	}
	return o.SignatureProviderOptions, true
}

// HasSignatureProviderOptions returns a boolean if a field has been set.
func (o *RecipientSignatureProvider) HasSignatureProviderOptions() bool {
	if o != nil && !IsNil(o.SignatureProviderOptions) {
		return true
	}

	return false
}

// SetSignatureProviderOptions gets a reference to the given RecipientSignatureProviderOptions and assigns it to the SignatureProviderOptions field.
func (o *RecipientSignatureProvider) SetSignatureProviderOptions(v RecipientSignatureProviderOptions) {
	o.SignatureProviderOptions = &v
}

func (o RecipientSignatureProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientSignatureProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SealDocumentsWithTabsOnly) {
		toSerialize["sealDocumentsWithTabsOnly"] = o.SealDocumentsWithTabsOnly
	}
	if !IsNil(o.SealName) {
		toSerialize["sealName"] = o.SealName
	}
	if !IsNil(o.SignatureProviderName) {
		toSerialize["signatureProviderName"] = o.SignatureProviderName
	}
	if !IsNil(o.SignatureProviderNameMetadata) {
		toSerialize["signatureProviderNameMetadata"] = o.SignatureProviderNameMetadata
	}
	if !IsNil(o.SignatureProviderOptions) {
		toSerialize["signatureProviderOptions"] = o.SignatureProviderOptions
	}
	return toSerialize, nil
}

type NullableRecipientSignatureProvider struct {
	value *RecipientSignatureProvider
	isSet bool
}

func (v NullableRecipientSignatureProvider) Get() *RecipientSignatureProvider {
	return v.value
}

func (v *NullableRecipientSignatureProvider) Set(val *RecipientSignatureProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientSignatureProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientSignatureProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientSignatureProvider(val *RecipientSignatureProvider) *NullableRecipientSignatureProvider {
	return &NullableRecipientSignatureProvider{value: val, isSet: true}
}

func (v NullableRecipientSignatureProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientSignatureProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


