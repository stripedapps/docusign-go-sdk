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
	"time"
)

// checks if the UserDrilldownResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDrilldownResponse{}

// UserDrilldownResponse Information about a user.
type UserDrilldownResponse struct {
	// The user's unique ID.
	Id *string `json:"id,omitempty"`
	// The site ID of the organization.
	SiteId *int32 `json:"site_id,omitempty"`
	// The site name of the account.
	SiteName *string `json:"site_name,omitempty"`
	// The full name of the user.
	UserName *string `json:"user_name,omitempty"`
	// The user's first name.
	FirstName *string `json:"first_name,omitempty"`
	// The user's last name.
	LastName *string `json:"last_name,omitempty"`
	// The user's status. One of:  - `active` - `created` - `closed` 
	UserStatus *string `json:"user_status,omitempty"`
	// The ID of the user's default account.
	DefaultAccountId *string `json:"default_account_id,omitempty"`
	// The name of the user's default account.
	DefaultAccountName *string `json:"default_account_name,omitempty"`
	// The language and culture of the user.    * Chinese Simplified: `zh_CN`   * Chinese Traditional: `zh_TW`   * Dutch: `nl`   * English: `en`   * French: `fr`   * German: `de`   * Italian: `it`   * Japanese: `ja`   * Korean: `ko`   * Portuguese: `pt`   * Portuguese Brazil: `pt_BR`   * Russian: `ru`   * Spanish: `es` 
	LanguageCulture *string `json:"language_culture,omitempty"`
	// 
	SelectedLanguages *string `json:"selected_languages,omitempty"`
	// The user's federated status. One of:  - `RemoveStatus` - `FedAuthRequired` - `FedAuthBypass` - `Evicted` 
	FederatedStatus *string `json:"federated_status,omitempty"`
	// When **true,** the user has organization administration privileges.
	IsOrganizationAdmin *bool `json:"is_organization_admin,omitempty"`
	// The date the user's account was created.
	CreatedOn *time.Time `json:"created_on,omitempty"`
	// This property has been deprecated.
	LastLogin *time.Time `json:"last_login,omitempty"`
	// A list of organizations that have groups that the user is a member of.
	Memberships []MembershipResponse `json:"memberships,omitempty"`
	// A list of identities associated with the user.
	Identities []UserIdentityResponse `json:"identities,omitempty"`
	// 
	DeviceVerificationEnabled *bool `json:"device_verification_enabled,omitempty"`
}

// NewUserDrilldownResponse instantiates a new UserDrilldownResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDrilldownResponse() *UserDrilldownResponse {
	this := UserDrilldownResponse{}
	return &this
}

// NewUserDrilldownResponseWithDefaults instantiates a new UserDrilldownResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDrilldownResponseWithDefaults() *UserDrilldownResponse {
	this := UserDrilldownResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserDrilldownResponse) SetId(v string) {
	o.Id = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetSiteId() int32 {
	if o == nil || IsNil(o.SiteId) {
		var ret int32
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetSiteIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given int32 and assigns it to the SiteId field.
func (o *UserDrilldownResponse) SetSiteId(v int32) {
	o.SiteId = &v
}

// GetSiteName returns the SiteName field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetSiteName() string {
	if o == nil || IsNil(o.SiteName) {
		var ret string
		return ret
	}
	return *o.SiteName
}

// GetSiteNameOk returns a tuple with the SiteName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetSiteNameOk() (*string, bool) {
	if o == nil || IsNil(o.SiteName) {
		return nil, false
	}
	return o.SiteName, true
}

// HasSiteName returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasSiteName() bool {
	if o != nil && !IsNil(o.SiteName) {
		return true
	}

	return false
}

