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

// checks if the MergeField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MergeField{}

// MergeField Contains information for transferring values between Salesforce data fields and DocuSign tabs. 
type MergeField struct {
	// When **true,** the sender can modify the value of the `mergeField` tab during the sending process.
	AllowSenderToEdit *string `json:"allowSenderToEdit,omitempty"`
	AllowSenderToEditMetadata *PropertyMetadata `json:"allowSenderToEditMetadata,omitempty"`
	// If you are using merge fields, this property specifies the type of the merge field. The only supported value is `salesforce`.
	ConfigurationType *string `json:"configurationType,omitempty"`
	ConfigurationTypeMetadata *PropertyMetadata `json:"configurationTypeMetadata,omitempty"`
	// Sets the object associated with the custom tab. Currently this is the Salesforce Object.
	Path *string `json:"path,omitempty"`
	// Reserved for DocuSign.
	PathExtended []PathExtendedElement `json:"pathExtended,omitempty"`
	PathExtendedMetadata *PropertyMetadata `json:"pathExtendedMetadata,omitempty"`
	PathMetadata *PropertyMetadata `json:"pathMetadata,omitempty"`
	// Specifies the row number in a Salesforce table that the merge field value corresponds to.
	Row *string `json:"row,omitempty"`
	RowMetadata *PropertyMetadata `json:"rowMetadata,omitempty"`
	// When **true,** data entered into the merge field during Signing will update the mapped Salesforce field.
	WriteBack *string `json:"writeBack,omitempty"`
	WriteBackMetadata *PropertyMetadata `json:"writeBackMetadata,omitempty"`
}

// NewMergeField instantiates a new MergeField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeField() *MergeField {
	this := MergeField{}
	return &this
}

// NewMergeFieldWithDefaults instantiates a new MergeField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeFieldWithDefaults() *MergeField {
	this := MergeField{}
	return &this
}

// GetAllowSenderToEdit returns the AllowSenderToEdit field value if set, zero value otherwise.
func (o *MergeField) GetAllowSenderToEdit() string {
	if o == nil || IsNil(o.AllowSenderToEdit) {
		var ret string
		return ret
	}
	return *o.AllowSenderToEdit
}

// GetAllowSenderToEditOk returns a tuple with the AllowSenderToEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetAllowSenderToEditOk() (*string, bool) {
	if o == nil || IsNil(o.AllowSenderToEdit) {
		return nil, false
	}
	return o.AllowSenderToEdit, true
}

// HasAllowSenderToEdit returns a boolean if a field has been set.
func (o *MergeField) HasAllowSenderToEdit() bool {
	if o != nil && !IsNil(o.AllowSenderToEdit) {
		return true
	}

	return false
}

// SetAllowSenderToEdit gets a reference to the given string and assigns it to the AllowSenderToEdit field.
func (o *MergeField) SetAllowSenderToEdit(v string) {
	o.AllowSenderToEdit = &v
}

// GetAllowSenderToEditMetadata returns the AllowSenderToEditMetadata field value if set, zero value otherwise.
func (o *MergeField) GetAllowSenderToEditMetadata() PropertyMetadata {
	if o == nil || IsNil(o.AllowSenderToEditMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.AllowSenderToEditMetadata
}

// GetAllowSenderToEditMetadataOk returns a tuple with the AllowSenderToEditMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetAllowSenderToEditMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.AllowSenderToEditMetadata) {
		return nil, false
	}
	return o.AllowSenderToEditMetadata, true
}

// HasAllowSenderToEditMetadata returns a boolean if a field has been set.
func (o *MergeField) HasAllowSenderToEditMetadata() bool {
	if o != nil && !IsNil(o.AllowSenderToEditMetadata) {
		return true
	}

	return false
}

// SetAllowSenderToEditMetadata gets a reference to the given PropertyMetadata and assigns it to the AllowSenderToEditMetadata field.
func (o *MergeField) SetAllowSenderToEditMetadata(v PropertyMetadata) {
	o.AllowSenderToEditMetadata = &v
}

// GetConfigurationType returns the ConfigurationType field value if set, zero value otherwise.
func (o *MergeField) GetConfigurationType() string {
	if o == nil || IsNil(o.ConfigurationType) {
		var ret string
		return ret
	}
	return *o.ConfigurationType
}

// GetConfigurationTypeOk returns a tuple with the ConfigurationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetConfigurationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigurationType) {
		return nil, false
	}
	return o.ConfigurationType, true
}

// HasConfigurationType returns a boolean if a field has been set.
func (o *MergeField) HasConfigurationType() bool {
	if o != nil && !IsNil(o.ConfigurationType) {
		return true
	}

	return false
}

