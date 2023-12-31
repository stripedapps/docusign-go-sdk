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

// checks if the BulkSendBatchSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkSendBatchSummary{}

// BulkSendBatchSummary Summary status of a single batch.
type BulkSendBatchSummary struct {
	// 
	Action *string `json:"action,omitempty"`
	// 
	ActionStatus *string `json:"actionStatus,omitempty"`
	// The batch ID.
	BatchId *string `json:"batchId,omitempty"`
	// The name of the batch.
	BatchName *string `json:"batchName,omitempty"`
	// The number of envelopes in the batch.
	BatchSize *string `json:"batchSize,omitempty"`
	// The batch details URI.
	BatchUri *string `json:"batchUri,omitempty"`
	// Number of envelopes that failed to send.
	Failed *string `json:"failed,omitempty"`
	// Number of envelopes peding processing. 
	Queued *string `json:"queued,omitempty"`
	// Number of envelopes that have been sent.
	Sent *string `json:"sent,omitempty"`
	// The time stamp of when the batch was created in ISO 8601 format.
	SubmittedDate *string `json:"submittedDate,omitempty"`
}

// NewBulkSendBatchSummary instantiates a new BulkSendBatchSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSendBatchSummary() *BulkSendBatchSummary {
	this := BulkSendBatchSummary{}
	return &this
}

// NewBulkSendBatchSummaryWithDefaults instantiates a new BulkSendBatchSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSendBatchSummaryWithDefaults() *BulkSendBatchSummary {
	this := BulkSendBatchSummary{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *BulkSendBatchSummary) SetAction(v string) {
	o.Action = &v
}

// GetActionStatus returns the ActionStatus field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetActionStatus() string {
	if o == nil || IsNil(o.ActionStatus) {
		var ret string
		return ret
	}
	return *o.ActionStatus
}

// GetActionStatusOk returns a tuple with the ActionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetActionStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ActionStatus) {
		return nil, false
	}
	return o.ActionStatus, true
}

// HasActionStatus returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasActionStatus() bool {
	if o != nil && !IsNil(o.ActionStatus) {
		return true
	}

	return false
}

// SetActionStatus gets a reference to the given string and assigns it to the ActionStatus field.
func (o *BulkSendBatchSummary) SetActionStatus(v string) {
	o.ActionStatus = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *BulkSendBatchSummary) SetBatchId(v string) {
	o.BatchId = &v
}

// GetBatchName returns the BatchName field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetBatchName() string {
	if o == nil || IsNil(o.BatchName) {
		var ret string
		return ret
	}
	return *o.BatchName
}

// GetBatchNameOk returns a tuple with the BatchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetBatchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BatchName) {
		return nil, false
	}
	return o.BatchName, true
}

// HasBatchName returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasBatchName() bool {
	if o != nil && !IsNil(o.BatchName) {
		return true
	}

	return false
}

// SetBatchName gets a reference to the given string and assigns it to the BatchName field.
func (o *BulkSendBatchSummary) SetBatchName(v string) {
	o.BatchName = &v
}

// GetBatchSize returns the BatchSize field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetBatchSize() string {
	if o == nil || IsNil(o.BatchSize) {
		var ret string
		return ret
	}
	return *o.BatchSize
}

// GetBatchSizeOk returns a tuple with the BatchSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetBatchSizeOk() (*string, bool) {
	if o == nil || IsNil(o.BatchSize) {
		return nil, false
	}
	return o.BatchSize, true
}

// HasBatchSize returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasBatchSize() bool {
	if o != nil && !IsNil(o.BatchSize) {
		return true
	}

	return false
}

// SetBatchSize gets a reference to the given string and assigns it to the BatchSize field.
func (o *BulkSendBatchSummary) SetBatchSize(v string) {
	o.BatchSize = &v
}

// GetBatchUri returns the BatchUri field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetBatchUri() string {
	if o == nil || IsNil(o.BatchUri) {
		var ret string
		return ret
	}
	return *o.BatchUri
}

// GetBatchUriOk returns a tuple with the BatchUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetBatchUriOk() (*string, bool) {
	if o == nil || IsNil(o.BatchUri) {
		return nil, false
	}
	return o.BatchUri, true
}

