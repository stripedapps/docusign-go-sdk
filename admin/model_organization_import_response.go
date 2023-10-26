/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the OrganizationImportResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationImportResponse{}

// OrganizationImportResponse 
type OrganizationImportResponse struct {
	// 
	Id *string `json:"id,omitempty"`
	// 
	Type *string `json:"type,omitempty"`
	Requestor *OrganizationImportResponseRequestor `json:"requestor,omitempty"`
	// 
	Created *time.Time `json:"created,omitempty"`
	// 
	LastModified *time.Time `json:"last_modified,omitempty"`
	// Status.
	Status *string `json:"status,omitempty"`
	// 
	UserCount *int32 `json:"user_count,omitempty"`
	// 
	ProcessedUserCount *int32 `json:"processed_user_count,omitempty"`
	// 
	AddedUserCount *int32 `json:"added_user_count,omitempty"`
	// 
	UpdatedUserCount *int32 `json:"updated_user_count,omitempty"`
	// 
	ClosedUserCount *int32 `json:"closed_user_count,omitempty"`
	// 
	NoActionRequiredUserCount *int32 `json:"no_action_required_user_count,omitempty"`
	// 
	ErrorCount *int32 `json:"error_count,omitempty"`
	// 
	WarningCount *int32 `json:"warning_count,omitempty"`
	// 
	InvalidColumnHeaders *string `json:"invalid_column_headers,omitempty"`
	// 
	ImportsNotFoundOrNotAvailableForAccounts *string `json:"imports_not_found_or_not_available_for_accounts,omitempty"`
	// 
	ImportsFailedForAccounts *string `json:"imports_failed_for_accounts,omitempty"`
	// 
	ImportsTimedOutForAccounts *string `json:"imports_timed_out_for_accounts,omitempty"`
	// 
	ImportsNotFoundOrNotAvailableForSites *string `json:"imports_not_found_or_not_available_for_sites,omitempty"`
	// 
	ImportsFailedForSites *string `json:"imports_failed_for_sites,omitempty"`
	// 
	ImportsTimedOutForSites *string `json:"imports_timed_out_for_sites,omitempty"`
	// 
	FileLevelErrorRollups []OrganizationImportResponseErrorRollup `json:"file_level_error_rollups,omitempty"`
	// 
	UserLevelErrorRollups []OrganizationImportResponseErrorRollup `json:"user_level_error_rollups,omitempty"`
	// 
	UserLevelWarningRollups []OrganizationImportResponseWarningRollup `json:"user_level_warning_rollups,omitempty"`
	// 
	HasCsvResults *bool `json:"has_csv_results,omitempty"`
	// 
	ResultsUri *string `json:"results_uri,omitempty"`
}

// NewOrganizationImportResponse instantiates a new OrganizationImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationImportResponse() *OrganizationImportResponse {
	this := OrganizationImportResponse{}
	return &this
}

// NewOrganizationImportResponseWithDefaults instantiates a new OrganizationImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationImportResponseWithDefaults() *OrganizationImportResponse {
	this := OrganizationImportResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationImportResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrganizationImportResponse) SetType(v string) {
	o.Type = &v
}

// GetRequestor returns the Requestor field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetRequestor() OrganizationImportResponseRequestor {
	if o == nil || IsNil(o.Requestor) {
		var ret OrganizationImportResponseRequestor
		return ret
	}
	return *o.Requestor
}

// GetRequestorOk returns a tuple with the Requestor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetRequestorOk() (*OrganizationImportResponseRequestor, bool) {
	if o == nil || IsNil(o.Requestor) {
		return nil, false
	}
	return o.Requestor, true
}

// HasRequestor returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasRequestor() bool {
	if o != nil && !IsNil(o.Requestor) {
		return true
	}

	return false
}

