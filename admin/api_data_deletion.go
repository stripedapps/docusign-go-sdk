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
)


// DataDeletionAPIService DataDeletionAPI service
type DataDeletionAPIService service

type ApiDataRedactionRedactIndividualMembershipDataRequest struct {
	ctx context.Context
	ApiService *DataDeletionAPIService
	accountId string
	requestModel *IndividualMembershipDataRedactionRequest
}

func (r ApiDataRedactionRedactIndividualMembershipDataRequest) RequestModel(requestModel IndividualMembershipDataRedactionRequest) ApiDataRedactionRedactIndividualMembershipDataRequest {
	r.requestModel = &requestModel
	return r
}

func (r ApiDataRedactionRedactIndividualMembershipDataRequest) Execute() (*IndividualUserDataRedactionResponse, *http.Response, error) {
	return r.ApiService.DataRedactionRedactIndividualMembershipDataExecute(r)
}

/*
DataRedactionRedactIndividualMembershipData Deletes membership data for a user on an account.

Deletes the data for a single account membership for a specified user.

To call this endpoint: 
* You must be an administrator (or delegated administrator) of the specified account.
* The user must have been closed for at least 24 hours.

**Note:** After a user's data is deleted for every account to which they belong, their user-level data will automatically be deleted.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `user_data_redact`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The ID of the account from which to delete the user's data. 
 @return ApiDataRedactionRedactIndividualMembershipDataRequest
*/
func (a *DataDeletionAPIService) DataRedactionRedactIndividualMembershipData(ctx context.Context, accountId string) ApiDataRedactionRedactIndividualMembershipDataRequest {
	return ApiDataRedactionRedactIndividualMembershipDataRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return IndividualUserDataRedactionResponse
func (a *DataDeletionAPIService) DataRedactionRedactIndividualMembershipDataExecute(r ApiDataRedactionRedactIndividualMembershipDataRequest) (*IndividualUserDataRedactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IndividualUserDataRedactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataDeletionAPIService.DataRedactionRedactIndividualMembershipData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/data_redaction/accounts/{accountId}/user"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestModel == nil {
		return localVarReturnValue, nil, reportError("requestModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestModel
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

type ApiDataRedactionRedactIndividualUserDataRequest struct {
	ctx context.Context
	ApiService *DataDeletionAPIService
	organizationId string
	requestModel *IndividualUserDataRedactionRequest
}

func (r ApiDataRedactionRedactIndividualUserDataRequest) RequestModel(requestModel IndividualUserDataRedactionRequest) ApiDataRedactionRedactIndividualUserDataRequest {
	r.requestModel = &requestModel
	return r
}

func (r ApiDataRedactionRedactIndividualUserDataRequest) Execute() (*IndividualUserDataRedactionResponse, *http.Response, error) {
	return r.ApiService.DataRedactionRedactIndividualUserDataExecute(r)
}

/*
DataRedactionRedactIndividualUserData Deletes data for one or more of a user's account memberships.

Deletes data for one or more of a user's account memberships.

To call this endpoint: 
* You must be an organization administrator (or delegated administrator) with permission to manage the specified accounts or the user's email domain.
* The user must have been closed for at least 24 hours.

**Note:** After a user's data is deleted for every account to which they belong, their user-level data will automatically be deleted.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `user_data_redact`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization ID Guid
 @return ApiDataRedactionRedactIndividualUserDataRequest
*/
func (a *DataDeletionAPIService) DataRedactionRedactIndividualUserData(ctx context.Context, organizationId string) ApiDataRedactionRedactIndividualUserDataRequest {
	return ApiDataRedactionRedactIndividualUserDataRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return IndividualUserDataRedactionResponse
func (a *DataDeletionAPIService) DataRedactionRedactIndividualUserDataExecute(r ApiDataRedactionRedactIndividualUserDataRequest) (*IndividualUserDataRedactionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IndividualUserDataRedactionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DataDeletionAPIService.DataRedactionRedactIndividualUserData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/data_redaction/organizations/{organizationId}/user"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestModel == nil {
		return localVarReturnValue, nil, reportError("requestModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.requestModel
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