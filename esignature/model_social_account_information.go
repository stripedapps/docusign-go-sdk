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

// checks if the SocialAccountInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SocialAccountInformation{}

// SocialAccountInformation 
type SocialAccountInformation struct {
	// The users email address.
	Email *string `json:"email,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The social account provider (Facebook, Yahoo, etc.)
	Provider *string `json:"provider,omitempty"`
	// The ID provided by the Socal Account.
	SocialId *string `json:"socialId,omitempty"`
	// The full user name for the account.
	UserName *string `json:"userName,omitempty"`
}

// NewSocialAccountInformation instantiates a new SocialAccountInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSocialAccountInformation() *SocialAccountInformation {
	this := SocialAccountInformation{}
	return &this
}

// NewSocialAccountInformationWithDefaults instantiates a new SocialAccountInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSocialAccountInformationWithDefaults() *SocialAccountInformation {
	this := SocialAccountInformation{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SocialAccountInformation) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAccountInformation) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SocialAccountInformation) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SocialAccountInformation) SetEmail(v string) {
	o.Email = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *SocialAccountInformation) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAccountInformation) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *SocialAccountInformation) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *SocialAccountInformation) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *SocialAccountInformation) GetProvider() string {
	if o == nil || IsNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAccountInformation) GetProviderOk() (*string, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *SocialAccountInformation) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *SocialAccountInformation) SetProvider(v string) {
	o.Provider = &v
}

// GetSocialId returns the SocialId field value if set, zero value otherwise.
func (o *SocialAccountInformation) GetSocialId() string {
	if o == nil || IsNil(o.SocialId) {
		var ret string
		return ret
	}
	return *o.SocialId
}

// GetSocialIdOk returns a tuple with the SocialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAccountInformation) GetSocialIdOk() (*string, bool) {
	if o == nil || IsNil(o.SocialId) {
		return nil, false
	}
	return o.SocialId, true
}

// HasSocialId returns a boolean if a field has been set.
func (o *SocialAccountInformation) HasSocialId() bool {
	if o != nil && !IsNil(o.SocialId) {
		return true
	}

	return false
}

// SetSocialId gets a reference to the given string and assigns it to the SocialId field.
func (o *SocialAccountInformation) SetSocialId(v string) {
	o.SocialId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *SocialAccountInformation) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SocialAccountInformation) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *SocialAccountInformation) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *SocialAccountInformation) SetUserName(v string) {
	o.UserName = &v
}

func (o SocialAccountInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SocialAccountInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.SocialId) {
		toSerialize["socialId"] = o.SocialId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

type NullableSocialAccountInformation struct {
	value *SocialAccountInformation
	isSet bool
}

func (v NullableSocialAccountInformation) Get() *SocialAccountInformation {
	return v.value
}

func (v *NullableSocialAccountInformation) Set(val *SocialAccountInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableSocialAccountInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableSocialAccountInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSocialAccountInformation(val *SocialAccountInformation) *NullableSocialAccountInformation {
	return &NullableSocialAccountInformation{value: val, isSet: true}
}

func (v NullableSocialAccountInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSocialAccountInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

