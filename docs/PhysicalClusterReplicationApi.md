# PhysicalClusterReplication

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationStream**](PhysicalClusterReplicationApi.md#CreateReplicationStream) | **Post** /api/v1/replication-streams | Create a replication stream
[**GetReplicationStream**](PhysicalClusterReplicationApi.md#GetReplicationStream) | **Get** /api/v1/replication-streams/{id} | Get a replication stream
[**UpdateReplicationStream**](PhysicalClusterReplicationApi.md#UpdateReplicationStream) | **Patch** /api/v1/replication-streams/{id} | Update a replication stream



## CreateReplicationStream

> ReplicationStream CreateReplicationStream(ctx).CreateReplicationStreamRequest(createReplicationStreamRequest).Execute()

Create a replication stream

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN


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
    createReplicationStreamRequest := *openapiclient.NewCreateReplicationStreamRequest("SourceClusterId_example", "TargetClusterId_example") // CreateReplicationStreamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.CreateReplicationStream(context.Background()).CreateReplicationStreamRequest(createReplicationStreamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.CreateReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReplicationStream`: ReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.CreateReplicationStream`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createReplicationStreamRequest** | [**CreateReplicationStreamRequest**](CreateReplicationStreamRequest.md) |  | 

### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetReplicationStream

> ReplicationStream GetReplicationStream(ctx, id).Execute()

Get a replication stream

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
    id := "id_example" // string | id is the ID of the replication stream to get.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.GetReplicationStream(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.GetReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationStream`: ReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.GetReplicationStream`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id is the ID of the replication stream to get. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateReplicationStream

> ReplicationStream UpdateReplicationStream(ctx, id).UpdateReplicationStreamSpec(updateReplicationStreamSpec).Execute()

Update a replication stream

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN


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
    id := "id_example" // string | id is the ID of the replication stream to update.
    updateReplicationStreamSpec := *openapiclient.NewUpdateReplicationStreamSpec() // UpdateReplicationStreamSpec | spec contains the information that is being updated for the given replication stream.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.UpdateReplicationStream(context.Background(), id).UpdateReplicationStreamSpec(updateReplicationStreamSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.UpdateReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReplicationStream`: ReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.UpdateReplicationStream`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id is the ID of the replication stream to update. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReplicationStreamSpec** | [**UpdateReplicationStreamSpec**](UpdateReplicationStreamSpec.md) | spec contains the information that is being updated for the given replication stream. | 

### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

