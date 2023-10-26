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

// checks if the FavoriteTemplates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FavoriteTemplates{}

// FavoriteTemplates 
type FavoriteTemplates struct {
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// 
	FavoriteTemplates []FavoriteTemplatesContentItem `json:"favoriteTemplates,omitempty"`
	// The number of templates successfully updated by the call. This property is read-only.
	TemplatesUpdatedCount *int32 `json:"templatesUpdatedCount,omitempty"`
}

// NewFavoriteTemplates instantiates a new FavoriteTemplates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFavoriteTemplates() *FavoriteTemplates {
	this := FavoriteTemplates{}
	return &this
}

// NewFavoriteTemplatesWithDefaults instantiates a new FavoriteTemplates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFavoriteTemplatesWithDefaults() *FavoriteTemplates {
	this := FavoriteTemplates{}
	return &this
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *FavoriteTemplates) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoriteTemplates) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *FavoriteTemplates) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *FavoriteTemplates) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetFavoriteTemplates returns the FavoriteTemplates field value if set, zero value otherwise.
func (o *FavoriteTemplates) GetFavoriteTemplates() []FavoriteTemplatesContentItem {
	if o == nil || IsNil(o.FavoriteTemplates) {
		var ret []FavoriteTemplatesContentItem
		return ret
	}
	return o.FavoriteTemplates
}

// GetFavoriteTemplatesOk returns a tuple with the FavoriteTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoriteTemplates) GetFavoriteTemplatesOk() ([]FavoriteTemplatesContentItem, bool) {
	if o == nil || IsNil(o.FavoriteTemplates) {
		return nil, false
	}
	return o.FavoriteTemplates, true
}

// HasFavoriteTemplates returns a boolean if a field has been set.
func (o *FavoriteTemplates) HasFavoriteTemplates() bool {
	if o != nil && !IsNil(o.FavoriteTemplates) {
		return true
	}

	return false
}

// SetFavoriteTemplates gets a reference to the given []FavoriteTemplatesContentItem and assigns it to the FavoriteTemplates field.
func (o *FavoriteTemplates) SetFavoriteTemplates(v []FavoriteTemplatesContentItem) {
	o.FavoriteTemplates = v
}

// GetTemplatesUpdatedCount returns the TemplatesUpdatedCount field value if set, zero value otherwise.
func (o *FavoriteTemplates) GetTemplatesUpdatedCount() int32 {
	if o == nil || IsNil(o.TemplatesUpdatedCount) {
		var ret int32
		return ret
	}
	return *o.TemplatesUpdatedCount
}

// GetTemplatesUpdatedCountOk returns a tuple with the TemplatesUpdatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FavoriteTemplates) GetTemplatesUpdatedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplatesUpdatedCount) {
		return nil, false
	}
	return o.TemplatesUpdatedCount, true
}

// HasTemplatesUpdatedCount returns a boolean if a field has been set.
func (o *FavoriteTemplates) HasTemplatesUpdatedCount() bool {
	if o != nil && !IsNil(o.TemplatesUpdatedCount) {
		return true
	}

	return false
}

// SetTemplatesUpdatedCount gets a reference to the given int32 and assigns it to the TemplatesUpdatedCount field.
func (o *FavoriteTemplates) SetTemplatesUpdatedCount(v int32) {
	o.TemplatesUpdatedCount = &v
}

func (o FavoriteTemplates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FavoriteTemplates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.FavoriteTemplates) {
		toSerialize["favoriteTemplates"] = o.FavoriteTemplates
	}
	if !IsNil(o.TemplatesUpdatedCount) {
		toSerialize["templatesUpdatedCount"] = o.TemplatesUpdatedCount
	}
	return toSerialize, nil
}

type NullableFavoriteTemplates struct {
	value *FavoriteTemplates
	isSet bool
}

func (v NullableFavoriteTemplates) Get() *FavoriteTemplates {
	return v.value
}

func (v *NullableFavoriteTemplates) Set(val *FavoriteTemplates) {
	v.value = val
	v.isSet = true
}

func (v NullableFavoriteTemplates) IsSet() bool {
	return v.isSet
}

func (v *NullableFavoriteTemplates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFavoriteTemplates(val *FavoriteTemplates) *NullableFavoriteTemplates {
	return &NullableFavoriteTemplates{value: val, isSet: true}
}

func (v NullableFavoriteTemplates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFavoriteTemplates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

