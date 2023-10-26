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

// checks if the UserSharedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSharedItem{}

// UserSharedItem Information about a shared item.
type UserSharedItem struct {
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// How the item is shared. One of:  - `not_shared` - `shared_to` - `shared_from` - `shared_to_and_from` 
	Shared *string `json:"shared,omitempty"`
	User *UserInfo `json:"user,omitempty"`
}

// NewUserSharedItem instantiates a new UserSharedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSharedItem() *UserSharedItem {
	this := UserSharedItem{}
	return &this
}

// NewUserSharedItemWithDefaults instantiates a new UserSharedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSharedItemWithDefaults() *UserSharedItem {
	this := UserSharedItem{}
	return &this
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *UserSharedItem) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSharedItem) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *UserSharedItem) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *UserSharedItem) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetShared returns the Shared field value if set, zero value otherwise.
func (o *UserSharedItem) GetShared() string {
	if o == nil || IsNil(o.Shared) {
		var ret string
		return ret
	}
	return *o.Shared
}

// GetSharedOk returns a tuple with the Shared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSharedItem) GetSharedOk() (*string, bool) {
	if o == nil || IsNil(o.Shared) {
		return nil, false
	}
	return o.Shared, true
}

// HasShared returns a boolean if a field has been set.
func (o *UserSharedItem) HasShared() bool {
	if o != nil && !IsNil(o.Shared) {
		return true
	}

	return false
}

// SetShared gets a reference to the given string and assigns it to the Shared field.
func (o *UserSharedItem) SetShared(v string) {
	o.Shared = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserSharedItem) GetUser() UserInfo {
	if o == nil || IsNil(o.User) {
		var ret UserInfo
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSharedItem) GetUserOk() (*UserInfo, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserSharedItem) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserInfo and assigns it to the User field.
func (o *UserSharedItem) SetUser(v UserInfo) {
	o.User = &v
}

func (o UserSharedItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSharedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Shared) {
		toSerialize["shared"] = o.Shared
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserSharedItem struct {
	value *UserSharedItem
	isSet bool
}

func (v NullableUserSharedItem) Get() *UserSharedItem {
	return v.value
}

func (v *NullableUserSharedItem) Set(val *UserSharedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSharedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSharedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSharedItem(val *UserSharedItem) *NullableUserSharedItem {
	return &NullableUserSharedItem{value: val, isSet: true}
}

func (v NullableUserSharedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSharedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

