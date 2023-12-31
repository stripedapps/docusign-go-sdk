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

// checks if the EnvelopePublishTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopePublishTransaction{}

// EnvelopePublishTransaction 
type EnvelopePublishTransaction struct {
	// 
	ApplyConnectSettings *string `json:"applyConnectSettings,omitempty"`
	// 
	EnvelopeCount *string `json:"envelopeCount,omitempty"`
	// 
	EnvelopeLevelErrorRollups []EnvelopePublishTransactionErrorRollup `json:"envelopeLevelErrorRollups,omitempty"`
	// The ID of the publish transaction.
	EnvelopePublishTransactionId *string `json:"envelopePublishTransactionId,omitempty"`
	// 
	ErrorCount *string `json:"errorCount,omitempty"`
	// 
	FileLevelErrors []string `json:"fileLevelErrors,omitempty"`
	// 
	NoActionRequiredEnvelopeCount *string `json:"noActionRequiredEnvelopeCount,omitempty"`
	// 
	ProcessedEnvelopeCount *string `json:"processedEnvelopeCount,omitempty"`
	// The status of the transaction. Valid values:  * `unprocessed` * `processing` * `complete` * `fatal_error` 
	ProcessingStatus *string `json:"processingStatus,omitempty"`
	// 
	ResultsUri *string `json:"resultsUri,omitempty"`
	// 
	SubmissionDate *string `json:"submissionDate,omitempty"`
	SubmittedByUserInfo *UserInfo `json:"submittedByUserInfo,omitempty"`
	// 
	SubmittedForPublishingEnvelopeCount *string `json:"submittedForPublishingEnvelopeCount,omitempty"`
}

// NewEnvelopePublishTransaction instantiates a new EnvelopePublishTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopePublishTransaction() *EnvelopePublishTransaction {
	this := EnvelopePublishTransaction{}
	return &this
}

// NewEnvelopePublishTransactionWithDefaults instantiates a new EnvelopePublishTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopePublishTransactionWithDefaults() *EnvelopePublishTransaction {
	this := EnvelopePublishTransaction{}
	return &this
}

// GetApplyConnectSettings returns the ApplyConnectSettings field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetApplyConnectSettings() string {
	if o == nil || IsNil(o.ApplyConnectSettings) {
		var ret string
		return ret
	}
	return *o.ApplyConnectSettings
}

// GetApplyConnectSettingsOk returns a tuple with the ApplyConnectSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetApplyConnectSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.ApplyConnectSettings) {
		return nil, false
	}
	return o.ApplyConnectSettings, true
}

// HasApplyConnectSettings returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasApplyConnectSettings() bool {
	if o != nil && !IsNil(o.ApplyConnectSettings) {
		return true
	}

	return false
}

// SetApplyConnectSettings gets a reference to the given string and assigns it to the ApplyConnectSettings field.
func (o *EnvelopePublishTransaction) SetApplyConnectSettings(v string) {
	o.ApplyConnectSettings = &v
}

// GetEnvelopeCount returns the EnvelopeCount field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetEnvelopeCount() string {
	if o == nil || IsNil(o.EnvelopeCount) {
		var ret string
		return ret
	}
	return *o.EnvelopeCount
}

// GetEnvelopeCountOk returns a tuple with the EnvelopeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetEnvelopeCountOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeCount) {
		return nil, false
	}
	return o.EnvelopeCount, true
}

// HasEnvelopeCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasEnvelopeCount() bool {
	if o != nil && !IsNil(o.EnvelopeCount) {
		return true
	}

	return false
}

// SetEnvelopeCount gets a reference to the given string and assigns it to the EnvelopeCount field.
func (o *EnvelopePublishTransaction) SetEnvelopeCount(v string) {
	o.EnvelopeCount = &v
}

// GetEnvelopeLevelErrorRollups returns the EnvelopeLevelErrorRollups field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetEnvelopeLevelErrorRollups() []EnvelopePublishTransactionErrorRollup {
	if o == nil || IsNil(o.EnvelopeLevelErrorRollups) {
		var ret []EnvelopePublishTransactionErrorRollup
		return ret
	}
	return o.EnvelopeLevelErrorRollups
}