// SetConfigurationType gets a reference to the given string and assigns it to the ConfigurationType field.
func (o *MergeField) SetConfigurationType(v string) {
	o.ConfigurationType = &v
}

// GetConfigurationTypeMetadata returns the ConfigurationTypeMetadata field value if set, zero value otherwise.
func (o *MergeField) GetConfigurationTypeMetadata() PropertyMetadata {
	if o == nil || IsNil(o.ConfigurationTypeMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.ConfigurationTypeMetadata
}

// GetConfigurationTypeMetadataOk returns a tuple with the ConfigurationTypeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetConfigurationTypeMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.ConfigurationTypeMetadata) {
		return nil, false
	}
	return o.ConfigurationTypeMetadata, true
}

// HasConfigurationTypeMetadata returns a boolean if a field has been set.
func (o *MergeField) HasConfigurationTypeMetadata() bool {
	if o != nil && !IsNil(o.ConfigurationTypeMetadata) {
		return true
	}

	return false
}

// SetConfigurationTypeMetadata gets a reference to the given PropertyMetadata and assigns it to the ConfigurationTypeMetadata field.
func (o *MergeField) SetConfigurationTypeMetadata(v PropertyMetadata) {
	o.ConfigurationTypeMetadata = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *MergeField) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *MergeField) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *MergeField) SetPath(v string) {
	o.Path = &v
}

// GetPathExtended returns the PathExtended field value if set, zero value otherwise.
func (o *MergeField) GetPathExtended() []PathExtendedElement {
	if o == nil || IsNil(o.PathExtended) {
		var ret []PathExtendedElement
		return ret
	}
	return o.PathExtended
}

// GetPathExtendedOk returns a tuple with the PathExtended field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetPathExtendedOk() ([]PathExtendedElement, bool) {
	if o == nil || IsNil(o.PathExtended) {
		return nil, false
	}
	return o.PathExtended, true
}

// HasPathExtended returns a boolean if a field has been set.
func (o *MergeField) HasPathExtended() bool {
	if o != nil && !IsNil(o.PathExtended) {
		return true
	}

	return false
}

// SetPathExtended gets a reference to the given []PathExtendedElement and assigns it to the PathExtended field.
func (o *MergeField) SetPathExtended(v []PathExtendedElement) {
	o.PathExtended = v
}

// GetPathExtendedMetadata returns the PathExtendedMetadata field value if set, zero value otherwise.
func (o *MergeField) GetPathExtendedMetadata() PropertyMetadata {
	if o == nil || IsNil(o.PathExtendedMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.PathExtendedMetadata
}

// GetPathExtendedMetadataOk returns a tuple with the PathExtendedMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetPathExtendedMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.PathExtendedMetadata) {
		return nil, false
	}
	return o.PathExtendedMetadata, true
}

// HasPathExtendedMetadata returns a boolean if a field has been set.
func (o *MergeField) HasPathExtendedMetadata() bool {
	if o != nil && !IsNil(o.PathExtendedMetadata) {
		return true
	}

	return false
}

// SetPathExtendedMetadata gets a reference to the given PropertyMetadata and assigns it to the PathExtendedMetadata field.
func (o *MergeField) SetPathExtendedMetadata(v PropertyMetadata) {
	o.PathExtendedMetadata = &v
}

// GetPathMetadata returns the PathMetadata field value if set, zero value otherwise.
func (o *MergeField) GetPathMetadata() PropertyMetadata {
	if o == nil || IsNil(o.PathMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.PathMetadata
}

// GetPathMetadataOk returns a tuple with the PathMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetPathMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.PathMetadata) {
		return nil, false
	}
	return o.PathMetadata, true
}

// HasPathMetadata returns a boolean if a field has been set.
func (o *MergeField) HasPathMetadata() bool {
	if o != nil && !IsNil(o.PathMetadata) {
		return true
	}

	return false
}

// SetPathMetadata gets a reference to the given PropertyMetadata and assigns it to the PathMetadata field.
func (o *MergeField) SetPathMetadata(v PropertyMetadata) {
	o.PathMetadata = &v
}

// GetRow returns the Row field value if set, zero value otherwise.
func (o *MergeField) GetRow() string {
	if o == nil || IsNil(o.Row) {
		var ret string
		return ret
	}
	return *o.Row
}

// GetRowOk returns a tuple with the Row field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetRowOk() (*string, bool) {
	if o == nil || IsNil(o.Row) {
		return nil, false
	}
	return o.Row, true
}

// HasRow returns a boolean if a field has been set.
func (o *MergeField) HasRow() bool {
	if o != nil && !IsNil(o.Row) {
		return true
	}

	return false
}

