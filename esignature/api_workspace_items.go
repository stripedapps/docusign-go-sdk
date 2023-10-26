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


// WorkspaceItemsAPIService WorkspaceItemsAPI service
type WorkspaceItemsAPIService service

type ApiWorkspaceFileGetWorkspaceFileRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	fileId string
	folderId string
	workspaceId string
	isDownload *string
	pdfVersion *string
}

// When **true,** the &#x60;Content-Disposition&#x60; header is set in the response. The value of the header provides the filename of the file. The default is **false.**
func (r ApiWorkspaceFileGetWorkspaceFileRequest) IsDownload(isDownload string) ApiWorkspaceFileGetWorkspaceFileRequest {
	r.isDownload = &isDownload
	return r
}

// When **true** the file is returned in PDF format.
func (r ApiWorkspaceFileGetWorkspaceFileRequest) PdfVersion(pdfVersion string) ApiWorkspaceFileGetWorkspaceFileRequest {
	r.pdfVersion = &pdfVersion
	return r
}

func (r ApiWorkspaceFileGetWorkspaceFileRequest) Execute() (*http.Response, error) {
	return r.ApiService.WorkspaceFileGetWorkspaceFileExecute(r)
}

/*
WorkspaceFileGetWorkspaceFile Gets a workspace file

This method returns a binary version of a file in a workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param fileId The ID of the file.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFileGetWorkspaceFileRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFileGetWorkspaceFile(ctx context.Context, accountId string, fileId string, folderId string, workspaceId string) ApiWorkspaceFileGetWorkspaceFileRequest {
	return ApiWorkspaceFileGetWorkspaceFileRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		fileId: fileId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceItemsAPIService) WorkspaceFileGetWorkspaceFileExecute(r ApiWorkspaceFileGetWorkspaceFileRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFileGetWorkspaceFile")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}/files/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isDownload != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_download", r.isDownload, "")
	}
	if r.pdfVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pdf_version", r.pdfVersion, "")
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

type ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	fileId string
	folderId string
	workspaceId string
	count *string
	dpi *string
	maxHeight *string
	maxWidth *string
	startPosition *string
}

// The maximum number of results to return.  Use &#x60;start_position&#x60; to specify the number of results to skip. 
func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) Count(count string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	r.count = &count
	return r
}

// The number of dots per inch (DPI) for the resulting images. Valid values are 1-310 DPI. The default value is 94.
func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) Dpi(dpi string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	r.dpi = &dpi
	return r
}

// Sets the maximum height of the returned images in pixels.
func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) MaxHeight(maxHeight string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	r.maxHeight = &maxHeight
	return r
}

// Sets the maximum width of the returned images in pixels.
func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) MaxWidth(maxWidth string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	r.maxWidth = &maxWidth
	return r
}

// The zero-based index of the result from which to start returning results.  Use with &#x60;count&#x60; to limit the number of results.  The default value is &#x60;0&#x60;. 
func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) StartPosition(startPosition string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	r.startPosition = &startPosition
	return r
}

func (r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) Execute() (*PageImages, *http.Response, error) {
	return r.ApiService.WorkspaceFilePagesGetWorkspaceFilePagesExecute(r)
}

/*
WorkspaceFilePagesGetWorkspaceFilePages List File Pages

This method returns a workspace file as rasterized pages.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param fileId The ID of the file.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFilePagesGetWorkspaceFilePages(ctx context.Context, accountId string, fileId string, folderId string, workspaceId string) ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest {
	return ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		fileId: fileId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return PageImages
func (a *WorkspaceItemsAPIService) WorkspaceFilePagesGetWorkspaceFilePagesExecute(r ApiWorkspaceFilePagesGetWorkspaceFilePagesRequest) (*PageImages, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PageImages
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFilePagesGetWorkspaceFilePages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}/files/{fileId}/pages"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.dpi != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dpi", r.dpi, "")
	}
	if r.maxHeight != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_height", r.maxHeight, "")
	}
	if r.maxWidth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max_width", r.maxWidth, "")
	}
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

type ApiWorkspaceFilePostWorkspaceFilesRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	folderId string
	workspaceId string
}

func (r ApiWorkspaceFilePostWorkspaceFilesRequest) Execute() (*WorkspaceItem, *http.Response, error) {
	return r.ApiService.WorkspaceFilePostWorkspaceFilesExecute(r)
}

/*
WorkspaceFilePostWorkspaceFiles Creates a workspace file.

This method adds a file to a workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFilePostWorkspaceFilesRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFilePostWorkspaceFiles(ctx context.Context, accountId string, folderId string, workspaceId string) ApiWorkspaceFilePostWorkspaceFilesRequest {
	return ApiWorkspaceFilePostWorkspaceFilesRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return WorkspaceItem
func (a *WorkspaceItemsAPIService) WorkspaceFilePostWorkspaceFilesExecute(r ApiWorkspaceFilePostWorkspaceFilesRequest) (*WorkspaceItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkspaceItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFilePostWorkspaceFiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}/files"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

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

type ApiWorkspaceFilePutWorkspaceFileRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	fileId string
	folderId string
	workspaceId string
}

func (r ApiWorkspaceFilePutWorkspaceFileRequest) Execute() (*WorkspaceItem, *http.Response, error) {
	return r.ApiService.WorkspaceFilePutWorkspaceFileExecute(r)
}

/*
WorkspaceFilePutWorkspaceFile Update workspace file or folder metadata

This method updates the metadata for one or more specific files or folders in a workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param fileId The ID of the file.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFilePutWorkspaceFileRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFilePutWorkspaceFile(ctx context.Context, accountId string, fileId string, folderId string, workspaceId string) ApiWorkspaceFilePutWorkspaceFileRequest {
	return ApiWorkspaceFilePutWorkspaceFileRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		fileId: fileId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return WorkspaceItem
func (a *WorkspaceItemsAPIService) WorkspaceFilePutWorkspaceFileExecute(r ApiWorkspaceFilePutWorkspaceFileRequest) (*WorkspaceItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkspaceItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFilePutWorkspaceFile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}/files/{fileId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fileId"+"}", url.PathEscape(parameterValueToString(r.fileId, "fileId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

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

type ApiWorkspaceFolderDeleteWorkspaceItemsRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	folderId string
	workspaceId string
	workspaceItemList *WorkspaceItemList
}

func (r ApiWorkspaceFolderDeleteWorkspaceItemsRequest) WorkspaceItemList(workspaceItemList WorkspaceItemList) ApiWorkspaceFolderDeleteWorkspaceItemsRequest {
	r.workspaceItemList = &workspaceItemList
	return r
}

func (r ApiWorkspaceFolderDeleteWorkspaceItemsRequest) Execute() (*http.Response, error) {
	return r.ApiService.WorkspaceFolderDeleteWorkspaceItemsExecute(r)
}

/*
WorkspaceFolderDeleteWorkspaceItems Deletes files or sub-folders from a workspace.

This method deletes one or more files or sub-folders from a workspace folder or root.

Note: To delete items from a workspace, the `status` of the workspace must be `active`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFolderDeleteWorkspaceItemsRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFolderDeleteWorkspaceItems(ctx context.Context, accountId string, folderId string, workspaceId string) ApiWorkspaceFolderDeleteWorkspaceItemsRequest {
	return ApiWorkspaceFolderDeleteWorkspaceItemsRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
func (a *WorkspaceItemsAPIService) WorkspaceFolderDeleteWorkspaceItemsExecute(r ApiWorkspaceFolderDeleteWorkspaceItemsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFolderDeleteWorkspaceItems")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

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
	localVarPostBody = r.workspaceItemList
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

type ApiWorkspaceFolderGetWorkspaceFolderRequest struct {
	ctx context.Context
	ApiService *WorkspaceItemsAPIService
	accountId string
	folderId string
	workspaceId string
	count *string
	includeFiles *string
	includeSubFolders *string
	includeThumbnails *string
	includeUserDetail *string
	startPosition *string
	workspaceUserId *string
}

// The maximum number of results to return.  Use &#x60;start_position&#x60; to specify the number of results to skip. 
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) Count(count string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.count = &count
	return r
}

// When **true,** the response includes file information (in addition to folder information). The default is **false.**
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) IncludeFiles(includeFiles string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.includeFiles = &includeFiles
	return r
}

// When **true,** the response includes information about the sub-folders of the current folder. The default is **false.**
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) IncludeSubFolders(includeSubFolders string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.includeSubFolders = &includeSubFolders
	return r
}

// When **true,** the response returns thumbnails.  The default is **false.**
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) IncludeThumbnails(includeThumbnails string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.includeThumbnails = &includeThumbnails
	return r
}

// When **true,** the response includes extended details about the user. The default is **false.**
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) IncludeUserDetail(includeUserDetail string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.includeUserDetail = &includeUserDetail
	return r
}

// The zero-based index of the result from which to start returning results.  Use with &#x60;count&#x60; to limit the number of results.  The default value is &#x60;0&#x60;. 
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) StartPosition(startPosition string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.startPosition = &startPosition
	return r
}

// If set, the response only includes results associated with the &#x60;userId&#x60; that you specify.
func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) WorkspaceUserId(workspaceUserId string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	r.workspaceUserId = &workspaceUserId
	return r
}

func (r ApiWorkspaceFolderGetWorkspaceFolderRequest) Execute() (*WorkspaceFolderContents, *http.Response, error) {
	return r.ApiService.WorkspaceFolderGetWorkspaceFolderExecute(r)
}

/*
WorkspaceFolderGetWorkspaceFolder List workspace folder contents

This method returns the contents of a workspace folder, which can include sub-folders and files.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The external account number (int) or account ID GUID.
 @param folderId The ID of the folder.
 @param workspaceId The ID of the workspace.
 @return ApiWorkspaceFolderGetWorkspaceFolderRequest
*/
func (a *WorkspaceItemsAPIService) WorkspaceFolderGetWorkspaceFolder(ctx context.Context, accountId string, folderId string, workspaceId string) ApiWorkspaceFolderGetWorkspaceFolderRequest {
	return ApiWorkspaceFolderGetWorkspaceFolderRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		folderId: folderId,
		workspaceId: workspaceId,
	}
}

// Execute executes the request
//  @return WorkspaceFolderContents
func (a *WorkspaceItemsAPIService) WorkspaceFolderGetWorkspaceFolderExecute(r ApiWorkspaceFolderGetWorkspaceFolderRequest) (*WorkspaceFolderContents, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkspaceFolderContents
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WorkspaceItemsAPIService.WorkspaceFolderGetWorkspaceFolder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2.1/accounts/{accountId}/workspaces/{workspaceId}/folders/{folderId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", url.PathEscape(parameterValueToString(r.accountId, "accountId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"folderId"+"}", url.PathEscape(parameterValueToString(r.folderId, "folderId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspaceId"+"}", url.PathEscape(parameterValueToString(r.workspaceId, "workspaceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	}
	if r.includeFiles != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_files", r.includeFiles, "")
	}
	if r.includeSubFolders != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_sub_folders", r.includeSubFolders, "")
	}
	if r.includeThumbnails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_thumbnails", r.includeThumbnails, "")
	}
	if r.includeUserDetail != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_user_detail", r.includeUserDetail, "")
	}
	if r.startPosition != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_position", r.startPosition, "")
	}
	if r.workspaceUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "workspace_user_id", r.workspaceUserId, "")
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