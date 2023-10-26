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

// checks if the Brand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Brand{}

// Brand Information about a brand that is associated with an account. A brand applies custom styles and text to an envelope.
type Brand struct {
	// The name of the company associated with the brand.
	BrandCompany *string `json:"brandCompany,omitempty"`
	// The ID used to identify a specific brand in API calls.
	BrandId *string `json:"brandId,omitempty"`
	// An array of two-letter codes for the languages that you want to use with the brand. The supported languages are:  - Arabic (`ar`) - Armenian (`hy`) - Bahasa Indonesia (`id`) - Bahasa Malay (`ms`) - Bulgarian (`bg`) - Chinese Simplified (`zh_CN`) - Chinese Traditional (`zh_TW`) - Croatian (`hr`) - Czech (`cs`) - Danish (`da`) - Dutch (`nl`) - English UK (`en_GB`) - English US (`en`) - Estonian (`et`) - Farsi (`fa`) - Finnish (`fi`) - French (`fr`) - French Canada (`fr_CA`) - German (`de`) - Greek (`el`) - Hebrew (`he`) - Hindi (`hi`) - Hungarian (`hu`) - Italian (`it`) - Japanese (`ja`) - Korean (`ko`) - Latvian (`lv`) - Lithuanian (`lt`) - Norwegian (`no`) - Polish (`pl`) - Portuguese (`pt`) - Portuguese Brasil (`pt_BR`) - Romanian (`ro`) - Russian (`ru`) - Serbian (`sr`) - Slovak (`sk`) - Slovenian (`sl`) - Spanish (`es`) - Spanish Latin America (`es_MX`) - Swedish (`sv`) - Thai (`th`) - Turkish (`tr`) - Ukranian (`uk`) - Vietnamese (`vi`)
	BrandLanguages []string `json:"brandLanguages,omitempty"`
	// The name of the brand.
	BrandName *string `json:"brandName,omitempty"`
	// An array of name-value pairs specifying the colors that the brand uses for the following elements:  - Button background - Button text - Header background - Header text
	Colors []NameValue `json:"colors,omitempty"`
	// The two-letter code for the language that you want to use as the brand default. The supported languages are:  - Arabic (`ar`) - Armenian (`hy`) - Bahasa Indonesia (`id`) - Bahasa Malay (`ms`) - Bulgarian (`bg`) - Chinese Simplified (`zh_CN`) - Chinese Traditional (`zh_TW`) - Croatian (`hr`) - Czech (`cs`) - Danish (`da`) - Dutch (`nl`) - English UK (`en_GB`) - English US (`en`) - Estonian (`et`) - Farsi (`fa`) - Finnish (`fi`) - French (`fr`) - French Canada (`fr_CA`) - German (`de`) - Greek (`el`) - Hebrew (`he`) - Hindi (`hi`) - Hungarian (`hu`) - Italian (`it`) - Japanese (`ja`) - Korean (`ko`) - Latvian (`lv`) - Lithuanian (`lt`) - Norwegian (`no`) - Polish (`pl`) - Portuguese (`pt`) - Portuguese Brasil (`pt_BR`) - Romanian (`ro`) - Russian (`ru`) - Serbian (`sr`) - Slovak (`sk`) - Slovenian (`sl`) - Spanish (`es`) - Spanish Latin America (`es_MX`) - Swedish (`sv`) - Thai (`th`) - Turkish (`tr`) - Ukranian (`uk`) - Vietnamese (`vi`)
	DefaultBrandLanguage *string `json:"defaultBrandLanguage,omitempty"`
	// Deprecated.
	EmailContent []BrandEmailContent `json:"emailContent,omitempty"`
	ErrorDetails *ErrorDetails `json:"errorDetails,omitempty"`
	// 
	IsOrganizationBrand *string `json:"isOrganizationBrand,omitempty"`
	// When **true,** the `brandCompany` property is overriding the name of the company in the account settings.
	IsOverridingCompanyName *bool `json:"isOverridingCompanyName,omitempty"`
	// When **true,** the sending brand is the default brand for sending new envelopes.
	IsSendingDefault *bool `json:"isSendingDefault,omitempty"`
	// When **true,** the siging brand is the default brand for the signing experience.
	IsSigningDefault *bool `json:"isSigningDefault,omitempty"`
	// An array of name/value pairs specifying the pages to which the user is redirected after the following events occur:  - Signing Completed - Viewed Exit - Finish Later - Decline - Session Timeout - Authentication Failure  If you do not specify landing pages, the DocuSign default pages are used.
	LandingPages []NameValue `json:"landingPages,omitempty"`
	// An array of `brandLink` objects that contain information about the links that the brand uses.
	Links []BrandLink `json:"links,omitempty"`
	Logos *BrandLogos `json:"logos,omitempty"`
	// 
	OrganizationBrandLogo *string `json:"organizationBrandLogo,omitempty"`
	Resources *BrandResourceUrls `json:"resources,omitempty"`
}

