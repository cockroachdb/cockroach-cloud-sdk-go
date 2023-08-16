# Default

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiOidcConfig**](DefaultApi.md#CreateApiOidcConfig) | **Post** /api/v1/api_oidc | Create an API OIDC configuration
[**DeleteApiOidcConfig**](DefaultApi.md#DeleteApiOidcConfig) | **Delete** /api/v1/api_oidc/{id} | Delete an API OIDC configuration
[**GetApiOidcConfig**](DefaultApi.md#GetApiOidcConfig) | **Get** /api/v1/api_oidc/{id} | Get an API OIDC configuration
[**GetClusterVersionDeferral**](DefaultApi.md#GetClusterVersionDeferral) | **Get** /api/v1/clusters/{cluster_id}/version-deferral | 
[**ListApiOidcConfig**](DefaultApi.md#ListApiOidcConfig) | **Get** /api/v1/api_oidc | List all API OIDC configurations
[**SetClusterVersionDeferral**](DefaultApi.md#SetClusterVersionDeferral) | **Put** /api/v1/clusters/{cluster_id}/version-deferral | 
[**UpdateApiOidcConfig**](DefaultApi.md#UpdateApiOidcConfig) | **Put** /api/v1/api_oidc/{id} | Update an API OIDC configuration



## CreateApiOidcConfig

> ApiOidcConfig CreateApiOidcConfig(ctx).CreateApiOidcConfigRequest(createApiOidcConfigRequest).Execute()

Create an API OIDC configuration

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
    resp, r, err := api_client.DefaultApi.CreateApiOidcConfig(context.Background()).CreateApiOidcConfigRequest(createApiOidcConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiOidcConfig struct via the builder pattern


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

Delete an API OIDC configuration

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
    resp, r, err := api_client.DefaultApi.DeleteApiOidcConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiOidcConfig struct via the builder pattern


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

Get an API OIDC configuration

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
    resp, r, err := api_client.DefaultApi.GetApiOidcConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiOidcConfig struct via the builder pattern


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


## GetClusterVersionDeferral

> ClusterVersionDeferral GetClusterVersionDeferral(ctx, clusterId).Execute()



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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DefaultApi.GetClusterVersionDeferral(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetClusterVersionDeferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterVersionDeferral`: ClusterVersionDeferral
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetClusterVersionDeferral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterVersionDeferral struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterVersionDeferral**](ClusterVersionDeferral.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListApiOidcConfig

> ListApiOidcConfigResponse ListApiOidcConfig(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all API OIDC configurations

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
    resp, r, err := api_client.DefaultApi.ListApiOidcConfig(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiOidcConfig`: ListApiOidcConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiOidcConfig struct via the builder pattern


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


## SetClusterVersionDeferral

> ClusterVersionDeferral SetClusterVersionDeferral(ctx, clusterId).ClusterVersionDeferral(clusterVersionDeferral).Execute()



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
    clusterId := "clusterId_example" // string | 
    clusterVersionDeferral := *openapiclient.NewClusterVersionDeferral(openapiclient.ClusterVersionDeferralPolicy.Type("NOT_DEFERRED")) // ClusterVersionDeferral | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DefaultApi.SetClusterVersionDeferral(context.Background(), clusterId).ClusterVersionDeferral(clusterVersionDeferral).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SetClusterVersionDeferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetClusterVersionDeferral`: ClusterVersionDeferral
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SetClusterVersionDeferral`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClusterVersionDeferral struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterVersionDeferral** | [**ClusterVersionDeferral**](ClusterVersionDeferral.md) |  | 

### Return type

[**ClusterVersionDeferral**](ClusterVersionDeferral.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateApiOidcConfig

> ApiOidcConfig UpdateApiOidcConfig(ctx, id).ApiOidcConfig1(apiOidcConfig1).Execute()

Update an API OIDC configuration

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
    resp, r, err := api_client.DefaultApi.UpdateApiOidcConfig(context.Background(), id).ApiOidcConfig1(apiOidcConfig1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateApiOidcConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiOidcConfig`: ApiOidcConfig
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateApiOidcConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiOidcConfig struct via the builder pattern


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

