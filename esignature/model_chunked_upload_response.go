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

// checks if the ChunkedUploadResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChunkedUploadResponse{}

// ChunkedUploadResponse This response object is returned after you upload a chunked upload.
type ChunkedUploadResponse struct {
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
	// The UTC time at which the chunked upload expires and is no longer addressable.   **Note:** You must fully upload and use a chunked upload within 20 minutes of initializing it. 
	ExpirationDateTime *string `json:"expirationDateTime,omitempty"`
	// The maximum number of parts allowed for a chunked upload. This value is configurable per DocuSign environment, account, or integrator. The default value is 128. The maximum possible value is 256.  
	MaxChunkedUploadParts *string `json:"maxChunkedUploadParts,omitempty"`
	// The maximum total size allowed for a chunked upload. This value is configured per DocuSign environment, account, or integrator. The default value is 50 MB.
	MaxTotalSize *string `json:"maxTotalSize,omitempty"`
	// The total size of the parts of the chunked upload.  **Note:** When a chunked upload is used as an envelope document, it is subject to the PDF size limit (25 MB) and page count limit that apply to all envelope documents.
	TotalSize *string `json:"totalSize,omitempty"`
}

// NewChunkedUploadResponse instantiates a new ChunkedUploadResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChunkedUploadResponse() *ChunkedUploadResponse {
	this := ChunkedUploadResponse{}
	return &this
}

// NewChunkedUploadResponseWithDefaults instantiates a new ChunkedUploadResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChunkedUploadResponseWithDefaults() *ChunkedUploadResponse {
	this := ChunkedUploadResponse{}
	return &this
}

// GetChecksum returns the Checksum field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetChecksum() string {
	if o == nil || IsNil(o.Checksum) {
		var ret string
		return ret
	}
	return *o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetChecksumOk() (*string, bool) {
	if o == nil || IsNil(o.Checksum) {
		return nil, false
	}
	return o.Checksum, true
}

// HasChecksum returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasChecksum() bool {
	if o != nil && !IsNil(o.Checksum) {
		return true
	}

	return false
}

// SetChecksum gets a reference to the given string and assigns it to the Checksum field.
func (o *ChunkedUploadResponse) SetChecksum(v string) {
	o.Checksum = &v
}

// GetChunkedUploadId returns the ChunkedUploadId field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetChunkedUploadId() string {
	if o == nil || IsNil(o.ChunkedUploadId) {
		var ret string
		return ret
	}
	return *o.ChunkedUploadId
}

// GetChunkedUploadIdOk returns a tuple with the ChunkedUploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetChunkedUploadIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChunkedUploadId) {
		return nil, false
	}
	return o.ChunkedUploadId, true
}

// HasChunkedUploadId returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasChunkedUploadId() bool {
	if o != nil && !IsNil(o.ChunkedUploadId) {
		return true
	}

	return false
}

// SetChunkedUploadId gets a reference to the given string and assigns it to the ChunkedUploadId field.
func (o *ChunkedUploadResponse) SetChunkedUploadId(v string) {
	o.ChunkedUploadId = &v
}

// GetChunkedUploadParts returns the ChunkedUploadParts field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetChunkedUploadParts() []ChunkedUploadPart {
	if o == nil || IsNil(o.ChunkedUploadParts) {
		var ret []ChunkedUploadPart
		return ret
	}
	return o.ChunkedUploadParts
}

// GetChunkedUploadPartsOk returns a tuple with the ChunkedUploadParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetChunkedUploadPartsOk() ([]ChunkedUploadPart, bool) {
	if o == nil || IsNil(o.ChunkedUploadParts) {
		return nil, false
	}
	return o.ChunkedUploadParts, true
}

// HasChunkedUploadParts returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasChunkedUploadParts() bool {
	if o != nil && !IsNil(o.ChunkedUploadParts) {
		return true
	}

	return false
}

// SetChunkedUploadParts gets a reference to the given []ChunkedUploadPart and assigns it to the ChunkedUploadParts field.
func (o *ChunkedUploadResponse) SetChunkedUploadParts(v []ChunkedUploadPart) {
	o.ChunkedUploadParts = v
}

// GetChunkedUploadUri returns the ChunkedUploadUri field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetChunkedUploadUri() string {
	if o == nil || IsNil(o.ChunkedUploadUri) {
		var ret string
		return ret
	}
	return *o.ChunkedUploadUri
}

// GetChunkedUploadUriOk returns a tuple with the ChunkedUploadUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetChunkedUploadUriOk() (*string, bool) {
	if o == nil || IsNil(o.ChunkedUploadUri) {
		return nil, false
	}
	return o.ChunkedUploadUri, true
}

