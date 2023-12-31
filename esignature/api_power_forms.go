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


// PowerFormsAPIService PowerFormsAPI service
type PowerFormsAPIService service

type ApiPowerFormsDeletePowerFormRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	powerFormId string
}

func (r ApiPowerFormsDeletePowerFormRequest) Execute() (*http.Response, error) {
	return r.ApiService.PowerFormsDeletePowerFormExecute(r)
}

/*
PowerFormsDeletePowerForm Deletes a PowerForm.

This method deletes a PowerForm.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param powerFormId The ID of the PowerForm.
 @return ApiPowerFormsDeletePowerFormRequest
*/
func (a *PowerFormsAPIService) PowerFormsDeletePowerForm(ctx context.Context, accountId string, powerFormId string) ApiPowerFormsDeletePowerFormRequest {
	return ApiPowerFormsDeletePowerFormRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		powerFormId: powerFormId,
	}
}

// Execute executes the request
func (a *PowerFormsAPIService) PowerFormsDeletePowerFormExecute(r ApiPowerFormsDeletePowerFormRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsDeletePowerForm")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms/{powerFormId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"powerFormId"+"}", url.PathEscape(parameterValueToString(r.powerFormId, "powerFormId")), -1)

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

type ApiPowerFormsDeletePowerFormsListRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	powerFormsRequest *PowerFormsRequest
}

func (r ApiPowerFormsDeletePowerFormsListRequest) PowerFormsRequest(powerFormsRequest PowerFormsRequest) ApiPowerFormsDeletePowerFormsListRequest {
	r.powerFormsRequest = &powerFormsRequest
	return r
}

func (r ApiPowerFormsDeletePowerFormsListRequest) Execute() (*PowerFormsResponse, *http.Response, error) {
	return r.ApiService.PowerFormsDeletePowerFormsListExecute(r)
}

/*
PowerFormsDeletePowerFormsList Deletes one or more PowerForms.

This method deletes one or more PowerForms. The request body takes an array of PowerForm objects that are deleted based on the `powerFormId`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @return ApiPowerFormsDeletePowerFormsListRequest
*/
func (a *PowerFormsAPIService) PowerFormsDeletePowerFormsList(ctx context.Context, accountId string) ApiPowerFormsDeletePowerFormsListRequest {
	return ApiPowerFormsDeletePowerFormsListRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return PowerFormsResponse
func (a *PowerFormsAPIService) PowerFormsDeletePowerFormsListExecute(r ApiPowerFormsDeletePowerFormsListRequest) (*PowerFormsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerFormsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsDeletePowerFormsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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
	localVarPostBody = r.powerFormsRequest
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

type ApiPowerFormsGetPowerFormRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	powerFormId string
}

func (r ApiPowerFormsGetPowerFormRequest) Execute() (*PowerForm, *http.Response, error) {
	return r.ApiService.PowerFormsGetPowerFormExecute(r)
}

/*
PowerFormsGetPowerForm Returns a single PowerForm.

This method returns detailed information about a specific PowerForm.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param powerFormId The ID of the PowerForm.
 @return ApiPowerFormsGetPowerFormRequest
*/
func (a *PowerFormsAPIService) PowerFormsGetPowerForm(ctx context.Context, accountId string, powerFormId string) ApiPowerFormsGetPowerFormRequest {
	return ApiPowerFormsGetPowerFormRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		powerFormId: powerFormId,
	}
}

// Execute executes the request
//  @return PowerForm
func (a *PowerFormsAPIService) PowerFormsGetPowerFormExecute(r ApiPowerFormsGetPowerFormRequest) (*PowerForm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerForm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsGetPowerForm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms/{powerFormId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"powerFormId"+"}", url.PathEscape(parameterValueToString(r.powerFormId, "powerFormId")), -1)

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

type ApiPowerFormsGetPowerFormsListRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	fromDate *string
	order *string
	orderBy *string
	searchFields *string
	searchText *string
	toDate *string
}

// The start date for a date range.  **Note:** If no value is provided, no date filtering is applied.
func (r ApiPowerFormsGetPowerFormsListRequest) FromDate(fromDate string) ApiPowerFormsGetPowerFormsListRequest {
	r.fromDate = &fromDate
	return r
}

// The order in which to sort the results.  Valid values are:    * &#x60;asc&#x60;: Ascending order. * &#x60;desc&#x60;: Descending order. 
func (r ApiPowerFormsGetPowerFormsListRequest) Order(order string) ApiPowerFormsGetPowerFormsListRequest {
	r.order = &order
	return r
}

