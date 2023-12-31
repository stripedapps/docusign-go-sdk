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

// checks if the RecipientViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientViewRequest{}

// RecipientViewRequest The request body for the [EnvelopeViews: createRecipient](/docs/esign-rest-api/reference/envelopes/envelopeviews/createrecipient/) and [EnvelopeViews: createSharedRecipient](/docs/esign-rest-api/reference/envelopes/envelopeviews/createsharedrecipient/) methods.
type RecipientViewRequest struct {
	// A unique identifier of the authentication event executed by the client application.
	AssertionId *string `json:"assertionId,omitempty"`
	// A sender-generated value that indicates the date and time that the signer was authenticated.
	AuthenticationInstant *string `json:"authenticationInstant,omitempty"`
	// Required. Choose a value that most closely matches the technique your application used to authenticate the recipient / signer.   Choose a value from this list:  * Biometric  * Email * HTTPBasicAuth * Kerberos * KnowledgeBasedAuth * None * PaperDocuments * Password * RSASecureID * SingleSignOn_CASiteminder * SingleSignOn_InfoCard * SingleSignOn_MicrosoftActiveDirectory * SingleSignOn_Other * SingleSignOn_Passport * SingleSignOn_SAML * Smartcard * SSLMutualAuth * X509Certificate  This information is included in the Certificate of Completion.
	AuthenticationMethod *string `json:"authenticationMethod,omitempty"`
	ClientURLs *RecipientTokenClientURLs `json:"clientURLs,omitempty"`
	// A sender-created value. If provided, the recipient is treated as an embedded (captive) recipient or signer.  Use your application's client ID (user ID) for the recipient. Doing so enables the details of your application's authentication of the recipient to be connected to the recipient's signature if the signature is disputed or repudiated.  Maximum length: 100 characters.
	ClientUserId *string `json:"clientUserId,omitempty"`
	// (Required) Specifies the email of the recipient. You can use either `email` and `userName` or `userId` to identify the recipient.
	Email *string `json:"email,omitempty"`
	// 
	FrameAncestors []string `json:"frameAncestors,omitempty"`
	// 
	MessageOrigins []string `json:"messageOrigins,omitempty"`
	// Only used if `pingUrl` is specified. This is the interval, in seconds, between pings on the `pingUrl`.  The default is `300` seconds. Valid values are 60-1200 seconds.
	PingFrequency *string `json:"pingFrequency,omitempty"`
	// The client URL that the DocuSign Signing experience should ping to indicate to the client that Signing is active. An HTTP GET call is executed against the client. The response from the client is ignored. The intent is for the client to reset its session timer when the request is received.
	PingUrl *string `json:"pingUrl,omitempty"`
	// A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
	RecipientId *string `json:"recipientId,omitempty"`
	// (Required) The URL to which the user should be redirected after the signing session has ended.  Maximum Length: 470 characters. If the `returnUrl` exceeds this limit, the user is redirected to a truncated URL Be sure to include `https://` in the URL or redirecting might fail on some browsers.   When DocuSign redirects to this URL, it will include an `event` query parameter that your app can use:  * `access_code_failed`: Recipient used incorrect access code. * `cancel`: Recipient canceled the signing operation,   possibly by using the **Finish Later** option. * `decline`: Recipient declined to sign. * `exception`: A system error occurred during the signing process. * `fax_pending`: Recipient has a fax pending. * `id_check_failed`: Recipient failed an ID check. * `session_timeout`: The session timed out. An account can control this timeout by using the **Signer Session Timeout** option. * `signing_complete`: The recipient completed the signing ceremony. * `ttl_expired`: The Time To Live token for the envelope has expired.   After being successfully invoked, these tokens expire   after five minutes. * `viewing_complete`: The recipient completed viewing an envelope   that is in a read-only/terminal state,   such as completed, declined, or voided.  
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// The domain in which the user authenticated.
	SecurityDomain *string `json:"securityDomain,omitempty"`
	// The user ID of the recipient. You can use either the user ID or email and user name to identify the recipient.   If `userId` is used and a `clientUserId` is provided, the value in the `userId` property must match a `recipientId` (which you can retrieve with a GET recipients call) for the envelope.   If a `userId` is used and a `clientUserId` is not provided, the `userId` must match the user ID of the authenticating user.
	UserId *string `json:"userId,omitempty"`
	// The username of the recipient. You can use either `email` and `userName` or `userId` to identify the recipient.
	UserName *string `json:"userName,omitempty"`
	// Specifies whether a browser should be allowed to render a page in a frame or IFrame. Setting this property ensures that your content is not embedded into unauthorized pages or frames.  Valid values are:  - `deny`: The page cannot be displayed in a frame. - `same_origin`: The page can only be displayed in a frame on the same origin as the page itself. - `allow_from`: The page can only be displayed in a frame on the origin specified by the `xFrameOptionsAllowFromUrl` property.
	XFrameOptions *string `json:"xFrameOptions,omitempty"`
	// When the value of `xFrameOptions` is `allow_from`, this property specifies the origin on which the page is allowed to display in a frame. If the value of `xFrameOptions` is `allow_from`, you must include a value for this property.
	XFrameOptionsAllowFromUrl *string `json:"xFrameOptionsAllowFromUrl,omitempty"`
}

