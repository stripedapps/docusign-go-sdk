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

// checks if the ChunkedUploads type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChunkedUploads{}

// ChunkedUploads The ChunkedUploads resource provides methods to complete integrity checks, and to add, commit, retrieve, initiate and delete chunked uploads.
type ChunkedUploads struct {
	// A 64-byte, Secure Hash Algorithm 256 (SHA256) checksum that the caller computes across the entirety of the original content that has been uploaded to the chunked upload. DocuSign compares this value to its own computation. If the two values are not equal, the original content and received content are not the same and the commit action is refused.
	Checksum *string `json:"checksum,omitempty"`
	// The ID of the chunked upload. 
	ChunkedUploadId *string `json:"chunkedUploadId,omitempty"`
	// A list of the parts that compose the chunked upload, including their byte sizes. The list must be contiguous before you can commit the chunked upload.
	ChunkedUploadParts []ChunkedUploadPart `json:"chunkedUploadParts,omitempty"`
	// The URI that you use to reference the chunked upload in other API requests, such as envelope document and envelope attachment requests. 
	ChunkedUploadUri *string `json:"chunkedUploadUri,omitempty"`
	// When **true,** the chunked upload has been committed. A committed chunked upload can no longer receive any additional parts and is ready for use within other API requests. 
	Committed *string `json:"committed,omitempty"`
	// The UTC time at which the chunked upload expires and is no longer addressable.   **Note:** The length of time before expiration is configurable, and begins when you initiate the chunked upload. You must fully upload and use a chunked upload within this time. The default value for this duration is 20 minutes.
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
	// The maximum number of parts allowed for a chunked upload. This value is configurable per DocuSign environment, account, or integrator. The default value is 128. The maximum possible value is 256.   
	MaxChunkedUploadParts *string `json:"maxChunkedUploadParts,omitempty"`
	// The maximum total size allowed for a chunked upload. This value is configured per DocuSign environment, account, or integrator. The default value is 50 MB.
	MaxTotalSize *string `json:"maxTotalSize,omitempty"`
	// The total size of the parts of the chunked upload.  **Note:** When a chunked upload is used as an envelope document, it is subject to the PDF size limit (25 MB) and page count limit that apply to all envelope documents.
	TotalSize *string `json:"totalSize,omitempty"`
}

// NewChunkedUploads instantiates a new ChunkedUploads object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkedUploads() *ChunkedUploads {
	this := ChunkedUploads{}
	return &this
}

// NewChunkedUploadsWithDefaults instantiates a new ChunkedUploads object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkedUploadsWithDefaults() *ChunkedUploads {
	this := ChunkedUploads{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *ChunkedUploads) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *ChunkedUploads) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *ChunkedUploads) SetChecksum(v string) {
	o.Checksum = &v
}

// GetChunkedUploadId returns the ChunkedUploadId field value if set, zero value otherwise.
func (o *ChunkedUploads) GetChunkedUploadId() string {
	if o == nil || IsNil(o.ChunkedUploadId) {
		var ret string
		return ret
	}
	return *o.ChunkedUploadId
}

// GetChunkedUploadIdOk returns a tuple with the ChunkedUploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetChunkedUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChunkedUploadId) {
		return nil, false
	}
	return o.ChunkedUploadId, true
}

// HasChunkedUploadId returns a boolean if a field has been set.
func (o *ChunkedUploads) HasChunkedUploadId() bool {
	if o != nil && !IsNil(o.ChunkedUploadId) {
		return true
	}

	return false
}

// SetChunkedUploadId gets a reference to the given string and assigns it to the ChunkedUploadId field.
func (o *ChunkedUploads) SetChunkedUploadId(v string) {
	o.ChunkedUploadId = &v
}

// GetChunkedUploadParts returns the ChunkedUploadParts field value if set, zero value otherwise.
func (o *ChunkedUploads) GetChunkedUploadParts() []ChunkedUploadPart {
	if o == nil || IsNil(o.ChunkedUploadParts) {
		var ret []ChunkedUploadPart
		return ret
	}
	return o.ChunkedUploadParts
}

// GetChunkedUploadPartsOk returns a tuple with the ChunkedUploadParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetChunkedUploadPartsOk() ([]ChunkedUploadPart, bool) {
	if o == nil || IsNil(o.ChunkedUploadParts) {
		return nil, false
	}
	return o.ChunkedUploadParts, true
}

// HasChunkedUploadParts returns a boolean if a field has been set.
func (o *ChunkedUploads) HasChunkedUploadParts() bool {
	if o != nil && !IsNil(o.ChunkedUploadParts) {
		return true
	}

	return false
}

// SetChunkedUploadParts gets a reference to the given []ChunkedUploadPart and assigns it to the ChunkedUploadParts field.
func (o *ChunkedUploads) SetChunkedUploadParts(v []ChunkedUploadPart) {
	o.ChunkedUploadParts = v
}

// GetChunkedUploadUri returns the ChunkedUploadUri field value if set, zero value otherwise.
func (o *ChunkedUploads) GetChunkedUploadUri() string {
	if o == nil || IsNil(o.ChunkedUploadUri) {
		var ret string
		return ret
	}
	return *o.ChunkedUploadUri
}

// GetChunkedUploadUriOk returns a tuple with the ChunkedUploadUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetChunkedUploadUriOk() (*string, bool) {
	if o == nil || IsNil(o.ChunkedUploadUri) {
		return nil, false
	}
	return o.ChunkedUploadUri, true
}