// The file attribute to use to sort the results.  Valid values are:  - &#x60;sender&#x60; - &#x60;auth&#x60; - &#x60;used&#x60; - &#x60;remaining&#x60; - &#x60;lastused&#x60; - &#x60;status&#x60; - &#x60;type&#x60; - &#x60;templatename&#x60; - &#x60;created&#x60;
func (r ApiPowerFormsGetPowerFormsListRequest) OrderBy(orderBy string) ApiPowerFormsGetPowerFormsListRequest {
	r.orderBy = &orderBy
	return r
}

// A comma-separated list of additional properties to include in a search.  - &#x60;sender&#x60;: Include sender name and email in the search. - &#x60;recipients&#x60;: Include recipient names and emails in the search. - &#x60;envelope&#x60;: Include envelope information in the search. 
func (r ApiPowerFormsGetPowerFormsListRequest) SearchFields(searchFields string) ApiPowerFormsGetPowerFormsListRequest {
	r.searchFields = &searchFields
	return r
}

// Use this parameter to search for specific text.
func (r ApiPowerFormsGetPowerFormsListRequest) SearchText(searchText string) ApiPowerFormsGetPowerFormsListRequest {
	r.searchText = &searchText
	return r
}

// The end date for a date range.  **Note:** If no value is provided, this property defaults to the current date.
func (r ApiPowerFormsGetPowerFormsListRequest) ToDate(toDate string) ApiPowerFormsGetPowerFormsListRequest {
	r.toDate = &toDate
	return r
}

func (r ApiPowerFormsGetPowerFormsListRequest) Execute() (*PowerFormsResponse, *http.Response, error) {
	return r.ApiService.PowerFormsGetPowerFormsListExecute(r)
}

/*
PowerFormsGetPowerFormsList Returns a list of PowerForms.

This method returns a list of PowerForms that are available to the user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @return ApiPowerFormsGetPowerFormsListRequest
*/
func (a *PowerFormsAPIService) PowerFormsGetPowerFormsList(ctx context.Context, accountId string) ApiPowerFormsGetPowerFormsListRequest {
	return ApiPowerFormsGetPowerFormsListRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return PowerFormsResponse
func (a *PowerFormsAPIService) PowerFormsGetPowerFormsListExecute(r ApiPowerFormsGetPowerFormsListRequest) (*PowerFormsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerFormsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsGetPowerFormsList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "from_date", r.fromDate, "")
	}
	if r.order != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order", r.order, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.searchFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_fields", r.searchFields, "")
	}
	if r.searchText != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search_text", r.searchText, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "to_date", r.toDate, "")
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

type ApiPowerFormsGetPowerFormsSendersRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	startPosition *string
}

// The position within the total result set from which to start returning values. The value **thumbnail** may be used to return the page image.
func (r ApiPowerFormsGetPowerFormsSendersRequest) StartPosition(startPosition string) ApiPowerFormsGetPowerFormsSendersRequest {
	r.startPosition = &startPosition
	return r
}

func (r ApiPowerFormsGetPowerFormsSendersRequest) Execute() (*PowerFormSendersResponse, *http.Response, error) {
	return r.ApiService.PowerFormsGetPowerFormsSendersExecute(r)
}

/*
PowerFormsGetPowerFormsSenders Gets PowerForm senders.

This method returns a list of users who have sent PowerForms.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @return ApiPowerFormsGetPowerFormsSendersRequest
*/
func (a *PowerFormsAPIService) PowerFormsGetPowerFormsSenders(ctx context.Context, accountId string) ApiPowerFormsGetPowerFormsSendersRequest {
	return ApiPowerFormsGetPowerFormsSendersRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return PowerFormSendersResponse
func (a *PowerFormsAPIService) PowerFormsGetPowerFormsSendersExecute(r ApiPowerFormsGetPowerFormsSendersRequest) (*PowerFormSendersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerFormSendersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsGetPowerFormsSenders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms/senders"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startPosition != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_position", r.startPosition, "")
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

type ApiPowerFormsPostPowerFormRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	powerForm *PowerForm
}

// Information about any PowerForms that are included in the envelope.
func (r ApiPowerFormsPostPowerFormRequest) PowerForm(powerForm PowerForm) ApiPowerFormsPostPowerFormRequest {
	r.powerForm = &powerForm
	return r
}

func (r ApiPowerFormsPostPowerFormRequest) Execute() (*PowerForm, *http.Response, error) {
	return r.ApiService.PowerFormsPostPowerFormExecute(r)
}

