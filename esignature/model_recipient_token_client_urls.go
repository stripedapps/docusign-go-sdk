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

// checks if the RecipientTokenClientURLs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientTokenClientURLs{}

// RecipientTokenClientURLs 
type RecipientTokenClientURLs struct {
	// 
	OnAccessCodeFailed *string `json:"onAccessCodeFailed,omitempty"`
	// 
	OnCancel *string `json:"onCancel,omitempty"`
	// 
	OnDecline *string `json:"onDecline,omitempty"`
	// 
	OnException *string `json:"onException,omitempty"`
	// 
	OnFaxPending *string `json:"onFaxPending,omitempty"`
	// 
	OnIdCheckFailed *string `json:"onIdCheckFailed,omitempty"`
	// 
	OnSessionTimeout *string `json:"onSessionTimeout,omitempty"`
	// 
	OnSigningComplete *string `json:"onSigningComplete,omitempty"`
	// 
	OnTTLExpired *string `json:"onTTLExpired,omitempty"`
	// 
	OnViewingComplete *string `json:"onViewingComplete,omitempty"`
}

// NewRecipientTokenClientURLs instantiates a new RecipientTokenClientURLs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientTokenClientURLs() *RecipientTokenClientURLs {
	this := RecipientTokenClientURLs{}
	return &this
}

// NewRecipientTokenClientURLsWithDefaults instantiates a new RecipientTokenClientURLs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientTokenClientURLsWithDefaults() *RecipientTokenClientURLs {
	this := RecipientTokenClientURLs{}
	return &this
}

// GetOnAccessCodeFailed returns the OnAccessCodeFailed field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnAccessCodeFailed() string {
	if o == nil || IsNil(o.OnAccessCodeFailed) {
		var ret string
		return ret
	}
	return *o.OnAccessCodeFailed
}

// GetOnAccessCodeFailedOk returns a tuple with the OnAccessCodeFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnAccessCodeFailedOk() (*string, bool) {
	if o == nil || IsNil(o.OnAccessCodeFailed) {
		return nil, false
	}
	return o.OnAccessCodeFailed, true
}

// HasOnAccessCodeFailed returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnAccessCodeFailed() bool {
	if o != nil && !IsNil(o.OnAccessCodeFailed) {
		return true
	}

	return false
}

// SetOnAccessCodeFailed gets a reference to the given string and assigns it to the OnAccessCodeFailed field.
func (o *RecipientTokenClientURLs) SetOnAccessCodeFailed(v string) {
	o.OnAccessCodeFailed = &v
}

// GetOnCancel returns the OnCancel field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnCancel() string {
	if o == nil || IsNil(o.OnCancel) {
		var ret string
		return ret
	}
	return *o.OnCancel
}

// GetOnCancelOk returns a tuple with the OnCancel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnCancelOk() (*string, bool) {
	if o == nil || IsNil(o.OnCancel) {
		return nil, false
	}
	return o.OnCancel, true
}

// HasOnCancel returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnCancel() bool {
	if o != nil && !IsNil(o.OnCancel) {
		return true
	}

	return false
}

// SetOnCancel gets a reference to the given string and assigns it to the OnCancel field.
func (o *RecipientTokenClientURLs) SetOnCancel(v string) {
	o.OnCancel = &v
}

// GetOnDecline returns the OnDecline field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnDecline() string {
	if o == nil || IsNil(o.OnDecline) {
		var ret string
		return ret
	}
	return *o.OnDecline
}

// GetOnDeclineOk returns a tuple with the OnDecline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnDeclineOk() (*string, bool) {
	if o == nil || IsNil(o.OnDecline) {
		return nil, false
	}
	return o.OnDecline, true
}

// HasOnDecline returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnDecline() bool {
	if o != nil && !IsNil(o.OnDecline) {
		return true
	}

	return false
}

// SetOnDecline gets a reference to the given string and assigns it to the OnDecline field.
func (o *RecipientTokenClientURLs) SetOnDecline(v string) {
	o.OnDecline = &v
}

