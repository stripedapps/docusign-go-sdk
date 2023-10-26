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

// checks if the ConnectOAuthConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectOAuthConfig{}

// ConnectOAuthConfig A complex object describing a Connect OAuth configuration.
type ConnectOAuthConfig struct {
	// The token URL for your authorization server or OAuth service.  This property is required.
	AuthorizationServerUrl *string `json:"authorizationServerUrl,omitempty"`
	// The client ID assigned to your app by your authorization server or OAuth service.  This property is required.
	ClientId *string `json:"clientId,omitempty"`
	// The secret value provided by your authorization server.  This property is required.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// 
	CustomParameters *string `json:"customParameters,omitempty"`
	// The scopes that your app will request from the authorization server.  This property is optional.  **Note:** If you are using Azure, this value is the application ID URI of the secified resource affixed with the `.default`. For example: `api://{{clientId}}/.default`
	Scope *string `json:"scope,omitempty"`
}

// NewConnectOAuthConfig instantiates a new ConnectOAuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectOAuthConfig() *ConnectOAuthConfig {
	this := ConnectOAuthConfig{}
	return &this
}

// NewConnectOAuthConfigWithDefaults instantiates a new ConnectOAuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectOAuthConfigWithDefaults() *ConnectOAuthConfig {
	this := ConnectOAuthConfig{}
	return &this
}

// GetAuthorizationServerUrl returns the AuthorizationServerUrl field value if set, zero value otherwise.
func (o *ConnectOAuthConfig) GetAuthorizationServerUrl() string {
	if o == nil || IsNil(o.AuthorizationServerUrl) {
		var ret string
		return ret
	}
	return *o.AuthorizationServerUrl
}

// GetAuthorizationServerUrlOk returns a tuple with the AuthorizationServerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectOAuthConfig) GetAuthorizationServerUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationServerUrl) {
		return nil, false
	}
	return o.AuthorizationServerUrl, true
}

// HasAuthorizationServerUrl returns a boolean if a field has been set.
func (o *ConnectOAuthConfig) HasAuthorizationServerUrl() bool {
	if o != nil && !IsNil(o.AuthorizationServerUrl) {
		return true
	}

	return false
}

// SetAuthorizationServerUrl gets a reference to the given string and assigns it to the AuthorizationServerUrl field.
func (o *ConnectOAuthConfig) SetAuthorizationServerUrl(v string) {
	o.AuthorizationServerUrl = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *ConnectOAuthConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectOAuthConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *ConnectOAuthConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *ConnectOAuthConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ConnectOAuthConfig) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectOAuthConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ConnectOAuthConfig) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ConnectOAuthConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetCustomParameters returns the CustomParameters field value if set, zero value otherwise.
func (o *ConnectOAuthConfig) GetCustomParameters() string {
	if o == nil || IsNil(o.CustomParameters) {
		var ret string
		return ret
	}
	return *o.CustomParameters
}

// GetCustomParametersOk returns a tuple with the CustomParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectOAuthConfig) GetCustomParametersOk() (*string, bool) {
	if o == nil || IsNil(o.CustomParameters) {
		return nil, false
	}
	return o.CustomParameters, true
}

// HasCustomParameters returns a boolean if a field has been set.
func (o *ConnectOAuthConfig) HasCustomParameters() bool {
	if o != nil && !IsNil(o.CustomParameters) {
		return true
	}

	return false
}

// SetCustomParameters gets a reference to the given string and assigns it to the CustomParameters field.
func (o *ConnectOAuthConfig) SetCustomParameters(v string) {
	o.CustomParameters = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ConnectOAuthConfig) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectOAuthConfig) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ConnectOAuthConfig) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ConnectOAuthConfig) SetScope(v string) {
	o.Scope = &v
}

func (o ConnectOAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectOAuthConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationServerUrl) {
		toSerialize["authorizationServerUrl"] = o.AuthorizationServerUrl
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.CustomParameters) {
		toSerialize["customParameters"] = o.CustomParameters
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	return toSerialize, nil
}

type NullableConnectOAuthConfig struct {
	value *ConnectOAuthConfig
	isSet bool
}

func (v NullableConnectOAuthConfig) Get() *ConnectOAuthConfig {
	return v.value
}

func (v *NullableConnectOAuthConfig) Set(val *ConnectOAuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectOAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectOAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectOAuthConfig(val *ConnectOAuthConfig) *NullableConnectOAuthConfig {
	return &NullableConnectOAuthConfig{value: val, isSet: true}
}

func (v NullableConnectOAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectOAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


