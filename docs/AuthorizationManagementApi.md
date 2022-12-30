# \AuthorizationManagementApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAuthorization**](AuthorizationManagementApi.md#CreateAuthorization) | **Post** /iso/dsm/v2/applications/{application_id}/authorizations | Create Authorizations
[**DeleteAuthorization**](AuthorizationManagementApi.md#DeleteAuthorization) | **Delete** /iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id} | Delete/Disable an Authorization for an Application
[**GetAuthorization**](AuthorizationManagementApi.md#GetAuthorization) | **Get** /iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id} | Get an Authorization from an Application
[**ListAuthorizations**](AuthorizationManagementApi.md#ListAuthorizations) | **Get** /iso/dsm/v2/applications/{application_id}/authorizations | List Authorizations
[**UpdateAuthorization**](AuthorizationManagementApi.md#UpdateAuthorization) | **Put** /iso/dsm/v2/applications/{application_id}/authorizations/{authorization_id} | Update an Authorization for an Application



## CreateAuthorization

> AuthorizationResponse CreateAuthorization(ctx, applicationId).AuthorizationRequest(authorizationRequest).Execute()

Create Authorizations



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
    authorizationRequest := *openapiclient.NewAuthorizationRequest("Monitoring", "Development", false, "Permissions_example") // AuthorizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationManagementApi.CreateAuthorization(context.Background(), applicationId).AuthorizationRequest(authorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationManagementApi.CreateAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAuthorization`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationManagementApi.CreateAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorizationRequest** | [**AuthorizationRequest**](AuthorizationRequest.md) |  | 

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


## DeleteAuthorization

> Response DeleteAuthorization(ctx, applicationId, authorizationId).Execute()

Delete/Disable an Authorization for an Application



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationManagementApi.DeleteAuthorization(context.Background(), applicationId, authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationManagementApi.DeleteAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthorization`: Response
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationManagementApi.DeleteAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 
**authorizationId** | **int32** | Authorization ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthorizationRequest struct via the builder pattern


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


## GetAuthorization

> AuthorizationResponse GetAuthorization(ctx, applicationId, authorizationId).Execute()

Get an Authorization from an Application



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationManagementApi.GetAuthorization(context.Background(), applicationId, authorizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationManagementApi.GetAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAuthorization`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationManagementApi.GetAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 
**authorizationId** | **int32** | Authorization ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AuthorizationResponse**](AuthorizationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuthorizations

> AuthorizationsResponse ListAuthorizations(ctx, applicationId).Execute()

List Authorizations



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationManagementApi.ListAuthorizations(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationManagementApi.ListAuthorizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthorizations`: AuthorizationsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationManagementApi.ListAuthorizations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthorizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AuthorizationsResponse**](AuthorizationsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAuthorization

> AuthorizationResponse UpdateAuthorization(ctx, applicationId, authorizationId).AuthorizationRequest(authorizationRequest).Execute()

Update an Authorization for an Application



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
    authorizationRequest := *openapiclient.NewAuthorizationRequest("Monitoring", "Development", false, "Permissions_example") // AuthorizationRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationManagementApi.UpdateAuthorization(context.Background(), applicationId, authorizationId).AuthorizationRequest(authorizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationManagementApi.UpdateAuthorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAuthorization`: AuthorizationResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationManagementApi.UpdateAuthorization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **int32** | Application ID generated by senhasegura | 
**authorizationId** | **int32** | Authorization ID generated by senhasegura | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAuthorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorizationRequest** | [**AuthorizationRequest**](AuthorizationRequest.md) |  | 

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
