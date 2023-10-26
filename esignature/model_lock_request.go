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

// checks if the LockRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LockRequest{}

// LockRequest This request object contains information about the lock that you want to create or update.
type LockRequest struct {
	// The number of seconds to lock the envelope for editing.  Must be greater than 0 seconds.
	LockDurationInSeconds *string `json:"lockDurationInSeconds,omitempty"`
	// A friendly name of the application used to lock the envelope.  Will be used in error messages to the user when lock conflicts occur.
	LockedByApp *string `json:"lockedByApp,omitempty"`
	// The type of lock.  Currently `edit` is the only supported type.
	LockType *string `json:"lockType,omitempty"`
	// The [password for the template](https://support.docusign.com/s/document-item?bundleId=xry1643227563338&topicId=xwo1578456395432.html). If you are using a lock for a template that has a password or an envelope that is based on a template that has a password, you must enter the `templatePassword` to save the changes.
	TemplatePassword *string `json:"templatePassword,omitempty"`
	// When **true,** a scratchpad is used to edit information.  
	UseScratchPad *string `json:"useScratchPad,omitempty"`
}

// NewLockRequest instantiates a new LockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLockRequest() *LockRequest {
	this := LockRequest{}
	return &this
}

// NewLockRequestWithDefaults instantiates a new LockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLockRequestWithDefaults() *LockRequest {
	this := LockRequest{}
	return &this
}

// GetLockDurationInSeconds returns the LockDurationInSeconds field value if set, zero value otherwise.
func (o *LockRequest) GetLockDurationInSeconds() string {
	if o == nil || IsNil(o.LockDurationInSeconds) {
		var ret string
		return ret
	}
	return *o.LockDurationInSeconds
}

// GetLockDurationInSecondsOk returns a tuple with the LockDurationInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockRequest) GetLockDurationInSecondsOk() (*string, bool) {
	if o == nil || IsNil(o.LockDurationInSeconds) {
		return nil, false
	}
	return o.LockDurationInSeconds, true
}

// HasLockDurationInSeconds returns a boolean if a field has been set.
func (o *LockRequest) HasLockDurationInSeconds() bool {
	if o != nil && !IsNil(o.LockDurationInSeconds) {
		return true
	}

	return false
}

// SetLockDurationInSeconds gets a reference to the given string and assigns it to the LockDurationInSeconds field.
func (o *LockRequest) SetLockDurationInSeconds(v string) {
	o.LockDurationInSeconds = &v
}

// GetLockedByApp returns the LockedByApp field value if set, zero value otherwise.
func (o *LockRequest) GetLockedByApp() string {
	if o == nil || IsNil(o.LockedByApp) {
		var ret string
		return ret
	}
	return *o.LockedByApp
}

// GetLockedByAppOk returns a tuple with the LockedByApp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockRequest) GetLockedByAppOk() (*string, bool) {
	if o == nil || IsNil(o.LockedByApp) {
		return nil, false
	}
	return o.LockedByApp, true
}

// HasLockedByApp returns a boolean if a field has been set.
func (o *LockRequest) HasLockedByApp() bool {
	if o != nil && !IsNil(o.LockedByApp) {
		return true
	}

	return false
}

// SetLockedByApp gets a reference to the given string and assigns it to the LockedByApp field.
func (o *LockRequest) SetLockedByApp(v string) {
	o.LockedByApp = &v
}

// GetLockType returns the LockType field value if set, zero value otherwise.
func (o *LockRequest) GetLockType() string {
	if o == nil || IsNil(o.LockType) {
		var ret string
		return ret
	}
	return *o.LockType
}

// GetLockTypeOk returns a tuple with the LockType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockRequest) GetLockTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LockType) {
		return nil, false
	}
	return o.LockType, true
}

// HasLockType returns a boolean if a field has been set.
func (o *LockRequest) HasLockType() bool {
	if o != nil && !IsNil(o.LockType) {
		return true
	}

	return false
}

// SetLockType gets a reference to the given string and assigns it to the LockType field.
func (o *LockRequest) SetLockType(v string) {
	o.LockType = &v
}

// GetTemplatePassword returns the TemplatePassword field value if set, zero value otherwise.
func (o *LockRequest) GetTemplatePassword() string {
	if o == nil || IsNil(o.TemplatePassword) {
		var ret string
		return ret
	}
	return *o.TemplatePassword
}

// GetTemplatePasswordOk returns a tuple with the TemplatePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockRequest) GetTemplatePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.TemplatePassword) {
		return nil, false
	}
	return o.TemplatePassword, true
}

// HasTemplatePassword returns a boolean if a field has been set.
func (o *LockRequest) HasTemplatePassword() bool {
	if o != nil && !IsNil(o.TemplatePassword) {
		return true
	}

	return false
}

// SetTemplatePassword gets a reference to the given string and assigns it to the TemplatePassword field.
func (o *LockRequest) SetTemplatePassword(v string) {
	o.TemplatePassword = &v
}

// GetUseScratchPad returns the UseScratchPad field value if set, zero value otherwise.
func (o *LockRequest) GetUseScratchPad() string {
	if o == nil || IsNil(o.UseScratchPad) {
		var ret string
		return ret
	}
	return *o.UseScratchPad
}

// GetUseScratchPadOk returns a tuple with the UseScratchPad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LockRequest) GetUseScratchPadOk() (*string, bool) {
	if o == nil || IsNil(o.UseScratchPad) {
		return nil, false
	}
	return o.UseScratchPad, true
}

// HasUseScratchPad returns a boolean if a field has been set.
func (o *LockRequest) HasUseScratchPad() bool {
	if o != nil && !IsNil(o.UseScratchPad) {
		return true
	}

	return false
}

// SetUseScratchPad gets a reference to the given string and assigns it to the UseScratchPad field.
func (o *LockRequest) SetUseScratchPad(v string) {
	o.UseScratchPad = &v
}

func (o LockRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LockRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LockDurationInSeconds) {
		toSerialize["lockDurationInSeconds"] = o.LockDurationInSeconds
	}
	if !IsNil(o.LockedByApp) {
		toSerialize["lockedByApp"] = o.LockedByApp
	}
	if !IsNil(o.LockType) {
		toSerialize["lockType"] = o.LockType
	}
	if !IsNil(o.TemplatePassword) {
		toSerialize["templatePassword"] = o.TemplatePassword
	}
	if !IsNil(o.UseScratchPad) {
		toSerialize["useScratchPad"] = o.UseScratchPad
	}
	return toSerialize, nil
}

type NullableLockRequest struct {
	value *LockRequest
	isSet bool
}

func (v NullableLockRequest) Get() *LockRequest {
	return v.value
}

func (v *NullableLockRequest) Set(val *LockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableLockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableLockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLockRequest(val *LockRequest) *NullableLockRequest {
	return &NullableLockRequest{value: val, isSet: true}
}

func (v NullableLockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


