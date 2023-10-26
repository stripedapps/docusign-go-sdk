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

// checks if the FolderItemV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FolderItemV2{}

// FolderItemV2 Information about folder item results.
type FolderItemV2 struct {
	// If the item is an envelope, this is the UTC DateTime when the envelope was completed.
	CompletedDateTime *string `json:"completedDateTime,omitempty"`
	// The UTC DateTime when the item was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	// If the item is an envelope, this is the ID of the envelope.
	EnvelopeId *string `json:"envelopeId,omitempty"`
	// If the item is an envelope, this is the URI for retrieving it.
	EnvelopeUri *string `json:"envelopeUri,omitempty"`
	// The date and time the envelope is set to expire.
	ExpireDateTime *string `json:"expireDateTime,omitempty"`
	// The ID of the folder.
	FolderId *string `json:"folderId,omitempty"`
	// If the item is a subfolder, this is the URI for retrieving it.
	FolderUri *string `json:"folderUri,omitempty"`
	// When **true,** indicates compliance with United States Food and Drug Administration (FDA) regulations on electronic records and electronic signatures (ERES).
	Is21CFRPart11 *string `json:"is21CFRPart11,omitempty"`
	// The date and time that the item was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	// The name of the user who owns the folder.
	OwnerName *string `json:"ownerName,omitempty"`
	Recipients *EnvelopeRecipients `json:"recipients,omitempty"`
	// Contains a URI for an endpoint that you can use to retrieve the recipients.
	RecipientsUri *string `json:"recipientsUri,omitempty"`
	// The name of the sender's company.
	SenderCompany *string `json:"senderCompany,omitempty"`
	// The sender's email address.
	SenderEmail *string `json:"senderEmail,omitempty"`
	// The sender's name.
	SenderName *string `json:"senderName,omitempty"`
	// The sender's id.
	SenderUserId *string `json:"senderUserId,omitempty"`
	// The UTC DateTime when the envelope was sent. This property is read-only.
	SentDateTime *string `json:"sentDateTime,omitempty"`
	// Indicates the envelope status. Valid values are:  * sent - The envelope is sent to the recipients.  * created - The envelope is saved as a draft and can be modified and sent later.
	Status *string `json:"status,omitempty"`
	// The subject of the envelope.
	Subject *string `json:"subject,omitempty"`
	// The unique identifier of the template. If this is not provided, DocuSign will generate a value. 
	TemplateId *string `json:"templateId,omitempty"`
	// The URI for retrieving the template.
	TemplateUri *string `json:"templateUri,omitempty"`
}

// NewFolderItemV2 instantiates a new FolderItemV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFolderItemV2() *FolderItemV2 {
	this := FolderItemV2{}
	return &this
}

// NewFolderItemV2WithDefaults instantiates a new FolderItemV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFolderItemV2WithDefaults() *FolderItemV2 {
	this := FolderItemV2{}
	return &this
}

// GetCompletedDateTime returns the CompletedDateTime field value if set, zero value otherwise.
func (o *FolderItemV2) GetCompletedDateTime() string {
	if o == nil || IsNil(o.CompletedDateTime) {
		var ret string
		return ret
	}
	return *o.CompletedDateTime
}

// GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetCompletedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CompletedDateTime) {
		return nil, false
	}
	return o.CompletedDateTime, true
}

// HasCompletedDateTime returns a boolean if a field has been set.
func (o *FolderItemV2) HasCompletedDateTime() bool {
	if o != nil && !IsNil(o.CompletedDateTime) {
		return true
	}

	return false
}

// SetCompletedDateTime gets a reference to the given string and assigns it to the CompletedDateTime field.
func (o *FolderItemV2) SetCompletedDateTime(v string) {
	o.CompletedDateTime = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *FolderItemV2) GetCreatedDateTime() string {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *FolderItemV2) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *FolderItemV2) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetEnvelopeId returns the EnvelopeId field value if set, zero value otherwise.
func (o *FolderItemV2) GetEnvelopeId() string {
	if o == nil || IsNil(o.EnvelopeId) {
		var ret string
		return ret
	}
	return *o.EnvelopeId
}

// GetEnvelopeIdOk returns a tuple with the EnvelopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetEnvelopeIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeId) {
		return nil, false
	}
	return o.EnvelopeId, true
}

