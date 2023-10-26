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

// checks if the TemplateLocks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateLocks{}

// TemplateLocks This section provides information about template locks. You use template locks to prevent others from making changes to a template while you are modifying it.
type TemplateLocks struct {
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The number of seconds until the lock expires when there is no activity on the template.  If no value is entered, then the default value of 300 seconds is used. The maximum value is 1,800 seconds.  The lock duration can be extended. 
	LockDurationInSeconds *string `json:"lockDurationInSeconds,omitempty"`
	// Specifies the friendly name of  the application that is locking the envelope.
	LockedByApp *string `json:"lockedByApp,omitempty"`
	LockedByUser *UserInfo `json:"lockedByUser,omitempty"`
	// The date and time that the lock expires.
	LockedUntilDateTime *string `json:"lockedUntilDateTime,omitempty"`
	// A unique identifier provided to the owner of the lock. You must use this token with subsequent calls to prove ownership of the lock.
	LockToken *string `json:"lockToken,omitempty"`
	// The type of lock.  Currently `edit` is the only supported type.
	LockType *string `json:"lockType,omitempty"`
	// When **true,** a scratchpad is used to edit information.  
	UseScratchPad *string `json:"useScratchPad,omitempty"`
}

// NewTemplateLocks instantiates a new TemplateLocks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateLocks() *TemplateLocks {
	this := TemplateLocks{}
	return &this
}

// NewTemplateLocksWithDefaults instantiates a new TemplateLocks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateLocksWithDefaults() *TemplateLocks {
	this := TemplateLocks{}
	return &this
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *TemplateLocks) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *TemplateLocks) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *TemplateLocks) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetLockDurationInSeconds returns the LockDurationInSeconds field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockDurationInSeconds() string {
	if o == nil || IsNil(o.LockDurationInSeconds) {
		var ret string
		return ret
	}
	return *o.LockDurationInSeconds
}

// GetLockDurationInSecondsOk returns a tuple with the LockDurationInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockDurationInSecondsOk() (*string, bool) {
	if o == nil || IsNil(o.LockDurationInSeconds) {
		return nil, false
	}
	return o.LockDurationInSeconds, true
}

// HasLockDurationInSeconds returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockDurationInSeconds() bool {
	if o != nil && !IsNil(o.LockDurationInSeconds) {
		return true
	}

	return false
}

// SetLockDurationInSeconds gets a reference to the given string and assigns it to the LockDurationInSeconds field.
func (o *TemplateLocks) SetLockDurationInSeconds(v string) {
	o.LockDurationInSeconds = &v
}

// GetLockedByApp returns the LockedByApp field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockedByApp() string {
	if o == nil || IsNil(o.LockedByApp) {
		var ret string
		return ret
	}
	return *o.LockedByApp
}

// GetLockedByAppOk returns a tuple with the LockedByApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockedByAppOk() (*string, bool) {
	if o == nil || IsNil(o.LockedByApp) {
		return nil, false
	}
	return o.LockedByApp, true
}

// HasLockedByApp returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockedByApp() bool {
	if o != nil && !IsNil(o.LockedByApp) {
		return true
	}

	return false
}

// SetLockedByApp gets a reference to the given string and assigns it to the LockedByApp field.
func (o *TemplateLocks) SetLockedByApp(v string) {
	o.LockedByApp = &v
}

// GetLockedByUser returns the LockedByUser field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockedByUser() UserInfo {
	if o == nil || IsNil(o.LockedByUser) {
		var ret UserInfo
		return ret
	}
	return *o.LockedByUser
}

// GetLockedByUserOk returns a tuple with the LockedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockedByUserOk() (*UserInfo, bool) {
	if o == nil || IsNil(o.LockedByUser) {
		return nil, false
	}
	return o.LockedByUser, true
}

// HasLockedByUser returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockedByUser() bool {
	if o != nil && !IsNil(o.LockedByUser) {
		return true
	}

	return false
}

// SetLockedByUser gets a reference to the given UserInfo and assigns it to the LockedByUser field.
func (o *TemplateLocks) SetLockedByUser(v UserInfo) {
	o.LockedByUser = &v
}