// SetRequestor gets a reference to the given OrganizationImportResponseRequestor and assigns it to the Requestor field.
func (o *OrganizationImportResponse) SetRequestor(v OrganizationImportResponseRequestor) {
	o.Requestor = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *OrganizationImportResponse) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetLastModified() time.Time {
	if o == nil || IsNil(o.LastModified) {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModified) {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasLastModified() bool {
	if o != nil && !IsNil(o.LastModified) {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *OrganizationImportResponse) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *OrganizationImportResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetUserCount() int32 {
	if o == nil || IsNil(o.UserCount) {
		var ret int32
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UserCount) {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasUserCount() bool {
	if o != nil && !IsNil(o.UserCount) {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int32 and assigns it to the UserCount field.
func (o *OrganizationImportResponse) SetUserCount(v int32) {
	o.UserCount = &v
}

// GetProcessedUserCount returns the ProcessedUserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetProcessedUserCount() int32 {
	if o == nil || IsNil(o.ProcessedUserCount) {
		var ret int32
		return ret
	}
	return *o.ProcessedUserCount
}

// GetProcessedUserCountOk returns a tuple with the ProcessedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetProcessedUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ProcessedUserCount) {
		return nil, false
	}
	return o.ProcessedUserCount, true
}

// HasProcessedUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasProcessedUserCount() bool {
	if o != nil && !IsNil(o.ProcessedUserCount) {
		return true
	}

	return false
}

// SetProcessedUserCount gets a reference to the given int32 and assigns it to the ProcessedUserCount field.
func (o *OrganizationImportResponse) SetProcessedUserCount(v int32) {
	o.ProcessedUserCount = &v
}

// GetAddedUserCount returns the AddedUserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetAddedUserCount() int32 {
	if o == nil || IsNil(o.AddedUserCount) {
		var ret int32
		return ret
	}
	return *o.AddedUserCount
}

// GetAddedUserCountOk returns a tuple with the AddedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetAddedUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AddedUserCount) {
		return nil, false
	}
	return o.AddedUserCount, true
}

// HasAddedUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasAddedUserCount() bool {
	if o != nil && !IsNil(o.AddedUserCount) {
		return true
	}

	return false
}

// SetAddedUserCount gets a reference to the given int32 and assigns it to the AddedUserCount field.
func (o *OrganizationImportResponse) SetAddedUserCount(v int32) {
	o.AddedUserCount = &v
}

// GetUpdatedUserCount returns the UpdatedUserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetUpdatedUserCount() int32 {
	if o == nil || IsNil(o.UpdatedUserCount) {
		var ret int32
		return ret
	}
	return *o.UpdatedUserCount
}

// GetUpdatedUserCountOk returns a tuple with the UpdatedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetUpdatedUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedUserCount) {
		return nil, false
	}
	return o.UpdatedUserCount, true
}

// HasUpdatedUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasUpdatedUserCount() bool {
	if o != nil && !IsNil(o.UpdatedUserCount) {
		return true
	}

	return false
}

// SetUpdatedUserCount gets a reference to the given int32 and assigns it to the UpdatedUserCount field.
func (o *OrganizationImportResponse) SetUpdatedUserCount(v int32) {
	o.UpdatedUserCount = &v
}

// GetClosedUserCount returns the ClosedUserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetClosedUserCount() int32 {
	if o == nil || IsNil(o.ClosedUserCount) {
		var ret int32
		return ret
	}
	return *o.ClosedUserCount
}

// GetClosedUserCountOk returns a tuple with the ClosedUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetClosedUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ClosedUserCount) {
		return nil, false
	}
	return o.ClosedUserCount, true
}

// HasClosedUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasClosedUserCount() bool {
	if o != nil && !IsNil(o.ClosedUserCount) {
		return true
	}

	return false
}

// SetClosedUserCount gets a reference to the given int32 and assigns it to the ClosedUserCount field.
func (o *OrganizationImportResponse) SetClosedUserCount(v int32) {
	o.ClosedUserCount = &v
}

// GetNoActionRequiredUserCount returns the NoActionRequiredUserCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetNoActionRequiredUserCount() int32 {
	if o == nil || IsNil(o.NoActionRequiredUserCount) {
		var ret int32
		return ret
	}
	return *o.NoActionRequiredUserCount
}

// GetNoActionRequiredUserCountOk returns a tuple with the NoActionRequiredUserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetNoActionRequiredUserCountOk() (*int32, bool) {
	if o == nil || IsNil(o.NoActionRequiredUserCount) {
		return nil, false
	}
	return o.NoActionRequiredUserCount, true
}

// HasNoActionRequiredUserCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasNoActionRequiredUserCount() bool {
	if o != nil && !IsNil(o.NoActionRequiredUserCount) {
		return true
	}

	return false
}

// SetNoActionRequiredUserCount gets a reference to the given int32 and assigns it to the NoActionRequiredUserCount field.
func (o *OrganizationImportResponse) SetNoActionRequiredUserCount(v int32) {
	o.NoActionRequiredUserCount = &v
}

