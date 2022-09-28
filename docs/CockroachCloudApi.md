# \CockroachCloudApi

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistEntry**](CockroachCloudApi.md#AddAllowlistEntry) | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist.
[**AddAllowlistEntry2**](CockroachCloudApi.md#AddAllowlistEntry2) | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist.
[**CreateCluster**](CockroachCloudApi.md#CreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster.
[**CreateDatabase**](CockroachCloudApi.md#CreateDatabase) | **Post** /api/v1/clusters/{cluster_id}/databases | Create a new database.
[**CreatePrivateEndpointServices**](CockroachCloudApi.md#CreatePrivateEndpointServices) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Creates all PrivateEndpointServices for a given cluster.
[**CreateSQLUser**](CockroachCloudApi.md#CreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user.
[**DeleteAllowlistEntry**](CockroachCloudApi.md#DeleteAllowlistEntry) | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry.
[**DeleteCluster**](CockroachCloudApi.md#DeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data.
[**DeleteDatabase**](CockroachCloudApi.md#DeleteDatabase) | **Delete** /api/v1/clusters/{cluster_id}/databases/{name} | Delete a database.
[**DeleteLogExport**](CockroachCloudApi.md#DeleteLogExport) | **Delete** /api/v1/clusters/{cluster_id}/logexport | 
[**DeleteSQLUser**](CockroachCloudApi.md#DeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user.
[**EditDatabase**](CockroachCloudApi.md#EditDatabase) | **Patch** /api/v1/clusters/{cluster_id}/databases | Update a database.
[**EnableCMEKSpec**](CockroachCloudApi.md#EnableCMEKSpec) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster.
[**EnableLogExport**](CockroachCloudApi.md#EnableLogExport) | **Post** /api/v1/clusters/{cluster_id}/logexport | 
[**GetCMEKClusterInfo**](CockroachCloudApi.md#GetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster.
[**GetCluster**](CockroachCloudApi.md#GetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster.
[**GetInvoice**](CockroachCloudApi.md#GetInvoice) | **Get** /api/v1/invoices/{invoice_id} | Gets a specific invoice for an organization.
[**GetLogExportInfo**](CockroachCloudApi.md#GetLogExportInfo) | **Get** /api/v1/clusters/{cluster_id}/logexport | 
[**ListAllowlistEntries**](CockroachCloudApi.md#ListAllowlistEntries) | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster.
[**ListAvailableRegions**](CockroachCloudApi.md#ListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes.
[**ListAwsEndpointConnections**](CockroachCloudApi.md#ListAwsEndpointConnections) | **Get** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections | Lists all AwsEndpointConnections for a given cluster.
[**ListClusterNodes**](CockroachCloudApi.md#ListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster.
[**ListClusters**](CockroachCloudApi.md#ListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization.
[**ListDatabases**](CockroachCloudApi.md#ListDatabases) | **Get** /api/v1/clusters/{cluster_id}/databases | List databases for a cluster.
[**ListInvoices**](CockroachCloudApi.md#ListInvoices) | **Get** /api/v1/invoices | Billing List invoices for a given organization.
[**ListPrivateEndpointServices**](CockroachCloudApi.md#ListPrivateEndpointServices) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Lists all PrivateEndpointServices for a given cluster.
[**ListSQLUsers**](CockroachCloudApi.md#ListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster.
[**SetAwsEndpointConnectionState**](CockroachCloudApi.md#SetAwsEndpointConnectionState) | **Patch** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections/{endpoint_id} | Sets the AWS Endpoint Connection state based on what is passed in the body: accepted or rejected. The \&quot;status\&quot; in the returned proto does not reflect the latest post-update status, but rather the status before the state is transitioned.
[**UpdateAllowlistEntry**](CockroachCloudApi.md#UpdateAllowlistEntry) | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry.
[**UpdateCMEKSpec**](CockroachCloudApi.md#UpdateCMEKSpec) | **Put** /api/v1/clusters/{cluster_id}/cmek | Enable or update the CMEK spec for a cluster.
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

> AllowlistEntry AddAllowlistEntry2(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()

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
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddAllowlistEntry2(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()
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



 **allowlistEntry1** | [**AllowlistEntry1**](AllowlistEntry1.md) | AllowlistEntry | 

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
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.api.CloudProvider("GCP"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

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


## CreateDatabase

> ApiDatabase CreateDatabase(ctx, clusterId).CreateDatabaseRequest(createDatabaseRequest).Execute()

Create a new database.

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
    createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest("Name_example") // CreateDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateDatabase(context.Background(), clusterId).CreateDatabaseRequest(createDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 

### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateEndpointServices

> PrivateEndpointServices CreatePrivateEndpointServices(ctx, clusterId).Body(body).Execute()

Creates all PrivateEndpointServices for a given cluster.

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
    clusterId := "clusterId_example" // string | ClusterID is the ID for the cluster.
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreatePrivateEndpointServices(context.Background(), clusterId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreatePrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreatePrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointServices struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

[**PrivateEndpointServices**](PrivateEndpointServices.md)

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


## DeleteDatabase

> ApiDatabase DeleteDatabase(ctx, clusterId, name).Execute()

Delete a database.

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
    resp, r, err := api_client.CockroachCloudApi.DeleteDatabase(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogExport

> LogExportClusterInfo DeleteLogExport(ctx, clusterId).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.DeleteLogExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

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


## EditDatabase

> ApiDatabase EditDatabase(ctx, clusterId).UpdateDatabaseRequest(updateDatabaseRequest).Execute()

Update a database.

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
    updateDatabaseRequest := *openapiclient.NewUpdateDatabaseRequest("Name_example", "NewName_example") // UpdateDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EditDatabase(context.Background(), clusterId).UpdateDatabaseRequest(updateDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EditDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EditDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseRequest** | [**UpdateDatabaseRequest**](UpdateDatabaseRequest.md) |  | 

### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCMEKSpec

> CMEKClusterInfo EnableCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

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
    resp, r, err := api_client.CockroachCloudApi.EnableCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableCMEKSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCMEKSpec struct via the builder pattern


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


## EnableLogExport

> LogExportClusterInfo EnableLogExport(ctx, clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()



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
    enableLogExportRequest := *openapiclient.NewEnableLogExportRequest() // EnableLogExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EnableLogExport(context.Background(), clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableLogExportRequest** | [**EnableLogExportRequest**](EnableLogExportRequest.md) |  | 

### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

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


## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Gets a specific invoice for an organization.

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
    invoiceId := "invoiceId_example" // string | InvoiceID is the unique ID representing the invoice. InvoiceID is used to retrieve a specific billing period's invoice.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | InvoiceID is the unique ID representing the invoice. InvoiceID is used to retrieve a specific billing period&#39;s invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoice struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogExportInfo

> LogExportClusterInfo GetLogExportInfo(ctx, clusterId).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.GetLogExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetLogExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogExportInfo`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetLogExportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogExportInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllowlistEntries

> ListAllowlistEntriesResponse ListAllowlistEntries(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAllowlistEntries(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
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

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

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

> ListAvailableRegionsResponse ListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

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
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. (optional)
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
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
 **provider** | **string** | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider. | 
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

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


## ListAwsEndpointConnections

> AwsEndpointConnections ListAwsEndpointConnections(ctx, clusterId).Execute()

Lists all AwsEndpointConnections for a given cluster.

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
    clusterId := "clusterId_example" // string | ClusterID is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAwsEndpointConnections(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAwsEndpointConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsEndpointConnections`: AwsEndpointConnections
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAwsEndpointConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the ID for the cluster. | 

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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterNodes

> ListClusterNodesResponse ListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
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
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

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

> ListClustersResponse ListClusters(ctx).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusters(context.Background()).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
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
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

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


## ListDatabases

> ApiListDatabasesResponse ListDatabases(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List databases for a cluster.



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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListDatabases(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: ApiListDatabasesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabases struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

### Return type

[**ApiListDatabasesResponse**](ApiListDatabasesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> ListInvoicesResponse ListInvoices(ctx).Execute()

Billing List invoices for a given organization.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoices struct via the builder pattern


### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateEndpointServices

> PrivateEndpointServices ListPrivateEndpointServices(ctx, clusterId).Execute()

Lists all PrivateEndpointServices for a given cluster.

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
    clusterId := "clusterId_example" // string | ClusterID is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListPrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListPrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListPrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the ID for the cluster. | 

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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSQLUsers

> ListSQLUsersResponse ListSQLUsers(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - DESC: Sort in descending order. The default order is ascending. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListSQLUsers(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
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

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - DESC: Sort in descending order. The default order is ascending. | 

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


## SetAwsEndpointConnectionState

> AwsEndpointConnection SetAwsEndpointConnectionState(ctx, clusterId, endpointId).CockroachCloudSetAwsEndpointConnectionStateRequest(cockroachCloudSetAwsEndpointConnectionStateRequest).Execute()

Sets the AWS Endpoint Connection state based on what is passed in the body: accepted or rejected. The \"status\" in the returned proto does not reflect the latest post-update status, but rather the status before the state is transitioned.

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
    clusterId := "clusterId_example" // string | ClusterID is the ID for the cluster.
    endpointId := "endpointId_example" // string | EndpointID is the ID for the VPC endpoint on the customer's side.
    cockroachCloudSetAwsEndpointConnectionStateRequest := *openapiclient.NewCockroachCloudSetAwsEndpointConnectionStateRequest() // CockroachCloudSetAwsEndpointConnectionStateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.SetAwsEndpointConnectionState(context.Background(), clusterId, endpointId).CockroachCloudSetAwsEndpointConnectionStateRequest(cockroachCloudSetAwsEndpointConnectionStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.SetAwsEndpointConnectionState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAwsEndpointConnectionState`: AwsEndpointConnection
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.SetAwsEndpointConnectionState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | ClusterID is the ID for the cluster. | 
**endpointId** | **string** | EndpointID is the ID for the VPC endpoint on the customer&#39;s side. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAwsEndpointConnectionState struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cockroachCloudSetAwsEndpointConnectionStateRequest** | [**CockroachCloudSetAwsEndpointConnectionStateRequest**](CockroachCloudSetAwsEndpointConnectionStateRequest.md) |  | 

### Return type

[**AwsEndpointConnection**](AwsEndpointConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllowlistEntry

> AllowlistEntry UpdateAllowlistEntry(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).FieldMask(fieldMask).Execute()

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
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry
    fieldMask := "fieldMask_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateAllowlistEntry(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).FieldMask(fieldMask).Execute()
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



 **allowlistEntry1** | [**AllowlistEntry1**](AllowlistEntry1.md) | AllowlistEntry | 
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


## UpdateCMEKSpec

> CMEKClusterInfo UpdateCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable or update the CMEK spec for a cluster.

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
    resp, r, err := api_client.CockroachCloudApi.UpdateCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCMEKSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCMEKSpec struct via the builder pattern


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
    updateCMEKStatusRequest := *openapiclient.NewUpdateCMEKStatusRequest(openapiclient.CMEKCustomerAction("REVOKE")) // UpdateCMEKStatusRequest | 

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

