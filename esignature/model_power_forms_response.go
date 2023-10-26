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

// checks if the PowerFormsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PowerFormsResponse{}

// PowerFormsResponse A list of PowerForms.
type PowerFormsResponse struct {
	// The last index position in the result set. 
	EndPosition *int32 `json:"endPosition,omitempty"`
	// The URI for the next chunk of records based on the search request. It is `null` if this is the last set of results for the search. 
	NextUri *string `json:"nextUri,omitempty"`
	// An array of PowerForm objects.
	PowerForms []PowerForm `json:"powerForms,omitempty"`
	// The URI for the prior chunk of records based on the search request. It is `null` if this is the first set of results for the search. 
	PreviousUri *string `json:"previousUri,omitempty"`
	// The number of results in this response. Because you can filter which entries are included in the response, this value is always less than or equal to the `totalSetSize`.
	ResultSetSize *int32 `json:"resultSetSize,omitempty"`
	// The starting index position of the current result set.
	StartPosition *int32 `json:"startPosition,omitempty"`
	// The total number of items in the result set. This value is always greater than or equal to the value of `resultSetSize`.
	TotalSetSize *int32 `json:"totalSetSize,omitempty"`
}

// NewPowerFormsResponse instantiates a new PowerFormsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPowerFormsResponse() *PowerFormsResponse {
	this := PowerFormsResponse{}
	return &this
}

// NewPowerFormsResponseWithDefaults instantiates a new PowerFormsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPowerFormsResponseWithDefaults() *PowerFormsResponse {
	this := PowerFormsResponse{}
	return &this
}

// GetEndPosition returns the EndPosition field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetEndPosition() int32 {
	if o == nil || IsNil(o.EndPosition) {
		var ret int32
		return ret
	}
	return *o.EndPosition
}

// GetEndPositionOk returns a tuple with the EndPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetEndPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.EndPosition) {
		return nil, false
	}
	return o.EndPosition, true
}

// HasEndPosition returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasEndPosition() bool {
	if o != nil && !IsNil(o.EndPosition) {
		return true
	}

	return false
}

// SetEndPosition gets a reference to the given int32 and assigns it to the EndPosition field.
func (o *PowerFormsResponse) SetEndPosition(v int32) {
	o.EndPosition = &v
}

// GetNextUri returns the NextUri field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetNextUri() string {
	if o == nil || IsNil(o.NextUri) {
		var ret string
		return ret
	}
	return *o.NextUri
}

// GetNextUriOk returns a tuple with the NextUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetNextUriOk() (*string, bool) {
	if o == nil || IsNil(o.NextUri) {
		return nil, false
	}
	return o.NextUri, true
}

// HasNextUri returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasNextUri() bool {
	if o != nil && !IsNil(o.NextUri) {
		return true
	}

	return false
}

// SetNextUri gets a reference to the given string and assigns it to the NextUri field.
func (o *PowerFormsResponse) SetNextUri(v string) {
	o.NextUri = &v
}

// GetPowerForms returns the PowerForms field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetPowerForms() []PowerForm {
	if o == nil || IsNil(o.PowerForms) {
		var ret []PowerForm
		return ret
	}
	return o.PowerForms
}

// GetPowerFormsOk returns a tuple with the PowerForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetPowerFormsOk() ([]PowerForm, bool) {
	if o == nil || IsNil(o.PowerForms) {
		return nil, false
	}
	return o.PowerForms, true
}

// HasPowerForms returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasPowerForms() bool {
	if o != nil && !IsNil(o.PowerForms) {
		return true
	}

	return false
}

// SetPowerForms gets a reference to the given []PowerForm and assigns it to the PowerForms field.
func (o *PowerFormsResponse) SetPowerForms(v []PowerForm) {
	o.PowerForms = v
}

// GetPreviousUri returns the PreviousUri field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetPreviousUri() string {
	if o == nil || IsNil(o.PreviousUri) {
		var ret string
		return ret
	}
	return *o.PreviousUri
}

// GetPreviousUriOk returns a tuple with the PreviousUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetPreviousUriOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousUri) {
		return nil, false
	}
	return o.PreviousUri, true
}

