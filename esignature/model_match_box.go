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

// checks if the MatchBox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatchBox{}

// MatchBox 
type MatchBox struct {
	// The height of the tab in pixels. Must be an integer.
	Height *string `json:"height,omitempty"`
	// Specifies the page number on which the tab is located. Must be 1 for supplemental documents. 
	PageNumber *string `json:"pageNumber,omitempty"`
	// The width of the tab in pixels. Must be an integer.
	Width *string `json:"width,omitempty"`
	// This property indicates the horizontal offset of the object on the page. DocuSign uses 72 DPI when determining position. Required. Must be an integer. May be zero. 
	XPosition *string `json:"xPosition,omitempty"`
	// This property indicates the vertical offset of the object on the page. DocuSign uses 72 DPI when determining position. Required. Must be an integer. May be zero. 
	YPosition *string `json:"yPosition,omitempty"`
}

// NewMatchBox instantiates a new MatchBox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatchBox() *MatchBox {
	this := MatchBox{}
	return &this
}

// NewMatchBoxWithDefaults instantiates a new MatchBox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchBoxWithDefaults() *MatchBox {
	this := MatchBox{}
	return &this
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *MatchBox) GetHeight() string {
	if o == nil || IsNil(o.Height) {
		var ret string
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchBox) GetHeightOk() (*string, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *MatchBox) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given string and assigns it to the Height field.
func (o *MatchBox) SetHeight(v string) {
	o.Height = &v
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *MatchBox) GetPageNumber() string {
	if o == nil || IsNil(o.PageNumber) {
		var ret string
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchBox) GetPageNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *MatchBox) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given string and assigns it to the PageNumber field.
func (o *MatchBox) SetPageNumber(v string) {
	o.PageNumber = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *MatchBox) GetWidth() string {
	if o == nil || IsNil(o.Width) {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchBox) GetWidthOk() (*string, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *MatchBox) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *MatchBox) SetWidth(v string) {
	o.Width = &v
}

// GetXPosition returns the XPosition field value if set, zero value otherwise.
func (o *MatchBox) GetXPosition() string {
	if o == nil || IsNil(o.XPosition) {
		var ret string
		return ret
	}
	return *o.XPosition
}

// GetXPositionOk returns a tuple with the XPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchBox) GetXPositionOk() (*string, bool) {
	if o == nil || IsNil(o.XPosition) {
		return nil, false
	}
	return o.XPosition, true
}

// HasXPosition returns a boolean if a field has been set.
func (o *MatchBox) HasXPosition() bool {
	if o != nil && !IsNil(o.XPosition) {
		return true
	}

	return false
}

// SetXPosition gets a reference to the given string and assigns it to the XPosition field.
func (o *MatchBox) SetXPosition(v string) {
	o.XPosition = &v
}

// GetYPosition returns the YPosition field value if set, zero value otherwise.
func (o *MatchBox) GetYPosition() string {
	if o == nil || IsNil(o.YPosition) {
		var ret string
		return ret
	}
	return *o.YPosition
}

// GetYPositionOk returns a tuple with the YPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchBox) GetYPositionOk() (*string, bool) {
	if o == nil || IsNil(o.YPosition) {
		return nil, false
	}
	return o.YPosition, true
}

// HasYPosition returns a boolean if a field has been set.
func (o *MatchBox) HasYPosition() bool {
	if o != nil && !IsNil(o.YPosition) {
		return true
	}

	return false
}

// SetYPosition gets a reference to the given string and assigns it to the YPosition field.
func (o *MatchBox) SetYPosition(v string) {
	o.YPosition = &v
}

func (o MatchBox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchBox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.XPosition) {
		toSerialize["xPosition"] = o.XPosition
	}
	if !IsNil(o.YPosition) {
		toSerialize["yPosition"] = o.YPosition
	}
	return toSerialize, nil
}

type NullableMatchBox struct {
	value *MatchBox
	isSet bool
}

func (v NullableMatchBox) Get() *MatchBox {
	return v.value
}

func (v *NullableMatchBox) Set(val *MatchBox) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchBox) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchBox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchBox(val *MatchBox) *NullableMatchBox {
	return &NullableMatchBox{value: val, isSet: true}
}

func (v NullableMatchBox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchBox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

