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


// TemplateViewsAPIService TemplateViewsAPI service
type TemplateViewsAPIService service

type ApiViewsPostTemplateEditViewRequest struct {
	ctx context.Context
	ApiService *TemplateViewsAPIService
	accountId string
	templateId string
	returnUrlRequest *ReturnUrlRequest
}

func (r ApiViewsPostTemplateEditViewRequest) ReturnUrlRequest(returnUrlRequest ReturnUrlRequest) ApiViewsPostTemplateEditViewRequest {
	r.returnUrlRequest = &returnUrlRequest
	return r
}

func (r ApiViewsPostTemplateEditViewRequest) Execute() (*ViewUrl, *http.Response, error) {
	return r.ApiService.ViewsPostTemplateEditViewExecute(r)
}

/*
ViewsPostTemplateEditView Gets a URL for a template edit view.

This method returns a URL for starting an edit view of a template that uses the DocuSign Template UI. The URL can only be used once.

To prevent the user from accessing the sending account, set the `returnUrl` value in the request body.

## Information security notice

If the `returnUrl` value is not set, this method provides full access to the sending account. If the account has administrative privileges, then this method also provides administrator access.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param templateId The ID of the template.
 @return ApiViewsPostTemplateEditViewRequest
*/
func (a *TemplateViewsAPIService) ViewsPostTemplateEditView(ctx context.Context, accountId string, templateId string) ApiViewsPostTemplateEditViewRequest {
	return ApiViewsPostTemplateEditViewRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return ViewUrl
func (a *TemplateViewsAPIService) ViewsPostTemplateEditViewExecute(r ApiViewsPostTemplateEditViewRequest) (*ViewUrl, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ViewUrl
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplateViewsAPIService.ViewsPostTemplateEditView")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/templates/{templateId}/views/edit"
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
	localVarPostBody = r.returnUrlRequest
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