// HasPreviousUri returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasPreviousUri() bool {
	if o != nil && !IsNil(o.PreviousUri) {
		return true
	}

	return false
}

// SetPreviousUri gets a reference to the given string and assigns it to the PreviousUri field.
func (o *PowerFormsResponse) SetPreviousUri(v string) {
	o.PreviousUri = &v
}

// GetResultSetSize returns the ResultSetSize field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetResultSetSize() int32 {
	if o == nil || IsNil(o.ResultSetSize) {
		var ret int32
		return ret
	}
	return *o.ResultSetSize
}

// GetResultSetSizeOk returns a tuple with the ResultSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetResultSetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ResultSetSize) {
		return nil, false
	}
	return o.ResultSetSize, true
}

// HasResultSetSize returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasResultSetSize() bool {
	if o != nil && !IsNil(o.ResultSetSize) {
		return true
	}

	return false
}

// SetResultSetSize gets a reference to the given int32 and assigns it to the ResultSetSize field.
func (o *PowerFormsResponse) SetResultSetSize(v int32) {
	o.ResultSetSize = &v
}

// GetStartPosition returns the StartPosition field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetStartPosition() int32 {
	if o == nil || IsNil(o.StartPosition) {
		var ret int32
		return ret
	}
	return *o.StartPosition
}

// GetStartPositionOk returns a tuple with the StartPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetStartPositionOk() (*int32, bool) {
	if o == nil || IsNil(o.StartPosition) {
		return nil, false
	}
	return o.StartPosition, true
}

// HasStartPosition returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasStartPosition() bool {
	if o != nil && !IsNil(o.StartPosition) {
		return true
	}

	return false
}

// SetStartPosition gets a reference to the given int32 and assigns it to the StartPosition field.
func (o *PowerFormsResponse) SetStartPosition(v int32) {
	o.StartPosition = &v
}

// GetTotalSetSize returns the TotalSetSize field value if set, zero value otherwise.
func (o *PowerFormsResponse) GetTotalSetSize() int32 {
	if o == nil || IsNil(o.TotalSetSize) {
		var ret int32
		return ret
	}
	return *o.TotalSetSize
}

// GetTotalSetSizeOk returns a tuple with the TotalSetSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PowerFormsResponse) GetTotalSetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSetSize) {
		return nil, false
	}
	return o.TotalSetSize, true
}

// HasTotalSetSize returns a boolean if a field has been set.
func (o *PowerFormsResponse) HasTotalSetSize() bool {
	if o != nil && !IsNil(o.TotalSetSize) {
		return true
	}

	return false
}

// SetTotalSetSize gets a reference to the given int32 and assigns it to the TotalSetSize field.
func (o *PowerFormsResponse) SetTotalSetSize(v int32) {
	o.TotalSetSize = &v
}

func (o PowerFormsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PowerFormsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndPosition) {
		toSerialize["endPosition"] = o.EndPosition
	}
	if !IsNil(o.NextUri) {
		toSerialize["nextUri"] = o.NextUri
	}
	if !IsNil(o.PowerForms) {
		toSerialize["powerForms"] = o.PowerForms
	}
	if !IsNil(o.PreviousUri) {
		toSerialize["previousUri"] = o.PreviousUri
	}
	if !IsNil(o.ResultSetSize) {
		toSerialize["resultSetSize"] = o.ResultSetSize
	}
	if !IsNil(o.StartPosition) {
		toSerialize["startPosition"] = o.StartPosition
	}
	if !IsNil(o.TotalSetSize) {
		toSerialize["totalSetSize"] = o.TotalSetSize
	}
	return toSerialize, nil
}

type NullablePowerFormsResponse struct {
	value *PowerFormsResponse
	isSet bool
}

func (v NullablePowerFormsResponse) Get() *PowerFormsResponse {
	return v.value
}

func (v *NullablePowerFormsResponse) Set(val *PowerFormsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePowerFormsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePowerFormsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePowerFormsResponse(val *PowerFormsResponse) *NullablePowerFormsResponse {
	return &NullablePowerFormsResponse{value: val, isSet: true}
}

func (v NullablePowerFormsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePowerFormsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


