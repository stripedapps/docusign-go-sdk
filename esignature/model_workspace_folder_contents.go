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

// checks if the WorkspaceFolderContents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceFolderContents{}

// WorkspaceFolderContents This object's properties describe the contents of a workspace folder.
type WorkspaceFolderContents struct {
	// The last index position in the result set. 
	EndPosition *string `json:"endPosition,omitempty"`
	Folder *WorkspaceItem `json:"folder,omitempty"`
	// A list of workspace items.
	Items []WorkspaceItem `json:"items,omitempty"`
	// 
	ParentFolders []WorkspaceItem `json:"parentFolders,omitempty"`
	// The number of results in this response. Because you can filter which entries are included in the response, this value is always less than or equal to the `totalSetSize`.
	ResultSetSize *string `json:"resultSetSize,omitempty"`
	// The starting index position of the current result set.
	StartPosition *string `json:"startPosition,omitempty"`
	// The total number of items in the result set. This value is always greater than or equal to the value of `resultSetSize`.
	TotalSetSize *string `json:"totalSetSize,omitempty"`
	// The ID of the workspace, always populated.
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

// NewWorkspaceFolderContents instantiates a new WorkspaceFolderContents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceFolderContents() *WorkspaceFolderContents {
	this := WorkspaceFolderContents{}
	return &this
}

// NewWorkspaceFolderContentsWithDefaults instantiates a new WorkspaceFolderContents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceFolderContentsWithDefaults() *WorkspaceFolderContents {
	this := WorkspaceFolderContents{}
	return &this
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetEndPosition() string {
	if o == nil || IsNil(o.EndPosition) {
		var ret string
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetEndPositionOk() (*string, bool) {
	if o == nil || IsNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasEndPosition() bool {
	if o != nil && !IsNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given string and assigns it to the EndPosition field.
func (o *WorkspaceFolderContents) SetEndPosition(v string) {
	o.EndPosition = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetFolder() WorkspaceItem {
	if o == nil || IsNil(o.Folder) {
		var ret WorkspaceItem
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetFolderOk() (*WorkspaceItem, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasFolder() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given WorkspaceItem and assigns it to the Folder field.
func (o *WorkspaceFolderContents) SetFolder(v WorkspaceItem) {
	o.Folder = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetItems() []WorkspaceItem {
	if o == nil || IsNil(o.Items) {
		var ret []WorkspaceItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetItemsOk() ([]WorkspaceItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceItem and assigns it to the Items field.
func (o *WorkspaceFolderContents) SetItems(v []WorkspaceItem) {
	o.Items = v
}

// GetParentFolders returns the ParentFolders field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetParentFolders() []WorkspaceItem {
	if o == nil || IsNil(o.ParentFolders) {
		var ret []WorkspaceItem
		return ret
	}
	return o.ParentFolders
}

// GetParentFoldersOk returns a tuple with the ParentFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetParentFoldersOk() ([]WorkspaceItem, bool) {
	if o == nil || IsNil(o.ParentFolders) {
		return nil, false
	}
	return o.ParentFolders, true
}

// HasParentFolders returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasParentFolders() bool {
	if o != nil && !IsNil(o.ParentFolders) {
		return true
	}

	return false
}

// SetParentFolders gets a reference to the given []WorkspaceItem and assigns it to the ParentFolders field.
func (o *WorkspaceFolderContents) SetParentFolders(v []WorkspaceItem) {
	o.ParentFolders = v
}

// GetResultSetSize returns the ResultSetSize field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetResultSetSize() string {
	if o == nil || IsNil(o.ResultSetSize) {
		var ret string
		return ret
	}
	return *o.ResultSetSize
}

// GetResultSetSizeOk returns a tuple with the ResultSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetResultSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultSetSize) {
		return nil, false
	}
	return o.ResultSetSize, true
}

// HasResultSetSize returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasResultSetSize() bool {
	if o != nil && !IsNil(o.ResultSetSize) {
		return true
	}

	return false
}

// SetResultSetSize gets a reference to the given string and assigns it to the ResultSetSize field.
func (o *WorkspaceFolderContents) SetResultSetSize(v string) {
	o.ResultSetSize = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetStartPosition() string {
	if o == nil || IsNil(o.StartPosition) {
		var ret string
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetStartPositionOk() (*string, bool) {
	if o == nil || IsNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasStartPosition() bool {
	if o != nil && !IsNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given string and assigns it to the StartPosition field.
func (o *WorkspaceFolderContents) SetStartPosition(v string) {
	o.StartPosition = &v
}

// GetTotalSetSize returns the TotalSetSize field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetTotalSetSize() string {
	if o == nil || IsNil(o.TotalSetSize) {
		var ret string
		return ret
	}
	return *o.TotalSetSize
}

// GetTotalSetSizeOk returns a tuple with the TotalSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetTotalSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSetSize) {
		return nil, false
	}
	return o.TotalSetSize, true
}

// HasTotalSetSize returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasTotalSetSize() bool {
	if o != nil && !IsNil(o.TotalSetSize) {
		return true
	}

	return false
}

// SetTotalSetSize gets a reference to the given string and assigns it to the TotalSetSize field.
func (o *WorkspaceFolderContents) SetTotalSetSize(v string) {
	o.TotalSetSize = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *WorkspaceFolderContents) GetWorkspaceId() string {
	if o == nil || IsNil(o.WorkspaceId) {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceFolderContents) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkspaceId) {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *WorkspaceFolderContents) HasWorkspaceId() bool {
	if o != nil && !IsNil(o.WorkspaceId) {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *WorkspaceFolderContents) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o WorkspaceFolderContents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceFolderContents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ParentFolders) {
		toSerialize["parentFolders"] = o.ParentFolders
	}
	if !IsNil(o.ResultSetSize) {
		toSerialize["resultSetSize"] = o.ResultSetSize
	}
	if !IsNil(o.StartPosition) {
		toSerialize["startPosition"] = o.StartPosition
	}
	if !IsNil(o.TotalSetSize) {
		toSerialize["totalSetSize"] = o.TotalSetSize
	}
	if !IsNil(o.WorkspaceId) {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return toSerialize, nil
}

type NullableWorkspaceFolderContents struct {
	value *WorkspaceFolderContents
	isSet bool
}

func (v NullableWorkspaceFolderContents) Get() *WorkspaceFolderContents {
	return v.value
}

func (v *NullableWorkspaceFolderContents) Set(val *WorkspaceFolderContents) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceFolderContents) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceFolderContents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceFolderContents(val *WorkspaceFolderContents) *NullableWorkspaceFolderContents {
	return &NullableWorkspaceFolderContents{value: val, isSet: true}
}

func (v NullableWorkspaceFolderContents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceFolderContents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

