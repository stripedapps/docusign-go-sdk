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

// checks if the ConnectSalesforceObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectSalesforceObject{}

// ConnectSalesforceObject A `connectSalesforceObject` is an object that updates envelope and document status or recipient status in your Salesforce account.  When you install DocuSign Connect for Salesforce, the service automatically sets up two Connect objects: one that updates envelope status and documents and one that updates recipient status. You can also customize DocuSign Connect for Salesforce by associating DocuSign objects with Salesforce objects so that DocuSign Connect for Salesforce updates or inserts the information into the Salesforce object. For more information, see  [DocuSign for Salesforce - Adding Completed Documents to the Notes and Attachments](https://support.docusign.com/s/articles/DocuSign-for-Salesforce-Adding-Completed-Documents-to-the-Notes-and-Attachments-New).
type ConnectSalesforceObject struct {
	// When **true,** the `connectSalesforceObject` is active.
	Active *string `json:"active,omitempty"`
	// A description of the `connectSalesforceObject`.
	Description *string `json:"description,omitempty"`
	// The ID of the `connectSalesforceObject`.
	Id *string `json:"id,omitempty"`
	// 
	Insert *string `json:"insert,omitempty"`
	// When **true,** Salesforce is updated only when the envelope is complete.
	OnCompleteOnly *string `json:"onCompleteOnly,omitempty"`
	// The DocuSign and Salesforce fields that you want to use to match a Salesforce object with DocuSign information. This information tells Connect when to send updates to Salesforce.
	SelectFields []ConnectSalesforceField `json:"selectFields,omitempty"`
	// The Salesforce.com object type, such as `case`, `contact`, or `opportunity`.
	SfObject *string `json:"sfObject,omitempty"`
	// A name for the Salesforce object.  **Note:** You can enter any name for the object. It does not have to match the `sfObject` property.
	SfObjectName *string `json:"sfObjectName,omitempty"`
	// The DocuSign and Salesforce fields that you want to update.   **Note:** You can choose to update SalesForce (with information from DocuSign) only, update DocuSign only, or both.
	UpdateFields []ConnectSalesforceField `json:"updateFields,omitempty"`
}

// NewConnectSalesforceObject instantiates a new ConnectSalesforceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectSalesforceObject() *ConnectSalesforceObject {
	this := ConnectSalesforceObject{}
	return &this
}

// NewConnectSalesforceObjectWithDefaults instantiates a new ConnectSalesforceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectSalesforceObjectWithDefaults() *ConnectSalesforceObject {
	this := ConnectSalesforceObject{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetActive() string {
	if o == nil || IsNil(o.Active) {
		var ret string
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetActiveOk() (*string, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given string and assigns it to the Active field.
func (o *ConnectSalesforceObject) SetActive(v string) {
	o.Active = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectSalesforceObject) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectSalesforceObject) SetId(v string) {
	o.Id = &v
}

// GetInsert returns the Insert field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetInsert() string {
	if o == nil || IsNil(o.Insert) {
		var ret string
		return ret
	}
	return *o.Insert
}

// GetInsertOk returns a tuple with the Insert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetInsertOk() (*string, bool) {
	if o == nil || IsNil(o.Insert) {
		return nil, false
	}
	return o.Insert, true
}

// HasInsert returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasInsert() bool {
	if o != nil && !IsNil(o.Insert) {
		return true
	}

	return false
}

// SetInsert gets a reference to the given string and assigns it to the Insert field.
func (o *ConnectSalesforceObject) SetInsert(v string) {
	o.Insert = &v
}

// GetOnCompleteOnly returns the OnCompleteOnly field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetOnCompleteOnly() string {
	if o == nil || IsNil(o.OnCompleteOnly) {
		var ret string
		return ret
	}
	return *o.OnCompleteOnly
}

// GetOnCompleteOnlyOk returns a tuple with the OnCompleteOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetOnCompleteOnlyOk() (*string, bool) {
	if o == nil || IsNil(o.OnCompleteOnly) {
		return nil, false
	}
	return o.OnCompleteOnly, true
}

// HasOnCompleteOnly returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasOnCompleteOnly() bool {
	if o != nil && !IsNil(o.OnCompleteOnly) {
		return true
	}

	return false
}

// SetOnCompleteOnly gets a reference to the given string and assigns it to the OnCompleteOnly field.
func (o *ConnectSalesforceObject) SetOnCompleteOnly(v string) {
	o.OnCompleteOnly = &v
}

// GetSelectFields returns the SelectFields field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetSelectFields() []ConnectSalesforceField {
	if o == nil || IsNil(o.SelectFields) {
		var ret []ConnectSalesforceField
		return ret
	}
	return o.SelectFields
}

