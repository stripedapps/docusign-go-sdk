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

// checks if the ExternalFolder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalFolder{}

// ExternalFolder 
type ExternalFolder struct {
	// The last index position in the result set. 
	EndPosition *string `json:"endPosition,omitempty"`
	ErrorDetails *ExternalDocServiceErrorDetails `json:"errorDetails,omitempty"`
	// A unique ID for the Salesforce object.
	Id *string `json:"id,omitempty"`
	// If the tab is a list, this represents the values that are possible for the tab.
	Items []ExternalFile `json:"items,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
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

// NewExternalFolder instantiates a new ExternalFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalFolder() *ExternalFolder {
	this := ExternalFolder{}
	return &this
}

// NewExternalFolderWithDefaults instantiates a new ExternalFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalFolderWithDefaults() *ExternalFolder {
	this := ExternalFolder{}
	return &this
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *ExternalFolder) GetEndPosition() string {
	if o == nil || IsNil(o.EndPosition) {
		var ret string
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetEndPositionOk() (*string, bool) {
	if o == nil || IsNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *ExternalFolder) HasEndPosition() bool {
	if o != nil && !IsNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given string and assigns it to the EndPosition field.
func (o *ExternalFolder) SetEndPosition(v string) {
	o.EndPosition = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ExternalFolder) GetErrorDetails() ExternalDocServiceErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ExternalDocServiceErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetErrorDetailsOk() (*ExternalDocServiceErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ExternalFolder) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ExternalDocServiceErrorDetails and assigns it to the ErrorDetails field.
func (o *ExternalFolder) SetErrorDetails(v ExternalDocServiceErrorDetails) {
	o.ErrorDetails = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExternalFolder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExternalFolder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ExternalFolder) SetId(v string) {
	o.Id = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ExternalFolder) GetItems() []ExternalFile {
	if o == nil || IsNil(o.Items) {
		var ret []ExternalFile
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetItemsOk() ([]ExternalFile, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ExternalFolder) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ExternalFile and assigns it to the Items field.
func (o *ExternalFolder) SetItems(v []ExternalFile) {
	o.Items = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExternalFolder) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExternalFolder) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExternalFolder) SetName(v string) {
	o.Name = &v
}

// GetNextUri returns the NextUri field value if set, zero value otherwise.
func (o *ExternalFolder) GetNextUri() string {
	if o == nil || IsNil(o.NextUri) {
		var ret string
		return ret
	}
	return *o.NextUri
}

// GetNextUriOk returns a tuple with the NextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetNextUriOk() (*string, bool) {
	if o == nil || IsNil(o.NextUri) {
		return nil, false
	}
	return o.NextUri, true
}

// HasNextUri returns a boolean if a field has been set.
func (o *ExternalFolder) HasNextUri() bool {
	if o != nil && !IsNil(o.NextUri) {
		return true
	}

	return false
}

// SetNextUri gets a reference to the given string and assigns it to the NextUri field.
func (o *ExternalFolder) SetNextUri(v string) {
	o.NextUri = &v
}

// GetPreviousUri returns the PreviousUri field value if set, zero value otherwise.
func (o *ExternalFolder) GetPreviousUri() string {
	if o == nil || IsNil(o.PreviousUri) {
		var ret string
		return ret
	}
	return *o.PreviousUri
}

// GetPreviousUriOk returns a tuple with the PreviousUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetPreviousUriOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousUri) {
		return nil, false
	}
	return o.PreviousUri, true
}

// HasPreviousUri returns a boolean if a field has been set.
func (o *ExternalFolder) HasPreviousUri() bool {
	if o != nil && !IsNil(o.PreviousUri) {
		return true
	}

	return false
}

// SetPreviousUri gets a reference to the given string and assigns it to the PreviousUri field.
func (o *ExternalFolder) SetPreviousUri(v string) {
	o.PreviousUri = &v
}

// GetResultSetSize returns the ResultSetSize field value if set, zero value otherwise.
func (o *ExternalFolder) GetResultSetSize() string {
	if o == nil || IsNil(o.ResultSetSize) {
		var ret string
		return ret
	}
	return *o.ResultSetSize
}

// GetResultSetSizeOk returns a tuple with the ResultSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetResultSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultSetSize) {
		return nil, false
	}
	return o.ResultSetSize, true
}

// HasResultSetSize returns a boolean if a field has been set.
func (o *ExternalFolder) HasResultSetSize() bool {
	if o != nil && !IsNil(o.ResultSetSize) {
		return true
	}

	return false
}

// SetResultSetSize gets a reference to the given string and assigns it to the ResultSetSize field.
func (o *ExternalFolder) SetResultSetSize(v string) {
	o.ResultSetSize = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *ExternalFolder) GetStartPosition() string {
	if o == nil || IsNil(o.StartPosition) {
		var ret string
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetStartPositionOk() (*string, bool) {
	if o == nil || IsNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *ExternalFolder) HasStartPosition() bool {
	if o != nil && !IsNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given string and assigns it to the StartPosition field.
func (o *ExternalFolder) SetStartPosition(v string) {
	o.StartPosition = &v
}

// GetTotalSetSize returns the TotalSetSize field value if set, zero value otherwise.
func (o *ExternalFolder) GetTotalSetSize() string {
	if o == nil || IsNil(o.TotalSetSize) {
		var ret string
		return ret
	}
	return *o.TotalSetSize
}

// GetTotalSetSizeOk returns a tuple with the TotalSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalFolder) GetTotalSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSetSize) {
		return nil, false
	}
	return o.TotalSetSize, true
}

// HasTotalSetSize returns a boolean if a field has been set.
func (o *ExternalFolder) HasTotalSetSize() bool {
	if o != nil && !IsNil(o.TotalSetSize) {
		return true
	}

	return false
}

// SetTotalSetSize gets a reference to the given string and assigns it to the TotalSetSize field.
func (o *ExternalFolder) SetTotalSetSize(v string) {
	o.TotalSetSize = &v
}

func (o ExternalFolder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalFolder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

type NullableExternalFolder struct {
	value *ExternalFolder
	isSet bool
}

func (v NullableExternalFolder) Get() *ExternalFolder {
	return v.value
}

func (v *NullableExternalFolder) Set(val *ExternalFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalFolder(val *ExternalFolder) *NullableExternalFolder {
	return &NullableExternalFolder{value: val, isSet: true}
}

func (v NullableExternalFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


