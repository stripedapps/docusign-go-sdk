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

// checks if the ConditionalRecipientRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionalRecipientRule{}

// ConditionalRecipientRule A rule that defines a set of recipients and the conditions under which they will be used for the envelope.
type ConditionalRecipientRule struct {
	// An array of conditions that define when the recipients will be used.
	Conditions []ConditionalRecipientRuleCondition `json:"conditions,omitempty"`
	// An integer that specifies the order in which rules are processed. Lower values are processed before higher values.
	Order *string `json:"order,omitempty"`
	RecipientGroup *RecipientGroup `json:"recipientGroup,omitempty"`
	// A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
	RecipientId *string `json:"recipientId,omitempty"`
}

// NewConditionalRecipientRule instantiates a new ConditionalRecipientRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionalRecipientRule() *ConditionalRecipientRule {
	this := ConditionalRecipientRule{}
	return &this
}

// NewConditionalRecipientRuleWithDefaults instantiates a new ConditionalRecipientRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionalRecipientRuleWithDefaults() *ConditionalRecipientRule {
	this := ConditionalRecipientRule{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ConditionalRecipientRule) GetConditions() []ConditionalRecipientRuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []ConditionalRecipientRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRule) GetConditionsOk() ([]ConditionalRecipientRuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ConditionalRecipientRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ConditionalRecipientRuleCondition and assigns it to the Conditions field.
func (o *ConditionalRecipientRule) SetConditions(v []ConditionalRecipientRuleCondition) {
	o.Conditions = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *ConditionalRecipientRule) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRule) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *ConditionalRecipientRule) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *ConditionalRecipientRule) SetOrder(v string) {
	o.Order = &v
}

// GetRecipientGroup returns the RecipientGroup field value if set, zero value otherwise.
func (o *ConditionalRecipientRule) GetRecipientGroup() RecipientGroup {
	if o == nil || IsNil(o.RecipientGroup) {
		var ret RecipientGroup
		return ret
	}
	return *o.RecipientGroup
}

// GetRecipientGroupOk returns a tuple with the RecipientGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRule) GetRecipientGroupOk() (*RecipientGroup, bool) {
	if o == nil || IsNil(o.RecipientGroup) {
		return nil, false
	}
	return o.RecipientGroup, true
}

// HasRecipientGroup returns a boolean if a field has been set.
func (o *ConditionalRecipientRule) HasRecipientGroup() bool {
	if o != nil && !IsNil(o.RecipientGroup) {
		return true
	}

	return false
}

// SetRecipientGroup gets a reference to the given RecipientGroup and assigns it to the RecipientGroup field.
func (o *ConditionalRecipientRule) SetRecipientGroup(v RecipientGroup) {
	o.RecipientGroup = &v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *ConditionalRecipientRule) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionalRecipientRule) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *ConditionalRecipientRule) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *ConditionalRecipientRule) SetRecipientId(v string) {
	o.RecipientId = &v
}

func (o ConditionalRecipientRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionalRecipientRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !IsNil(o.RecipientGroup) {
		toSerialize["recipientGroup"] = o.RecipientGroup
	}
	if !IsNil(o.RecipientId) {
		toSerialize["recipientId"] = o.RecipientId
	}
	return toSerialize, nil
}

type NullableConditionalRecipientRule struct {
	value *ConditionalRecipientRule
	isSet bool
}

func (v NullableConditionalRecipientRule) Get() *ConditionalRecipientRule {
	return v.value
}

func (v *NullableConditionalRecipientRule) Set(val *ConditionalRecipientRule) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionalRecipientRule) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionalRecipientRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionalRecipientRule(val *ConditionalRecipientRule) *NullableConditionalRecipientRule {
	return &NullableConditionalRecipientRule{value: val, isSet: true}
}

func (v NullableConditionalRecipientRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionalRecipientRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


