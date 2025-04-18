# ClusterDisruption

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterDisruptionInfo**](ClusterDisruptionApi.md#GetClusterDisruptionInfo) | **Get** /api/v1/clusters/{cluster_id}/disrupt | Get disruption specifications for a cluster
[**UpdateClusterDisruption**](ClusterDisruptionApi.md#UpdateClusterDisruption) | **Put** /api/v1/clusters/{cluster_id}/disrupt | Update disruption specifications for a cluster



## GetClusterDisruptionInfo

> ClusterDisruptionInfo GetClusterDisruptionInfo(ctx, clusterId).Execute()

Get disruption specifications for a cluster

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
    clusterId := "clusterId_example" // string | cluster_id is the cluster we are requesting disruption information for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClusterDisruptionApi.GetClusterDisruptionInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterDisruptionApi.GetClusterDisruptionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterDisruptionInfo`: ClusterDisruptionInfo
    fmt.Fprintf(os.Stdout, "Response from `ClusterDisruptionApi.GetClusterDisruptionInfo`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the cluster we are requesting disruption information for. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterDisruptionInfo**](ClusterDisruptionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateClusterDisruption

> ClusterDisruptionInfo UpdateClusterDisruption(ctx, clusterId).CockroachCloudUpdateClusterDisruptionRequest(cockroachCloudUpdateClusterDisruptionRequest).Execute()

Update disruption specifications for a cluster

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
    clusterId := "clusterId_example" // string | cluster_id specifies the cluster for this request.
    cockroachCloudUpdateClusterDisruptionRequest := *openapiclient.NewCockroachCloudUpdateClusterDisruptionRequest() // CockroachCloudUpdateClusterDisruptionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClusterDisruptionApi.UpdateClusterDisruption(context.Background(), clusterId).CockroachCloudUpdateClusterDisruptionRequest(cockroachCloudUpdateClusterDisruptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterDisruptionApi.UpdateClusterDisruption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClusterDisruption`: ClusterDisruptionInfo
    fmt.Fprintf(os.Stdout, "Response from `ClusterDisruptionApi.UpdateClusterDisruption`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id specifies the cluster for this request. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cockroachCloudUpdateClusterDisruptionRequest** | [**CockroachCloudUpdateClusterDisruptionRequest**](CockroachCloudUpdateClusterDisruptionRequest.md) |  | 

### Return type

[**ClusterDisruptionInfo**](ClusterDisruptionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

