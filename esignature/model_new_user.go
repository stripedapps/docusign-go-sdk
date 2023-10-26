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

// checks if the NewUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewUser{}

// NewUser Object representing a new user.
type NewUser struct {
	// Contains a token that can be used for authentication in API calls instead of using the user name and password.
	ApiPassword *string `json:"apiPassword,omitempty"`
	// The UTC DateTime when the item was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	// The user's email address.
	Email *string `json:"email,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The user's membership ID.
	MembershipId *string `json:"membershipId,omitempty"`
	// The ID of the permission profile.  Use [AccountPermissionProfiles: list](/docs/esign-rest-api/reference/accounts/accountpermissionprofiles/list/) to get a list of permission profiles and their IDs.  You can also download a CSV file of all permission profiles and their IDs from the **Settings > Permission Profiles** page of your eSignature account page. 
	PermissionProfileId *string `json:"permissionProfileId,omitempty"`
	// The name of the account permission profile.   Example: `Account Administrator`
	PermissionProfileName *string `json:"permissionProfileName,omitempty"`
	// A URI containing the user ID.
	Uri *string `json:"uri,omitempty"`
	// Specifies the user ID for the new user.
	UserId *string `json:"userId,omitempty"`
	// The name of the user.
	UserName *string `json:"userName,omitempty"`
	// Status of the user's account. One of:  - `ActivationRequired` - `ActivationSent` - `Active` - `Closed` - `Disabled` 
	UserStatus *string `json:"userStatus,omitempty"`
}

// NewNewUser instantiates a new NewUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewUser() *NewUser {
	this := NewUser{}
	return &this
}

// NewNewUserWithDefaults instantiates a new NewUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewUserWithDefaults() *NewUser {
	this := NewUser{}
	return &this
}

// GetApiPassword returns the ApiPassword field value if set, zero value otherwise.
func (o *NewUser) GetApiPassword() string {
	if o == nil || IsNil(o.ApiPassword) {
		var ret string
		return ret
	}
	return *o.ApiPassword
}

// GetApiPasswordOk returns a tuple with the ApiPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetApiPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ApiPassword) {
		return nil, false
	}
	return o.ApiPassword, true
}

// HasApiPassword returns a boolean if a field has been set.
func (o *NewUser) HasApiPassword() bool {
	if o != nil && !IsNil(o.ApiPassword) {
		return true
	}

	return false
}

// SetApiPassword gets a reference to the given string and assigns it to the ApiPassword field.
func (o *NewUser) SetApiPassword(v string) {
	o.ApiPassword = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *NewUser) GetCreatedDateTime() string {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *NewUser) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *NewUser) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *NewUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *NewUser) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *NewUser) SetEmail(v string) {
	o.Email = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *NewUser) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *NewUser) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *NewUser) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetMembershipId returns the MembershipId field value if set, zero value otherwise.
func (o *NewUser) GetMembershipId() string {
	if o == nil || IsNil(o.MembershipId) {
		var ret string
		return ret
	}
	return *o.MembershipId
}

// GetMembershipIdOk returns a tuple with the MembershipId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetMembershipIdOk() (*string, bool) {
	if o == nil || IsNil(o.MembershipId) {
		return nil, false
	}
	return o.MembershipId, true
}

// HasMembershipId returns a boolean if a field has been set.
func (o *NewUser) HasMembershipId() bool {
	if o != nil && !IsNil(o.MembershipId) {
		return true
	}

	return false
}

// SetMembershipId gets a reference to the given string and assigns it to the MembershipId field.
func (o *NewUser) SetMembershipId(v string) {
	o.MembershipId = &v
}

// GetPermissionProfileId returns the PermissionProfileId field value if set, zero value otherwise.
func (o *NewUser) GetPermissionProfileId() string {
	if o == nil || IsNil(o.PermissionProfileId) {
		var ret string
		return ret
	}
	return *o.PermissionProfileId
}

// GetPermissionProfileIdOk returns a tuple with the PermissionProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetPermissionProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionProfileId) {
		return nil, false
	}
	return o.PermissionProfileId, true
}

// HasPermissionProfileId returns a boolean if a field has been set.
func (o *NewUser) HasPermissionProfileId() bool {
	if o != nil && !IsNil(o.PermissionProfileId) {
		return true
	}

	return false
}

// SetPermissionProfileId gets a reference to the given string and assigns it to the PermissionProfileId field.
func (o *NewUser) SetPermissionProfileId(v string) {
	o.PermissionProfileId = &v
}

// GetPermissionProfileName returns the PermissionProfileName field value if set, zero value otherwise.
func (o *NewUser) GetPermissionProfileName() string {
	if o == nil || IsNil(o.PermissionProfileName) {
		var ret string
		return ret
	}
	return *o.PermissionProfileName
}

// GetPermissionProfileNameOk returns a tuple with the PermissionProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetPermissionProfileNameOk() (*string, bool) {
	if o == nil || IsNil(o.PermissionProfileName) {
		return nil, false
	}
	return o.PermissionProfileName, true
}

// HasPermissionProfileName returns a boolean if a field has been set.
func (o *NewUser) HasPermissionProfileName() bool {
	if o != nil && !IsNil(o.PermissionProfileName) {
		return true
	}

	return false
}

// SetPermissionProfileName gets a reference to the given string and assigns it to the PermissionProfileName field.
func (o *NewUser) SetPermissionProfileName(v string) {
	o.PermissionProfileName = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *NewUser) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *NewUser) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *NewUser) SetUri(v string) {
	o.Uri = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *NewUser) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *NewUser) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *NewUser) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *NewUser) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *NewUser) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *NewUser) SetUserName(v string) {
	o.UserName = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *NewUser) GetUserStatus() string {
	if o == nil || IsNil(o.UserStatus) {
		var ret string
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewUser) GetUserStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UserStatus) {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *NewUser) HasUserStatus() bool {
	if o != nil && !IsNil(o.UserStatus) {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given string and assigns it to the UserStatus field.
func (o *NewUser) SetUserStatus(v string) {
	o.UserStatus = &v
}

func (o NewUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiPassword) {
		toSerialize["apiPassword"] = o.ApiPassword
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.MembershipId) {
		toSerialize["membershipId"] = o.MembershipId
	}
	if !IsNil(o.PermissionProfileId) {
		toSerialize["permissionProfileId"] = o.PermissionProfileId
	}
	if !IsNil(o.PermissionProfileName) {
		toSerialize["permissionProfileName"] = o.PermissionProfileName
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.UserStatus) {
		toSerialize["userStatus"] = o.UserStatus
	}
	return toSerialize, nil
}

type NullableNewUser struct {
	value *NewUser
	isSet bool
}

func (v NullableNewUser) Get() *NewUser {
	return v.value
}

func (v *NullableNewUser) Set(val *NewUser) {
	v.value = val
	v.isSet = true
}

func (v NullableNewUser) IsSet() bool {
	return v.isSet
}

func (v *NullableNewUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewUser(val *NewUser) *NullableNewUser {
	return &NullableNewUser{value: val, isSet: true}
}

func (v NullableNewUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


