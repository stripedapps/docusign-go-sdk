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

// checks if the DocumentHtmlDisplayAnchor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentHtmlDisplayAnchor{}

// DocumentHtmlDisplayAnchor 
type DocumentHtmlDisplayAnchor struct {
	// When **true,** the start or end anchor strings must match the strings specified by the start and end anchor settings in case as well as in content.
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	DisplaySettings *DocumentHtmlDisplaySettings `json:"displaySettings,omitempty"`
	// Specifies the end of the area in the HTML where the display settings will be applied. If you do not specify an end anchor, the end of the document will be used by default.  **Note:** A start anchor, an end anchor, or both are required.
	EndAnchor *string `json:"endAnchor,omitempty"`
	// When **true,** removes the end anchor string for the Smart Section from the HTML, preventing it from displaying.
	RemoveEndAnchor *bool `json:"removeEndAnchor,omitempty"`
	// When **true,** removes the start anchor string for the Smart Section from the HTML, preventing it from displaying.
	RemoveStartAnchor *bool `json:"removeStartAnchor,omitempty"`
	// Specifies the beginning of the area in the HTML where the display settings will be applied. If you do not specify a start anchor, the beginning of the document will be used by default.  **Note:** A start anchor, an end anchor, or both are required.
	StartAnchor *string `json:"startAnchor,omitempty"`
}

// NewDocumentHtmlDisplayAnchor instantiates a new DocumentHtmlDisplayAnchor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentHtmlDisplayAnchor() *DocumentHtmlDisplayAnchor {
	this := DocumentHtmlDisplayAnchor{}
	return &this
}

// NewDocumentHtmlDisplayAnchorWithDefaults instantiates a new DocumentHtmlDisplayAnchor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentHtmlDisplayAnchorWithDefaults() *DocumentHtmlDisplayAnchor {
	this := DocumentHtmlDisplayAnchor{}
	return &this
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetCaseSensitive() bool {
	if o == nil || IsNil(o.CaseSensitive) {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseSensitive) {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasCaseSensitive() bool {
	if o != nil && !IsNil(o.CaseSensitive) {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *DocumentHtmlDisplayAnchor) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

// GetDisplaySettings returns the DisplaySettings field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetDisplaySettings() DocumentHtmlDisplaySettings {
	if o == nil || IsNil(o.DisplaySettings) {
		var ret DocumentHtmlDisplaySettings
		return ret
	}
	return *o.DisplaySettings
}

// GetDisplaySettingsOk returns a tuple with the DisplaySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetDisplaySettingsOk() (*DocumentHtmlDisplaySettings, bool) {
	if o == nil || IsNil(o.DisplaySettings) {
		return nil, false
	}
	return o.DisplaySettings, true
}

// HasDisplaySettings returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasDisplaySettings() bool {
	if o != nil && !IsNil(o.DisplaySettings) {
		return true
	}

	return false
}

// SetDisplaySettings gets a reference to the given DocumentHtmlDisplaySettings and assigns it to the DisplaySettings field.
func (o *DocumentHtmlDisplayAnchor) SetDisplaySettings(v DocumentHtmlDisplaySettings) {
	o.DisplaySettings = &v
}

// GetEndAnchor returns the EndAnchor field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetEndAnchor() string {
	if o == nil || IsNil(o.EndAnchor) {
		var ret string
		return ret
	}
	return *o.EndAnchor
}

// GetEndAnchorOk returns a tuple with the EndAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetEndAnchorOk() (*string, bool) {
	if o == nil || IsNil(o.EndAnchor) {
		return nil, false
	}
	return o.EndAnchor, true
}

// HasEndAnchor returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasEndAnchor() bool {
	if o != nil && !IsNil(o.EndAnchor) {
		return true
	}

	return false
}

// SetEndAnchor gets a reference to the given string and assigns it to the EndAnchor field.
func (o *DocumentHtmlDisplayAnchor) SetEndAnchor(v string) {
	o.EndAnchor = &v
}

// GetRemoveEndAnchor returns the RemoveEndAnchor field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetRemoveEndAnchor() bool {
	if o == nil || IsNil(o.RemoveEndAnchor) {
		var ret bool
		return ret
	}
	return *o.RemoveEndAnchor
}