// GetErrorCount returns the ErrorCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetErrorCount() int32 {
	if o == nil || IsNil(o.ErrorCount) {
		var ret int32
		return ret
	}
	return *o.ErrorCount
}

// GetErrorCountOk returns a tuple with the ErrorCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetErrorCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCount) {
		return nil, false
	}
	return o.ErrorCount, true
}

// HasErrorCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasErrorCount() bool {
	if o != nil && !IsNil(o.ErrorCount) {
		return true
	}

	return false
}

// SetErrorCount gets a reference to the given int32 and assigns it to the ErrorCount field.
func (o *OrganizationImportResponse) SetErrorCount(v int32) {
	o.ErrorCount = &v
}

// GetWarningCount returns the WarningCount field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetWarningCount() int32 {
	if o == nil || IsNil(o.WarningCount) {
		var ret int32
		return ret
	}
	return *o.WarningCount
}

// GetWarningCountOk returns a tuple with the WarningCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetWarningCountOk() (*int32, bool) {
	if o == nil || IsNil(o.WarningCount) {
		return nil, false
	}
	return o.WarningCount, true
}

// HasWarningCount returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasWarningCount() bool {
	if o != nil && !IsNil(o.WarningCount) {
		return true
	}

	return false
}

// SetWarningCount gets a reference to the given int32 and assigns it to the WarningCount field.
func (o *OrganizationImportResponse) SetWarningCount(v int32) {
	o.WarningCount = &v
}

// GetInvalidColumnHeaders returns the InvalidColumnHeaders field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetInvalidColumnHeaders() string {
	if o == nil || IsNil(o.InvalidColumnHeaders) {
		var ret string
		return ret
	}
	return *o.InvalidColumnHeaders
}

// GetInvalidColumnHeadersOk returns a tuple with the InvalidColumnHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetInvalidColumnHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.InvalidColumnHeaders) {
		return nil, false
	}
	return o.InvalidColumnHeaders, true
}

// HasInvalidColumnHeaders returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasInvalidColumnHeaders() bool {
	if o != nil && !IsNil(o.InvalidColumnHeaders) {
		return true
	}

	return false
}

// SetInvalidColumnHeaders gets a reference to the given string and assigns it to the InvalidColumnHeaders field.
func (o *OrganizationImportResponse) SetInvalidColumnHeaders(v string) {
	o.InvalidColumnHeaders = &v
}

// GetImportsNotFoundOrNotAvailableForAccounts returns the ImportsNotFoundOrNotAvailableForAccounts field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsNotFoundOrNotAvailableForAccounts() string {
	if o == nil || IsNil(o.ImportsNotFoundOrNotAvailableForAccounts) {
		var ret string
		return ret
	}
	return *o.ImportsNotFoundOrNotAvailableForAccounts
}

// GetImportsNotFoundOrNotAvailableForAccountsOk returns a tuple with the ImportsNotFoundOrNotAvailableForAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsNotFoundOrNotAvailableForAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsNotFoundOrNotAvailableForAccounts) {
		return nil, false
	}
	return o.ImportsNotFoundOrNotAvailableForAccounts, true
}

// HasImportsNotFoundOrNotAvailableForAccounts returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsNotFoundOrNotAvailableForAccounts() bool {
	if o != nil && !IsNil(o.ImportsNotFoundOrNotAvailableForAccounts) {
		return true
	}

	return false
}

// SetImportsNotFoundOrNotAvailableForAccounts gets a reference to the given string and assigns it to the ImportsNotFoundOrNotAvailableForAccounts field.
func (o *OrganizationImportResponse) SetImportsNotFoundOrNotAvailableForAccounts(v string) {
	o.ImportsNotFoundOrNotAvailableForAccounts = &v
}

// GetImportsFailedForAccounts returns the ImportsFailedForAccounts field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsFailedForAccounts() string {
	if o == nil || IsNil(o.ImportsFailedForAccounts) {
		var ret string
		return ret
	}
	return *o.ImportsFailedForAccounts
}

// GetImportsFailedForAccountsOk returns a tuple with the ImportsFailedForAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsFailedForAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsFailedForAccounts) {
		return nil, false
	}
	return o.ImportsFailedForAccounts, true
}

// HasImportsFailedForAccounts returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsFailedForAccounts() bool {
	if o != nil && !IsNil(o.ImportsFailedForAccounts) {
		return true
	}

	return false
}

