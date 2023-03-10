/*
senhasegura DSM API Reference

This is a senhasegura new API design

API version: 2.0.22
Contact: egoncalves@senhasegura.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AuthorizationManagementApiService AuthorizationManagementApi service
type AuthorizationManagementApiService service

type ApiCreateAuthorizationRequest struct {
	ctx context.Context
	ApiService *AuthorizationManagementApiService
	applicationId int32
	authorizationRequest *AuthorizationRequest
}

func (r ApiCreateAuthorizationRequest) AuthorizationRequest(authorizationRequest AuthorizationRequest) ApiCreateAuthorizationRequest {
	r.authorizationRequest = &authorizationRequest
	return r
}

func (r ApiCreateAuthorizationRequest) Execute() (*AuthorizationResponse, *http.Response, error) {
	return r.ApiService.CreateAuthorizationExecute(r)
}

/*
CreateAuthorization Create Authorizations

Create an authorization for an application on senhasegura DSM. Accessing applications/authorizations not related to the requesting authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID generated by senhasegura
 @return ApiCreateAuthorizationRequest
*/
func (a *AuthorizationManagementApiService) CreateAuthorization(ctx context.Context, applicationId int32) ApiCreateAuthorizationRequest {
	return ApiCreateAuthorizationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return AuthorizationResponse
func (a *AuthorizationManagementApiService) CreateAuthorizationExecute(r ApiCreateAuthorizationRequest) (*AuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationManagementApiService.CreateAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/applications/{application_id}/authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.authorizationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeleteAuthorizationRequest struct {
	ctx context.Context
	ApiService *AuthorizationManagementApiService
	applicationId int32
	authorizationId int32
}

func (r ApiDeleteAuthorizationRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.DeleteAuthorizationExecute(r)
}

/*
DeleteAuthorization Delete/Disable an Authorization for an Application

Delete/Disable an authorization from an application on senhasegura DSM. Accessing applications not related to the requesting authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID generated by senhasegura
 @param authorizationId Authorization ID generated by senhasegura
 @return ApiDeleteAuthorizationRequest
*/
func (a *AuthorizationManagementApiService) DeleteAuthorization(ctx context.Context, applicationId int32, authorizationId int32) ApiDeleteAuthorizationRequest {
	return ApiDeleteAuthorizationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
//  @return Response
func (a *AuthorizationManagementApiService) DeleteAuthorizationExecute(r ApiDeleteAuthorizationRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationManagementApiService.DeleteAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorization_id"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGetAuthorizationRequest struct {
	ctx context.Context
	ApiService *AuthorizationManagementApiService
	applicationId int32
	authorizationId int32
}

func (r ApiGetAuthorizationRequest) Execute() (*AuthorizationResponse, *http.Response, error) {
	return r.ApiService.GetAuthorizationExecute(r)
}

/*
GetAuthorization Get an Authorization from an Application

Return details of an authorization from an application on senhasegura DSM. Accessing applications/authorizations not related to the requesting authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID generated by senhasegura
 @param authorizationId Authorization ID generated by senhasegura
 @return ApiGetAuthorizationRequest
*/
func (a *AuthorizationManagementApiService) GetAuthorization(ctx context.Context, applicationId int32, authorizationId int32) ApiGetAuthorizationRequest {
	return ApiGetAuthorizationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
//  @return AuthorizationResponse
func (a *AuthorizationManagementApiService) GetAuthorizationExecute(r ApiGetAuthorizationRequest) (*AuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationManagementApiService.GetAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorization_id"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiListAuthorizationsRequest struct {
	ctx context.Context
	ApiService *AuthorizationManagementApiService
	applicationId int32
}

func (r ApiListAuthorizationsRequest) Execute() (*AuthorizationsResponse, *http.Response, error) {
	return r.ApiService.ListAuthorizationsExecute(r)
}

/*
ListAuthorizations List Authorizations

Return a list of all authorizations related to an application on senhasegura DSM. Accessing applications not related to the requesting authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID generated by senhasegura
 @return ApiListAuthorizationsRequest
*/
func (a *AuthorizationManagementApiService) ListAuthorizations(ctx context.Context, applicationId int32) ApiListAuthorizationsRequest {
	return ApiListAuthorizationsRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return AuthorizationsResponse
func (a *AuthorizationManagementApiService) ListAuthorizationsExecute(r ApiListAuthorizationsRequest) (*AuthorizationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationManagementApiService.ListAuthorizations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/applications/{application_id}/authorizations"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiUpdateAuthorizationRequest struct {
	ctx context.Context
	ApiService *AuthorizationManagementApiService
	applicationId int32
	authorizationId int32
	authorizationRequest *AuthorizationRequest
}

func (r ApiUpdateAuthorizationRequest) AuthorizationRequest(authorizationRequest AuthorizationRequest) ApiUpdateAuthorizationRequest {
	r.authorizationRequest = &authorizationRequest
	return r
}

func (r ApiUpdateAuthorizationRequest) Execute() (*AuthorizationResponse, *http.Response, error) {
	return r.ApiService.UpdateAuthorizationExecute(r)
}

/*
UpdateAuthorization Update an Authorization for an Application

Update an authorization from an application on senhasegura DSM. Accessing applications not related to the requesting authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Application ID generated by senhasegura
 @param authorizationId Authorization ID generated by senhasegura
 @return ApiUpdateAuthorizationRequest
*/
func (a *AuthorizationManagementApiService) UpdateAuthorization(ctx context.Context, applicationId int32, authorizationId int32) ApiUpdateAuthorizationRequest {
	return ApiUpdateAuthorizationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		authorizationId: authorizationId,
	}
}

// Execute executes the request
//  @return AuthorizationResponse
func (a *AuthorizationManagementApiService) UpdateAuthorizationExecute(r ApiUpdateAuthorizationRequest) (*AuthorizationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationManagementApiService.UpdateAuthorization")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"application_id"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"authorization_id"+"}", url.PathEscape(parameterValueToString(r.authorizationId, "authorizationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.authorizationRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
