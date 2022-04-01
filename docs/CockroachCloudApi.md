# \CockroachCloudApi

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CockroachCloudCreateCluster**](CockroachCloudApi.md#CockroachCloudCreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster.
[**CockroachCloudCreateSQLUser**](CockroachCloudApi.md#CockroachCloudCreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user.
[**CockroachCloudDeleteCluster**](CockroachCloudApi.md#CockroachCloudDeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data.
[**CockroachCloudDeleteSQLUser**](CockroachCloudApi.md#CockroachCloudDeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user.
[**CockroachCloudEnableCMEK**](CockroachCloudApi.md#CockroachCloudEnableCMEK) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster.
[**CockroachCloudGetCMEKClusterInfo**](CockroachCloudApi.md#CockroachCloudGetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster.
[**CockroachCloudGetCluster**](CockroachCloudApi.md#CockroachCloudGetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster.
[**CockroachCloudListAvailableRegions**](CockroachCloudApi.md#CockroachCloudListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes.
[**CockroachCloudListClusterNodes**](CockroachCloudApi.md#CockroachCloudListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster.
[**CockroachCloudListClusters**](CockroachCloudApi.md#CockroachCloudListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization.
[**CockroachCloudListSQLUsers**](CockroachCloudApi.md#CockroachCloudListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster.
[**CockroachCloudSetNodeCount**](CockroachCloudApi.md#CockroachCloudSetNodeCount) | **Put** /api/v1/clusters/{cluster_id}/nodes | Set the number of nodes in each region for a dedicated cluster.
[**CockroachCloudUpdateCMEKStatus**](CockroachCloudApi.md#CockroachCloudUpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster.
[**CockroachCloudUpdateCluster**](CockroachCloudApi.md#CockroachCloudUpdateCluster) | **Patch** /api/v1/clusters/{cluster_id} | Scale or edit a cluster.
[**CockroachCloudUpdateSQLUserPassword**](CockroachCloudApi.md#CockroachCloudUpdateSQLUserPassword) | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Updates a SQL user&#39;s password.



## CockroachCloudCreateCluster

> Cluster CockroachCloudCreateCluster(ctx).Body(body).Execute()

Create and initialize a new cluster.

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
    body := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.CloudProvider("CLOUD_PROVIDER_UNSPECIFIED"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudCreateCluster(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudCreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudCreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudCreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudCreateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudCreateSQLUser

> map[string]interface{} CockroachCloudCreateSQLUser(ctx, clusterId).Body(body).Execute()

Create a new SQL user.

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
    body := *openapiclient.NewCreateSQLUserRequest(*openapiclient.NewSQLUser("Name_example"), "Password_example") // CreateSQLUserRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudCreateSQLUser(context.Background(), clusterId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudCreateSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudCreateSQLUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudCreateSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudCreateSQLUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSQLUserRequest**](CreateSQLUserRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudDeleteCluster

> map[string]interface{} CockroachCloudDeleteCluster(ctx, clusterId).Execute()

Delete a cluster and all of its data.

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudDeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudDeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudDeleteCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudDeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudDeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudDeleteSQLUser

> map[string]interface{} CockroachCloudDeleteSQLUser(ctx, clusterId, name).Execute()

Delete a SQL user.

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudDeleteSQLUser(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudDeleteSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudDeleteSQLUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudDeleteSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudDeleteSQLUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudEnableCMEK

> CMEKClusterInfo CockroachCloudEnableCMEK(ctx, clusterId).Body(body).Execute()

Enable CMEK for a cluster.

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
    body := *openapiclient.NewCMEKClusterSpecification([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // CMEKClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudEnableCMEK(context.Background(), clusterId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudEnableCMEK``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudEnableCMEK`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudEnableCMEK`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudEnableCMEKRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CMEKClusterSpecification**](CMEKClusterSpecification.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudGetCMEKClusterInfo

> CMEKClusterInfo CockroachCloudGetCMEKClusterInfo(ctx, clusterId).Execute()

Get CMEK-related information for a cluster.

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudGetCMEKClusterInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudGetCMEKClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudGetCMEKClusterInfo`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudGetCMEKClusterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudGetCMEKClusterInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudGetCluster

> Cluster CockroachCloudGetCluster(ctx, clusterId).Execute()

Get extended information about a cluster.

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
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudGetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudGetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudGetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudGetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudGetClusterRequest struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListAvailableRegions

> ListAvailableRegionsResponse CockroachCloudListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List the regions available for new clusters and nodes.

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
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - CLOUD_PROVIDER_GCP: The Google Cloud Platform cloud provider.  - CLOUD_PROVIDER_AWS: The Amazon Web Services cloud provider. (optional) (default to "CLOUD_PROVIDER_UNSPECIFIED")
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationStart := int32(56) // int32 | Index of the first entry to return. (optional)
    paginationLimit := int32(56) // int32 | Number of entries to return. (optional)
    paginationTime := time.Now() // time.Time | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an `AS OF SYSTEM TIME` clause. Currently non-functional. (optional)
    paginationOrder := "paginationOrder_example" // string | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListAvailableRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListAvailableRegions`: ListAvailableRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListAvailableRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListAvailableRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Optional CloudProvider for filtering.   - CLOUD_PROVIDER_GCP: The Google Cloud Platform cloud provider.  - CLOUD_PROVIDER_AWS: The Amazon Web Services cloud provider. | [default to &quot;CLOUD_PROVIDER_UNSPECIFIED&quot;]
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationStart** | **int32** | Index of the first entry to return. | 
 **paginationLimit** | **int32** | Number of entries to return. | 
 **paginationTime** | **time.Time** | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an &#x60;AS OF SYSTEM TIME&#x60; clause. Currently non-functional. | 
 **paginationOrder** | **string** | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListAvailableRegionsResponse**](ListAvailableRegionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListClusterNodes

> ListClusterNodesResponse CockroachCloudListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List nodes for a cluster.

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
    paginationStart := int32(56) // int32 | Index of the first entry to return. (optional)
    paginationLimit := int32(56) // int32 | Number of entries to return. (optional)
    paginationTime := time.Now() // time.Time | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an `AS OF SYSTEM TIME` clause. Currently non-functional. (optional)
    paginationOrder := "paginationOrder_example" // string | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListClusterNodes`: ListClusterNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListClusterNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionName** | **string** | Optional filter to limit response to a single region. | 
 **paginationStart** | **int32** | Index of the first entry to return. | 
 **paginationLimit** | **int32** | Number of entries to return. | 
 **paginationTime** | **time.Time** | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an &#x60;AS OF SYSTEM TIME&#x60; clause. Currently non-functional. | 
 **paginationOrder** | **string** | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListClusterNodesResponse**](ListClusterNodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListClusters

> ListClustersResponse CockroachCloudListClusters(ctx).ShowInactive(showInactive).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List clusters owned by an organization.

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
    showInactive := true // bool | If `true`, show clusters that have been deleted or failed to initialize. (optional) (default to false)
    paginationStart := int32(56) // int32 | Index of the first entry to return. (optional)
    paginationLimit := int32(56) // int32 | Number of entries to return. (optional)
    paginationTime := time.Now() // time.Time | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an `AS OF SYSTEM TIME` clause. Currently non-functional. (optional)
    paginationOrder := "paginationOrder_example" // string | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListClusters(context.Background()).ShowInactive(showInactive).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListClusters`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInactive** | **bool** | If &#x60;true&#x60;, show clusters that have been deleted or failed to initialize. | [default to false]
 **paginationStart** | **int32** | Index of the first entry to return. | 
 **paginationLimit** | **int32** | Number of entries to return. | 
 **paginationTime** | **time.Time** | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an &#x60;AS OF SYSTEM TIME&#x60; clause. Currently non-functional. | 
 **paginationOrder** | **string** | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudListSQLUsers

> ListSQLUsersResponse CockroachCloudListSQLUsers(ctx, clusterId).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

List SQL users for a cluster.

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
    paginationStart := int32(56) // int32 | Index of the first entry to return. (optional)
    paginationLimit := int32(56) // int32 | Number of entries to return. (optional)
    paginationTime := time.Now() // time.Time | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an `AS OF SYSTEM TIME` clause. Currently non-functional. (optional)
    paginationOrder := "paginationOrder_example" // string | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudListSQLUsers(context.Background(), clusterId).PaginationStart(paginationStart).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudListSQLUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudListSQLUsers`: ListSQLUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudListSQLUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudListSQLUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationStart** | **int32** | Index of the first entry to return. | 
 **paginationLimit** | **int32** | Number of entries to return. | 
 **paginationTime** | **time.Time** | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an &#x60;AS OF SYSTEM TIME&#x60; clause. Currently non-functional. | 
 **paginationOrder** | **string** | ASC (ascending) or DESC (descending) sort order.   - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListSQLUsersResponse**](ListSQLUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudSetNodeCount

> Cluster CockroachCloudSetNodeCount(ctx, clusterId).Body(body).Execute()

Set the number of nodes in each region for a dedicated cluster.

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
    body := *openapiclient.NewSetNodeCountRequest(map[string]int32{"key": int32(123)}) // SetNodeCountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudSetNodeCount(context.Background(), clusterId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudSetNodeCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudSetNodeCount`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudSetNodeCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudSetNodeCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SetNodeCountRequest**](SetNodeCountRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateCMEKStatus

> CMEKClusterInfo CockroachCloudUpdateCMEKStatus(ctx, clusterId).Body(body).Execute()

Update the CMEK-related status for a cluster.

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
    body := *openapiclient.NewUpdateCMEKStatusRequest(openapiclient.CMEKCustomerAction("UNKNOWN_ACTION")) // UpdateCMEKStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateCMEKStatus(context.Background(), clusterId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateCMEKStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateCMEKStatus`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateCMEKStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateCMEKStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateCMEKStatusRequest**](UpdateCMEKStatusRequest.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateCluster

> map[string]interface{} CockroachCloudUpdateCluster(ctx, clusterId).Body(body).FieldMask(fieldMask).Execute()

Scale or edit a cluster.

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
    body := *openapiclient.NewUpdateClusterSpecification() // UpdateClusterSpecification | 
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateCluster(context.Background(), clusterId).Body(body).FieldMask(fieldMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateCluster`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateClusterSpecification**](UpdateClusterSpecification.md) |  | 
 **fieldMask** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CockroachCloudUpdateSQLUserPassword

> map[string]interface{} CockroachCloudUpdateSQLUserPassword(ctx, clusterId, name).Body(body).Execute()

Updates a SQL user's password.

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
    name := "name_example" // string | 
    body := *openapiclient.NewUpdateSQLUserPasswordRequest("Password_example") // UpdateSQLUserPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CockroachCloudApi.CockroachCloudUpdateSQLUserPassword(context.Background(), clusterId, name).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CockroachCloudUpdateSQLUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CockroachCloudUpdateSQLUserPassword`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CockroachCloudUpdateSQLUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCockroachCloudUpdateSQLUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateSQLUserPasswordRequest**](UpdateSQLUserPasswordRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

