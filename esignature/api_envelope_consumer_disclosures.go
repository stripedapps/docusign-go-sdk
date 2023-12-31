/*
DocuSign REST API

The DocuSign REST API provides you with a powerful, convenient, and simple Web services API for interacting with DocuSign.

API version: v2.1
Contact: devcenter@docusign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EnvelopeConsumerDisclosuresAPIService EnvelopeConsumerDisclosuresAPI service
type EnvelopeConsumerDisclosuresAPIService service

type ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest struct {
	ctx context.Context
	ApiService *EnvelopeConsumerDisclosuresAPIService
	accountId string
	envelopeId string
	recipientId string
	langCode *string
}

// (Optional) The code for the signer language version of the disclosure that you want to retrieve. The following languages are supported:  - Arabic (&#x60;ar&#x60;) - Bulgarian (&#x60;bg&#x60;) - Czech (&#x60;cs&#x60;) - Chinese Simplified (&#x60;zh_CN&#x60;) - Chinese Traditional (&#x60;zh_TW&#x60;) - Croatian (&#x60;hr&#x60;) - Danish (&#x60;da&#x60;) - Dutch (&#x60;nl&#x60;) - English US (&#x60;en&#x60;) - English UK (&#x60;en_GB&#x60;) - Estonian (&#x60;et&#x60;) - Farsi (&#x60;fa&#x60;) - Finnish (&#x60;fi&#x60;) - French (&#x60;fr&#x60;) - French Canadian (&#x60;fr_CA&#x60;) - German (&#x60;de&#x60;) - Greek (&#x60;el&#x60;) - Hebrew (&#x60;he&#x60;) - Hindi (&#x60;hi&#x60;) - Hungarian (&#x60;hu&#x60;) - Bahasa Indonesian (&#x60;id&#x60;) - Italian (&#x60;it&#x60;) - Japanese (&#x60;ja&#x60;) - Korean (&#x60;ko&#x60;) - Latvian (&#x60;lv&#x60;) - Lithuanian (&#x60;lt&#x60;) - Bahasa Melayu (&#x60;ms&#x60;) - Norwegian (&#x60;no&#x60;) - Polish (&#x60;pl&#x60;) - Portuguese (&#x60;pt&#x60;) - Portuguese Brazil (&#x60;pt_BR&#x60;) - Romanian (&#x60;ro&#x60;) - Russian (&#x60;ru&#x60;) - Serbian (&#x60;sr&#x60;) - Slovak (&#x60;sk&#x60;) - Slovenian (&#x60;sl&#x60;) - Spanish (&#x60;es&#x60;) - Spanish Latin America (&#x60;es_MX&#x60;) - Swedish (&#x60;sv&#x60;) - Thai (&#x60;th&#x60;) - Turkish (&#x60;tr&#x60;) - Ukrainian (&#x60;uk&#x60;)  - Vietnamese (&#x60;vi&#x60;)  Additionally, you can automatically detect the browser language being used by the viewer and display the disclosure in that language by setting the value to &#x60;browser&#x60;.
func (r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest) LangCode(langCode string) ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest {
	r.langCode = &langCode
	return r
}

func (r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest) Execute() (*ConsumerDisclosure, *http.Response, error) {
	return r.ApiService.ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdExecute(r)
}

/*
ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientId Gets the default Electronic Record and Signature Disclosure for an envelope.

Retrieves the default, HTML-formatted Electronic Record and Signature Disclosure (ERSD) for the envelope that you specify. 

This is the default ERSD disclosure that DocuSign provides for the convenience of U.S.-based customers only. This default disclosure is only valid for transactions between U.S.-based parties.

To set the language of the disclosure that you want to retrieve, use the optional `langCode` query parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @param recipientId A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
 @return ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest
*/
func (a *EnvelopeConsumerDisclosuresAPIService) ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientId(ctx context.Context, accountId string, envelopeId string, recipientId string) ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest {
	return ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envelopeId: envelopeId,
		recipientId: recipientId,
	}
}