// GetSelectFieldsOk returns a tuple with the SelectFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetSelectFieldsOk() ([]ConnectSalesforceField, bool) {
	if o == nil || IsNil(o.SelectFields) {
		return nil, false
	}
	return o.SelectFields, true
}

// HasSelectFields returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasSelectFields() bool {
	if o != nil && !IsNil(o.SelectFields) {
		return true
	}

	return false
}

// SetSelectFields gets a reference to the given []ConnectSalesforceField and assigns it to the SelectFields field.
func (o *ConnectSalesforceObject) SetSelectFields(v []ConnectSalesforceField) {
	o.SelectFields = v
}

// GetSfObject returns the SfObject field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetSfObject() string {
	if o == nil || IsNil(o.SfObject) {
		var ret string
		return ret
	}
	return *o.SfObject
}

// GetSfObjectOk returns a tuple with the SfObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetSfObjectOk() (*string, bool) {
	if o == nil || IsNil(o.SfObject) {
		return nil, false
	}
	return o.SfObject, true
}

// HasSfObject returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasSfObject() bool {
	if o != nil && !IsNil(o.SfObject) {
		return true
	}

	return false
}

// SetSfObject gets a reference to the given string and assigns it to the SfObject field.
func (o *ConnectSalesforceObject) SetSfObject(v string) {
	o.SfObject = &v
}

// GetSfObjectName returns the SfObjectName field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetSfObjectName() string {
	if o == nil || IsNil(o.SfObjectName) {
		var ret string
		return ret
	}
	return *o.SfObjectName
}

// GetSfObjectNameOk returns a tuple with the SfObjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetSfObjectNameOk() (*string, bool) {
	if o == nil || IsNil(o.SfObjectName) {
		return nil, false
	}
	return o.SfObjectName, true
}

// HasSfObjectName returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasSfObjectName() bool {
	if o != nil && !IsNil(o.SfObjectName) {
		return true
	}

	return false
}

// SetSfObjectName gets a reference to the given string and assigns it to the SfObjectName field.
func (o *ConnectSalesforceObject) SetSfObjectName(v string) {
	o.SfObjectName = &v
}

// GetUpdateFields returns the UpdateFields field value if set, zero value otherwise.
func (o *ConnectSalesforceObject) GetUpdateFields() []ConnectSalesforceField {
	if o == nil || IsNil(o.UpdateFields) {
		var ret []ConnectSalesforceField
		return ret
	}
	return o.UpdateFields
}

// GetUpdateFieldsOk returns a tuple with the UpdateFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectSalesforceObject) GetUpdateFieldsOk() ([]ConnectSalesforceField, bool) {
	if o == nil || IsNil(o.UpdateFields) {
		return nil, false
	}
	return o.UpdateFields, true
}

// HasUpdateFields returns a boolean if a field has been set.
func (o *ConnectSalesforceObject) HasUpdateFields() bool {
	if o != nil && !IsNil(o.UpdateFields) {
		return true
	}

	return false
}

// SetUpdateFields gets a reference to the given []ConnectSalesforceField and assigns it to the UpdateFields field.
func (o *ConnectSalesforceObject) SetUpdateFields(v []ConnectSalesforceField) {
	o.UpdateFields = v
}

func (o ConnectSalesforceObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectSalesforceObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Insert) {
		toSerialize["insert"] = o.Insert
	}
	if !IsNil(o.OnCompleteOnly) {
		toSerialize["onCompleteOnly"] = o.OnCompleteOnly
	}
	if !IsNil(o.SelectFields) {
		toSerialize["selectFields"] = o.SelectFields
	}
	if !IsNil(o.SfObject) {
		toSerialize["sfObject"] = o.SfObject
	}
	if !IsNil(o.SfObjectName) {
		toSerialize["sfObjectName"] = o.SfObjectName
	}
	if !IsNil(o.UpdateFields) {
		toSerialize["updateFields"] = o.UpdateFields
	}
	return toSerialize, nil
}

type NullableConnectSalesforceObject struct {
	value *ConnectSalesforceObject
	isSet bool
}

func (v NullableConnectSalesforceObject) Get() *ConnectSalesforceObject {
	return v.value
}

func (v *NullableConnectSalesforceObject) Set(val *ConnectSalesforceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectSalesforceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectSalesforceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectSalesforceObject(val *ConnectSalesforceObject) *NullableConnectSalesforceObject {
	return &NullableConnectSalesforceObject{value: val, isSet: true}
}

func (v NullableConnectSalesforceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectSalesforceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


