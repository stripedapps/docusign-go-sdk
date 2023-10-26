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

// checks if the UserProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserProfile{}

// UserProfile 
type UserProfile struct {
	Address *AddressInformation `json:"address,omitempty"`
	// Indicates the authentication methods that the user uses. These properties cannot be modified by the PUT operation. 
	AuthenticationMethods []AuthenticationMethod `json:"authenticationMethods,omitempty"`
	// The name of the user's company.
	CompanyName *string `json:"companyName,omitempty"`
	//  When **true,** the user's company and title information display on the ID card. 
	DisplayOrganizationInfo *string `json:"displayOrganizationInfo,omitempty"`
	// When **true,** the user's address and phone number display on the ID card.
	DisplayPersonalInfo *string `json:"displayPersonalInfo,omitempty"`
	// When **true,** the user's ID card can be viewed from signed documents and envelope history.
	DisplayProfile *string `json:"displayProfile,omitempty"`
	// When **true,** the user's usage information displays on the ID card.
	DisplayUsageHistory *string `json:"displayUsageHistory,omitempty"`
	// The URL for retrieving the user's profile image.
	ProfileImageUri *string `json:"profileImageUri,omitempty"`
	// The title of the user.
	Title *string `json:"title,omitempty"`
	UsageHistory *UsageHistory `json:"usageHistory,omitempty"`
	UserDetails *UserInformation `json:"userDetails,omitempty"`
	// The date and time that the user's profile was last modified.
	UserProfileLastModifiedDate *string `json:"userProfileLastModifiedDate,omitempty"`
}

// NewUserProfile instantiates a new UserProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserProfile() *UserProfile {
	this := UserProfile{}
	return &this
}

// NewUserProfileWithDefaults instantiates a new UserProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserProfileWithDefaults() *UserProfile {
	this := UserProfile{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UserProfile) GetAddress() AddressInformation {
	if o == nil || IsNil(o.Address) {
		var ret AddressInformation
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetAddressOk() (*AddressInformation, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UserProfile) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given AddressInformation and assigns it to the Address field.
func (o *UserProfile) SetAddress(v AddressInformation) {
	o.Address = &v
}

// GetAuthenticationMethods returns the AuthenticationMethods field value if set, zero value otherwise.
func (o *UserProfile) GetAuthenticationMethods() []AuthenticationMethod {
	if o == nil || IsNil(o.AuthenticationMethods) {
		var ret []AuthenticationMethod
		return ret
	}
	return o.AuthenticationMethods
}

// GetAuthenticationMethodsOk returns a tuple with the AuthenticationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetAuthenticationMethodsOk() ([]AuthenticationMethod, bool) {
	if o == nil || IsNil(o.AuthenticationMethods) {
		return nil, false
	}
	return o.AuthenticationMethods, true
}

// HasAuthenticationMethods returns a boolean if a field has been set.
func (o *UserProfile) HasAuthenticationMethods() bool {
	if o != nil && !IsNil(o.AuthenticationMethods) {
		return true
	}

	return false
}

// SetAuthenticationMethods gets a reference to the given []AuthenticationMethod and assigns it to the AuthenticationMethods field.
func (o *UserProfile) SetAuthenticationMethods(v []AuthenticationMethod) {
	o.AuthenticationMethods = v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *UserProfile) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *UserProfile) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *UserProfile) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetDisplayOrganizationInfo returns the DisplayOrganizationInfo field value if set, zero value otherwise.
func (o *UserProfile) GetDisplayOrganizationInfo() string {
	if o == nil || IsNil(o.DisplayOrganizationInfo) {
		var ret string
		return ret
	}
	return *o.DisplayOrganizationInfo
}

// GetDisplayOrganizationInfoOk returns a tuple with the DisplayOrganizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDisplayOrganizationInfoOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayOrganizationInfo) {
		return nil, false
	}
	return o.DisplayOrganizationInfo, true
}

// HasDisplayOrganizationInfo returns a boolean if a field has been set.
func (o *UserProfile) HasDisplayOrganizationInfo() bool {
	if o != nil && !IsNil(o.DisplayOrganizationInfo) {
		return true
	}

	return false
}

// SetDisplayOrganizationInfo gets a reference to the given string and assigns it to the DisplayOrganizationInfo field.
func (o *UserProfile) SetDisplayOrganizationInfo(v string) {
	o.DisplayOrganizationInfo = &v
}