// SetImportsFailedForAccounts gets a reference to the given string and assigns it to the ImportsFailedForAccounts field.
func (o *OrganizationImportResponse) SetImportsFailedForAccounts(v string) {
	o.ImportsFailedForAccounts = &v
}

// GetImportsTimedOutForAccounts returns the ImportsTimedOutForAccounts field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsTimedOutForAccounts() string {
	if o == nil || IsNil(o.ImportsTimedOutForAccounts) {
		var ret string
		return ret
	}
	return *o.ImportsTimedOutForAccounts
}

// GetImportsTimedOutForAccountsOk returns a tuple with the ImportsTimedOutForAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsTimedOutForAccountsOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsTimedOutForAccounts) {
		return nil, false
	}
	return o.ImportsTimedOutForAccounts, true
}

// HasImportsTimedOutForAccounts returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsTimedOutForAccounts() bool {
	if o != nil && !IsNil(o.ImportsTimedOutForAccounts) {
		return true
	}

	return false
}

// SetImportsTimedOutForAccounts gets a reference to the given string and assigns it to the ImportsTimedOutForAccounts field.
func (o *OrganizationImportResponse) SetImportsTimedOutForAccounts(v string) {
	o.ImportsTimedOutForAccounts = &v
}

// GetImportsNotFoundOrNotAvailableForSites returns the ImportsNotFoundOrNotAvailableForSites field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsNotFoundOrNotAvailableForSites() string {
	if o == nil || IsNil(o.ImportsNotFoundOrNotAvailableForSites) {
		var ret string
		return ret
	}
	return *o.ImportsNotFoundOrNotAvailableForSites
}

// GetImportsNotFoundOrNotAvailableForSitesOk returns a tuple with the ImportsNotFoundOrNotAvailableForSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsNotFoundOrNotAvailableForSitesOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsNotFoundOrNotAvailableForSites) {
		return nil, false
	}
	return o.ImportsNotFoundOrNotAvailableForSites, true
}

// HasImportsNotFoundOrNotAvailableForSites returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsNotFoundOrNotAvailableForSites() bool {
	if o != nil && !IsNil(o.ImportsNotFoundOrNotAvailableForSites) {
		return true
	}

	return false
}

// SetImportsNotFoundOrNotAvailableForSites gets a reference to the given string and assigns it to the ImportsNotFoundOrNotAvailableForSites field.
func (o *OrganizationImportResponse) SetImportsNotFoundOrNotAvailableForSites(v string) {
	o.ImportsNotFoundOrNotAvailableForSites = &v
}

// GetImportsFailedForSites returns the ImportsFailedForSites field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsFailedForSites() string {
	if o == nil || IsNil(o.ImportsFailedForSites) {
		var ret string
		return ret
	}
	return *o.ImportsFailedForSites
}

// GetImportsFailedForSitesOk returns a tuple with the ImportsFailedForSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsFailedForSitesOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsFailedForSites) {
		return nil, false
	}
	return o.ImportsFailedForSites, true
}

// HasImportsFailedForSites returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsFailedForSites() bool {
	if o != nil && !IsNil(o.ImportsFailedForSites) {
		return true
	}

	return false
}

// SetImportsFailedForSites gets a reference to the given string and assigns it to the ImportsFailedForSites field.
func (o *OrganizationImportResponse) SetImportsFailedForSites(v string) {
	o.ImportsFailedForSites = &v
}

// GetImportsTimedOutForSites returns the ImportsTimedOutForSites field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetImportsTimedOutForSites() string {
	if o == nil || IsNil(o.ImportsTimedOutForSites) {
		var ret string
		return ret
	}
	return *o.ImportsTimedOutForSites
}

// GetImportsTimedOutForSitesOk returns a tuple with the ImportsTimedOutForSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetImportsTimedOutForSitesOk() (*string, bool) {
	if o == nil || IsNil(o.ImportsTimedOutForSites) {
		return nil, false
	}
	return o.ImportsTimedOutForSites, true
}

// HasImportsTimedOutForSites returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasImportsTimedOutForSites() bool {
	if o != nil && !IsNil(o.ImportsTimedOutForSites) {
		return true
	}

	return false
}

// SetImportsTimedOutForSites gets a reference to the given string and assigns it to the ImportsTimedOutForSites field.
func (o *OrganizationImportResponse) SetImportsTimedOutForSites(v string) {
	o.ImportsTimedOutForSites = &v
}

