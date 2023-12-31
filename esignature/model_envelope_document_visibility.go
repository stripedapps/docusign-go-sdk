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

// checks if the EnvelopeDocumentVisibility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvelopeDocumentVisibility{}

// EnvelopeDocumentVisibility Document Visibility enables senders to control the visibility of the documents in an envelope at the recipient level. For example, if the parties associated with a legal proceeding should have access to different documents, the Document Visibility feature enables you to keep all of the documents in the same envelope and set view permissions for the documents by recipient. This functionality is enabled for envelopes and templates. It is not available for PowerForms.  **Note:** Before you use Document Visibility, you should be aware of the following information:  - Document Visibility must be enabled for your account by your DocuSign administrator.  - A document cannot be hidden from a recipient if the recipient has tabs assigned to them on the document.  - When the Document Visibility setting hides a document from a recipient, the document also does not appear in the recipient's list of envelopes, documents, or page images. - Carbon Copy, Certified Delivery (Needs to Sign), Editor, and Agent recipients can always see all of the documents associated with the envelope or template.  The Document Visibility feature has multiple settings that specify the options that senders have when sending documents. For more information, see [Use Document Visibility to Control Recipient Access](https://support.docusign.com/s/document-item?bundleId=gbo1643332197980&topicId=eui1578456411411.html).
type EnvelopeDocumentVisibility struct {
	// An array of `documentVisibility` objects that specifies which documents are visible to which recipients.
	DocumentVisibility []DocumentVisibility `json:"documentVisibility,omitempty"`
}

// NewEnvelopeDocumentVisibility instantiates a new EnvelopeDocumentVisibility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvelopeDocumentVisibility() *EnvelopeDocumentVisibility {
	this := EnvelopeDocumentVisibility{}
	return &this
}

// NewEnvelopeDocumentVisibilityWithDefaults instantiates a new EnvelopeDocumentVisibility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvelopeDocumentVisibilityWithDefaults() *EnvelopeDocumentVisibility {
	this := EnvelopeDocumentVisibility{}
	return &this
}

// GetDocumentVisibility returns the DocumentVisibility field value if set, zero value otherwise.
func (o *EnvelopeDocumentVisibility) GetDocumentVisibility() []DocumentVisibility {
	if o == nil || IsNil(o.DocumentVisibility) {
		var ret []DocumentVisibility
		return ret
	}
	return o.DocumentVisibility
}

// GetDocumentVisibilityOk returns a tuple with the DocumentVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvelopeDocumentVisibility) GetDocumentVisibilityOk() ([]DocumentVisibility, bool) {
	if o == nil || IsNil(o.DocumentVisibility) {
		return nil, false
	}
	return o.DocumentVisibility, true
}

// HasDocumentVisibility returns a boolean if a field has been set.
func (o *EnvelopeDocumentVisibility) HasDocumentVisibility() bool {
	if o != nil && !IsNil(o.DocumentVisibility) {
		return true
	}

	return false
}

// SetDocumentVisibility gets a reference to the given []DocumentVisibility and assigns it to the DocumentVisibility field.
func (o *EnvelopeDocumentVisibility) SetDocumentVisibility(v []DocumentVisibility) {
	o.DocumentVisibility = v
}

func (o EnvelopeDocumentVisibility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvelopeDocumentVisibility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentVisibility) {
		toSerialize["documentVisibility"] = o.DocumentVisibility
	}
	return toSerialize, nil
}

type NullableEnvelopeDocumentVisibility struct {
	value *EnvelopeDocumentVisibility
	isSet bool
}

func (v NullableEnvelopeDocumentVisibility) Get() *EnvelopeDocumentVisibility {
	return v.value
}

func (v *NullableEnvelopeDocumentVisibility) Set(val *EnvelopeDocumentVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvelopeDocumentVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvelopeDocumentVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvelopeDocumentVisibility(val *EnvelopeDocumentVisibility) *NullableEnvelopeDocumentVisibility {
	return &NullableEnvelopeDocumentVisibility{value: val, isSet: true}
}

func (v NullableEnvelopeDocumentVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvelopeDocumentVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


