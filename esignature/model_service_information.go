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

// checks if the ServiceInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceInformation{}

// ServiceInformation 
type ServiceInformation struct {
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

// NewServiceInformation instantiates a new ServiceInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInformation() *ServiceInformation {
	this := ServiceInformation{}
	return &this
}

// NewServiceInformationWithDefaults instantiates a new ServiceInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInformationWithDefaults() *ServiceInformation {
	this := ServiceInformation{}
	return &this
}

// GetBuildBranch returns the BuildBranch field value if set, zero value otherwise.
func (o *ServiceInformation) GetBuildBranch() string {
	if o == nil || IsNil(o.BuildBranch) {
		var ret string
		return ret
	}
	return *o.BuildBranch
}

// GetBuildBranchOk returns a tuple with the BuildBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetBuildBranchOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranch) {
		return nil, false
	}
	return o.BuildBranch, true
}

// HasBuildBranch returns a boolean if a field has been set.
func (o *ServiceInformation) HasBuildBranch() bool {
	if o != nil && !IsNil(o.BuildBranch) {
		return true
	}

	return false
}

// SetBuildBranch gets a reference to the given string and assigns it to the BuildBranch field.
func (o *ServiceInformation) SetBuildBranch(v string) {
	o.BuildBranch = &v
}

// GetBuildBranchDeployedDateTime returns the BuildBranchDeployedDateTime field value if set, zero value otherwise.
func (o *ServiceInformation) GetBuildBranchDeployedDateTime() string {
	if o == nil || IsNil(o.BuildBranchDeployedDateTime) {
		var ret string
		return ret
	}
	return *o.BuildBranchDeployedDateTime
}

// GetBuildBranchDeployedDateTimeOk returns a tuple with the BuildBranchDeployedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetBuildBranchDeployedDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.BuildBranchDeployedDateTime) {
		return nil, false
	}
	return o.BuildBranchDeployedDateTime, true
}

// HasBuildBranchDeployedDateTime returns a boolean if a field has been set.
func (o *ServiceInformation) HasBuildBranchDeployedDateTime() bool {
	if o != nil && !IsNil(o.BuildBranchDeployedDateTime) {
		return true
	}

	return false
}

// SetBuildBranchDeployedDateTime gets a reference to the given string and assigns it to the BuildBranchDeployedDateTime field.
func (o *ServiceInformation) SetBuildBranchDeployedDateTime(v string) {
	o.BuildBranchDeployedDateTime = &v
}

// GetBuildSHA returns the BuildSHA field value if set, zero value otherwise.
func (o *ServiceInformation) GetBuildSHA() string {
	if o == nil || IsNil(o.BuildSHA) {
		var ret string
		return ret
	}
	return *o.BuildSHA
}

// GetBuildSHAOk returns a tuple with the BuildSHA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetBuildSHAOk() (*string, bool) {
	if o == nil || IsNil(o.BuildSHA) {
		return nil, false
	}
	return o.BuildSHA, true
}

// HasBuildSHA returns a boolean if a field has been set.
func (o *ServiceInformation) HasBuildSHA() bool {
	if o != nil && !IsNil(o.BuildSHA) {
		return true
	}

	return false
}

// SetBuildSHA gets a reference to the given string and assigns it to the BuildSHA field.
func (o *ServiceInformation) SetBuildSHA(v string) {
	o.BuildSHA = &v
}

// GetBuildVersion returns the BuildVersion field value if set, zero value otherwise.
func (o *ServiceInformation) GetBuildVersion() string {
	if o == nil || IsNil(o.BuildVersion) {
		var ret string
		return ret
	}
	return *o.BuildVersion
}

// GetBuildVersionOk returns a tuple with the BuildVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetBuildVersionOk() (*string, bool) {
	if o == nil || IsNil(o.BuildVersion) {
		return nil, false
	}
	return o.BuildVersion, true
}

// HasBuildVersion returns a boolean if a field has been set.
func (o *ServiceInformation) HasBuildVersion() bool {
	if o != nil && !IsNil(o.BuildVersion) {
		return true
	}

	return false
}

// SetBuildVersion gets a reference to the given string and assigns it to the BuildVersion field.
func (o *ServiceInformation) SetBuildVersion(v string) {
	o.BuildVersion = &v
}

// GetLinkedSites returns the LinkedSites field value if set, zero value otherwise.
func (o *ServiceInformation) GetLinkedSites() []string {
	if o == nil || IsNil(o.LinkedSites) {
		var ret []string
		return ret
	}
	return o.LinkedSites
}

// GetLinkedSitesOk returns a tuple with the LinkedSites field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetLinkedSitesOk() ([]string, bool) {
	if o == nil || IsNil(o.LinkedSites) {
		return nil, false
	}
	return o.LinkedSites, true
}

// HasLinkedSites returns a boolean if a field has been set.
func (o *ServiceInformation) HasLinkedSites() bool {
	if o != nil && !IsNil(o.LinkedSites) {
		return true
	}

	return false
}

// SetLinkedSites gets a reference to the given []string and assigns it to the LinkedSites field.
func (o *ServiceInformation) SetLinkedSites(v []string) {
	o.LinkedSites = v
}

// GetServiceVersions returns the ServiceVersions field value if set, zero value otherwise.
func (o *ServiceInformation) GetServiceVersions() []ServiceVersion {
	if o == nil || IsNil(o.ServiceVersions) {
		var ret []ServiceVersion
		return ret
	}
	return o.ServiceVersions
}

// GetServiceVersionsOk returns a tuple with the ServiceVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInformation) GetServiceVersionsOk() ([]ServiceVersion, bool) {
	if o == nil || IsNil(o.ServiceVersions) {
		return nil, false
	}
	return o.ServiceVersions, true
}

// HasServiceVersions returns a boolean if a field has been set.
func (o *ServiceInformation) HasServiceVersions() bool {
	if o != nil && !IsNil(o.ServiceVersions) {
		return true
	}

	return false
}

// SetServiceVersions gets a reference to the given []ServiceVersion and assigns it to the ServiceVersions field.
func (o *ServiceInformation) SetServiceVersions(v []ServiceVersion) {
	o.ServiceVersions = v
}

func (o ServiceInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceInformation) ToMap() (map[string]interface{}, error) {
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

type NullableServiceInformation struct {
	value *ServiceInformation
	isSet bool
}

func (v NullableServiceInformation) Get() *ServiceInformation {
	return v.value
}

func (v *NullableServiceInformation) Set(val *ServiceInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInformation(val *ServiceInformation) *NullableServiceInformation {
	return &NullableServiceInformation{value: val, isSet: true}
}

func (v NullableServiceInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