// GetRemoveEndAnchorOk returns a tuple with the RemoveEndAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetRemoveEndAnchorOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveEndAnchor) {
		return nil, false
	}
	return o.RemoveEndAnchor, true
}

// HasRemoveEndAnchor returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasRemoveEndAnchor() bool {
	if o != nil && !IsNil(o.RemoveEndAnchor) {
		return true
	}

	return false
}

// SetRemoveEndAnchor gets a reference to the given bool and assigns it to the RemoveEndAnchor field.
func (o *DocumentHtmlDisplayAnchor) SetRemoveEndAnchor(v bool) {
	o.RemoveEndAnchor = &v
}

// GetRemoveStartAnchor returns the RemoveStartAnchor field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetRemoveStartAnchor() bool {
	if o == nil || IsNil(o.RemoveStartAnchor) {
		var ret bool
		return ret
	}
	return *o.RemoveStartAnchor
}

// GetRemoveStartAnchorOk returns a tuple with the RemoveStartAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetRemoveStartAnchorOk() (*bool, bool) {
	if o == nil || IsNil(o.RemoveStartAnchor) {
		return nil, false
	}
	return o.RemoveStartAnchor, true
}

// HasRemoveStartAnchor returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasRemoveStartAnchor() bool {
	if o != nil && !IsNil(o.RemoveStartAnchor) {
		return true
	}

	return false
}

// SetRemoveStartAnchor gets a reference to the given bool and assigns it to the RemoveStartAnchor field.
func (o *DocumentHtmlDisplayAnchor) SetRemoveStartAnchor(v bool) {
	o.RemoveStartAnchor = &v
}

// GetStartAnchor returns the StartAnchor field value if set, zero value otherwise.
func (o *DocumentHtmlDisplayAnchor) GetStartAnchor() string {
	if o == nil || IsNil(o.StartAnchor) {
		var ret string
		return ret
	}
	return *o.StartAnchor
}

// GetStartAnchorOk returns a tuple with the StartAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentHtmlDisplayAnchor) GetStartAnchorOk() (*string, bool) {
	if o == nil || IsNil(o.StartAnchor) {
		return nil, false
	}
	return o.StartAnchor, true
}

// HasStartAnchor returns a boolean if a field has been set.
func (o *DocumentHtmlDisplayAnchor) HasStartAnchor() bool {
	if o != nil && !IsNil(o.StartAnchor) {
		return true
	}

	return false
}

// SetStartAnchor gets a reference to the given string and assigns it to the StartAnchor field.
func (o *DocumentHtmlDisplayAnchor) SetStartAnchor(v string) {
	o.StartAnchor = &v
}

func (o DocumentHtmlDisplayAnchor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentHtmlDisplayAnchor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CaseSensitive) {
		toSerialize["caseSensitive"] = o.CaseSensitive
	}
	if !IsNil(o.DisplaySettings) {
		toSerialize["displaySettings"] = o.DisplaySettings
	}
	if !IsNil(o.EndAnchor) {
		toSerialize["endAnchor"] = o.EndAnchor
	}
	if !IsNil(o.RemoveEndAnchor) {
		toSerialize["removeEndAnchor"] = o.RemoveEndAnchor
	}
	if !IsNil(o.RemoveStartAnchor) {
		toSerialize["removeStartAnchor"] = o.RemoveStartAnchor
	}
	if !IsNil(o.StartAnchor) {
		toSerialize["startAnchor"] = o.StartAnchor
	}
	return toSerialize, nil
}

type NullableDocumentHtmlDisplayAnchor struct {
	value *DocumentHtmlDisplayAnchor
	isSet bool
}

func (v NullableDocumentHtmlDisplayAnchor) Get() *DocumentHtmlDisplayAnchor {
	return v.value
}

func (v *NullableDocumentHtmlDisplayAnchor) Set(val *DocumentHtmlDisplayAnchor) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentHtmlDisplayAnchor) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentHtmlDisplayAnchor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentHtmlDisplayAnchor(val *DocumentHtmlDisplayAnchor) *NullableDocumentHtmlDisplayAnchor {
	return &NullableDocumentHtmlDisplayAnchor{value: val, isSet: true}
}

func (v NullableDocumentHtmlDisplayAnchor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentHtmlDisplayAnchor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