// GetEnvelopeLevelErrorRollupsOk returns a tuple with the EnvelopeLevelErrorRollups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetEnvelopeLevelErrorRollupsOk() ([]EnvelopePublishTransactionErrorRollup, bool) {
	if o == nil || IsNil(o.EnvelopeLevelErrorRollups) {
		return nil, false
	}
	return o.EnvelopeLevelErrorRollups, true
}

// HasEnvelopeLevelErrorRollups returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasEnvelopeLevelErrorRollups() bool {
	if o != nil && !IsNil(o.EnvelopeLevelErrorRollups) {
		return true
	}

	return false
}

// SetEnvelopeLevelErrorRollups gets a reference to the given []EnvelopePublishTransactionErrorRollup and assigns it to the EnvelopeLevelErrorRollups field.
func (o *EnvelopePublishTransaction) SetEnvelopeLevelErrorRollups(v []EnvelopePublishTransactionErrorRollup) {
	o.EnvelopeLevelErrorRollups = v
}

// GetEnvelopePublishTransactionId returns the EnvelopePublishTransactionId field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetEnvelopePublishTransactionId() string {
	if o == nil || IsNil(o.EnvelopePublishTransactionId) {
		var ret string
		return ret
	}
	return *o.EnvelopePublishTransactionId
}

// GetEnvelopePublishTransactionIdOk returns a tuple with the EnvelopePublishTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetEnvelopePublishTransactionIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopePublishTransactionId) {
		return nil, false
	}
	return o.EnvelopePublishTransactionId, true
}

// HasEnvelopePublishTransactionId returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasEnvelopePublishTransactionId() bool {
	if o != nil && !IsNil(o.EnvelopePublishTransactionId) {
		return true
	}

	return false
}

// SetEnvelopePublishTransactionId gets a reference to the given string and assigns it to the EnvelopePublishTransactionId field.
func (o *EnvelopePublishTransaction) SetEnvelopePublishTransactionId(v string) {
	o.EnvelopePublishTransactionId = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetErrorCount() string {
	if o == nil || IsNil(o.ErrorCount) {
		var ret string
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetErrorCountOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasErrorCount() bool {
	if o != nil && !IsNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given string and assigns it to the ErrorCount field.
func (o *EnvelopePublishTransaction) SetErrorCount(v string) {
	o.ErrorCount = &v
}

// GetFileLevelErrors returns the FileLevelErrors field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetFileLevelErrors() []string {
	if o == nil || IsNil(o.FileLevelErrors) {
		var ret []string
		return ret
	}
	return o.FileLevelErrors
}

// GetFileLevelErrorsOk returns a tuple with the FileLevelErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetFileLevelErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.FileLevelErrors) {
		return nil, false
	}
	return o.FileLevelErrors, true
}

// HasFileLevelErrors returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasFileLevelErrors() bool {
	if o != nil && !IsNil(o.FileLevelErrors) {
		return true
	}

	return false
}

// SetFileLevelErrors gets a reference to the given []string and assigns it to the FileLevelErrors field.
func (o *EnvelopePublishTransaction) SetFileLevelErrors(v []string) {
	o.FileLevelErrors = v
}

// GetNoActionRequiredEnvelopeCount returns the NoActionRequiredEnvelopeCount field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetNoActionRequiredEnvelopeCount() string {
	if o == nil || IsNil(o.NoActionRequiredEnvelopeCount) {
		var ret string
		return ret
	}
	return *o.NoActionRequiredEnvelopeCount
}

// GetNoActionRequiredEnvelopeCountOk returns a tuple with the NoActionRequiredEnvelopeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetNoActionRequiredEnvelopeCountOk() (*string, bool) {
	if o == nil || IsNil(o.NoActionRequiredEnvelopeCount) {
		return nil, false
	}
	return o.NoActionRequiredEnvelopeCount, true
}

// HasNoActionRequiredEnvelopeCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasNoActionRequiredEnvelopeCount() bool {
	if o != nil && !IsNil(o.NoActionRequiredEnvelopeCount) {
		return true
	}

	return false
}