// HasEnvelopeId returns a boolean if a field has been set.
func (o *FolderItemV2) HasEnvelopeId() bool {
	if o != nil && !IsNil(o.EnvelopeId) {
		return true
	}

	return false
}

// SetEnvelopeId gets a reference to the given string and assigns it to the EnvelopeId field.
func (o *FolderItemV2) SetEnvelopeId(v string) {
	o.EnvelopeId = &v
}

// GetEnvelopeUri returns the EnvelopeUri field value if set, zero value otherwise.
func (o *FolderItemV2) GetEnvelopeUri() string {
	if o == nil || IsNil(o.EnvelopeUri) {
		var ret string
		return ret
	}
	return *o.EnvelopeUri
}

// GetEnvelopeUriOk returns a tuple with the EnvelopeUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetEnvelopeUriOk() (*string, bool) {
	if o == nil || IsNil(o.EnvelopeUri) {
		return nil, false
	}
	return o.EnvelopeUri, true
}

// HasEnvelopeUri returns a boolean if a field has been set.
func (o *FolderItemV2) HasEnvelopeUri() bool {
	if o != nil && !IsNil(o.EnvelopeUri) {
		return true
	}

	return false
}

// SetEnvelopeUri gets a reference to the given string and assigns it to the EnvelopeUri field.
func (o *FolderItemV2) SetEnvelopeUri(v string) {
	o.EnvelopeUri = &v
}

// GetExpireDateTime returns the ExpireDateTime field value if set, zero value otherwise.
func (o *FolderItemV2) GetExpireDateTime() string {
	if o == nil || IsNil(o.ExpireDateTime) {
		var ret string
		return ret
	}
	return *o.ExpireDateTime
}

// GetExpireDateTimeOk returns a tuple with the ExpireDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetExpireDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireDateTime) {
		return nil, false
	}
	return o.ExpireDateTime, true
}

// HasExpireDateTime returns a boolean if a field has been set.
func (o *FolderItemV2) HasExpireDateTime() bool {
	if o != nil && !IsNil(o.ExpireDateTime) {
		return true
	}

	return false
}

