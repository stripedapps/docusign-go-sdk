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


// TemplateDocumentTabsAPIService TemplateDocumentTabsAPI service
type TemplateDocumentTabsAPIService service

type ApiTabsDeleteTemplateDocumentTabsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentTabsAPIService
	accountId string
	documentId string
	templateId string
	templateTabs *TemplateTabs
}

func (r ApiTabsDeleteTemplateDocumentTabsRequest) TemplateTabs(templateTabs TemplateTabs) ApiTabsDeleteTemplateDocumentTabsRequest {
	r.templateTabs = &templateTabs
	return r
}

func (r ApiTabsDeleteTemplateDocumentTabsRequest) Execute() (*Tabs, *http.Response, error) {
	return r.ApiService.TabsDeleteTemplateDocumentTabsExecute(r)
}

/*
TabsDeleteTemplateDocumentTabs Deletes tabs from a template.

Deletes tabs from the document specified by `documentId` in the
template specified by `templateId`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiTabsDeleteTemplateDocumentTabsRequest
*/
func (a *TemplateDocumentTabsAPIService) TabsDeleteTemplateDocumentTabs(ctx context.Context, accountId string, documentId string, templateId string) ApiTabsDeleteTemplateDocumentTabsRequest {
	return ApiTabsDeleteTemplateDocumentTabsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return Tabs
func (a *TemplateDocumentTabsAPIService) TabsDeleteTemplateDocumentTabsExecute(r ApiTabsDeleteTemplateDocumentTabsRequest) (*Tabs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tabs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentTabsAPIService.TabsDeleteTemplateDocumentTabs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/tabs"
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
	localVarPostBody = r.templateTabs
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

type ApiTabsGetTemplateDocumentTabsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentTabsAPIService
	accountId string
	documentId string
	templateId string
	pageNumbers *string
}

// Filters for tabs that occur on the pages that you specify. Enter as a comma-separated list of page Guids.  Example: &#x60;page_numbers&#x3D;2,6&#x60;
func (r ApiTabsGetTemplateDocumentTabsRequest) PageNumbers(pageNumbers string) ApiTabsGetTemplateDocumentTabsRequest {
	r.pageNumbers = &pageNumbers
	return r
}

func (r ApiTabsGetTemplateDocumentTabsRequest) Execute() (*TemplateDocumentTabs, *http.Response, error) {
	return r.ApiService.TabsGetTemplateDocumentTabsExecute(r)
}

/*
TabsGetTemplateDocumentTabs Returns tabs on a template.

Returns the tabs on the document specified by `documentId` in the
template specified by `templateId`.



 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiTabsGetTemplateDocumentTabsRequest
*/
func (a *TemplateDocumentTabsAPIService) TabsGetTemplateDocumentTabs(ctx context.Context, accountId string, documentId string, templateId string) ApiTabsGetTemplateDocumentTabsRequest {
	return ApiTabsGetTemplateDocumentTabsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateDocumentTabs
func (a *TemplateDocumentTabsAPIService) TabsGetTemplateDocumentTabsExecute(r ApiTabsGetTemplateDocumentTabsRequest) (*TemplateDocumentTabs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateDocumentTabs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentTabsAPIService.TabsGetTemplateDocumentTabs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/tabs"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"templateId"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageNumbers != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page_numbers", r.pageNumbers, "")
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

type ApiTabsGetTemplatePageTabsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentTabsAPIService
	accountId string
	documentId string
	pageNumber string
	templateId string
}

func (r ApiTabsGetTemplatePageTabsRequest) Execute() (*TemplateDocumentTabs, *http.Response, error) {
	return r.ApiService.TabsGetTemplatePageTabsExecute(r)
}

/*
TabsGetTemplatePageTabs Returns tabs on the specified page.

Returns the tabs from the page specified by `pageNumber` of the document specified by `documentId` in the
template specified by `templateId`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param pageNumber The page number being accessed.
 @param templateId The ID of the template.
 @return ApiTabsGetTemplatePageTabsRequest
*/
func (a *TemplateDocumentTabsAPIService) TabsGetTemplatePageTabs(ctx context.Context, accountId string, documentId string, pageNumber string, templateId string) ApiTabsGetTemplatePageTabsRequest {
	return ApiTabsGetTemplatePageTabsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		pageNumber: pageNumber,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateDocumentTabs
func (a *TemplateDocumentTabsAPIService) TabsGetTemplatePageTabsExecute(r ApiTabsGetTemplatePageTabsRequest) (*TemplateDocumentTabs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateDocumentTabs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentTabsAPIService.TabsGetTemplatePageTabs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/pages/{pageNumber}/tabs"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"documentId"+"}", url.PathEscape(parameterValueToString(r.documentId, "documentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", url.PathEscape(parameterValueToString(r.pageNumber, "pageNumber")), -1)
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

type ApiTabsPostTemplateDocumentTabsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentTabsAPIService
	accountId string
	documentId string
	templateId string
	templateTabs *TemplateTabs
}

func (r ApiTabsPostTemplateDocumentTabsRequest) TemplateTabs(templateTabs TemplateTabs) ApiTabsPostTemplateDocumentTabsRequest {
	r.templateTabs = &templateTabs
	return r
}

func (r ApiTabsPostTemplateDocumentTabsRequest) Execute() (*Tabs, *http.Response, error) {
	return r.ApiService.TabsPostTemplateDocumentTabsExecute(r)
}

/*
TabsPostTemplateDocumentTabs Adds tabs to a document in a template.

Adds tabs to the document specified by `documentId` in the
template specified by `templateId`.

In the request body, you only need to specify the tabs that your
are adding. For example, to add a text
[prefill tab](/docs/esign-rest-api/reference/templates/templatedocumenttabs/create/#definition__templatetabs_prefilltabs),
your request body might look like this:

```
{
  "prefillTabs": {
    "textTabs": [
      {
        "value": "a prefill text tab",
        "pageNumber": "1",
        "documentId": "1",
        "xPosition": 316,
        "yPosition": 97
      }
    ]
  }
}
```


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiTabsPostTemplateDocumentTabsRequest
*/
func (a *TemplateDocumentTabsAPIService) TabsPostTemplateDocumentTabs(ctx context.Context, accountId string, documentId string, templateId string) ApiTabsPostTemplateDocumentTabsRequest {
	return ApiTabsPostTemplateDocumentTabsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return Tabs
func (a *TemplateDocumentTabsAPIService) TabsPostTemplateDocumentTabsExecute(r ApiTabsPostTemplateDocumentTabsRequest) (*Tabs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tabs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentTabsAPIService.TabsPostTemplateDocumentTabs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/tabs"
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
	localVarPostBody = r.templateTabs
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

type ApiTabsPutTemplateDocumentTabsRequest struct {
	ctx context.Context
	ApiService *TemplateDocumentTabsAPIService
	accountId string
	documentId string
	templateId string
	templateTabs *TemplateTabs
}

func (r ApiTabsPutTemplateDocumentTabsRequest) TemplateTabs(templateTabs TemplateTabs) ApiTabsPutTemplateDocumentTabsRequest {
	r.templateTabs = &templateTabs
	return r
}

func (r ApiTabsPutTemplateDocumentTabsRequest) Execute() (*Tabs, *http.Response, error) {
	return r.ApiService.TabsPutTemplateDocumentTabsExecute(r)
}

/*
TabsPutTemplateDocumentTabs Updates the tabs for a template.

Updates tabs in the document specified by `documentId` in the
template specified by `templateId`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param documentId The unique ID of the document within the envelope.  Unlike other IDs in the eSignature API, you specify the `documentId` yourself. Typically the first document has the ID `1`, the second document `2`, and so on, but you can use any numbering scheme that fits within a 32-bit signed integer (1 through 2147483647).   Tab objects have a `documentId` property that specifies the document on which to place the tab. 
 @param templateId The ID of the template.
 @return ApiTabsPutTemplateDocumentTabsRequest
*/
func (a *TemplateDocumentTabsAPIService) TabsPutTemplateDocumentTabs(ctx context.Context, accountId string, documentId string, templateId string) ApiTabsPutTemplateDocumentTabsRequest {
	return ApiTabsPutTemplateDocumentTabsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		documentId: documentId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return Tabs
func (a *TemplateDocumentTabsAPIService) TabsPutTemplateDocumentTabsExecute(r ApiTabsPutTemplateDocumentTabsRequest) (*Tabs, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tabs
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateDocumentTabsAPIService.TabsPutTemplateDocumentTabs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/documents/{documentId}/tabs"
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
	localVarPostBody = r.templateTabs
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
