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

// checks if the ListCustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListCustomField{}

// ListCustomField This object represents a list custom field from which envelope creators and senders can select custom data.
type ListCustomField struct {
	// If you are using merge fields, this property specifies the type of the merge field. The only supported value is `salesforce`.
	ConfigurationType *string `json:"configurationType,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// The ID of the custom field.
	FieldId *string `json:"fieldId,omitempty"`
	// An array of strings that represents the options in a list.  Maximum length: 2048 characters, but each individual option string can only be a maximum of 100 characters.
	ListItems []string `json:"listItems,omitempty"`
	// The name of the custom field.
	Name *string `json:"name,omitempty"`
	// When **true,** senders are required to select an option from the list before they can send the envelope.
	Required *string `json:"required,omitempty"`
	// When **true,** the field displays in the **Envelope Custom Fields** section when a user creates or sends an envelope.
	Show *string `json:"show,omitempty"`
	// The value of the custom field. This is the value that the user who creates or sends the envelope selects from the list.
	Value *string `json:"value,omitempty"`
}

// NewListCustomField instantiates a new ListCustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListCustomField() *ListCustomField {
	this := ListCustomField{}
	return &this
}

// NewListCustomFieldWithDefaults instantiates a new ListCustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListCustomFieldWithDefaults() *ListCustomField {
	this := ListCustomField{}
	return &this
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *ListCustomField) GetConfigurationType() string {
	if o == nil || IsNil(o.ConfigurationType) {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationType) {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *ListCustomField) HasConfigurationType() bool {
	if o != nil && !IsNil(o.ConfigurationType) {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *ListCustomField) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *ListCustomField) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *ListCustomField) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *ListCustomField) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *ListCustomField) GetFieldId() string {
	if o == nil || IsNil(o.FieldId) {
		var ret string
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetFieldIdOk() (*string, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *ListCustomField) HasFieldId() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given string and assigns it to the FieldId field.
func (o *ListCustomField) SetFieldId(v string) {
	o.FieldId = &v
}

// GetListItems returns the ListItems field value if set, zero value otherwise.
func (o *ListCustomField) GetListItems() []string {
	if o == nil || IsNil(o.ListItems) {
		var ret []string
		return ret
	}
	return o.ListItems
}

// GetListItemsOk returns a tuple with the ListItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetListItemsOk() ([]string, bool) {
	if o == nil || IsNil(o.ListItems) {
		return nil, false
	}
	return o.ListItems, true
}

// HasListItems returns a boolean if a field has been set.
func (o *ListCustomField) HasListItems() bool {
	if o != nil && !IsNil(o.ListItems) {
		return true
	}

	return false
}

// SetListItems gets a reference to the given []string and assigns it to the ListItems field.
func (o *ListCustomField) SetListItems(v []string) {
	o.ListItems = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListCustomField) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListCustomField) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListCustomField) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *ListCustomField) GetRequired() string {
	if o == nil || IsNil(o.Required) {
		var ret string
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetRequiredOk() (*string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *ListCustomField) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given string and assigns it to the Required field.
func (o *ListCustomField) SetRequired(v string) {
	o.Required = &v
}

// GetShow returns the Show field value if set, zero value otherwise.
func (o *ListCustomField) GetShow() string {
	if o == nil || IsNil(o.Show) {
		var ret string
		return ret
	}
	return *o.Show
}

// GetShowOk returns a tuple with the Show field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetShowOk() (*string, bool) {
	if o == nil || IsNil(o.Show) {
		return nil, false
	}
	return o.Show, true
}

// HasShow returns a boolean if a field has been set.
func (o *ListCustomField) HasShow() bool {
	if o != nil && !IsNil(o.Show) {
		return true
	}

	return false
}

// SetShow gets a reference to the given string and assigns it to the Show field.
func (o *ListCustomField) SetShow(v string) {
	o.Show = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListCustomField) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListCustomField) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListCustomField) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListCustomField) SetValue(v string) {
	o.Value = &v
}

func (o ListCustomField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListCustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConfigurationType) {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.FieldId) {
		toSerialize["fieldId"] = o.FieldId
	}
	if !IsNil(o.ListItems) {
		toSerialize["listItems"] = o.ListItems
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Show) {
		toSerialize["show"] = o.Show
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableListCustomField struct {
	value *ListCustomField
	isSet bool
}

func (v NullableListCustomField) Get() *ListCustomField {
	return v.value
}

func (v *NullableListCustomField) Set(val *ListCustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableListCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableListCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListCustomField(val *ListCustomField) *NullableListCustomField {
	return &NullableListCustomField{value: val, isSet: true}
}

func (v NullableListCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


