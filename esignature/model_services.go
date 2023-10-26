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

// checks if the Services type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Services{}

// Services API service information
type Services struct {
	// Reserved for DocuSign.
	BuildBranch *string `json:"buildBranch,omitempty"`
	// Reserved for DocuSign.
	BuildBranchDeployedDateTime *string `json:"buildBranchDeployedDateTime,omitempty"`
	// Reserved for DocuSign.
	BuildSHA *string `json:"buildSHA,omitempty"`
	// Reserved for DocuSign.
	BuildVersion *string `json:"buildVersion,omitempty"`
	// 
	LinkedSites []string `json:"linkedSites,omitempty"`
	// 
	ServiceVersions []ServiceVersion `json:"serviceVersions,omitempty"`
}

// NewServices instantiates a new Services object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServices() *Services {
	this := Services{}
	return &this
}

// NewServicesWithDefaults instantiates a new Services object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesWithDefaults() *Services {
	this := Services{}
	return &this
}

// GetBuildBranch returns the BuildBranch field value if set, zero value otherwise.
func (o *Services) GetBuildBranch() string {
	if o == nil || IsNil(o.BuildBranch) {
		var ret string
		return ret
	}
	return *o.BuildBranch
}

// GetBuildBranchOk returns a tuple with the BuildBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetBuildBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranch) {
		return nil, false
	}
	return o.BuildBranch, true
}

// HasBuildBranch returns a boolean if a field has been set.
func (o *Services) HasBuildBranch() bool {
	if o != nil && !IsNil(o.BuildBranch) {
		return true
	}

	return false
}

// SetBuildBranch gets a reference to the given string and assigns it to the BuildBranch field.
func (o *Services) SetBuildBranch(v string) {
	o.BuildBranch = &v
}

// GetBuildBranchDeployedDateTime returns the BuildBranchDeployedDateTime field value if set, zero value otherwise.
func (o *Services) GetBuildBranchDeployedDateTime() string {
	if o == nil || IsNil(o.BuildBranchDeployedDateTime) {
		var ret string
		return ret
	}
	return *o.BuildBranchDeployedDateTime
}

// GetBuildBranchDeployedDateTimeOk returns a tuple with the BuildBranchDeployedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetBuildBranchDeployedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranchDeployedDateTime) {
		return nil, false
	}
	return o.BuildBranchDeployedDateTime, true
}

// HasBuildBranchDeployedDateTime returns a boolean if a field has been set.
func (o *Services) HasBuildBranchDeployedDateTime() bool {
	if o != nil && !IsNil(o.BuildBranchDeployedDateTime) {
		return true
	}

	return false
}

// SetBuildBranchDeployedDateTime gets a reference to the given string and assigns it to the BuildBranchDeployedDateTime field.
func (o *Services) SetBuildBranchDeployedDateTime(v string) {
	o.BuildBranchDeployedDateTime = &v
}

// GetBuildSHA returns the BuildSHA field value if set, zero value otherwise.
func (o *Services) GetBuildSHA() string {
	if o == nil || IsNil(o.BuildSHA) {
		var ret string
		return ret
	}
	return *o.BuildSHA
}

// GetBuildSHAOk returns a tuple with the BuildSHA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetBuildSHAOk() (*string, bool) {
	if o == nil || IsNil(o.BuildSHA) {
		return nil, false
	}
	return o.BuildSHA, true
}

// HasBuildSHA returns a boolean if a field has been set.
func (o *Services) HasBuildSHA() bool {
	if o != nil && !IsNil(o.BuildSHA) {
		return true
	}

	return false
}

// SetBuildSHA gets a reference to the given string and assigns it to the BuildSHA field.
func (o *Services) SetBuildSHA(v string) {
	o.BuildSHA = &v
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *Services) GetBuildVersion() string {
	if o == nil || IsNil(o.BuildVersion) {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildVersion) {
		return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *Services) HasBuildVersion() bool {
	if o != nil && !IsNil(o.BuildVersion) {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *Services) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetLinkedSites returns the LinkedSites field value if set, zero value otherwise.
func (o *Services) GetLinkedSites() []string {
	if o == nil || IsNil(o.LinkedSites) {
		var ret []string
		return ret
	}
	return o.LinkedSites
}

// GetLinkedSitesOk returns a tuple with the LinkedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetLinkedSitesOk() ([]string, bool) {
	if o == nil || IsNil(o.LinkedSites) {
		return nil, false
	}
	return o.LinkedSites, true
}

// HasLinkedSites returns a boolean if a field has been set.
func (o *Services) HasLinkedSites() bool {
	if o != nil && !IsNil(o.LinkedSites) {
		return true
	}

	return false
}

// SetLinkedSites gets a reference to the given []string and assigns it to the LinkedSites field.
func (o *Services) SetLinkedSites(v []string) {
	o.LinkedSites = v
}

// GetServiceVersions returns the ServiceVersions field value if set, zero value otherwise.
func (o *Services) GetServiceVersions() []ServiceVersion {
	if o == nil || IsNil(o.ServiceVersions) {
		var ret []ServiceVersion
		return ret
	}
	return o.ServiceVersions
}

// GetServiceVersionsOk returns a tuple with the ServiceVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetServiceVersionsOk() ([]ServiceVersion, bool) {
	if o == nil || IsNil(o.ServiceVersions) {
		return nil, false
	}
	return o.ServiceVersions, true
}

// HasServiceVersions returns a boolean if a field has been set.
func (o *Services) HasServiceVersions() bool {
	if o != nil && !IsNil(o.ServiceVersions) {
		return true
	}

	return false
}

// SetServiceVersions gets a reference to the given []ServiceVersion and assigns it to the ServiceVersions field.
func (o *Services) SetServiceVersions(v []ServiceVersion) {
	o.ServiceVersions = v
}

func (o Services) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Services) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BuildBranch) {
		toSerialize["buildBranch"] = o.BuildBranch
	}
	if !IsNil(o.BuildBranchDeployedDateTime) {
		toSerialize["buildBranchDeployedDateTime"] = o.BuildBranchDeployedDateTime
	}
	if !IsNil(o.BuildSHA) {
		toSerialize["buildSHA"] = o.BuildSHA
	}
	if !IsNil(o.BuildVersion) {
		toSerialize["buildVersion"] = o.BuildVersion
	}
	if !IsNil(o.LinkedSites) {
		toSerialize["linkedSites"] = o.LinkedSites
	}
	if !IsNil(o.ServiceVersions) {
		toSerialize["serviceVersions"] = o.ServiceVersions
	}
	return toSerialize, nil
}

type NullableServices struct {
	value *Services
	isSet bool
}

func (v NullableServices) Get() *Services {
	return v.value
}

func (v *NullableServices) Set(val *Services) {
	v.value = val
	v.isSet = true
}

func (v NullableServices) IsSet() bool {
	return v.isSet
}

func (v *NullableServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServices(val *Services) *NullableServices {
	return &NullableServices{value: val, isSet: true}
}

func (v NullableServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

