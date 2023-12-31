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

// checks if the NotaryJournalMetaData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotaryJournalMetaData{}

// NotaryJournalMetaData 
type NotaryJournalMetaData struct {
	// A freeform comment that the notary can add to the journal entry.
	Comment *string `json:"comment,omitempty"`
	// An array of witnesses.
	CredibleWitnesses []NotaryJournalCredibleWitness `json:"credibleWitnesses,omitempty"`
	// A base64-encoded image of the signature.
	SignatureImage *string `json:"signatureImage,omitempty"`
	// A string that describes the ID that the signer presented. For example `drivers license` or `military ID`.
	SignerIdType *string `json:"signerIdType,omitempty"`
}

// NewNotaryJournalMetaData instantiates a new NotaryJournalMetaData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotaryJournalMetaData() *NotaryJournalMetaData {
	this := NotaryJournalMetaData{}
	return &this
}

// NewNotaryJournalMetaDataWithDefaults instantiates a new NotaryJournalMetaData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotaryJournalMetaDataWithDefaults() *NotaryJournalMetaData {
	this := NotaryJournalMetaData{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NotaryJournalMetaData) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotaryJournalMetaData) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NotaryJournalMetaData) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NotaryJournalMetaData) SetComment(v string) {
	o.Comment = &v
}

// GetCredibleWitnesses returns the CredibleWitnesses field value if set, zero value otherwise.
func (o *NotaryJournalMetaData) GetCredibleWitnesses() []NotaryJournalCredibleWitness {
	if o == nil || IsNil(o.CredibleWitnesses) {
		var ret []NotaryJournalCredibleWitness
		return ret
	}
	return o.CredibleWitnesses
}

// GetCredibleWitnessesOk returns a tuple with the CredibleWitnesses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotaryJournalMetaData) GetCredibleWitnessesOk() ([]NotaryJournalCredibleWitness, bool) {
	if o == nil || IsNil(o.CredibleWitnesses) {
		return nil, false
	}
	return o.CredibleWitnesses, true
}

// HasCredibleWitnesses returns a boolean if a field has been set.
func (o *NotaryJournalMetaData) HasCredibleWitnesses() bool {
	if o != nil && !IsNil(o.CredibleWitnesses) {
		return true
	}

	return false
}

// SetCredibleWitnesses gets a reference to the given []NotaryJournalCredibleWitness and assigns it to the CredibleWitnesses field.
func (o *NotaryJournalMetaData) SetCredibleWitnesses(v []NotaryJournalCredibleWitness) {
	o.CredibleWitnesses = v
}

// GetSignatureImage returns the SignatureImage field value if set, zero value otherwise.
func (o *NotaryJournalMetaData) GetSignatureImage() string {
	if o == nil || IsNil(o.SignatureImage) {
		var ret string
		return ret
	}
	return *o.SignatureImage
}

// GetSignatureImageOk returns a tuple with the SignatureImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotaryJournalMetaData) GetSignatureImageOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureImage) {
		return nil, false
	}
	return o.SignatureImage, true
}

// HasSignatureImage returns a boolean if a field has been set.
func (o *NotaryJournalMetaData) HasSignatureImage() bool {
	if o != nil && !IsNil(o.SignatureImage) {
		return true
	}

	return false
}

// SetSignatureImage gets a reference to the given string and assigns it to the SignatureImage field.
func (o *NotaryJournalMetaData) SetSignatureImage(v string) {
	o.SignatureImage = &v
}

// GetSignerIdType returns the SignerIdType field value if set, zero value otherwise.
func (o *NotaryJournalMetaData) GetSignerIdType() string {
	if o == nil || IsNil(o.SignerIdType) {
		var ret string
		return ret
	}
	return *o.SignerIdType
}

// GetSignerIdTypeOk returns a tuple with the SignerIdType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotaryJournalMetaData) GetSignerIdTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SignerIdType) {
		return nil, false
	}
	return o.SignerIdType, true
}

// HasSignerIdType returns a boolean if a field has been set.
func (o *NotaryJournalMetaData) HasSignerIdType() bool {
	if o != nil && !IsNil(o.SignerIdType) {
		return true
	}

	return false
}

// SetSignerIdType gets a reference to the given string and assigns it to the SignerIdType field.
func (o *NotaryJournalMetaData) SetSignerIdType(v string) {
	o.SignerIdType = &v
}

func (o NotaryJournalMetaData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotaryJournalMetaData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CredibleWitnesses) {
		toSerialize["credibleWitnesses"] = o.CredibleWitnesses
	}
	if !IsNil(o.SignatureImage) {
		toSerialize["signatureImage"] = o.SignatureImage
	}
	if !IsNil(o.SignerIdType) {
		toSerialize["signerIdType"] = o.SignerIdType
	}
	return toSerialize, nil
}

type NullableNotaryJournalMetaData struct {
	value *NotaryJournalMetaData
	isSet bool
}

func (v NullableNotaryJournalMetaData) Get() *NotaryJournalMetaData {
	return v.value
}

func (v *NullableNotaryJournalMetaData) Set(val *NotaryJournalMetaData) {
	v.value = val
	v.isSet = true
}

func (v NullableNotaryJournalMetaData) IsSet() bool {
	return v.isSet
}

func (v *NullableNotaryJournalMetaData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotaryJournalMetaData(val *NotaryJournalMetaData) *NullableNotaryJournalMetaData {
	return &NullableNotaryJournalMetaData{value: val, isSet: true}
}

func (v NullableNotaryJournalMetaData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotaryJournalMetaData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


