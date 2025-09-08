# EgressPrivateEndpoints

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEgressPrivateEndpoint**](EgressPrivateEndpointsApi.md#CreateEgressPrivateEndpoint) | **Post** /api/v1/clusters/{cluster_id}/networking/egress-private-endpoints | Create an egress private endpoint
[**DeleteEgressPrivateEndpoint**](EgressPrivateEndpointsApi.md#DeleteEgressPrivateEndpoint) | **Delete** /api/v1/clusters/{cluster_id}/networking/egress-private-endpoints/{id} | Delete an egress private endpoint
[**GetEgressPrivateEndpoint**](EgressPrivateEndpointsApi.md#GetEgressPrivateEndpoint) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-private-endpoints/{id} | Get egress private endpoint
[**ListEgressPrivateEndpoints**](EgressPrivateEndpointsApi.md#ListEgressPrivateEndpoints) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-private-endpoints | List egress private endpoints
[**UpdateEgressPrivateEndpointDomainNames**](EgressPrivateEndpointsApi.md#UpdateEgressPrivateEndpointDomainNames) | **Patch** /api/v1/clusters/{cluster_id}/networking/egress-private-endpoints/{id}/domain-names | Update egress private endpoint domain names



## CreateEgressPrivateEndpoint

> EgressPrivateEndpoint CreateEgressPrivateEndpoint(ctx, clusterId).CreateEgressPrivateEndpointRequest(createEgressPrivateEndpointRequest).Execute()

Create an egress private endpoint

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
    clusterId := "clusterId_example" // string | cluster_id identifies the cluster to which this egress private endpoint applies.
    createEgressPrivateEndpointRequest := *openapiclient.NewCreateEgressPrivateEndpointRequest("Region_example", "TargetServiceIdentifier_example", openapiclient.EgressPrivateEndpointTargetServiceType.Type("PRIVATE_SERVICE")) // CreateEgressPrivateEndpointRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressPrivateEndpointsApi.CreateEgressPrivateEndpoint(context.Background(), clusterId).CreateEgressPrivateEndpointRequest(createEgressPrivateEndpointRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressPrivateEndpointsApi.CreateEgressPrivateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEgressPrivateEndpoint`: EgressPrivateEndpoint
    fmt.Fprintf(os.Stdout, "Response from `EgressPrivateEndpointsApi.CreateEgressPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the cluster to which this egress private endpoint applies. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEgressPrivateEndpointRequest** | [**CreateEgressPrivateEndpointRequest**](CreateEgressPrivateEndpointRequest.md) |  | 

### Return type

[**EgressPrivateEndpoint**](EgressPrivateEndpoint.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteEgressPrivateEndpoint

> map[string]interface{} DeleteEgressPrivateEndpoint(ctx, clusterId, id).Execute()

Delete an egress private endpoint

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
    clusterId := "clusterId_example" // string | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint.
    id := "id_example" // string | id is the UUID value of the egress private endpoint in CockroachDB Cloud.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressPrivateEndpointsApi.DeleteEgressPrivateEndpoint(context.Background(), clusterId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressPrivateEndpointsApi.DeleteEgressPrivateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEgressPrivateEndpoint`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EgressPrivateEndpointsApi.DeleteEgressPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint. | 
**id** | **string** | id is the UUID value of the egress private endpoint in CockroachDB Cloud. | 

### Other Parameters


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
[[Back to README]](../README.md)


## GetEgressPrivateEndpoint

> GetEgressPrivateEndpointResponse GetEgressPrivateEndpoint(ctx, clusterId, id).Execute()

Get egress private endpoint

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
    clusterId := "clusterId_example" // string | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint.
    id := "id_example" // string | id is the UUID value of the egress private endpoint in CockroachDB Cloud.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressPrivateEndpointsApi.GetEgressPrivateEndpoint(context.Background(), clusterId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressPrivateEndpointsApi.GetEgressPrivateEndpoint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEgressPrivateEndpoint`: GetEgressPrivateEndpointResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressPrivateEndpointsApi.GetEgressPrivateEndpoint`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint. | 
**id** | **string** | id is the UUID value of the egress private endpoint in CockroachDB Cloud. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEgressPrivateEndpointResponse**](GetEgressPrivateEndpointResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEgressPrivateEndpoints

> ListEgressPrivateEndpointsResponse ListEgressPrivateEndpoints(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List egress private endpoints

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
    "time"
    openapiclient "./openapi"
)

func main() {
    clusterId := "clusterId_example" // string | cluster_id identifies the CockroachDB Cloud cluster whose egress private endpoints to list.
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressPrivateEndpointsApi.ListEgressPrivateEndpoints(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressPrivateEndpointsApi.ListEgressPrivateEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEgressPrivateEndpoints`: ListEgressPrivateEndpointsResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressPrivateEndpointsApi.ListEgressPrivateEndpoints`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the CockroachDB Cloud cluster whose egress private endpoints to list. | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListEgressPrivateEndpointsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListEgressPrivateEndpointsResponse**](ListEgressPrivateEndpointsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateEgressPrivateEndpointDomainNames

> map[string]interface{} UpdateEgressPrivateEndpointDomainNames(ctx, clusterId, id).UpdateEgressPrivateEndpointDomainNamesRequest(updateEgressPrivateEndpointDomainNamesRequest).Execute()

Update egress private endpoint domain names

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
    clusterId := "clusterId_example" // string | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint.
    id := "id_example" // string | id is the UUID value of the egress private endpoint in CockroachDB Cloud.
    updateEgressPrivateEndpointDomainNamesRequest := *openapiclient.NewUpdateEgressPrivateEndpointDomainNamesRequest([]string{"DomainNames_example"}) // UpdateEgressPrivateEndpointDomainNamesRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressPrivateEndpointsApi.UpdateEgressPrivateEndpointDomainNames(context.Background(), clusterId, id).UpdateEgressPrivateEndpointDomainNamesRequest(updateEgressPrivateEndpointDomainNamesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressPrivateEndpointsApi.UpdateEgressPrivateEndpointDomainNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEgressPrivateEndpointDomainNames`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EgressPrivateEndpointsApi.UpdateEgressPrivateEndpointDomainNames`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the CockroachDB Cloud cluster owning the egress private endpoint. | 
**id** | **string** | id is the UUID value of the egress private endpoint in CockroachDB Cloud. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateEgressPrivateEndpointDomainNamesRequest** | [**UpdateEgressPrivateEndpointDomainNamesRequest**](UpdateEgressPrivateEndpointDomainNamesRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

