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


// SecretDataManagementApiService SecretDataManagementApi service
type SecretDataManagementApiService service

type ApiRegisterSecretCloudCredentialRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	secretCloudCredentialRequest *SecretCloudCredentialRequest
}

func (r ApiRegisterSecretCloudCredentialRequest) SecretCloudCredentialRequest(secretCloudCredentialRequest SecretCloudCredentialRequest) ApiRegisterSecretCloudCredentialRequest {
	r.secretCloudCredentialRequest = &secretCloudCredentialRequest
	return r
}

func (r ApiRegisterSecretCloudCredentialRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.RegisterSecretCloudCredentialExecute(r)
}

/*
RegisterSecretCloudCredential Register a Cloud Credential on a Secret

Register a cloud credential on a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @return ApiRegisterSecretCloudCredentialRequest
*/
func (a *SecretDataManagementApiService) RegisterSecretCloudCredential(ctx context.Context, secretId int32) ApiRegisterSecretCloudCredentialRequest {
	return ApiRegisterSecretCloudCredentialRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *SecretDataManagementApiService) RegisterSecretCloudCredentialExecute(r ApiRegisterSecretCloudCredentialRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RegisterSecretCloudCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/credentials/cloud"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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
	localVarPostBody = r.secretCloudCredentialRequest
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

type ApiRegisterSecretKVRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	requestBody *map[string]string
}

func (r ApiRegisterSecretKVRequest) RequestBody(requestBody map[string]string) ApiRegisterSecretKVRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiRegisterSecretKVRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.RegisterSecretKVExecute(r)
}

/*
RegisterSecretKV Create a Key/Value Pair on a Secret

Create a key/value pair on a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @return ApiRegisterSecretKVRequest
*/
func (a *SecretDataManagementApiService) RegisterSecretKV(ctx context.Context, secretId int32) ApiRegisterSecretKVRequest {
	return ApiRegisterSecretKVRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *SecretDataManagementApiService) RegisterSecretKVExecute(r ApiRegisterSecretKVRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RegisterSecretKV")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/kv"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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
	localVarPostBody = r.requestBody
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

type ApiRegisterSecretPAMCredentialRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	secretPAMCredentialRequest *SecretPAMCredentialRequest
}

func (r ApiRegisterSecretPAMCredentialRequest) SecretPAMCredentialRequest(secretPAMCredentialRequest SecretPAMCredentialRequest) ApiRegisterSecretPAMCredentialRequest {
	r.secretPAMCredentialRequest = &secretPAMCredentialRequest
	return r
}

func (r ApiRegisterSecretPAMCredentialRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.RegisterSecretPAMCredentialExecute(r)
}

/*
RegisterSecretPAMCredential Register PAM Credentials on a Secret

Register a credential on a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @return ApiRegisterSecretPAMCredentialRequest
*/
func (a *SecretDataManagementApiService) RegisterSecretPAMCredential(ctx context.Context, secretId int32) ApiRegisterSecretPAMCredentialRequest {
	return ApiRegisterSecretPAMCredentialRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *SecretDataManagementApiService) RegisterSecretPAMCredentialExecute(r ApiRegisterSecretPAMCredentialRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RegisterSecretPAMCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/credentials/pam"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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
	localVarPostBody = r.secretPAMCredentialRequest
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

type ApiRegisterSecretSSHKeyRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	secretPAMSSHKeyRequest *SecretPAMSSHKeyRequest
}

func (r ApiRegisterSecretSSHKeyRequest) SecretPAMSSHKeyRequest(secretPAMSSHKeyRequest SecretPAMSSHKeyRequest) ApiRegisterSecretSSHKeyRequest {
	r.secretPAMSSHKeyRequest = &secretPAMSSHKeyRequest
	return r
}

func (r ApiRegisterSecretSSHKeyRequest) Execute() (*SecretResponse, *http.Response, error) {
	return r.ApiService.RegisterSecretSSHKeyExecute(r)
}

/*
RegisterSecretSSHKey Register SSH Keys on a Secret

Register a credential on a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @return ApiRegisterSecretSSHKeyRequest
*/
func (a *SecretDataManagementApiService) RegisterSecretSSHKey(ctx context.Context, secretId int32) ApiRegisterSecretSSHKeyRequest {
	return ApiRegisterSecretSSHKeyRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
	}
}

// Execute executes the request
//  @return SecretResponse
func (a *SecretDataManagementApiService) RegisterSecretSSHKeyExecute(r ApiRegisterSecretSSHKeyRequest) (*SecretResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecretResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RegisterSecretSSHKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/keys"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)

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
	localVarPostBody = r.secretPAMSSHKeyRequest
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

type ApiRemoveSecretCloudCredentialRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	cloudCredentialId int32
}

func (r ApiRemoveSecretCloudCredentialRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveSecretCloudCredentialExecute(r)
}

/*
RemoveSecretCloudCredential Remove a Cloud Credential from a Secret

Delete/Disable a cloud credential from a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @param cloudCredentialId Cloud Credential ID generated by senhasegura
 @return ApiRemoveSecretCloudCredentialRequest
*/
func (a *SecretDataManagementApiService) RemoveSecretCloudCredential(ctx context.Context, secretId int32, cloudCredentialId int32) ApiRemoveSecretCloudCredentialRequest {
	return ApiRemoveSecretCloudCredentialRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
		cloudCredentialId: cloudCredentialId,
	}
}

// Execute executes the request
//  @return Response
func (a *SecretDataManagementApiService) RemoveSecretCloudCredentialExecute(r ApiRemoveSecretCloudCredentialRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RemoveSecretCloudCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/credentials/cloud/{cloud_credential_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloud_credential_id"+"}", url.PathEscape(parameterValueToString(r.cloudCredentialId, "cloudCredentialId")), -1)

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

type ApiRemoveSecretKVRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	keyIdentifier int32
}

func (r ApiRemoveSecretKVRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveSecretKVExecute(r)
}

/*
RemoveSecretKV Remove a Key/Value Pair on a Secret

Delete a key/value pair from a secret. Accessing secrets not related to this application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @param keyIdentifier Key Identifier registered on senhasegura
 @return ApiRemoveSecretKVRequest
*/
func (a *SecretDataManagementApiService) RemoveSecretKV(ctx context.Context, secretId int32, keyIdentifier int32) ApiRemoveSecretKVRequest {
	return ApiRemoveSecretKVRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
		keyIdentifier: keyIdentifier,
	}
}

// Execute executes the request
//  @return Response
func (a *SecretDataManagementApiService) RemoveSecretKVExecute(r ApiRemoveSecretKVRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RemoveSecretKV")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/kv/{key_identifier}"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_identifier"+"}", url.PathEscape(parameterValueToString(r.keyIdentifier, "keyIdentifier")), -1)

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

type ApiRemoveSecretPAMCredentialRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	pamCredentialId int32
}

func (r ApiRemoveSecretPAMCredentialRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveSecretPAMCredentialExecute(r)
}

/*
RemoveSecretPAMCredential Remove a PAM Credential from a Secret

Delete a credential from a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @param pamCredentialId Credential ID generated by senhasegura
 @return ApiRemoveSecretPAMCredentialRequest
*/
func (a *SecretDataManagementApiService) RemoveSecretPAMCredential(ctx context.Context, secretId int32, pamCredentialId int32) ApiRemoveSecretPAMCredentialRequest {
	return ApiRemoveSecretPAMCredentialRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
		pamCredentialId: pamCredentialId,
	}
}

// Execute executes the request
//  @return Response
func (a *SecretDataManagementApiService) RemoveSecretPAMCredentialExecute(r ApiRemoveSecretPAMCredentialRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RemoveSecretPAMCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/credentials/pam/{pam_credential_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pam_credential_id"+"}", url.PathEscape(parameterValueToString(r.pamCredentialId, "pamCredentialId")), -1)

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

type ApiRemoveSecretSSHKeyRequest struct {
	ctx context.Context
	ApiService *SecretDataManagementApiService
	secretId int32
	keyId string
}

func (r ApiRemoveSecretSSHKeyRequest) Execute() (*Response, *http.Response, error) {
	return r.ApiService.RemoveSecretSSHKeyExecute(r)
}

/*
RemoveSecretSSHKey Remove a SSH Key from a Secret

Delete a credential from a secret. Accessing secrets not related to the requesting application/authorization requires system-wide permission.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param secretId Secret ID generated by senhasegura
 @param keyId
 @return ApiRemoveSecretSSHKeyRequest
*/
func (a *SecretDataManagementApiService) RemoveSecretSSHKey(ctx context.Context, secretId int32, keyId string) ApiRemoveSecretSSHKeyRequest {
	return ApiRemoveSecretSSHKeyRequest{
		ApiService: a,
		ctx: ctx,
		secretId: secretId,
		keyId: keyId,
	}
}

// Execute executes the request
//  @return Response
func (a *SecretDataManagementApiService) RemoveSecretSSHKeyExecute(r ApiRemoveSecretSSHKeyRequest) (*Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SecretDataManagementApiService.RemoveSecretSSHKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/iso/dsm/v2/secrets/{secret_id}/keys/{key_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"secret_id"+"}", url.PathEscape(parameterValueToString(r.secretId, "secretId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_id"+"}", url.PathEscape(parameterValueToString(r.keyId, "keyId")), -1)

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