// SetSiteName gets a reference to the given string and assigns it to the SiteName field.
func (o *UserDrilldownResponse) SetSiteName(v string) {
	o.SiteName = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UserDrilldownResponse) SetUserName(v string) {
	o.UserName = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UserDrilldownResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UserDrilldownResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetUserStatus() string {
	if o == nil || IsNil(o.UserStatus) {
		var ret string
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetUserStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UserStatus) {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasUserStatus() bool {
	if o != nil && !IsNil(o.UserStatus) {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given string and assigns it to the UserStatus field.
func (o *UserDrilldownResponse) SetUserStatus(v string) {
	o.UserStatus = &v
}

// GetDefaultAccountId returns the DefaultAccountId field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetDefaultAccountId() string {
	if o == nil || IsNil(o.DefaultAccountId) {
		var ret string
		return ret
	}
	return *o.DefaultAccountId
}

// GetDefaultAccountIdOk returns a tuple with the DefaultAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetDefaultAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAccountId) {
		return nil, false
	}
	return o.DefaultAccountId, true
}

// HasDefaultAccountId returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasDefaultAccountId() bool {
	if o != nil && !IsNil(o.DefaultAccountId) {
		return true
	}

	return false
}

// SetDefaultAccountId gets a reference to the given string and assigns it to the DefaultAccountId field.
func (o *UserDrilldownResponse) SetDefaultAccountId(v string) {
	o.DefaultAccountId = &v
}

// GetDefaultAccountName returns the DefaultAccountName field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetDefaultAccountName() string {
	if o == nil || IsNil(o.DefaultAccountName) {
		var ret string
		return ret
	}
	return *o.DefaultAccountName
}

// GetDefaultAccountNameOk returns a tuple with the DefaultAccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetDefaultAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAccountName) {
		return nil, false
	}
	return o.DefaultAccountName, true
}

// HasDefaultAccountName returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasDefaultAccountName() bool {
	if o != nil && !IsNil(o.DefaultAccountName) {
		return true
	}

	return false
}

// SetDefaultAccountName gets a reference to the given string and assigns it to the DefaultAccountName field.
func (o *UserDrilldownResponse) SetDefaultAccountName(v string) {
	o.DefaultAccountName = &v
}

// GetLanguageCulture returns the LanguageCulture field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetLanguageCulture() string {
	if o == nil || IsNil(o.LanguageCulture) {
		var ret string
		return ret
	}
	return *o.LanguageCulture
}

// GetLanguageCultureOk returns a tuple with the LanguageCulture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetLanguageCultureOk() (*string, bool) {
	if o == nil || IsNil(o.LanguageCulture) {
		return nil, false
	}
	return o.LanguageCulture, true
}

// HasLanguageCulture returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasLanguageCulture() bool {
	if o != nil && !IsNil(o.LanguageCulture) {
		return true
	}

	return false
}

// SetLanguageCulture gets a reference to the given string and assigns it to the LanguageCulture field.
func (o *UserDrilldownResponse) SetLanguageCulture(v string) {
	o.LanguageCulture = &v
}

// GetSelectedLanguages returns the SelectedLanguages field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetSelectedLanguages() string {
	if o == nil || IsNil(o.SelectedLanguages) {
		var ret string
		return ret
	}
	return *o.SelectedLanguages
}

// GetSelectedLanguagesOk returns a tuple with the SelectedLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetSelectedLanguagesOk() (*string, bool) {
	if o == nil || IsNil(o.SelectedLanguages) {
		return nil, false
	}
	return o.SelectedLanguages, true
}

// HasSelectedLanguages returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasSelectedLanguages() bool {
	if o != nil && !IsNil(o.SelectedLanguages) {
		return true
	}

	return false
}

// SetSelectedLanguages gets a reference to the given string and assigns it to the SelectedLanguages field.
func (o *UserDrilldownResponse) SetSelectedLanguages(v string) {
	o.SelectedLanguages = &v
}

// GetFederatedStatus returns the FederatedStatus field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetFederatedStatus() string {
	if o == nil || IsNil(o.FederatedStatus) {
		var ret string
		return ret
	}
	return *o.FederatedStatus
}