// NewRecipientViewRequest instantiates a new RecipientViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientViewRequest() *RecipientViewRequest {
	this := RecipientViewRequest{}
	return &this
}

// NewRecipientViewRequestWithDefaults instantiates a new RecipientViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientViewRequestWithDefaults() *RecipientViewRequest {
	this := RecipientViewRequest{}
	return &this
}

// GetAssertionId returns the AssertionId field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetAssertionId() string {
	if o == nil || IsNil(o.AssertionId) {
		var ret string
		return ret
	}
	return *o.AssertionId
}

// GetAssertionIdOk returns a tuple with the AssertionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetAssertionIdOk() (*string, bool) {
	if o == nil || IsNil(o.AssertionId) {
		return nil, false
	}
	return o.AssertionId, true
}

// HasAssertionId returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasAssertionId() bool {
	if o != nil && !IsNil(o.AssertionId) {
		return true
	}

	return false
}

// SetAssertionId gets a reference to the given string and assigns it to the AssertionId field.
func (o *RecipientViewRequest) SetAssertionId(v string) {
	o.AssertionId = &v
}

// GetAuthenticationInstant returns the AuthenticationInstant field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetAuthenticationInstant() string {
	if o == nil || IsNil(o.AuthenticationInstant) {
		var ret string
		return ret
	}
	return *o.AuthenticationInstant
}

// GetAuthenticationInstantOk returns a tuple with the AuthenticationInstant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetAuthenticationInstantOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationInstant) {
		return nil, false
	}
	return o.AuthenticationInstant, true
}

// HasAuthenticationInstant returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasAuthenticationInstant() bool {
	if o != nil && !IsNil(o.AuthenticationInstant) {
		return true
	}

	return false
}

// SetAuthenticationInstant gets a reference to the given string and assigns it to the AuthenticationInstant field.
func (o *RecipientViewRequest) SetAuthenticationInstant(v string) {
	o.AuthenticationInstant = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetAuthenticationMethod() string {
	if o == nil || IsNil(o.AuthenticationMethod) {
		var ret string
		return ret
	}
	return *o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetAuthenticationMethodOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationMethod) {
		return nil, false
	}
	return o.AuthenticationMethod, true
}

// HasAuthenticationMethod returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasAuthenticationMethod() bool {
	if o != nil && !IsNil(o.AuthenticationMethod) {
		return true
	}

	return false
}

// SetAuthenticationMethod gets a reference to the given string and assigns it to the AuthenticationMethod field.
func (o *RecipientViewRequest) SetAuthenticationMethod(v string) {
	o.AuthenticationMethod = &v
}

// GetClientURLs returns the ClientURLs field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetClientURLs() RecipientTokenClientURLs {
	if o == nil || IsNil(o.ClientURLs) {
		var ret RecipientTokenClientURLs
		return ret
	}
	return *o.ClientURLs
}

// GetClientURLsOk returns a tuple with the ClientURLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetClientURLsOk() (*RecipientTokenClientURLs, bool) {
	if o == nil || IsNil(o.ClientURLs) {
		return nil, false
	}
	return o.ClientURLs, true
}

// HasClientURLs returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasClientURLs() bool {
	if o != nil && !IsNil(o.ClientURLs) {
		return true
	}

	return false
}

// SetClientURLs gets a reference to the given RecipientTokenClientURLs and assigns it to the ClientURLs field.
func (o *RecipientViewRequest) SetClientURLs(v RecipientTokenClientURLs) {
	o.ClientURLs = &v
}

// GetClientUserId returns the ClientUserId field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetClientUserId() string {
	if o == nil || IsNil(o.ClientUserId) {
		var ret string
		return ret
	}
	return *o.ClientUserId
}

// GetClientUserIdOk returns a tuple with the ClientUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetClientUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientUserId) {
		return nil, false
	}
	return o.ClientUserId, true
}

// HasClientUserId returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasClientUserId() bool {
	if o != nil && !IsNil(o.ClientUserId) {
		return true
	}

	return false
}

// SetClientUserId gets a reference to the given string and assigns it to the ClientUserId field.
func (o *RecipientViewRequest) SetClientUserId(v string) {
	o.ClientUserId = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RecipientViewRequest) SetEmail(v string) {
	o.Email = &v
}

