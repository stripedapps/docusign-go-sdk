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

// checks if the FileTypeList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileTypeList{}

// FileTypeList 
type FileTypeList struct {
	// A collection of file types.
	FileTypes []FileType `json:"fileTypes,omitempty"`
}

// NewFileTypeList instantiates a new FileTypeList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileTypeList() *FileTypeList {
	this := FileTypeList{}
	return &this
}

// NewFileTypeListWithDefaults instantiates a new FileTypeList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTypeListWithDefaults() *FileTypeList {
	this := FileTypeList{}
	return &this
}

// GetFileTypes returns the FileTypes field value if set, zero value otherwise.
func (o *FileTypeList) GetFileTypes() []FileType {
	if o == nil || IsNil(o.FileTypes) {
		var ret []FileType
		return ret
	}
	return o.FileTypes
}

// GetFileTypesOk returns a tuple with the FileTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileTypeList) GetFileTypesOk() ([]FileType, bool) {
	if o == nil || IsNil(o.FileTypes) {
		return nil, false
	}
	return o.FileTypes, true
}

// HasFileTypes returns a boolean if a field has been set.
func (o *FileTypeList) HasFileTypes() bool {
	if o != nil && !IsNil(o.FileTypes) {
		return true
	}

	return false
}

// SetFileTypes gets a reference to the given []FileType and assigns it to the FileTypes field.
func (o *FileTypeList) SetFileTypes(v []FileType) {
	o.FileTypes = v
}

func (o FileTypeList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileTypeList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileTypes) {
		toSerialize["fileTypes"] = o.FileTypes
	}
	return toSerialize, nil
}

type NullableFileTypeList struct {
	value *FileTypeList
	isSet bool
}

func (v NullableFileTypeList) Get() *FileTypeList {
	return v.value
}

func (v *NullableFileTypeList) Set(val *FileTypeList) {
	v.value = val
	v.isSet = true
}

func (v NullableFileTypeList) IsSet() bool {
	return v.isSet
}

func (v *NullableFileTypeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileTypeList(val *FileTypeList) *NullableFileTypeList {
	return &NullableFileTypeList{value: val, isSet: true}
}

func (v NullableFileTypeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileTypeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

