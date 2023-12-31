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

// checks if the AccountSignatureDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountSignatureDefinition{}

// AccountSignatureDefinition 
type AccountSignatureDefinition struct {
	DateStampProperties *DateStampProperties `json:"dateStampProperties,omitempty"`
	// When **true,** users may not resize the stamp.
	DisallowUserResizeStamp *string `json:"disallowUserResizeStamp,omitempty"`
	// Optionally specify an external identifier for the user's signature.
	ExternalID *string `json:"externalID,omitempty"`
	// Specificies the type of image. Valid values:  - `stamp_image` - `signature_image` - `initials_image`
	ImageType *string `json:"imageType,omitempty"`
	// Boolean that specifies whether the signature is the default signature for the user.
	IsDefault *string `json:"isDefault,omitempty"`
	// The National Association of Realtors (NAR) membership ID for a user who is a realtor.
	NrdsId *string `json:"nrdsId,omitempty"`
	// The realtor's last name.
	NrdsLastName *string `json:"nrdsLastName,omitempty"`
	// The phonetic spelling of the `signatureName`.
	PhoneticName *string `json:"phoneticName,omitempty"`
	// The font type to use for the signature if the signature is not drawn. The following font styles  are supported. The quotes are to indicate that these values are strings, not `enums`.  - `\"1_DocuSign\"` - `\"2_DocuSign\"` - `\"3_DocuSign\"` - `\"4_DocuSign\"` - `\"5_DocuSign\"` - `\"6_DocuSign\"` - `\"7_DocuSign\"` - `\"8_DocuSign\"` - `\"Mistral\"` - `\"Rage Italic\"` 
	SignatureFont *string `json:"signatureFont,omitempty"`
	// 
	SignatureGroups []SignatureGroupDef `json:"signatureGroups,omitempty"`
	// Specifies the signature ID associated with the signature name. You can use the signature ID in the URI in place of the signature name, and the value stored in the `signatureName` property in the body is used. This allows the use of special characters (such as \"&\", \"<\", \">\") in a the signature name. Note that with each update to signatures, the returned signature ID might change, so the caller will need to trigger off the signature name to get the new signature ID.
	SignatureId *string `json:"signatureId,omitempty"`
	// Specifies the user's signature in initials format.
	SignatureInitials *string `json:"signatureInitials,omitempty"`
	// Specifies the user's signature name.
	SignatureName *string `json:"signatureName,omitempty"`
	// Specifies the type of signature.
	SignatureType *string `json:"signatureType,omitempty"`
	// 
	SignatureUsers []SignatureUserDef `json:"signatureUsers,omitempty"`
	// The format of a stamp. Valid values are:  - `NameHanko`: The stamp represents only the signer's name. - `NameDateHanko`: The stamp represents the signer's name and the date. 
	StampFormat *string `json:"stampFormat,omitempty"`
	// The physical height of the stamp image (in millimeters) that the stamp vendor recommends for displaying the image in PDF documents.
	StampSizeMM *string `json:"stampSizeMM,omitempty"`
}

// NewAccountSignatureDefinition instantiates a new AccountSignatureDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountSignatureDefinition() *AccountSignatureDefinition {
	this := AccountSignatureDefinition{}
	return &this
}

// NewAccountSignatureDefinitionWithDefaults instantiates a new AccountSignatureDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountSignatureDefinitionWithDefaults() *AccountSignatureDefinition {
	this := AccountSignatureDefinition{}
	return &this
}

// GetDateStampProperties returns the DateStampProperties field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetDateStampProperties() DateStampProperties {
	if o == nil || IsNil(o.DateStampProperties) {
		var ret DateStampProperties
		return ret
	}
	return *o.DateStampProperties
}

// GetDateStampPropertiesOk returns a tuple with the DateStampProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetDateStampPropertiesOk() (*DateStampProperties, bool) {
	if o == nil || IsNil(o.DateStampProperties) {
		return nil, false
	}
	return o.DateStampProperties, true
}

// HasDateStampProperties returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasDateStampProperties() bool {
	if o != nil && !IsNil(o.DateStampProperties) {
		return true
	}

	return false
}

// SetDateStampProperties gets a reference to the given DateStampProperties and assigns it to the DateStampProperties field.
func (o *AccountSignatureDefinition) SetDateStampProperties(v DateStampProperties) {
	o.DateStampProperties = &v
}

// GetDisallowUserResizeStamp returns the DisallowUserResizeStamp field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetDisallowUserResizeStamp() string {
	if o == nil || IsNil(o.DisallowUserResizeStamp) {
		var ret string
		return ret
	}
	return *o.DisallowUserResizeStamp
}