// SetNoActionRequiredEnvelopeCount gets a reference to the given string and assigns it to the NoActionRequiredEnvelopeCount field.
func (o *EnvelopePublishTransaction) SetNoActionRequiredEnvelopeCount(v string) {
	o.NoActionRequiredEnvelopeCount = &v
}

// GetProcessedEnvelopeCount returns the ProcessedEnvelopeCount field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetProcessedEnvelopeCount() string {
	if o == nil || IsNil(o.ProcessedEnvelopeCount) {
		var ret string
		return ret
	}
	return *o.ProcessedEnvelopeCount
}

// GetProcessedEnvelopeCountOk returns a tuple with the ProcessedEnvelopeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetProcessedEnvelopeCountOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessedEnvelopeCount) {
		return nil, false
	}
	return o.ProcessedEnvelopeCount, true
}

// HasProcessedEnvelopeCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasProcessedEnvelopeCount() bool {
	if o != nil && !IsNil(o.ProcessedEnvelopeCount) {
		return true
	}

	return false
}

// SetProcessedEnvelopeCount gets a reference to the given string and assigns it to the ProcessedEnvelopeCount field.
func (o *EnvelopePublishTransaction) SetProcessedEnvelopeCount(v string) {
	o.ProcessedEnvelopeCount = &v
}

// GetProcessingStatus returns the ProcessingStatus field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetProcessingStatus() string {
	if o == nil || IsNil(o.ProcessingStatus) {
		var ret string
		return ret
	}
	return *o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetProcessingStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ProcessingStatus) {
		return nil, false
	}
	return o.ProcessingStatus, true
}

// HasProcessingStatus returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasProcessingStatus() bool {
	if o != nil && !IsNil(o.ProcessingStatus) {
		return true
	}

	return false
}

// SetProcessingStatus gets a reference to the given string and assigns it to the ProcessingStatus field.
func (o *EnvelopePublishTransaction) SetProcessingStatus(v string) {
	o.ProcessingStatus = &v
}

// GetResultsUri returns the ResultsUri field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetResultsUri() string {
	if o == nil || IsNil(o.ResultsUri) {
		var ret string
		return ret
	}
	return *o.ResultsUri
}

// GetResultsUriOk returns a tuple with the ResultsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetResultsUriOk() (*string, bool) {
	if o == nil || IsNil(o.ResultsUri) {
		return nil, false
	}
	return o.ResultsUri, true
}

// HasResultsUri returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasResultsUri() bool {
	if o != nil && !IsNil(o.ResultsUri) {
		return true
	}

	return false
}

