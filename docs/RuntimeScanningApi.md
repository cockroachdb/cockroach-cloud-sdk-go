# RuntimeScanning

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableClusterRuntimeScanning**](RuntimeScanningApi.md#DisableClusterRuntimeScanning) | **Delete** /api/v1/clusters/{cluster_id}/runtime-scanning | Disable runtime scanning on a cluster.
[**EnableClusterRuntimeScanning**](RuntimeScanningApi.md#EnableClusterRuntimeScanning) | **Post** /api/v1/clusters/{cluster_id}/runtime-scanning | Enable runtime scanning on a cluster.
[**GetClusterRuntimeScanning**](RuntimeScanningApi.md#GetClusterRuntimeScanning) | **Get** /api/v1/clusters/{cluster_id}/runtime-scanning | Get the runtime scanning status of a cluster.



## DisableClusterRuntimeScanning

> RuntimeScanningInfo DisableClusterRuntimeScanning(ctx, clusterId).Execute()

Disable runtime scanning on a cluster.

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
    clusterId := "clusterId_example" // string | cluster_id is the ID of the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RuntimeScanningApi.DisableClusterRuntimeScanning(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeScanningApi.DisableClusterRuntimeScanning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableClusterRuntimeScanning`: RuntimeScanningInfo
    fmt.Fprintf(os.Stdout, "Response from `RuntimeScanningApi.DisableClusterRuntimeScanning`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID of the cluster. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuntimeScanningInfo**](RuntimeScanningInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## EnableClusterRuntimeScanning

> RuntimeScanningInfo EnableClusterRuntimeScanning(ctx, clusterId).EnableClusterRuntimeScanningBody(enableClusterRuntimeScanningBody).Execute()

Enable runtime scanning on a cluster.

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
    clusterId := "clusterId_example" // string | cluster_id is the ID of the cluster.
    enableClusterRuntimeScanningBody := *openapiclient.NewEnableClusterRuntimeScanningBody(openapiclient.RuntimeScanning.Type("NONE")) // EnableClusterRuntimeScanningBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RuntimeScanningApi.EnableClusterRuntimeScanning(context.Background(), clusterId).EnableClusterRuntimeScanningBody(enableClusterRuntimeScanningBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeScanningApi.EnableClusterRuntimeScanning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableClusterRuntimeScanning`: RuntimeScanningInfo
    fmt.Fprintf(os.Stdout, "Response from `RuntimeScanningApi.EnableClusterRuntimeScanning`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID of the cluster. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableClusterRuntimeScanningBody** | [**EnableClusterRuntimeScanningBody**](EnableClusterRuntimeScanningBody.md) |  | 

### Return type

[**RuntimeScanningInfo**](RuntimeScanningInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetClusterRuntimeScanning

> RuntimeScanningInfo GetClusterRuntimeScanning(ctx, clusterId).Execute()

Get the runtime scanning status of a cluster.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER


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
    clusterId := "clusterId_example" // string | cluster_id is the ID of the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RuntimeScanningApi.GetClusterRuntimeScanning(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RuntimeScanningApi.GetClusterRuntimeScanning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterRuntimeScanning`: RuntimeScanningInfo
    fmt.Fprintf(os.Stdout, "Response from `RuntimeScanningApi.GetClusterRuntimeScanning`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID of the cluster. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RuntimeScanningInfo**](RuntimeScanningInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

