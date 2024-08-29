# OpenIDConnectConfiguration

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOidcConfig**](OpenIDConnectConfigurationApi.md#CreateApiOidcConfig) | **Post** /api/v1/api_oidc | Create an API OIDC configuration (Deprecated)
[**DeleteApiOidcConfig**](OpenIDConnectConfigurationApi.md#DeleteApiOidcConfig) | **Delete** /api/v1/api_oidc/{id} | Delete an API OIDC configuration (Deprecated)
[**GetApiOidcConfig**](OpenIDConnectConfigurationApi.md#GetApiOidcConfig) | **Get** /api/v1/api_oidc/{id} | Get an API OIDC configuration (Deprecated)
[**ListApiOidcConfig**](OpenIDConnectConfigurationApi.md#ListApiOidcConfig) | **Get** /api/v1/api_oidc | List all API OIDC configurations (Deprecated)
[**UpdateApiOidcConfig**](OpenIDConnectConfigurationApi.md#UpdateApiOidcConfig) | **Put** /api/v1/api_oidc/{id} | Update an API OIDC configuration (Deprecated)



## CreateApiOidcConfig

> ApiOidcConfig CreateApiOidcConfig(ctx).CreateApiOidcConfigRequest(createApiOidcConfigRequest).Execute()

Create an API OIDC configuration (Deprecated)

This endpoint has been deprecated in favor of /jwt-issuers

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    createApiOidcConfigRequest := *openapiclient.NewCreateApiOidcConfigRequest("Audience_example", "Issuer_example", "Jwks_example") // CreateApiOidcConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.OpenIDConnectConfigurationApi.CreateApiOidcConfig(context.Background()).CreateApiOidcConfigRequest(createApiOidcConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectConfigurationApi.CreateApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectConfigurationApi.CreateApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiOidcConfigRequest** | [**CreateApiOidcConfigRequest**](CreateApiOidcConfigRequest.md) |  | 

### Return type

[**ApiOidcConfig**](ApiOidcConfig.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteApiOidcConfig

> ApiOidcConfig DeleteApiOidcConfig(ctx, id).Execute()

Delete an API OIDC configuration (Deprecated)

This endpoint has been deprecated in favor of /jwt-issuers

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.OpenIDConnectConfigurationApi.DeleteApiOidcConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectConfigurationApi.DeleteApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectConfigurationApi.DeleteApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiOidcConfig**](ApiOidcConfig.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetApiOidcConfig

> ApiOidcConfig GetApiOidcConfig(ctx, id).Execute()

Get an API OIDC configuration (Deprecated)

This endpoint has been deprecated in favor of /jwt-issuers

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.OpenIDConnectConfigurationApi.GetApiOidcConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectConfigurationApi.GetApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectConfigurationApi.GetApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiOidcConfig**](ApiOidcConfig.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListApiOidcConfig

> ListApiOidcConfigResponse ListApiOidcConfig(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all API OIDC configurations (Deprecated)

This endpoint has been deprecated in favor of /jwt-issuers

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.OpenIDConnectConfigurationApi.ListApiOidcConfig(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectConfigurationApi.ListApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiOidcConfig`: ListApiOidcConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectConfigurationApi.ListApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListApiOidcConfigOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListApiOidcConfigResponse**](ListApiOidcConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateApiOidcConfig

> ApiOidcConfig UpdateApiOidcConfig(ctx, id).ApiOidcConfig1(apiOidcConfig1).Execute()

Update an API OIDC configuration (Deprecated)

This endpoint has been deprecated in favor of /jwt-issuers

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    id := "id_example" // string | 
    apiOidcConfig1 := *openapiclient.NewApiOidcConfig1("Audience_example", "Issuer_example", "Jwks_example") // ApiOidcConfig1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.OpenIDConnectConfigurationApi.UpdateApiOidcConfig(context.Background(), id).ApiOidcConfig1(apiOidcConfig1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectConfigurationApi.UpdateApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectConfigurationApi.UpdateApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiOidcConfig1** | [**ApiOidcConfig1**](ApiOidcConfig1.md) |  | 

### Return type

[**ApiOidcConfig**](ApiOidcConfig.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

