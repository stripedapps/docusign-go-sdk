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

// checks if the AuthenticationMethod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationMethod{}

// AuthenticationMethod Contains information about the method used for authentication.
type AuthenticationMethod struct {
	// Indicates the type of authentication. Valid values are: PhoneAuth, STAN, ISCheck, OFAC, AccessCode, AgeVerify, or SSOAuth. 
	AuthenticationType *string `json:"authenticationType,omitempty"`
	// The last provider that authenticated the user. 
	LastProvider *string `json:"lastProvider,omitempty"`
	//  The data and time the user last used the authentication method. 
	LastTimestamp *string `json:"lastTimestamp,omitempty"`
	// The number of times the authentication method was used. 
	TotalCount *string `json:"totalCount,omitempty"`
}

// NewAuthenticationMethod instantiates a new AuthenticationMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationMethod() *AuthenticationMethod {
	this := AuthenticationMethod{}
	return &this
}

// NewAuthenticationMethodWithDefaults instantiates a new AuthenticationMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationMethodWithDefaults() *AuthenticationMethod {
	this := AuthenticationMethod{}
	return &this
}

// GetAuthenticationType returns the AuthenticationType field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetAuthenticationType() string {
	if o == nil || IsNil(o.AuthenticationType) {
		var ret string
		return ret
	}
	return *o.AuthenticationType
}

// GetAuthenticationTypeOk returns a tuple with the AuthenticationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetAuthenticationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationType) {
		return nil, false
	}
	return o.AuthenticationType, true
}

// HasAuthenticationType returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasAuthenticationType() bool {
	if o != nil && !IsNil(o.AuthenticationType) {
		return true
	}

	return false
}

// SetAuthenticationType gets a reference to the given string and assigns it to the AuthenticationType field.
func (o *AuthenticationMethod) SetAuthenticationType(v string) {
	o.AuthenticationType = &v
}

// GetLastProvider returns the LastProvider field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetLastProvider() string {
	if o == nil || IsNil(o.LastProvider) {
		var ret string
		return ret
	}
	return *o.LastProvider
}

// GetLastProviderOk returns a tuple with the LastProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetLastProviderOk() (*string, bool) {
	if o == nil || IsNil(o.LastProvider) {
		return nil, false
	}
	return o.LastProvider, true
}

// HasLastProvider returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasLastProvider() bool {
	if o != nil && !IsNil(o.LastProvider) {
		return true
	}

	return false
}

// SetLastProvider gets a reference to the given string and assigns it to the LastProvider field.
func (o *AuthenticationMethod) SetLastProvider(v string) {
	o.LastProvider = &v
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetLastTimestamp() string {
	if o == nil || IsNil(o.LastTimestamp) {
		var ret string
		return ret
	}
	return *o.LastTimestamp
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetLastTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.LastTimestamp) {
		return nil, false
	}
	return o.LastTimestamp, true
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasLastTimestamp() bool {
	if o != nil && !IsNil(o.LastTimestamp) {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given string and assigns it to the LastTimestamp field.
func (o *AuthenticationMethod) SetLastTimestamp(v string) {
	o.LastTimestamp = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *AuthenticationMethod) GetTotalCount() string {
	if o == nil || IsNil(o.TotalCount) {
		var ret string
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationMethod) GetTotalCountOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *AuthenticationMethod) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given string and assigns it to the TotalCount field.
func (o *AuthenticationMethod) SetTotalCount(v string) {
	o.TotalCount = &v
}

func (o AuthenticationMethod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationMethod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthenticationType) {
		toSerialize["authenticationType"] = o.AuthenticationType
	}
	if !IsNil(o.LastProvider) {
		toSerialize["lastProvider"] = o.LastProvider
	}
	if !IsNil(o.LastTimestamp) {
		toSerialize["lastTimestamp"] = o.LastTimestamp
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableAuthenticationMethod struct {
	value *AuthenticationMethod
	isSet bool
}

func (v NullableAuthenticationMethod) Get() *AuthenticationMethod {
	return v.value
}

func (v *NullableAuthenticationMethod) Set(val *AuthenticationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationMethod(val *AuthenticationMethod) *NullableAuthenticationMethod {
	return &NullableAuthenticationMethod{value: val, isSet: true}
}

func (v NullableAuthenticationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