/*
PowerFormsPostPowerForm Creates a new PowerForm

This method creates a new PowerForm.

You create a PowerForm from an existing DocuSign [template](/docs/esign-rest-api/reference/templates/templates/create/), based on the `templateId` in the request body.

 PowerForms that you create from a template are referred to as *web PowerForms*.

**Note:** The DocuSign Admin console also supports creating a PowerForm by uploading a PDF file that has active form fields (referred to as a *PDF PowerForm*). However, PDF PowerForms are deprecated and are not supported in the API.

**Note:** A PowerForm can have only one sender. (Because PowerForms are not necessarily sent by email, this user is also referred to as the PowerForm *initiator*.) If you need to associate multiple senders with a PowerForm, create multiple copies of the PowerForm by using the same template (one copy for each sender). By default, the sender is the PowerForm Administrator who creates the PowerForm.


### Signing modes

You can use one of the following signing modes for a PowerForm:

**`email`**

This mode verifies the recipient's identity by using email authentication before the recipient can sign a document. The recipient enters their email address on the landing page and then clicks **Begin Signing** to begin the signing process. The system then sends an email message with a validation code to the recipient. If the recipient does not provide a valid email address, they do not receive the email message containing the access code and are not able to open and sign the document.

Alternatively, you can make the process easier for signers by using email authentication only and omitting the access code. To do this, you append the `activateonly` flag to the PowerForm URL and set it to true by passing in the value `1`.  When the flag is active, the first recipient receives an email with a link that initiates the signing session without having to enter access code.

Example: `activateonly=1`

**`direct`**

This mode does not require any verification. After a recipient enters their email address on the landing page and clicks **Begin Signing,** a new browser tab opens and the recipient can immediately begin the signing process.

Because the `direct` signing mode does not verify the recipient's identity by using email authentication, we strongly recommend that you use this mode only when the PowerForm is accessible behind a secure portal where the recipient's identity is already authenticated, or where another form of authentication is specified for the recipient in the DocuSign template (for example, an access code, phone authentication, or ID check).

**Note:** In the account settings, `enablePowerFormDirect` must be **true** to use `direct` as the `signingMode`.

### Redirect URLs

You can control the URL to which signers are redirected after signing your PowerForm. However, the URL is specified elsewhere, outside of the PowerForm creation process. For details, see [How do I specify a URL to redirect to when a PowerForm is completed?](https://support.docusign.com/s/articles/How-do-I-specify-a-URL-to-redirect-to-when-a-Powerform-is-completed).

### More information

For more information about creating PowerForms, see [Create a PowerForm](https://support.docusign.com/en/guides/ndse-user-guide-create-a-powerform).



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @return ApiPowerFormsPostPowerFormRequest
*/
func (a *PowerFormsAPIService) PowerFormsPostPowerForm(ctx context.Context, accountId string) ApiPowerFormsPostPowerFormRequest {
	return ApiPowerFormsPostPowerFormRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return PowerForm
func (a *PowerFormsAPIService) PowerFormsPostPowerFormExecute(r ApiPowerFormsPostPowerFormRequest) (*PowerForm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerForm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsPostPowerForm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)

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
	localVarPostBody = r.powerForm
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

type ApiPowerFormsPutPowerFormRequest struct {
	ctx context.Context
	ApiService *PowerFormsAPIService
	accountId string
	powerFormId string
	powerForm *PowerForm
}

// Information about any PowerForms that are included in the envelope.
func (r ApiPowerFormsPutPowerFormRequest) PowerForm(powerForm PowerForm) ApiPowerFormsPutPowerFormRequest {
	r.powerForm = &powerForm
	return r
}

func (r ApiPowerFormsPutPowerFormRequest) Execute() (*PowerForm, *http.Response, error) {
	return r.ApiService.PowerFormsPutPowerFormExecute(r)
}

/*
PowerFormsPutPowerForm Updates an existing PowerForm.

This method updates an existing PowerForm.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param powerFormId The ID of the PowerForm.
 @return ApiPowerFormsPutPowerFormRequest
*/
func (a *PowerFormsAPIService) PowerFormsPutPowerForm(ctx context.Context, accountId string, powerFormId string) ApiPowerFormsPutPowerFormRequest {
	return ApiPowerFormsPutPowerFormRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		powerFormId: powerFormId,
	}
}

// Execute executes the request
//  @return PowerForm
func (a *PowerFormsAPIService) PowerFormsPutPowerFormExecute(r ApiPowerFormsPutPowerFormRequest) (*PowerForm, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PowerForm
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PowerFormsAPIService.PowerFormsPutPowerForm")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/powerforms/{powerFormId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"powerFormId"+"}", url.PathEscape(parameterValueToString(r.powerFormId, "powerFormId")), -1)

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
	localVarPostBody = r.powerForm
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
