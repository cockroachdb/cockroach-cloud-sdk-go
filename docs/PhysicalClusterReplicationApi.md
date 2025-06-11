# PhysicalClusterReplication

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhysicalReplicationStream**](PhysicalClusterReplicationApi.md#CreatePhysicalReplicationStream) | **Post** /api/v1/physical-replication-streams | Create a physical replication stream
[**CreateReplicationStream**](PhysicalClusterReplicationApi.md#CreateReplicationStream) | **Post** /api/v1/replication-streams | Create a replication stream
[**GetPhysicalReplicationStream**](PhysicalClusterReplicationApi.md#GetPhysicalReplicationStream) | **Get** /api/v1/physical-replication-streams/{id} | Get a physical replication stream
[**GetReplicationStream**](PhysicalClusterReplicationApi.md#GetReplicationStream) | **Get** /api/v1/replication-streams/{id} | Get a replication stream
[**ListPhysicalReplicationStreams**](PhysicalClusterReplicationApi.md#ListPhysicalReplicationStreams) | **Get** /api/v1/physical-replication-streams | List physical replication streams
[**ListReplicationStreams**](PhysicalClusterReplicationApi.md#ListReplicationStreams) | **Get** /api/v1/replication-streams | List replication streams
[**UpdatePhysicalReplicationStream**](PhysicalClusterReplicationApi.md#UpdatePhysicalReplicationStream) | **Patch** /api/v1/physical-replication-streams/{id} | Update a physical replication stream
[**UpdateReplicationStream**](PhysicalClusterReplicationApi.md#UpdateReplicationStream) | **Patch** /api/v1/replication-streams/{id} | Update a replication stream



## CreatePhysicalReplicationStream

> PhysicalReplicationStream CreatePhysicalReplicationStream(ctx).CreatePhysicalReplicationStreamRequest(createPhysicalReplicationStreamRequest).Execute()

Create a physical replication stream

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
    createPhysicalReplicationStreamRequest := *openapiclient.NewCreatePhysicalReplicationStreamRequest("PrimaryClusterId_example", "StandbyClusterId_example") // CreatePhysicalReplicationStreamRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.CreatePhysicalReplicationStream(context.Background()).CreatePhysicalReplicationStreamRequest(createPhysicalReplicationStreamRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.CreatePhysicalReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePhysicalReplicationStream`: PhysicalReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.CreatePhysicalReplicationStream`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPhysicalReplicationStreamRequest** | [**CreatePhysicalReplicationStreamRequest**](CreatePhysicalReplicationStreamRequest.md) |  | 

### Return type

[**PhysicalReplicationStream**](PhysicalReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


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


## GetPhysicalReplicationStream

> PhysicalReplicationStream GetPhysicalReplicationStream(ctx, id).Execute()

Get a physical replication stream

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
    resp, r, err := api_client.PhysicalClusterReplicationApi.GetPhysicalReplicationStream(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.GetPhysicalReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhysicalReplicationStream`: PhysicalReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.GetPhysicalReplicationStream`: %v\n", resp)
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

[**PhysicalReplicationStream**](PhysicalReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
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


## ListPhysicalReplicationStreams

> PhysicalReplicationStreamList ListPhysicalReplicationStreams(ctx).PrimaryClusterId(primaryClusterId).StandbyClusterId(standbyClusterId).ClusterId(clusterId).ShowCompleted(showCompleted).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List physical replication streams

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
    "time"
    openapiclient "./openapi"
)

func main() {
    primaryClusterId := "primaryClusterId_example" // string | primary_cluster_id, if set, will cause only replication streams with this cluster as the primary to be returned. (optional)
    standbyClusterId := "standbyClusterId_example" // string | standby_cluster_id, if set, will cause only replication streams with this cluster as the standby to be returned. (optional)
    clusterId := "clusterId_example" // string | cluster_id, if set, will cause replication streams with this cluster as the primary or the standby to be returned. (optional)
    showCompleted := true // bool | show_completed specifies whether or not replication streams in the completed state are shown. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.ListPhysicalReplicationStreams(context.Background()).PrimaryClusterId(primaryClusterId).StandbyClusterId(standbyClusterId).ClusterId(clusterId).ShowCompleted(showCompleted).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.ListPhysicalReplicationStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPhysicalReplicationStreams`: PhysicalReplicationStreamList
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.ListPhysicalReplicationStreams`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListPhysicalReplicationStreamsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **primaryClusterId** | **string** | primary_cluster_id, if set, will cause only replication streams with this cluster as the primary to be returned. | 
 **standbyClusterId** | **string** | standby_cluster_id, if set, will cause only replication streams with this cluster as the standby to be returned. | 
 **clusterId** | **string** | cluster_id, if set, will cause replication streams with this cluster as the primary or the standby to be returned. | 
 **showCompleted** | **bool** | show_completed specifies whether or not replication streams in the completed state are shown. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**PhysicalReplicationStreamList**](PhysicalReplicationStreamList.md)

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
    "time"
    openapiclient "./openapi"
)

func main() {
    sourceClusterId := "sourceClusterId_example" // string | source_cluster_id, if set, will cause only replication streams with this cluster as the source to be returned. (optional)
    targetClusterId := "targetClusterId_example" // string | target_cluster_id, if set, will cause only replication streams with this cluster as the target to be returned. (optional)
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListReplicationStreamsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceClusterId** | **string** | source_cluster_id, if set, will cause only replication streams with this cluster as the source to be returned. | 
 **targetClusterId** | **string** | target_cluster_id, if set, will cause only replication streams with this cluster as the target to be returned. | 
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


## UpdatePhysicalReplicationStream

> PhysicalReplicationStream UpdatePhysicalReplicationStream(ctx, id).UpdatePhysicalReplicationStreamSpec(updatePhysicalReplicationStreamSpec).Execute()

Update a physical replication stream

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
    updatePhysicalReplicationStreamSpec := *openapiclient.NewUpdatePhysicalReplicationStreamSpec() // UpdatePhysicalReplicationStreamSpec | spec contains the information that is being updated for the given replication stream.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PhysicalClusterReplicationApi.UpdatePhysicalReplicationStream(context.Background(), id).UpdatePhysicalReplicationStreamSpec(updatePhysicalReplicationStreamSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PhysicalClusterReplicationApi.UpdatePhysicalReplicationStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePhysicalReplicationStream`: PhysicalReplicationStream
    fmt.Fprintf(os.Stdout, "Response from `PhysicalClusterReplicationApi.UpdatePhysicalReplicationStream`: %v\n", resp)
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

 **updatePhysicalReplicationStreamSpec** | [**UpdatePhysicalReplicationStreamSpec**](UpdatePhysicalReplicationStreamSpec.md) | spec contains the information that is being updated for the given replication stream. | 

### Return type

[**PhysicalReplicationStream**](PhysicalReplicationStream.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
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