// GetFileLevelErrorRollups returns the FileLevelErrorRollups field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetFileLevelErrorRollups() []OrganizationImportResponseErrorRollup {
	if o == nil || IsNil(o.FileLevelErrorRollups) {
		var ret []OrganizationImportResponseErrorRollup
		return ret
	}
	return o.FileLevelErrorRollups
}

// GetFileLevelErrorRollupsOk returns a tuple with the FileLevelErrorRollups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetFileLevelErrorRollupsOk() ([]OrganizationImportResponseErrorRollup, bool) {
	if o == nil || IsNil(o.FileLevelErrorRollups) {
		return nil, false
	}
	return o.FileLevelErrorRollups, true
}

// HasFileLevelErrorRollups returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasFileLevelErrorRollups() bool {
	if o != nil && !IsNil(o.FileLevelErrorRollups) {
		return true
	}

	return false
}

// SetFileLevelErrorRollups gets a reference to the given []OrganizationImportResponseErrorRollup and assigns it to the FileLevelErrorRollups field.
func (o *OrganizationImportResponse) SetFileLevelErrorRollups(v []OrganizationImportResponseErrorRollup) {
	o.FileLevelErrorRollups = v
}

// GetUserLevelErrorRollups returns the UserLevelErrorRollups field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetUserLevelErrorRollups() []OrganizationImportResponseErrorRollup {
	if o == nil || IsNil(o.UserLevelErrorRollups) {
		var ret []OrganizationImportResponseErrorRollup
		return ret
	}
	return o.UserLevelErrorRollups
}

// GetUserLevelErrorRollupsOk returns a tuple with the UserLevelErrorRollups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetUserLevelErrorRollupsOk() ([]OrganizationImportResponseErrorRollup, bool) {
	if o == nil || IsNil(o.UserLevelErrorRollups) {
		return nil, false
	}
	return o.UserLevelErrorRollups, true
}

// HasUserLevelErrorRollups returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasUserLevelErrorRollups() bool {
	if o != nil && !IsNil(o.UserLevelErrorRollups) {
		return true
	}

	return false
}

// SetUserLevelErrorRollups gets a reference to the given []OrganizationImportResponseErrorRollup and assigns it to the UserLevelErrorRollups field.
func (o *OrganizationImportResponse) SetUserLevelErrorRollups(v []OrganizationImportResponseErrorRollup) {
	o.UserLevelErrorRollups = v
}

// GetUserLevelWarningRollups returns the UserLevelWarningRollups field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetUserLevelWarningRollups() []OrganizationImportResponseWarningRollup {
	if o == nil || IsNil(o.UserLevelWarningRollups) {
		var ret []OrganizationImportResponseWarningRollup
		return ret
	}
	return o.UserLevelWarningRollups
}

// GetUserLevelWarningRollupsOk returns a tuple with the UserLevelWarningRollups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetUserLevelWarningRollupsOk() ([]OrganizationImportResponseWarningRollup, bool) {
	if o == nil || IsNil(o.UserLevelWarningRollups) {
		return nil, false
	}
	return o.UserLevelWarningRollups, true
}

// HasUserLevelWarningRollups returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasUserLevelWarningRollups() bool {
	if o != nil && !IsNil(o.UserLevelWarningRollups) {
		return true
	}

	return false
}

// SetUserLevelWarningRollups gets a reference to the given []OrganizationImportResponseWarningRollup and assigns it to the UserLevelWarningRollups field.
func (o *OrganizationImportResponse) SetUserLevelWarningRollups(v []OrganizationImportResponseWarningRollup) {
	o.UserLevelWarningRollups = v
}

// GetHasCsvResults returns the HasCsvResults field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetHasCsvResults() bool {
	if o == nil || IsNil(o.HasCsvResults) {
		var ret bool
		return ret
	}
	return *o.HasCsvResults
}

// GetHasCsvResultsOk returns a tuple with the HasCsvResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetHasCsvResultsOk() (*bool, bool) {
	if o == nil || IsNil(o.HasCsvResults) {
		return nil, false
	}
	return o.HasCsvResults, true
}

// HasHasCsvResults returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasHasCsvResults() bool {
	if o != nil && !IsNil(o.HasCsvResults) {
		return true
	}

	return false
}

// SetHasCsvResults gets a reference to the given bool and assigns it to the HasCsvResults field.
func (o *OrganizationImportResponse) SetHasCsvResults(v bool) {
	o.HasCsvResults = &v
}

