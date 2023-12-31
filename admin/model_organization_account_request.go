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

// checks if the OrganizationAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationAccountRequest{}

// OrganizationAccountRequest 
type OrganizationAccountRequest struct {
	// Select users that are members of the specified account. At least one of `email`, `account_id` or `organization_reserved_domain_id` must be specified.
	AccountId string `json:"account_id"`
}

// NewOrganizationAccountRequest instantiates a new OrganizationAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAccountRequest(accountId string) *OrganizationAccountRequest {
	this := OrganizationAccountRequest{}
	this.AccountId = accountId
	return &this
}

// NewOrganizationAccountRequestWithDefaults instantiates a new OrganizationAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAccountRequestWithDefaults() *OrganizationAccountRequest {
	this := OrganizationAccountRequest{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *OrganizationAccountRequest) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *OrganizationAccountRequest) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *OrganizationAccountRequest) SetAccountId(v string) {
	o.AccountId = v
}

func (o OrganizationAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	return toSerialize, nil
}

type NullableOrganizationAccountRequest struct {
	value *OrganizationAccountRequest
	isSet bool
}

func (v NullableOrganizationAccountRequest) Get() *OrganizationAccountRequest {
	return v.value
}

func (v *NullableOrganizationAccountRequest) Set(val *OrganizationAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAccountRequest(val *OrganizationAccountRequest) *NullableOrganizationAccountRequest {
	return &NullableOrganizationAccountRequest{value: val, isSet: true}
}

func (v NullableOrganizationAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


