# VersionDeferral

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterVersionDeferral**](VersionDeferralApi.md#GetClusterVersionDeferral) | **Get** /api/v1/clusters/{cluster_id}/version-deferral | Get the version upgrade deferral policy for a cluster.
[**SetClusterVersionDeferral**](VersionDeferralApi.md#SetClusterVersionDeferral) | **Put** /api/v1/clusters/{cluster_id}/version-deferral | Set the version upgrade deferral policy for a cluster



## GetClusterVersionDeferral

> ClusterVersionDeferral GetClusterVersionDeferral(ctx, clusterId).Execute()

Get the version upgrade deferral policy for a cluster.

Can be used by the following roles assigned at the organization or cluster scope:
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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.VersionDeferralApi.GetClusterVersionDeferral(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionDeferralApi.GetClusterVersionDeferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClusterVersionDeferral`: ClusterVersionDeferral
    fmt.Fprintf(os.Stdout, "Response from `VersionDeferralApi.GetClusterVersionDeferral`: %v\n", resp)
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


## SetClusterVersionDeferral

> ClusterVersionDeferral SetClusterVersionDeferral(ctx, clusterId).ClusterVersionDeferral(clusterVersionDeferral).Execute()

Set the version upgrade deferral policy for a cluster

Can be used by the following roles assigned at the organization or cluster scope:
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
    clusterVersionDeferral := *openapiclient.NewClusterVersionDeferral(openapiclient.ClusterVersionDeferralPolicy.Type("NOT_DEFERRED")) // ClusterVersionDeferral | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.VersionDeferralApi.SetClusterVersionDeferral(context.Background(), clusterId).ClusterVersionDeferral(clusterVersionDeferral).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VersionDeferralApi.SetClusterVersionDeferral``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetClusterVersionDeferral`: ClusterVersionDeferral
    fmt.Fprintf(os.Stdout, "Response from `VersionDeferralApi.SetClusterVersionDeferral`: %v\n", resp)
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