// HasChunkedUploadUri returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasChunkedUploadUri() bool {
	if o != nil && !IsNil(o.ChunkedUploadUri) {
		return true
	}

	return false
}

// SetChunkedUploadUri gets a reference to the given string and assigns it to the ChunkedUploadUri field.
func (o *ChunkedUploadResponse) SetChunkedUploadUri(v string) {
	o.ChunkedUploadUri = &v
}

// GetCommitted returns the Committed field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetCommitted() string {
	if o == nil || IsNil(o.Committed) {
		var ret string
		return ret
	}
	return *o.Committed
}

// GetCommittedOk returns a tuple with the Committed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetCommittedOk() (*string, bool) {
	if o == nil || IsNil(o.Committed) {
		return nil, false
	}
	return o.Committed, true
}

// HasCommitted returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasCommitted() bool {
	if o != nil && !IsNil(o.Committed) {
		return true
	}

	return false
}

// SetCommitted gets a reference to the given string and assigns it to the Committed field.
func (o *ChunkedUploadResponse) SetCommitted(v string) {
	o.Committed = &v
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetExpirationDateTime() string {
	if o == nil || IsNil(o.ExpirationDateTime) {
		var ret string
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetExpirationDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDateTime) {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasExpirationDateTime() bool {
	if o != nil && !IsNil(o.ExpirationDateTime) {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given string and assigns it to the ExpirationDateTime field.
func (o *ChunkedUploadResponse) SetExpirationDateTime(v string) {
	o.ExpirationDateTime = &v
}

// GetMaxChunkedUploadParts returns the MaxChunkedUploadParts field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetMaxChunkedUploadParts() string {
	if o == nil || IsNil(o.MaxChunkedUploadParts) {
		var ret string
		return ret
	}
	return *o.MaxChunkedUploadParts
}

// GetMaxChunkedUploadPartsOk returns a tuple with the MaxChunkedUploadParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetMaxChunkedUploadPartsOk() (*string, bool) {
	if o == nil || IsNil(o.MaxChunkedUploadParts) {
		return nil, false
	}
	return o.MaxChunkedUploadParts, true
}

// HasMaxChunkedUploadParts returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasMaxChunkedUploadParts() bool {
	if o != nil && !IsNil(o.MaxChunkedUploadParts) {
		return true
	}

	return false
}

// SetMaxChunkedUploadParts gets a reference to the given string and assigns it to the MaxChunkedUploadParts field.
func (o *ChunkedUploadResponse) SetMaxChunkedUploadParts(v string) {
	o.MaxChunkedUploadParts = &v
}

// GetMaxTotalSize returns the MaxTotalSize field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetMaxTotalSize() string {
	if o == nil || IsNil(o.MaxTotalSize) {
		var ret string
		return ret
	}
	return *o.MaxTotalSize
}

// GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetMaxTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.MaxTotalSize) {
		return nil, false
	}
	return o.MaxTotalSize, true
}

// HasMaxTotalSize returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasMaxTotalSize() bool {
	if o != nil && !IsNil(o.MaxTotalSize) {
		return true
	}

	return false
}

// SetMaxTotalSize gets a reference to the given string and assigns it to the MaxTotalSize field.
func (o *ChunkedUploadResponse) SetMaxTotalSize(v string) {
	o.MaxTotalSize = &v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ChunkedUploadResponse) GetTotalSize() string {
	if o == nil || IsNil(o.TotalSize) {
		var ret string
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChunkedUploadResponse) GetTotalSizeOk() (*string, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ChunkedUploadResponse) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given string and assigns it to the TotalSize field.
func (o *ChunkedUploadResponse) SetTotalSize(v string) {
	o.TotalSize = &v
}

func (o ChunkedUploadResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChunkedUploadResponse) ToMap() (map[string]interface{}, error) {
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

type NullableChunkedUploadResponse struct {
	value *ChunkedUploadResponse
	isSet bool
}

func (v NullableChunkedUploadResponse) Get() *ChunkedUploadResponse {
	return v.value
}

func (v *NullableChunkedUploadResponse) Set(val *ChunkedUploadResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChunkedUploadResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChunkedUploadResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChunkedUploadResponse(val *ChunkedUploadResponse) *NullableChunkedUploadResponse {
	return &NullableChunkedUploadResponse{value: val, isSet: true}
}

func (v NullableChunkedUploadResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChunkedUploadResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


