# BlackoutWindows

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlackoutWindow**](BlackoutWindowsApi.md#CreateBlackoutWindow) | **Post** /api/v1/clusters/{cluster_id}/blackout-windows | Create a blackout window for a cluster
[**DeleteBlackoutWindow**](BlackoutWindowsApi.md#DeleteBlackoutWindow) | **Delete** /api/v1/clusters/{cluster_id}/blackout-windows/{blackout_window_id} | Delete a blackout window for a cluster
[**GetBlackoutWindow**](BlackoutWindowsApi.md#GetBlackoutWindow) | **Get** /api/v1/clusters/{cluster_id}/blackout-windows/{blackout_window_id} | Get a blackout window by its ID for a cluster
[**ListBlackoutWindows**](BlackoutWindowsApi.md#ListBlackoutWindows) | **Get** /api/v1/clusters/{cluster_id}/blackout-windows | List all blackout windows for a cluster
[**UpdateBlackoutWindow**](BlackoutWindowsApi.md#UpdateBlackoutWindow) | **Patch** /api/v1/clusters/{cluster_id}/blackout-windows/{blackout_window_id} | Update a blackout window for a cluster



## CreateBlackoutWindow

> BlackoutWindow CreateBlackoutWindow(ctx, clusterId).CreateBlackoutWindowRequest(createBlackoutWindowRequest).Execute()

Create a blackout window for a cluster

Blackout windows are supported for ADVANCED clusters only.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | 
    createBlackoutWindowRequest := *openapiclient.NewCreateBlackoutWindowRequest(time.Now(), time.Now()) // CreateBlackoutWindowRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BlackoutWindowsApi.CreateBlackoutWindow(context.Background(), clusterId).CreateBlackoutWindowRequest(createBlackoutWindowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlackoutWindowsApi.CreateBlackoutWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBlackoutWindow`: BlackoutWindow
    fmt.Fprintf(os.Stdout, "Response from `BlackoutWindowsApi.CreateBlackoutWindow`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBlackoutWindowRequest** | [**CreateBlackoutWindowRequest**](CreateBlackoutWindowRequest.md) |  | 

### Return type

[**BlackoutWindow**](BlackoutWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBlackoutWindow

> BlackoutWindow DeleteBlackoutWindow(ctx, clusterId, blackoutWindowId).Execute()

Delete a blackout window for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    blackoutWindowId := "blackoutWindowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BlackoutWindowsApi.DeleteBlackoutWindow(context.Background(), clusterId, blackoutWindowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlackoutWindowsApi.DeleteBlackoutWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBlackoutWindow`: BlackoutWindow
    fmt.Fprintf(os.Stdout, "Response from `BlackoutWindowsApi.DeleteBlackoutWindow`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**blackoutWindowId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlackoutWindow**](BlackoutWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBlackoutWindow

> BlackoutWindow GetBlackoutWindow(ctx, clusterId, blackoutWindowId).Execute()

Get a blackout window by its ID for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    blackoutWindowId := "blackoutWindowId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BlackoutWindowsApi.GetBlackoutWindow(context.Background(), clusterId, blackoutWindowId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlackoutWindowsApi.GetBlackoutWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBlackoutWindow`: BlackoutWindow
    fmt.Fprintf(os.Stdout, "Response from `BlackoutWindowsApi.GetBlackoutWindow`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**blackoutWindowId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlackoutWindow**](BlackoutWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBlackoutWindows

> ListBlackoutWindowsResponse ListBlackoutWindows(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all blackout windows for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BlackoutWindowsApi.ListBlackoutWindows(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlackoutWindowsApi.ListBlackoutWindows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBlackoutWindows`: ListBlackoutWindowsResponse
    fmt.Fprintf(os.Stdout, "Response from `BlackoutWindowsApi.ListBlackoutWindows`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListBlackoutWindowsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListBlackoutWindowsResponse**](ListBlackoutWindowsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBlackoutWindow

> BlackoutWindow UpdateBlackoutWindow(ctx, clusterId, blackoutWindowId).UpdateBlackoutWindowRequest(updateBlackoutWindowRequest).Execute()

Update a blackout window for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    blackoutWindowId := "blackoutWindowId_example" // string | 
    updateBlackoutWindowRequest := *openapiclient.NewUpdateBlackoutWindowRequest() // UpdateBlackoutWindowRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BlackoutWindowsApi.UpdateBlackoutWindow(context.Background(), clusterId, blackoutWindowId).UpdateBlackoutWindowRequest(updateBlackoutWindowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BlackoutWindowsApi.UpdateBlackoutWindow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBlackoutWindow`: BlackoutWindow
    fmt.Fprintf(os.Stdout, "Response from `BlackoutWindowsApi.UpdateBlackoutWindow`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**blackoutWindowId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBlackoutWindowRequest** | [**UpdateBlackoutWindowRequest**](UpdateBlackoutWindowRequest.md) |  | 

### Return type

[**BlackoutWindow**](BlackoutWindow.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