// SetExpireDateTime gets a reference to the given string and assigns it to the ExpireDateTime field.
func (o *FolderItemV2) SetExpireDateTime(v string) {
	o.ExpireDateTime = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *FolderItemV2) GetFolderId() string {
	if o == nil || IsNil(o.FolderId) {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetFolderIdOk() (*string, bool) {
	if o == nil || IsNil(o.FolderId) {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *FolderItemV2) HasFolderId() bool {
	if o != nil && !IsNil(o.FolderId) {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *FolderItemV2) SetFolderId(v string) {
	o.FolderId = &v
}

// GetFolderUri returns the FolderUri field value if set, zero value otherwise.
func (o *FolderItemV2) GetFolderUri() string {
	if o == nil || IsNil(o.FolderUri) {
		var ret string
		return ret
	}
	return *o.FolderUri
}

// GetFolderUriOk returns a tuple with the FolderUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetFolderUriOk() (*string, bool) {
	if o == nil || IsNil(o.FolderUri) {
		return nil, false
	}
	return o.FolderUri, true
}

// HasFolderUri returns a boolean if a field has been set.
func (o *FolderItemV2) HasFolderUri() bool {
	if o != nil && !IsNil(o.FolderUri) {
		return true
	}

	return false
}

// SetFolderUri gets a reference to the given string and assigns it to the FolderUri field.
func (o *FolderItemV2) SetFolderUri(v string) {
	o.FolderUri = &v
}

// GetIs21CFRPart11 returns the Is21CFRPart11 field value if set, zero value otherwise.
func (o *FolderItemV2) GetIs21CFRPart11() string {
	if o == nil || IsNil(o.Is21CFRPart11) {
		var ret string
		return ret
	}
	return *o.Is21CFRPart11
}

// GetIs21CFRPart11Ok returns a tuple with the Is21CFRPart11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetIs21CFRPart11Ok() (*string, bool) {
	if o == nil || IsNil(o.Is21CFRPart11) {
		return nil, false
	}
	return o.Is21CFRPart11, true
}

// HasIs21CFRPart11 returns a boolean if a field has been set.
func (o *FolderItemV2) HasIs21CFRPart11() bool {
	if o != nil && !IsNil(o.Is21CFRPart11) {
		return true
	}

	return false
}

// SetIs21CFRPart11 gets a reference to the given string and assigns it to the Is21CFRPart11 field.
func (o *FolderItemV2) SetIs21CFRPart11(v string) {
	o.Is21CFRPart11 = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *FolderItemV2) GetLastModifiedDateTime() string {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		var ret string
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetLastModifiedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *FolderItemV2) HasLastModifiedDateTime() bool {
	if o != nil && !IsNil(o.LastModifiedDateTime) {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given string and assigns it to the LastModifiedDateTime field.
func (o *FolderItemV2) SetLastModifiedDateTime(v string) {
	o.LastModifiedDateTime = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *FolderItemV2) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *FolderItemV2) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *FolderItemV2) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *FolderItemV2) GetRecipients() EnvelopeRecipients {
	if o == nil || IsNil(o.Recipients) {
		var ret EnvelopeRecipients
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetRecipientsOk() (*EnvelopeRecipients, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *FolderItemV2) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given EnvelopeRecipients and assigns it to the Recipients field.
func (o *FolderItemV2) SetRecipients(v EnvelopeRecipients) {
	o.Recipients = &v
}

// GetRecipientsUri returns the RecipientsUri field value if set, zero value otherwise.
func (o *FolderItemV2) GetRecipientsUri() string {
	if o == nil || IsNil(o.RecipientsUri) {
		var ret string
		return ret
	}
	return *o.RecipientsUri
}

// GetRecipientsUriOk returns a tuple with the RecipientsUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetRecipientsUriOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientsUri) {
		return nil, false
	}
	return o.RecipientsUri, true
}

// HasRecipientsUri returns a boolean if a field has been set.
func (o *FolderItemV2) HasRecipientsUri() bool {
	if o != nil && !IsNil(o.RecipientsUri) {
		return true
	}

	return false
}

// SetRecipientsUri gets a reference to the given string and assigns it to the RecipientsUri field.
func (o *FolderItemV2) SetRecipientsUri(v string) {
	o.RecipientsUri = &v
}

// GetSenderCompany returns the SenderCompany field value if set, zero value otherwise.
func (o *FolderItemV2) GetSenderCompany() string {
	if o == nil || IsNil(o.SenderCompany) {
		var ret string
		return ret
	}
	return *o.SenderCompany
}

// GetSenderCompanyOk returns a tuple with the SenderCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSenderCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.SenderCompany) {
		return nil, false
	}
	return o.SenderCompany, true
}

// HasSenderCompany returns a boolean if a field has been set.
func (o *FolderItemV2) HasSenderCompany() bool {
	if o != nil && !IsNil(o.SenderCompany) {
		return true
	}

	return false
}

// SetSenderCompany gets a reference to the given string and assigns it to the SenderCompany field.
func (o *FolderItemV2) SetSenderCompany(v string) {
	o.SenderCompany = &v
}

// GetSenderEmail returns the SenderEmail field value if set, zero value otherwise.
func (o *FolderItemV2) GetSenderEmail() string {
	if o == nil || IsNil(o.SenderEmail) {
		var ret string
		return ret
	}
	return *o.SenderEmail
}

// GetSenderEmailOk returns a tuple with the SenderEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSenderEmailOk() (*string, bool) {
	if o == nil || IsNil(o.SenderEmail) {
		return nil, false
	}
	return o.SenderEmail, true
}

// HasSenderEmail returns a boolean if a field has been set.
func (o *FolderItemV2) HasSenderEmail() bool {
	if o != nil && !IsNil(o.SenderEmail) {
		return true
	}

	return false
}

// SetSenderEmail gets a reference to the given string and assigns it to the SenderEmail field.
func (o *FolderItemV2) SetSenderEmail(v string) {
	o.SenderEmail = &v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *FolderItemV2) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *FolderItemV2) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *FolderItemV2) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSenderUserId returns the SenderUserId field value if set, zero value otherwise.
func (o *FolderItemV2) GetSenderUserId() string {
	if o == nil || IsNil(o.SenderUserId) {
		var ret string
		return ret
	}
	return *o.SenderUserId
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSenderUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.SenderUserId) {
		return nil, false
	}
	return o.SenderUserId, true
}

// HasSenderUserId returns a boolean if a field has been set.
func (o *FolderItemV2) HasSenderUserId() bool {
	if o != nil && !IsNil(o.SenderUserId) {
		return true
	}

	return false
}