// NewBrand instantiates a new Brand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrand() *Brand {
	this := Brand{}
	return &this
}

// NewBrandWithDefaults instantiates a new Brand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandWithDefaults() *Brand {
	this := Brand{}
	return &this
}

// GetBrandCompany returns the BrandCompany field value if set, zero value otherwise.
func (o *Brand) GetBrandCompany() string {
	if o == nil || IsNil(o.BrandCompany) {
		var ret string
		return ret
	}
	return *o.BrandCompany
}

// GetBrandCompanyOk returns a tuple with the BrandCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.BrandCompany) {
		return nil, false
	}
	return o.BrandCompany, true
}

// HasBrandCompany returns a boolean if a field has been set.
func (o *Brand) HasBrandCompany() bool {
	if o != nil && !IsNil(o.BrandCompany) {
		return true
	}

	return false
}

// SetBrandCompany gets a reference to the given string and assigns it to the BrandCompany field.
func (o *Brand) SetBrandCompany(v string) {
	o.BrandCompany = &v
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *Brand) GetBrandId() string {
	if o == nil || IsNil(o.BrandId) {
		var ret string
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandIdOk() (*string, bool) {
	if o == nil || IsNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *Brand) HasBrandId() bool {
	if o != nil && !IsNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given string and assigns it to the BrandId field.
func (o *Brand) SetBrandId(v string) {
	o.BrandId = &v
}

// GetBrandLanguages returns the BrandLanguages field value if set, zero value otherwise.
func (o *Brand) GetBrandLanguages() []string {
	if o == nil || IsNil(o.BrandLanguages) {
		var ret []string
		return ret
	}
	return o.BrandLanguages
}

// GetBrandLanguagesOk returns a tuple with the BrandLanguages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandLanguagesOk() ([]string, bool) {
	if o == nil || IsNil(o.BrandLanguages) {
		return nil, false
	}
	return o.BrandLanguages, true
}

// HasBrandLanguages returns a boolean if a field has been set.
func (o *Brand) HasBrandLanguages() bool {
	if o != nil && !IsNil(o.BrandLanguages) {
		return true
	}

	return false
}

// SetBrandLanguages gets a reference to the given []string and assigns it to the BrandLanguages field.
func (o *Brand) SetBrandLanguages(v []string) {
	o.BrandLanguages = v
}

// GetBrandName returns the BrandName field value if set, zero value otherwise.
func (o *Brand) GetBrandName() string {
	if o == nil || IsNil(o.BrandName) {
		var ret string
		return ret
	}
	return *o.BrandName
}

// GetBrandNameOk returns a tuple with the BrandName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetBrandNameOk() (*string, bool) {
	if o == nil || IsNil(o.BrandName) {
		return nil, false
	}
	return o.BrandName, true
}

// HasBrandName returns a boolean if a field has been set.
func (o *Brand) HasBrandName() bool {
	if o != nil && !IsNil(o.BrandName) {
		return true
	}

	return false
}

// SetBrandName gets a reference to the given string and assigns it to the BrandName field.
func (o *Brand) SetBrandName(v string) {
	o.BrandName = &v
}

// GetColors returns the Colors field value if set, zero value otherwise.
func (o *Brand) GetColors() []NameValue {
	if o == nil || IsNil(o.Colors) {
		var ret []NameValue
		return ret
	}
	return o.Colors
}

// GetColorsOk returns a tuple with the Colors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetColorsOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.Colors) {
		return nil, false
	}
	return o.Colors, true
}

