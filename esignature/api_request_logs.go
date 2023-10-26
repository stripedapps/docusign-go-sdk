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
	"os"
)


// RequestLogsAPIService RequestLogsAPI service
type RequestLogsAPIService service

type ApiAPIRequestLogDeleteRequestLogsRequest struct {
	ctx context.Context
	ApiService *RequestLogsAPIService
}

func (r ApiAPIRequestLogDeleteRequestLogsRequest) Execute() (*http.Response, error) {
	return r.ApiService.APIRequestLogDeleteRequestLogsExecute(r)
}

/*
APIRequestLogDeleteRequestLogs Deletes the request log files.

Deletes the request log files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAPIRequestLogDeleteRequestLogsRequest
*/
func (a *RequestLogsAPIService) APIRequestLogDeleteRequestLogs(ctx context.Context) ApiAPIRequestLogDeleteRequestLogsRequest {
	return ApiAPIRequestLogDeleteRequestLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *RequestLogsAPIService) APIRequestLogDeleteRequestLogsExecute(r ApiAPIRequestLogDeleteRequestLogsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestLogsAPIService.APIRequestLogDeleteRequestLogs")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/diagnostics/request_logs"

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

type ApiAPIRequestLogGetRequestLogRequest struct {
	ctx context.Context
	ApiService *RequestLogsAPIService
	requestLogId string
}

func (r ApiAPIRequestLogGetRequestLogRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.APIRequestLogGetRequestLogExecute(r)
}

/*
APIRequestLogGetRequestLog Gets a request logging log file.

Retrieves information for a single log entry.

**Request**
The `requestLogId` property can be retrieved by getting the list of log entries. The Content-Transfer-Encoding header can be set to base64 to retrieve the API request/response as base 64 string. Otherwise the bytes of the request/response are returned.

**Response**
If the Content-Transfer-Encoding header was set to base64, the log is returned as a base64 string.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param requestLogId The ID of the log entry.
 @return ApiAPIRequestLogGetRequestLogRequest
*/
func (a *RequestLogsAPIService) APIRequestLogGetRequestLog(ctx context.Context, requestLogId string) ApiAPIRequestLogGetRequestLogRequest {
	return ApiAPIRequestLogGetRequestLogRequest{
		ApiService: a,
		ctx: ctx,
		requestLogId: requestLogId,
	}
}

// Execute executes the request
//  @return *os.File
func (a *RequestLogsAPIService) APIRequestLogGetRequestLogExecute(r ApiAPIRequestLogGetRequestLogRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestLogsAPIService.APIRequestLogGetRequestLog")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/diagnostics/request_logs/{requestLogId}"
	localVarPath = strings.Replace(localVarPath, "{"+"requestLogId"+"}", url.PathEscape(parameterValueToString(r.requestLogId, "requestLogId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"text/plain"}

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

type ApiAPIRequestLogGetRequestLogSettingsRequest struct {
	ctx context.Context
	ApiService *RequestLogsAPIService
}

func (r ApiAPIRequestLogGetRequestLogSettingsRequest) Execute() (*DiagnosticsSettingsInformation, *http.Response, error) {
	return r.ApiService.APIRequestLogGetRequestLogSettingsExecute(r)
}

/*
APIRequestLogGetRequestLogSettings Gets the API request logging settings.

Retrieves the current API request logging setting for the user and remaining log entries.

**Response**
The response includes the current API request logging setting for the user, along with the maximum log entries and remaining log entries.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAPIRequestLogGetRequestLogSettingsRequest
*/
func (a *RequestLogsAPIService) APIRequestLogGetRequestLogSettings(ctx context.Context) ApiAPIRequestLogGetRequestLogSettingsRequest {
	return ApiAPIRequestLogGetRequestLogSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiagnosticsSettingsInformation
func (a *RequestLogsAPIService) APIRequestLogGetRequestLogSettingsExecute(r ApiAPIRequestLogGetRequestLogSettingsRequest) (*DiagnosticsSettingsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiagnosticsSettingsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestLogsAPIService.APIRequestLogGetRequestLogSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/diagnostics/settings"

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

type ApiAPIRequestLogGetRequestLogsRequest struct {
	ctx context.Context
	ApiService *RequestLogsAPIService
	encoding *string
}

// Reserved for DocuSign.
func (r ApiAPIRequestLogGetRequestLogsRequest) Encoding(encoding string) ApiAPIRequestLogGetRequestLogsRequest {
	r.encoding = &encoding
	return r
}

func (r ApiAPIRequestLogGetRequestLogsRequest) Execute() (*ApiRequestLogsResult, *http.Response, error) {
	return r.ApiService.APIRequestLogGetRequestLogsExecute(r)
}

/*
APIRequestLogGetRequestLogs Gets the API request logging log files.

Retrieves a list of log entries as a JSON or XML object or as a zip file containing the entries.

If the Accept header is set to `application/zip`, the response is a zip file containing individual text files, each representing an API request.

If the Accept header is set to `application/json` or `application/xml`, the response returns list of log entries in either JSON or XML. An example JSON response body is shown below. 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAPIRequestLogGetRequestLogsRequest
*/
func (a *RequestLogsAPIService) APIRequestLogGetRequestLogs(ctx context.Context) ApiAPIRequestLogGetRequestLogsRequest {
	return ApiAPIRequestLogGetRequestLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiRequestLogsResult
func (a *RequestLogsAPIService) APIRequestLogGetRequestLogsExecute(r ApiAPIRequestLogGetRequestLogsRequest) (*ApiRequestLogsResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiRequestLogsResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestLogsAPIService.APIRequestLogGetRequestLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/diagnostics/request_logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.encoding != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "encoding", r.encoding, "")
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

type ApiAPIRequestLogPutRequestLogSettingsRequest struct {
	ctx context.Context
	ApiService *RequestLogsAPIService
	diagnosticsSettingsInformation *DiagnosticsSettingsInformation
}

func (r ApiAPIRequestLogPutRequestLogSettingsRequest) DiagnosticsSettingsInformation(diagnosticsSettingsInformation DiagnosticsSettingsInformation) ApiAPIRequestLogPutRequestLogSettingsRequest {
	r.diagnosticsSettingsInformation = &diagnosticsSettingsInformation
	return r
}

func (r ApiAPIRequestLogPutRequestLogSettingsRequest) Execute() (*DiagnosticsSettingsInformation, *http.Response, error) {
	return r.ApiService.APIRequestLogPutRequestLogSettingsExecute(r)
}

/*
APIRequestLogPutRequestLogSettings Enables or disables API request logging for troubleshooting.

Enables or disables API request logging for troubleshooting.

When enabled (`apiRequestLogging` is **true**),
REST API requests and responses for the user are added to a log.
A log can have up to 50 requests/responses
and the current number of log entries can be determined
by getting the settings.
Logging is automatically disabled when the log limit of 50 is reached.

You can call
[Diagnostics: getRequestLog](/docs/esign-rest-api/reference/diagnostics/requestlogs/get/)
or
[Diagnostics: listRequestLogs](/docs/esign-rest-api/reference/diagnostics/requestlogs/list/)
to download the log files (individually or as a zip file).
Call [Diagnostics: deleteRequestLogs](/docs/esign-rest-api/reference/diagnostics/requestlogs/delete/)
to clear the log by deleting current entries.

Private information, such as passwords and integration key information,
which is normally located in the call header is omitted from the request/response log.

API request logging only captures requests from the authenticated user.
Any call that does not authenticate the user and resolve a userId is not logged.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAPIRequestLogPutRequestLogSettingsRequest
*/
func (a *RequestLogsAPIService) APIRequestLogPutRequestLogSettings(ctx context.Context) ApiAPIRequestLogPutRequestLogSettingsRequest {
	return ApiAPIRequestLogPutRequestLogSettingsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DiagnosticsSettingsInformation
func (a *RequestLogsAPIService) APIRequestLogPutRequestLogSettingsExecute(r ApiAPIRequestLogPutRequestLogSettingsRequest) (*DiagnosticsSettingsInformation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DiagnosticsSettingsInformation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RequestLogsAPIService.APIRequestLogPutRequestLogSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/diagnostics/settings"

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
	localVarPostBody = r.diagnosticsSettingsInformation
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