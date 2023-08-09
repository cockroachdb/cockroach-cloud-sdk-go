# IPAllowlists

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistEntry**](IPAllowlistsApi.md#AddAllowlistEntry) | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist
[**AddAllowlistEntry2**](IPAllowlistsApi.md#AddAllowlistEntry2) | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist
[**DeleteAllowlistEntry**](IPAllowlistsApi.md#DeleteAllowlistEntry) | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry
[**ListAllowlistEntries**](IPAllowlistsApi.md#ListAllowlistEntries) | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster
[**UpdateAllowlistEntry**](IPAllowlistsApi.md#UpdateAllowlistEntry) | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry



## AddAllowlistEntry

> AllowlistEntry AddAllowlistEntry(ctx, clusterId).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist

Can be used by the following roles assigned at the organization or cluster scope:
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
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.IPAllowlistsApi.AddAllowlistEntry(context.Background(), clusterId).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistsApi.AddAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistsApi.AddAllowlistEntry`: %v\n", resp)
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
[[Back to README]](../README.md)


## AddAllowlistEntry2

> AllowlistEntry AddAllowlistEntry2(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()

Add a new CIDR address to the IP allowlist

Can be used by the following roles assigned at the organization or cluster scope:
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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.IPAllowlistsApi.AddAllowlistEntry2(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistsApi.AddAllowlistEntry2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry2`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistsApi.AddAllowlistEntry2`: %v\n", resp)
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
[[Back to README]](../README.md)


## DeleteAllowlistEntry

> AllowlistEntry DeleteAllowlistEntry(ctx, clusterId, cidrIp, cidrMask).Execute()

Delete an IP allowlist entry

Can be used by the following roles assigned at the organization or cluster scope:
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
    cidrIp := "cidrIp_example" // string | 
    cidrMask := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.IPAllowlistsApi.DeleteAllowlistEntry(context.Background(), clusterId, cidrIp, cidrMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistsApi.DeleteAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistsApi.DeleteAllowlistEntry`: %v\n", resp)
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
[[Back to README]](../README.md)


## ListAllowlistEntries

> ListAllowlistEntriesResponse ListAllowlistEntries(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

Get the IP allowlist and propagation status for a cluster

Sort order: CIDR address

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
    clusterId := "clusterId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.IPAllowlistsApi.ListAllowlistEntries(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistsApi.ListAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowlistEntries`: ListAllowlistEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistsApi.ListAllowlistEntries`: %v\n", resp)
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
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListAllowlistEntriesResponse**](ListAllowlistEntriesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateAllowlistEntry

> AllowlistEntry UpdateAllowlistEntry(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()

Update an IP allowlist entry

Can be used by the following roles assigned at the organization or cluster scope:
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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.IPAllowlistsApi.UpdateAllowlistEntry(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IPAllowlistsApi.UpdateAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `IPAllowlistsApi.UpdateAllowlistEntry`: %v\n", resp)
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

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

