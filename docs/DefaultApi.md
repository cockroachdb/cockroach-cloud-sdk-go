# Default

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterVersionDeferral**](DefaultApi.md#GetClusterVersionDeferral) | **Get** /api/v1/clusters/{cluster_id}/version-deferral | 
[**SetClusterVersionDeferral**](DefaultApi.md#SetClusterVersionDeferral) | **Put** /api/v1/clusters/{cluster_id}/version-deferral | 



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

