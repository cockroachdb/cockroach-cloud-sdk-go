# CustomerManagedEncryptionKeys

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnableCMEKSpec**](CustomerManagedEncryptionKeysApi.md#EnableCMEKSpec) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster
[**GetCMEKClusterInfo**](CustomerManagedEncryptionKeysApi.md#GetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster
[**UpdateCMEKSpec**](CustomerManagedEncryptionKeysApi.md#UpdateCMEKSpec) | **Put** /api/v1/clusters/{cluster_id}/cmek | Enable or update the CMEK spec for a cluster
[**UpdateCMEKStatus**](CustomerManagedEncryptionKeysApi.md#UpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster



## EnableCMEKSpec

> CMEKClusterInfo EnableCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable CMEK for a cluster



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
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.EnableCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
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
[[Back to README]](../README.md)


## GetCMEKClusterInfo

> CMEKClusterInfo GetCMEKClusterInfo(ctx, clusterId).Execute()

Get CMEK-related information for a cluster



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
[[Back to README]](../README.md)


## UpdateCMEKSpec

> CMEKClusterInfo UpdateCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable or update the CMEK spec for a cluster



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
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.UpdateCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
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
[[Back to README]](../README.md)


## UpdateCMEKStatus

> CMEKClusterInfo UpdateCMEKStatus(ctx, clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()

Update the CMEK-related status for a cluster



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
    resp, r, err := api_client.CustomerManagedEncryptionKeysApi.UpdateCMEKStatus(context.Background(), clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()
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
[[Back to README]](../README.md)

