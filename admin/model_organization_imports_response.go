/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the OrganizationImportsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationImportsResponse{}

// OrganizationImportsResponse 
type OrganizationImportsResponse struct {
	// A list of responses to user import requests.
	Imports []OrganizationImportResponse `json:"imports,omitempty"`
}

// NewOrganizationImportsResponse instantiates a new OrganizationImportsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationImportsResponse() *OrganizationImportsResponse {
	this := OrganizationImportsResponse{}
	return &this
}

// NewOrganizationImportsResponseWithDefaults instantiates a new OrganizationImportsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationImportsResponseWithDefaults() *OrganizationImportsResponse {
	this := OrganizationImportsResponse{}
	return &this
}

// GetImports returns the Imports field value if set, zero value otherwise.
func (o *OrganizationImportsResponse) GetImports() []OrganizationImportResponse {
	if o == nil || IsNil(o.Imports) {
		var ret []OrganizationImportResponse
		return ret
	}
	return o.Imports
}

// GetImportsOk returns a tuple with the Imports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportsResponse) GetImportsOk() ([]OrganizationImportResponse, bool) {
	if o == nil || IsNil(o.Imports) {
		return nil, false
	}
	return o.Imports, true
}

// HasImports returns a boolean if a field has been set.
func (o *OrganizationImportsResponse) HasImports() bool {
	if o != nil && !IsNil(o.Imports) {
		return true
	}

	return false
}

// SetImports gets a reference to the given []OrganizationImportResponse and assigns it to the Imports field.
func (o *OrganizationImportsResponse) SetImports(v []OrganizationImportResponse) {
	o.Imports = v
}

func (o OrganizationImportsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationImportsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Imports) {
		toSerialize["imports"] = o.Imports
	}
	return toSerialize, nil
}

type NullableOrganizationImportsResponse struct {
	value *OrganizationImportsResponse
	isSet bool
}

func (v NullableOrganizationImportsResponse) Get() *OrganizationImportsResponse {
	return v.value
}

func (v *NullableOrganizationImportsResponse) Set(val *OrganizationImportsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationImportsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationImportsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationImportsResponse(val *OrganizationImportsResponse) *NullableOrganizationImportsResponse {
	return &NullableOrganizationImportsResponse{value: val, isSet: true}
}

func (v NullableOrganizationImportsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationImportsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

