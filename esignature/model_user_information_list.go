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

// checks if the UserInformationList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserInformationList{}

// UserInformationList Contains a list of account users.
type UserInformationList struct {
	// The last index position in the result set. 
	EndPosition *string `json:"endPosition,omitempty"`
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
	// User management information.
	Users []UserInformation `json:"users,omitempty"`
}

// NewUserInformationList instantiates a new UserInformationList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserInformationList() *UserInformationList {
	this := UserInformationList{}
	return &this
}

// NewUserInformationListWithDefaults instantiates a new UserInformationList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserInformationListWithDefaults() *UserInformationList {
	this := UserInformationList{}
	return &this
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *UserInformationList) GetEndPosition() string {
	if o == nil || IsNil(o.EndPosition) {
		var ret string
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetEndPositionOk() (*string, bool) {
	if o == nil || IsNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *UserInformationList) HasEndPosition() bool {
	if o != nil && !IsNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given string and assigns it to the EndPosition field.
func (o *UserInformationList) SetEndPosition(v string) {
	o.EndPosition = &v
}

// GetNextUri returns the NextUri field value if set, zero value otherwise.
func (o *UserInformationList) GetNextUri() string {
	if o == nil || IsNil(o.NextUri) {
		var ret string
		return ret
	}
	return *o.NextUri
}

// GetNextUriOk returns a tuple with the NextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetNextUriOk() (*string, bool) {
	if o == nil || IsNil(o.NextUri) {
		return nil, false
	}
	return o.NextUri, true
}

// HasNextUri returns a boolean if a field has been set.
func (o *UserInformationList) HasNextUri() bool {
	if o != nil && !IsNil(o.NextUri) {
		return true
	}

	return false
}

// SetNextUri gets a reference to the given string and assigns it to the NextUri field.
func (o *UserInformationList) SetNextUri(v string) {
	o.NextUri = &v
}

// GetPreviousUri returns the PreviousUri field value if set, zero value otherwise.
func (o *UserInformationList) GetPreviousUri() string {
	if o == nil || IsNil(o.PreviousUri) {
		var ret string
		return ret
	}
	return *o.PreviousUri
}

// GetPreviousUriOk returns a tuple with the PreviousUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetPreviousUriOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousUri) {
		return nil, false
	}
	return o.PreviousUri, true
}

// HasPreviousUri returns a boolean if a field has been set.
func (o *UserInformationList) HasPreviousUri() bool {
	if o != nil && !IsNil(o.PreviousUri) {
		return true
	}

	return false
}

// SetPreviousUri gets a reference to the given string and assigns it to the PreviousUri field.
func (o *UserInformationList) SetPreviousUri(v string) {
	o.PreviousUri = &v
}

// GetResultSetSize returns the ResultSetSize field value if set, zero value otherwise.
func (o *UserInformationList) GetResultSetSize() string {
	if o == nil || IsNil(o.ResultSetSize) {
		var ret string
		return ret
	}
	return *o.ResultSetSize
}

// GetResultSetSizeOk returns a tuple with the ResultSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetResultSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.ResultSetSize) {
		return nil, false
	}
	return o.ResultSetSize, true
}

// HasResultSetSize returns a boolean if a field has been set.
func (o *UserInformationList) HasResultSetSize() bool {
	if o != nil && !IsNil(o.ResultSetSize) {
		return true
	}

	return false
}

// SetResultSetSize gets a reference to the given string and assigns it to the ResultSetSize field.
func (o *UserInformationList) SetResultSetSize(v string) {
	o.ResultSetSize = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *UserInformationList) GetStartPosition() string {
	if o == nil || IsNil(o.StartPosition) {
		var ret string
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetStartPositionOk() (*string, bool) {
	if o == nil || IsNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *UserInformationList) HasStartPosition() bool {
	if o != nil && !IsNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given string and assigns it to the StartPosition field.
func (o *UserInformationList) SetStartPosition(v string) {
	o.StartPosition = &v
}

// GetTotalSetSize returns the TotalSetSize field value if set, zero value otherwise.
func (o *UserInformationList) GetTotalSetSize() string {
	if o == nil || IsNil(o.TotalSetSize) {
		var ret string
		return ret
	}
	return *o.TotalSetSize
}

// GetTotalSetSizeOk returns a tuple with the TotalSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetTotalSetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSetSize) {
		return nil, false
	}
	return o.TotalSetSize, true
}

// HasTotalSetSize returns a boolean if a field has been set.
func (o *UserInformationList) HasTotalSetSize() bool {
	if o != nil && !IsNil(o.TotalSetSize) {
		return true
	}

	return false
}

// SetTotalSetSize gets a reference to the given string and assigns it to the TotalSetSize field.
func (o *UserInformationList) SetTotalSetSize(v string) {
	o.TotalSetSize = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UserInformationList) GetUsers() []UserInformation {
	if o == nil || IsNil(o.Users) {
		var ret []UserInformation
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserInformationList) GetUsersOk() ([]UserInformation, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserInformationList) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []UserInformation and assigns it to the Users field.
func (o *UserInformationList) SetUsers(v []UserInformation) {
	o.Users = v
}

func (o UserInformationList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserInformationList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
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
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableUserInformationList struct {
	value *UserInformationList
	isSet bool
}

func (v NullableUserInformationList) Get() *UserInformationList {
	return v.value
}

func (v *NullableUserInformationList) Set(val *UserInformationList) {
	v.value = val
	v.isSet = true
}

func (v NullableUserInformationList) IsSet() bool {
	return v.isSet
}

func (v *NullableUserInformationList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserInformationList(val *UserInformationList) *NullableUserInformationList {
	return &NullableUserInformationList{value: val, isSet: true}
}

func (v NullableUserInformationList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserInformationList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


