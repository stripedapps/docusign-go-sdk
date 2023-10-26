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

// checks if the SupportedLanguages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedLanguages{}

// SupportedLanguages A list of supported languages.
type SupportedLanguages struct {
	// A list of languages that you can use for a recipient's language setting. These are the languages that you can set for the standard email format and signing view for each recipient.  For example, in the recipient's email notification, this setting affects elements such as the standard introductory text describing the request to sign. It also determines the language used for buttons and tabs in both the email notification and the signing experience.  **Note:** Setting a language for a recipient affects only the DocuSign standard text. Any custom text that you enter for the `emailBody` and `emailSubject` of the notification is not translated, and appears exactly as you enter it.  Example:  ``` {     \"languages\": [         {             \"name\": \"Arabic (ar)\",             \"value\": \"ar\"         },         {             \"name\": \"Bulgarian (bg)\",             \"value\": \"bg\"         },         .         .         . } ```
	Languages []NameValue `json:"languages,omitempty"`
}

// NewSupportedLanguages instantiates a new SupportedLanguages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedLanguages() *SupportedLanguages {
	this := SupportedLanguages{}
	return &this
}

// NewSupportedLanguagesWithDefaults instantiates a new SupportedLanguages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedLanguagesWithDefaults() *SupportedLanguages {
	this := SupportedLanguages{}
	return &this
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *SupportedLanguages) GetLanguages() []NameValue {
	if o == nil || IsNil(o.Languages) {
		var ret []NameValue
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedLanguages) GetLanguagesOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.Languages) {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *SupportedLanguages) HasLanguages() bool {
	if o != nil && !IsNil(o.Languages) {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []NameValue and assigns it to the Languages field.
func (o *SupportedLanguages) SetLanguages(v []NameValue) {
	o.Languages = v
}

func (o SupportedLanguages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedLanguages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Languages) {
		toSerialize["languages"] = o.Languages
	}
	return toSerialize, nil
}

type NullableSupportedLanguages struct {
	value *SupportedLanguages
	isSet bool
}

func (v NullableSupportedLanguages) Get() *SupportedLanguages {
	return v.value
}

func (v *NullableSupportedLanguages) Set(val *SupportedLanguages) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedLanguages) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedLanguages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedLanguages(val *SupportedLanguages) *NullableSupportedLanguages {
	return &NullableSupportedLanguages{value: val, isSet: true}
}

func (v NullableSupportedLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedLanguages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

