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

// checks if the ConnectEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectEvents{}

// ConnectEvents Connect event logging information. This object contains sections for regular Connect logs and for Connect failures. 
type ConnectEvents struct {
	// A list of Connect failure logs.
	Failures []ConnectLog `json:"failures,omitempty"`
	// A list of Connect general logs.
	Logs []ConnectLog `json:"logs,omitempty"`
	// The count of records in the log list.
	TotalRecords *string `json:"totalRecords,omitempty"`
	// The type of this tab. Values are:  - Approve - CheckBox - Company - Date - DateSigned, Decline - Email, EmailAddress - EnvelopeId - FirstName - Formula - FullName, InitialHere - InitialHereOptional - LastName - List - Note - Number - Radio - SignerAttachment - SignHere - SignHereOptional - Ssn - Text - Title - Zip5 - Zip5Dash4 
	Type *string `json:"type,omitempty"`
}

// NewConnectEvents instantiates a new ConnectEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectEvents() *ConnectEvents {
	this := ConnectEvents{}
	return &this
}

// NewConnectEventsWithDefaults instantiates a new ConnectEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectEventsWithDefaults() *ConnectEvents {
	this := ConnectEvents{}
	return &this
}

// GetFailures returns the Failures field value if set, zero value otherwise.
func (o *ConnectEvents) GetFailures() []ConnectLog {
	if o == nil || IsNil(o.Failures) {
		var ret []ConnectLog
		return ret
	}
	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectEvents) GetFailuresOk() ([]ConnectLog, bool) {
	if o == nil || IsNil(o.Failures) {
		return nil, false
	}
	return o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *ConnectEvents) HasFailures() bool {
	if o != nil && !IsNil(o.Failures) {
		return true
	}

	return false
}

// SetFailures gets a reference to the given []ConnectLog and assigns it to the Failures field.
func (o *ConnectEvents) SetFailures(v []ConnectLog) {
	o.Failures = v
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *ConnectEvents) GetLogs() []ConnectLog {
	if o == nil || IsNil(o.Logs) {
		var ret []ConnectLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectEvents) GetLogsOk() ([]ConnectLog, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *ConnectEvents) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []ConnectLog and assigns it to the Logs field.
func (o *ConnectEvents) SetLogs(v []ConnectLog) {
	o.Logs = v
}

// GetTotalRecords returns the TotalRecords field value if set, zero value otherwise.
func (o *ConnectEvents) GetTotalRecords() string {
	if o == nil || IsNil(o.TotalRecords) {
		var ret string
		return ret
	}
	return *o.TotalRecords
}

// GetTotalRecordsOk returns a tuple with the TotalRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectEvents) GetTotalRecordsOk() (*string, bool) {
	if o == nil || IsNil(o.TotalRecords) {
		return nil, false
	}
	return o.TotalRecords, true
}

// HasTotalRecords returns a boolean if a field has been set.
func (o *ConnectEvents) HasTotalRecords() bool {
	if o != nil && !IsNil(o.TotalRecords) {
		return true
	}

	return false
}

// SetTotalRecords gets a reference to the given string and assigns it to the TotalRecords field.
func (o *ConnectEvents) SetTotalRecords(v string) {
	o.TotalRecords = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectEvents) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectEvents) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectEvents) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConnectEvents) SetType(v string) {
	o.Type = &v
}

func (o ConnectEvents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Failures) {
		toSerialize["failures"] = o.Failures
	}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	if !IsNil(o.TotalRecords) {
		toSerialize["totalRecords"] = o.TotalRecords
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableConnectEvents struct {
	value *ConnectEvents
	isSet bool
}

func (v NullableConnectEvents) Get() *ConnectEvents {
	return v.value
}

func (v *NullableConnectEvents) Set(val *ConnectEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectEvents(val *ConnectEvents) *NullableConnectEvents {
	return &NullableConnectEvents{value: val, isSet: true}
}

func (v NullableConnectEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

