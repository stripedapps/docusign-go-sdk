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


// TemplateCustomFieldsAPIService TemplateCustomFieldsAPI service
type TemplateCustomFieldsAPIService service

type ApiCustomFieldsDeleteTemplateCustomFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateCustomFieldsAPIService
	accountId string
	templateId string
	templateCustomFields *TemplateCustomFields
}

func (r ApiCustomFieldsDeleteTemplateCustomFieldsRequest) TemplateCustomFields(templateCustomFields TemplateCustomFields) ApiCustomFieldsDeleteTemplateCustomFieldsRequest {
	r.templateCustomFields = &templateCustomFields
	return r
}

func (r ApiCustomFieldsDeleteTemplateCustomFieldsRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsDeleteTemplateCustomFieldsExecute(r)
}

/*
CustomFieldsDeleteTemplateCustomFields Deletes envelope custom fields in a template.

Deletes envelope custom fields in a template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param templateId The ID of the template.
 @return ApiCustomFieldsDeleteTemplateCustomFieldsRequest
*/
func (a *TemplateCustomFieldsAPIService) CustomFieldsDeleteTemplateCustomFields(ctx context.Context, accountId string, templateId string) ApiCustomFieldsDeleteTemplateCustomFieldsRequest {
	return ApiCustomFieldsDeleteTemplateCustomFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *TemplateCustomFieldsAPIService) CustomFieldsDeleteTemplateCustomFieldsExecute(r ApiCustomFieldsDeleteTemplateCustomFieldsRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateCustomFieldsAPIService.CustomFieldsDeleteTemplateCustomFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/custom_fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
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
	localVarPostBody = r.templateCustomFields
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

type ApiCustomFieldsGetTemplateCustomFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateCustomFieldsAPIService
	accountId string
	templateId string
}

func (r ApiCustomFieldsGetTemplateCustomFieldsRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsGetTemplateCustomFieldsExecute(r)
}

/*
CustomFieldsGetTemplateCustomFields Gets the custom document fields from a template.

Retrieves the custom document field information from an existing template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param templateId The ID of the template.
 @return ApiCustomFieldsGetTemplateCustomFieldsRequest
*/
func (a *TemplateCustomFieldsAPIService) CustomFieldsGetTemplateCustomFields(ctx context.Context, accountId string, templateId string) ApiCustomFieldsGetTemplateCustomFieldsRequest {
	return ApiCustomFieldsGetTemplateCustomFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *TemplateCustomFieldsAPIService) CustomFieldsGetTemplateCustomFieldsExecute(r ApiCustomFieldsGetTemplateCustomFieldsRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateCustomFieldsAPIService.CustomFieldsGetTemplateCustomFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/custom_fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
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

type ApiCustomFieldsPostTemplateCustomFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateCustomFieldsAPIService
	accountId string
	templateId string
	templateCustomFields *TemplateCustomFields
}

func (r ApiCustomFieldsPostTemplateCustomFieldsRequest) TemplateCustomFields(templateCustomFields TemplateCustomFields) ApiCustomFieldsPostTemplateCustomFieldsRequest {
	r.templateCustomFields = &templateCustomFields
	return r
}

func (r ApiCustomFieldsPostTemplateCustomFieldsRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsPostTemplateCustomFieldsExecute(r)
}

/*
CustomFieldsPostTemplateCustomFields Creates custom document fields in an existing template document.

Creates custom document fields in an existing template document.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param templateId The ID of the template.
 @return ApiCustomFieldsPostTemplateCustomFieldsRequest
*/
func (a *TemplateCustomFieldsAPIService) CustomFieldsPostTemplateCustomFields(ctx context.Context, accountId string, templateId string) ApiCustomFieldsPostTemplateCustomFieldsRequest {
	return ApiCustomFieldsPostTemplateCustomFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *TemplateCustomFieldsAPIService) CustomFieldsPostTemplateCustomFieldsExecute(r ApiCustomFieldsPostTemplateCustomFieldsRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateCustomFieldsAPIService.CustomFieldsPostTemplateCustomFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/custom_fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
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
	localVarPostBody = r.templateCustomFields
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

type ApiCustomFieldsPutTemplateCustomFieldsRequest struct {
	ctx context.Context
	ApiService *TemplateCustomFieldsAPIService
	accountId string
	templateId string
	templateCustomFields *TemplateCustomFields
}

func (r ApiCustomFieldsPutTemplateCustomFieldsRequest) TemplateCustomFields(templateCustomFields TemplateCustomFields) ApiCustomFieldsPutTemplateCustomFieldsRequest {
	r.templateCustomFields = &templateCustomFields
	return r
}

func (r ApiCustomFieldsPutTemplateCustomFieldsRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsPutTemplateCustomFieldsExecute(r)
}

/*
CustomFieldsPutTemplateCustomFields Updates envelope custom fields in a template.

Updates the custom fields in a template.

Each custom field used in a template must have a unique name.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param templateId The ID of the template.
 @return ApiCustomFieldsPutTemplateCustomFieldsRequest
*/
func (a *TemplateCustomFieldsAPIService) CustomFieldsPutTemplateCustomFields(ctx context.Context, accountId string, templateId string) ApiCustomFieldsPutTemplateCustomFieldsRequest {
	return ApiCustomFieldsPutTemplateCustomFieldsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *TemplateCustomFieldsAPIService) CustomFieldsPutTemplateCustomFieldsExecute(r ApiCustomFieldsPutTemplateCustomFieldsRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateCustomFieldsAPIService.CustomFieldsPutTemplateCustomFields")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/custom_fields"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
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
	localVarPostBody = r.templateCustomFields
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
