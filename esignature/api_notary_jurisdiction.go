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


// NotaryJurisdictionAPIService NotaryJurisdictionAPI service
type NotaryJurisdictionAPIService service

type ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest struct {
	ctx context.Context
	ApiService *NotaryJurisdictionAPIService
	jurisdictionId string
}

func (r ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest) Execute() (*http.Response, error) {
	return r.ApiService.NotaryJurisdictionsDeleteNotaryJurisdictionExecute(r)
}

/*
NotaryJurisdictionsDeleteNotaryJurisdiction Deletes the specified jurisdiction.

Deletes the specified jurisdiction.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jurisdictionId The ID of the jurisdiction. The following jurisdictions are supported:  -  `5 - California` -  `6 - Colorado` -  `9 - Florida` -  `10 - Georgia` -  `12 - Idaho` -  `13 - Illinois` -  `14 - Indiana` -  `15 - Iowa` -  `17 - Kentucky` -  `23 - Minnesota` -  `25 - Missouri` -  `30 - New Jersey` -  `32 - New York` -  `33 - North Carolina` -  `35 - Ohio` -  `37 - Oregon` -  `38 - Pennsylvania` -  `40 - South Carolina` -  `43 - Texas` -  `44 - Utah` -  `47 - Washington` -  `48 - West Virginia` -  `49 - Wisconsin` -  `62 - Florida Commissioner of Deeds` 
 @return ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest
*/
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsDeleteNotaryJurisdiction(ctx context.Context, jurisdictionId string) ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest {
	return ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest{
		ApiService: a,
		ctx: ctx,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsDeleteNotaryJurisdictionExecute(r ApiNotaryJurisdictionsDeleteNotaryJurisdictionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotaryJurisdictionAPIService.NotaryJurisdictionsDeleteNotaryJurisdiction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/current_user/notary/jurisdictions/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", url.PathEscape(parameterValueToString(r.jurisdictionId, "jurisdictionId")), -1)

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

type ApiNotaryJurisdictionsGetNotaryJurisdictionRequest struct {
	ctx context.Context
	ApiService *NotaryJurisdictionAPIService
	jurisdictionId string
}

func (r ApiNotaryJurisdictionsGetNotaryJurisdictionRequest) Execute() (*NotaryJurisdiction, *http.Response, error) {
	return r.ApiService.NotaryJurisdictionsGetNotaryJurisdictionExecute(r)
}

/*
NotaryJurisdictionsGetNotaryJurisdiction Gets a jurisdiction object for the current user. The user must be a notary.

Gets a jurisdiction object for the current user.  The following restrictions apply:

- The current user must be a notary.
- The `jurisdictionId` must be a jurisdiction that the notary is registered for.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jurisdictionId The ID of the jurisdiction. The following jurisdictions are supported:  -  `5 - California` -  `6 - Colorado` -  `9 - Florida` -  `10 - Georgia` -  `12 - Idaho` -  `13 - Illinois` -  `14 - Indiana` -  `15 - Iowa` -  `17 - Kentucky` -  `23 - Minnesota` -  `25 - Missouri` -  `30 - New Jersey` -  `32 - New York` -  `33 - North Carolina` -  `35 - Ohio` -  `37 - Oregon` -  `38 - Pennsylvania` -  `40 - South Carolina` -  `43 - Texas` -  `44 - Utah` -  `47 - Washington` -  `48 - West Virginia` -  `49 - Wisconsin` -  `62 - Florida Commissioner of Deeds` 
 @return ApiNotaryJurisdictionsGetNotaryJurisdictionRequest
*/
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsGetNotaryJurisdiction(ctx context.Context, jurisdictionId string) ApiNotaryJurisdictionsGetNotaryJurisdictionRequest {
	return ApiNotaryJurisdictionsGetNotaryJurisdictionRequest{
		ApiService: a,
		ctx: ctx,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
//  @return NotaryJurisdiction
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsGetNotaryJurisdictionExecute(r ApiNotaryJurisdictionsGetNotaryJurisdictionRequest) (*NotaryJurisdiction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotaryJurisdiction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotaryJurisdictionAPIService.NotaryJurisdictionsGetNotaryJurisdiction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/current_user/notary/jurisdictions/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", url.PathEscape(parameterValueToString(r.jurisdictionId, "jurisdictionId")), -1)

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

type ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest struct {
	ctx context.Context
	ApiService *NotaryJurisdictionAPIService
}

func (r ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest) Execute() (*NotaryJurisdictionList, *http.Response, error) {
	return r.ApiService.NotaryJurisdictionsGetNotaryJurisdictionsExecute(r)
}

/*
NotaryJurisdictionsGetNotaryJurisdictions Returns a list of jurisdictions that the notary is registered in.

Returns a list of jurisdictions that the notary is registered in.
The current user must be a notary.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest
*/
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsGetNotaryJurisdictions(ctx context.Context) ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest {
	return ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotaryJurisdictionList
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsGetNotaryJurisdictionsExecute(r ApiNotaryJurisdictionsGetNotaryJurisdictionsRequest) (*NotaryJurisdictionList, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotaryJurisdictionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotaryJurisdictionAPIService.NotaryJurisdictionsGetNotaryJurisdictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/current_user/notary/jurisdictions"

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

type ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest struct {
	ctx context.Context
	ApiService *NotaryJurisdictionAPIService
	notaryJurisdiction *NotaryJurisdiction
}

func (r ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest) NotaryJurisdiction(notaryJurisdiction NotaryJurisdiction) ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest {
	r.notaryJurisdiction = &notaryJurisdiction
	return r
}

func (r ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest) Execute() (*NotaryJurisdiction, *http.Response, error) {
	return r.ApiService.NotaryJurisdictionsPostNotaryJurisdictionsExecute(r)
}

/*
NotaryJurisdictionsPostNotaryJurisdictions Creates a jurisdiction object.

Creates a jurisdiction object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest
*/
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsPostNotaryJurisdictions(ctx context.Context) ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest {
	return ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NotaryJurisdiction
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsPostNotaryJurisdictionsExecute(r ApiNotaryJurisdictionsPostNotaryJurisdictionsRequest) (*NotaryJurisdiction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotaryJurisdiction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotaryJurisdictionAPIService.NotaryJurisdictionsPostNotaryJurisdictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/current_user/notary/jurisdictions"

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
	localVarPostBody = r.notaryJurisdiction
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

type ApiNotaryJurisdictionsPutNotaryJurisdictionRequest struct {
	ctx context.Context
	ApiService *NotaryJurisdictionAPIService
	jurisdictionId string
	notaryJurisdiction *NotaryJurisdiction
}

func (r ApiNotaryJurisdictionsPutNotaryJurisdictionRequest) NotaryJurisdiction(notaryJurisdiction NotaryJurisdiction) ApiNotaryJurisdictionsPutNotaryJurisdictionRequest {
	r.notaryJurisdiction = &notaryJurisdiction
	return r
}

func (r ApiNotaryJurisdictionsPutNotaryJurisdictionRequest) Execute() (*NotaryJurisdiction, *http.Response, error) {
	return r.ApiService.NotaryJurisdictionsPutNotaryJurisdictionExecute(r)
}

/*
NotaryJurisdictionsPutNotaryJurisdiction Updates the jurisdiction information about a notary.

Updates the jurisdiction information about a notary.

The following restrictions apply:

- The current user must be a notary.
- The `jurisdictionId` path parameter must be a jurisdiction that the notary is registered for.
- The `jurisdictionId` path parameter must match the request body's `jurisdiction.jurisdictionId`.

The request body must have a full `jurisdiction` object for the jurisdiction property.
The best way to do this is to use `getNotaryJurisdiction` to obtain the current values and update the properties you want to change.

For example, assume `getNotaryJurisdiction` returns this:

```
{
    "jurisdiction": {
        "jurisdictionId": "15",
        "name": "Iowa",
        "county": "",
        "enabled": "true",
        "countyInSeal": "false",
        "commissionIdInSeal": "true",
        "stateNameInSeal": "true",
        "notaryPublicInSeal": "true",
        "allowSystemCreatedSeal": "true",
        "allowUserUploadedSeal": "false"
    },
    "commissionId": "123456",
    "commissionExpiration": "2020-08-31T07:00:00.0000000Z",
    "registeredName": "Bob Notary",
    "county": "Adams",
    "sealType": "system_created"
}
```

If you want to change the name of the notary from "Bob Notary" to "Robert Notary", your request body would be:

```
{
    "jurisdiction": {
        "jurisdictionId": "15",
        "name": "Iowa",
        "county": "",
        "enabled": "true",
        "countyInSeal": "false",
        "commissionIdInSeal": "true",
        "stateNameInSeal": "true",
        "notaryPublicInSeal": "true",
        "allowSystemCreatedSeal": "true",
        "allowUserUploadedSeal": "false"
    },
    "commissionId": "123456",
    "commissionExpiration": "2020-08-31T07:00:00.0000000Z",
    "registeredName": "Robert Notary",
    "county": "Adams",
    "sealType": "system_created"
}
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param jurisdictionId The ID of the jurisdiction. The following jurisdictions are supported:  -  `5 - California` -  `6 - Colorado` -  `9 - Florida` -  `10 - Georgia` -  `12 - Idaho` -  `13 - Illinois` -  `14 - Indiana` -  `15 - Iowa` -  `17 - Kentucky` -  `23 - Minnesota` -  `25 - Missouri` -  `30 - New Jersey` -  `32 - New York` -  `33 - North Carolina` -  `35 - Ohio` -  `37 - Oregon` -  `38 - Pennsylvania` -  `40 - South Carolina` -  `43 - Texas` -  `44 - Utah` -  `47 - Washington` -  `48 - West Virginia` -  `49 - Wisconsin` -  `62 - Florida Commissioner of Deeds` 
 @return ApiNotaryJurisdictionsPutNotaryJurisdictionRequest
*/
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsPutNotaryJurisdiction(ctx context.Context, jurisdictionId string) ApiNotaryJurisdictionsPutNotaryJurisdictionRequest {
	return ApiNotaryJurisdictionsPutNotaryJurisdictionRequest{
		ApiService: a,
		ctx: ctx,
		jurisdictionId: jurisdictionId,
	}
}

// Execute executes the request
//  @return NotaryJurisdiction
func (a *NotaryJurisdictionAPIService) NotaryJurisdictionsPutNotaryJurisdictionExecute(r ApiNotaryJurisdictionsPutNotaryJurisdictionRequest) (*NotaryJurisdiction, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NotaryJurisdiction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotaryJurisdictionAPIService.NotaryJurisdictionsPutNotaryJurisdiction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/current_user/notary/jurisdictions/{jurisdictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"jurisdictionId"+"}", url.PathEscape(parameterValueToString(r.jurisdictionId, "jurisdictionId")), -1)

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
	localVarPostBody = r.notaryJurisdiction
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