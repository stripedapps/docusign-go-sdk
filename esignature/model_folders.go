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

// checks if the Folders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Folders{}

// Folders Folders allow you to organize envelopes and templates.
type Folders struct {
	// The last index position in the result set. 
	EndPosition *string `json:"endPosition,omitempty"`
	// A list of envelopes in this folder.
	Envelopes []EnvelopeSummary `json:"envelopes,omitempty"`
	// A list of folder objects.
	Folders []Folder `json:"folders,omitempty"`
	// The URI for the next chunk of records based on the search request. It is `null` if this is the last set of results for the search. 
	NextUri *string `json:"nextUri,omitempty"`
	// The URI for the prior chunk of records based on the search request. It is `null` if this is the first set of results for the search. 
	PreviousUri *string `json:"previousUri,omitempty"`
	// The number of results in this response. Because you can filter which entries are included in the response, this value is always less than or equal to the `totalSetSize`.
	ResultSetSize *string `json:"resultSetSize,omitempty"`
	// The starting index position of the current result set.
	StartPosition *string `json:"startPosition,omitempty"`
	// The total number of items in the result set. This value is always greater than or equal to the value of `resultSetSize`.
	TotalSetSize *string `json:"totalSetSize,omitempty"`
}

// NewFolders instantiates a new Folders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolders() *Folders {
	this := Folders{}
	return &this
}