// Execute executes the request
//  @return ConsumerDisclosure
func (a *EnvelopeConsumerDisclosuresAPIService) ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdExecute(r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdRequest) (*ConsumerDisclosure, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsumerDisclosure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeConsumerDisclosuresAPIService.ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/recipients/{recipientId}/consumer_disclosure"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recipientId"+"}", url.PathEscape(parameterValueToString(r.recipientId, "recipientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.langCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "langCode", r.langCode, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest struct {
	ctx context.Context
	ApiService *EnvelopeConsumerDisclosuresAPIService
	accountId string
	envelopeId string
	langCode string
	recipientId string
	langCode2 *string
}

// (Optional) The code for the signer language version of the disclosure that you want to retrieve, as a query parameter. The following languages are supported:  - Arabic (&#x60;ar&#x60;) - Bulgarian (&#x60;bg&#x60;) - Czech (&#x60;cs&#x60;) - Chinese Simplified (&#x60;zh_CN&#x60;) - Chinese Traditional (&#x60;zh_TW&#x60;) - Croatian (&#x60;hr&#x60;) - Danish (&#x60;da&#x60;) - Dutch (&#x60;nl&#x60;) - English US (&#x60;en&#x60;) - English UK (&#x60;en_GB&#x60;) - Estonian (&#x60;et&#x60;) - Farsi (&#x60;fa&#x60;) - Finnish (&#x60;fi&#x60;) - French (&#x60;fr&#x60;) - French Canadian (&#x60;fr_CA&#x60;) - German (&#x60;de&#x60;) - Greek (&#x60;el&#x60;) - Hebrew (&#x60;he&#x60;) - Hindi (&#x60;hi&#x60;) - Hungarian (&#x60;hu&#x60;) - Bahasa Indonesian (&#x60;id&#x60;) - Italian (&#x60;it&#x60;) - Japanese (&#x60;ja&#x60;) - Korean (&#x60;ko&#x60;) - Latvian (&#x60;lv&#x60;) - Lithuanian (&#x60;lt&#x60;) - Bahasa Melayu (&#x60;ms&#x60;) - Norwegian (&#x60;no&#x60;) - Polish (&#x60;pl&#x60;) - Portuguese (&#x60;pt&#x60;) - Portuguese Brazil (&#x60;pt_BR&#x60;) - Romanian (&#x60;ro&#x60;) - Russian (&#x60;ru&#x60;) - Serbian (&#x60;sr&#x60;) - Slovak (&#x60;sk&#x60;) - Slovenian (&#x60;sl&#x60;) - Spanish (&#x60;es&#x60;) - Spanish Latin America (&#x60;es_MX&#x60;) - Swedish (&#x60;sv&#x60;) - Thai (&#x60;th&#x60;) - Turkish (&#x60;tr&#x60;) - Ukrainian (&#x60;uk&#x60;)  - Vietnamese (&#x60;vi&#x60;)  Additionally, you can automatically detect the browser language being used by the viewer and display the disclosure in that language by setting the value to &#x60;browser&#x60;.
func (r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest) LangCode2(langCode2 string) ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest {
	r.langCode2 = &langCode2
	return r
}

func (r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest) Execute() (*ConsumerDisclosure, *http.Response, error) {
	return r.ApiService.ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeExecute(r)
}

/*
ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCode Gets the Electronic Record and Signature Disclosure for a specific envelope recipient.

Retrieves the HTML-formatted Electronic Record and Signature Disclosure (ERSD) for the envelope recipient that you specify. This disclosure might differ from the account-level disclosure, based on the signing brand applied to the envelope and the recipient's language settings.

To set the language of the disclosure that you want to retrieve, specify the `langCode` as either a path or query parameter.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @param langCode (Optional) The code for the signer language version of the disclosure that you want to retrieve, as a path parameter. The following languages are supported:  - Arabic (`ar`) - Bulgarian (`bg`) - Czech (`cs`) - Chinese Simplified (`zh_CN`) - Chinese Traditional (`zh_TW`) - Croatian (`hr`) - Danish (`da`) - Dutch (`nl`) - English US (`en`) - English UK (`en_GB`) - Estonian (`et`) - Farsi (`fa`) - Finnish (`fi`) - French (`fr`) - French Canadian (`fr_CA`) - German (`de`) - Greek (`el`) - Hebrew (`he`) - Hindi (`hi`) - Hungarian (`hu`) - Bahasa Indonesian (`id`) - Italian (`it`) - Japanese (`ja`) - Korean (`ko`) - Latvian (`lv`) - Lithuanian (`lt`) - Bahasa Melayu (`ms`) - Norwegian (`no`) - Polish (`pl`) - Portuguese (`pt`) - Portuguese Brazil (`pt_BR`) - Romanian (`ro`) - Russian (`ru`) - Serbian (`sr`) - Slovak (`sk`) - Slovenian (`sl`) - Spanish (`es`) - Spanish Latin America (`es_MX`) - Swedish (`sv`) - Thai (`th`) - Turkish (`tr`) - Ukrainian (`uk`)  - Vietnamese (`vi`)  Additionally, you can automatically detect the browser language being used by the viewer and display the disclosure in that language by setting the value to `browser`.
 @param recipientId A local reference used to map recipients to other objects, such as specific document tabs.  A `recipientId` must be either an integer or a GUID, and the `recipientId` must be unique within an envelope.  For example, many envelopes assign the first recipient a `recipientId` of `1`. 
 @return ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest
*/
func (a *EnvelopeConsumerDisclosuresAPIService) ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCode(ctx context.Context, accountId string, envelopeId string, langCode string, recipientId string) ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest {
	return ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envelopeId: envelopeId,
		langCode: langCode,
		recipientId: recipientId,
	}
}

// Execute executes the request
//  @return ConsumerDisclosure
func (a *EnvelopeConsumerDisclosuresAPIService) ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeExecute(r ApiConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCodeRequest) (*ConsumerDisclosure, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConsumerDisclosure
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeConsumerDisclosuresAPIService.ConsumerDisclosureGetConsumerDisclosureEnvelopeIdRecipientIdLangCode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/recipients/{recipientId}/consumer_disclosure/{langCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"langCode"+"}", url.PathEscape(parameterValueToString(r.langCode, "langCode")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"recipientId"+"}", url.PathEscape(parameterValueToString(r.recipientId, "recipientId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.langCode2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "langCode", r.langCode2, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
