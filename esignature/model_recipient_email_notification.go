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

// checks if the RecipientEmailNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientEmailNotification{}

// RecipientEmailNotification Sets custom email subject and email body for individual recipients. **Note:** You must explicitly set `supportedLanguage` if you use this feature. 
type RecipientEmailNotification struct {
	// The body of the email message.
	EmailBody *string `json:"emailBody,omitempty"`
	EmailBodyMetadata *PropertyMetadata `json:"emailBodyMetadata,omitempty"`
	// The subject line for the email notification.
	EmailSubject *string `json:"emailSubject,omitempty"`
	EmailSubjectMetadata *PropertyMetadata `json:"emailSubjectMetadata,omitempty"`
	// The language to use for the standard email format and signing view for a recipient.  For example, this setting determines the language of the recipient's email notification message. It also determines the language used for buttons and tabs in both the email notification and the signing experience.  **Note:** This setting affects only DocuSign standard text. Any custom text that you enter for the `emailBody` and `emailSubject` of the notification is not translated, and appears exactly as you enter it.  To retrieve the possible values, use the [Accounts::listSupportedLanguages][ListLang] method.  [ListLang]: /docs/esign-rest-api/reference/accounts/accounts/listsupportedlanguages/ 
	SupportedLanguage *string `json:"supportedLanguage,omitempty"`
	SupportedLanguageMetadata *PropertyMetadata `json:"supportedLanguageMetadata,omitempty"`
}

// NewRecipientEmailNotification instantiates a new RecipientEmailNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientEmailNotification() *RecipientEmailNotification {
	this := RecipientEmailNotification{}
	return &this
}

// NewRecipientEmailNotificationWithDefaults instantiates a new RecipientEmailNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientEmailNotificationWithDefaults() *RecipientEmailNotification {
	this := RecipientEmailNotification{}
	return &this
}

// GetEmailBody returns the EmailBody field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetEmailBody() string {
	if o == nil || IsNil(o.EmailBody) {
		var ret string
		return ret
	}
	return *o.EmailBody
}

// GetEmailBodyOk returns a tuple with the EmailBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetEmailBodyOk() (*string, bool) {
	if o == nil || IsNil(o.EmailBody) {
		return nil, false
	}
	return o.EmailBody, true
}

// HasEmailBody returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasEmailBody() bool {
	if o != nil && !IsNil(o.EmailBody) {
		return true
	}

	return false
}

// SetEmailBody gets a reference to the given string and assigns it to the EmailBody field.
func (o *RecipientEmailNotification) SetEmailBody(v string) {
	o.EmailBody = &v
}

// GetEmailBodyMetadata returns the EmailBodyMetadata field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetEmailBodyMetadata() PropertyMetadata {
	if o == nil || IsNil(o.EmailBodyMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.EmailBodyMetadata
}

// GetEmailBodyMetadataOk returns a tuple with the EmailBodyMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetEmailBodyMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.EmailBodyMetadata) {
		return nil, false
	}
	return o.EmailBodyMetadata, true
}

// HasEmailBodyMetadata returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasEmailBodyMetadata() bool {
	if o != nil && !IsNil(o.EmailBodyMetadata) {
		return true
	}

	return false
}

// SetEmailBodyMetadata gets a reference to the given PropertyMetadata and assigns it to the EmailBodyMetadata field.
func (o *RecipientEmailNotification) SetEmailBodyMetadata(v PropertyMetadata) {
	o.EmailBodyMetadata = &v
}

// GetEmailSubject returns the EmailSubject field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetEmailSubject() string {
	if o == nil || IsNil(o.EmailSubject) {
		var ret string
		return ret
	}
	return *o.EmailSubject
}

// GetEmailSubjectOk returns a tuple with the EmailSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetEmailSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubject) {
		return nil, false
	}
	return o.EmailSubject, true
}

// HasEmailSubject returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasEmailSubject() bool {
	if o != nil && !IsNil(o.EmailSubject) {
		return true
	}

	return false
}

// SetEmailSubject gets a reference to the given string and assigns it to the EmailSubject field.
func (o *RecipientEmailNotification) SetEmailSubject(v string) {
	o.EmailSubject = &v
}