// HasColors returns a boolean if a field has been set.
func (o *Brand) HasColors() bool {
	if o != nil && !IsNil(o.Colors) {
		return true
	}

	return false
}

// SetColors gets a reference to the given []NameValue and assigns it to the Colors field.
func (o *Brand) SetColors(v []NameValue) {
	o.Colors = v
}

// GetDefaultBrandLanguage returns the DefaultBrandLanguage field value if set, zero value otherwise.
func (o *Brand) GetDefaultBrandLanguage() string {
	if o == nil || IsNil(o.DefaultBrandLanguage) {
		var ret string
		return ret
	}
	return *o.DefaultBrandLanguage
}

// GetDefaultBrandLanguageOk returns a tuple with the DefaultBrandLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetDefaultBrandLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBrandLanguage) {
		return nil, false
	}
	return o.DefaultBrandLanguage, true
}

// HasDefaultBrandLanguage returns a boolean if a field has been set.
func (o *Brand) HasDefaultBrandLanguage() bool {
	if o != nil && !IsNil(o.DefaultBrandLanguage) {
		return true
	}

	return false
}

// SetDefaultBrandLanguage gets a reference to the given string and assigns it to the DefaultBrandLanguage field.
func (o *Brand) SetDefaultBrandLanguage(v string) {
	o.DefaultBrandLanguage = &v
}

// GetEmailContent returns the EmailContent field value if set, zero value otherwise.
func (o *Brand) GetEmailContent() []BrandEmailContent {
	if o == nil || IsNil(o.EmailContent) {
		var ret []BrandEmailContent
		return ret
	}
	return o.EmailContent
}

// GetEmailContentOk returns a tuple with the EmailContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetEmailContentOk() ([]BrandEmailContent, bool) {
	if o == nil || IsNil(o.EmailContent) {
		return nil, false
	}
	return o.EmailContent, true
}

// HasEmailContent returns a boolean if a field has been set.
func (o *Brand) HasEmailContent() bool {
	if o != nil && !IsNil(o.EmailContent) {
		return true
	}

	return false
}

// SetEmailContent gets a reference to the given []BrandEmailContent and assigns it to the EmailContent field.
func (o *Brand) SetEmailContent(v []BrandEmailContent) {
	o.EmailContent = v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *Brand) GetErrorDetails() ErrorDetails {
	if o == nil || IsNil(o.ErrorDetails) {
		var ret ErrorDetails
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetErrorDetailsOk() (*ErrorDetails, bool) {
	if o == nil || IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *Brand) HasErrorDetails() bool {
	if o != nil && !IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given ErrorDetails and assigns it to the ErrorDetails field.
func (o *Brand) SetErrorDetails(v ErrorDetails) {
	o.ErrorDetails = &v
}

// GetIsOrganizationBrand returns the IsOrganizationBrand field value if set, zero value otherwise.
func (o *Brand) GetIsOrganizationBrand() string {
	if o == nil || IsNil(o.IsOrganizationBrand) {
		var ret string
		return ret
	}
	return *o.IsOrganizationBrand
}

// GetIsOrganizationBrandOk returns a tuple with the IsOrganizationBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIsOrganizationBrandOk() (*string, bool) {
	if o == nil || IsNil(o.IsOrganizationBrand) {
		return nil, false
	}
	return o.IsOrganizationBrand, true
}

// HasIsOrganizationBrand returns a boolean if a field has been set.
func (o *Brand) HasIsOrganizationBrand() bool {
	if o != nil && !IsNil(o.IsOrganizationBrand) {
		return true
	}

	return false
}

// SetIsOrganizationBrand gets a reference to the given string and assigns it to the IsOrganizationBrand field.
func (o *Brand) SetIsOrganizationBrand(v string) {
	o.IsOrganizationBrand = &v
}

// GetIsOverridingCompanyName returns the IsOverridingCompanyName field value if set, zero value otherwise.
func (o *Brand) GetIsOverridingCompanyName() bool {
	if o == nil || IsNil(o.IsOverridingCompanyName) {
		var ret bool
		return ret
	}
	return *o.IsOverridingCompanyName
}

// GetIsOverridingCompanyNameOk returns a tuple with the IsOverridingCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIsOverridingCompanyNameOk() (*bool, bool) {
	if o == nil || IsNil(o.IsOverridingCompanyName) {
		return nil, false
	}
	return o.IsOverridingCompanyName, true
}

// HasIsOverridingCompanyName returns a boolean if a field has been set.
func (o *Brand) HasIsOverridingCompanyName() bool {
	if o != nil && !IsNil(o.IsOverridingCompanyName) {
		return true
	}

	return false
}

// SetIsOverridingCompanyName gets a reference to the given bool and assigns it to the IsOverridingCompanyName field.
func (o *Brand) SetIsOverridingCompanyName(v bool) {
	o.IsOverridingCompanyName = &v
}

// GetIsSendingDefault returns the IsSendingDefault field value if set, zero value otherwise.
func (o *Brand) GetIsSendingDefault() bool {
	if o == nil || IsNil(o.IsSendingDefault) {
		var ret bool
		return ret
	}
	return *o.IsSendingDefault
}

// GetIsSendingDefaultOk returns a tuple with the IsSendingDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIsSendingDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSendingDefault) {
		return nil, false
	}
	return o.IsSendingDefault, true
}

