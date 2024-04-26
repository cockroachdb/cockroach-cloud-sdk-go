# PhysicalClusterReplication

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateReplicationStream**](PhysicalClusterReplicationApi.md#CreateReplicationStream) | **Post** /api/v1/replication-streams | Create a replication stream
[**GetReplicationStream**](PhysicalClusterReplicationApi.md#GetReplicationStream) | **Get** /api/v1/replication-streams/{id} | Get a replication stream
[**ListReplicationStreams**](PhysicalClusterReplicationApi.md#ListReplicationStreams) | **Get** /api/v1/replication-streams | List replication streams
[**UpdateReplicationStream**](PhysicalClusterReplicationApi.md#UpdateReplicationStream) | **Patch** /api/v1/replication-streams/{id} | Update a replication stream



## CreateReplicationStream

> ReplicationStream CreateReplicationStream(ctx).CreateReplicationStreamRequest(createReplicationStreamRequest).Execute()

Create a replication stream

Can be used by the following roles assigned at the organization or cluster scope:
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



### Other Parameters

Other parameters are passed through a pointer to a apiCreateReplicationStream struct via the builder pattern


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

Can be used by the following roles assigned at the organization or cluster scope:
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

Other parameters are passed through a pointer to a apiGetReplicationStream struct via the builder pattern


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


## ListReplicationStreams

> ReplicationStreamList ListReplicationStreams(ctx).SourceClusterId(sourceClusterId).TargetClusterId(targetClusterId).ClusterId(clusterId).ShowCompleted(showCompleted).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List replication streams

Can be used by the following roles assigned at the organization or cluster scope:
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
    "time"
    openapiclient "./openapi"
)

func main() {
    sourceClusterId := "sourceClusterId_example" // string | source_cluster_id, if set, will cause only replication streams with this cluster as the source to be returned. (optional)
    targetClusterId := "targetClusterId_example" // string | source_cluster_id, if set, will cause only replication streams with this cluster as the target to be returned. (optional)
    clusterId := "clusterId_example" // string | cluster_id, if set, will cause replication streams with this cluster as the source or the target to be returned. (optional)
    showCompleted := true // bool | show_completed specifies whether or not replication streams in the completed state are shown. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.ListReplicationStreams(context.Background()).SourceClusterId(sourceClusterId).TargetClusterId(targetClusterId).ClusterId(clusterId).ShowCompleted(showCompleted).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.ListReplicationStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReplicationStreams`: ReplicationStreamList
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.ListReplicationStreams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReplicationStreams struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceClusterId** | **string** | source_cluster_id, if set, will cause only replication streams with this cluster as the source to be returned. | 
 **targetClusterId** | **string** | source_cluster_id, if set, will cause only replication streams with this cluster as the target to be returned. | 
 **clusterId** | **string** | cluster_id, if set, will cause replication streams with this cluster as the source or the target to be returned. | 
 **showCompleted** | **bool** | show_completed specifies whether or not replication streams in the completed state are shown. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ReplicationStreamList**](ReplicationStreamList.md)

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

Can be used by the following roles assigned at the organization or cluster scope:
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
    updateReplicationStreamSpec := *openapiclient.NewUpdateReplicationStreamSpec(false) // UpdateReplicationStreamSpec | 

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

Other parameters are passed through a pointer to a apiUpdateReplicationStream struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateReplicationStreamSpec** | [**UpdateReplicationStreamSpec**](UpdateReplicationStreamSpec.md) |  | 

### Return type

[**ReplicationStream**](ReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

