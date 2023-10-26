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

// checks if the UserProductPermissionProfilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProductPermissionProfilesResponse{}

// UserProductPermissionProfilesResponse 
type UserProductPermissionProfilesResponse struct {
	// The ID of the user.
	UserId *string `json:"user_id,omitempty"`
	// The ID of the account.
	AccountId *string `json:"account_id,omitempty"`
	// A list of one or more products and their respective permissions.
	ProductPermissionProfiles []ProductPermissionProfileResponse `json:"product_permission_profiles,omitempty"`
}

// NewUserProductPermissionProfilesResponse instantiates a new UserProductPermissionProfilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProductPermissionProfilesResponse() *UserProductPermissionProfilesResponse {
	this := UserProductPermissionProfilesResponse{}
	return &this
}

// NewUserProductPermissionProfilesResponseWithDefaults instantiates a new UserProductPermissionProfilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProductPermissionProfilesResponseWithDefaults() *UserProductPermissionProfilesResponse {
	this := UserProductPermissionProfilesResponse{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserProductPermissionProfilesResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProductPermissionProfilesResponse) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserProductPermissionProfilesResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserProductPermissionProfilesResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *UserProductPermissionProfilesResponse) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProductPermissionProfilesResponse) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *UserProductPermissionProfilesResponse) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *UserProductPermissionProfilesResponse) SetAccountId(v string) {
	o.AccountId = &v
}

// GetProductPermissionProfiles returns the ProductPermissionProfiles field value if set, zero value otherwise.
func (o *UserProductPermissionProfilesResponse) GetProductPermissionProfiles() []ProductPermissionProfileResponse {
	if o == nil || IsNil(o.ProductPermissionProfiles) {
		var ret []ProductPermissionProfileResponse
		return ret
	}
	return o.ProductPermissionProfiles
}

// GetProductPermissionProfilesOk returns a tuple with the ProductPermissionProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProductPermissionProfilesResponse) GetProductPermissionProfilesOk() ([]ProductPermissionProfileResponse, bool) {
	if o == nil || IsNil(o.ProductPermissionProfiles) {
		return nil, false
	}
	return o.ProductPermissionProfiles, true
}

// HasProductPermissionProfiles returns a boolean if a field has been set.
func (o *UserProductPermissionProfilesResponse) HasProductPermissionProfiles() bool {
	if o != nil && !IsNil(o.ProductPermissionProfiles) {
		return true
	}

	return false
}

// SetProductPermissionProfiles gets a reference to the given []ProductPermissionProfileResponse and assigns it to the ProductPermissionProfiles field.
func (o *UserProductPermissionProfilesResponse) SetProductPermissionProfiles(v []ProductPermissionProfileResponse) {
	o.ProductPermissionProfiles = v
}

func (o UserProductPermissionProfilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProductPermissionProfilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.ProductPermissionProfiles) {
		toSerialize["product_permission_profiles"] = o.ProductPermissionProfiles
	}
	return toSerialize, nil
}

type NullableUserProductPermissionProfilesResponse struct {
	value *UserProductPermissionProfilesResponse
	isSet bool
}

func (v NullableUserProductPermissionProfilesResponse) Get() *UserProductPermissionProfilesResponse {
	return v.value
}

func (v *NullableUserProductPermissionProfilesResponse) Set(val *UserProductPermissionProfilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProductPermissionProfilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProductPermissionProfilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProductPermissionProfilesResponse(val *UserProductPermissionProfilesResponse) *NullableUserProductPermissionProfilesResponse {
	return &NullableUserProductPermissionProfilesResponse{value: val, isSet: true}
}

func (v NullableUserProductPermissionProfilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProductPermissionProfilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


