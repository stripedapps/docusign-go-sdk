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

// checks if the EnvelopeUpdateSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeUpdateSummary{}

// EnvelopeUpdateSummary 
type EnvelopeUpdateSummary struct {
	BulkEnvelopeStatus *BulkEnvelopeStatus `json:"bulkEnvelopeStatus,omitempty"`
	// The envelope ID of the envelope status that failed to post.
	EnvelopeId *string `json:"envelopeId,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// 
	ListCustomFieldUpdateResults []ListCustomField `json:"listCustomFieldUpdateResults,omitempty"`
	LockInformation *EnvelopeLocks `json:"lockInformation,omitempty"`
	// Shows the current purge state for the envelope. Valid values:  - `unpurged`: There has been no successful request to purge documents. - `documents_queued`: The envelope documents have been added to the purge queue, but have not been purged. - `documents_dequeued`: The envelope documents have been taken out of the purge queue. - `documents_purged`: The envelope documents have been successfully purged. - `documents_and_metadata_queued`: The envelope documents and metadata have been added to the purge queue, but have not yet been purged. - `documents_and_metadata_purged`: The envelope documents and metadata have been successfully purged. - `documents_and_metadata_and_redact_queued`: The envelope documents and metadata have been added to the purge queue, but have not yet been purged, nor has personal information been redacted. - `documents_and_metadata_and_redact_purged`: The envelope documents and metadata have been successfully purged, and personal information has been redacted.  **Related topics**  - [Purging documents (eSingature Concepts)](/docs/esign-rest-api/esign101/concepts/documents/purging/) - [Purging documents in an envelope (blog post)](https://www.docusign.com/blog/developers/purging-documents-envelope)  
	PurgeState *string `json:"purgeState,omitempty"`
	// An array of `recipientUpdateResults` objects that contain details about the recipients.
	RecipientUpdateResults []RecipientUpdateResponse `json:"recipientUpdateResults,omitempty"`
	TabUpdateResults *EnvelopeRecipientTabs `json:"tabUpdateResults,omitempty"`
	// 
	TextCustomFieldUpdateResults []TextCustomField `json:"textCustomFieldUpdateResults,omitempty"`
}

// NewEnvelopeUpdateSummary instantiates a new EnvelopeUpdateSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeUpdateSummary() *EnvelopeUpdateSummary {
	this := EnvelopeUpdateSummary{}
	return &this
}

// NewEnvelopeUpdateSummaryWithDefaults instantiates a new EnvelopeUpdateSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeUpdateSummaryWithDefaults() *EnvelopeUpdateSummary {
	this := EnvelopeUpdateSummary{}
	return &this
}

// GetBulkEnvelopeStatus returns the BulkEnvelopeStatus field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetBulkEnvelopeStatus() BulkEnvelopeStatus {
	if o == nil || IsNil(o.BulkEnvelopeStatus) {
		var ret BulkEnvelopeStatus
		return ret
	}
	return *o.BulkEnvelopeStatus
}

// GetBulkEnvelopeStatusOk returns a tuple with the BulkEnvelopeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetBulkEnvelopeStatusOk() (*BulkEnvelopeStatus, bool) {
	if o == nil || IsNil(o.BulkEnvelopeStatus) {
		return nil, false
	}
	return o.BulkEnvelopeStatus, true
}

// HasBulkEnvelopeStatus returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasBulkEnvelopeStatus() bool {
	if o != nil && !IsNil(o.BulkEnvelopeStatus) {
		return true
	}

	return false
}

// SetBulkEnvelopeStatus gets a reference to the given BulkEnvelopeStatus and assigns it to the BulkEnvelopeStatus field.
func (o *EnvelopeUpdateSummary) SetBulkEnvelopeStatus(v BulkEnvelopeStatus) {
	o.BulkEnvelopeStatus = &v
}

// GetEnvelopeId returns the EnvelopeId field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetEnvelopeId() string {
	if o == nil || IsNil(o.EnvelopeId) {
		var ret string
		return ret
	}
	return *o.EnvelopeId
}

// GetEnvelopeIdOk returns a tuple with the EnvelopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetEnvelopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeId) {
		return nil, false
	}
	return o.EnvelopeId, true
}

// HasEnvelopeId returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasEnvelopeId() bool {
	if o != nil && !IsNil(o.EnvelopeId) {
		return true
	}

	return false
}

// SetEnvelopeId gets a reference to the given string and assigns it to the EnvelopeId field.
func (o *EnvelopeUpdateSummary) SetEnvelopeId(v string) {
	o.EnvelopeId = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *EnvelopeUpdateSummary) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetListCustomFieldUpdateResults returns the ListCustomFieldUpdateResults field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetListCustomFieldUpdateResults() []ListCustomField {
	if o == nil || IsNil(o.ListCustomFieldUpdateResults) {
		var ret []ListCustomField
		return ret
	}
	return o.ListCustomFieldUpdateResults
}

// GetListCustomFieldUpdateResultsOk returns a tuple with the ListCustomFieldUpdateResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetListCustomFieldUpdateResultsOk() ([]ListCustomField, bool) {
	if o == nil || IsNil(o.ListCustomFieldUpdateResults) {
		return nil, false
	}
	return o.ListCustomFieldUpdateResults, true
}

// HasListCustomFieldUpdateResults returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasListCustomFieldUpdateResults() bool {
	if o != nil && !IsNil(o.ListCustomFieldUpdateResults) {
		return true
	}

	return false
}

// SetListCustomFieldUpdateResults gets a reference to the given []ListCustomField and assigns it to the ListCustomFieldUpdateResults field.
func (o *EnvelopeUpdateSummary) SetListCustomFieldUpdateResults(v []ListCustomField) {
	o.ListCustomFieldUpdateResults = v
}

// GetLockInformation returns the LockInformation field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetLockInformation() EnvelopeLocks {
	if o == nil || IsNil(o.LockInformation) {
		var ret EnvelopeLocks
		return ret
	}
	return *o.LockInformation
}

// GetLockInformationOk returns a tuple with the LockInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetLockInformationOk() (*EnvelopeLocks, bool) {
	if o == nil || IsNil(o.LockInformation) {
		return nil, false
	}
	return o.LockInformation, true
}

// HasLockInformation returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasLockInformation() bool {
	if o != nil && !IsNil(o.LockInformation) {
		return true
	}

	return false
}

// SetLockInformation gets a reference to the given EnvelopeLocks and assigns it to the LockInformation field.
func (o *EnvelopeUpdateSummary) SetLockInformation(v EnvelopeLocks) {
	o.LockInformation = &v
}

// GetPurgeState returns the PurgeState field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetPurgeState() string {
	if o == nil || IsNil(o.PurgeState) {
		var ret string
		return ret
	}
	return *o.PurgeState
}

// GetPurgeStateOk returns a tuple with the PurgeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetPurgeStateOk() (*string, bool) {
	if o == nil || IsNil(o.PurgeState) {
		return nil, false
	}
	return o.PurgeState, true
}

// HasPurgeState returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasPurgeState() bool {
	if o != nil && !IsNil(o.PurgeState) {
		return true
	}

	return false
}

// SetPurgeState gets a reference to the given string and assigns it to the PurgeState field.
func (o *EnvelopeUpdateSummary) SetPurgeState(v string) {
	o.PurgeState = &v
}

// GetRecipientUpdateResults returns the RecipientUpdateResults field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetRecipientUpdateResults() []RecipientUpdateResponse {
	if o == nil || IsNil(o.RecipientUpdateResults) {
		var ret []RecipientUpdateResponse
		return ret
	}
	return o.RecipientUpdateResults
}

// GetRecipientUpdateResultsOk returns a tuple with the RecipientUpdateResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetRecipientUpdateResultsOk() ([]RecipientUpdateResponse, bool) {
	if o == nil || IsNil(o.RecipientUpdateResults) {
		return nil, false
	}
	return o.RecipientUpdateResults, true
}

// HasRecipientUpdateResults returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasRecipientUpdateResults() bool {
	if o != nil && !IsNil(o.RecipientUpdateResults) {
		return true
	}

	return false
}

// SetRecipientUpdateResults gets a reference to the given []RecipientUpdateResponse and assigns it to the RecipientUpdateResults field.
func (o *EnvelopeUpdateSummary) SetRecipientUpdateResults(v []RecipientUpdateResponse) {
	o.RecipientUpdateResults = v
}

// GetTabUpdateResults returns the TabUpdateResults field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetTabUpdateResults() EnvelopeRecipientTabs {
	if o == nil || IsNil(o.TabUpdateResults) {
		var ret EnvelopeRecipientTabs
		return ret
	}
	return *o.TabUpdateResults
}

// GetTabUpdateResultsOk returns a tuple with the TabUpdateResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetTabUpdateResultsOk() (*EnvelopeRecipientTabs, bool) {
	if o == nil || IsNil(o.TabUpdateResults) {
		return nil, false
	}
	return o.TabUpdateResults, true
}

// HasTabUpdateResults returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasTabUpdateResults() bool {
	if o != nil && !IsNil(o.TabUpdateResults) {
		return true
	}

	return false
}

// SetTabUpdateResults gets a reference to the given EnvelopeRecipientTabs and assigns it to the TabUpdateResults field.
func (o *EnvelopeUpdateSummary) SetTabUpdateResults(v EnvelopeRecipientTabs) {
	o.TabUpdateResults = &v
}

// GetTextCustomFieldUpdateResults returns the TextCustomFieldUpdateResults field value if set, zero value otherwise.
func (o *EnvelopeUpdateSummary) GetTextCustomFieldUpdateResults() []TextCustomField {
	if o == nil || IsNil(o.TextCustomFieldUpdateResults) {
		var ret []TextCustomField
		return ret
	}
	return o.TextCustomFieldUpdateResults
}

// GetTextCustomFieldUpdateResultsOk returns a tuple with the TextCustomFieldUpdateResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeUpdateSummary) GetTextCustomFieldUpdateResultsOk() ([]TextCustomField, bool) {
	if o == nil || IsNil(o.TextCustomFieldUpdateResults) {
		return nil, false
	}
	return o.TextCustomFieldUpdateResults, true
}

// HasTextCustomFieldUpdateResults returns a boolean if a field has been set.
func (o *EnvelopeUpdateSummary) HasTextCustomFieldUpdateResults() bool {
	if o != nil && !IsNil(o.TextCustomFieldUpdateResults) {
		return true
	}

	return false
}

// SetTextCustomFieldUpdateResults gets a reference to the given []TextCustomField and assigns it to the TextCustomFieldUpdateResults field.
func (o *EnvelopeUpdateSummary) SetTextCustomFieldUpdateResults(v []TextCustomField) {
	o.TextCustomFieldUpdateResults = v
}

func (o EnvelopeUpdateSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeUpdateSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkEnvelopeStatus) {
		toSerialize["bulkEnvelopeStatus"] = o.BulkEnvelopeStatus
	}
	if !IsNil(o.EnvelopeId) {
		toSerialize["envelopeId"] = o.EnvelopeId
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.ListCustomFieldUpdateResults) {
		toSerialize["listCustomFieldUpdateResults"] = o.ListCustomFieldUpdateResults
	}
	if !IsNil(o.LockInformation) {
		toSerialize["lockInformation"] = o.LockInformation
	}
	if !IsNil(o.PurgeState) {
		toSerialize["purgeState"] = o.PurgeState
	}
	if !IsNil(o.RecipientUpdateResults) {
		toSerialize["recipientUpdateResults"] = o.RecipientUpdateResults
	}
	if !IsNil(o.TabUpdateResults) {
		toSerialize["tabUpdateResults"] = o.TabUpdateResults
	}
	if !IsNil(o.TextCustomFieldUpdateResults) {
		toSerialize["textCustomFieldUpdateResults"] = o.TextCustomFieldUpdateResults
	}
	return toSerialize, nil
}

type NullableEnvelopeUpdateSummary struct {
	value *EnvelopeUpdateSummary
	isSet bool
}

func (v NullableEnvelopeUpdateSummary) Get() *EnvelopeUpdateSummary {
	return v.value
}

func (v *NullableEnvelopeUpdateSummary) Set(val *EnvelopeUpdateSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeUpdateSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeUpdateSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeUpdateSummary(val *EnvelopeUpdateSummary) *NullableEnvelopeUpdateSummary {
	return &NullableEnvelopeUpdateSummary{value: val, isSet: true}
}

func (v NullableEnvelopeUpdateSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeUpdateSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