// HasIsSendingDefault returns a boolean if a field has been set.
func (o *Brand) HasIsSendingDefault() bool {
	if o != nil && !IsNil(o.IsSendingDefault) {
		return true
	}

	return false
}

// SetIsSendingDefault gets a reference to the given bool and assigns it to the IsSendingDefault field.
func (o *Brand) SetIsSendingDefault(v bool) {
	o.IsSendingDefault = &v
}

// GetIsSigningDefault returns the IsSigningDefault field value if set, zero value otherwise.
func (o *Brand) GetIsSigningDefault() bool {
	if o == nil || IsNil(o.IsSigningDefault) {
		var ret bool
		return ret
	}
	return *o.IsSigningDefault
}

// GetIsSigningDefaultOk returns a tuple with the IsSigningDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetIsSigningDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSigningDefault) {
		return nil, false
	}
	return o.IsSigningDefault, true
}

// HasIsSigningDefault returns a boolean if a field has been set.
func (o *Brand) HasIsSigningDefault() bool {
	if o != nil && !IsNil(o.IsSigningDefault) {
		return true
	}

	return false
}

// SetIsSigningDefault gets a reference to the given bool and assigns it to the IsSigningDefault field.
func (o *Brand) SetIsSigningDefault(v bool) {
	o.IsSigningDefault = &v
}

// GetLandingPages returns the LandingPages field value if set, zero value otherwise.
func (o *Brand) GetLandingPages() []NameValue {
	if o == nil || IsNil(o.LandingPages) {
		var ret []NameValue
		return ret
	}
	return o.LandingPages
}

// GetLandingPagesOk returns a tuple with the LandingPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetLandingPagesOk() ([]NameValue, bool) {
	if o == nil || IsNil(o.LandingPages) {
		return nil, false
	}
	return o.LandingPages, true
}

// HasLandingPages returns a boolean if a field has been set.
func (o *Brand) HasLandingPages() bool {
	if o != nil && !IsNil(o.LandingPages) {
		return true
	}

	return false
}

