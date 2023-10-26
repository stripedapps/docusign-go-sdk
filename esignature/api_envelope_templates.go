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


// EnvelopeTemplatesAPIService EnvelopeTemplatesAPI service
type EnvelopeTemplatesAPIService service

type ApiTemplatesDeleteDocumentTemplatesRequest struct {
	ctx context.Context
	ApiService *EnvelopeTemplatesAPIService
	accountId string
	documentId string
	envelopeId string
	templateId string
}

func (r ApiTemplatesDeleteDocumentTemplatesRequest) Execute() (*http.Response, error) {
	return r.ApiService.TemplatesDeleteDocumentTemplatesExecute(r)
}

/*
TemplatesDeleteDocumentTemplates Deletes a template from a document in an existing envelope.

Deletes the specified template from a document in an existing envelope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @param templateId The ID of the template.
 @return ApiTemplatesDeleteDocumentTemplatesRequest
*/
func (a *EnvelopeTemplatesAPIService) TemplatesDeleteDocumentTemplates(ctx context.Context, accountId string, documentId string, envelopeId string, templateId string) ApiTemplatesDeleteDocumentTemplatesRequest {
	return ApiTemplatesDeleteDocumentTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		envelopeId: envelopeId,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *EnvelopeTemplatesAPIService) TemplatesDeleteDocumentTemplatesExecute(r ApiTemplatesDeleteDocumentTemplatesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeTemplatesAPIService.TemplatesDeleteDocumentTemplates")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/documents/{documentId}/templates/{templateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiTemplatesGetDocumentTemplatesRequest struct {
	ctx context.Context
	ApiService *EnvelopeTemplatesAPIService
	accountId string
	documentId string
	envelopeId string
	include *string
}

// A comma-separated list that limits the results. Valid values are:  * &#x60;applied&#x60; * &#x60;matched&#x60; 
func (r ApiTemplatesGetDocumentTemplatesRequest) Include(include string) ApiTemplatesGetDocumentTemplatesRequest {
	r.include = &include
	return r
}

func (r ApiTemplatesGetDocumentTemplatesRequest) Execute() (*TemplateInformation, *http.Response, error) {
	return r.ApiService.TemplatesGetDocumentTemplatesExecute(r)
}

/*
TemplatesGetDocumentTemplates Gets the templates associated with a document in an existing envelope.

Retrieves the templates associated with a document in the specified envelope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @return ApiTemplatesGetDocumentTemplatesRequest
*/
func (a *EnvelopeTemplatesAPIService) TemplatesGetDocumentTemplates(ctx context.Context, accountId string, documentId string, envelopeId string) ApiTemplatesGetDocumentTemplatesRequest {
	return ApiTemplatesGetDocumentTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		envelopeId: envelopeId,
	}
}

// Execute executes the request
//  @return TemplateInformation
func (a *EnvelopeTemplatesAPIService) TemplatesGetDocumentTemplatesExecute(r ApiTemplatesGetDocumentTemplatesRequest) (*TemplateInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeTemplatesAPIService.TemplatesGetDocumentTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/documents/{documentId}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
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

type ApiTemplatesGetEnvelopeTemplatesRequest struct {
	ctx context.Context
	ApiService *EnvelopeTemplatesAPIService
	accountId string
	envelopeId string
	include *string
}

// The possible value is &#x60;matching_applied&#x60;, which returns template matching information for the template.
func (r ApiTemplatesGetEnvelopeTemplatesRequest) Include(include string) ApiTemplatesGetEnvelopeTemplatesRequest {
	r.include = &include
	return r
}

func (r ApiTemplatesGetEnvelopeTemplatesRequest) Execute() (*TemplateInformation, *http.Response, error) {
	return r.ApiService.TemplatesGetEnvelopeTemplatesExecute(r)
}

/*
TemplatesGetEnvelopeTemplates Get List of Templates used in an Envelope

This returns a list of the server-side templates, their name and ID, used in an envelope.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @return ApiTemplatesGetEnvelopeTemplatesRequest
*/
func (a *EnvelopeTemplatesAPIService) TemplatesGetEnvelopeTemplates(ctx context.Context, accountId string, envelopeId string) ApiTemplatesGetEnvelopeTemplatesRequest {
	return ApiTemplatesGetEnvelopeTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envelopeId: envelopeId,
	}
}

// Execute executes the request
//  @return TemplateInformation
func (a *EnvelopeTemplatesAPIService) TemplatesGetEnvelopeTemplatesExecute(r ApiTemplatesGetEnvelopeTemplatesRequest) (*TemplateInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeTemplatesAPIService.TemplatesGetEnvelopeTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "")
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

