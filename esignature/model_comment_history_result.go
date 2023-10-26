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

// checks if the CommentHistoryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentHistoryResult{}

// CommentHistoryResult 
type CommentHistoryResult struct {
	// An array of comment tabs that contain information about users' comments on documents.
	Comments []Comment `json:"comments,omitempty"`
	// The maximum number of results to return.
	Count *int32 `json:"count,omitempty"`
	// 
	EndTimetoken *string `json:"endTimetoken,omitempty"`
	// 
	StartTimetoken *string `json:"startTimetoken,omitempty"`
}

// NewCommentHistoryResult instantiates a new CommentHistoryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentHistoryResult() *CommentHistoryResult {
	this := CommentHistoryResult{}
	return &this
}

// NewCommentHistoryResultWithDefaults instantiates a new CommentHistoryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentHistoryResultWithDefaults() *CommentHistoryResult {
	this := CommentHistoryResult{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *CommentHistoryResult) GetComments() []Comment {
	if o == nil || IsNil(o.Comments) {
		var ret []Comment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentHistoryResult) GetCommentsOk() ([]Comment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *CommentHistoryResult) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given []Comment and assigns it to the Comments field.
func (o *CommentHistoryResult) SetComments(v []Comment) {
	o.Comments = v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CommentHistoryResult) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentHistoryResult) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CommentHistoryResult) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CommentHistoryResult) SetCount(v int32) {
	o.Count = &v
}

// GetEndTimetoken returns the EndTimetoken field value if set, zero value otherwise.
func (o *CommentHistoryResult) GetEndTimetoken() string {
	if o == nil || IsNil(o.EndTimetoken) {
		var ret string
		return ret
	}
	return *o.EndTimetoken
}

// GetEndTimetokenOk returns a tuple with the EndTimetoken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentHistoryResult) GetEndTimetokenOk() (*string, bool) {
	if o == nil || IsNil(o.EndTimetoken) {
		return nil, false
	}
	return o.EndTimetoken, true
}

// HasEndTimetoken returns a boolean if a field has been set.
func (o *CommentHistoryResult) HasEndTimetoken() bool {
	if o != nil && !IsNil(o.EndTimetoken) {
		return true
	}

	return false
}

// SetEndTimetoken gets a reference to the given string and assigns it to the EndTimetoken field.
func (o *CommentHistoryResult) SetEndTimetoken(v string) {
	o.EndTimetoken = &v
}

// GetStartTimetoken returns the StartTimetoken field value if set, zero value otherwise.
func (o *CommentHistoryResult) GetStartTimetoken() string {
	if o == nil || IsNil(o.StartTimetoken) {
		var ret string
		return ret
	}
	return *o.StartTimetoken
}

// GetStartTimetokenOk returns a tuple with the StartTimetoken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentHistoryResult) GetStartTimetokenOk() (*string, bool) {
	if o == nil || IsNil(o.StartTimetoken) {
		return nil, false
	}
	return o.StartTimetoken, true
}

// HasStartTimetoken returns a boolean if a field has been set.
func (o *CommentHistoryResult) HasStartTimetoken() bool {
	if o != nil && !IsNil(o.StartTimetoken) {
		return true
	}

	return false
}

// SetStartTimetoken gets a reference to the given string and assigns it to the StartTimetoken field.
func (o *CommentHistoryResult) SetStartTimetoken(v string) {
	o.StartTimetoken = &v
}

func (o CommentHistoryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentHistoryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.EndTimetoken) {
		toSerialize["endTimetoken"] = o.EndTimetoken
	}
	if !IsNil(o.StartTimetoken) {
		toSerialize["startTimetoken"] = o.StartTimetoken
	}
	return toSerialize, nil
}

type NullableCommentHistoryResult struct {
	value *CommentHistoryResult
	isSet bool
}

func (v NullableCommentHistoryResult) Get() *CommentHistoryResult {
	return v.value
}

func (v *NullableCommentHistoryResult) Set(val *CommentHistoryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentHistoryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentHistoryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentHistoryResult(val *CommentHistoryResult) *NullableCommentHistoryResult {
	return &NullableCommentHistoryResult{value: val, isSet: true}
}

func (v NullableCommentHistoryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentHistoryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