// SetLandingPages gets a reference to the given []NameValue and assigns it to the LandingPages field.
func (o *Brand) SetLandingPages(v []NameValue) {
	o.LandingPages = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Brand) GetLinks() []BrandLink {
	if o == nil || IsNil(o.Links) {
		var ret []BrandLink
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetLinksOk() ([]BrandLink, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Brand) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []BrandLink and assigns it to the Links field.
func (o *Brand) SetLinks(v []BrandLink) {
	o.Links = v
}

// GetLogos returns the Logos field value if set, zero value otherwise.
func (o *Brand) GetLogos() BrandLogos {
	if o == nil || IsNil(o.Logos) {
		var ret BrandLogos
		return ret
	}
	return *o.Logos
}

// GetLogosOk returns a tuple with the Logos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetLogosOk() (*BrandLogos, bool) {
	if o == nil || IsNil(o.Logos) {
		return nil, false
	}
	return o.Logos, true
}

// HasLogos returns a boolean if a field has been set.
func (o *Brand) HasLogos() bool {
	if o != nil && !IsNil(o.Logos) {
		return true
	}

	return false
}

// SetLogos gets a reference to the given BrandLogos and assigns it to the Logos field.
func (o *Brand) SetLogos(v BrandLogos) {
	o.Logos = &v
}

// GetOrganizationBrandLogo returns the OrganizationBrandLogo field value if set, zero value otherwise.
func (o *Brand) GetOrganizationBrandLogo() string {
	if o == nil || IsNil(o.OrganizationBrandLogo) {
		var ret string
		return ret
	}
	return *o.OrganizationBrandLogo
}

// GetOrganizationBrandLogoOk returns a tuple with the OrganizationBrandLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetOrganizationBrandLogoOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationBrandLogo) {
		return nil, false
	}
	return o.OrganizationBrandLogo, true
}

// HasOrganizationBrandLogo returns a boolean if a field has been set.
func (o *Brand) HasOrganizationBrandLogo() bool {
	if o != nil && !IsNil(o.OrganizationBrandLogo) {
		return true
	}

	return false
}

// SetOrganizationBrandLogo gets a reference to the given string and assigns it to the OrganizationBrandLogo field.
func (o *Brand) SetOrganizationBrandLogo(v string) {
	o.OrganizationBrandLogo = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Brand) GetResources() BrandResourceUrls {
	if o == nil || IsNil(o.Resources) {
		var ret BrandResourceUrls
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Brand) GetResourcesOk() (*BrandResourceUrls, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Brand) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given BrandResourceUrls and assigns it to the Resources field.
func (o *Brand) SetResources(v BrandResourceUrls) {
	o.Resources = &v
}

func (o Brand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Brand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BrandCompany) {
		toSerialize["brandCompany"] = o.BrandCompany
	}
	if !IsNil(o.BrandId) {
		toSerialize["brandId"] = o.BrandId
	}
	if !IsNil(o.BrandLanguages) {
		toSerialize["brandLanguages"] = o.BrandLanguages
	}
	if !IsNil(o.BrandName) {
		toSerialize["brandName"] = o.BrandName
	}
	if !IsNil(o.Colors) {
		toSerialize["colors"] = o.Colors
	}
	if !IsNil(o.DefaultBrandLanguage) {
		toSerialize["defaultBrandLanguage"] = o.DefaultBrandLanguage
	}
	if !IsNil(o.EmailContent) {
		toSerialize["emailContent"] = o.EmailContent
	}
	if !IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !IsNil(o.IsOrganizationBrand) {
		toSerialize["isOrganizationBrand"] = o.IsOrganizationBrand
	}
	if !IsNil(o.IsOverridingCompanyName) {
		toSerialize["isOverridingCompanyName"] = o.IsOverridingCompanyName
	}
	if !IsNil(o.IsSendingDefault) {
		toSerialize["isSendingDefault"] = o.IsSendingDefault
	}
	if !IsNil(o.IsSigningDefault) {
		toSerialize["isSigningDefault"] = o.IsSigningDefault
	}
	if !IsNil(o.LandingPages) {
		toSerialize["landingPages"] = o.LandingPages
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Logos) {
		toSerialize["logos"] = o.Logos
	}
	if !IsNil(o.OrganizationBrandLogo) {
		toSerialize["organizationBrandLogo"] = o.OrganizationBrandLogo
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableBrand struct {
	value *Brand
	isSet bool
}

func (v NullableBrand) Get() *Brand {
	return v.value
}

func (v *NullableBrand) Set(val *Brand) {
	v.value = val
	v.isSet = true
}

func (v NullableBrand) IsSet() bool {
	return v.isSet
}

func (v *NullableBrand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrand(val *Brand) *NullableBrand {
	return &NullableBrand{value: val, isSet: true}
}

func (v NullableBrand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

