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

// checks if the EnvelopeAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeAttachment{}

// EnvelopeAttachment 
type EnvelopeAttachment struct {
	// Valid values are `sender` and `senderAndAllRecipients`.
	AccessControl *string `json:"accessControl,omitempty"`
	// The unique identifier for the attachment.
	AttachmentId *string `json:"attachmentId,omitempty"`
	// Specifies the type of the attachment for the recipient. Possible values are:  - `.htm` - `.xml`
	AttachmentType *string `json:"attachmentType,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// 
	Label *string `json:"label,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
}

// NewEnvelopeAttachment instantiates a new EnvelopeAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeAttachment() *EnvelopeAttachment {
	this := EnvelopeAttachment{}
	return &this
}

// NewEnvelopeAttachmentWithDefaults instantiates a new EnvelopeAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeAttachmentWithDefaults() *EnvelopeAttachment {
	this := EnvelopeAttachment{}
	return &this
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetAccessControl() string {
	if o == nil || IsNil(o.AccessControl) {
		var ret string
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetAccessControlOk() (*string, bool) {
	if o == nil || IsNil(o.AccessControl) {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasAccessControl() bool {
	if o != nil && !IsNil(o.AccessControl) {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given string and assigns it to the AccessControl field.
func (o *EnvelopeAttachment) SetAccessControl(v string) {
	o.AccessControl = &v
}

// GetAttachmentId returns the AttachmentId field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetAttachmentId() string {
	if o == nil || IsNil(o.AttachmentId) {
		var ret string
		return ret
	}
	return *o.AttachmentId
}

// GetAttachmentIdOk returns a tuple with the AttachmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetAttachmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentId) {
		return nil, false
	}
	return o.AttachmentId, true
}

// HasAttachmentId returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasAttachmentId() bool {
	if o != nil && !IsNil(o.AttachmentId) {
		return true
	}

	return false
}

// SetAttachmentId gets a reference to the given string and assigns it to the AttachmentId field.
func (o *EnvelopeAttachment) SetAttachmentId(v string) {
	o.AttachmentId = &v
}

// GetAttachmentType returns the AttachmentType field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetAttachmentType() string {
	if o == nil || IsNil(o.AttachmentType) {
		var ret string
		return ret
	}
	return *o.AttachmentType
}

// GetAttachmentTypeOk returns a tuple with the AttachmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetAttachmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentType) {
		return nil, false
	}
	return o.AttachmentType, true
}

// HasAttachmentType returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasAttachmentType() bool {
	if o != nil && !IsNil(o.AttachmentType) {
		return true
	}

	return false
}

// SetAttachmentType gets a reference to the given string and assigns it to the AttachmentType field.
func (o *EnvelopeAttachment) SetAttachmentType(v string) {
	o.AttachmentType = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *EnvelopeAttachment) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EnvelopeAttachment) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvelopeAttachment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeAttachment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvelopeAttachment) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvelopeAttachment) SetName(v string) {
	o.Name = &v
}

func (o EnvelopeAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessControl) {
		toSerialize["accessControl"] = o.AccessControl
	}
	if !IsNil(o.AttachmentId) {
		toSerialize["attachmentId"] = o.AttachmentId
	}
	if !IsNil(o.AttachmentType) {
		toSerialize["attachmentType"] = o.AttachmentType
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEnvelopeAttachment struct {
	value *EnvelopeAttachment
	isSet bool
}

func (v NullableEnvelopeAttachment) Get() *EnvelopeAttachment {
	return v.value
}

func (v *NullableEnvelopeAttachment) Set(val *EnvelopeAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeAttachment(val *EnvelopeAttachment) *NullableEnvelopeAttachment {
	return &NullableEnvelopeAttachment{value: val, isSet: true}
}

func (v NullableEnvelopeAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


