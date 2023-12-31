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

// checks if the BulkSendingCopyRecipient type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkSendingCopyRecipient{}

// BulkSendingCopyRecipient This object contains details about a bulk send recipient.
type BulkSendingCopyRecipient struct {
	// If a value is provided, the recipient must enter the value as the access code to view and sign the envelope.   Maximum Length: 50 characters and it must conform to the account's access code format setting.  If blank, but the signer `accessCode` property is set in the envelope, then that value is used.  If blank and the signer `accessCode` property is not set, then the access code is not required.
	AccessCode *string `json:"accessCode,omitempty"`
	// Specifies whether the recipient is embedded or remote.   If the `clientUserId` property is not null then the recipient is embedded. Use this field to associate the signer with their userId in your app. Authenticating the user is the responsibility of your app when you use embedded signing.  If the `clientUserId` property is set and either `SignerMustHaveAccount` or `SignerMustLoginToSign` property of the account settings is set to  **true,** an error is generated on sending.  **Note:** This property is not returned by the [listStatusChanges](/docs/esign-rest-api/reference/envelopes/envelopes/liststatuschanges/) endpoint.  Maximum length: 100 characters.  
	ClientUserId *string `json:"clientUserId,omitempty"`
	// An optional array of strings that allows the sender to provide custom data about the recipient. This information is returned in the envelope status but otherwise not used by DocuSign. Each customField string can be a maximum of 100 characters.
	CustomFields []string `json:"customFields,omitempty"`
	// The delivery method. One of:  - `email` - `fax` - `SMS` - `WhatsApp` - `offline`  The `SMS` and `WhatsApp` delivery methods are limited to `signer`, `carbonCopy`, and `certifiedDelivery` recipients.  **Related topics**  - [Using SMS delivery with the eSignature API][smsconcept] - [How to request a signature by SMS delivery][howto]  [smsconcept]: /docs/esign-rest-api/esign101/concepts/sms-delivery/using-sms-esignature/ [howto]: /docs/esign-rest-api/how-to/request-signature-sms/
	DeliveryMethod *string `json:"deliveryMethod,omitempty"`
	// The recipient's email address.
	Email *string `json:"email,omitempty"`
	EmailNotification *RecipientEmailNotification `json:"emailNotification,omitempty"`
	// Specifies a sender-provided valid URL string for redirecting an embedded recipient. When using this option, the embedded recipient still receives an email from DocuSign, just as a remote recipient would. When the document link in the email is clicked the recipient is redirected, through DocuSign, to the supplied URL to complete their actions. When routing to the URL, the sender's system (the server responding to the URL) must request a recipient token to launch a signing session.   When `SIGN_AT_DOCUSIGN`, the recipient is directed to an embedded signing or viewing process directly at DocuSign. The signing or viewing action is initiated by the DocuSign system and the transaction activity and Certificate of Completion records will reflect this. In all other ways the process is identical to an embedded signing or viewing operation launched by a partner.  It is important to understand that in a typical embedded workflow, the authentication of an embedded recipient is the responsibility of the sending application. DocuSign expects that senders will follow their own processes for establishing the recipient's identity. In this workflow the recipient goes through the sending application before the embedded signing or viewing process is initiated. However, when the sending application sets `EmbeddedRecipientStartURL=SIGN_AT_DOCUSIGN`, the recipient goes directly to the embedded signing or viewing process, bypassing the sending application and any authentication steps the sending application would use. In this case, DocuSign recommends that you use one of the normal DocuSign authentication features (Access Code, Phone Authentication, SMS Authentication, etc.) to verify the identity of the recipient.  If the `clientUserId` property is NOT set, and the `embeddedRecipientStartURL` is set, DocuSign will ignore the redirect URL and launch the standard signing process for the email recipient. Information can be appended to the embedded recipient start URL using merge fields. The available merge fields items are: `envelopeId`, `recipientId`, `recipientName`, `recipientEmail`, and `customFields`. The `customFields` property must be set for the recipient or envelope. The merge fields are enclosed in double brackets.   *Example*:   `http://senderHost/[[mergeField1]]/ beginSigningSession? [[mergeField2]]&[[mergeField3]]` 
	EmbeddedRecipientStartURL *string `json:"embeddedRecipientStartURL,omitempty"`
	// Reserved for DocuSign.
	FaxNumber *string `json:"faxNumber,omitempty"`
	// The email address of the signing host. This is the DocuSign user that is hosting the in-person signing session.  Required when `inPersonSigningType` is `inPersonSigner`. For eNotary flow, use `email` instead.  Maximum Length: 100 characters. 
	HostEmail *string `json:"hostEmail,omitempty"`
	// The name of the signing host. This is the DocuSign user that is hosting the in-person signing session.  Required when `inPersonSigningType` is `inPersonSigner`. For eNotary flow, use `name` instead.  Maximum Length: 100 characters. 
	HostName *string `json:"hostName,omitempty"`
	// The name of the authentication check to use. This value must match one of the authentication types that the account uses. The names of these authentication types appear in the web console sending interface in the Identify list for a recipient. This setting overrides any default authentication setting. Valid values are:  - `Phone Auth $`: The recipient must authenticate by using two-factor authentication (2FA). You provide the phone number to use for 2FA in the `phoneAuthentication` object. - `SMS Auth $`: The recipient must authenticate via SMS. You provide the phone number to use in the `smsAuthentication` object. - `ID Check $`: The  recipient must answer detailed security questions.   **Example:** Your account has ID Check and SMS Authentication available. In the web console Identify list, these appear as ID Check $ and SMS Auth $. To use ID Check in an envelope, the idCheckConfigurationName should be ID Check $. For SMS, you would use SMS Auth $, and you would also need to add a phone number to the smsAuthentication node.
	IdCheckConfigurationName *string `json:"idCheckConfigurationName,omitempty"`
	IdCheckInformationInput *IdCheckInformationInput `json:"idCheckInformationInput,omitempty"`
	// 
	IdentificationMethod *string `json:"identificationMethod,omitempty"`
	IdentityVerification *RecipientIdentityVerification `json:"identityVerification,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
	// A note sent to the recipient in the signing email. This note is unique to this recipient. In the user interface, it appears near the upper left corner of the document on the signing screen.  Maximum Length: 1000 characters. 
	Note *string `json:"note,omitempty"`
	PhoneAuthentication *RecipientPhoneAuthentication `json:"phoneAuthentication,omitempty"`
	// A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
	RecipientId *string `json:"recipientId,omitempty"`
	// The default signature provider is the DocuSign Electronic signature system. This parameter is used to specify one or more Standards Based Signature (digital signature) providers for the signer to use. [More information.](/docs/esign-rest-api/esign101/concepts/standards-based-signatures/)
	RecipientSignatureProviders []RecipientSignatureProvider `json:"recipientSignatureProviders,omitempty"`
	// The name of the role associated with the recipient.  **Note:** Every recipient must be assigned either a `recipientId` or a `roleName` but not both. You cannot use `roleName` and `recipientId` in the same list.
	RoleName *string `json:"roleName,omitempty"`
	// The in-person signer's full legal name.  Required when `inPersonSigningType` is `inPersonSigner`. For eNotary flow, use `name` instead.  Maximum Length: 100 characters. 
	SignerName *string `json:"signerName,omitempty"`
	// The ID of the [signing group](https://support.docusign.com/s/document-item?bundleId=gav1643676262430&topicId=zgn1578456447934.html). 
	SigningGroupId *string `json:"signingGroupId,omitempty"`
	SmsAuthentication *RecipientSMSAuthentication `json:"smsAuthentication,omitempty"`
	// Deprecated.
	SocialAuthentications []SocialAuthentication `json:"socialAuthentications,omitempty"`
	// A list of tabs associated with the recipient. In a bulk send request, each of these recipient tabs must match a recipient tab on the envelope or template that you want to send. To match up, the `tabLabel` for this tab and the `tabLabel` for the corresponding tab on the envelope or template must be the same.  For example, if the envelope has a placeholder text tab with the `tabLabel` `childName`, you must assign the same `tabLabel` `childName` to the tab here that you are populating with that information.   You can use the following types of tabs to match bulk send recipients to an envelope:  - Text tabs - Radio group tabs (where the name of the `radioGroup` on the envelope is used as the `tabLabel` in the bulk send list) - List tabs
	Tabs []BulkSendingCopyTab `json:"tabs,omitempty"`
}

// NewBulkSendingCopyRecipient instantiates a new BulkSendingCopyRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSendingCopyRecipient() *BulkSendingCopyRecipient {
	this := BulkSendingCopyRecipient{}
	return &this
}

// NewBulkSendingCopyRecipientWithDefaults instantiates a new BulkSendingCopyRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSendingCopyRecipientWithDefaults() *BulkSendingCopyRecipient {
	this := BulkSendingCopyRecipient{}
	return &this
}

// GetAccessCode returns the AccessCode field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetAccessCode() string {
	if o == nil || IsNil(o.AccessCode) {
		var ret string
		return ret
	}
	return *o.AccessCode
}

// GetAccessCodeOk returns a tuple with the AccessCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetAccessCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessCode) {
		return nil, false
	}
	return o.AccessCode, true
}

// HasAccessCode returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasAccessCode() bool {
	if o != nil && !IsNil(o.AccessCode) {
		return true
	}

	return false
}

// SetAccessCode gets a reference to the given string and assigns it to the AccessCode field.
func (o *BulkSendingCopyRecipient) SetAccessCode(v string) {
	o.AccessCode = &v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetClientUserId() string {
	if o == nil || IsNil(o.ClientUserId) {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetClientUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientUserId) {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasClientUserId() bool {
	if o != nil && !IsNil(o.ClientUserId) {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *BulkSendingCopyRecipient) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetCustomFields() []string {
	if o == nil || IsNil(o.CustomFields) {
		var ret []string
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetCustomFieldsOk() ([]string, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []string and assigns it to the CustomFields field.
func (o *BulkSendingCopyRecipient) SetCustomFields(v []string) {
	o.CustomFields = v
}

// GetDeliveryMethod returns the DeliveryMethod field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetDeliveryMethod() string {
	if o == nil || IsNil(o.DeliveryMethod) {
		var ret string
		return ret
	}
	return *o.DeliveryMethod
}

// GetDeliveryMethodOk returns a tuple with the DeliveryMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetDeliveryMethodOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryMethod) {
		return nil, false
	}
	return o.DeliveryMethod, true
}

// HasDeliveryMethod returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasDeliveryMethod() bool {
	if o != nil && !IsNil(o.DeliveryMethod) {
		return true
	}

	return false
}

// SetDeliveryMethod gets a reference to the given string and assigns it to the DeliveryMethod field.
func (o *BulkSendingCopyRecipient) SetDeliveryMethod(v string) {
	o.DeliveryMethod = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BulkSendingCopyRecipient) SetEmail(v string) {
	o.Email = &v
}

// GetEmailNotification returns the EmailNotification field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetEmailNotification() RecipientEmailNotification {
	if o == nil || IsNil(o.EmailNotification) {
		var ret RecipientEmailNotification
		return ret
	}
	return *o.EmailNotification
}

// GetEmailNotificationOk returns a tuple with the EmailNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetEmailNotificationOk() (*RecipientEmailNotification, bool) {
	if o == nil || IsNil(o.EmailNotification) {
		return nil, false
	}
	return o.EmailNotification, true
}

// HasEmailNotification returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasEmailNotification() bool {
	if o != nil && !IsNil(o.EmailNotification) {
		return true
	}

	return false
}

// SetEmailNotification gets a reference to the given RecipientEmailNotification and assigns it to the EmailNotification field.
func (o *BulkSendingCopyRecipient) SetEmailNotification(v RecipientEmailNotification) {
	o.EmailNotification = &v
}

// GetEmbeddedRecipientStartURL returns the EmbeddedRecipientStartURL field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetEmbeddedRecipientStartURL() string {
	if o == nil || IsNil(o.EmbeddedRecipientStartURL) {
		var ret string
		return ret
	}
	return *o.EmbeddedRecipientStartURL
}

// GetEmbeddedRecipientStartURLOk returns a tuple with the EmbeddedRecipientStartURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetEmbeddedRecipientStartURLOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddedRecipientStartURL) {
		return nil, false
	}
	return o.EmbeddedRecipientStartURL, true
}

// HasEmbeddedRecipientStartURL returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasEmbeddedRecipientStartURL() bool {
	if o != nil && !IsNil(o.EmbeddedRecipientStartURL) {
		return true
	}

	return false
}

// SetEmbeddedRecipientStartURL gets a reference to the given string and assigns it to the EmbeddedRecipientStartURL field.
func (o *BulkSendingCopyRecipient) SetEmbeddedRecipientStartURL(v string) {
	o.EmbeddedRecipientStartURL = &v
}

// GetFaxNumber returns the FaxNumber field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetFaxNumber() string {
	if o == nil || IsNil(o.FaxNumber) {
		var ret string
		return ret
	}
	return *o.FaxNumber
}

// GetFaxNumberOk returns a tuple with the FaxNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetFaxNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FaxNumber) {
		return nil, false
	}
	return o.FaxNumber, true
}

// HasFaxNumber returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasFaxNumber() bool {
	if o != nil && !IsNil(o.FaxNumber) {
		return true
	}

	return false
}

// SetFaxNumber gets a reference to the given string and assigns it to the FaxNumber field.
func (o *BulkSendingCopyRecipient) SetFaxNumber(v string) {
	o.FaxNumber = &v
}

// GetHostEmail returns the HostEmail field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetHostEmail() string {
	if o == nil || IsNil(o.HostEmail) {
		var ret string
		return ret
	}
	return *o.HostEmail
}

// GetHostEmailOk returns a tuple with the HostEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetHostEmailOk() (*string, bool) {
	if o == nil || IsNil(o.HostEmail) {
		return nil, false
	}
	return o.HostEmail, true
}

// HasHostEmail returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasHostEmail() bool {
	if o != nil && !IsNil(o.HostEmail) {
		return true
	}

	return false
}

// SetHostEmail gets a reference to the given string and assigns it to the HostEmail field.
func (o *BulkSendingCopyRecipient) SetHostEmail(v string) {
	o.HostEmail = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetHostName() string {
	if o == nil || IsNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetHostNameOk() (*string, bool) {
	if o == nil || IsNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasHostName() bool {
	if o != nil && !IsNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *BulkSendingCopyRecipient) SetHostName(v string) {
	o.HostName = &v
}

// GetIdCheckConfigurationName returns the IdCheckConfigurationName field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetIdCheckConfigurationName() string {
	if o == nil || IsNil(o.IdCheckConfigurationName) {
		var ret string
		return ret
	}
	return *o.IdCheckConfigurationName
}

// GetIdCheckConfigurationNameOk returns a tuple with the IdCheckConfigurationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetIdCheckConfigurationNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdCheckConfigurationName) {
		return nil, false
	}
	return o.IdCheckConfigurationName, true
}

// HasIdCheckConfigurationName returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasIdCheckConfigurationName() bool {
	if o != nil && !IsNil(o.IdCheckConfigurationName) {
		return true
	}

	return false
}

// SetIdCheckConfigurationName gets a reference to the given string and assigns it to the IdCheckConfigurationName field.
func (o *BulkSendingCopyRecipient) SetIdCheckConfigurationName(v string) {
	o.IdCheckConfigurationName = &v
}

// GetIdCheckInformationInput returns the IdCheckInformationInput field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetIdCheckInformationInput() IdCheckInformationInput {
	if o == nil || IsNil(o.IdCheckInformationInput) {
		var ret IdCheckInformationInput
		return ret
	}
	return *o.IdCheckInformationInput
}

// GetIdCheckInformationInputOk returns a tuple with the IdCheckInformationInput field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetIdCheckInformationInputOk() (*IdCheckInformationInput, bool) {
	if o == nil || IsNil(o.IdCheckInformationInput) {
		return nil, false
	}
	return o.IdCheckInformationInput, true
}

// HasIdCheckInformationInput returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasIdCheckInformationInput() bool {
	if o != nil && !IsNil(o.IdCheckInformationInput) {
		return true
	}

	return false
}

// SetIdCheckInformationInput gets a reference to the given IdCheckInformationInput and assigns it to the IdCheckInformationInput field.
func (o *BulkSendingCopyRecipient) SetIdCheckInformationInput(v IdCheckInformationInput) {
	o.IdCheckInformationInput = &v
}

// GetIdentificationMethod returns the IdentificationMethod field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetIdentificationMethod() string {
	if o == nil || IsNil(o.IdentificationMethod) {
		var ret string
		return ret
	}
	return *o.IdentificationMethod
}

// GetIdentificationMethodOk returns a tuple with the IdentificationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetIdentificationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.IdentificationMethod) {
		return nil, false
	}
	return o.IdentificationMethod, true
}

// HasIdentificationMethod returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasIdentificationMethod() bool {
	if o != nil && !IsNil(o.IdentificationMethod) {
		return true
	}

	return false
}

// SetIdentificationMethod gets a reference to the given string and assigns it to the IdentificationMethod field.
func (o *BulkSendingCopyRecipient) SetIdentificationMethod(v string) {
	o.IdentificationMethod = &v
}

// GetIdentityVerification returns the IdentityVerification field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetIdentityVerification() RecipientIdentityVerification {
	if o == nil || IsNil(o.IdentityVerification) {
		var ret RecipientIdentityVerification
		return ret
	}
	return *o.IdentityVerification
}

// GetIdentityVerificationOk returns a tuple with the IdentityVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetIdentityVerificationOk() (*RecipientIdentityVerification, bool) {
	if o == nil || IsNil(o.IdentityVerification) {
		return nil, false
	}
	return o.IdentityVerification, true
}

// HasIdentityVerification returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasIdentityVerification() bool {
	if o != nil && !IsNil(o.IdentityVerification) {
		return true
	}

	return false
}

// SetIdentityVerification gets a reference to the given RecipientIdentityVerification and assigns it to the IdentityVerification field.
func (o *BulkSendingCopyRecipient) SetIdentityVerification(v RecipientIdentityVerification) {
	o.IdentityVerification = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BulkSendingCopyRecipient) SetName(v string) {
	o.Name = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *BulkSendingCopyRecipient) SetNote(v string) {
	o.Note = &v
}

// GetPhoneAuthentication returns the PhoneAuthentication field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetPhoneAuthentication() RecipientPhoneAuthentication {
	if o == nil || IsNil(o.PhoneAuthentication) {
		var ret RecipientPhoneAuthentication
		return ret
	}
	return *o.PhoneAuthentication
}

// GetPhoneAuthenticationOk returns a tuple with the PhoneAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetPhoneAuthenticationOk() (*RecipientPhoneAuthentication, bool) {
	if o == nil || IsNil(o.PhoneAuthentication) {
		return nil, false
	}
	return o.PhoneAuthentication, true
}

// HasPhoneAuthentication returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasPhoneAuthentication() bool {
	if o != nil && !IsNil(o.PhoneAuthentication) {
		return true
	}

	return false
}

// SetPhoneAuthentication gets a reference to the given RecipientPhoneAuthentication and assigns it to the PhoneAuthentication field.
func (o *BulkSendingCopyRecipient) SetPhoneAuthentication(v RecipientPhoneAuthentication) {
	o.PhoneAuthentication = &v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *BulkSendingCopyRecipient) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetRecipientSignatureProviders returns the RecipientSignatureProviders field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetRecipientSignatureProviders() []RecipientSignatureProvider {
	if o == nil || IsNil(o.RecipientSignatureProviders) {
		var ret []RecipientSignatureProvider
		return ret
	}
	return o.RecipientSignatureProviders
}

// GetRecipientSignatureProvidersOk returns a tuple with the RecipientSignatureProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetRecipientSignatureProvidersOk() ([]RecipientSignatureProvider, bool) {
	if o == nil || IsNil(o.RecipientSignatureProviders) {
		return nil, false
	}
	return o.RecipientSignatureProviders, true
}

// HasRecipientSignatureProviders returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasRecipientSignatureProviders() bool {
	if o != nil && !IsNil(o.RecipientSignatureProviders) {
		return true
	}

	return false
}

// SetRecipientSignatureProviders gets a reference to the given []RecipientSignatureProvider and assigns it to the RecipientSignatureProviders field.
func (o *BulkSendingCopyRecipient) SetRecipientSignatureProviders(v []RecipientSignatureProvider) {
	o.RecipientSignatureProviders = v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetRoleName() string {
	if o == nil || IsNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetRoleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RoleName) {
		return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasRoleName() bool {
	if o != nil && !IsNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *BulkSendingCopyRecipient) SetRoleName(v string) {
	o.RoleName = &v
}

// GetSignerName returns the SignerName field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetSignerName() string {
	if o == nil || IsNil(o.SignerName) {
		var ret string
		return ret
	}
	return *o.SignerName
}

// GetSignerNameOk returns a tuple with the SignerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetSignerNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignerName) {
		return nil, false
	}
	return o.SignerName, true
}

// HasSignerName returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasSignerName() bool {
	if o != nil && !IsNil(o.SignerName) {
		return true
	}

	return false
}

// SetSignerName gets a reference to the given string and assigns it to the SignerName field.
func (o *BulkSendingCopyRecipient) SetSignerName(v string) {
	o.SignerName = &v
}

// GetSigningGroupId returns the SigningGroupId field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetSigningGroupId() string {
	if o == nil || IsNil(o.SigningGroupId) {
		var ret string
		return ret
	}
	return *o.SigningGroupId
}

// GetSigningGroupIdOk returns a tuple with the SigningGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetSigningGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.SigningGroupId) {
		return nil, false
	}
	return o.SigningGroupId, true
}

// HasSigningGroupId returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasSigningGroupId() bool {
	if o != nil && !IsNil(o.SigningGroupId) {
		return true
	}

	return false
}

// SetSigningGroupId gets a reference to the given string and assigns it to the SigningGroupId field.
func (o *BulkSendingCopyRecipient) SetSigningGroupId(v string) {
	o.SigningGroupId = &v
}

// GetSmsAuthentication returns the SmsAuthentication field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetSmsAuthentication() RecipientSMSAuthentication {
	if o == nil || IsNil(o.SmsAuthentication) {
		var ret RecipientSMSAuthentication
		return ret
	}
	return *o.SmsAuthentication
}

// GetSmsAuthenticationOk returns a tuple with the SmsAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetSmsAuthenticationOk() (*RecipientSMSAuthentication, bool) {
	if o == nil || IsNil(o.SmsAuthentication) {
		return nil, false
	}
	return o.SmsAuthentication, true
}

// HasSmsAuthentication returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasSmsAuthentication() bool {
	if o != nil && !IsNil(o.SmsAuthentication) {
		return true
	}

	return false
}

// SetSmsAuthentication gets a reference to the given RecipientSMSAuthentication and assigns it to the SmsAuthentication field.
func (o *BulkSendingCopyRecipient) SetSmsAuthentication(v RecipientSMSAuthentication) {
	o.SmsAuthentication = &v
}

// GetSocialAuthentications returns the SocialAuthentications field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetSocialAuthentications() []SocialAuthentication {
	if o == nil || IsNil(o.SocialAuthentications) {
		var ret []SocialAuthentication
		return ret
	}
	return o.SocialAuthentications
}

// GetSocialAuthenticationsOk returns a tuple with the SocialAuthentications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetSocialAuthenticationsOk() ([]SocialAuthentication, bool) {
	if o == nil || IsNil(o.SocialAuthentications) {
		return nil, false
	}
	return o.SocialAuthentications, true
}

// HasSocialAuthentications returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasSocialAuthentications() bool {
	if o != nil && !IsNil(o.SocialAuthentications) {
		return true
	}

	return false
}

// SetSocialAuthentications gets a reference to the given []SocialAuthentication and assigns it to the SocialAuthentications field.
func (o *BulkSendingCopyRecipient) SetSocialAuthentications(v []SocialAuthentication) {
	o.SocialAuthentications = v
}

// GetTabs returns the Tabs field value if set, zero value otherwise.
func (o *BulkSendingCopyRecipient) GetTabs() []BulkSendingCopyTab {
	if o == nil || IsNil(o.Tabs) {
		var ret []BulkSendingCopyTab
		return ret
	}
	return o.Tabs
}

// GetTabsOk returns a tuple with the Tabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendingCopyRecipient) GetTabsOk() ([]BulkSendingCopyTab, bool) {
	if o == nil || IsNil(o.Tabs) {
		return nil, false
	}
	return o.Tabs, true
}

// HasTabs returns a boolean if a field has been set.
func (o *BulkSendingCopyRecipient) HasTabs() bool {
	if o != nil && !IsNil(o.Tabs) {
		return true
	}

	return false
}

// SetTabs gets a reference to the given []BulkSendingCopyTab and assigns it to the Tabs field.
func (o *BulkSendingCopyRecipient) SetTabs(v []BulkSendingCopyTab) {
	o.Tabs = v
}

func (o BulkSendingCopyRecipient) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkSendingCopyRecipient) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessCode) {
		toSerialize["accessCode"] = o.AccessCode
	}
	if !IsNil(o.ClientUserId) {
		toSerialize["clientUserId"] = o.ClientUserId
	}
	if !IsNil(o.CustomFields) {
		toSerialize["customFields"] = o.CustomFields
	}
	if !IsNil(o.DeliveryMethod) {
		toSerialize["deliveryMethod"] = o.DeliveryMethod
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.EmailNotification) {
		toSerialize["emailNotification"] = o.EmailNotification
	}
	if !IsNil(o.EmbeddedRecipientStartURL) {
		toSerialize["embeddedRecipientStartURL"] = o.EmbeddedRecipientStartURL
	}
	if !IsNil(o.FaxNumber) {
		toSerialize["faxNumber"] = o.FaxNumber
	}
	if !IsNil(o.HostEmail) {
		toSerialize["hostEmail"] = o.HostEmail
	}
	if !IsNil(o.HostName) {
		toSerialize["hostName"] = o.HostName
	}
	if !IsNil(o.IdCheckConfigurationName) {
		toSerialize["idCheckConfigurationName"] = o.IdCheckConfigurationName
	}
	if !IsNil(o.IdCheckInformationInput) {
		toSerialize["idCheckInformationInput"] = o.IdCheckInformationInput
	}
	if !IsNil(o.IdentificationMethod) {
		toSerialize["identificationMethod"] = o.IdentificationMethod
	}
	if !IsNil(o.IdentityVerification) {
		toSerialize["identityVerification"] = o.IdentityVerification
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.PhoneAuthentication) {
		toSerialize["phoneAuthentication"] = o.PhoneAuthentication
	}
	if !IsNil(o.RecipientId) {
		toSerialize["recipientId"] = o.RecipientId
	}
	if !IsNil(o.RecipientSignatureProviders) {
		toSerialize["recipientSignatureProviders"] = o.RecipientSignatureProviders
	}
	if !IsNil(o.RoleName) {
		toSerialize["roleName"] = o.RoleName
	}
	if !IsNil(o.SignerName) {
		toSerialize["signerName"] = o.SignerName
	}
	if !IsNil(o.SigningGroupId) {
		toSerialize["signingGroupId"] = o.SigningGroupId
	}
	if !IsNil(o.SmsAuthentication) {
		toSerialize["smsAuthentication"] = o.SmsAuthentication
	}
	if !IsNil(o.SocialAuthentications) {
		toSerialize["socialAuthentications"] = o.SocialAuthentications
	}
	if !IsNil(o.Tabs) {
		toSerialize["tabs"] = o.Tabs
	}
	return toSerialize, nil
}

type NullableBulkSendingCopyRecipient struct {
	value *BulkSendingCopyRecipient
	isSet bool
}

func (v NullableBulkSendingCopyRecipient) Get() *BulkSendingCopyRecipient {
	return v.value
}

func (v *NullableBulkSendingCopyRecipient) Set(val *BulkSendingCopyRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSendingCopyRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSendingCopyRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSendingCopyRecipient(val *BulkSendingCopyRecipient) *NullableBulkSendingCopyRecipient {
	return &NullableBulkSendingCopyRecipient{value: val, isSet: true}
}

func (v NullableBulkSendingCopyRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSendingCopyRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


