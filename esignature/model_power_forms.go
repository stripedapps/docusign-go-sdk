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

// checks if the PowerForms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerForms{}

// PowerForms The PowerForms resource enables you to create fillable forms that you can email or make available for self service on the web.
type PowerForms struct {
	// The ID of the user who created the PowerForm. This property is returned in a response only when you set the `include_created_by` query parameter to **true.**
	CreatedBy *string `json:"createdBy,omitempty"`
	// The date and time that the PowerForm was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	// For a PowerForm that is sent by email, this is the body of the email message sent to the recipients.  Maximum length: 10000 characters.
	EmailBody *string `json:"emailBody,omitempty"`
	// Sets the envelope name for the envelopes that the PowerForm generates. One option is to make this property the same as the subject from the template.  You can customize the subject line to include a recipient's name or email address by using merge fields. For information about adding merge fields to the email subject, see [Template Email Subject Merge Fields](/docs/esign-rest-api/reference/templates/templates/create/). 
	EmailSubject *string `json:"emailSubject,omitempty"`
	// An array of envelope objects that contain information about the envelopes that are associated with the PowerForm.
	Envelopes []Envelope `json:"envelopes,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The instructions that display on the landing page for the first recipient. These instructions are important if the recipient accesses the PowerForm by a method other than email. When you include instructions, they display as an introduction after the recipient accesses the PowerForm.
	Instructions *string `json:"instructions,omitempty"`
	// When **true,** indicates that the PowerForm is active and can be sent to recipients. This is the default value.   When **false,** the PowerForm cannot be emailed or accessed by a recipient, even if they arrive at the PowerForm URL.   If a recipient attempts to sign an inactive PowerForm, an error message informs the recipient that the document is not active and suggests that they contact the sender.
	IsActive *string `json:"isActive,omitempty"`
	// The date and time that the PowerForm was last used.
	LastUsed *string `json:"lastUsed,omitempty"`
	// The length of time before the same recipient can sign the same PowerForm again. This property is used in combination with the `limitUseIntervalUnits` property.
	LimitUseInterval *string `json:"limitUseInterval,omitempty"`
	// When **true,** the `limitUseInterval` is enabled.
	LimitUseIntervalEnabled *string `json:"limitUseIntervalEnabled,omitempty"`
	// The units associated with the `limitUseInterval`. Valid values are:  - `minutes` - `hours` - `days` - `weeks` - `months`  For example, to limit a recipient to signing once per year, set the `limitUseInterval` to 365 and the `limitUseIntervalUnits` to `days`. 
	LimitUseIntervalUnits *string `json:"limitUseIntervalUnits,omitempty"`
	// When **true,** you can set a maximum number of uses for the PowerForm.
	MaxUseEnabled *string `json:"maxUseEnabled,omitempty"`
	// The name of the PowerForm.
	Name *string `json:"name,omitempty"`
	// The ID of the PowerForm.
	PowerFormId *string `json:"powerFormId,omitempty"`
	// The URL for the PowerForm.
	PowerFormUrl *string `json:"powerFormUrl,omitempty"`
	// An array of `powerFormRecipient` objects.  **Note:** For self-service documents where you do not know who the recipients are in advance, you can enter generic information for the `role` property and leave other details (such as `name` and `email`) blank.
	Recipients []PowerFormRecipient `json:"recipients,omitempty"`
	// The name of the sender.   **Note:** The default sender for a PowerForm is the PowerForm Administrator who created it.
	SenderName *string `json:"senderName,omitempty"`
	// The ID of the sender.
	SenderUserId *string `json:"senderUserId,omitempty"`
	// The signing method to use. Valid values are:  - `email`: This mode verifies the recipient's identity by using email authentication before the recipient can sign a document.  - `direct`: This mode does not require any verification. DocuSign recommends that you use this signing method only when another form of authentication is in use.  **Note:** In the account settings, `enablePowerFormDirect` must be **true** to use `direct` as the `signingMode`.  For more information about signing modes, see the [overview of the Create method](/docs/esign-rest-api/reference/powerforms/powerforms/create/).
	SigningMode *string `json:"signingMode,omitempty"`
	// The ID of the template used to create the PowerForm.
	TemplateId *string `json:"templateId,omitempty"`
	// The name of the template used to create the PowerForm.
	TemplateName *string `json:"templateName,omitempty"`
	// The number of times the PowerForm has been used. 
	TimesUsed *string `json:"timesUsed,omitempty"`
	// The URI for the PowerForm.
	Uri *string `json:"uri,omitempty"`
	// The number of times that the PowerForm can still be used. If no use limit is set, the value is `Unlimited`. 
	UsesRemaining *string `json:"usesRemaining,omitempty"`
}

// NewPowerForms instantiates a new PowerForms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerForms() *PowerForms {
	this := PowerForms{}
	return &this
}

// NewPowerFormsWithDefaults instantiates a new PowerForms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFormsWithDefaults() *PowerForms {
	this := PowerForms{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *PowerForms) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PowerForms) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *PowerForms) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *PowerForms) GetCreatedDateTime() string {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *PowerForms) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *PowerForms) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetEmailBody returns the EmailBody field value if set, zero value otherwise.
func (o *PowerForms) GetEmailBody() string {
	if o == nil || IsNil(o.EmailBody) {
		var ret string
		return ret
	}
	return *o.EmailBody
}

// GetEmailBodyOk returns a tuple with the EmailBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetEmailBodyOk() (*string, bool) {
	if o == nil || IsNil(o.EmailBody) {
		return nil, false
	}
	return o.EmailBody, true
}

// HasEmailBody returns a boolean if a field has been set.
func (o *PowerForms) HasEmailBody() bool {
	if o != nil && !IsNil(o.EmailBody) {
		return true
	}

	return false
}

// SetEmailBody gets a reference to the given string and assigns it to the EmailBody field.
func (o *PowerForms) SetEmailBody(v string) {
	o.EmailBody = &v
}

// GetEmailSubject returns the EmailSubject field value if set, zero value otherwise.
func (o *PowerForms) GetEmailSubject() string {
	if o == nil || IsNil(o.EmailSubject) {
		var ret string
		return ret
	}
	return *o.EmailSubject
}

// GetEmailSubjectOk returns a tuple with the EmailSubject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetEmailSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.EmailSubject) {
		return nil, false
	}
	return o.EmailSubject, true
}

// HasEmailSubject returns a boolean if a field has been set.
func (o *PowerForms) HasEmailSubject() bool {
	if o != nil && !IsNil(o.EmailSubject) {
		return true
	}

	return false
}

// SetEmailSubject gets a reference to the given string and assigns it to the EmailSubject field.
func (o *PowerForms) SetEmailSubject(v string) {
	o.EmailSubject = &v
}

// GetEnvelopes returns the Envelopes field value if set, zero value otherwise.
func (o *PowerForms) GetEnvelopes() []Envelope {
	if o == nil || IsNil(o.Envelopes) {
		var ret []Envelope
		return ret
	}
	return o.Envelopes
}

// GetEnvelopesOk returns a tuple with the Envelopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetEnvelopesOk() ([]Envelope, bool) {
	if o == nil || IsNil(o.Envelopes) {
		return nil, false
	}
	return o.Envelopes, true
}

// HasEnvelopes returns a boolean if a field has been set.
func (o *PowerForms) HasEnvelopes() bool {
	if o != nil && !IsNil(o.Envelopes) {
		return true
	}

	return false
}

// SetEnvelopes gets a reference to the given []Envelope and assigns it to the Envelopes field.
func (o *PowerForms) SetEnvelopes(v []Envelope) {
	o.Envelopes = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *PowerForms) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *PowerForms) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *PowerForms) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *PowerForms) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *PowerForms) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *PowerForms) SetInstructions(v string) {
	o.Instructions = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *PowerForms) GetIsActive() string {
	if o == nil || IsNil(o.IsActive) {
		var ret string
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetIsActiveOk() (*string, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *PowerForms) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given string and assigns it to the IsActive field.
func (o *PowerForms) SetIsActive(v string) {
	o.IsActive = &v
}

// GetLastUsed returns the LastUsed field value if set, zero value otherwise.
func (o *PowerForms) GetLastUsed() string {
	if o == nil || IsNil(o.LastUsed) {
		var ret string
		return ret
	}
	return *o.LastUsed
}

// GetLastUsedOk returns a tuple with the LastUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetLastUsedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsed) {
		return nil, false
	}
	return o.LastUsed, true
}

// HasLastUsed returns a boolean if a field has been set.
func (o *PowerForms) HasLastUsed() bool {
	if o != nil && !IsNil(o.LastUsed) {
		return true
	}

	return false
}

// SetLastUsed gets a reference to the given string and assigns it to the LastUsed field.
func (o *PowerForms) SetLastUsed(v string) {
	o.LastUsed = &v
}

// GetLimitUseInterval returns the LimitUseInterval field value if set, zero value otherwise.
func (o *PowerForms) GetLimitUseInterval() string {
	if o == nil || IsNil(o.LimitUseInterval) {
		var ret string
		return ret
	}
	return *o.LimitUseInterval
}

// GetLimitUseIntervalOk returns a tuple with the LimitUseInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetLimitUseIntervalOk() (*string, bool) {
	if o == nil || IsNil(o.LimitUseInterval) {
		return nil, false
	}
	return o.LimitUseInterval, true
}

// HasLimitUseInterval returns a boolean if a field has been set.
func (o *PowerForms) HasLimitUseInterval() bool {
	if o != nil && !IsNil(o.LimitUseInterval) {
		return true
	}

	return false
}

// SetLimitUseInterval gets a reference to the given string and assigns it to the LimitUseInterval field.
func (o *PowerForms) SetLimitUseInterval(v string) {
	o.LimitUseInterval = &v
}

// GetLimitUseIntervalEnabled returns the LimitUseIntervalEnabled field value if set, zero value otherwise.
func (o *PowerForms) GetLimitUseIntervalEnabled() string {
	if o == nil || IsNil(o.LimitUseIntervalEnabled) {
		var ret string
		return ret
	}
	return *o.LimitUseIntervalEnabled
}

// GetLimitUseIntervalEnabledOk returns a tuple with the LimitUseIntervalEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetLimitUseIntervalEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.LimitUseIntervalEnabled) {
		return nil, false
	}
	return o.LimitUseIntervalEnabled, true
}

// HasLimitUseIntervalEnabled returns a boolean if a field has been set.
func (o *PowerForms) HasLimitUseIntervalEnabled() bool {
	if o != nil && !IsNil(o.LimitUseIntervalEnabled) {
		return true
	}

	return false
}

// SetLimitUseIntervalEnabled gets a reference to the given string and assigns it to the LimitUseIntervalEnabled field.
func (o *PowerForms) SetLimitUseIntervalEnabled(v string) {
	o.LimitUseIntervalEnabled = &v
}

// GetLimitUseIntervalUnits returns the LimitUseIntervalUnits field value if set, zero value otherwise.
func (o *PowerForms) GetLimitUseIntervalUnits() string {
	if o == nil || IsNil(o.LimitUseIntervalUnits) {
		var ret string
		return ret
	}
	return *o.LimitUseIntervalUnits
}

// GetLimitUseIntervalUnitsOk returns a tuple with the LimitUseIntervalUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetLimitUseIntervalUnitsOk() (*string, bool) {
	if o == nil || IsNil(o.LimitUseIntervalUnits) {
		return nil, false
	}
	return o.LimitUseIntervalUnits, true
}

// HasLimitUseIntervalUnits returns a boolean if a field has been set.
func (o *PowerForms) HasLimitUseIntervalUnits() bool {
	if o != nil && !IsNil(o.LimitUseIntervalUnits) {
		return true
	}

	return false
}

// SetLimitUseIntervalUnits gets a reference to the given string and assigns it to the LimitUseIntervalUnits field.
func (o *PowerForms) SetLimitUseIntervalUnits(v string) {
	o.LimitUseIntervalUnits = &v
}

// GetMaxUseEnabled returns the MaxUseEnabled field value if set, zero value otherwise.
func (o *PowerForms) GetMaxUseEnabled() string {
	if o == nil || IsNil(o.MaxUseEnabled) {
		var ret string
		return ret
	}
	return *o.MaxUseEnabled
}

// GetMaxUseEnabledOk returns a tuple with the MaxUseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetMaxUseEnabledOk() (*string, bool) {
	if o == nil || IsNil(o.MaxUseEnabled) {
		return nil, false
	}
	return o.MaxUseEnabled, true
}

// HasMaxUseEnabled returns a boolean if a field has been set.
func (o *PowerForms) HasMaxUseEnabled() bool {
	if o != nil && !IsNil(o.MaxUseEnabled) {
		return true
	}

	return false
}

// SetMaxUseEnabled gets a reference to the given string and assigns it to the MaxUseEnabled field.
func (o *PowerForms) SetMaxUseEnabled(v string) {
	o.MaxUseEnabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PowerForms) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PowerForms) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PowerForms) SetName(v string) {
	o.Name = &v
}

// GetPowerFormId returns the PowerFormId field value if set, zero value otherwise.
func (o *PowerForms) GetPowerFormId() string {
	if o == nil || IsNil(o.PowerFormId) {
		var ret string
		return ret
	}
	return *o.PowerFormId
}

// GetPowerFormIdOk returns a tuple with the PowerFormId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetPowerFormIdOk() (*string, bool) {
	if o == nil || IsNil(o.PowerFormId) {
		return nil, false
	}
	return o.PowerFormId, true
}

// HasPowerFormId returns a boolean if a field has been set.
func (o *PowerForms) HasPowerFormId() bool {
	if o != nil && !IsNil(o.PowerFormId) {
		return true
	}

	return false
}

// SetPowerFormId gets a reference to the given string and assigns it to the PowerFormId field.
func (o *PowerForms) SetPowerFormId(v string) {
	o.PowerFormId = &v
}

// GetPowerFormUrl returns the PowerFormUrl field value if set, zero value otherwise.
func (o *PowerForms) GetPowerFormUrl() string {
	if o == nil || IsNil(o.PowerFormUrl) {
		var ret string
		return ret
	}
	return *o.PowerFormUrl
}

// GetPowerFormUrlOk returns a tuple with the PowerFormUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetPowerFormUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PowerFormUrl) {
		return nil, false
	}
	return o.PowerFormUrl, true
}

// HasPowerFormUrl returns a boolean if a field has been set.
func (o *PowerForms) HasPowerFormUrl() bool {
	if o != nil && !IsNil(o.PowerFormUrl) {
		return true
	}

	return false
}

// SetPowerFormUrl gets a reference to the given string and assigns it to the PowerFormUrl field.
func (o *PowerForms) SetPowerFormUrl(v string) {
	o.PowerFormUrl = &v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *PowerForms) GetRecipients() []PowerFormRecipient {
	if o == nil || IsNil(o.Recipients) {
		var ret []PowerFormRecipient
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetRecipientsOk() ([]PowerFormRecipient, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *PowerForms) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []PowerFormRecipient and assigns it to the Recipients field.
func (o *PowerForms) SetRecipients(v []PowerFormRecipient) {
	o.Recipients = v
}

// GetSenderName returns the SenderName field value if set, zero value otherwise.
func (o *PowerForms) GetSenderName() string {
	if o == nil || IsNil(o.SenderName) {
		var ret string
		return ret
	}
	return *o.SenderName
}

// GetSenderNameOk returns a tuple with the SenderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetSenderNameOk() (*string, bool) {
	if o == nil || IsNil(o.SenderName) {
		return nil, false
	}
	return o.SenderName, true
}

// HasSenderName returns a boolean if a field has been set.
func (o *PowerForms) HasSenderName() bool {
	if o != nil && !IsNil(o.SenderName) {
		return true
	}

	return false
}

// SetSenderName gets a reference to the given string and assigns it to the SenderName field.
func (o *PowerForms) SetSenderName(v string) {
	o.SenderName = &v
}

// GetSenderUserId returns the SenderUserId field value if set, zero value otherwise.
func (o *PowerForms) GetSenderUserId() string {
	if o == nil || IsNil(o.SenderUserId) {
		var ret string
		return ret
	}
	return *o.SenderUserId
}

// GetSenderUserIdOk returns a tuple with the SenderUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetSenderUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.SenderUserId) {
		return nil, false
	}
	return o.SenderUserId, true
}

// HasSenderUserId returns a boolean if a field has been set.
func (o *PowerForms) HasSenderUserId() bool {
	if o != nil && !IsNil(o.SenderUserId) {
		return true
	}

	return false
}

// SetSenderUserId gets a reference to the given string and assigns it to the SenderUserId field.
func (o *PowerForms) SetSenderUserId(v string) {
	o.SenderUserId = &v
}

// GetSigningMode returns the SigningMode field value if set, zero value otherwise.
func (o *PowerForms) GetSigningMode() string {
	if o == nil || IsNil(o.SigningMode) {
		var ret string
		return ret
	}
	return *o.SigningMode
}

// GetSigningModeOk returns a tuple with the SigningMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetSigningModeOk() (*string, bool) {
	if o == nil || IsNil(o.SigningMode) {
		return nil, false
	}
	return o.SigningMode, true
}

// HasSigningMode returns a boolean if a field has been set.
func (o *PowerForms) HasSigningMode() bool {
	if o != nil && !IsNil(o.SigningMode) {
		return true
	}

	return false
}

// SetSigningMode gets a reference to the given string and assigns it to the SigningMode field.
func (o *PowerForms) SetSigningMode(v string) {
	o.SigningMode = &v
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *PowerForms) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *PowerForms) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *PowerForms) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetTemplateName returns the TemplateName field value if set, zero value otherwise.
func (o *PowerForms) GetTemplateName() string {
	if o == nil || IsNil(o.TemplateName) {
		var ret string
		return ret
	}
	return *o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetTemplateNameOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateName) {
		return nil, false
	}
	return o.TemplateName, true
}

// HasTemplateName returns a boolean if a field has been set.
func (o *PowerForms) HasTemplateName() bool {
	if o != nil && !IsNil(o.TemplateName) {
		return true
	}

	return false
}

// SetTemplateName gets a reference to the given string and assigns it to the TemplateName field.
func (o *PowerForms) SetTemplateName(v string) {
	o.TemplateName = &v
}

// GetTimesUsed returns the TimesUsed field value if set, zero value otherwise.
func (o *PowerForms) GetTimesUsed() string {
	if o == nil || IsNil(o.TimesUsed) {
		var ret string
		return ret
	}
	return *o.TimesUsed
}

// GetTimesUsedOk returns a tuple with the TimesUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetTimesUsedOk() (*string, bool) {
	if o == nil || IsNil(o.TimesUsed) {
		return nil, false
	}
	return o.TimesUsed, true
}

// HasTimesUsed returns a boolean if a field has been set.
func (o *PowerForms) HasTimesUsed() bool {
	if o != nil && !IsNil(o.TimesUsed) {
		return true
	}

	return false
}

// SetTimesUsed gets a reference to the given string and assigns it to the TimesUsed field.
func (o *PowerForms) SetTimesUsed(v string) {
	o.TimesUsed = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *PowerForms) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *PowerForms) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *PowerForms) SetUri(v string) {
	o.Uri = &v
}

// GetUsesRemaining returns the UsesRemaining field value if set, zero value otherwise.
func (o *PowerForms) GetUsesRemaining() string {
	if o == nil || IsNil(o.UsesRemaining) {
		var ret string
		return ret
	}
	return *o.UsesRemaining
}

// GetUsesRemainingOk returns a tuple with the UsesRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerForms) GetUsesRemainingOk() (*string, bool) {
	if o == nil || IsNil(o.UsesRemaining) {
		return nil, false
	}
	return o.UsesRemaining, true
}

// HasUsesRemaining returns a boolean if a field has been set.
func (o *PowerForms) HasUsesRemaining() bool {
	if o != nil && !IsNil(o.UsesRemaining) {
		return true
	}

	return false
}

// SetUsesRemaining gets a reference to the given string and assigns it to the UsesRemaining field.
func (o *PowerForms) SetUsesRemaining(v string) {
	o.UsesRemaining = &v
}

func (o PowerForms) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerForms) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.EmailBody) {
		toSerialize["emailBody"] = o.EmailBody
	}
	if !IsNil(o.EmailSubject) {
		toSerialize["emailSubject"] = o.EmailSubject
	}
	if !IsNil(o.Envelopes) {
		toSerialize["envelopes"] = o.Envelopes
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.IsActive) {
		toSerialize["isActive"] = o.IsActive
	}
	if !IsNil(o.LastUsed) {
		toSerialize["lastUsed"] = o.LastUsed
	}
	if !IsNil(o.LimitUseInterval) {
		toSerialize["limitUseInterval"] = o.LimitUseInterval
	}
	if !IsNil(o.LimitUseIntervalEnabled) {
		toSerialize["limitUseIntervalEnabled"] = o.LimitUseIntervalEnabled
	}
	if !IsNil(o.LimitUseIntervalUnits) {
		toSerialize["limitUseIntervalUnits"] = o.LimitUseIntervalUnits
	}
	if !IsNil(o.MaxUseEnabled) {
		toSerialize["maxUseEnabled"] = o.MaxUseEnabled
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PowerFormId) {
		toSerialize["powerFormId"] = o.PowerFormId
	}
	if !IsNil(o.PowerFormUrl) {
		toSerialize["powerFormUrl"] = o.PowerFormUrl
	}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.SenderName) {
		toSerialize["senderName"] = o.SenderName
	}
	if !IsNil(o.SenderUserId) {
		toSerialize["senderUserId"] = o.SenderUserId
	}
	if !IsNil(o.SigningMode) {
		toSerialize["signingMode"] = o.SigningMode
	}
	if !IsNil(o.TemplateId) {
		toSerialize["templateId"] = o.TemplateId
	}
	if !IsNil(o.TemplateName) {
		toSerialize["templateName"] = o.TemplateName
	}
	if !IsNil(o.TimesUsed) {
		toSerialize["timesUsed"] = o.TimesUsed
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.UsesRemaining) {
		toSerialize["usesRemaining"] = o.UsesRemaining
	}
	return toSerialize, nil
}

type NullablePowerForms struct {
	value *PowerForms
	isSet bool
}

func (v NullablePowerForms) Get() *PowerForms {
	return v.value
}

func (v *NullablePowerForms) Set(val *PowerForms) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerForms) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerForms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerForms(val *PowerForms) *NullablePowerForms {
	return &NullablePowerForms{value: val, isSet: true}
}

func (v NullablePowerForms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerForms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

