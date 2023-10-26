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

// checks if the UserSignatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserSignatures{}

// UserSignatures Users' signatures
type UserSignatures struct {
	// The date and time the user adopted their signature.
	AdoptedDateTime *string `json:"adoptedDateTime,omitempty"`
	// The UTC date and time when the user created the signature.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	// Serialized information about any custom [eHanko stamps](https://support.docusign.com/s/articles/Sending-and-Signing-with-eHanko) that have been ordered from an eHanko provider, including the order status, purchase order id, time created, and time modified.
	CustomField *string `json:"customField,omitempty"`
	DateStampProperties *DateStampProperties `json:"dateStampProperties,omitempty"`
	// When **true,** users may not resize the stamp.
	DisallowUserResizeStamp *string `json:"disallowUserResizeStamp,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// An external ID for the signature or stamp.  **Note:** If a recipient uses a stamp instead of a signature, this is the stamp vendor's serial number for the stamp.
	ExternalID *string `json:"externalID,omitempty"`
	// A Base64-encoded representation of the signature image.
	ImageBase64 *string `json:"imageBase64,omitempty"`
	// The format of the signature image, such as:  - `GIF` - `PNG` - `JPG` - `PDF` - `BMP`
	ImageType *string `json:"imageType,omitempty"`
	// The ID of the user's initials image.
	Initials150ImageId *string `json:"initials150ImageId,omitempty"`
	// The URI for retrieving the image of the user's initials.
	InitialsImageUri *string `json:"initialsImageUri,omitempty"`
	// Boolean that specifies whether the signature is the default signature for the user.
	IsDefault *string `json:"isDefault,omitempty"`
	// The UTC date and time when the signature was last modified.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	// The National Association of Realtors (NAR) membership ID for a user who is a realtor.
	NrdsId *string `json:"nrdsId,omitempty"`
	// The realtor's last name.
	NrdsLastName *string `json:"nrdsLastName,omitempty"`
	// The realtor's NAR membership status. The value `active` verifies that the user is a current NAR member. Valid values are:  - `Active` - `Inactive` - `Terminate` - `Provisional` - `Deceased` - `Suspend` - `Unknown`
	NrdsStatus *string `json:"nrdsStatus,omitempty"`
	// The phonetic spelling of the `signatureName`.
	PhoneticName *string `json:"phoneticName,omitempty"`
	// The ID of the user's signature image.
	Signature150ImageId *string `json:"signature150ImageId,omitempty"`
	// The font type for the signature, if the signature is not drawn. The supported font types are:  \"7_DocuSign\", \"1_DocuSign\", \"6_DocuSign\", \"8_DocuSign\", \"3_DocuSign\", \"Mistral\", \"4_DocuSign\", \"2_DocuSign\", \"5_DocuSign\", \"Rage Italic\" 
	SignatureFont *string `json:"signatureFont,omitempty"`
	// The ID associated with the signature name. You can use this property in the URI in place of the signature name. This enables the use of special characters (such as \"&\", \"<\", and \">\") in a signature name.  **Note:** When you update a signature, its signature ID might change. In that case you need to use `signatureName` to get the new `signatureId`.
	SignatureId *string `json:"signatureId,omitempty"`
	// An endpoint URI that you can use to retrieve the user's signature image.
	SignatureImageUri *string `json:"signatureImageUri,omitempty"`
	//  The initials associated with the signature.
	SignatureInitials *string `json:"signatureInitials,omitempty"`
	// Specifies the user's signature name.
	SignatureName *string `json:"signatureName,omitempty"`
	// The rights that the user has to the signature. Valid values are:  - `none` - `read` - `admin`
	SignatureRights *string `json:"signatureRights,omitempty"`
	// Specifies the type of signature. Possible values include:  - `RubberStamp`: A DocuSign pre-formatted signature style. This is the default value. - `Imported`: A signature image that the user uploaded. - `Drawn`: A freehand drawing of the user's signature and initials.
	SignatureType *string `json:"signatureType,omitempty"`
	// The format of a stamp. Valid values are:  - `NameHanko`: The stamp represents only the signer's name. - `NameDateHanko`: The stamp represents the signer's name and the date. 
	StampFormat *string `json:"stampFormat,omitempty"`
	// The URI for retrieving the image of the user's stamp.
	StampImageUri *string `json:"stampImageUri,omitempty"`
	// The physical height of the stamp image (in millimeters) that the stamp vendor recommends for displaying the image in PDF documents.
	StampSizeMM *string `json:"stampSizeMM,omitempty"`
	// The type of stamp. Valid values are:  - `signature`: A signature image. This is the default value. - `stamp`: A stamp image. - null
	StampType *string `json:"stampType,omitempty"`
	// Indicates the envelope status. Valid values are:  * sent - The envelope is sent to the recipients.  * created - The envelope is saved as a draft and can be modified and sent later.
	Status *string `json:"status,omitempty"`
}

// NewUserSignatures instantiates a new UserSignatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserSignatures() *UserSignatures {
	this := UserSignatures{}
	return &this
}

// NewUserSignaturesWithDefaults instantiates a new UserSignatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserSignaturesWithDefaults() *UserSignatures {
	this := UserSignatures{}
	return &this
}

// GetAdoptedDateTime returns the AdoptedDateTime field value if set, zero value otherwise.
func (o *UserSignatures) GetAdoptedDateTime() string {
	if o == nil || IsNil(o.AdoptedDateTime) {
		var ret string
		return ret
	}
	return *o.AdoptedDateTime
}

// GetAdoptedDateTimeOk returns a tuple with the AdoptedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetAdoptedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.AdoptedDateTime) {
		return nil, false
	}
	return o.AdoptedDateTime, true
}

// HasAdoptedDateTime returns a boolean if a field has been set.
func (o *UserSignatures) HasAdoptedDateTime() bool {
	if o != nil && !IsNil(o.AdoptedDateTime) {
		return true
	}

	return false
}

// SetAdoptedDateTime gets a reference to the given string and assigns it to the AdoptedDateTime field.
func (o *UserSignatures) SetAdoptedDateTime(v string) {
	o.AdoptedDateTime = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *UserSignatures) GetCreatedDateTime() string {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret string
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetCreatedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *UserSignatures) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given string and assigns it to the CreatedDateTime field.
func (o *UserSignatures) SetCreatedDateTime(v string) {
	o.CreatedDateTime = &v
}

// GetCustomField returns the CustomField field value if set, zero value otherwise.
func (o *UserSignatures) GetCustomField() string {
	if o == nil || IsNil(o.CustomField) {
		var ret string
		return ret
	}
	return *o.CustomField
}

// GetCustomFieldOk returns a tuple with the CustomField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetCustomFieldOk() (*string, bool) {
	if o == nil || IsNil(o.CustomField) {
		return nil, false
	}
	return o.CustomField, true
}

// HasCustomField returns a boolean if a field has been set.
func (o *UserSignatures) HasCustomField() bool {
	if o != nil && !IsNil(o.CustomField) {
		return true
	}

	return false
}

// SetCustomField gets a reference to the given string and assigns it to the CustomField field.
func (o *UserSignatures) SetCustomField(v string) {
	o.CustomField = &v
}

// GetDateStampProperties returns the DateStampProperties field value if set, zero value otherwise.
func (o *UserSignatures) GetDateStampProperties() DateStampProperties {
	if o == nil || IsNil(o.DateStampProperties) {
		var ret DateStampProperties
		return ret
	}
	return *o.DateStampProperties
}

// GetDateStampPropertiesOk returns a tuple with the DateStampProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetDateStampPropertiesOk() (*DateStampProperties, bool) {
	if o == nil || IsNil(o.DateStampProperties) {
		return nil, false
	}
	return o.DateStampProperties, true
}

// HasDateStampProperties returns a boolean if a field has been set.
func (o *UserSignatures) HasDateStampProperties() bool {
	if o != nil && !IsNil(o.DateStampProperties) {
		return true
	}

	return false
}

// SetDateStampProperties gets a reference to the given DateStampProperties and assigns it to the DateStampProperties field.
func (o *UserSignatures) SetDateStampProperties(v DateStampProperties) {
	o.DateStampProperties = &v
}

// GetDisallowUserResizeStamp returns the DisallowUserResizeStamp field value if set, zero value otherwise.
func (o *UserSignatures) GetDisallowUserResizeStamp() string {
	if o == nil || IsNil(o.DisallowUserResizeStamp) {
		var ret string
		return ret
	}
	return *o.DisallowUserResizeStamp
}

// GetDisallowUserResizeStampOk returns a tuple with the DisallowUserResizeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetDisallowUserResizeStampOk() (*string, bool) {
	if o == nil || IsNil(o.DisallowUserResizeStamp) {
		return nil, false
	}
	return o.DisallowUserResizeStamp, true
}

// HasDisallowUserResizeStamp returns a boolean if a field has been set.
func (o *UserSignatures) HasDisallowUserResizeStamp() bool {
	if o != nil && !IsNil(o.DisallowUserResizeStamp) {
		return true
	}

	return false
}

// SetDisallowUserResizeStamp gets a reference to the given string and assigns it to the DisallowUserResizeStamp field.
func (o *UserSignatures) SetDisallowUserResizeStamp(v string) {
	o.DisallowUserResizeStamp = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *UserSignatures) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *UserSignatures) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *UserSignatures) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *UserSignatures) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *UserSignatures) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *UserSignatures) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetImageBase64 returns the ImageBase64 field value if set, zero value otherwise.
func (o *UserSignatures) GetImageBase64() string {
	if o == nil || IsNil(o.ImageBase64) {
		var ret string
		return ret
	}
	return *o.ImageBase64
}

// GetImageBase64Ok returns a tuple with the ImageBase64 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetImageBase64Ok() (*string, bool) {
	if o == nil || IsNil(o.ImageBase64) {
		return nil, false
	}
	return o.ImageBase64, true
}

// HasImageBase64 returns a boolean if a field has been set.
func (o *UserSignatures) HasImageBase64() bool {
	if o != nil && !IsNil(o.ImageBase64) {
		return true
	}

	return false
}

// SetImageBase64 gets a reference to the given string and assigns it to the ImageBase64 field.
func (o *UserSignatures) SetImageBase64(v string) {
	o.ImageBase64 = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *UserSignatures) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *UserSignatures) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *UserSignatures) SetImageType(v string) {
	o.ImageType = &v
}

// GetInitials150ImageId returns the Initials150ImageId field value if set, zero value otherwise.
func (o *UserSignatures) GetInitials150ImageId() string {
	if o == nil || IsNil(o.Initials150ImageId) {
		var ret string
		return ret
	}
	return *o.Initials150ImageId
}

// GetInitials150ImageIdOk returns a tuple with the Initials150ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetInitials150ImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.Initials150ImageId) {
		return nil, false
	}
	return o.Initials150ImageId, true
}

// HasInitials150ImageId returns a boolean if a field has been set.
func (o *UserSignatures) HasInitials150ImageId() bool {
	if o != nil && !IsNil(o.Initials150ImageId) {
		return true
	}

	return false
}

// SetInitials150ImageId gets a reference to the given string and assigns it to the Initials150ImageId field.
func (o *UserSignatures) SetInitials150ImageId(v string) {
	o.Initials150ImageId = &v
}

// GetInitialsImageUri returns the InitialsImageUri field value if set, zero value otherwise.
func (o *UserSignatures) GetInitialsImageUri() string {
	if o == nil || IsNil(o.InitialsImageUri) {
		var ret string
		return ret
	}
	return *o.InitialsImageUri
}

// GetInitialsImageUriOk returns a tuple with the InitialsImageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetInitialsImageUriOk() (*string, bool) {
	if o == nil || IsNil(o.InitialsImageUri) {
		return nil, false
	}
	return o.InitialsImageUri, true
}

// HasInitialsImageUri returns a boolean if a field has been set.
func (o *UserSignatures) HasInitialsImageUri() bool {
	if o != nil && !IsNil(o.InitialsImageUri) {
		return true
	}

	return false
}

// SetInitialsImageUri gets a reference to the given string and assigns it to the InitialsImageUri field.
func (o *UserSignatures) SetInitialsImageUri(v string) {
	o.InitialsImageUri = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *UserSignatures) GetIsDefault() string {
	if o == nil || IsNil(o.IsDefault) {
		var ret string
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetIsDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *UserSignatures) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given string and assigns it to the IsDefault field.
func (o *UserSignatures) SetIsDefault(v string) {
	o.IsDefault = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *UserSignatures) GetLastModifiedDateTime() string {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		var ret string
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetLastModifiedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.LastModifiedDateTime) {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *UserSignatures) HasLastModifiedDateTime() bool {
	if o != nil && !IsNil(o.LastModifiedDateTime) {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given string and assigns it to the LastModifiedDateTime field.
func (o *UserSignatures) SetLastModifiedDateTime(v string) {
	o.LastModifiedDateTime = &v
}

// GetNrdsId returns the NrdsId field value if set, zero value otherwise.
func (o *UserSignatures) GetNrdsId() string {
	if o == nil || IsNil(o.NrdsId) {
		var ret string
		return ret
	}
	return *o.NrdsId
}

// GetNrdsIdOk returns a tuple with the NrdsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetNrdsIdOk() (*string, bool) {
	if o == nil || IsNil(o.NrdsId) {
		return nil, false
	}
	return o.NrdsId, true
}

// HasNrdsId returns a boolean if a field has been set.
func (o *UserSignatures) HasNrdsId() bool {
	if o != nil && !IsNil(o.NrdsId) {
		return true
	}

	return false
}

// SetNrdsId gets a reference to the given string and assigns it to the NrdsId field.
func (o *UserSignatures) SetNrdsId(v string) {
	o.NrdsId = &v
}

// GetNrdsLastName returns the NrdsLastName field value if set, zero value otherwise.
func (o *UserSignatures) GetNrdsLastName() string {
	if o == nil || IsNil(o.NrdsLastName) {
		var ret string
		return ret
	}
	return *o.NrdsLastName
}

// GetNrdsLastNameOk returns a tuple with the NrdsLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetNrdsLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.NrdsLastName) {
		return nil, false
	}
	return o.NrdsLastName, true
}

// HasNrdsLastName returns a boolean if a field has been set.
func (o *UserSignatures) HasNrdsLastName() bool {
	if o != nil && !IsNil(o.NrdsLastName) {
		return true
	}

	return false
}

// SetNrdsLastName gets a reference to the given string and assigns it to the NrdsLastName field.
func (o *UserSignatures) SetNrdsLastName(v string) {
	o.NrdsLastName = &v
}

// GetNrdsStatus returns the NrdsStatus field value if set, zero value otherwise.
func (o *UserSignatures) GetNrdsStatus() string {
	if o == nil || IsNil(o.NrdsStatus) {
		var ret string
		return ret
	}
	return *o.NrdsStatus
}

// GetNrdsStatusOk returns a tuple with the NrdsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetNrdsStatusOk() (*string, bool) {
	if o == nil || IsNil(o.NrdsStatus) {
		return nil, false
	}
	return o.NrdsStatus, true
}

// HasNrdsStatus returns a boolean if a field has been set.
func (o *UserSignatures) HasNrdsStatus() bool {
	if o != nil && !IsNil(o.NrdsStatus) {
		return true
	}

	return false
}

// SetNrdsStatus gets a reference to the given string and assigns it to the NrdsStatus field.
func (o *UserSignatures) SetNrdsStatus(v string) {
	o.NrdsStatus = &v
}

// GetPhoneticName returns the PhoneticName field value if set, zero value otherwise.
func (o *UserSignatures) GetPhoneticName() string {
	if o == nil || IsNil(o.PhoneticName) {
		var ret string
		return ret
	}
	return *o.PhoneticName
}

// GetPhoneticNameOk returns a tuple with the PhoneticName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetPhoneticNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneticName) {
		return nil, false
	}
	return o.PhoneticName, true
}

// HasPhoneticName returns a boolean if a field has been set.
func (o *UserSignatures) HasPhoneticName() bool {
	if o != nil && !IsNil(o.PhoneticName) {
		return true
	}

	return false
}

// SetPhoneticName gets a reference to the given string and assigns it to the PhoneticName field.
func (o *UserSignatures) SetPhoneticName(v string) {
	o.PhoneticName = &v
}

// GetSignature150ImageId returns the Signature150ImageId field value if set, zero value otherwise.
func (o *UserSignatures) GetSignature150ImageId() string {
	if o == nil || IsNil(o.Signature150ImageId) {
		var ret string
		return ret
	}
	return *o.Signature150ImageId
}

// GetSignature150ImageIdOk returns a tuple with the Signature150ImageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignature150ImageIdOk() (*string, bool) {
	if o == nil || IsNil(o.Signature150ImageId) {
		return nil, false
	}
	return o.Signature150ImageId, true
}

// HasSignature150ImageId returns a boolean if a field has been set.
func (o *UserSignatures) HasSignature150ImageId() bool {
	if o != nil && !IsNil(o.Signature150ImageId) {
		return true
	}

	return false
}

// SetSignature150ImageId gets a reference to the given string and assigns it to the Signature150ImageId field.
func (o *UserSignatures) SetSignature150ImageId(v string) {
	o.Signature150ImageId = &v
}

// GetSignatureFont returns the SignatureFont field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureFont() string {
	if o == nil || IsNil(o.SignatureFont) {
		var ret string
		return ret
	}
	return *o.SignatureFont
}

// GetSignatureFontOk returns a tuple with the SignatureFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureFontOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureFont) {
		return nil, false
	}
	return o.SignatureFont, true
}

// HasSignatureFont returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureFont() bool {
	if o != nil && !IsNil(o.SignatureFont) {
		return true
	}

	return false
}

// SetSignatureFont gets a reference to the given string and assigns it to the SignatureFont field.
func (o *UserSignatures) SetSignatureFont(v string) {
	o.SignatureFont = &v
}

// GetSignatureId returns the SignatureId field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureId() string {
	if o == nil || IsNil(o.SignatureId) {
		var ret string
		return ret
	}
	return *o.SignatureId
}

// GetSignatureIdOk returns a tuple with the SignatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureId) {
		return nil, false
	}
	return o.SignatureId, true
}

// HasSignatureId returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureId() bool {
	if o != nil && !IsNil(o.SignatureId) {
		return true
	}

	return false
}

// SetSignatureId gets a reference to the given string and assigns it to the SignatureId field.
func (o *UserSignatures) SetSignatureId(v string) {
	o.SignatureId = &v
}

// GetSignatureImageUri returns the SignatureImageUri field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureImageUri() string {
	if o == nil || IsNil(o.SignatureImageUri) {
		var ret string
		return ret
	}
	return *o.SignatureImageUri
}

// GetSignatureImageUriOk returns a tuple with the SignatureImageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureImageUriOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureImageUri) {
		return nil, false
	}
	return o.SignatureImageUri, true
}

// HasSignatureImageUri returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureImageUri() bool {
	if o != nil && !IsNil(o.SignatureImageUri) {
		return true
	}

	return false
}

// SetSignatureImageUri gets a reference to the given string and assigns it to the SignatureImageUri field.
func (o *UserSignatures) SetSignatureImageUri(v string) {
	o.SignatureImageUri = &v
}

// GetSignatureInitials returns the SignatureInitials field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureInitials() string {
	if o == nil || IsNil(o.SignatureInitials) {
		var ret string
		return ret
	}
	return *o.SignatureInitials
}

// GetSignatureInitialsOk returns a tuple with the SignatureInitials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureInitialsOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureInitials) {
		return nil, false
	}
	return o.SignatureInitials, true
}

// HasSignatureInitials returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureInitials() bool {
	if o != nil && !IsNil(o.SignatureInitials) {
		return true
	}

	return false
}

// SetSignatureInitials gets a reference to the given string and assigns it to the SignatureInitials field.
func (o *UserSignatures) SetSignatureInitials(v string) {
	o.SignatureInitials = &v
}

// GetSignatureName returns the SignatureName field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureName() string {
	if o == nil || IsNil(o.SignatureName) {
		var ret string
		return ret
	}
	return *o.SignatureName
}

// GetSignatureNameOk returns a tuple with the SignatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureName) {
		return nil, false
	}
	return o.SignatureName, true
}

// HasSignatureName returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureName() bool {
	if o != nil && !IsNil(o.SignatureName) {
		return true
	}

	return false
}

// SetSignatureName gets a reference to the given string and assigns it to the SignatureName field.
func (o *UserSignatures) SetSignatureName(v string) {
	o.SignatureName = &v
}

// GetSignatureRights returns the SignatureRights field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureRights() string {
	if o == nil || IsNil(o.SignatureRights) {
		var ret string
		return ret
	}
	return *o.SignatureRights
}

// GetSignatureRightsOk returns a tuple with the SignatureRights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureRightsOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureRights) {
		return nil, false
	}
	return o.SignatureRights, true
}

// HasSignatureRights returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureRights() bool {
	if o != nil && !IsNil(o.SignatureRights) {
		return true
	}

	return false
}

// SetSignatureRights gets a reference to the given string and assigns it to the SignatureRights field.
func (o *UserSignatures) SetSignatureRights(v string) {
	o.SignatureRights = &v
}

// GetSignatureType returns the SignatureType field value if set, zero value otherwise.
func (o *UserSignatures) GetSignatureType() string {
	if o == nil || IsNil(o.SignatureType) {
		var ret string
		return ret
	}
	return *o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetSignatureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureType) {
		return nil, false
	}
	return o.SignatureType, true
}

// HasSignatureType returns a boolean if a field has been set.
func (o *UserSignatures) HasSignatureType() bool {
	if o != nil && !IsNil(o.SignatureType) {
		return true
	}

	return false
}

// SetSignatureType gets a reference to the given string and assigns it to the SignatureType field.
func (o *UserSignatures) SetSignatureType(v string) {
	o.SignatureType = &v
}

// GetStampFormat returns the StampFormat field value if set, zero value otherwise.
func (o *UserSignatures) GetStampFormat() string {
	if o == nil || IsNil(o.StampFormat) {
		var ret string
		return ret
	}
	return *o.StampFormat
}

// GetStampFormatOk returns a tuple with the StampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetStampFormatOk() (*string, bool) {
	if o == nil || IsNil(o.StampFormat) {
		return nil, false
	}
	return o.StampFormat, true
}

// HasStampFormat returns a boolean if a field has been set.
func (o *UserSignatures) HasStampFormat() bool {
	if o != nil && !IsNil(o.StampFormat) {
		return true
	}

	return false
}

// SetStampFormat gets a reference to the given string and assigns it to the StampFormat field.
func (o *UserSignatures) SetStampFormat(v string) {
	o.StampFormat = &v
}

// GetStampImageUri returns the StampImageUri field value if set, zero value otherwise.
func (o *UserSignatures) GetStampImageUri() string {
	if o == nil || IsNil(o.StampImageUri) {
		var ret string
		return ret
	}
	return *o.StampImageUri
}

// GetStampImageUriOk returns a tuple with the StampImageUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetStampImageUriOk() (*string, bool) {
	if o == nil || IsNil(o.StampImageUri) {
		return nil, false
	}
	return o.StampImageUri, true
}

// HasStampImageUri returns a boolean if a field has been set.
func (o *UserSignatures) HasStampImageUri() bool {
	if o != nil && !IsNil(o.StampImageUri) {
		return true
	}

	return false
}

// SetStampImageUri gets a reference to the given string and assigns it to the StampImageUri field.
func (o *UserSignatures) SetStampImageUri(v string) {
	o.StampImageUri = &v
}

// GetStampSizeMM returns the StampSizeMM field value if set, zero value otherwise.
func (o *UserSignatures) GetStampSizeMM() string {
	if o == nil || IsNil(o.StampSizeMM) {
		var ret string
		return ret
	}
	return *o.StampSizeMM
}

// GetStampSizeMMOk returns a tuple with the StampSizeMM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetStampSizeMMOk() (*string, bool) {
	if o == nil || IsNil(o.StampSizeMM) {
		return nil, false
	}
	return o.StampSizeMM, true
}

// HasStampSizeMM returns a boolean if a field has been set.
func (o *UserSignatures) HasStampSizeMM() bool {
	if o != nil && !IsNil(o.StampSizeMM) {
		return true
	}

	return false
}

// SetStampSizeMM gets a reference to the given string and assigns it to the StampSizeMM field.
func (o *UserSignatures) SetStampSizeMM(v string) {
	o.StampSizeMM = &v
}

// GetStampType returns the StampType field value if set, zero value otherwise.
func (o *UserSignatures) GetStampType() string {
	if o == nil || IsNil(o.StampType) {
		var ret string
		return ret
	}
	return *o.StampType
}

// GetStampTypeOk returns a tuple with the StampType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetStampTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StampType) {
		return nil, false
	}
	return o.StampType, true
}

// HasStampType returns a boolean if a field has been set.
func (o *UserSignatures) HasStampType() bool {
	if o != nil && !IsNil(o.StampType) {
		return true
	}

	return false
}

// SetStampType gets a reference to the given string and assigns it to the StampType field.
func (o *UserSignatures) SetStampType(v string) {
	o.StampType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UserSignatures) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserSignatures) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UserSignatures) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UserSignatures) SetStatus(v string) {
	o.Status = &v
}

func (o UserSignatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserSignatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdoptedDateTime) {
		toSerialize["adoptedDateTime"] = o.AdoptedDateTime
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.CustomField) {
		toSerialize["customField"] = o.CustomField
	}
	if !IsNil(o.DateStampProperties) {
		toSerialize["dateStampProperties"] = o.DateStampProperties
	}
	if !IsNil(o.DisallowUserResizeStamp) {
		toSerialize["disallowUserResizeStamp"] = o.DisallowUserResizeStamp
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.ImageBase64) {
		toSerialize["imageBase64"] = o.ImageBase64
	}
	if !IsNil(o.ImageType) {
		toSerialize["imageType"] = o.ImageType
	}
	if !IsNil(o.Initials150ImageId) {
		toSerialize["initials150ImageId"] = o.Initials150ImageId
	}
	if !IsNil(o.InitialsImageUri) {
		toSerialize["initialsImageUri"] = o.InitialsImageUri
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.LastModifiedDateTime) {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if !IsNil(o.NrdsId) {
		toSerialize["nrdsId"] = o.NrdsId
	}
	if !IsNil(o.NrdsLastName) {
		toSerialize["nrdsLastName"] = o.NrdsLastName
	}
	if !IsNil(o.NrdsStatus) {
		toSerialize["nrdsStatus"] = o.NrdsStatus
	}
	if !IsNil(o.PhoneticName) {
		toSerialize["phoneticName"] = o.PhoneticName
	}
	if !IsNil(o.Signature150ImageId) {
		toSerialize["signature150ImageId"] = o.Signature150ImageId
	}
	if !IsNil(o.SignatureFont) {
		toSerialize["signatureFont"] = o.SignatureFont
	}
	if !IsNil(o.SignatureId) {
		toSerialize["signatureId"] = o.SignatureId
	}
	if !IsNil(o.SignatureImageUri) {
		toSerialize["signatureImageUri"] = o.SignatureImageUri
	}
	if !IsNil(o.SignatureInitials) {
		toSerialize["signatureInitials"] = o.SignatureInitials
	}
	if !IsNil(o.SignatureName) {
		toSerialize["signatureName"] = o.SignatureName
	}
	if !IsNil(o.SignatureRights) {
		toSerialize["signatureRights"] = o.SignatureRights
	}
	if !IsNil(o.SignatureType) {
		toSerialize["signatureType"] = o.SignatureType
	}
	if !IsNil(o.StampFormat) {
		toSerialize["stampFormat"] = o.StampFormat
	}
	if !IsNil(o.StampImageUri) {
		toSerialize["stampImageUri"] = o.StampImageUri
	}
	if !IsNil(o.StampSizeMM) {
		toSerialize["stampSizeMM"] = o.StampSizeMM
	}
	if !IsNil(o.StampType) {
		toSerialize["stampType"] = o.StampType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUserSignatures struct {
	value *UserSignatures
	isSet bool
}

func (v NullableUserSignatures) Get() *UserSignatures {
	return v.value
}

func (v *NullableUserSignatures) Set(val *UserSignatures) {
	v.value = val
	v.isSet = true
}

func (v NullableUserSignatures) IsSet() bool {
	return v.isSet
}

func (v *NullableUserSignatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserSignatures(val *UserSignatures) *NullableUserSignatures {
	return &NullableUserSignatures{value: val, isSet: true}
}

func (v NullableUserSignatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserSignatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