// GetOnException returns the OnException field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnException() string {
	if o == nil || IsNil(o.OnException) {
		var ret string
		return ret
	}
	return *o.OnException
}

// GetOnExceptionOk returns a tuple with the OnException field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnExceptionOk() (*string, bool) {
	if o == nil || IsNil(o.OnException) {
		return nil, false
	}
	return o.OnException, true
}

// HasOnException returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnException() bool {
	if o != nil && !IsNil(o.OnException) {
		return true
	}

	return false
}

// SetOnException gets a reference to the given string and assigns it to the OnException field.
func (o *RecipientTokenClientURLs) SetOnException(v string) {
	o.OnException = &v
}

// GetOnFaxPending returns the OnFaxPending field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnFaxPending() string {
	if o == nil || IsNil(o.OnFaxPending) {
		var ret string
		return ret
	}
	return *o.OnFaxPending
}

// GetOnFaxPendingOk returns a tuple with the OnFaxPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnFaxPendingOk() (*string, bool) {
	if o == nil || IsNil(o.OnFaxPending) {
		return nil, false
	}
	return o.OnFaxPending, true
}

// HasOnFaxPending returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnFaxPending() bool {
	if o != nil && !IsNil(o.OnFaxPending) {
		return true
	}

	return false
}

// SetOnFaxPending gets a reference to the given string and assigns it to the OnFaxPending field.
func (o *RecipientTokenClientURLs) SetOnFaxPending(v string) {
	o.OnFaxPending = &v
}

// GetOnIdCheckFailed returns the OnIdCheckFailed field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnIdCheckFailed() string {
	if o == nil || IsNil(o.OnIdCheckFailed) {
		var ret string
		return ret
	}
	return *o.OnIdCheckFailed
}

// GetOnIdCheckFailedOk returns a tuple with the OnIdCheckFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnIdCheckFailedOk() (*string, bool) {
	if o == nil || IsNil(o.OnIdCheckFailed) {
		return nil, false
	}
	return o.OnIdCheckFailed, true
}

// HasOnIdCheckFailed returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnIdCheckFailed() bool {
	if o != nil && !IsNil(o.OnIdCheckFailed) {
		return true
	}

	return false
}

// SetOnIdCheckFailed gets a reference to the given string and assigns it to the OnIdCheckFailed field.
func (o *RecipientTokenClientURLs) SetOnIdCheckFailed(v string) {
	o.OnIdCheckFailed = &v
}

// GetOnSessionTimeout returns the OnSessionTimeout field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnSessionTimeout() string {
	if o == nil || IsNil(o.OnSessionTimeout) {
		var ret string
		return ret
	}
	return *o.OnSessionTimeout
}

// GetOnSessionTimeoutOk returns a tuple with the OnSessionTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnSessionTimeoutOk() (*string, bool) {
	if o == nil || IsNil(o.OnSessionTimeout) {
		return nil, false
	}
	return o.OnSessionTimeout, true
}

// HasOnSessionTimeout returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnSessionTimeout() bool {
	if o != nil && !IsNil(o.OnSessionTimeout) {
		return true
	}

	return false
}

// SetOnSessionTimeout gets a reference to the given string and assigns it to the OnSessionTimeout field.
func (o *RecipientTokenClientURLs) SetOnSessionTimeout(v string) {
	o.OnSessionTimeout = &v
}

// GetOnSigningComplete returns the OnSigningComplete field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnSigningComplete() string {
	if o == nil || IsNil(o.OnSigningComplete) {
		var ret string
		return ret
	}
	return *o.OnSigningComplete
}

// GetOnSigningCompleteOk returns a tuple with the OnSigningComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnSigningCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.OnSigningComplete) {
		return nil, false
	}
	return o.OnSigningComplete, true
}

// HasOnSigningComplete returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnSigningComplete() bool {
	if o != nil && !IsNil(o.OnSigningComplete) {
		return true
	}

	return false
}

// SetOnSigningComplete gets a reference to the given string and assigns it to the OnSigningComplete field.
func (o *RecipientTokenClientURLs) SetOnSigningComplete(v string) {
	o.OnSigningComplete = &v
}

