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

// checks if the DiagnosticsSettingsInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosticsSettingsInformation{}

// DiagnosticsSettingsInformation 
type DiagnosticsSettingsInformation struct {
	//  When **true,** enables API request logging for the user. 
	ApiRequestLogging *string `json:"apiRequestLogging,omitempty"`
	// Specifies the maximum number of API requests to log.
	ApiRequestLogMaxEntries *string `json:"apiRequestLogMaxEntries,omitempty"`
	// Indicates the remaining number of API requests that can be logged.
	ApiRequestLogRemainingEntries *string `json:"apiRequestLogRemainingEntries,omitempty"`
}

// NewDiagnosticsSettingsInformation instantiates a new DiagnosticsSettingsInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosticsSettingsInformation() *DiagnosticsSettingsInformation {
	this := DiagnosticsSettingsInformation{}
	return &this
}

// NewDiagnosticsSettingsInformationWithDefaults instantiates a new DiagnosticsSettingsInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosticsSettingsInformationWithDefaults() *DiagnosticsSettingsInformation {
	this := DiagnosticsSettingsInformation{}
	return &this
}

// GetApiRequestLogging returns the ApiRequestLogging field value if set, zero value otherwise.
func (o *DiagnosticsSettingsInformation) GetApiRequestLogging() string {
	if o == nil || IsNil(o.ApiRequestLogging) {
		var ret string
		return ret
	}
	return *o.ApiRequestLogging
}

// GetApiRequestLoggingOk returns a tuple with the ApiRequestLogging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsSettingsInformation) GetApiRequestLoggingOk() (*string, bool) {
	if o == nil || IsNil(o.ApiRequestLogging) {
		return nil, false
	}
	return o.ApiRequestLogging, true
}

// HasApiRequestLogging returns a boolean if a field has been set.
func (o *DiagnosticsSettingsInformation) HasApiRequestLogging() bool {
	if o != nil && !IsNil(o.ApiRequestLogging) {
		return true
	}

	return false
}

// SetApiRequestLogging gets a reference to the given string and assigns it to the ApiRequestLogging field.
func (o *DiagnosticsSettingsInformation) SetApiRequestLogging(v string) {
	o.ApiRequestLogging = &v
}

// GetApiRequestLogMaxEntries returns the ApiRequestLogMaxEntries field value if set, zero value otherwise.
func (o *DiagnosticsSettingsInformation) GetApiRequestLogMaxEntries() string {
	if o == nil || IsNil(o.ApiRequestLogMaxEntries) {
		var ret string
		return ret
	}
	return *o.ApiRequestLogMaxEntries
}

// GetApiRequestLogMaxEntriesOk returns a tuple with the ApiRequestLogMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsSettingsInformation) GetApiRequestLogMaxEntriesOk() (*string, bool) {
	if o == nil || IsNil(o.ApiRequestLogMaxEntries) {
		return nil, false
	}
	return o.ApiRequestLogMaxEntries, true
}

// HasApiRequestLogMaxEntries returns a boolean if a field has been set.
func (o *DiagnosticsSettingsInformation) HasApiRequestLogMaxEntries() bool {
	if o != nil && !IsNil(o.ApiRequestLogMaxEntries) {
		return true
	}

	return false
}

// SetApiRequestLogMaxEntries gets a reference to the given string and assigns it to the ApiRequestLogMaxEntries field.
func (o *DiagnosticsSettingsInformation) SetApiRequestLogMaxEntries(v string) {
	o.ApiRequestLogMaxEntries = &v
}

// GetApiRequestLogRemainingEntries returns the ApiRequestLogRemainingEntries field value if set, zero value otherwise.
func (o *DiagnosticsSettingsInformation) GetApiRequestLogRemainingEntries() string {
	if o == nil || IsNil(o.ApiRequestLogRemainingEntries) {
		var ret string
		return ret
	}
	return *o.ApiRequestLogRemainingEntries
}

// GetApiRequestLogRemainingEntriesOk returns a tuple with the ApiRequestLogRemainingEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticsSettingsInformation) GetApiRequestLogRemainingEntriesOk() (*string, bool) {
	if o == nil || IsNil(o.ApiRequestLogRemainingEntries) {
		return nil, false
	}
	return o.ApiRequestLogRemainingEntries, true
}

// HasApiRequestLogRemainingEntries returns a boolean if a field has been set.
func (o *DiagnosticsSettingsInformation) HasApiRequestLogRemainingEntries() bool {
	if o != nil && !IsNil(o.ApiRequestLogRemainingEntries) {
		return true
	}

	return false
}

// SetApiRequestLogRemainingEntries gets a reference to the given string and assigns it to the ApiRequestLogRemainingEntries field.
func (o *DiagnosticsSettingsInformation) SetApiRequestLogRemainingEntries(v string) {
	o.ApiRequestLogRemainingEntries = &v
}

func (o DiagnosticsSettingsInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosticsSettingsInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiRequestLogging) {
		toSerialize["apiRequestLogging"] = o.ApiRequestLogging
	}
	if !IsNil(o.ApiRequestLogMaxEntries) {
		toSerialize["apiRequestLogMaxEntries"] = o.ApiRequestLogMaxEntries
	}
	if !IsNil(o.ApiRequestLogRemainingEntries) {
		toSerialize["apiRequestLogRemainingEntries"] = o.ApiRequestLogRemainingEntries
	}
	return toSerialize, nil
}

type NullableDiagnosticsSettingsInformation struct {
	value *DiagnosticsSettingsInformation
	isSet bool
}

func (v NullableDiagnosticsSettingsInformation) Get() *DiagnosticsSettingsInformation {
	return v.value
}

func (v *NullableDiagnosticsSettingsInformation) Set(val *DiagnosticsSettingsInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosticsSettingsInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosticsSettingsInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosticsSettingsInformation(val *DiagnosticsSettingsInformation) *NullableDiagnosticsSettingsInformation {
	return &NullableDiagnosticsSettingsInformation{value: val, isSet: true}
}

func (v NullableDiagnosticsSettingsInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosticsSettingsInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


