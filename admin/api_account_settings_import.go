/*
DocuSign Admin API

An API for an organization administrator to manage organizations, accounts and users

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


// AccountSettingsImportAPIService AccountSettingsImportAPI service
type AccountSettingsImportAPIService service

type ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest struct {
	ctx context.Context
	ApiService *AccountSettingsImportAPIService
	organizationId string
	importId string
}

func (r ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportAccountSettingsDeleteByIdExecute(r)
}

/*
OrganizationImportOrganizationImportAccountSettingsDeleteById Deletes a Bulk Account Settings Import request.

Deletes a single account settings import request.
Any data associated with the request is also deleted.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `account_write`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @param importId The import ID GUID for the request.
 @return ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest
*/
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsDeleteById(ctx context.Context, organizationId string, importId string) ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest {
	return ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		importId: importId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsDeleteByIdExecute(r ApiOrganizationImportOrganizationImportAccountSettingsDeleteByIdRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountSettingsImportAPIService.OrganizationImportOrganizationImportAccountSettingsDeleteById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/imports/account_settings/{importId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", url.PathEscape(parameterValueToString(r.importId, "importId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiOrganizationImportOrganizationImportAccountSettingsGetRequest struct {
	ctx context.Context
	ApiService *AccountSettingsImportAPIService
	organizationId string
}

func (r ApiOrganizationImportOrganizationImportAccountSettingsGetRequest) Execute() ([]OrganizationAccountSettingsImportResponse, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportAccountSettingsGetExecute(r)
}

/*
OrganizationImportOrganizationImportAccountSettingsGet Returns the details and metadata for Bulk Account Settings Import requests in the organization.

Returns the details and metadata for Bulk Account Settings Import requests in the organization.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `account_read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @return ApiOrganizationImportOrganizationImportAccountSettingsGetRequest
*/
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsGet(ctx context.Context, organizationId string) ApiOrganizationImportOrganizationImportAccountSettingsGetRequest {
	return ApiOrganizationImportOrganizationImportAccountSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []OrganizationAccountSettingsImportResponse
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsGetExecute(r ApiOrganizationImportOrganizationImportAccountSettingsGetRequest) ([]OrganizationAccountSettingsImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []OrganizationAccountSettingsImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountSettingsImportAPIService.OrganizationImportOrganizationImportAccountSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/imports/account_settings"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest struct {
	ctx context.Context
	ApiService *AccountSettingsImportAPIService
	organizationId string
	importId string
}

func (r ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest) Execute() (*OrganizationAccountSettingsImportResponse, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportAccountSettingsGetByIdExecute(r)
}

/*
OrganizationImportOrganizationImportAccountSettingsGetById Returns the details/metadata for a Bulk Account Settings Import request.

Returns the details/metadata for a Bulk Account Settings Import request.

Given an import ID, this method returns the results of an account settings import request.
To get a list of all the import requests, use `getAccountSettingsImports`.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `account_read`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @param importId The import ID GUID for the request.
 @return ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest
*/
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsGetById(ctx context.Context, organizationId string, importId string) ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest {
	return ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		importId: importId,
	}
}

// Execute executes the request
//  @return OrganizationAccountSettingsImportResponse
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsGetByIdExecute(r ApiOrganizationImportOrganizationImportAccountSettingsGetByIdRequest) (*OrganizationAccountSettingsImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationAccountSettingsImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountSettingsImportAPIService.OrganizationImportOrganizationImportAccountSettingsGetById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/imports/account_settings/{importId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"importId"+"}", url.PathEscape(parameterValueToString(r.importId, "importId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type ApiOrganizationImportOrganizationImportAccountSettingsPostRequest struct {
	ctx context.Context
	ApiService *AccountSettingsImportAPIService
	organizationId string
	fileCsv *os.File
}

// CSV file.
func (r ApiOrganizationImportOrganizationImportAccountSettingsPostRequest) FileCsv(fileCsv *os.File) ApiOrganizationImportOrganizationImportAccountSettingsPostRequest {
	r.fileCsv = fileCsv
	return r
}

func (r ApiOrganizationImportOrganizationImportAccountSettingsPostRequest) Execute() (*OrganizationAccountSettingsImportResponse, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportAccountSettingsPostExecute(r)
}

/*
OrganizationImportOrganizationImportAccountSettingsPost Creates a new account settings import request.

Creates a new account settings import request.

The request for this method consists of
the contents of a CSV
file.

You can export your current settings as a CSV file with
[AccountSettingsExport: createAccountSettingsExport](/docs/admin-api/reference/bulkoperations/accountsettingsexport/createaccountsettingsexport/)
and use it as the basis
of the changes you want to import.


To learn more about the format
of a settings CSV file, see
[Preparing a CSV for account settings import][guide-import]
in the DocuSign Admin Guide.

You can import settings for up to 40 accounts at a time.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `account_write`.

### Related topics

- [Account Settings import][guide-import] guide
- [Account Settings Export][guide-export] guide

[guide-import]: https://support.docusign.com/s/document-item?bundleId=rrf1583359212854&topicId=nwl1583359167434.html
[guide-export]: https://support.docusign.com/s/document-item?bundleId=rrf1583359212854&topicId=kfj1583359164049.html


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID GUID.
 @return ApiOrganizationImportOrganizationImportAccountSettingsPostRequest
*/
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsPost(ctx context.Context, organizationId string) ApiOrganizationImportOrganizationImportAccountSettingsPostRequest {
	return ApiOrganizationImportOrganizationImportAccountSettingsPostRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return OrganizationAccountSettingsImportResponse
func (a *AccountSettingsImportAPIService) OrganizationImportOrganizationImportAccountSettingsPostExecute(r ApiOrganizationImportOrganizationImportAccountSettingsPostRequest) (*OrganizationAccountSettingsImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationAccountSettingsImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountSettingsImportAPIService.OrganizationImportOrganizationImportAccountSettingsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/imports/account_settings"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fileCsv == nil {
		return localVarReturnValue, nil, reportError("fileCsv is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	var fileCsvLocalVarFormFileName string
	var fileCsvLocalVarFileName     string
	var fileCsvLocalVarFileBytes    []byte

	fileCsvLocalVarFormFileName = "file.csv"


	fileCsvLocalVarFile := r.fileCsv

	if fileCsvLocalVarFile != nil {
		fbs, _ := io.ReadAll(fileCsvLocalVarFile)

		fileCsvLocalVarFileBytes = fbs
		fileCsvLocalVarFileName = fileCsvLocalVarFile.Name()
		fileCsvLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileCsvLocalVarFileBytes, fileName: fileCsvLocalVarFileName, formFileName: fileCsvLocalVarFormFileName})
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