// GetEmailSubjectMetadata returns the EmailSubjectMetadata field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetEmailSubjectMetadata() PropertyMetadata {
	if o == nil || IsNil(o.EmailSubjectMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.EmailSubjectMetadata
}

// GetEmailSubjectMetadataOk returns a tuple with the EmailSubjectMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetEmailSubjectMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.EmailSubjectMetadata) {
		return nil, false
	}
	return o.EmailSubjectMetadata, true
}

// HasEmailSubjectMetadata returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasEmailSubjectMetadata() bool {
	if o != nil && !IsNil(o.EmailSubjectMetadata) {
		return true
	}

	return false
}

// SetEmailSubjectMetadata gets a reference to the given PropertyMetadata and assigns it to the EmailSubjectMetadata field.
func (o *RecipientEmailNotification) SetEmailSubjectMetadata(v PropertyMetadata) {
	o.EmailSubjectMetadata = &v
}

// GetSupportedLanguage returns the SupportedLanguage field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetSupportedLanguage() string {
	if o == nil || IsNil(o.SupportedLanguage) {
		var ret string
		return ret
	}
	return *o.SupportedLanguage
}

// GetSupportedLanguageOk returns a tuple with the SupportedLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetSupportedLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedLanguage) {
		return nil, false
	}
	return o.SupportedLanguage, true
}

// HasSupportedLanguage returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasSupportedLanguage() bool {
	if o != nil && !IsNil(o.SupportedLanguage) {
		return true
	}

	return false
}

// SetSupportedLanguage gets a reference to the given string and assigns it to the SupportedLanguage field.
func (o *RecipientEmailNotification) SetSupportedLanguage(v string) {
	o.SupportedLanguage = &v
}

// GetSupportedLanguageMetadata returns the SupportedLanguageMetadata field value if set, zero value otherwise.
func (o *RecipientEmailNotification) GetSupportedLanguageMetadata() PropertyMetadata {
	if o == nil || IsNil(o.SupportedLanguageMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.SupportedLanguageMetadata
}

// GetSupportedLanguageMetadataOk returns a tuple with the SupportedLanguageMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientEmailNotification) GetSupportedLanguageMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.SupportedLanguageMetadata) {
		return nil, false
	}
	return o.SupportedLanguageMetadata, true
}

// HasSupportedLanguageMetadata returns a boolean if a field has been set.
func (o *RecipientEmailNotification) HasSupportedLanguageMetadata() bool {
	if o != nil && !IsNil(o.SupportedLanguageMetadata) {
		return true
	}

	return false
}

// SetSupportedLanguageMetadata gets a reference to the given PropertyMetadata and assigns it to the SupportedLanguageMetadata field.
func (o *RecipientEmailNotification) SetSupportedLanguageMetadata(v PropertyMetadata) {
	o.SupportedLanguageMetadata = &v
}

func (o RecipientEmailNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientEmailNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailBody) {
		toSerialize["emailBody"] = o.EmailBody
	}
	if !IsNil(o.EmailBodyMetadata) {
		toSerialize["emailBodyMetadata"] = o.EmailBodyMetadata
	}
	if !IsNil(o.EmailSubject) {
		toSerialize["emailSubject"] = o.EmailSubject
	}
	if !IsNil(o.EmailSubjectMetadata) {
		toSerialize["emailSubjectMetadata"] = o.EmailSubjectMetadata
	}
	if !IsNil(o.SupportedLanguage) {
		toSerialize["supportedLanguage"] = o.SupportedLanguage
	}
	if !IsNil(o.SupportedLanguageMetadata) {
		toSerialize["supportedLanguageMetadata"] = o.SupportedLanguageMetadata
	}
	return toSerialize, nil
}

type NullableRecipientEmailNotification struct {
	value *RecipientEmailNotification
	isSet bool
}

func (v NullableRecipientEmailNotification) Get() *RecipientEmailNotification {
	return v.value
}

func (v *NullableRecipientEmailNotification) Set(val *RecipientEmailNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientEmailNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientEmailNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientEmailNotification(val *RecipientEmailNotification) *NullableRecipientEmailNotification {
	return &NullableRecipientEmailNotification{value: val, isSet: true}
}

func (v NullableRecipientEmailNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientEmailNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


