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

// checks if the RemoveUserProductsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveUserProductsResponse{}

// RemoveUserProductsResponse 
type RemoveUserProductsResponse struct {
	// When **true,** indicates that the user's access to the specified products was successfully revoked.
	IsSuccess bool `json:"is_success"`
	// The user's email address.
	UserEmail *string `json:"user_email,omitempty"`
	// The user's ID.
	UserId *string `json:"user_id,omitempty"`
	// The results of the request to revoke product access.  Each key in the object is the ID of a product specified in the request. Each corresponding value indicates whether the user's access for that product was successfully revoked. If successful, the value is `deleted`. Otherwise, the value describes the error that occurred.  For example: ``` {         \"230546a7-xxxx-xxxx-xxxx-af205d5494ad\"\": \"deleted\",         \"984800b7-xxxx-xxxx-xxxx-kt374a5922lk\": \"Invalid product id\" } ```
	UserProductResults map[string]string `json:"user_product_results"`
}

// NewRemoveUserProductsResponse instantiates a new RemoveUserProductsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveUserProductsResponse(isSuccess bool, userProductResults map[string]string) *RemoveUserProductsResponse {
	this := RemoveUserProductsResponse{}
	this.IsSuccess = isSuccess
	this.UserProductResults = userProductResults
	return &this
}

// NewRemoveUserProductsResponseWithDefaults instantiates a new RemoveUserProductsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveUserProductsResponseWithDefaults() *RemoveUserProductsResponse {
	this := RemoveUserProductsResponse{}
	return &this
}

// GetIsSuccess returns the IsSuccess field value
func (o *RemoveUserProductsResponse) GetIsSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSuccess
}

// GetIsSuccessOk returns a tuple with the IsSuccess field value
// and a boolean to check if the value has been set.
func (o *RemoveUserProductsResponse) GetIsSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSuccess, true
}

// SetIsSuccess sets field value
func (o *RemoveUserProductsResponse) SetIsSuccess(v bool) {
	o.IsSuccess = v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *RemoveUserProductsResponse) GetUserEmail() string {
	if o == nil || IsNil(o.UserEmail) {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveUserProductsResponse) GetUserEmailOk() (*string, bool) {
	if o == nil || IsNil(o.UserEmail) {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *RemoveUserProductsResponse) HasUserEmail() bool {
	if o != nil && !IsNil(o.UserEmail) {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *RemoveUserProductsResponse) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RemoveUserProductsResponse) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemoveUserProductsResponse) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RemoveUserProductsResponse) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RemoveUserProductsResponse) SetUserId(v string) {
	o.UserId = &v
}

// GetUserProductResults returns the UserProductResults field value
func (o *RemoveUserProductsResponse) GetUserProductResults() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.UserProductResults
}

// GetUserProductResultsOk returns a tuple with the UserProductResults field value
// and a boolean to check if the value has been set.
func (o *RemoveUserProductsResponse) GetUserProductResultsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserProductResults, true
}

// SetUserProductResults sets field value
func (o *RemoveUserProductsResponse) SetUserProductResults(v map[string]string) {
	o.UserProductResults = v
}

func (o RemoveUserProductsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveUserProductsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_success"] = o.IsSuccess
	if !IsNil(o.UserEmail) {
		toSerialize["user_email"] = o.UserEmail
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	toSerialize["user_product_results"] = o.UserProductResults
	return toSerialize, nil
}

type NullableRemoveUserProductsResponse struct {
	value *RemoveUserProductsResponse
	isSet bool
}

func (v NullableRemoveUserProductsResponse) Get() *RemoveUserProductsResponse {
	return v.value
}

func (v *NullableRemoveUserProductsResponse) Set(val *RemoveUserProductsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveUserProductsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveUserProductsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveUserProductsResponse(val *RemoveUserProductsResponse) *NullableRemoveUserProductsResponse {
	return &NullableRemoveUserProductsResponse{value: val, isSet: true}
}

func (v NullableRemoveUserProductsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveUserProductsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

