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


// TemplateDocumentFieldsAPIService TemplateDocumentFieldsAPI service
type TemplateDocumentFieldsAPIService service

type ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentFieldsAPIService
	accountId string
	documentId string
	templateId string
	documentFieldsInformation *DocumentFieldsInformation
}

func (r ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest) DocumentFieldsInformation(documentFieldsInformation DocumentFieldsInformation) ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest {
	r.documentFieldsInformation = &documentFieldsInformation
	return r
}

func (r ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest) Execute() (*DocumentFieldsInformation, *http.Response, error) {
	return r.ApiService.DocumentFieldsDeleteTemplateDocumentFieldsExecute(r)
}

/*
DocumentFieldsDeleteTemplateDocumentFields Deletes custom document fields from an existing template document.

Deletes custom document fields from an existing template document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest
*/
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsDeleteTemplateDocumentFields(ctx context.Context, accountId string, documentId string, templateId string) ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest {
	return ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return DocumentFieldsInformation
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsDeleteTemplateDocumentFieldsExecute(r ApiDocumentFieldsDeleteTemplateDocumentFieldsRequest) (*DocumentFieldsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentFieldsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentFieldsAPIService.DocumentFieldsDeleteTemplateDocumentFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.documentFieldsInformation
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

type ApiDocumentFieldsGetTemplateDocumentFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentFieldsAPIService
	accountId string
	documentId string
	templateId string
}

func (r ApiDocumentFieldsGetTemplateDocumentFieldsRequest) Execute() (*DocumentFieldsInformation, *http.Response, error) {
	return r.ApiService.DocumentFieldsGetTemplateDocumentFieldsExecute(r)
}

/*
DocumentFieldsGetTemplateDocumentFields Gets the custom document fields for a an existing template document.

This method retrieves the custom document fields for an existing template document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiDocumentFieldsGetTemplateDocumentFieldsRequest
*/
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsGetTemplateDocumentFields(ctx context.Context, accountId string, documentId string, templateId string) ApiDocumentFieldsGetTemplateDocumentFieldsRequest {
	return ApiDocumentFieldsGetTemplateDocumentFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return DocumentFieldsInformation
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsGetTemplateDocumentFieldsExecute(r ApiDocumentFieldsGetTemplateDocumentFieldsRequest) (*DocumentFieldsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentFieldsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentFieldsAPIService.DocumentFieldsGetTemplateDocumentFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
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

type ApiDocumentFieldsPostTemplateDocumentFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentFieldsAPIService
	accountId string
	documentId string
	templateId string
	documentFieldsInformation *DocumentFieldsInformation
}

func (r ApiDocumentFieldsPostTemplateDocumentFieldsRequest) DocumentFieldsInformation(documentFieldsInformation DocumentFieldsInformation) ApiDocumentFieldsPostTemplateDocumentFieldsRequest {
	r.documentFieldsInformation = &documentFieldsInformation
	return r
}

func (r ApiDocumentFieldsPostTemplateDocumentFieldsRequest) Execute() (*DocumentFieldsInformation, *http.Response, error) {
	return r.ApiService.DocumentFieldsPostTemplateDocumentFieldsExecute(r)
}

/*
DocumentFieldsPostTemplateDocumentFields Creates custom document fields in an existing template document.

Creates custom document fields in an existing template document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiDocumentFieldsPostTemplateDocumentFieldsRequest
*/
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsPostTemplateDocumentFields(ctx context.Context, accountId string, documentId string, templateId string) ApiDocumentFieldsPostTemplateDocumentFieldsRequest {
	return ApiDocumentFieldsPostTemplateDocumentFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return DocumentFieldsInformation
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsPostTemplateDocumentFieldsExecute(r ApiDocumentFieldsPostTemplateDocumentFieldsRequest) (*DocumentFieldsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentFieldsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentFieldsAPIService.DocumentFieldsPostTemplateDocumentFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.documentFieldsInformation
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

type ApiDocumentFieldsPutTemplateDocumentFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentFieldsAPIService
	accountId string
	documentId string
	templateId string
	documentFieldsInformation *DocumentFieldsInformation
}

func (r ApiDocumentFieldsPutTemplateDocumentFieldsRequest) DocumentFieldsInformation(documentFieldsInformation DocumentFieldsInformation) ApiDocumentFieldsPutTemplateDocumentFieldsRequest {
	r.documentFieldsInformation = &documentFieldsInformation
	return r
}

func (r ApiDocumentFieldsPutTemplateDocumentFieldsRequest) Execute() (*DocumentFieldsInformation, *http.Response, error) {
	return r.ApiService.DocumentFieldsPutTemplateDocumentFieldsExecute(r)
}

/*
DocumentFieldsPutTemplateDocumentFields Updates existing custom document fields in an existing template document.

Updates existing custom document fields in an existing template document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiDocumentFieldsPutTemplateDocumentFieldsRequest
*/
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsPutTemplateDocumentFields(ctx context.Context, accountId string, documentId string, templateId string) ApiDocumentFieldsPutTemplateDocumentFieldsRequest {
	return ApiDocumentFieldsPutTemplateDocumentFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return DocumentFieldsInformation
func (a *TemplateDocumentFieldsAPIService) DocumentFieldsPutTemplateDocumentFieldsExecute(r ApiDocumentFieldsPutTemplateDocumentFieldsRequest) (*DocumentFieldsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DocumentFieldsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentFieldsAPIService.DocumentFieldsPutTemplateDocumentFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.documentFieldsInformation
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
