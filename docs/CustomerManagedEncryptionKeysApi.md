# CustomerManagedEncryptionKeys

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableCMEKSpec**](CustomerManagedEncryptionKeysApi.md#EnableCMEKSpec) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster
[**GetCMEKClusterInfo**](CustomerManagedEncryptionKeysApi.md#GetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster
[**UpdateCMEKSpec**](CustomerManagedEncryptionKeysApi.md#UpdateCMEKSpec) | **Put** /api/v1/clusters/{cluster_id}/cmek | Enable or update the CMEK spec for a cluster
[**UpdateCMEKStatus**](CustomerManagedEncryptionKeysApi.md#UpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster



## EnableCMEKSpec

> CMEKClusterInfo EnableCMEKSpec(ctx, clusterId).EnableCMEKSpecBody(enableCMEKSpecBody).Execute()

Enable CMEK for a cluster

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
    enableCMEKSpecBody := *openapiclient.NewEnableCMEKSpecBody([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // EnableCMEKSpecBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.EnableCMEKSpec(context.Background(), clusterId).EnableCMEKSpecBody(enableCMEKSpecBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagedEncryptionKeysApi.EnableCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagedEncryptionKeysApi.EnableCMEKSpec`: %v\n", resp)
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

 **enableCMEKSpecBody** | [**EnableCMEKSpecBody**](EnableCMEKSpecBody.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCMEKClusterInfo

> CMEKClusterInfo GetCMEKClusterInfo(ctx, clusterId).Execute()

Get CMEK-related information for a cluster

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.GetCMEKClusterInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagedEncryptionKeysApi.GetCMEKClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCMEKClusterInfo`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagedEncryptionKeysApi.GetCMEKClusterInfo`: %v\n", resp)
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

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCMEKSpec

> CMEKClusterInfo UpdateCMEKSpec(ctx, clusterId).UpdateCMEKSpecBody(updateCMEKSpecBody).Execute()

Enable or update the CMEK spec for a cluster

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
    updateCMEKSpecBody := *openapiclient.NewUpdateCMEKSpecBody([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // UpdateCMEKSpecBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.UpdateCMEKSpec(context.Background(), clusterId).UpdateCMEKSpecBody(updateCMEKSpecBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagedEncryptionKeysApi.UpdateCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagedEncryptionKeysApi.UpdateCMEKSpec`: %v\n", resp)
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

 **updateCMEKSpecBody** | [**UpdateCMEKSpecBody**](UpdateCMEKSpecBody.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCMEKStatus

> CMEKClusterInfo UpdateCMEKStatus(ctx, clusterId).UpdateCMEKStatusBody(updateCMEKStatusBody).Execute()

Update the CMEK-related status for a cluster

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
    updateCMEKStatusBody := *openapiclient.NewUpdateCMEKStatusBody(openapiclient.CMEKCustomerAction("REVOKE")) // UpdateCMEKStatusBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.UpdateCMEKStatus(context.Background(), clusterId).UpdateCMEKStatusBody(updateCMEKStatusBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomerManagedEncryptionKeysApi.UpdateCMEKStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKStatus`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CustomerManagedEncryptionKeysApi.UpdateCMEKStatus`: %v\n", resp)
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

 **updateCMEKStatusBody** | [**UpdateCMEKStatusBody**](UpdateCMEKStatusBody.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

