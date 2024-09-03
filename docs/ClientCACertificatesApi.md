# ClientCACertificates

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteClientCACert**](ClientCACertificatesApi.md#DeleteClientCACert) | **Delete** /api/v1/clusters/{cluster_id}/client-ca-cert | Delete Client CA Cert for a cluster
[**GetClientCACert**](ClientCACertificatesApi.md#GetClientCACert) | **Get** /api/v1/clusters/{cluster_id}/client-ca-cert | Get Client CA Cert information for a cluster
[**SetClientCACert**](ClientCACertificatesApi.md#SetClientCACert) | **Post** /api/v1/clusters/{cluster_id}/client-ca-cert | Set Client CA Cert for a cluster
[**UpdateClientCACert**](ClientCACertificatesApi.md#UpdateClientCACert) | **Patch** /api/v1/clusters/{cluster_id}/client-ca-cert | Update Client CA Cert for a cluster



## DeleteClientCACert

> ClientCACertInfo DeleteClientCACert(ctx, clusterId).Execute()

Delete Client CA Cert for a cluster

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
    resp, r, err := api_client.ClientCACertificatesApi.DeleteClientCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientCACertificatesApi.DeleteClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `ClientCACertificatesApi.DeleteClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetClientCACert

> ClientCACertInfo GetClientCACert(ctx, clusterId).Execute()

Get Client CA Cert information for a cluster

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
    resp, r, err := api_client.ClientCACertificatesApi.GetClientCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientCACertificatesApi.GetClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `ClientCACertificatesApi.GetClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetClientCACert

> ClientCACertInfo SetClientCACert(ctx, clusterId).SetClientCACertRequest(setClientCACertRequest).Execute()

Set Client CA Cert for a cluster

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
    setClientCACertRequest := *openapiclient.NewSetClientCACertRequest("X509PemCert_example") // SetClientCACertRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClientCACertificatesApi.SetClientCACert(context.Background(), clusterId).SetClientCACertRequest(setClientCACertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientCACertificatesApi.SetClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `ClientCACertificatesApi.SetClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setClientCACertRequest** | [**SetClientCACertRequest**](SetClientCACertRequest.md) |  | 

### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateClientCACert

> ClientCACertInfo UpdateClientCACert(ctx, clusterId).UpdateClientCACertRequest(updateClientCACertRequest).Execute()

Update Client CA Cert for a cluster

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
    updateClientCACertRequest := *openapiclient.NewUpdateClientCACertRequest() // UpdateClientCACertRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ClientCACertificatesApi.UpdateClientCACert(context.Background(), clusterId).UpdateClientCACertRequest(updateClientCACertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClientCACertificatesApi.UpdateClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `ClientCACertificatesApi.UpdateClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClientCACertRequest** | [**UpdateClientCACertRequest**](UpdateClientCACertRequest.md) |  | 

### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

