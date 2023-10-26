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
	"time"
)


// AccountCloningAPIService AccountCloningAPI service
type AccountCloningAPIService service

type ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest struct {
	ctx context.Context
	ApiService *AccountCloningAPIService
	organizationId string
	request *AssetGroupAccountClone
}

func (r ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest) Request(request AssetGroupAccountClone) ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest {
	r.request = &request
	return r
}

func (r ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest) Execute() (*AssetGroupAccountClone, *http.Response, error) {
	return r.ApiService.OrganizationProvisionAssetGroupCloneAssetGroupAccountExecute(r)
}

/*
OrganizationProvisionAssetGroupCloneAssetGroupAccount Clone an existing DocuSign account.

Clones an existing DocuSign account. Cloning an account copies the plan items and eSignature settings. Users, templates, and permission profiles are _not_ copied into the target account. The new account will linked to the same organization.

In the request body, specify the source account you are cloning by its ID. You can get the IDs for your organization's asset group accounts using the [getAssetGroupAccounts](/docs/admin-api/reference/accountprovisioning/accountcloning/getassetgroupaccounts/) endpoint. 

You also need to specify information about the new target account, including the name, location, and administrator. To set the location, provide either the `countryCode` or `region` properties. (If both are specified, the `region` value will be used.)

The request body looks like this:
```
{
    "sourceAccount": {
        "id": "624e3e00-xxxx-xxxx-xxxx-43918c520dab"
    },
    "targetAccount": {
        "name": "My Cloned Account",
        "admin": {
            "firstName": "Francis",
            "lastName": "Beagle",
            "email": "francis@example.com"
        },
        "region": "NA"
    }
}
```

To call this endpoint:
* You must be an administrator for the specified organization.
* The source account must be on the same plan as your contract.
* The source account must be an asset group account. An asset group is a set of modules, products, plans, and charges purchased for your organization. An asset group account is an account that has been linked to an asset group.

[Required scopes](/docs/admin-api/admin101/auth/): `asset_group_account_clone_write`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization's GUID. 
 @return ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest
*/
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupCloneAssetGroupAccount(ctx context.Context, organizationId string) ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest {
	return ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return AssetGroupAccountClone
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupCloneAssetGroupAccountExecute(r ApiOrganizationProvisionAssetGroupCloneAssetGroupAccountRequest) (*AssetGroupAccountClone, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetGroupAccountClone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountCloningAPIService.OrganizationProvisionAssetGroupCloneAssetGroupAccount")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationId}/assetGroups/accountClone"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
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
	localVarPostBody = r.request
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

type ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest struct {
	ctx context.Context
	ApiService *AccountCloningAPIService
	organizationId string
	assetGroupId string
	assetGroupWorkId string
	includeDetails *bool
}

// When **true,** include additional details about the cloned account. The default value is **false.**
func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest) IncludeDetails(includeDetails bool) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest {
	r.includeDetails = &includeDetails
	return r
}

func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest) Execute() (*AssetGroupAccountClone, *http.Response, error) {
	return r.ApiService.OrganizationProvisionAssetGroupGetAssetGroupAccountCloneExecute(r)
}

/*
OrganizationProvisionAssetGroupGetAssetGroupAccountClone Gets information about a single cloned account.

Gets information about a cloned account by the `assetGroupWorkId`.

To call this endpoint:
* You must be an administrator for the specified organization.

[Required scopes](/docs/admin-api/admin101/auth/): `asset_group_account_clone_read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization's GUID. 
 @param assetGroupId The ID of the asset group.
 @param assetGroupWorkId The ID of the asset group account clone request.
 @return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest
*/
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountClone(ctx context.Context, organizationId string, assetGroupId string, assetGroupWorkId string) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest {
	return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
		assetGroupId: assetGroupId,
		assetGroupWorkId: assetGroupWorkId,
	}
}

// Execute executes the request
//  @return AssetGroupAccountClone
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountCloneExecute(r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountCloneRequest) (*AssetGroupAccountClone, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetGroupAccountClone
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountCloningAPIService.OrganizationProvisionAssetGroupGetAssetGroupAccountClone")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationId}/assetGroups/{assetGroupId}/accountClones/{assetGroupWorkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assetGroupId"+"}", url.PathEscape(parameterValueToString(r.assetGroupId, "assetGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"assetGroupWorkId"+"}", url.PathEscape(parameterValueToString(r.assetGroupWorkId, "assetGroupWorkId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_details", r.includeDetails, "")
	}
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

type ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest struct {
	ctx context.Context
	ApiService *AccountCloningAPIService
	organizationId string
	sinceUpdatedDate *time.Time
	includeDetails *bool
}

// Use this parameter to retrieve only account clones that were created on or after a specified date.
func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest) SinceUpdatedDate(sinceUpdatedDate time.Time) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest {
	r.sinceUpdatedDate = &sinceUpdatedDate
	return r
}

// When **true,** include additional details for the asset group account clones. The default value is **false.**
func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest) IncludeDetails(includeDetails bool) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest {
	r.includeDetails = &includeDetails
	return r
}

func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest) Execute() (*AssetGroupAccountClones, *http.Response, error) {
	return r.ApiService.OrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdExecute(r)
}

/*
OrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgId Gets all asset group account clones for an organization.

Retrieves all the cloned accounts for an organization.

To call this endpoint:
* You must be an administrator for the specified organization.

[Required scopes](/docs/admin-api/admin101/auth/): `asset_group_account_clone_read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization's GUID. 
 @return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest
*/
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgId(ctx context.Context, organizationId string) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest {
	return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return AssetGroupAccountClones
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdExecute(r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgIdRequest) (*AssetGroupAccountClones, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetGroupAccountClones
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountCloningAPIService.OrganizationProvisionAssetGroupGetAssetGroupAccountClonesByOrgId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationId}/assetGroups/accountClones"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sinceUpdatedDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since_updated_date", r.sinceUpdatedDate, "")
	}
	if r.includeDetails != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include_details", r.includeDetails, "")
	}
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

type ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest struct {
	ctx context.Context
	ApiService *AccountCloningAPIService
	organizationId string
	compliant *bool
}

// When **true,** only compliant accounts are returned and account responses do not include the &#x60;compliant&#x60; field. The default value is **false.**
func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest) Compliant(compliant bool) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest {
	r.compliant = &compliant
	return r
}

func (r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest) Execute() (*AssetGroupAccountsResponse, *http.Response, error) {
	return r.ApiService.OrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgExecute(r)
}

/*
OrganizationProvisionAssetGroupGetAssetGroupAccountsByOrg Get asset group accounts for an organization.

Returns a list of asset group accounts for an organization.

An asset group is a set of modules, products, plans, and charges purchased for your organization. An asset group account is an account that has been linked to an asset group.

To call this endpoint:
* You must be an administrator for the specified organization.

[Required authentication scopes](/docs/admin-api/admin101/auth/): `asset_group_account_read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId The organization's GUID. 
 @return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest
*/
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountsByOrg(ctx context.Context, organizationId string) ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest {
	return ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return AssetGroupAccountsResponse
func (a *AccountCloningAPIService) OrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgExecute(r ApiOrganizationProvisionAssetGroupGetAssetGroupAccountsByOrgRequest) (*AssetGroupAccountsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssetGroupAccountsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountCloningAPIService.OrganizationProvisionAssetGroupGetAssetGroupAccountsByOrg")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/organizations/{organizationId}/assetGroups/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.compliant != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "compliant", r.compliant, "")
	}
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