// HasBatchUri returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasBatchUri() bool {
	if o != nil && !IsNil(o.BatchUri) {
		return true
	}

	return false
}

// SetBatchUri gets a reference to the given string and assigns it to the BatchUri field.
func (o *BulkSendBatchSummary) SetBatchUri(v string) {
	o.BatchUri = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetFailed() string {
	if o == nil || IsNil(o.Failed) {
		var ret string
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetFailedOk() (*string, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given string and assigns it to the Failed field.
func (o *BulkSendBatchSummary) SetFailed(v string) {
	o.Failed = &v
}

// GetQueued returns the Queued field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetQueued() string {
	if o == nil || IsNil(o.Queued) {
		var ret string
		return ret
	}
	return *o.Queued
}

// GetQueuedOk returns a tuple with the Queued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetQueuedOk() (*string, bool) {
	if o == nil || IsNil(o.Queued) {
		return nil, false
	}
	return o.Queued, true
}

// HasQueued returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasQueued() bool {
	if o != nil && !IsNil(o.Queued) {
		return true
	}

	return false
}

// SetQueued gets a reference to the given string and assigns it to the Queued field.
func (o *BulkSendBatchSummary) SetQueued(v string) {
	o.Queued = &v
}

// GetSent returns the Sent field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetSent() string {
	if o == nil || IsNil(o.Sent) {
		var ret string
		return ret
	}
	return *o.Sent
}

// GetSentOk returns a tuple with the Sent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetSentOk() (*string, bool) {
	if o == nil || IsNil(o.Sent) {
		return nil, false
	}
	return o.Sent, true
}

// HasSent returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasSent() bool {
	if o != nil && !IsNil(o.Sent) {
		return true
	}

	return false
}

// SetSent gets a reference to the given string and assigns it to the Sent field.
func (o *BulkSendBatchSummary) SetSent(v string) {
	o.Sent = &v
}

// GetSubmittedDate returns the SubmittedDate field value if set, zero value otherwise.
func (o *BulkSendBatchSummary) GetSubmittedDate() string {
	if o == nil || IsNil(o.SubmittedDate) {
		var ret string
		return ret
	}
	return *o.SubmittedDate
}

// GetSubmittedDateOk returns a tuple with the SubmittedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendBatchSummary) GetSubmittedDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubmittedDate) {
		return nil, false
	}
	return o.SubmittedDate, true
}

// HasSubmittedDate returns a boolean if a field has been set.
func (o *BulkSendBatchSummary) HasSubmittedDate() bool {
	if o != nil && !IsNil(o.SubmittedDate) {
		return true
	}

	return false
}

// SetSubmittedDate gets a reference to the given string and assigns it to the SubmittedDate field.
func (o *BulkSendBatchSummary) SetSubmittedDate(v string) {
	o.SubmittedDate = &v
}

func (o BulkSendBatchSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkSendBatchSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ActionStatus) {
		toSerialize["actionStatus"] = o.ActionStatus
	}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.BatchName) {
		toSerialize["batchName"] = o.BatchName
	}
	if !IsNil(o.BatchSize) {
		toSerialize["batchSize"] = o.BatchSize
	}
	if !IsNil(o.BatchUri) {
		toSerialize["batchUri"] = o.BatchUri
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Queued) {
		toSerialize["queued"] = o.Queued
	}
	if !IsNil(o.Sent) {
		toSerialize["sent"] = o.Sent
	}
	if !IsNil(o.SubmittedDate) {
		toSerialize["submittedDate"] = o.SubmittedDate
	}
	return toSerialize, nil
}

type NullableBulkSendBatchSummary struct {
	value *BulkSendBatchSummary
	isSet bool
}

func (v NullableBulkSendBatchSummary) Get() *BulkSendBatchSummary {
	return v.value
}

func (v *NullableBulkSendBatchSummary) Set(val *BulkSendBatchSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSendBatchSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSendBatchSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSendBatchSummary(val *BulkSendBatchSummary) *NullableBulkSendBatchSummary {
	return &NullableBulkSendBatchSummary{value: val, isSet: true}
}

func (v NullableBulkSendBatchSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSendBatchSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