// GetOnTTLExpired returns the OnTTLExpired field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnTTLExpired() string {
	if o == nil || IsNil(o.OnTTLExpired) {
		var ret string
		return ret
	}
	return *o.OnTTLExpired
}

// GetOnTTLExpiredOk returns a tuple with the OnTTLExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnTTLExpiredOk() (*string, bool) {
	if o == nil || IsNil(o.OnTTLExpired) {
		return nil, false
	}
	return o.OnTTLExpired, true
}

// HasOnTTLExpired returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnTTLExpired() bool {
	if o != nil && !IsNil(o.OnTTLExpired) {
		return true
	}

	return false
}

// SetOnTTLExpired gets a reference to the given string and assigns it to the OnTTLExpired field.
func (o *RecipientTokenClientURLs) SetOnTTLExpired(v string) {
	o.OnTTLExpired = &v
}

// GetOnViewingComplete returns the OnViewingComplete field value if set, zero value otherwise.
func (o *RecipientTokenClientURLs) GetOnViewingComplete() string {
	if o == nil || IsNil(o.OnViewingComplete) {
		var ret string
		return ret
	}
	return *o.OnViewingComplete
}

// GetOnViewingCompleteOk returns a tuple with the OnViewingComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientTokenClientURLs) GetOnViewingCompleteOk() (*string, bool) {
	if o == nil || IsNil(o.OnViewingComplete) {
		return nil, false
	}
	return o.OnViewingComplete, true
}

// HasOnViewingComplete returns a boolean if a field has been set.
func (o *RecipientTokenClientURLs) HasOnViewingComplete() bool {
	if o != nil && !IsNil(o.OnViewingComplete) {
		return true
	}

	return false
}

// SetOnViewingComplete gets a reference to the given string and assigns it to the OnViewingComplete field.
func (o *RecipientTokenClientURLs) SetOnViewingComplete(v string) {
	o.OnViewingComplete = &v
}

func (o RecipientTokenClientURLs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientTokenClientURLs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OnAccessCodeFailed) {
		toSerialize["onAccessCodeFailed"] = o.OnAccessCodeFailed
	}
	if !IsNil(o.OnCancel) {
		toSerialize["onCancel"] = o.OnCancel
	}
	if !IsNil(o.OnDecline) {
		toSerialize["onDecline"] = o.OnDecline
	}
	if !IsNil(o.OnException) {
		toSerialize["onException"] = o.OnException
	}
	if !IsNil(o.OnFaxPending) {
		toSerialize["onFaxPending"] = o.OnFaxPending
	}
	if !IsNil(o.OnIdCheckFailed) {
		toSerialize["onIdCheckFailed"] = o.OnIdCheckFailed
	}
	if !IsNil(o.OnSessionTimeout) {
		toSerialize["onSessionTimeout"] = o.OnSessionTimeout
	}
	if !IsNil(o.OnSigningComplete) {
		toSerialize["onSigningComplete"] = o.OnSigningComplete
	}
	if !IsNil(o.OnTTLExpired) {
		toSerialize["onTTLExpired"] = o.OnTTLExpired
	}
	if !IsNil(o.OnViewingComplete) {
		toSerialize["onViewingComplete"] = o.OnViewingComplete
	}
	return toSerialize, nil
}

type NullableRecipientTokenClientURLs struct {
	value *RecipientTokenClientURLs
	isSet bool
}

func (v NullableRecipientTokenClientURLs) Get() *RecipientTokenClientURLs {
	return v.value
}

func (v *NullableRecipientTokenClientURLs) Set(val *RecipientTokenClientURLs) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientTokenClientURLs) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientTokenClientURLs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientTokenClientURLs(val *RecipientTokenClientURLs) *NullableRecipientTokenClientURLs {
	return &NullableRecipientTokenClientURLs{value: val, isSet: true}
}

func (v NullableRecipientTokenClientURLs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientTokenClientURLs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