type ApiTemplatesPostDocumentTemplatesRequest struct {
	ctx context.Context
	ApiService *EnvelopeTemplatesAPIService
	accountId string
	documentId string
	envelopeId string
	preserveTemplateRecipient *string
	documentTemplateList *DocumentTemplateList
}

// If omitted or set to false (the default), envelope recipients _will be removed_ if the template being applied includes only  tabs positioned via anchor text for the recipient, and none of the documents include the anchor text.   When **true,** the recipients _will be preserved_ after the template is applied.  
func (r ApiTemplatesPostDocumentTemplatesRequest) PreserveTemplateRecipient(preserveTemplateRecipient string) ApiTemplatesPostDocumentTemplatesRequest {
	r.preserveTemplateRecipient = &preserveTemplateRecipient
	return r
}

func (r ApiTemplatesPostDocumentTemplatesRequest) DocumentTemplateList(documentTemplateList DocumentTemplateList) ApiTemplatesPostDocumentTemplatesRequest {
	r.documentTemplateList = &documentTemplateList
	return r
}

func (r ApiTemplatesPostDocumentTemplatesRequest) Execute() (*DocumentTemplateList, *http.Response, error) {
	return r.ApiService.TemplatesPostDocumentTemplatesExecute(r)
}

/*
TemplatesPostDocumentTemplates Adds templates to a document in an  envelope.

Adds templates to a document in the specified envelope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @return ApiTemplatesPostDocumentTemplatesRequest
*/
func (a *EnvelopeTemplatesAPIService) TemplatesPostDocumentTemplates(ctx context.Context, accountId string, documentId string, envelopeId string) ApiTemplatesPostDocumentTemplatesRequest {
	return ApiTemplatesPostDocumentTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		envelopeId: envelopeId,
	}
}

// Execute executes the request
//  @return DocumentTemplateList
func (a *EnvelopeTemplatesAPIService) TemplatesPostDocumentTemplatesExecute(r ApiTemplatesPostDocumentTemplatesRequest) (*DocumentTemplateList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentTemplateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeTemplatesAPIService.TemplatesPostDocumentTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/documents/{documentId}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.preserveTemplateRecipient != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "preserve_template_recipient", r.preserveTemplateRecipient, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.documentTemplateList
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

type ApiTemplatesPostEnvelopeTemplatesRequest struct {
	ctx context.Context
	ApiService *EnvelopeTemplatesAPIService
	accountId string
	envelopeId string
	preserveTemplateRecipient *string
	documentTemplateList *DocumentTemplateList
}

// If omitted or set to false (the default), envelope recipients _will be removed_ if the template being applied includes only  tabs positioned via anchor text for the recipient, and none of the documents include the anchor text.   When **true,** the recipients _will be preserved_ after the template is applied.  
func (r ApiTemplatesPostEnvelopeTemplatesRequest) PreserveTemplateRecipient(preserveTemplateRecipient string) ApiTemplatesPostEnvelopeTemplatesRequest {
	r.preserveTemplateRecipient = &preserveTemplateRecipient
	return r
}

func (r ApiTemplatesPostEnvelopeTemplatesRequest) DocumentTemplateList(documentTemplateList DocumentTemplateList) ApiTemplatesPostEnvelopeTemplatesRequest {
	r.documentTemplateList = &documentTemplateList
	return r
}

func (r ApiTemplatesPostEnvelopeTemplatesRequest) Execute() (*DocumentTemplateList, *http.Response, error) {
	return r.ApiService.TemplatesPostEnvelopeTemplatesExecute(r)
}

/*
TemplatesPostEnvelopeTemplates Adds templates to an envelope.

Adds templates to the specified envelope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param envelopeId The envelope's GUID.   Example: `93be49ab-xxxx-xxxx-xxxx-f752070d71ec` 
 @return ApiTemplatesPostEnvelopeTemplatesRequest
*/
func (a *EnvelopeTemplatesAPIService) TemplatesPostEnvelopeTemplates(ctx context.Context, accountId string, envelopeId string) ApiTemplatesPostEnvelopeTemplatesRequest {
	return ApiTemplatesPostEnvelopeTemplatesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		envelopeId: envelopeId,
	}
}

// Execute executes the request
//  @return DocumentTemplateList
func (a *EnvelopeTemplatesAPIService) TemplatesPostEnvelopeTemplatesExecute(r ApiTemplatesPostEnvelopeTemplatesRequest) (*DocumentTemplateList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentTemplateList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EnvelopeTemplatesAPIService.TemplatesPostEnvelopeTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/envelopes/{envelopeId}/templates"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"envelopeId"+"}", url.PathEscape(parameterValueToString(r.envelopeId, "envelopeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.preserveTemplateRecipient != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "preserve_template_recipient", r.preserveTemplateRecipient, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	// body params
	localVarPostBody = r.documentTemplateList
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
