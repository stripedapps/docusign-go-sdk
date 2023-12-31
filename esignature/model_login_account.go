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

// checks if the LoginAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoginAccount{}

// LoginAccount 
type LoginAccount struct {
	// The account ID associated with the envelope.
	AccountId *string `json:"accountId,omitempty"`
	// The GUID associated with the account ID.
	AccountIdGuid *string `json:"accountIdGuid,omitempty"`
	// The URL that should be used for successive calls to this account. It includes the protocal (https), the DocuSign server where the account is located, and the account number. Use this Url to make API calls against this account. Many of the API calls provide Uri's that are relative to this baseUrl.
	BaseUrl *string `json:"baseUrl,omitempty"`
	// The email address for the user.
	Email *string `json:"email,omitempty"`
	// This value is true if this is the default account for the user, otherwise false is returned.
	IsDefault *string `json:"isDefault,omitempty"`
	// A list of settings on the account that indicate what features are available.
	LoginAccountSettings []NameValue `json:"loginAccountSettings,omitempty"`
	// A list of user-level settings that indicate what user-specific features are available.
	LoginUserSettings []NameValue `json:"loginUserSettings,omitempty"`
	// The name associated with the account.
	Name *string `json:"name,omitempty"`
	// An optional descirption of the site that hosts the account.
	SiteDescription *string `json:"siteDescription,omitempty"`
	// The ID of the user to access.  **Note:** Users can only access their own information. A user, even one with Admin rights, cannot access another user's settings.
	UserId *string `json:"userId,omitempty"`
	// The name of this user as defined by the account.
	UserName *string `json:"userName,omitempty"`
}

// NewLoginAccount instantiates a new LoginAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoginAccount() *LoginAccount {
	this := LoginAccount{}
	return &this
}

// NewLoginAccountWithDefaults instantiates a new LoginAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoginAccountWithDefaults() *LoginAccount {
	this := LoginAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *LoginAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *LoginAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *LoginAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetAccountIdGuid returns the AccountIdGuid field value if set, zero value otherwise.
func (o *LoginAccount) GetAccountIdGuid() string {
	if o == nil || IsNil(o.AccountIdGuid) {
		var ret string
		return ret
	}
	return *o.AccountIdGuid
}

// GetAccountIdGuidOk returns a tuple with the AccountIdGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetAccountIdGuidOk() (*string, bool) {
	if o == nil || IsNil(o.AccountIdGuid) {
		return nil, false
	}
	return o.AccountIdGuid, true
}

// HasAccountIdGuid returns a boolean if a field has been set.
func (o *LoginAccount) HasAccountIdGuid() bool {
	if o != nil && !IsNil(o.AccountIdGuid) {
		return true
	}

	return false
}

// SetAccountIdGuid gets a reference to the given string and assigns it to the AccountIdGuid field.
func (o *LoginAccount) SetAccountIdGuid(v string) {
	o.AccountIdGuid = &v
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *LoginAccount) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl) {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BaseUrl) {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *LoginAccount) HasBaseUrl() bool {
	if o != nil && !IsNil(o.BaseUrl) {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *LoginAccount) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *LoginAccount) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *LoginAccount) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *LoginAccount) SetEmail(v string) {
	o.Email = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *LoginAccount) GetIsDefault() string {
	if o == nil || IsNil(o.IsDefault) {
		var ret string
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetIsDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *LoginAccount) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given string and assigns it to the IsDefault field.
func (o *LoginAccount) SetIsDefault(v string) {
	o.IsDefault = &v
}

// GetLoginAccountSettings returns the LoginAccountSettings field value if set, zero value otherwise.
func (o *LoginAccount) GetLoginAccountSettings() []NameValue {
	if o == nil || IsNil(o.LoginAccountSettings) {
		var ret []NameValue
		return ret
	}
	return o.LoginAccountSettings
}

// GetLoginAccountSettingsOk returns a tuple with the LoginAccountSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetLoginAccountSettingsOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.LoginAccountSettings) {
		return nil, false
	}
	return o.LoginAccountSettings, true
}

// HasLoginAccountSettings returns a boolean if a field has been set.
func (o *LoginAccount) HasLoginAccountSettings() bool {
	if o != nil && !IsNil(o.LoginAccountSettings) {
		return true
	}

	return false
}

// SetLoginAccountSettings gets a reference to the given []NameValue and assigns it to the LoginAccountSettings field.
func (o *LoginAccount) SetLoginAccountSettings(v []NameValue) {
	o.LoginAccountSettings = v
}

// GetLoginUserSettings returns the LoginUserSettings field value if set, zero value otherwise.
func (o *LoginAccount) GetLoginUserSettings() []NameValue {
	if o == nil || IsNil(o.LoginUserSettings) {
		var ret []NameValue
		return ret
	}
	return o.LoginUserSettings
}

// GetLoginUserSettingsOk returns a tuple with the LoginUserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetLoginUserSettingsOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.LoginUserSettings) {
		return nil, false
	}
	return o.LoginUserSettings, true
}

// HasLoginUserSettings returns a boolean if a field has been set.
func (o *LoginAccount) HasLoginUserSettings() bool {
	if o != nil && !IsNil(o.LoginUserSettings) {
		return true
	}

	return false
}

// SetLoginUserSettings gets a reference to the given []NameValue and assigns it to the LoginUserSettings field.
func (o *LoginAccount) SetLoginUserSettings(v []NameValue) {
	o.LoginUserSettings = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoginAccount) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoginAccount) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoginAccount) SetName(v string) {
	o.Name = &v
}

// GetSiteDescription returns the SiteDescription field value if set, zero value otherwise.
func (o *LoginAccount) GetSiteDescription() string {
	if o == nil || IsNil(o.SiteDescription) {
		var ret string
		return ret
	}
	return *o.SiteDescription
}

// GetSiteDescriptionOk returns a tuple with the SiteDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetSiteDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SiteDescription) {
		return nil, false
	}
	return o.SiteDescription, true
}

// HasSiteDescription returns a boolean if a field has been set.
func (o *LoginAccount) HasSiteDescription() bool {
	if o != nil && !IsNil(o.SiteDescription) {
		return true
	}

	return false
}

// SetSiteDescription gets a reference to the given string and assigns it to the SiteDescription field.
func (o *LoginAccount) SetSiteDescription(v string) {
	o.SiteDescription = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *LoginAccount) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *LoginAccount) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *LoginAccount) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *LoginAccount) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoginAccount) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *LoginAccount) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *LoginAccount) SetUserName(v string) {
	o.UserName = &v
}

func (o LoginAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoginAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.AccountIdGuid) {
		toSerialize["accountIdGuid"] = o.AccountIdGuid
	}
	if !IsNil(o.BaseUrl) {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.LoginAccountSettings) {
		toSerialize["loginAccountSettings"] = o.LoginAccountSettings
	}
	if !IsNil(o.LoginUserSettings) {
		toSerialize["loginUserSettings"] = o.LoginUserSettings
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.SiteDescription) {
		toSerialize["siteDescription"] = o.SiteDescription
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	return toSerialize, nil
}

type NullableLoginAccount struct {
	value *LoginAccount
	isSet bool
}

func (v NullableLoginAccount) Get() *LoginAccount {
	return v.value
}

func (v *NullableLoginAccount) Set(val *LoginAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableLoginAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableLoginAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoginAccount(val *LoginAccount) *NullableLoginAccount {
	return &NullableLoginAccount{value: val, isSet: true}
}

func (v NullableLoginAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoginAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