// SetRow gets a reference to the given string and assigns it to the Row field.
func (o *MergeField) SetRow(v string) {
	o.Row = &v
}

// GetRowMetadata returns the RowMetadata field value if set, zero value otherwise.
func (o *MergeField) GetRowMetadata() PropertyMetadata {
	if o == nil || IsNil(o.RowMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.RowMetadata
}

// GetRowMetadataOk returns a tuple with the RowMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetRowMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.RowMetadata) {
		return nil, false
	}
	return o.RowMetadata, true
}

// HasRowMetadata returns a boolean if a field has been set.
func (o *MergeField) HasRowMetadata() bool {
	if o != nil && !IsNil(o.RowMetadata) {
		return true
	}

	return false
}

// SetRowMetadata gets a reference to the given PropertyMetadata and assigns it to the RowMetadata field.
func (o *MergeField) SetRowMetadata(v PropertyMetadata) {
	o.RowMetadata = &v
}

// GetWriteBack returns the WriteBack field value if set, zero value otherwise.
func (o *MergeField) GetWriteBack() string {
	if o == nil || IsNil(o.WriteBack) {
		var ret string
		return ret
	}
	return *o.WriteBack
}

// GetWriteBackOk returns a tuple with the WriteBack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetWriteBackOk() (*string, bool) {
	if o == nil || IsNil(o.WriteBack) {
		return nil, false
	}
	return o.WriteBack, true
}

// HasWriteBack returns a boolean if a field has been set.
func (o *MergeField) HasWriteBack() bool {
	if o != nil && !IsNil(o.WriteBack) {
		return true
	}

	return false
}

// SetWriteBack gets a reference to the given string and assigns it to the WriteBack field.
func (o *MergeField) SetWriteBack(v string) {
	o.WriteBack = &v
}

// GetWriteBackMetadata returns the WriteBackMetadata field value if set, zero value otherwise.
func (o *MergeField) GetWriteBackMetadata() PropertyMetadata {
	if o == nil || IsNil(o.WriteBackMetadata) {
		var ret PropertyMetadata
		return ret
	}
	return *o.WriteBackMetadata
}

// GetWriteBackMetadataOk returns a tuple with the WriteBackMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeField) GetWriteBackMetadataOk() (*PropertyMetadata, bool) {
	if o == nil || IsNil(o.WriteBackMetadata) {
		return nil, false
	}
	return o.WriteBackMetadata, true
}

// HasWriteBackMetadata returns a boolean if a field has been set.
func (o *MergeField) HasWriteBackMetadata() bool {
	if o != nil && !IsNil(o.WriteBackMetadata) {
		return true
	}

	return false
}

// SetWriteBackMetadata gets a reference to the given PropertyMetadata and assigns it to the WriteBackMetadata field.
func (o *MergeField) SetWriteBackMetadata(v PropertyMetadata) {
	o.WriteBackMetadata = &v
}

func (o MergeField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MergeField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowSenderToEdit) {
		toSerialize["allowSenderToEdit"] = o.AllowSenderToEdit
	}
	if !IsNil(o.AllowSenderToEditMetadata) {
		toSerialize["allowSenderToEditMetadata"] = o.AllowSenderToEditMetadata
	}
	if !IsNil(o.ConfigurationType) {
		toSerialize["configurationType"] = o.ConfigurationType
	}
	if !IsNil(o.ConfigurationTypeMetadata) {
		toSerialize["configurationTypeMetadata"] = o.ConfigurationTypeMetadata
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PathExtended) {
		toSerialize["pathExtended"] = o.PathExtended
	}
	if !IsNil(o.PathExtendedMetadata) {
		toSerialize["pathExtendedMetadata"] = o.PathExtendedMetadata
	}
	if !IsNil(o.PathMetadata) {
		toSerialize["pathMetadata"] = o.PathMetadata
	}
	if !IsNil(o.Row) {
		toSerialize["row"] = o.Row
	}
	if !IsNil(o.RowMetadata) {
		toSerialize["rowMetadata"] = o.RowMetadata
	}
	if !IsNil(o.WriteBack) {
		toSerialize["writeBack"] = o.WriteBack
	}
	if !IsNil(o.WriteBackMetadata) {
		toSerialize["writeBackMetadata"] = o.WriteBackMetadata
	}
	return toSerialize, nil
}

type NullableMergeField struct {
	value *MergeField
	isSet bool
}

func (v NullableMergeField) Get() *MergeField {
	return v.value
}

func (v *NullableMergeField) Set(val *MergeField) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeField) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeField(val *MergeField) *NullableMergeField {
	return &NullableMergeField{value: val, isSet: true}
}

func (v NullableMergeField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


