# PrivateEndpointServices

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePrivateEndpointServices**](PrivateEndpointServicesApi.md#CreatePrivateEndpointServices) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Create all PrivateEndpointServices for a cluster
[**ListAwsEndpointConnections**](PrivateEndpointServicesApi.md#ListAwsEndpointConnections) | **Get** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections | List all AwsEndpointConnections for a cluster
[**ListPrivateEndpointServices**](PrivateEndpointServicesApi.md#ListPrivateEndpointServices) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | List all PrivateEndpointServices for a cluster
[**SetAwsEndpointConnectionState**](PrivateEndpointServicesApi.md#SetAwsEndpointConnectionState) | **Patch** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections/{endpoint_id} | Set the AWS Endpoint Connection state



## CreatePrivateEndpointServices

> PrivateEndpointServices CreatePrivateEndpointServices(ctx, clusterId).Execute()

Create all PrivateEndpointServices for a cluster



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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.CreatePrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreatePrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreatePrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointServices struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateEndpointServices**](PrivateEndpointServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListAwsEndpointConnections

> AwsEndpointConnections ListAwsEndpointConnections(ctx, clusterId).Execute()

List all AwsEndpointConnections for a cluster



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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListAwsEndpointConnections(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListAwsEndpointConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsEndpointConnections`: AwsEndpointConnections
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListAwsEndpointConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsEndpointConnections struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsEndpointConnections**](AwsEndpointConnections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPrivateEndpointServices

> PrivateEndpointServices ListPrivateEndpointServices(ctx, clusterId).Execute()

List all PrivateEndpointServices for a cluster



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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListPrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListPrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListPrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointServices struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateEndpointServices**](PrivateEndpointServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetAwsEndpointConnectionState

> AwsEndpointConnection SetAwsEndpointConnectionState(ctx, clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()

Set the AWS Endpoint Connection state



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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.
    endpointId := "endpointId_example" // string | endpoint_id is the ID for the VPC endpoint on the customer's side.
    setAwsEndpointConnectionStateRequest := *openapiclient.NewSetAwsEndpointConnectionStateRequest(openapiclient.SetAWSEndpointConnectionStatus.Type("AVAILABLE")) // SetAwsEndpointConnectionStateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.SetAwsEndpointConnectionState(context.Background(), clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.SetAwsEndpointConnectionState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAwsEndpointConnectionState`: AwsEndpointConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.SetAwsEndpointConnectionState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 
**endpointId** | **string** | endpoint_id is the ID for the VPC endpoint on the customer&#39;s side. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAwsEndpointConnectionState struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setAwsEndpointConnectionStateRequest** | [**SetAwsEndpointConnectionStateRequest**](SetAwsEndpointConnectionStateRequest.md) |  | 

### Return type

[**AwsEndpointConnection**](AwsEndpointConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

