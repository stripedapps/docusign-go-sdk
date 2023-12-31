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

// checks if the FileType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileType{}

// FileType 
type FileType struct {
	// 
	FileExtension *string `json:"fileExtension,omitempty"`
	// The mime-type of a file type listed in a fileTypes collection.
	MimeType *string `json:"mimeType,omitempty"`
}

// NewFileType instantiates a new FileType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileType() *FileType {
	this := FileType{}
	return &this
}

// NewFileTypeWithDefaults instantiates a new FileType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileTypeWithDefaults() *FileType {
	this := FileType{}
	return &this
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *FileType) GetFileExtension() string {
	if o == nil || IsNil(o.FileExtension) {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileType) GetFileExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.FileExtension) {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *FileType) HasFileExtension() bool {
	if o != nil && !IsNil(o.FileExtension) {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *FileType) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetMimeType returns the MimeType field value if set, zero value otherwise.
func (o *FileType) GetMimeType() string {
	if o == nil || IsNil(o.MimeType) {
		var ret string
		return ret
	}
	return *o.MimeType
}

// GetMimeTypeOk returns a tuple with the MimeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileType) GetMimeTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MimeType) {
		return nil, false
	}
	return o.MimeType, true
}

// HasMimeType returns a boolean if a field has been set.
func (o *FileType) HasMimeType() bool {
	if o != nil && !IsNil(o.MimeType) {
		return true
	}

	return false
}

// SetMimeType gets a reference to the given string and assigns it to the MimeType field.
func (o *FileType) SetMimeType(v string) {
	o.MimeType = &v
}

func (o FileType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileExtension) {
		toSerialize["fileExtension"] = o.FileExtension
	}
	if !IsNil(o.MimeType) {
		toSerialize["mimeType"] = o.MimeType
	}
	return toSerialize, nil
}

type NullableFileType struct {
	value *FileType
	isSet bool
}

func (v NullableFileType) Get() *FileType {
	return v.value
}

func (v *NullableFileType) Set(val *FileType) {
	v.value = val
	v.isSet = true
}

func (v NullableFileType) IsSet() bool {
	return v.isSet
}

func (v *NullableFileType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileType(val *FileType) *NullableFileType {
	return &NullableFileType{value: val, isSet: true}
}

func (v NullableFileType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