// GetDisplayPersonalInfo returns the DisplayPersonalInfo field value if set, zero value otherwise.
func (o *UserProfile) GetDisplayPersonalInfo() string {
	if o == nil || IsNil(o.DisplayPersonalInfo) {
		var ret string
		return ret
	}
	return *o.DisplayPersonalInfo
}

// GetDisplayPersonalInfoOk returns a tuple with the DisplayPersonalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDisplayPersonalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayPersonalInfo) {
		return nil, false
	}
	return o.DisplayPersonalInfo, true
}

// HasDisplayPersonalInfo returns a boolean if a field has been set.
func (o *UserProfile) HasDisplayPersonalInfo() bool {
	if o != nil && !IsNil(o.DisplayPersonalInfo) {
		return true
	}

	return false
}

// SetDisplayPersonalInfo gets a reference to the given string and assigns it to the DisplayPersonalInfo field.
func (o *UserProfile) SetDisplayPersonalInfo(v string) {
	o.DisplayPersonalInfo = &v
}

// GetDisplayProfile returns the DisplayProfile field value if set, zero value otherwise.
func (o *UserProfile) GetDisplayProfile() string {
	if o == nil || IsNil(o.DisplayProfile) {
		var ret string
		return ret
	}
	return *o.DisplayProfile
}

// GetDisplayProfileOk returns a tuple with the DisplayProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDisplayProfileOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayProfile) {
		return nil, false
	}
	return o.DisplayProfile, true
}

// HasDisplayProfile returns a boolean if a field has been set.
func (o *UserProfile) HasDisplayProfile() bool {
	if o != nil && !IsNil(o.DisplayProfile) {
		return true
	}

	return false
}

// SetDisplayProfile gets a reference to the given string and assigns it to the DisplayProfile field.
func (o *UserProfile) SetDisplayProfile(v string) {
	o.DisplayProfile = &v
}

// GetDisplayUsageHistory returns the DisplayUsageHistory field value if set, zero value otherwise.
func (o *UserProfile) GetDisplayUsageHistory() string {
	if o == nil || IsNil(o.DisplayUsageHistory) {
		var ret string
		return ret
	}
	return *o.DisplayUsageHistory
}

// GetDisplayUsageHistoryOk returns a tuple with the DisplayUsageHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetDisplayUsageHistoryOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUsageHistory) {
		return nil, false
	}
	return o.DisplayUsageHistory, true
}

// HasDisplayUsageHistory returns a boolean if a field has been set.
func (o *UserProfile) HasDisplayUsageHistory() bool {
	if o != nil && !IsNil(o.DisplayUsageHistory) {
		return true
	}

	return false
}

// SetDisplayUsageHistory gets a reference to the given string and assigns it to the DisplayUsageHistory field.
func (o *UserProfile) SetDisplayUsageHistory(v string) {
	o.DisplayUsageHistory = &v
}

// GetProfileImageUri returns the ProfileImageUri field value if set, zero value otherwise.
func (o *UserProfile) GetProfileImageUri() string {
	if o == nil || IsNil(o.ProfileImageUri) {
		var ret string
		return ret
	}
	return *o.ProfileImageUri
}

// GetProfileImageUriOk returns a tuple with the ProfileImageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetProfileImageUriOk() (*string, bool) {
	if o == nil || IsNil(o.ProfileImageUri) {
		return nil, false
	}
	return o.ProfileImageUri, true
}

// HasProfileImageUri returns a boolean if a field has been set.
func (o *UserProfile) HasProfileImageUri() bool {
	if o != nil && !IsNil(o.ProfileImageUri) {
		return true
	}

	return false
}

// SetProfileImageUri gets a reference to the given string and assigns it to the ProfileImageUri field.
func (o *UserProfile) SetProfileImageUri(v string) {
	o.ProfileImageUri = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserProfile) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserProfile) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserProfile) SetTitle(v string) {
	o.Title = &v
}

// GetUsageHistory returns the UsageHistory field value if set, zero value otherwise.
func (o *UserProfile) GetUsageHistory() UsageHistory {
	if o == nil || IsNil(o.UsageHistory) {
		var ret UsageHistory
		return ret
	}
	return *o.UsageHistory
}

