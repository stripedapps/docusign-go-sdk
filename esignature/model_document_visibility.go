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

// checks if the DocumentVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentVisibility{}

// DocumentVisibility This object configures a recipient's read/write access to a document.
type DocumentVisibility struct {
	// Specifies the document ID number that the tab is placed on. This must refer to an existing Document's ID attribute.
	DocumentId *string `json:"documentId,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
	RecipientId *string `json:"recipientId,omitempty"`
	// Indicates whether the document is editable:  - `editable` - `read_only`
	Rights *string `json:"rights,omitempty"`
	// When **true,** the document is visible to the recipient.
	Visible *string `json:"visible,omitempty"`
}

// NewDocumentVisibility instantiates a new DocumentVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentVisibility() *DocumentVisibility {
	this := DocumentVisibility{}
	return &this
}

// NewDocumentVisibilityWithDefaults instantiates a new DocumentVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentVisibilityWithDefaults() *DocumentVisibility {
	this := DocumentVisibility{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *DocumentVisibility) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVisibility) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *DocumentVisibility) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *DocumentVisibility) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *DocumentVisibility) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVisibility) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *DocumentVisibility) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *DocumentVisibility) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *DocumentVisibility) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVisibility) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *DocumentVisibility) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *DocumentVisibility) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetRights returns the Rights field value if set, zero value otherwise.
func (o *DocumentVisibility) GetRights() string {
	if o == nil || IsNil(o.Rights) {
		var ret string
		return ret
	}
	return *o.Rights
}

// GetRightsOk returns a tuple with the Rights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVisibility) GetRightsOk() (*string, bool) {
	if o == nil || IsNil(o.Rights) {
		return nil, false
	}
	return o.Rights, true
}

// HasRights returns a boolean if a field has been set.
func (o *DocumentVisibility) HasRights() bool {
	if o != nil && !IsNil(o.Rights) {
		return true
	}

	return false
}

// SetRights gets a reference to the given string and assigns it to the Rights field.
func (o *DocumentVisibility) SetRights(v string) {
	o.Rights = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *DocumentVisibility) GetVisible() string {
	if o == nil || IsNil(o.Visible) {
		var ret string
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentVisibility) GetVisibleOk() (*string, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *DocumentVisibility) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given string and assigns it to the Visible field.
func (o *DocumentVisibility) SetVisible(v string) {
	o.Visible = &v
}

func (o DocumentVisibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentId) {
		toSerialize["documentId"] = o.DocumentId
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.RecipientId) {
		toSerialize["recipientId"] = o.RecipientId
	}
	if !IsNil(o.Rights) {
		toSerialize["rights"] = o.Rights
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	return toSerialize, nil
}

type NullableDocumentVisibility struct {
	value *DocumentVisibility
	isSet bool
}

func (v NullableDocumentVisibility) Get() *DocumentVisibility {
	return v.value
}

func (v *NullableDocumentVisibility) Set(val *DocumentVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentVisibility(val *DocumentVisibility) *NullableDocumentVisibility {
	return &NullableDocumentVisibility{value: val, isSet: true}
}

func (v NullableDocumentVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


