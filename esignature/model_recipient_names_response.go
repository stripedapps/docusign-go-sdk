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

// checks if the RecipientNamesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientNamesResponse{}

// RecipientNamesResponse This response object contains a list of recipients.
type RecipientNamesResponse struct {
	// When **true,** the email address is used by more than one user.
	MultipleUsers *string `json:"multipleUsers,omitempty"`
	// The names of the recipients associated with the email address.
	RecipientNames []string `json:"recipientNames,omitempty"`
	// When **true,** new names cannot be added to the email address.
	ReservedRecipientEmail *string `json:"reservedRecipientEmail,omitempty"`
}

// NewRecipientNamesResponse instantiates a new RecipientNamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientNamesResponse() *RecipientNamesResponse {
	this := RecipientNamesResponse{}
	return &this
}

// NewRecipientNamesResponseWithDefaults instantiates a new RecipientNamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientNamesResponseWithDefaults() *RecipientNamesResponse {
	this := RecipientNamesResponse{}
	return &this
}

// GetMultipleUsers returns the MultipleUsers field value if set, zero value otherwise.
func (o *RecipientNamesResponse) GetMultipleUsers() string {
	if o == nil || IsNil(o.MultipleUsers) {
		var ret string
		return ret
	}
	return *o.MultipleUsers
}

// GetMultipleUsersOk returns a tuple with the MultipleUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientNamesResponse) GetMultipleUsersOk() (*string, bool) {
	if o == nil || IsNil(o.MultipleUsers) {
		return nil, false
	}
	return o.MultipleUsers, true
}

// HasMultipleUsers returns a boolean if a field has been set.
func (o *RecipientNamesResponse) HasMultipleUsers() bool {
	if o != nil && !IsNil(o.MultipleUsers) {
		return true
	}

	return false
}

// SetMultipleUsers gets a reference to the given string and assigns it to the MultipleUsers field.
func (o *RecipientNamesResponse) SetMultipleUsers(v string) {
	o.MultipleUsers = &v
}

// GetRecipientNames returns the RecipientNames field value if set, zero value otherwise.
func (o *RecipientNamesResponse) GetRecipientNames() []string {
	if o == nil || IsNil(o.RecipientNames) {
		var ret []string
		return ret
	}
	return o.RecipientNames
}

// GetRecipientNamesOk returns a tuple with the RecipientNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientNamesResponse) GetRecipientNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.RecipientNames) {
		return nil, false
	}
	return o.RecipientNames, true
}

// HasRecipientNames returns a boolean if a field has been set.
func (o *RecipientNamesResponse) HasRecipientNames() bool {
	if o != nil && !IsNil(o.RecipientNames) {
		return true
	}

	return false
}

// SetRecipientNames gets a reference to the given []string and assigns it to the RecipientNames field.
func (o *RecipientNamesResponse) SetRecipientNames(v []string) {
	o.RecipientNames = v
}

// GetReservedRecipientEmail returns the ReservedRecipientEmail field value if set, zero value otherwise.
func (o *RecipientNamesResponse) GetReservedRecipientEmail() string {
	if o == nil || IsNil(o.ReservedRecipientEmail) {
		var ret string
		return ret
	}
	return *o.ReservedRecipientEmail
}

// GetReservedRecipientEmailOk returns a tuple with the ReservedRecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientNamesResponse) GetReservedRecipientEmailOk() (*string, bool) {
	if o == nil || IsNil(o.ReservedRecipientEmail) {
		return nil, false
	}
	return o.ReservedRecipientEmail, true
}

// HasReservedRecipientEmail returns a boolean if a field has been set.
func (o *RecipientNamesResponse) HasReservedRecipientEmail() bool {
	if o != nil && !IsNil(o.ReservedRecipientEmail) {
		return true
	}

	return false
}

// SetReservedRecipientEmail gets a reference to the given string and assigns it to the ReservedRecipientEmail field.
func (o *RecipientNamesResponse) SetReservedRecipientEmail(v string) {
	o.ReservedRecipientEmail = &v
}

func (o RecipientNamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientNamesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MultipleUsers) {
		toSerialize["multipleUsers"] = o.MultipleUsers
	}
	if !IsNil(o.RecipientNames) {
		toSerialize["recipientNames"] = o.RecipientNames
	}
	if !IsNil(o.ReservedRecipientEmail) {
		toSerialize["reservedRecipientEmail"] = o.ReservedRecipientEmail
	}
	return toSerialize, nil
}

type NullableRecipientNamesResponse struct {
	value *RecipientNamesResponse
	isSet bool
}

func (v NullableRecipientNamesResponse) Get() *RecipientNamesResponse {
	return v.value
}

func (v *NullableRecipientNamesResponse) Set(val *RecipientNamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientNamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientNamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientNamesResponse(val *RecipientNamesResponse) *NullableRecipientNamesResponse {
	return &NullableRecipientNamesResponse{value: val, isSet: true}
}

func (v NullableRecipientNamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientNamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