// SetSenderUserId gets a reference to the given string and assigns it to the SenderUserId field.
func (o *FolderItemV2) SetSenderUserId(v string) {
	o.SenderUserId = &v
}

// GetSentDateTime returns the SentDateTime field value if set, zero value otherwise.
func (o *FolderItemV2) GetSentDateTime() string {
	if o == nil || IsNil(o.SentDateTime) {
		var ret string
		return ret
	}
	return *o.SentDateTime
}

// GetSentDateTimeOk returns a tuple with the SentDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSentDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.SentDateTime) {
		return nil, false
	}
	return o.SentDateTime, true
}

// HasSentDateTime returns a boolean if a field has been set.
func (o *FolderItemV2) HasSentDateTime() bool {
	if o != nil && !IsNil(o.SentDateTime) {
		return true
	}

	return false
}

// SetSentDateTime gets a reference to the given string and assigns it to the SentDateTime field.
func (o *FolderItemV2) SetSentDateTime(v string) {
	o.SentDateTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FolderItemV2) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FolderItemV2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FolderItemV2) SetStatus(v string) {
	o.Status = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *FolderItemV2) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *FolderItemV2) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *FolderItemV2) SetSubject(v string) {
	o.Subject = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *FolderItemV2) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *FolderItemV2) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *FolderItemV2) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateUri returns the TemplateUri field value if set, zero value otherwise.
func (o *FolderItemV2) GetTemplateUri() string {
	if o == nil || IsNil(o.TemplateUri) {
		var ret string
		return ret
	}
	return *o.TemplateUri
}

// GetTemplateUriOk returns a tuple with the TemplateUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FolderItemV2) GetTemplateUriOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateUri) {
		return nil, false
	}
	return o.TemplateUri, true
}

// HasTemplateUri returns a boolean if a field has been set.
func (o *FolderItemV2) HasTemplateUri() bool {
	if o != nil && !IsNil(o.TemplateUri) {
		return true
	}

	return false
}

// SetTemplateUri gets a reference to the given string and assigns it to the TemplateUri field.
func (o *FolderItemV2) SetTemplateUri(v string) {
	o.TemplateUri = &v
}

func (o FolderItemV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FolderItemV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CompletedDateTime) {
		toSerialize["completedDateTime"] = o.CompletedDateTime
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.EnvelopeId) {
		toSerialize["envelopeId"] = o.EnvelopeId
	}
	if !IsNil(o.EnvelopeUri) {
		toSerialize["envelopeUri"] = o.EnvelopeUri
	}
	if !IsNil(o.ExpireDateTime) {
		toSerialize["expireDateTime"] = o.ExpireDateTime
	}
	if !IsNil(o.FolderId) {
		toSerialize["folderId"] = o.FolderId
	}
	if !IsNil(o.FolderUri) {
		toSerialize["folderUri"] = o.FolderUri
	}
	if !IsNil(o.Is21CFRPart11) {
		toSerialize["is21CFRPart11"] = o.Is21CFRPart11
	}
	if !IsNil(o.LastModifiedDateTime) {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.RecipientsUri) {
		toSerialize["recipientsUri"] = o.RecipientsUri
	}
	if !IsNil(o.SenderCompany) {
		toSerialize["senderCompany"] = o.SenderCompany
	}
	if !IsNil(o.SenderEmail) {
		toSerialize["senderEmail"] = o.SenderEmail
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	if !IsNil(o.SenderUserId) {
		toSerialize["senderUserId"] = o.SenderUserId
	}
	if !IsNil(o.SentDateTime) {
		toSerialize["sentDateTime"] = o.SentDateTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.TemplateUri) {
		toSerialize["templateUri"] = o.TemplateUri
	}
	return toSerialize, nil
}

type NullableFolderItemV2 struct {
	value *FolderItemV2
	isSet bool
}

func (v NullableFolderItemV2) Get() *FolderItemV2 {
	return v.value
}

func (v *NullableFolderItemV2) Set(val *FolderItemV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFolderItemV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFolderItemV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFolderItemV2(val *FolderItemV2) *NullableFolderItemV2 {
	return &NullableFolderItemV2{value: val, isSet: true}
}

func (v NullableFolderItemV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFolderItemV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