// GetDisallowUserResizeStampOk returns a tuple with the DisallowUserResizeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetDisallowUserResizeStampOk() (*string, bool) {
	if o == nil || IsNil(o.DisallowUserResizeStamp) {
		return nil, false
	}
	return o.DisallowUserResizeStamp, true
}

// HasDisallowUserResizeStamp returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasDisallowUserResizeStamp() bool {
	if o != nil && !IsNil(o.DisallowUserResizeStamp) {
		return true
	}

	return false
}

// SetDisallowUserResizeStamp gets a reference to the given string and assigns it to the DisallowUserResizeStamp field.
func (o *AccountSignatureDefinition) SetDisallowUserResizeStamp(v string) {
	o.DisallowUserResizeStamp = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetExternalID() string {
	if o == nil || IsNil(o.ExternalID) {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetExternalIDOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalID) {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasExternalID() bool {
	if o != nil && !IsNil(o.ExternalID) {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *AccountSignatureDefinition) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetImageType returns the ImageType field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetImageType() string {
	if o == nil || IsNil(o.ImageType) {
		var ret string
		return ret
	}
	return *o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetImageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ImageType) {
		return nil, false
	}
	return o.ImageType, true
}

// HasImageType returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasImageType() bool {
	if o != nil && !IsNil(o.ImageType) {
		return true
	}

	return false
}

// SetImageType gets a reference to the given string and assigns it to the ImageType field.
func (o *AccountSignatureDefinition) SetImageType(v string) {
	o.ImageType = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetIsDefault() string {
	if o == nil || IsNil(o.IsDefault) {
		var ret string
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetIsDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasIsDefault() bool {
	if o != nil && !IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given string and assigns it to the IsDefault field.
func (o *AccountSignatureDefinition) SetIsDefault(v string) {
	o.IsDefault = &v
}

// GetNrdsId returns the NrdsId field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetNrdsId() string {
	if o == nil || IsNil(o.NrdsId) {
		var ret string
		return ret
	}
	return *o.NrdsId
}

// GetNrdsIdOk returns a tuple with the NrdsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetNrdsIdOk() (*string, bool) {
	if o == nil || IsNil(o.NrdsId) {
		return nil, false
	}
	return o.NrdsId, true
}

// HasNrdsId returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasNrdsId() bool {
	if o != nil && !IsNil(o.NrdsId) {
		return true
	}

	return false
}

// SetNrdsId gets a reference to the given string and assigns it to the NrdsId field.
func (o *AccountSignatureDefinition) SetNrdsId(v string) {
	o.NrdsId = &v
}

// GetNrdsLastName returns the NrdsLastName field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetNrdsLastName() string {
	if o == nil || IsNil(o.NrdsLastName) {
		var ret string
		return ret
	}
	return *o.NrdsLastName
}

// GetNrdsLastNameOk returns a tuple with the NrdsLastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetNrdsLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.NrdsLastName) {
		return nil, false
	}
	return o.NrdsLastName, true
}

// HasNrdsLastName returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasNrdsLastName() bool {
	if o != nil && !IsNil(o.NrdsLastName) {
		return true
	}

	return false
}

// SetNrdsLastName gets a reference to the given string and assigns it to the NrdsLastName field.
func (o *AccountSignatureDefinition) SetNrdsLastName(v string) {
	o.NrdsLastName = &v
}

// GetPhoneticName returns the PhoneticName field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetPhoneticName() string {
	if o == nil || IsNil(o.PhoneticName) {
		var ret string
		return ret
	}
	return *o.PhoneticName
}

// GetPhoneticNameOk returns a tuple with the PhoneticName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetPhoneticNameOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneticName) {
		return nil, false
	}
	return o.PhoneticName, true
}

// HasPhoneticName returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasPhoneticName() bool {
	if o != nil && !IsNil(o.PhoneticName) {
		return true
	}

	return false
}

// SetPhoneticName gets a reference to the given string and assigns it to the PhoneticName field.
func (o *AccountSignatureDefinition) SetPhoneticName(v string) {
	o.PhoneticName = &v
}

// GetSignatureFont returns the SignatureFont field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureFont() string {
	if o == nil || IsNil(o.SignatureFont) {
		var ret string
		return ret
	}
	return *o.SignatureFont
}

// GetSignatureFontOk returns a tuple with the SignatureFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureFontOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureFont) {
		return nil, false
	}
	return o.SignatureFont, true
}

