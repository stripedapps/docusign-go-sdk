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


// IdentityVerificationsAPIService IdentityVerificationsAPI service
type IdentityVerificationsAPIService service

type ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest struct {
	ctx context.Context
	ApiService *IdentityVerificationsAPIService
	accountId string
	identityVerificationWorkflowStatus *string
}

func (r ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest) IdentityVerificationWorkflowStatus(identityVerificationWorkflowStatus string) ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest {
	r.identityVerificationWorkflowStatus = &identityVerificationWorkflowStatus
	return r
}

func (r ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest) Execute() (*AccountIdentityVerificationResponse, *http.Response, error) {
	return r.ApiService.AccountIdentityVerificationGetAccountIdentityVerificationExecute(r)
}

/*
AccountIdentityVerificationGetAccountIdentityVerification Retrieves the Identity Verification workflows available to an account.

This method returns a list of Identity Verification workflows that are available to an account.

**Note:** To use this method, you must either be an account administrator or a sender.

### Related topics

- [How to require ID Verification (IDV) for a recipient](/docs/esign-rest-api/how-to/id-verification/)



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @return ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest
*/
func (a *IdentityVerificationsAPIService) AccountIdentityVerificationGetAccountIdentityVerification(ctx context.Context, accountId string) ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest {
	return ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return AccountIdentityVerificationResponse
func (a *IdentityVerificationsAPIService) AccountIdentityVerificationGetAccountIdentityVerificationExecute(r ApiAccountIdentityVerificationGetAccountIdentityVerificationRequest) (*AccountIdentityVerificationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountIdentityVerificationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IdentityVerificationsAPIService.AccountIdentityVerificationGetAccountIdentityVerification")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/identity_verification"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.identityVerificationWorkflowStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identity_verification_workflow_status", r.identityVerificationWorkflowStatus, "")
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