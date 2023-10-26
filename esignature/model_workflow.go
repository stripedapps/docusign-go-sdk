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

// checks if the Workflow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Workflow{}

// Workflow Describes the workflow for an envelope.
type Workflow struct {
	// The `workflowStepId` of the current step. This is not an index into the `workflowSteps` array in this object. See the `workflowStep` object.
	CurrentWorkflowStepId *string `json:"currentWorkflowStepId,omitempty"`
	// The ISO 8601 timestamp of when the envelope is scheduled to be sent, if applicable. Its value is the maximum of the `resumeDate` property on `scheduledSending` and the `resumeDate` property on the current `workflowStep`.  This property is read-only.
	ResumeDate *string `json:"resumeDate,omitempty"`
	ScheduledSending *ScheduledSending `json:"scheduledSending,omitempty"`
	// The status of the workflow:  - `paused` if the workflow is paused - `in_progress` if the workflow is in progress
	WorkflowStatus *string `json:"workflowStatus,omitempty"`
	// An array of workflow steps.
	WorkflowSteps []WorkflowStep `json:"workflowSteps,omitempty"`
}

// NewWorkflow instantiates a new Workflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow() *Workflow {
	this := Workflow{}
	return &this
}

// NewWorkflowWithDefaults instantiates a new Workflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWithDefaults() *Workflow {
	this := Workflow{}
	return &this
}

// GetCurrentWorkflowStepId returns the CurrentWorkflowStepId field value if set, zero value otherwise.
func (o *Workflow) GetCurrentWorkflowStepId() string {
	if o == nil || IsNil(o.CurrentWorkflowStepId) {
		var ret string
		return ret
	}
	return *o.CurrentWorkflowStepId
}

// GetCurrentWorkflowStepIdOk returns a tuple with the CurrentWorkflowStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetCurrentWorkflowStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentWorkflowStepId) {
		return nil, false
	}
	return o.CurrentWorkflowStepId, true
}

// HasCurrentWorkflowStepId returns a boolean if a field has been set.
func (o *Workflow) HasCurrentWorkflowStepId() bool {
	if o != nil && !IsNil(o.CurrentWorkflowStepId) {
		return true
	}

	return false
}

// SetCurrentWorkflowStepId gets a reference to the given string and assigns it to the CurrentWorkflowStepId field.
func (o *Workflow) SetCurrentWorkflowStepId(v string) {
	o.CurrentWorkflowStepId = &v
}

// GetResumeDate returns the ResumeDate field value if set, zero value otherwise.
func (o *Workflow) GetResumeDate() string {
	if o == nil || IsNil(o.ResumeDate) {
		var ret string
		return ret
	}
	return *o.ResumeDate
}

// GetResumeDateOk returns a tuple with the ResumeDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetResumeDateOk() (*string, bool) {
	if o == nil || IsNil(o.ResumeDate) {
		return nil, false
	}
	return o.ResumeDate, true
}

// HasResumeDate returns a boolean if a field has been set.
func (o *Workflow) HasResumeDate() bool {
	if o != nil && !IsNil(o.ResumeDate) {
		return true
	}

	return false
}

// SetResumeDate gets a reference to the given string and assigns it to the ResumeDate field.
func (o *Workflow) SetResumeDate(v string) {
	o.ResumeDate = &v
}

// GetScheduledSending returns the ScheduledSending field value if set, zero value otherwise.
func (o *Workflow) GetScheduledSending() ScheduledSending {
	if o == nil || IsNil(o.ScheduledSending) {
		var ret ScheduledSending
		return ret
	}
	return *o.ScheduledSending
}

// GetScheduledSendingOk returns a tuple with the ScheduledSending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetScheduledSendingOk() (*ScheduledSending, bool) {
	if o == nil || IsNil(o.ScheduledSending) {
		return nil, false
	}
	return o.ScheduledSending, true
}

// HasScheduledSending returns a boolean if a field has been set.
func (o *Workflow) HasScheduledSending() bool {
	if o != nil && !IsNil(o.ScheduledSending) {
		return true
	}

	return false
}

// SetScheduledSending gets a reference to the given ScheduledSending and assigns it to the ScheduledSending field.
func (o *Workflow) SetScheduledSending(v ScheduledSending) {
	o.ScheduledSending = &v
}

// GetWorkflowStatus returns the WorkflowStatus field value if set, zero value otherwise.
func (o *Workflow) GetWorkflowStatus() string {
	if o == nil || IsNil(o.WorkflowStatus) {
		var ret string
		return ret
	}
	return *o.WorkflowStatus
}

// GetWorkflowStatusOk returns a tuple with the WorkflowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetWorkflowStatusOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowStatus) {
		return nil, false
	}
	return o.WorkflowStatus, true
}

// HasWorkflowStatus returns a boolean if a field has been set.
func (o *Workflow) HasWorkflowStatus() bool {
	if o != nil && !IsNil(o.WorkflowStatus) {
		return true
	}

	return false
}

// SetWorkflowStatus gets a reference to the given string and assigns it to the WorkflowStatus field.
func (o *Workflow) SetWorkflowStatus(v string) {
	o.WorkflowStatus = &v
}

// GetWorkflowSteps returns the WorkflowSteps field value if set, zero value otherwise.
func (o *Workflow) GetWorkflowSteps() []WorkflowStep {
	if o == nil || IsNil(o.WorkflowSteps) {
		var ret []WorkflowStep
		return ret
	}
	return o.WorkflowSteps
}

// GetWorkflowStepsOk returns a tuple with the WorkflowSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetWorkflowStepsOk() ([]WorkflowStep, bool) {
	if o == nil || IsNil(o.WorkflowSteps) {
		return nil, false
	}
	return o.WorkflowSteps, true
}

// HasWorkflowSteps returns a boolean if a field has been set.
func (o *Workflow) HasWorkflowSteps() bool {
	if o != nil && !IsNil(o.WorkflowSteps) {
		return true
	}

	return false
}

// SetWorkflowSteps gets a reference to the given []WorkflowStep and assigns it to the WorkflowSteps field.
func (o *Workflow) SetWorkflowSteps(v []WorkflowStep) {
	o.WorkflowSteps = v
}

func (o Workflow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Workflow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentWorkflowStepId) {
		toSerialize["currentWorkflowStepId"] = o.CurrentWorkflowStepId
	}
	if !IsNil(o.ResumeDate) {
		toSerialize["resumeDate"] = o.ResumeDate
	}
	if !IsNil(o.ScheduledSending) {
		toSerialize["scheduledSending"] = o.ScheduledSending
	}
	if !IsNil(o.WorkflowStatus) {
		toSerialize["workflowStatus"] = o.WorkflowStatus
	}
	if !IsNil(o.WorkflowSteps) {
		toSerialize["workflowSteps"] = o.WorkflowSteps
	}
	return toSerialize, nil
}

type NullableWorkflow struct {
	value *Workflow
	isSet bool
}

func (v NullableWorkflow) Get() *Workflow {
	return v.value
}

func (v *NullableWorkflow) Set(val *Workflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow(val *Workflow) *NullableWorkflow {
	return &NullableWorkflow{value: val, isSet: true}
}

func (v NullableWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