// GetResultsUri returns the ResultsUri field value if set, zero value otherwise.
func (o *OrganizationImportResponse) GetResultsUri() string {
	if o == nil || IsNil(o.ResultsUri) {
		var ret string
		return ret
	}
	return *o.ResultsUri
}

// GetResultsUriOk returns a tuple with the ResultsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationImportResponse) GetResultsUriOk() (*string, bool) {
	if o == nil || IsNil(o.ResultsUri) {
		return nil, false
	}
	return o.ResultsUri, true
}

// HasResultsUri returns a boolean if a field has been set.
func (o *OrganizationImportResponse) HasResultsUri() bool {
	if o != nil && !IsNil(o.ResultsUri) {
		return true
	}

	return false
}

// SetResultsUri gets a reference to the given string and assigns it to the ResultsUri field.
func (o *OrganizationImportResponse) SetResultsUri(v string) {
	o.ResultsUri = &v
}

func (o OrganizationImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationImportResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Requestor) {
		toSerialize["requestor"] = o.Requestor
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.LastModified) {
		toSerialize["last_modified"] = o.LastModified
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UserCount) {
		toSerialize["user_count"] = o.UserCount
	}
	if !IsNil(o.ProcessedUserCount) {
		toSerialize["processed_user_count"] = o.ProcessedUserCount
	}
	if !IsNil(o.AddedUserCount) {
		toSerialize["added_user_count"] = o.AddedUserCount
	}
	if !IsNil(o.UpdatedUserCount) {
		toSerialize["updated_user_count"] = o.UpdatedUserCount
	}
	if !IsNil(o.ClosedUserCount) {
		toSerialize["closed_user_count"] = o.ClosedUserCount
	}
	if !IsNil(o.NoActionRequiredUserCount) {
		toSerialize["no_action_required_user_count"] = o.NoActionRequiredUserCount
	}
	if !IsNil(o.ErrorCount) {
		toSerialize["error_count"] = o.ErrorCount
	}
	if !IsNil(o.WarningCount) {
		toSerialize["warning_count"] = o.WarningCount
	}
	if !IsNil(o.InvalidColumnHeaders) {
		toSerialize["invalid_column_headers"] = o.InvalidColumnHeaders
	}
	if !IsNil(o.ImportsNotFoundOrNotAvailableForAccounts) {
		toSerialize["imports_not_found_or_not_available_for_accounts"] = o.ImportsNotFoundOrNotAvailableForAccounts
	}
	if !IsNil(o.ImportsFailedForAccounts) {
		toSerialize["imports_failed_for_accounts"] = o.ImportsFailedForAccounts
	}
	if !IsNil(o.ImportsTimedOutForAccounts) {
		toSerialize["imports_timed_out_for_accounts"] = o.ImportsTimedOutForAccounts
	}
	if !IsNil(o.ImportsNotFoundOrNotAvailableForSites) {
		toSerialize["imports_not_found_or_not_available_for_sites"] = o.ImportsNotFoundOrNotAvailableForSites
	}
	if !IsNil(o.ImportsFailedForSites) {
		toSerialize["imports_failed_for_sites"] = o.ImportsFailedForSites
	}
	if !IsNil(o.ImportsTimedOutForSites) {
		toSerialize["imports_timed_out_for_sites"] = o.ImportsTimedOutForSites
	}
	if !IsNil(o.FileLevelErrorRollups) {
		toSerialize["file_level_error_rollups"] = o.FileLevelErrorRollups
	}
	if !IsNil(o.UserLevelErrorRollups) {
		toSerialize["user_level_error_rollups"] = o.UserLevelErrorRollups
	}
	if !IsNil(o.UserLevelWarningRollups) {
		toSerialize["user_level_warning_rollups"] = o.UserLevelWarningRollups
	}
	if !IsNil(o.HasCsvResults) {
		toSerialize["has_csv_results"] = o.HasCsvResults
	}
	if !IsNil(o.ResultsUri) {
		toSerialize["results_uri"] = o.ResultsUri
	}
	return toSerialize, nil
}

type NullableOrganizationImportResponse struct {
	value *OrganizationImportResponse
	isSet bool
}

func (v NullableOrganizationImportResponse) Get() *OrganizationImportResponse {
	return v.value
}

func (v *NullableOrganizationImportResponse) Set(val *OrganizationImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationImportResponse(val *OrganizationImportResponse) *NullableOrganizationImportResponse {
	return &NullableOrganizationImportResponse{value: val, isSet: true}
}

func (v NullableOrganizationImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