// GetFrameAncestors returns the FrameAncestors field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetFrameAncestors() []string {
	if o == nil || IsNil(o.FrameAncestors) {
		var ret []string
		return ret
	}
	return o.FrameAncestors
}

// GetFrameAncestorsOk returns a tuple with the FrameAncestors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetFrameAncestorsOk() ([]string, bool) {
	if o == nil || IsNil(o.FrameAncestors) {
		return nil, false
	}
	return o.FrameAncestors, true
}

// HasFrameAncestors returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasFrameAncestors() bool {
	if o != nil && !IsNil(o.FrameAncestors) {
		return true
	}

	return false
}

// SetFrameAncestors gets a reference to the given []string and assigns it to the FrameAncestors field.
func (o *RecipientViewRequest) SetFrameAncestors(v []string) {
	o.FrameAncestors = v
}

// GetMessageOrigins returns the MessageOrigins field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetMessageOrigins() []string {
	if o == nil || IsNil(o.MessageOrigins) {
		var ret []string
		return ret
	}
	return o.MessageOrigins
}

// GetMessageOriginsOk returns a tuple with the MessageOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetMessageOriginsOk() ([]string, bool) {
	if o == nil || IsNil(o.MessageOrigins) {
		return nil, false
	}
	return o.MessageOrigins, true
}

// HasMessageOrigins returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasMessageOrigins() bool {
	if o != nil && !IsNil(o.MessageOrigins) {
		return true
	}

	return false
}

// SetMessageOrigins gets a reference to the given []string and assigns it to the MessageOrigins field.
func (o *RecipientViewRequest) SetMessageOrigins(v []string) {
	o.MessageOrigins = v
}

// GetPingFrequency returns the PingFrequency field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetPingFrequency() string {
	if o == nil || IsNil(o.PingFrequency) {
		var ret string
		return ret
	}
	return *o.PingFrequency
}

// GetPingFrequencyOk returns a tuple with the PingFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetPingFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.PingFrequency) {
		return nil, false
	}
	return o.PingFrequency, true
}

// HasPingFrequency returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasPingFrequency() bool {
	if o != nil && !IsNil(o.PingFrequency) {
		return true
	}

	return false
}

// SetPingFrequency gets a reference to the given string and assigns it to the PingFrequency field.
func (o *RecipientViewRequest) SetPingFrequency(v string) {
	o.PingFrequency = &v
}

// GetPingUrl returns the PingUrl field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetPingUrl() string {
	if o == nil || IsNil(o.PingUrl) {
		var ret string
		return ret
	}
	return *o.PingUrl
}

// GetPingUrlOk returns a tuple with the PingUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetPingUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PingUrl) {
		return nil, false
	}
	return o.PingUrl, true
}

// HasPingUrl returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasPingUrl() bool {
	if o != nil && !IsNil(o.PingUrl) {
		return true
	}

	return false
}

// SetPingUrl gets a reference to the given string and assigns it to the PingUrl field.
func (o *RecipientViewRequest) SetPingUrl(v string) {
	o.PingUrl = &v
}

// GetRecipientId returns the RecipientId field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetRecipientId() string {
	if o == nil || IsNil(o.RecipientId) {
		var ret string
		return ret
	}
	return *o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetRecipientIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientId) {
		return nil, false
	}
	return o.RecipientId, true
}

// HasRecipientId returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasRecipientId() bool {
	if o != nil && !IsNil(o.RecipientId) {
		return true
	}

	return false
}

// SetRecipientId gets a reference to the given string and assigns it to the RecipientId field.
func (o *RecipientViewRequest) SetRecipientId(v string) {
	o.RecipientId = &v
}

// GetReturnUrl returns the ReturnUrl field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetReturnUrl() string {
	if o == nil || IsNil(o.ReturnUrl) {
		var ret string
		return ret
	}
	return *o.ReturnUrl
}

// GetReturnUrlOk returns a tuple with the ReturnUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetReturnUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReturnUrl) {
		return nil, false
	}
	return o.ReturnUrl, true
}

// HasReturnUrl returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasReturnUrl() bool {
	if o != nil && !IsNil(o.ReturnUrl) {
		return true
	}

	return false
}

// SetReturnUrl gets a reference to the given string and assigns it to the ReturnUrl field.
func (o *RecipientViewRequest) SetReturnUrl(v string) {
	o.ReturnUrl = &v
}

// GetSecurityDomain returns the SecurityDomain field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetSecurityDomain() string {
	if o == nil || IsNil(o.SecurityDomain) {
		var ret string
		return ret
	}
	return *o.SecurityDomain
}

// GetSecurityDomainOk returns a tuple with the SecurityDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetSecurityDomainOk() (*string, bool) {
	if o == nil || IsNil(o.SecurityDomain) {
		return nil, false
	}
	return o.SecurityDomain, true
}

