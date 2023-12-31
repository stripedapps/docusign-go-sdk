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

// checks if the ConditionalRecipientRuleFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionalRecipientRuleFilter{}

// ConditionalRecipientRuleFilter 
type ConditionalRecipientRuleFilter struct {
	// How the tab value is compared to the `value` property. Valid values:  * `equals` * `greaterThan` * `greaterThanEquals` * `lessThan` * `lessThanEquals` * `filled` * `selected`
	Operator *string `json:"operator,omitempty"`
	// A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
	RecipientId *string `json:"recipientId,omitempty"`
	// The scope under which the condition is evaluated. Valid values:  * `tabs`
	Scope *string `json:"scope,omitempty"`
	// The unique identifier for the tab.
	TabId *string `json:"tabId,omitempty"`
	// The label associated with the tab. This value may be an empty string. If no value is provided, the tab type is used as the value.  Maximum Length: 500 characters. 
	TabLabel *string `json:"tabLabel,omitempty"`
	// Indicates the type of tab (for example, `signHere` or `initialHere`).
	TabType *string `json:"tabType,omitempty"`
	// A set value to which the tab's actual value is compared.
	Value *string `json:"value,omitempty"`
}

// NewConditionalRecipientRuleFilter instantiates a new ConditionalRecipientRuleFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalRecipientRuleFilter() *ConditionalRecipientRuleFilter {
	this := ConditionalRecipientRuleFilter{}
	return &this
}

// NewConditionalRecipientRuleFilterWithDefaults instantiates a new ConditionalRecipientRuleFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalRecipientRuleFilterWithDefaults() *ConditionalRecipientRuleFilter {
	this := ConditionalRecipientRuleFilter{}
	return &this
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetOperator() string {
	if o == nil || IsNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetOperatorOk() (*string, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *ConditionalRecipientRuleFilter) SetOperator(v string) {
	o.Operator = &v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *ConditionalRecipientRuleFilter) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetScope() string {
	if o == nil || IsNil(o.Scope) {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetScopeOk() (*string, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *ConditionalRecipientRuleFilter) SetScope(v string) {
	o.Scope = &v
}

// GetTabId returns the TabId field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetTabId() string {
	if o == nil || IsNil(o.TabId) {
		var ret string
		return ret
	}
	return *o.TabId
}

// GetTabIdOk returns a tuple with the TabId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetTabIdOk() (*string, bool) {
	if o == nil || IsNil(o.TabId) {
		return nil, false
	}
	return o.TabId, true
}

// HasTabId returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasTabId() bool {
	if o != nil && !IsNil(o.TabId) {
		return true
	}

	return false
}

// SetTabId gets a reference to the given string and assigns it to the TabId field.
func (o *ConditionalRecipientRuleFilter) SetTabId(v string) {
	o.TabId = &v
}

// GetTabLabel returns the TabLabel field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetTabLabel() string {
	if o == nil || IsNil(o.TabLabel) {
		var ret string
		return ret
	}
	return *o.TabLabel
}

// GetTabLabelOk returns a tuple with the TabLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetTabLabelOk() (*string, bool) {
	if o == nil || IsNil(o.TabLabel) {
		return nil, false
	}
	return o.TabLabel, true
}

// HasTabLabel returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasTabLabel() bool {
	if o != nil && !IsNil(o.TabLabel) {
		return true
	}

	return false
}

// SetTabLabel gets a reference to the given string and assigns it to the TabLabel field.
func (o *ConditionalRecipientRuleFilter) SetTabLabel(v string) {
	o.TabLabel = &v
}

// GetTabType returns the TabType field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetTabType() string {
	if o == nil || IsNil(o.TabType) {
		var ret string
		return ret
	}
	return *o.TabType
}

// GetTabTypeOk returns a tuple with the TabType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetTabTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TabType) {
		return nil, false
	}
	return o.TabType, true
}

// HasTabType returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasTabType() bool {
	if o != nil && !IsNil(o.TabType) {
		return true
	}

	return false
}

// SetTabType gets a reference to the given string and assigns it to the TabType field.
func (o *ConditionalRecipientRuleFilter) SetTabType(v string) {
	o.TabType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ConditionalRecipientRuleFilter) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRuleFilter) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ConditionalRecipientRuleFilter) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ConditionalRecipientRuleFilter) SetValue(v string) {
	o.Value = &v
}

func (o ConditionalRecipientRuleFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalRecipientRuleFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.RecipientId) {
		toSerialize["recipientId"] = o.RecipientId
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.TabId) {
		toSerialize["tabId"] = o.TabId
	}
	if !IsNil(o.TabLabel) {
		toSerialize["tabLabel"] = o.TabLabel
	}
	if !IsNil(o.TabType) {
		toSerialize["tabType"] = o.TabType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableConditionalRecipientRuleFilter struct {
	value *ConditionalRecipientRuleFilter
	isSet bool
}

func (v NullableConditionalRecipientRuleFilter) Get() *ConditionalRecipientRuleFilter {
	return v.value
}

func (v *NullableConditionalRecipientRuleFilter) Set(val *ConditionalRecipientRuleFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalRecipientRuleFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalRecipientRuleFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalRecipientRuleFilter(val *ConditionalRecipientRuleFilter) *NullableConditionalRecipientRuleFilter {
	return &NullableConditionalRecipientRuleFilter{value: val, isSet: true}
}

func (v NullableConditionalRecipientRuleFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalRecipientRuleFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