// GetUsageHistoryOk returns a tuple with the UsageHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUsageHistoryOk() (*UsageHistory, bool) {
	if o == nil || IsNil(o.UsageHistory) {
		return nil, false
	}
	return o.UsageHistory, true
}

// HasUsageHistory returns a boolean if a field has been set.
func (o *UserProfile) HasUsageHistory() bool {
	if o != nil && !IsNil(o.UsageHistory) {
		return true
	}

	return false
}

// SetUsageHistory gets a reference to the given UsageHistory and assigns it to the UsageHistory field.
func (o *UserProfile) SetUsageHistory(v UsageHistory) {
	o.UsageHistory = &v
}

// GetUserDetails returns the UserDetails field value if set, zero value otherwise.
func (o *UserProfile) GetUserDetails() UserInformation {
	if o == nil || IsNil(o.UserDetails) {
		var ret UserInformation
		return ret
	}
	return *o.UserDetails
}

// GetUserDetailsOk returns a tuple with the UserDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserDetailsOk() (*UserInformation, bool) {
	if o == nil || IsNil(o.UserDetails) {
		return nil, false
	}
	return o.UserDetails, true
}

// HasUserDetails returns a boolean if a field has been set.
func (o *UserProfile) HasUserDetails() bool {
	if o != nil && !IsNil(o.UserDetails) {
		return true
	}

	return false
}

// SetUserDetails gets a reference to the given UserInformation and assigns it to the UserDetails field.
func (o *UserProfile) SetUserDetails(v UserInformation) {
	o.UserDetails = &v
}

// GetUserProfileLastModifiedDate returns the UserProfileLastModifiedDate field value if set, zero value otherwise.
func (o *UserProfile) GetUserProfileLastModifiedDate() string {
	if o == nil || IsNil(o.UserProfileLastModifiedDate) {
		var ret string
		return ret
	}
	return *o.UserProfileLastModifiedDate
}

// GetUserProfileLastModifiedDateOk returns a tuple with the UserProfileLastModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserProfile) GetUserProfileLastModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.UserProfileLastModifiedDate) {
		return nil, false
	}
	return o.UserProfileLastModifiedDate, true
}

// HasUserProfileLastModifiedDate returns a boolean if a field has been set.
func (o *UserProfile) HasUserProfileLastModifiedDate() bool {
	if o != nil && !IsNil(o.UserProfileLastModifiedDate) {
		return true
	}

	return false
}

// SetUserProfileLastModifiedDate gets a reference to the given string and assigns it to the UserProfileLastModifiedDate field.
func (o *UserProfile) SetUserProfileLastModifiedDate(v string) {
	o.UserProfileLastModifiedDate = &v
}

func (o UserProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.AuthenticationMethods) {
		toSerialize["authenticationMethods"] = o.AuthenticationMethods
	}
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.DisplayOrganizationInfo) {
		toSerialize["displayOrganizationInfo"] = o.DisplayOrganizationInfo
	}
	if !IsNil(o.DisplayPersonalInfo) {
		toSerialize["displayPersonalInfo"] = o.DisplayPersonalInfo
	}
	if !IsNil(o.DisplayProfile) {
		toSerialize["displayProfile"] = o.DisplayProfile
	}
	if !IsNil(o.DisplayUsageHistory) {
		toSerialize["displayUsageHistory"] = o.DisplayUsageHistory
	}
	if !IsNil(o.ProfileImageUri) {
		toSerialize["profileImageUri"] = o.ProfileImageUri
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UsageHistory) {
		toSerialize["usageHistory"] = o.UsageHistory
	}
	if !IsNil(o.UserDetails) {
		toSerialize["userDetails"] = o.UserDetails
	}
	if !IsNil(o.UserProfileLastModifiedDate) {
		toSerialize["userProfileLastModifiedDate"] = o.UserProfileLastModifiedDate
	}
	return toSerialize, nil
}

type NullableUserProfile struct {
	value *UserProfile
	isSet bool
}

func (v NullableUserProfile) Get() *UserProfile {
	return v.value
}

func (v *NullableUserProfile) Set(val *UserProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableUserProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableUserProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserProfile(val *UserProfile) *NullableUserProfile {
	return &NullableUserProfile{value: val, isSet: true}
}

func (v NullableUserProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


