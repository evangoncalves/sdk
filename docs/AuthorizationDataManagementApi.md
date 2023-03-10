# \AuthorizationDataManagementApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegisterSecret**](AuthorizationDataManagementApi.md#RegisterSecret) | **Post** /iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id}/secrets | Register Secrets on an Authorization for an Application
[**RemoveSecret**](AuthorizationDataManagementApi.md#RemoveSecret) | **Delete** /iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id}/secrets/{secret_id} | Remove a Secret from an Authorization for an Application



## RegisterSecret

> AuthorizationResponse RegisterSecret(ctx, applicationId, authorizationId).RegisterSecretRequest(registerSecretRequest).Execute()

Register Secrets on an Authorization for an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationId := int32(56) // int32 | Application ID generated by senhasegura
    authorizationId := int32(56) // int32 | Authorization ID generated by senhasegura
    registerSecretRequest := *openapiclient.NewRegisterSecretRequest(int32(123)) // RegisterSecretRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationDataManagementApi.RegisterSecret(context.Background(), applicationId, authorizationId).RegisterSecretRequest(registerSecretRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationDataManagementApi.RegisterSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSecret`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationDataManagementApi.RegisterSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 
**authorizationId** | **int32** | Authorization ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **registerSecretRequest** | [**RegisterSecretRequest**](RegisterSecretRequest.md) |  | 

### Return type

[**AuthorizationResponse**](AuthorizationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSecret

> Response RemoveSecret(ctx, secretId, applicationId, authorizationId).Execute()

Remove a Secret from an Authorization for an Application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    secretId := int32(56) // int32 | Secret ID generated by senhasegura
    applicationId := int32(56) // int32 | Application ID generated by senhasegura
    authorizationId := int32(56) // int32 | Authorization ID generated by senhasegura

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationDataManagementApi.RemoveSecret(context.Background(), secretId, applicationId, authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationDataManagementApi.RemoveSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveSecret`: Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationDataManagementApi.RemoveSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **int32** | Secret ID generated by senhasegura | 
**applicationId** | **int32** | Application ID generated by senhasegura | 
**authorizationId** | **int32** | Authorization ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Response**](Response.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