// HasSecurityDomain returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasSecurityDomain() bool {
	if o != nil && !IsNil(o.SecurityDomain) {
		return true
	}

	return false
}

// SetSecurityDomain gets a reference to the given string and assigns it to the SecurityDomain field.
func (o *RecipientViewRequest) SetSecurityDomain(v string) {
	o.SecurityDomain = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *RecipientViewRequest) SetUserId(v string) {
	o.UserId = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *RecipientViewRequest) SetUserName(v string) {
	o.UserName = &v
}

// GetXFrameOptions returns the XFrameOptions field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetXFrameOptions() string {
	if o == nil || IsNil(o.XFrameOptions) {
		var ret string
		return ret
	}
	return *o.XFrameOptions
}

// GetXFrameOptionsOk returns a tuple with the XFrameOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetXFrameOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.XFrameOptions) {
		return nil, false
	}
	return o.XFrameOptions, true
}

// HasXFrameOptions returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasXFrameOptions() bool {
	if o != nil && !IsNil(o.XFrameOptions) {
		return true
	}

	return false
}

// SetXFrameOptions gets a reference to the given string and assigns it to the XFrameOptions field.
func (o *RecipientViewRequest) SetXFrameOptions(v string) {
	o.XFrameOptions = &v
}

// GetXFrameOptionsAllowFromUrl returns the XFrameOptionsAllowFromUrl field value if set, zero value otherwise.
func (o *RecipientViewRequest) GetXFrameOptionsAllowFromUrl() string {
	if o == nil || IsNil(o.XFrameOptionsAllowFromUrl) {
		var ret string
		return ret
	}
	return *o.XFrameOptionsAllowFromUrl
}

// GetXFrameOptionsAllowFromUrlOk returns a tuple with the XFrameOptionsAllowFromUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientViewRequest) GetXFrameOptionsAllowFromUrlOk() (*string, bool) {
	if o == nil || IsNil(o.XFrameOptionsAllowFromUrl) {
		return nil, false
	}
	return o.XFrameOptionsAllowFromUrl, true
}

// HasXFrameOptionsAllowFromUrl returns a boolean if a field has been set.
func (o *RecipientViewRequest) HasXFrameOptionsAllowFromUrl() bool {
	if o != nil && !IsNil(o.XFrameOptionsAllowFromUrl) {
		return true
	}

	return false
}

// SetXFrameOptionsAllowFromUrl gets a reference to the given string and assigns it to the XFrameOptionsAllowFromUrl field.
func (o *RecipientViewRequest) SetXFrameOptionsAllowFromUrl(v string) {
	o.XFrameOptionsAllowFromUrl = &v
}

func (o RecipientViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssertionId) {
		toSerialize["assertionId"] = o.AssertionId
	}
	if !IsNil(o.AuthenticationInstant) {
		toSerialize["authenticationInstant"] = o.AuthenticationInstant
	}
	if !IsNil(o.AuthenticationMethod) {
		toSerialize["authenticationMethod"] = o.AuthenticationMethod
	}
	if !IsNil(o.ClientURLs) {
		toSerialize["clientURLs"] = o.ClientURLs
	}
	if !IsNil(o.ClientUserId) {
		toSerialize["clientUserId"] = o.ClientUserId
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FrameAncestors) {
		toSerialize["frameAncestors"] = o.FrameAncestors
	}
	if !IsNil(o.MessageOrigins) {
		toSerialize["messageOrigins"] = o.MessageOrigins
	}
	if !IsNil(o.PingFrequency) {
		toSerialize["pingFrequency"] = o.PingFrequency
	}
	if !IsNil(o.PingUrl) {
		toSerialize["pingUrl"] = o.PingUrl
	}
	if !IsNil(o.RecipientId) {
		toSerialize["recipientId"] = o.RecipientId
	}
	if !IsNil(o.ReturnUrl) {
		toSerialize["returnUrl"] = o.ReturnUrl
	}
	if !IsNil(o.SecurityDomain) {
		toSerialize["securityDomain"] = o.SecurityDomain
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	if !IsNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if !IsNil(o.XFrameOptions) {
		toSerialize["xFrameOptions"] = o.XFrameOptions
	}
	if !IsNil(o.XFrameOptionsAllowFromUrl) {
		toSerialize["xFrameOptionsAllowFromUrl"] = o.XFrameOptionsAllowFromUrl
	}
	return toSerialize, nil
}

type NullableRecipientViewRequest struct {
	value *RecipientViewRequest
	isSet bool
}

func (v NullableRecipientViewRequest) Get() *RecipientViewRequest {
	return v.value
}

func (v *NullableRecipientViewRequest) Set(val *RecipientViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientViewRequest(val *RecipientViewRequest) *NullableRecipientViewRequest {
	return &NullableRecipientViewRequest{value: val, isSet: true}
}

func (v NullableRecipientViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


