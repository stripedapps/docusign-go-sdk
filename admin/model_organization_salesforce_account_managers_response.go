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

// checks if the OrganizationSalesforceAccountManagersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationSalesforceAccountManagersResponse{}

// OrganizationSalesforceAccountManagersResponse 
type OrganizationSalesforceAccountManagersResponse struct {
	// Select users that are members of the specified account. At least one of `email`, `account_id` or `organization_reserved_domain_id` must be specified.
	AccountId *string `json:"account_id,omitempty"`
	// 
	AccountName *string `json:"account_name,omitempty"`
	// 
	AccountType *string `json:"account_type,omitempty"`
	AccountOwner *OSAMRContact `json:"account_owner,omitempty"`
	AccountManager *OSAMRContact `json:"account_manager,omitempty"`
	ParentAccount *OrganizationSalesforceAccountManagersResponse `json:"parent_account,omitempty"`
}

// NewOrganizationSalesforceAccountManagersResponse instantiates a new OrganizationSalesforceAccountManagersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationSalesforceAccountManagersResponse() *OrganizationSalesforceAccountManagersResponse {
	this := OrganizationSalesforceAccountManagersResponse{}
	return &this
}

// NewOrganizationSalesforceAccountManagersResponseWithDefaults instantiates a new OrganizationSalesforceAccountManagersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationSalesforceAccountManagersResponseWithDefaults() *OrganizationSalesforceAccountManagersResponse {
	this := OrganizationSalesforceAccountManagersResponse{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *OrganizationSalesforceAccountManagersResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *OrganizationSalesforceAccountManagersResponse) SetAccountName(v string) {
	o.AccountName = &v
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *OrganizationSalesforceAccountManagersResponse) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAccountOwner returns the AccountOwner field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountOwner() OSAMRContact {
	if o == nil || IsNil(o.AccountOwner) {
		var ret OSAMRContact
		return ret
	}
	return *o.AccountOwner
}

// GetAccountOwnerOk returns a tuple with the AccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountOwnerOk() (*OSAMRContact, bool) {
	if o == nil || IsNil(o.AccountOwner) {
		return nil, false
	}
	return o.AccountOwner, true
}

// HasAccountOwner returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasAccountOwner() bool {
	if o != nil && !IsNil(o.AccountOwner) {
		return true
	}

	return false
}

// SetAccountOwner gets a reference to the given OSAMRContact and assigns it to the AccountOwner field.
func (o *OrganizationSalesforceAccountManagersResponse) SetAccountOwner(v OSAMRContact) {
	o.AccountOwner = &v
}

// GetAccountManager returns the AccountManager field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountManager() OSAMRContact {
	if o == nil || IsNil(o.AccountManager) {
		var ret OSAMRContact
		return ret
	}
	return *o.AccountManager
}

// GetAccountManagerOk returns a tuple with the AccountManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetAccountManagerOk() (*OSAMRContact, bool) {
	if o == nil || IsNil(o.AccountManager) {
		return nil, false
	}
	return o.AccountManager, true
}

// HasAccountManager returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasAccountManager() bool {
	if o != nil && !IsNil(o.AccountManager) {
		return true
	}

	return false
}

// SetAccountManager gets a reference to the given OSAMRContact and assigns it to the AccountManager field.
func (o *OrganizationSalesforceAccountManagersResponse) SetAccountManager(v OSAMRContact) {
	o.AccountManager = &v
}

// GetParentAccount returns the ParentAccount field value if set, zero value otherwise.
func (o *OrganizationSalesforceAccountManagersResponse) GetParentAccount() OrganizationSalesforceAccountManagersResponse {
	if o == nil || IsNil(o.ParentAccount) {
		var ret OrganizationSalesforceAccountManagersResponse
		return ret
	}
	return *o.ParentAccount
}

// GetParentAccountOk returns a tuple with the ParentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSalesforceAccountManagersResponse) GetParentAccountOk() (*OrganizationSalesforceAccountManagersResponse, bool) {
	if o == nil || IsNil(o.ParentAccount) {
		return nil, false
	}
	return o.ParentAccount, true
}

// HasParentAccount returns a boolean if a field has been set.
func (o *OrganizationSalesforceAccountManagersResponse) HasParentAccount() bool {
	if o != nil && !IsNil(o.ParentAccount) {
		return true
	}

	return false
}

// SetParentAccount gets a reference to the given OrganizationSalesforceAccountManagersResponse and assigns it to the ParentAccount field.
func (o *OrganizationSalesforceAccountManagersResponse) SetParentAccount(v OrganizationSalesforceAccountManagersResponse) {
	o.ParentAccount = &v
}

func (o OrganizationSalesforceAccountManagersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationSalesforceAccountManagersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	if !IsNil(o.AccountOwner) {
		toSerialize["account_owner"] = o.AccountOwner
	}
	if !IsNil(o.AccountManager) {
		toSerialize["account_manager"] = o.AccountManager
	}
	if !IsNil(o.ParentAccount) {
		toSerialize["parent_account"] = o.ParentAccount
	}
	return toSerialize, nil
}

type NullableOrganizationSalesforceAccountManagersResponse struct {
	value *OrganizationSalesforceAccountManagersResponse
	isSet bool
}

func (v NullableOrganizationSalesforceAccountManagersResponse) Get() *OrganizationSalesforceAccountManagersResponse {
	return v.value
}

func (v *NullableOrganizationSalesforceAccountManagersResponse) Set(val *OrganizationSalesforceAccountManagersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationSalesforceAccountManagersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationSalesforceAccountManagersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationSalesforceAccountManagersResponse(val *OrganizationSalesforceAccountManagersResponse) *NullableOrganizationSalesforceAccountManagersResponse {
	return &NullableOrganizationSalesforceAccountManagersResponse{value: val, isSet: true}
}

func (v NullableOrganizationSalesforceAccountManagersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationSalesforceAccountManagersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


