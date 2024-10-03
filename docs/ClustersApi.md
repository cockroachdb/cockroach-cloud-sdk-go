# Clusters

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCluster**](ClustersApi.md#CreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster
[**GetConnectionString**](ClustersApi.md#GetConnectionString) | **Get** /api/v1/clusters/{cluster_id}/connection-string | Get a formatted generic connection string for a cluster
[**ListAvailableRegions**](ClustersApi.md#ListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes
[**ListClusterNodes**](ClustersApi.md#ListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster
[**ListClusters**](ClustersApi.md#ListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization
[**ListMajorClusterVersions**](ClustersApi.md#ListMajorClusterVersions) | **Get** /api/v1/cluster-versions | List available major cluster versions
[**UpdateCluster**](ClustersApi.md#UpdateCluster) | **Patch** /api/v1/clusters/{cluster_id} | Scale, edit or upgrade a cluster



## CreateCluster

> Cluster CreateCluster(ctx).CreateClusterRequest(createClusterRequest).Execute()

Create and initialize a new cluster

Can be used by the following roles assigned at the organization scope:
- CLUSTER_ADMIN
- CLUSTER_CREATOR


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
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.CloudProvider.Type("GCP"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.CreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCluster

> Cluster DeleteCluster(ctx, clusterId).Execute()

Delete a cluster and all of its data

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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.DeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.DeleteCluster`: %v\n", resp)
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


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCluster

> Cluster GetCluster(ctx, clusterId).Execute()

Get extended information about a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- FOLDER_ADMIN
- FOLDER_MOVER


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
    resp, r, err := api_client.ClustersApi.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetCluster`: %v\n", resp)
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


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetConnectionString

> GetConnectionStringResponse GetConnectionString(ctx, clusterId).Database(database).SqlUser(sqlUser).Os(os).Execute()

Get a formatted generic connection string for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- FOLDER_ADMIN
- FOLDER_MOVER


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
    database := "database_example" // string |  (optional) (default to "defaultdb")
    sqlUser := "sqlUser_example" // string |  (optional)
    os := "os_example" // string | os indicates the target operating system, used with formatting the default SSL certificate path. Required only for dedicated clusters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.GetConnectionString(context.Background(), clusterId).Database(database).SqlUser(sqlUser).Os(os).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.GetConnectionString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionString`: GetConnectionStringResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.GetConnectionString`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the GetConnectionStringOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **string** |  | [default to &quot;defaultdb&quot;]
 **sqlUser** | **string** |  | 
 **os** | **string** | os indicates the target operating system, used with formatting the default SSL certificate path. Required only for dedicated clusters. | 

### Return type

[**GetConnectionStringResponse**](GetConnectionStringResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListAvailableRegions

> ListAvailableRegionsResponse ListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List the regions available for new clusters and nodes

Sort order: Distance (based on client IP address)

This endpoint may be used by any member of the organization.

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
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider.  - AZURE: The Azure cloud provider. (optional)
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.ListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListAvailableRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableRegions`: ListAvailableRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListAvailableRegions`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListAvailableRegionsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider.  - AZURE: The Azure cloud provider. | 
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListAvailableRegionsResponse**](ListAvailableRegionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListClusterNodes

> ListClusterNodesResponse ListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List nodes for a cluster

Sort order: Region name, node name

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | 
    regionName := "regionName_example" // string | Optional filter to limit response to a single region. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterNodes`: ListClusterNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusterNodes`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListClusterNodesOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionName** | **string** | Optional filter to limit response to a single region. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListClusterNodesResponse**](ListClusterNodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListClusters

> ListClustersResponse ListClusters(ctx).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List clusters owned by an organization

Sort order: Cluster name

Returns all clusters that the user has read access over

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
    showInactive := true // bool | If `true`, show clusters that have been deleted or failed to initialize. Note that inactive clusters will only be included if the requesting user has organization-scoped cluster read permissions. (optional) (default to false)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.ListClusters(context.Background()).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListClusters`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListClustersOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInactive** | **bool** | If &#x60;true&#x60;, show clusters that have been deleted or failed to initialize. Note that inactive clusters will only be included if the requesting user has organization-scoped cluster read permissions. | [default to false]
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListMajorClusterVersions

> ListMajorClusterVersionsResponse ListMajorClusterVersions(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List available major cluster versions

Sort order: Version number descending

This endpoint may be used by any member of the organization.

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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.ListMajorClusterVersions(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ListMajorClusterVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMajorClusterVersions`: ListMajorClusterVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ListMajorClusterVersions`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListMajorClusterVersionsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListMajorClusterVersionsResponse**](ListMajorClusterVersionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx, clusterId).UpdateClusterSpecification(updateClusterSpecification).Execute()

Scale, edit or upgrade a cluster

In addition to adding nodes and changing cluster fields, the PATCH Cluster endpoint can be used to upgrade the cluster version. A cluster can be upgraded when its `upgrade_status` field is equal to `UPGRADE_AVAILABLE`. The `/api/v1/cluster-versions` endpoint can be used to enumerate versions which are valid to upgrade to. To begin the upgrade, PATCH the desired version into `cockroach_version`.  For example `{"cockroach_version": "v24.2"}`. Multi-node clusters will undergo a rolling upgrade and will remain available, but single-node clusters will be briefly unavailable while the upgrade takes place. Upgrades will be finalized automatically after 72 hours but can be manually finalized by sending a PATCH containing `{"upgrade_status": "FINALIZED"}` to this endpoint. Before the cluster is finalized, it can be rolled back by either sending a PATCH of the previous version via `cockroach_version` or sending a PATCH containing `{"upgrade_status": "ROLLBACK_RUNNING"}`. Version upgrade operations cannot be performed simultaneously with other update operations. Only one of `upgrade_status` or `cockroach_version` is allowed in the request.

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
    updateClusterSpecification := *openapiclient.NewUpdateClusterSpecification() // UpdateClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClustersApi.UpdateCluster(context.Background(), clusterId).UpdateClusterSpecification(updateClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.UpdateCluster`: %v\n", resp)
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

 **updateClusterSpecification** | [**UpdateClusterSpecification**](UpdateClusterSpecification.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