// HasSignatureFont returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureFont() bool {
	if o != nil && !IsNil(o.SignatureFont) {
		return true
	}

	return false
}

// SetSignatureFont gets a reference to the given string and assigns it to the SignatureFont field.
func (o *AccountSignatureDefinition) SetSignatureFont(v string) {
	o.SignatureFont = &v
}

// GetSignatureGroups returns the SignatureGroups field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureGroups() []SignatureGroupDef {
	if o == nil || IsNil(o.SignatureGroups) {
		var ret []SignatureGroupDef
		return ret
	}
	return o.SignatureGroups
}

// GetSignatureGroupsOk returns a tuple with the SignatureGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureGroupsOk() ([]SignatureGroupDef, bool) {
	if o == nil || IsNil(o.SignatureGroups) {
		return nil, false
	}
	return o.SignatureGroups, true
}

// HasSignatureGroups returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureGroups() bool {
	if o != nil && !IsNil(o.SignatureGroups) {
		return true
	}

	return false
}

// SetSignatureGroups gets a reference to the given []SignatureGroupDef and assigns it to the SignatureGroups field.
func (o *AccountSignatureDefinition) SetSignatureGroups(v []SignatureGroupDef) {
	o.SignatureGroups = v
}

// GetSignatureId returns the SignatureId field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureId() string {
	if o == nil || IsNil(o.SignatureId) {
		var ret string
		return ret
	}
	return *o.SignatureId
}

// GetSignatureIdOk returns a tuple with the SignatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureIdOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureId) {
		return nil, false
	}
	return o.SignatureId, true
}

// HasSignatureId returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureId() bool {
	if o != nil && !IsNil(o.SignatureId) {
		return true
	}

	return false
}

// SetSignatureId gets a reference to the given string and assigns it to the SignatureId field.
func (o *AccountSignatureDefinition) SetSignatureId(v string) {
	o.SignatureId = &v
}

// GetSignatureInitials returns the SignatureInitials field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureInitials() string {
	if o == nil || IsNil(o.SignatureInitials) {
		var ret string
		return ret
	}
	return *o.SignatureInitials
}

// GetSignatureInitialsOk returns a tuple with the SignatureInitials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureInitialsOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureInitials) {
		return nil, false
	}
	return o.SignatureInitials, true
}

// HasSignatureInitials returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureInitials() bool {
	if o != nil && !IsNil(o.SignatureInitials) {
		return true
	}

	return false
}

// SetSignatureInitials gets a reference to the given string and assigns it to the SignatureInitials field.
func (o *AccountSignatureDefinition) SetSignatureInitials(v string) {
	o.SignatureInitials = &v
}

// GetSignatureName returns the SignatureName field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureName() string {
	if o == nil || IsNil(o.SignatureName) {
		var ret string
		return ret
	}
	return *o.SignatureName
}

// GetSignatureNameOk returns a tuple with the SignatureName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureNameOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureName) {
		return nil, false
	}
	return o.SignatureName, true
}

// HasSignatureName returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureName() bool {
	if o != nil && !IsNil(o.SignatureName) {
		return true
	}

	return false
}

// SetSignatureName gets a reference to the given string and assigns it to the SignatureName field.
func (o *AccountSignatureDefinition) SetSignatureName(v string) {
	o.SignatureName = &v
}

// GetSignatureType returns the SignatureType field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureType() string {
	if o == nil || IsNil(o.SignatureType) {
		var ret string
		return ret
	}
	return *o.SignatureType
}

// GetSignatureTypeOk returns a tuple with the SignatureType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureType) {
		return nil, false
	}
	return o.SignatureType, true
}

// HasSignatureType returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureType() bool {
	if o != nil && !IsNil(o.SignatureType) {
		return true
	}

	return false
}

// SetSignatureType gets a reference to the given string and assigns it to the SignatureType field.
func (o *AccountSignatureDefinition) SetSignatureType(v string) {
	o.SignatureType = &v
}

// GetSignatureUsers returns the SignatureUsers field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetSignatureUsers() []SignatureUserDef {
	if o == nil || IsNil(o.SignatureUsers) {
		var ret []SignatureUserDef
		return ret
	}
	return o.SignatureUsers
}

// GetSignatureUsersOk returns a tuple with the SignatureUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetSignatureUsersOk() ([]SignatureUserDef, bool) {
	if o == nil || IsNil(o.SignatureUsers) {
		return nil, false
	}
	return o.SignatureUsers, true
}

// HasSignatureUsers returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasSignatureUsers() bool {
	if o != nil && !IsNil(o.SignatureUsers) {
		return true
	}

	return false
}

