# \CockroachCloudApi

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistEntry**](CockroachCloudApi.md#AddAllowlistEntry) | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist.
[**AddAllowlistEntry2**](CockroachCloudApi.md#AddAllowlistEntry2) | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist.
[**CreateCluster**](CockroachCloudApi.md#CreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster.
[**CreateSQLUser**](CockroachCloudApi.md#CreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user.
[**DeleteAllowlistEntry**](CockroachCloudApi.md#DeleteAllowlistEntry) | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry.
[**DeleteCluster**](CockroachCloudApi.md#DeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data.
[**DeleteSQLUser**](CockroachCloudApi.md#DeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user.
[**EnableCMEK**](CockroachCloudApi.md#EnableCMEK) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster.
[**GetCMEKClusterInfo**](CockroachCloudApi.md#GetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster.
[**GetCluster**](CockroachCloudApi.md#GetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster.
[**ListAllowlistEntries**](CockroachCloudApi.md#ListAllowlistEntries) | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster.
[**ListAvailableRegions**](CockroachCloudApi.md#ListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes.
[**ListClusterNodes**](CockroachCloudApi.md#ListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster.
[**ListClusters**](CockroachCloudApi.md#ListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization.
[**ListSQLUsers**](CockroachCloudApi.md#ListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster.
[**UpdateAllowlistEntry**](CockroachCloudApi.md#UpdateAllowlistEntry) | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry.
[**UpdateCMEKStatus**](CockroachCloudApi.md#UpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster.
[**UpdateCluster**](CockroachCloudApi.md#UpdateCluster) | **Patch** /api/v1/clusters/{cluster_id} | Scale or edit a cluster.
[**UpdateSQLUserPassword**](CockroachCloudApi.md#UpdateSQLUserPassword) | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Update a SQL user&#39;s password.



## AddAllowlistEntry

> AllowlistEntry AddAllowlistEntry(ctx, clusterId).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist.

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
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddAllowlistEntry(context.Background(), clusterId).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAllowlistEntry2

> AllowlistEntry AddAllowlistEntry2(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist.

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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddAllowlistEntry2(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddAllowlistEntry2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry2`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddAllowlistEntry2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAllowlistEntry2 struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCluster

> Cluster CreateCluster(ctx).CreateClusterRequest(createClusterRequest).Execute()

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
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.api.CloudProvider("CLOUD_PROVIDER_UNSPECIFIED"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCluster struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSQLUser

> SQLUser CreateSQLUser(ctx, clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()

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
    createSQLUserRequest := *openapiclient.NewCreateSQLUserRequest("Name_example", "Password_example") // CreateSQLUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateSQLUser(context.Background(), clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSQLUserRequest** | [**CreateSQLUserRequest**](CreateSQLUserRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistEntry

> AllowlistEntry DeleteAllowlistEntry(ctx, clusterId, cidrIp, cidrMask).Execute()

Delete an IP allowlist entry.

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
    cidrIp := "cidrIp_example" // string | 
    cidrMask := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteAllowlistEntry(context.Background(), clusterId, cidrIp, cidrMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**cidrIp** | **string** |  | 
**cidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> Cluster DeleteCluster(ctx, clusterId).Execute()

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
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCluster struct via the builder pattern


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


## DeleteSQLUser

> SQLUser DeleteSQLUser(ctx, clusterId, name).Execute()

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
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteSQLUser(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCMEK

> CMEKClusterInfo EnableCMEK(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

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
    cMEKClusterSpecification := *openapiclient.NewCMEKClusterSpecification([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // CMEKClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EnableCMEK(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableCMEK``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCMEK`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableCMEK`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCMEK struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cMEKClusterSpecification** | [**CMEKClusterSpecification**](CMEKClusterSpecification.md) |  | 

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


## GetCMEKClusterInfo

> CMEKClusterInfo GetCMEKClusterInfo(ctx, clusterId).Execute()

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
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetCMEKClusterInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetCMEKClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCMEKClusterInfo`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetCMEKClusterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCMEKClusterInfo struct via the builder pattern


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


## GetCluster

> Cluster GetCluster(ctx, clusterId).Execute()

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
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCluster struct via the builder pattern


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


## ListAllowlistEntries

> ListAllowlistEntriesResponse ListAllowlistEntries(ctx, clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

Get the IP allowlist and propagation status for a cluster.



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
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAllowlistEntries(context.Background(), clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowlistEntries`: ListAllowlistEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowlistEntries struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

### Return type

[**ListAllowlistEntriesResponse**](ListAllowlistEntriesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableRegions

> ListAvailableRegionsResponse ListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

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
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. (optional) (default to "CLOUD_PROVIDER_UNSPECIFIED")
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAvailableRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableRegions`: ListAvailableRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAvailableRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableRegions struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. | [default to &quot;CLOUD_PROVIDER_UNSPECIFIED&quot;]
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

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


## ListClusterNodes

> ListClusterNodesResponse ListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

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
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterNodes`: ListClusterNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterNodes struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionName** | **string** | Optional filter to limit response to a single region. | 
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

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


## ListClusters

> ListClustersResponse ListClusters(ctx).ShowInactive(showInactive).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

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
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusters(context.Background()).ShowInactive(showInactive).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusters struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInactive** | **bool** | If &#x60;true&#x60;, show clusters that have been deleted or failed to initialize. | [default to false]
 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

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


## ListSQLUsers

> ListSQLUsersResponse ListSQLUsers(ctx, clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()

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
    paginationStartKey := "paginationStartKey_example" // string |  (optional)
    paginationDirection := "paginationDirection_example" // string |  (optional) (default to "PAGE_DIRECTION_NEXT")
    paginationLimit := int32(56) // int32 |  (optional)
    paginationTime := time.Now() // time.Time |  (optional)
    paginationOrder := "paginationOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional) (default to "ASC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListSQLUsers(context.Background(), clusterId).PaginationStartKey(paginationStartKey).PaginationDirection(paginationDirection).PaginationLimit(paginationLimit).PaginationTime(paginationTime).PaginationOrder(paginationOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListSQLUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSQLUsers`: ListSQLUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListSQLUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSQLUsers struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationStartKey** | **string** |  | 
 **paginationDirection** | **string** |  | [default to &quot;PAGE_DIRECTION_NEXT&quot;]
 **paginationLimit** | **int32** |  | 
 **paginationTime** | **time.Time** |  | 
 **paginationOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | [default to &quot;ASC&quot;]

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


## UpdateAllowlistEntry

> AllowlistEntry UpdateAllowlistEntry(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).FieldMask(fieldMask).Execute()

Update an IP allowlist entry.

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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateAllowlistEntry(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry(allowlistEntry).FieldMask(fieldMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 
 **fieldMask** | **string** |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCMEKStatus

> CMEKClusterInfo UpdateCMEKStatus(ctx, clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()

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
    updateCMEKStatusRequest := *openapiclient.NewUpdateCMEKStatusRequest(openapiclient.CMEKCustomerAction("UNKNOWN_ACTION")) // UpdateCMEKStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateCMEKStatus(context.Background(), clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCMEKStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKStatus`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCMEKStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCMEKStatus struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCMEKStatusRequest** | [**UpdateCMEKStatusRequest**](UpdateCMEKStatusRequest.md) |  | 

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


## UpdateCluster

> Cluster UpdateCluster(ctx, clusterId).UpdateClusterSpecification(updateClusterSpecification).FieldMask(fieldMask).Execute()

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
    updateClusterSpecification := *openapiclient.NewUpdateClusterSpecification() // UpdateClusterSpecification | 
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateCluster(context.Background(), clusterId).UpdateClusterSpecification(updateClusterSpecification).FieldMask(fieldMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCluster struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterSpecification** | [**UpdateClusterSpecification**](UpdateClusterSpecification.md) |  | 
 **fieldMask** | **string** |  | 

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


## UpdateSQLUserPassword

> SQLUser UpdateSQLUserPassword(ctx, clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()

Update a SQL user's password.

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
    updateSQLUserPasswordRequest := *openapiclient.NewUpdateSQLUserPasswordRequest("Password_example") // UpdateSQLUserPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateSQLUserPassword(context.Background(), clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateSQLUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSQLUserPassword`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateSQLUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSQLUserPassword struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSQLUserPasswordRequest** | [**UpdateSQLUserPasswordRequest**](UpdateSQLUserPasswordRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