// NewFoldersWithDefaults instantiates a new Folders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFoldersWithDefaults() *Folders {
	this := Folders{}
	return &this
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *Folders) GetEndPosition() string {
	if o == nil || IsNil(o.EndPosition) {
		var ret string
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetEndPositionOk() (*string, bool) {
	if o == nil || IsNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *Folders) HasEndPosition() bool {
	if o != nil && !IsNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given string and assigns it to the EndPosition field.
func (o *Folders) SetEndPosition(v string) {
	o.EndPosition = &v
}

// GetEnvelopes returns the Envelopes field value if set, zero value otherwise.
func (o *Folders) GetEnvelopes() []EnvelopeSummary {
	if o == nil || IsNil(o.Envelopes) {
		var ret []EnvelopeSummary
		return ret
	}
	return o.Envelopes
}

// GetEnvelopesOk returns a tuple with the Envelopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetEnvelopesOk() ([]EnvelopeSummary, bool) {
	if o == nil || IsNil(o.Envelopes) {
		return nil, false
	}
	return o.Envelopes, true
}

// HasEnvelopes returns a boolean if a field has been set.
func (o *Folders) HasEnvelopes() bool {
	if o != nil && !IsNil(o.Envelopes) {
		return true
	}

	return false
}

// SetEnvelopes gets a reference to the given []EnvelopeSummary and assigns it to the Envelopes field.
func (o *Folders) SetEnvelopes(v []EnvelopeSummary) {
	o.Envelopes = v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *Folders) GetFolders() []Folder {
	if o == nil || IsNil(o.Folders) {
		var ret []Folder
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetFoldersOk() ([]Folder, bool) {
	if o == nil || IsNil(o.Folders) {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *Folders) HasFolders() bool {
	if o != nil && !IsNil(o.Folders) {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []Folder and assigns it to the Folders field.
func (o *Folders) SetFolders(v []Folder) {
	o.Folders = v
}

// GetNextUri returns the NextUri field value if set, zero value otherwise.
func (o *Folders) GetNextUri() string {
	if o == nil || IsNil(o.NextUri) {
		var ret string
		return ret
	}
	return *o.NextUri
}

// GetNextUriOk returns a tuple with the NextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetNextUriOk() (*string, bool) {
	if o == nil || IsNil(o.NextUri) {
		return nil, false
	}
	return o.NextUri, true
}

// HasNextUri returns a boolean if a field has been set.
func (o *Folders) HasNextUri() bool {
	if o != nil && !IsNil(o.NextUri) {
		return true
	}

	return false
}

// SetNextUri gets a reference to the given string and assigns it to the NextUri field.
func (o *Folders) SetNextUri(v string) {
	o.NextUri = &v
}

// GetPreviousUri returns the PreviousUri field value if set, zero value otherwise.
func (o *Folders) GetPreviousUri() string {
	if o == nil || IsNil(o.PreviousUri) {
		var ret string
		return ret
	}
	return *o.PreviousUri
}

// GetPreviousUriOk returns a tuple with the PreviousUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetPreviousUriOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousUri) {
		return nil, false
	}
	return o.PreviousUri, true
}

// HasPreviousUri returns a boolean if a field has been set.
func (o *Folders) HasPreviousUri() bool {
	if o != nil && !IsNil(o.PreviousUri) {
		return true
	}

	return false
}

// SetPreviousUri gets a reference to the given string and assigns it to the PreviousUri field.
func (o *Folders) SetPreviousUri(v string) {
	o.PreviousUri = &v
}

// GetResultSetSize returns the ResultSetSize field value if set, zero value otherwise.
func (o *Folders) GetResultSetSize() string {
	if o == nil || IsNil(o.ResultSetSize) {
		var ret string
		return ret
	}
	return *o.ResultSetSize
}

// GetResultSetSizeOk returns a tuple with the ResultSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetResultSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultSetSize) {
		return nil, false
	}
	return o.ResultSetSize, true
}

// HasResultSetSize returns a boolean if a field has been set.
func (o *Folders) HasResultSetSize() bool {
	if o != nil && !IsNil(o.ResultSetSize) {
		return true
	}

	return false
}

// SetResultSetSize gets a reference to the given string and assigns it to the ResultSetSize field.
func (o *Folders) SetResultSetSize(v string) {
	o.ResultSetSize = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *Folders) GetStartPosition() string {
	if o == nil || IsNil(o.StartPosition) {
		var ret string
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetStartPositionOk() (*string, bool) {
	if o == nil || IsNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *Folders) HasStartPosition() bool {
	if o != nil && !IsNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given string and assigns it to the StartPosition field.
func (o *Folders) SetStartPosition(v string) {
	o.StartPosition = &v
}

// GetTotalSetSize returns the TotalSetSize field value if set, zero value otherwise.
func (o *Folders) GetTotalSetSize() string {
	if o == nil || IsNil(o.TotalSetSize) {
		var ret string
		return ret
	}
	return *o.TotalSetSize
}

// GetTotalSetSizeOk returns a tuple with the TotalSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Folders) GetTotalSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSetSize) {
		return nil, false
	}
	return o.TotalSetSize, true
}

// HasTotalSetSize returns a boolean if a field has been set.
func (o *Folders) HasTotalSetSize() bool {
	if o != nil && !IsNil(o.TotalSetSize) {
		return true
	}

	return false
}

// SetTotalSetSize gets a reference to the given string and assigns it to the TotalSetSize field.
func (o *Folders) SetTotalSetSize(v string) {
	o.TotalSetSize = &v
}

func (o Folders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Folders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !IsNil(o.Envelopes) {
		toSerialize["envelopes"] = o.Envelopes
	}
	if !IsNil(o.Folders) {
		toSerialize["folders"] = o.Folders
	}
	if !IsNil(o.NextUri) {
		toSerialize["nextUri"] = o.NextUri
	}
	if !IsNil(o.PreviousUri) {
		toSerialize["previousUri"] = o.PreviousUri
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
	return toSerialize, nil
}

type NullableFolders struct {
	value *Folders
	isSet bool
}

func (v NullableFolders) Get() *Folders {
	return v.value
}

func (v *NullableFolders) Set(val *Folders) {
	v.value = val
	v.isSet = true
}

func (v NullableFolders) IsSet() bool {
	return v.isSet
}

func (v *NullableFolders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolders(val *Folders) *NullableFolders {
	return &NullableFolders{value: val, isSet: true}
}

func (v NullableFolders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