// SetResultsUri gets a reference to the given string and assigns it to the ResultsUri field.
func (o *EnvelopePublishTransaction) SetResultsUri(v string) {
	o.ResultsUri = &v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetSubmissionDate() string {
	if o == nil || IsNil(o.SubmissionDate) {
		var ret string
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetSubmissionDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubmissionDate) {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasSubmissionDate() bool {
	if o != nil && !IsNil(o.SubmissionDate) {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given string and assigns it to the SubmissionDate field.
func (o *EnvelopePublishTransaction) SetSubmissionDate(v string) {
	o.SubmissionDate = &v
}

// GetSubmittedByUserInfo returns the SubmittedByUserInfo field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetSubmittedByUserInfo() UserInfo {
	if o == nil || IsNil(o.SubmittedByUserInfo) {
		var ret UserInfo
		return ret
	}
	return *o.SubmittedByUserInfo
}

// GetSubmittedByUserInfoOk returns a tuple with the SubmittedByUserInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetSubmittedByUserInfoOk() (*UserInfo, bool) {
	if o == nil || IsNil(o.SubmittedByUserInfo) {
		return nil, false
	}
	return o.SubmittedByUserInfo, true
}

// HasSubmittedByUserInfo returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasSubmittedByUserInfo() bool {
	if o != nil && !IsNil(o.SubmittedByUserInfo) {
		return true
	}

	return false
}

// SetSubmittedByUserInfo gets a reference to the given UserInfo and assigns it to the SubmittedByUserInfo field.
func (o *EnvelopePublishTransaction) SetSubmittedByUserInfo(v UserInfo) {
	o.SubmittedByUserInfo = &v
}

// GetSubmittedForPublishingEnvelopeCount returns the SubmittedForPublishingEnvelopeCount field value if set, zero value otherwise.
func (o *EnvelopePublishTransaction) GetSubmittedForPublishingEnvelopeCount() string {
	if o == nil || IsNil(o.SubmittedForPublishingEnvelopeCount) {
		var ret string
		return ret
	}
	return *o.SubmittedForPublishingEnvelopeCount
}

// GetSubmittedForPublishingEnvelopeCountOk returns a tuple with the SubmittedForPublishingEnvelopeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopePublishTransaction) GetSubmittedForPublishingEnvelopeCountOk() (*string, bool) {
	if o == nil || IsNil(o.SubmittedForPublishingEnvelopeCount) {
		return nil, false
	}
	return o.SubmittedForPublishingEnvelopeCount, true
}

// HasSubmittedForPublishingEnvelopeCount returns a boolean if a field has been set.
func (o *EnvelopePublishTransaction) HasSubmittedForPublishingEnvelopeCount() bool {
	if o != nil && !IsNil(o.SubmittedForPublishingEnvelopeCount) {
		return true
	}

	return false
}

// SetSubmittedForPublishingEnvelopeCount gets a reference to the given string and assigns it to the SubmittedForPublishingEnvelopeCount field.
func (o *EnvelopePublishTransaction) SetSubmittedForPublishingEnvelopeCount(v string) {
	o.SubmittedForPublishingEnvelopeCount = &v
}

func (o EnvelopePublishTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopePublishTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyConnectSettings) {
		toSerialize["applyConnectSettings"] = o.ApplyConnectSettings
	}
	if !IsNil(o.EnvelopeCount) {
		toSerialize["envelopeCount"] = o.EnvelopeCount
	}
	if !IsNil(o.EnvelopeLevelErrorRollups) {
		toSerialize["envelopeLevelErrorRollups"] = o.EnvelopeLevelErrorRollups
	}
	if !IsNil(o.EnvelopePublishTransactionId) {
		toSerialize["envelopePublishTransactionId"] = o.EnvelopePublishTransactionId
	}
	if !IsNil(o.ErrorCount) {
		toSerialize["errorCount"] = o.ErrorCount
	}
	if !IsNil(o.FileLevelErrors) {
		toSerialize["fileLevelErrors"] = o.FileLevelErrors
	}
	if !IsNil(o.NoActionRequiredEnvelopeCount) {
		toSerialize["noActionRequiredEnvelopeCount"] = o.NoActionRequiredEnvelopeCount
	}
	if !IsNil(o.ProcessedEnvelopeCount) {
		toSerialize["processedEnvelopeCount"] = o.ProcessedEnvelopeCount
	}
	if !IsNil(o.ProcessingStatus) {
		toSerialize["processingStatus"] = o.ProcessingStatus
	}
	if !IsNil(o.ResultsUri) {
		toSerialize["resultsUri"] = o.ResultsUri
	}
	if !IsNil(o.SubmissionDate) {
		toSerialize["submissionDate"] = o.SubmissionDate
	}
	if !IsNil(o.SubmittedByUserInfo) {
		toSerialize["submittedByUserInfo"] = o.SubmittedByUserInfo
	}
	if !IsNil(o.SubmittedForPublishingEnvelopeCount) {
		toSerialize["submittedForPublishingEnvelopeCount"] = o.SubmittedForPublishingEnvelopeCount
	}
	return toSerialize, nil
}

type NullableEnvelopePublishTransaction struct {
	value *EnvelopePublishTransaction
	isSet bool
}

func (v NullableEnvelopePublishTransaction) Get() *EnvelopePublishTransaction {
	return v.value
}

func (v *NullableEnvelopePublishTransaction) Set(val *EnvelopePublishTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopePublishTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopePublishTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopePublishTransaction(val *EnvelopePublishTransaction) *NullableEnvelopePublishTransaction {
	return &NullableEnvelopePublishTransaction{value: val, isSet: true}
}

func (v NullableEnvelopePublishTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopePublishTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