// GetFederatedStatusOk returns a tuple with the FederatedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetFederatedStatusOk() (*string, bool) {
	if o == nil || IsNil(o.FederatedStatus) {
		return nil, false
	}
	return o.FederatedStatus, true
}

// HasFederatedStatus returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasFederatedStatus() bool {
	if o != nil && !IsNil(o.FederatedStatus) {
		return true
	}

	return false
}

// SetFederatedStatus gets a reference to the given string and assigns it to the FederatedStatus field.
func (o *UserDrilldownResponse) SetFederatedStatus(v string) {
	o.FederatedStatus = &v
}

// GetIsOrganizationAdmin returns the IsOrganizationAdmin field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetIsOrganizationAdmin() bool {
	if o == nil || IsNil(o.IsOrganizationAdmin) {
		var ret bool
		return ret
	}
	return *o.IsOrganizationAdmin
}

// GetIsOrganizationAdminOk returns a tuple with the IsOrganizationAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetIsOrganizationAdminOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOrganizationAdmin) {
		return nil, false
	}
	return o.IsOrganizationAdmin, true
}

// HasIsOrganizationAdmin returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasIsOrganizationAdmin() bool {
	if o != nil && !IsNil(o.IsOrganizationAdmin) {
		return true
	}

	return false
}

// SetIsOrganizationAdmin gets a reference to the given bool and assigns it to the IsOrganizationAdmin field.
func (o *UserDrilldownResponse) SetIsOrganizationAdmin(v bool) {
	o.IsOrganizationAdmin = &v
}

// GetCreatedOn returns the CreatedOn field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetCreatedOn() time.Time {
	if o == nil || IsNil(o.CreatedOn) {
		var ret time.Time
		return ret
	}
	return *o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedOn) {
		return nil, false
	}
	return o.CreatedOn, true
}

// HasCreatedOn returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasCreatedOn() bool {
	if o != nil && !IsNil(o.CreatedOn) {
		return true
	}

	return false
}

// SetCreatedOn gets a reference to the given time.Time and assigns it to the CreatedOn field.
func (o *UserDrilldownResponse) SetCreatedOn(v time.Time) {
	o.CreatedOn = &v
}

// GetLastLogin returns the LastLogin field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetLastLogin() time.Time {
	if o == nil || IsNil(o.LastLogin) {
		var ret time.Time
		return ret
	}
	return *o.LastLogin
}

// GetLastLoginOk returns a tuple with the LastLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetLastLoginOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastLogin) {
		return nil, false
	}
	return o.LastLogin, true
}

// HasLastLogin returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasLastLogin() bool {
	if o != nil && !IsNil(o.LastLogin) {
		return true
	}

	return false
}

// SetLastLogin gets a reference to the given time.Time and assigns it to the LastLogin field.
func (o *UserDrilldownResponse) SetLastLogin(v time.Time) {
	o.LastLogin = &v
}

// GetMemberships returns the Memberships field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetMemberships() []MembershipResponse {
	if o == nil || IsNil(o.Memberships) {
		var ret []MembershipResponse
		return ret
	}
	return o.Memberships
}

// GetMembershipsOk returns a tuple with the Memberships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetMembershipsOk() ([]MembershipResponse, bool) {
	if o == nil || IsNil(o.Memberships) {
		return nil, false
	}
	return o.Memberships, true
}

// HasMemberships returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasMemberships() bool {
	if o != nil && !IsNil(o.Memberships) {
		return true
	}

	return false
}

// SetMemberships gets a reference to the given []MembershipResponse and assigns it to the Memberships field.
func (o *UserDrilldownResponse) SetMemberships(v []MembershipResponse) {
	o.Memberships = v
}