// HasChunkedUploadUri returns a boolean if a field has been set.
func (o *ChunkedUploads) HasChunkedUploadUri() bool {
	if o != nil && !IsNil(o.ChunkedUploadUri) {
		return true
	}

	return false
}

// SetChunkedUploadUri gets a reference to the given string and assigns it to the ChunkedUploadUri field.
func (o *ChunkedUploads) SetChunkedUploadUri(v string) {
	o.ChunkedUploadUri = &v
}

// GetCommitted returns the Committed field value if set, zero value otherwise.
func (o *ChunkedUploads) GetCommitted() string {
	if o == nil || IsNil(o.Committed) {
		var ret string
		return ret
	}
	return *o.Committed
}

// GetCommittedOk returns a tuple with the Committed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetCommittedOk() (*string, bool) {
	if o == nil || IsNil(o.Committed) {
		return nil, false
	}
	return o.Committed, true
}

// HasCommitted returns a boolean if a field has been set.
func (o *ChunkedUploads) HasCommitted() bool {
	if o != nil && !IsNil(o.Committed) {
		return true
	}

	return false
}

// SetCommitted gets a reference to the given string and assigns it to the Committed field.
func (o *ChunkedUploads) SetCommitted(v string) {
	o.Committed = &v
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *ChunkedUploads) GetExpirationDateTime() string {
	if o == nil || IsNil(o.ExpirationDateTime) {
		var ret string
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetExpirationDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDateTime) {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *ChunkedUploads) HasExpirationDateTime() bool {
	if o != nil && !IsNil(o.ExpirationDateTime) {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given string and assigns it to the ExpirationDateTime field.
func (o *ChunkedUploads) SetExpirationDateTime(v string) {
	o.ExpirationDateTime = &v
}

// GetMaxChunkedUploadParts returns the MaxChunkedUploadParts field value if set, zero value otherwise.
func (o *ChunkedUploads) GetMaxChunkedUploadParts() string {
	if o == nil || IsNil(o.MaxChunkedUploadParts) {
		var ret string
		return ret
	}
	return *o.MaxChunkedUploadParts
}

// GetMaxChunkedUploadPartsOk returns a tuple with the MaxChunkedUploadParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetMaxChunkedUploadPartsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxChunkedUploadParts) {
		return nil, false
	}
	return o.MaxChunkedUploadParts, true
}

// HasMaxChunkedUploadParts returns a boolean if a field has been set.
func (o *ChunkedUploads) HasMaxChunkedUploadParts() bool {
	if o != nil && !IsNil(o.MaxChunkedUploadParts) {
		return true
	}

	return false
}

// SetMaxChunkedUploadParts gets a reference to the given string and assigns it to the MaxChunkedUploadParts field.
func (o *ChunkedUploads) SetMaxChunkedUploadParts(v string) {
	o.MaxChunkedUploadParts = &v
}

// GetMaxTotalSize returns the MaxTotalSize field value if set, zero value otherwise.
func (o *ChunkedUploads) GetMaxTotalSize() string {
	if o == nil || IsNil(o.MaxTotalSize) {
		var ret string
		return ret
	}
	return *o.MaxTotalSize
}

// GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetMaxTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxTotalSize) {
		return nil, false
	}
	return o.MaxTotalSize, true
}

// HasMaxTotalSize returns a boolean if a field has been set.
func (o *ChunkedUploads) HasMaxTotalSize() bool {
	if o != nil && !IsNil(o.MaxTotalSize) {
		return true
	}

	return false
}

// SetMaxTotalSize gets a reference to the given string and assigns it to the MaxTotalSize field.
func (o *ChunkedUploads) SetMaxTotalSize(v string) {
	o.MaxTotalSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ChunkedUploads) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploads) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ChunkedUploads) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *ChunkedUploads) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o ChunkedUploads) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChunkedUploads) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checksum) {
		toSerialize["checksum"] = o.Checksum
	}
	if !IsNil(o.ChunkedUploadId) {
		toSerialize["chunkedUploadId"] = o.ChunkedUploadId
	}
	if !IsNil(o.ChunkedUploadParts) {
		toSerialize["chunkedUploadParts"] = o.ChunkedUploadParts
	}
	if !IsNil(o.ChunkedUploadUri) {
		toSerialize["chunkedUploadUri"] = o.ChunkedUploadUri
	}
	if !IsNil(o.Committed) {
		toSerialize["committed"] = o.Committed
	}
	if !IsNil(o.ExpirationDateTime) {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if !IsNil(o.MaxChunkedUploadParts) {
		toSerialize["maxChunkedUploadParts"] = o.MaxChunkedUploadParts
	}
	if !IsNil(o.MaxTotalSize) {
		toSerialize["maxTotalSize"] = o.MaxTotalSize
	}
	if !IsNil(o.TotalSize) {
		toSerialize["totalSize"] = o.TotalSize
	}
	return toSerialize, nil
}

type NullableChunkedUploads struct {
	value *ChunkedUploads
	isSet bool
}

func (v NullableChunkedUploads) Get() *ChunkedUploads {
	return v.value
}

func (v *NullableChunkedUploads) Set(val *ChunkedUploads) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkedUploads) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkedUploads) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkedUploads(val *ChunkedUploads) *NullableChunkedUploads {
	return &NullableChunkedUploads{value: val, isSet: true}
}

func (v NullableChunkedUploads) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkedUploads) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