// GetLockedUntilDateTime returns the LockedUntilDateTime field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockedUntilDateTime() string {
	if o == nil || IsNil(o.LockedUntilDateTime) {
		var ret string
		return ret
	}
	return *o.LockedUntilDateTime
}

// GetLockedUntilDateTimeOk returns a tuple with the LockedUntilDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockedUntilDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LockedUntilDateTime) {
		return nil, false
	}
	return o.LockedUntilDateTime, true
}

// HasLockedUntilDateTime returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockedUntilDateTime() bool {
	if o != nil && !IsNil(o.LockedUntilDateTime) {
		return true
	}

	return false
}

// SetLockedUntilDateTime gets a reference to the given string and assigns it to the LockedUntilDateTime field.
func (o *TemplateLocks) SetLockedUntilDateTime(v string) {
	o.LockedUntilDateTime = &v
}

// GetLockToken returns the LockToken field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockToken() string {
	if o == nil || IsNil(o.LockToken) {
		var ret string
		return ret
	}
	return *o.LockToken
}

// GetLockTokenOk returns a tuple with the LockToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockTokenOk() (*string, bool) {
	if o == nil || IsNil(o.LockToken) {
		return nil, false
	}
	return o.LockToken, true
}

// HasLockToken returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockToken() bool {
	if o != nil && !IsNil(o.LockToken) {
		return true
	}

	return false
}

// SetLockToken gets a reference to the given string and assigns it to the LockToken field.
func (o *TemplateLocks) SetLockToken(v string) {
	o.LockToken = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *TemplateLocks) GetLockType() string {
	if o == nil || IsNil(o.LockType) {
		var ret string
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetLockTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LockType) {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *TemplateLocks) HasLockType() bool {
	if o != nil && !IsNil(o.LockType) {
		return true
	}

	return false
}

// SetLockType gets a reference to the given string and assigns it to the LockType field.
func (o *TemplateLocks) SetLockType(v string) {
	o.LockType = &v
}

// GetUseScratchPad returns the UseScratchPad field value if set, zero value otherwise.
func (o *TemplateLocks) GetUseScratchPad() string {
	if o == nil || IsNil(o.UseScratchPad) {
		var ret string
		return ret
	}
	return *o.UseScratchPad
}

// GetUseScratchPadOk returns a tuple with the UseScratchPad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateLocks) GetUseScratchPadOk() (*string, bool) {
	if o == nil || IsNil(o.UseScratchPad) {
		return nil, false
	}
	return o.UseScratchPad, true
}

// HasUseScratchPad returns a boolean if a field has been set.
func (o *TemplateLocks) HasUseScratchPad() bool {
	if o != nil && !IsNil(o.UseScratchPad) {
		return true
	}

	return false
}

// SetUseScratchPad gets a reference to the given string and assigns it to the UseScratchPad field.
func (o *TemplateLocks) SetUseScratchPad(v string) {
	o.UseScratchPad = &v
}

func (o TemplateLocks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateLocks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.LockDurationInSeconds) {
		toSerialize["lockDurationInSeconds"] = o.LockDurationInSeconds
	}
	if !IsNil(o.LockedByApp) {
		toSerialize["lockedByApp"] = o.LockedByApp
	}
	if !IsNil(o.LockedByUser) {
		toSerialize["lockedByUser"] = o.LockedByUser
	}
	if !IsNil(o.LockedUntilDateTime) {
		toSerialize["lockedUntilDateTime"] = o.LockedUntilDateTime
	}
	if !IsNil(o.LockToken) {
		toSerialize["lockToken"] = o.LockToken
	}
	if !IsNil(o.LockType) {
		toSerialize["lockType"] = o.LockType
	}
	if !IsNil(o.UseScratchPad) {
		toSerialize["useScratchPad"] = o.UseScratchPad
	}
	return toSerialize, nil
}

type NullableTemplateLocks struct {
	value *TemplateLocks
	isSet bool
}

func (v NullableTemplateLocks) Get() *TemplateLocks {
	return v.value
}

func (v *NullableTemplateLocks) Set(val *TemplateLocks) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateLocks) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateLocks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateLocks(val *TemplateLocks) *NullableTemplateLocks {
	return &NullableTemplateLocks{value: val, isSet: true}
}

func (v NullableTemplateLocks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateLocks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