// GetIdentities returns the Identities field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetIdentities() []UserIdentityResponse {
	if o == nil || IsNil(o.Identities) {
		var ret []UserIdentityResponse
		return ret
	}
	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetIdentitiesOk() ([]UserIdentityResponse, bool) {
	if o == nil || IsNil(o.Identities) {
		return nil, false
	}
	return o.Identities, true
}

// HasIdentities returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasIdentities() bool {
	if o != nil && !IsNil(o.Identities) {
		return true
	}

	return false
}

// SetIdentities gets a reference to the given []UserIdentityResponse and assigns it to the Identities field.
func (o *UserDrilldownResponse) SetIdentities(v []UserIdentityResponse) {
	o.Identities = v
}

// GetDeviceVerificationEnabled returns the DeviceVerificationEnabled field value if set, zero value otherwise.
func (o *UserDrilldownResponse) GetDeviceVerificationEnabled() bool {
	if o == nil || IsNil(o.DeviceVerificationEnabled) {
		var ret bool
		return ret
	}
	return *o.DeviceVerificationEnabled
}

// GetDeviceVerificationEnabledOk returns a tuple with the DeviceVerificationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDrilldownResponse) GetDeviceVerificationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DeviceVerificationEnabled) {
		return nil, false
	}
	return o.DeviceVerificationEnabled, true
}

// HasDeviceVerificationEnabled returns a boolean if a field has been set.
func (o *UserDrilldownResponse) HasDeviceVerificationEnabled() bool {
	if o != nil && !IsNil(o.DeviceVerificationEnabled) {
		return true
	}

	return false
}

// SetDeviceVerificationEnabled gets a reference to the given bool and assigns it to the DeviceVerificationEnabled field.
func (o *UserDrilldownResponse) SetDeviceVerificationEnabled(v bool) {
	o.DeviceVerificationEnabled = &v
}

func (o UserDrilldownResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDrilldownResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.SiteName) {
		toSerialize["site_name"] = o.SiteName
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.UserStatus) {
		toSerialize["user_status"] = o.UserStatus
	}
	if !IsNil(o.DefaultAccountId) {
		toSerialize["default_account_id"] = o.DefaultAccountId
	}
	if !IsNil(o.DefaultAccountName) {
		toSerialize["default_account_name"] = o.DefaultAccountName
	}
	if !IsNil(o.LanguageCulture) {
		toSerialize["language_culture"] = o.LanguageCulture
	}
	if !IsNil(o.SelectedLanguages) {
		toSerialize["selected_languages"] = o.SelectedLanguages
	}
	if !IsNil(o.FederatedStatus) {
		toSerialize["federated_status"] = o.FederatedStatus
	}
	if !IsNil(o.IsOrganizationAdmin) {
		toSerialize["is_organization_admin"] = o.IsOrganizationAdmin
	}
	if !IsNil(o.CreatedOn) {
		toSerialize["created_on"] = o.CreatedOn
	}
	if !IsNil(o.LastLogin) {
		toSerialize["last_login"] = o.LastLogin
	}
	if !IsNil(o.Memberships) {
		toSerialize["memberships"] = o.Memberships
	}
	if !IsNil(o.Identities) {
		toSerialize["identities"] = o.Identities
	}
	if !IsNil(o.DeviceVerificationEnabled) {
		toSerialize["device_verification_enabled"] = o.DeviceVerificationEnabled
	}
	return toSerialize, nil
}

type NullableUserDrilldownResponse struct {
	value *UserDrilldownResponse
	isSet bool
}

func (v NullableUserDrilldownResponse) Get() *UserDrilldownResponse {
	return v.value
}

func (v *NullableUserDrilldownResponse) Set(val *UserDrilldownResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDrilldownResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDrilldownResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDrilldownResponse(val *UserDrilldownResponse) *NullableUserDrilldownResponse {
	return &NullableUserDrilldownResponse{value: val, isSet: true}
}

func (v NullableUserDrilldownResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDrilldownResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

