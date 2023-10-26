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

// checks if the IdEvidenceResourceToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdEvidenceResourceToken{}

// IdEvidenceResourceToken 
type IdEvidenceResourceToken struct {
	// 
	ProofBaseURI *string `json:"proofBaseURI,omitempty"`
	// 
	ResourceToken *string `json:"resourceToken,omitempty"`
}

// NewIdEvidenceResourceToken instantiates a new IdEvidenceResourceToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdEvidenceResourceToken() *IdEvidenceResourceToken {
	this := IdEvidenceResourceToken{}
	return &this
}

// NewIdEvidenceResourceTokenWithDefaults instantiates a new IdEvidenceResourceToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdEvidenceResourceTokenWithDefaults() *IdEvidenceResourceToken {
	this := IdEvidenceResourceToken{}
	return &this
}

// GetProofBaseURI returns the ProofBaseURI field value if set, zero value otherwise.
func (o *IdEvidenceResourceToken) GetProofBaseURI() string {
	if o == nil || IsNil(o.ProofBaseURI) {
		var ret string
		return ret
	}
	return *o.ProofBaseURI
}

// GetProofBaseURIOk returns a tuple with the ProofBaseURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdEvidenceResourceToken) GetProofBaseURIOk() (*string, bool) {
	if o == nil || IsNil(o.ProofBaseURI) {
		return nil, false
	}
	return o.ProofBaseURI, true
}

// HasProofBaseURI returns a boolean if a field has been set.
func (o *IdEvidenceResourceToken) HasProofBaseURI() bool {
	if o != nil && !IsNil(o.ProofBaseURI) {
		return true
	}

	return false
}

// SetProofBaseURI gets a reference to the given string and assigns it to the ProofBaseURI field.
func (o *IdEvidenceResourceToken) SetProofBaseURI(v string) {
	o.ProofBaseURI = &v
}

// GetResourceToken returns the ResourceToken field value if set, zero value otherwise.
func (o *IdEvidenceResourceToken) GetResourceToken() string {
	if o == nil || IsNil(o.ResourceToken) {
		var ret string
		return ret
	}
	return *o.ResourceToken
}

// GetResourceTokenOk returns a tuple with the ResourceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdEvidenceResourceToken) GetResourceTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceToken) {
		return nil, false
	}
	return o.ResourceToken, true
}

// HasResourceToken returns a boolean if a field has been set.
func (o *IdEvidenceResourceToken) HasResourceToken() bool {
	if o != nil && !IsNil(o.ResourceToken) {
		return true
	}

	return false
}

// SetResourceToken gets a reference to the given string and assigns it to the ResourceToken field.
func (o *IdEvidenceResourceToken) SetResourceToken(v string) {
	o.ResourceToken = &v
}

func (o IdEvidenceResourceToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdEvidenceResourceToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProofBaseURI) {
		toSerialize["proofBaseURI"] = o.ProofBaseURI
	}
	if !IsNil(o.ResourceToken) {
		toSerialize["resourceToken"] = o.ResourceToken
	}
	return toSerialize, nil
}

type NullableIdEvidenceResourceToken struct {
	value *IdEvidenceResourceToken
	isSet bool
}

func (v NullableIdEvidenceResourceToken) Get() *IdEvidenceResourceToken {
	return v.value
}

func (v *NullableIdEvidenceResourceToken) Set(val *IdEvidenceResourceToken) {
	v.value = val
	v.isSet = true
}

func (v NullableIdEvidenceResourceToken) IsSet() bool {
	return v.isSet
}

func (v *NullableIdEvidenceResourceToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdEvidenceResourceToken(val *IdEvidenceResourceToken) *NullableIdEvidenceResourceToken {
	return &NullableIdEvidenceResourceToken{value: val, isSet: true}
}

func (v NullableIdEvidenceResourceToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdEvidenceResourceToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