// SetSignatureUsers gets a reference to the given []SignatureUserDef and assigns it to the SignatureUsers field.
func (o *AccountSignatureDefinition) SetSignatureUsers(v []SignatureUserDef) {
	o.SignatureUsers = v
}

// GetStampFormat returns the StampFormat field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetStampFormat() string {
	if o == nil || IsNil(o.StampFormat) {
		var ret string
		return ret
	}
	return *o.StampFormat
}

// GetStampFormatOk returns a tuple with the StampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetStampFormatOk() (*string, bool) {
	if o == nil || IsNil(o.StampFormat) {
		return nil, false
	}
	return o.StampFormat, true
}

// HasStampFormat returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasStampFormat() bool {
	if o != nil && !IsNil(o.StampFormat) {
		return true
	}

	return false
}

// SetStampFormat gets a reference to the given string and assigns it to the StampFormat field.
func (o *AccountSignatureDefinition) SetStampFormat(v string) {
	o.StampFormat = &v
}

// GetStampSizeMM returns the StampSizeMM field value if set, zero value otherwise.
func (o *AccountSignatureDefinition) GetStampSizeMM() string {
	if o == nil || IsNil(o.StampSizeMM) {
		var ret string
		return ret
	}
	return *o.StampSizeMM
}

// GetStampSizeMMOk returns a tuple with the StampSizeMM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountSignatureDefinition) GetStampSizeMMOk() (*string, bool) {
	if o == nil || IsNil(o.StampSizeMM) {
		return nil, false
	}
	return o.StampSizeMM, true
}

// HasStampSizeMM returns a boolean if a field has been set.
func (o *AccountSignatureDefinition) HasStampSizeMM() bool {
	if o != nil && !IsNil(o.StampSizeMM) {
		return true
	}

	return false
}

// SetStampSizeMM gets a reference to the given string and assigns it to the StampSizeMM field.
func (o *AccountSignatureDefinition) SetStampSizeMM(v string) {
	o.StampSizeMM = &v
}

func (o AccountSignatureDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountSignatureDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DateStampProperties) {
		toSerialize["dateStampProperties"] = o.DateStampProperties
	}
	if !IsNil(o.DisallowUserResizeStamp) {
		toSerialize["disallowUserResizeStamp"] = o.DisallowUserResizeStamp
	}
	if !IsNil(o.ExternalID) {
		toSerialize["externalID"] = o.ExternalID
	}
	if !IsNil(o.ImageType) {
		toSerialize["imageType"] = o.ImageType
	}
	if !IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	if !IsNil(o.NrdsId) {
		toSerialize["nrdsId"] = o.NrdsId
	}
	if !IsNil(o.NrdsLastName) {
		toSerialize["nrdsLastName"] = o.NrdsLastName
	}
	if !IsNil(o.PhoneticName) {
		toSerialize["phoneticName"] = o.PhoneticName
	}
	if !IsNil(o.SignatureFont) {
		toSerialize["signatureFont"] = o.SignatureFont
	}
	if !IsNil(o.SignatureGroups) {
		toSerialize["signatureGroups"] = o.SignatureGroups
	}
	if !IsNil(o.SignatureId) {
		toSerialize["signatureId"] = o.SignatureId
	}
	if !IsNil(o.SignatureInitials) {
		toSerialize["signatureInitials"] = o.SignatureInitials
	}
	if !IsNil(o.SignatureName) {
		toSerialize["signatureName"] = o.SignatureName
	}
	if !IsNil(o.SignatureType) {
		toSerialize["signatureType"] = o.SignatureType
	}
	if !IsNil(o.SignatureUsers) {
		toSerialize["signatureUsers"] = o.SignatureUsers
	}
	if !IsNil(o.StampFormat) {
		toSerialize["stampFormat"] = o.StampFormat
	}
	if !IsNil(o.StampSizeMM) {
		toSerialize["stampSizeMM"] = o.StampSizeMM
	}
	return toSerialize, nil
}

type NullableAccountSignatureDefinition struct {
	value *AccountSignatureDefinition
	isSet bool
}

func (v NullableAccountSignatureDefinition) Get() *AccountSignatureDefinition {
	return v.value
}

func (v *NullableAccountSignatureDefinition) Set(val *AccountSignatureDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSignatureDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSignatureDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSignatureDefinition(val *AccountSignatureDefinition) *NullableAccountSignatureDefinition {
	return &NullableAccountSignatureDefinition{value: val, isSet: true}
}

func (v NullableAccountSignatureDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSignatureDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


