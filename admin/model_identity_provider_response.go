/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the IdentityProviderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderResponse{}

// IdentityProviderResponse Information about a single identity provider.
type IdentityProviderResponse struct {
	// The unique ID of the identity provider.
	Id *string `json:"id,omitempty"`
	// The human-readable name of the identity provider.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// When **true,** users who use this identity provider are automatically provisioned. 
	AutoProvisionUsers *bool `json:"auto_provision_users,omitempty"`
	// The type of the identity provider. One of:  - `none` - `saml_20` - `saml_11` - `saml_10` - `ws_federation` - `open_id_connect`
	Type *string `json:"type,omitempty"`
	Saml20 *Saml2IdentityProviderResponse `json:"saml_20,omitempty"`
	// A list of useful URLs.
	Links []LinkResponse `json:"links,omitempty"`
}

// NewIdentityProviderResponse instantiates a new IdentityProviderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderResponse() *IdentityProviderResponse {
	this := IdentityProviderResponse{}
	return &this
}

// NewIdentityProviderResponseWithDefaults instantiates a new IdentityProviderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderResponseWithDefaults() *IdentityProviderResponse {
	this := IdentityProviderResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderResponse) SetId(v string) {
	o.Id = &v
}

// GetFriendlyName returns the FriendlyName field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetFriendlyName() string {
	if o == nil || IsNil(o.FriendlyName) {
		var ret string
		return ret
	}
	return *o.FriendlyName
}

// GetFriendlyNameOk returns a tuple with the FriendlyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetFriendlyNameOk() (*string, bool) {
	if o == nil || IsNil(o.FriendlyName) {
		return nil, false
	}
	return o.FriendlyName, true
}

// HasFriendlyName returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasFriendlyName() bool {
	if o != nil && !IsNil(o.FriendlyName) {
		return true
	}

	return false
}

// SetFriendlyName gets a reference to the given string and assigns it to the FriendlyName field.
func (o *IdentityProviderResponse) SetFriendlyName(v string) {
	o.FriendlyName = &v
}

// GetAutoProvisionUsers returns the AutoProvisionUsers field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetAutoProvisionUsers() bool {
	if o == nil || IsNil(o.AutoProvisionUsers) {
		var ret bool
		return ret
	}
	return *o.AutoProvisionUsers
}

// GetAutoProvisionUsersOk returns a tuple with the AutoProvisionUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetAutoProvisionUsersOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoProvisionUsers) {
		return nil, false
	}
	return o.AutoProvisionUsers, true
}

// HasAutoProvisionUsers returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasAutoProvisionUsers() bool {
	if o != nil && !IsNil(o.AutoProvisionUsers) {
		return true
	}

	return false
}

// SetAutoProvisionUsers gets a reference to the given bool and assigns it to the AutoProvisionUsers field.
func (o *IdentityProviderResponse) SetAutoProvisionUsers(v bool) {
	o.AutoProvisionUsers = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IdentityProviderResponse) SetType(v string) {
	o.Type = &v
}

// GetSaml20 returns the Saml20 field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetSaml20() Saml2IdentityProviderResponse {
	if o == nil || IsNil(o.Saml20) {
		var ret Saml2IdentityProviderResponse
		return ret
	}
	return *o.Saml20
}

// GetSaml20Ok returns a tuple with the Saml20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetSaml20Ok() (*Saml2IdentityProviderResponse, bool) {
	if o == nil || IsNil(o.Saml20) {
		return nil, false
	}
	return o.Saml20, true
}

// HasSaml20 returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasSaml20() bool {
	if o != nil && !IsNil(o.Saml20) {
		return true
	}

	return false
}

// SetSaml20 gets a reference to the given Saml2IdentityProviderResponse and assigns it to the Saml20 field.
func (o *IdentityProviderResponse) SetSaml20(v Saml2IdentityProviderResponse) {
	o.Saml20 = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityProviderResponse) GetLinks() []LinkResponse {
	if o == nil || IsNil(o.Links) {
		var ret []LinkResponse
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderResponse) GetLinksOk() ([]LinkResponse, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityProviderResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []LinkResponse and assigns it to the Links field.
func (o *IdentityProviderResponse) SetLinks(v []LinkResponse) {
	o.Links = v
}

func (o IdentityProviderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FriendlyName) {
		toSerialize["friendly_name"] = o.FriendlyName
	}
	if !IsNil(o.AutoProvisionUsers) {
		toSerialize["auto_provision_users"] = o.AutoProvisionUsers
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Saml20) {
		toSerialize["saml_20"] = o.Saml20
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableIdentityProviderResponse struct {
	value *IdentityProviderResponse
	isSet bool
}

func (v NullableIdentityProviderResponse) Get() *IdentityProviderResponse {
	return v.value
}

func (v *NullableIdentityProviderResponse) Set(val *IdentityProviderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderResponse(val *IdentityProviderResponse) *NullableIdentityProviderResponse {
	return &NullableIdentityProviderResponse{value: val, isSet: true}
}

func (v NullableIdentityProviderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


