# LogExport

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLogExport**](LogExportApi.md#DeleteLogExport) | **Delete** /api/v1/clusters/{cluster_id}/logexport | Delete the Log Export configuration for a cluster
[**EnableLogExport**](LogExportApi.md#EnableLogExport) | **Post** /api/v1/clusters/{cluster_id}/logexport | Create or update the Log Export configuration for a cluster
[**GetLogExportInfo**](LogExportApi.md#GetLogExportInfo) | **Get** /api/v1/clusters/{cluster_id}/logexport | Get the Log Export configuration for a cluster



## DeleteLogExport

> LogExportClusterInfo DeleteLogExport(ctx, clusterId).Execute()

Delete the Log Export configuration for a cluster

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
    resp, r, err := api_client.LogExportApi.DeleteLogExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogExportApi.DeleteLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `LogExportApi.DeleteLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## EnableLogExport

> LogExportClusterInfo EnableLogExport(ctx, clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()

Create or update the Log Export configuration for a cluster

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
    enableLogExportRequest := *openapiclient.NewEnableLogExportRequest("AuthPrincipal_example", "LogName_example", openapiclient.LogExportType("AWS_CLOUDWATCH")) // EnableLogExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.LogExportApi.EnableLogExport(context.Background(), clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogExportApi.EnableLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `LogExportApi.EnableLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableLogExportRequest** | [**EnableLogExportRequest**](EnableLogExportRequest.md) |  | 

### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLogExportInfo

> LogExportClusterInfo GetLogExportInfo(ctx, clusterId).Execute()

Get the Log Export configuration for a cluster

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
    resp, r, err := api_client.LogExportApi.GetLogExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogExportApi.GetLogExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogExportInfo`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `LogExportApi.GetLogExportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogExportInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

