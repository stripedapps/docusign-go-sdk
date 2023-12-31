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


// SingleAccountUserImportAPIService SingleAccountUserImportAPI service
type SingleAccountUserImportAPIService service

type ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest struct {
	ctx context.Context
	ApiService *SingleAccountUserImportAPIService
	organizationId string
	accountId string
	fileCsv *os.File
}

// CSV file.
func (r ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest) FileCsv(fileCsv *os.File) ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest {
	r.fileCsv = fileCsv
	return r
}

func (r ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest) Execute() (*OrganizationImportResponse, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportSingleAccountUsersInsertExecute(r)
}

/*
OrganizationImportOrganizationImportSingleAccountUsersInsert Import request for adding a user to a single account within the organization.

Import request for adding a user to a single account within the organization.
This method lets you upload user information without requiring an AccountId column.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `user_write`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @param accountId The account ID Guid
 @return ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest
*/
func (a *SingleAccountUserImportAPIService) OrganizationImportOrganizationImportSingleAccountUsersInsert(ctx context.Context, organizationId string, accountId string) ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest {
	return ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return OrganizationImportResponse
func (a *SingleAccountUserImportAPIService) OrganizationImportOrganizationImportSingleAccountUsersInsertExecute(r ApiOrganizationImportOrganizationImportSingleAccountUsersInsertRequest) (*OrganizationImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SingleAccountUserImportAPIService.OrganizationImportOrganizationImportSingleAccountUsersInsert")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/accounts/{accountId}/imports/bulk_users/add"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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

type ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest struct {
	ctx context.Context
	ApiService *SingleAccountUserImportAPIService
	organizationId string
	accountId string
	fileCsv *os.File
}

// CSV file.
func (r ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest) FileCsv(fileCsv *os.File) ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest {
	r.fileCsv = fileCsv
	return r
}

func (r ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest) Execute() (*OrganizationImportResponse, *http.Response, error) {
	return r.ApiService.OrganizationImportOrganizationImportSingleAccountUsersUpdateExecute(r)
}

/*
OrganizationImportOrganizationImportSingleAccountUsersUpdate Import request for updating users for a single account within the organization.

Import request for updating users for a single account
within the organization.

This method lets you upload user information without requiring an `AccountId` column.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `user_write`.



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @param accountId The account ID Guid
 @return ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest
*/
func (a *SingleAccountUserImportAPIService) OrganizationImportOrganizationImportSingleAccountUsersUpdate(ctx context.Context, organizationId string, accountId string) ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest {
	return ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return OrganizationImportResponse
func (a *SingleAccountUserImportAPIService) OrganizationImportOrganizationImportSingleAccountUsersUpdateExecute(r ApiOrganizationImportOrganizationImportSingleAccountUsersUpdateRequest) (*OrganizationImportResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OrganizationImportResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SingleAccountUserImportAPIService.OrganizationImportOrganizationImportSingleAccountUsersUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/organizations/{organizationId}/accounts/{accountId}/imports/bulk_users/update"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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
