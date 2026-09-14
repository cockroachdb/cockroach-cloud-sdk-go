# MigrationAssistant

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMigrationAssistant**](MigrationAssistantApi.md#CreateMigrationAssistant) | **Post** /api/v1/clusters/{cluster_id}/migration-assistant | Provision a Migration Assistant for a cluster.
[**DeleteMigrationAssistant**](MigrationAssistantApi.md#DeleteMigrationAssistant) | **Delete** /api/v1/clusters/{cluster_id}/migration-assistant | Delete the Migration Assistant for a cluster.
[**DeleteMigrationAssistantSourceCACert**](MigrationAssistantApi.md#DeleteMigrationAssistantSourceCACert) | **Delete** /api/v1/clusters/{cluster_id}/migration-assistant/source-ca-cert | Delete the source-database CA certificate for a cluster&#39;s Migration Assistant.
[**GetMigrationAssistant**](MigrationAssistantApi.md#GetMigrationAssistant) | **Get** /api/v1/clusters/{cluster_id}/migration-assistant | Get the Migration Assistant for a cluster.
[**GetMigrationAssistantSourceCACert**](MigrationAssistantApi.md#GetMigrationAssistantSourceCACert) | **Get** /api/v1/clusters/{cluster_id}/migration-assistant/source-ca-cert | Get the source-database CA certificate for a cluster&#39;s Migration Assistant.
[**UpdateMigrationAssistantSourceCACert**](MigrationAssistantApi.md#UpdateMigrationAssistantSourceCACert) | **Patch** /api/v1/clusters/{cluster_id}/migration-assistant/source-ca-cert | Update the source-database CA certificate for a cluster&#39;s Migration Assistant.



## CreateMigrationAssistant

> MigrationAssistant CreateMigrationAssistant(ctx, clusterId).CreateMigrationAssistantBody(createMigrationAssistantBody).Execute()

Provision a Migration Assistant for a cluster.

The assistant is a per-cluster, dedicated-cluster-only migration tool. Poll
GetMigrationAssistant until its status is RUNNING to retrieve the URL and
credentials used to sign in, or FAILED to inspect last_error.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster to deploy the assistant on.
    createMigrationAssistantBody := *openapiclient.NewCreateMigrationAssistantBody() // CreateMigrationAssistantBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.CreateMigrationAssistant(context.Background(), clusterId).CreateMigrationAssistantBody(createMigrationAssistantBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.CreateMigrationAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMigrationAssistant`: MigrationAssistant
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.CreateMigrationAssistant`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster to deploy the assistant on. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createMigrationAssistantBody** | [**CreateMigrationAssistantBody**](CreateMigrationAssistantBody.md) |  | 

### Return type

[**MigrationAssistant**](MigrationAssistant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteMigrationAssistant

> MigrationAssistant DeleteMigrationAssistant(ctx, clusterId).Execute()

Delete the Migration Assistant for a cluster.

Tears down the deployment and returns the assistant in its DELETING status.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster whose assistant to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.DeleteMigrationAssistant(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.DeleteMigrationAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMigrationAssistant`: MigrationAssistant
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.DeleteMigrationAssistant`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster whose assistant to delete. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MigrationAssistant**](MigrationAssistant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteMigrationAssistantSourceCACert

> MigrationAssistantSourceCACert DeleteMigrationAssistantSourceCACert(ctx, clusterId).Execute()

Delete the source-database CA certificate for a cluster's Migration Assistant.

Reverts the assistant to the system trust roots and returns the deleted
certificate bundle.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster whose source-database CA certificate to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.DeleteMigrationAssistantSourceCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.DeleteMigrationAssistantSourceCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMigrationAssistantSourceCACert`: MigrationAssistantSourceCACert
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.DeleteMigrationAssistantSourceCACert`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster whose source-database CA certificate to delete. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MigrationAssistantSourceCACert**](MigrationAssistantSourceCACert.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMigrationAssistant

> MigrationAssistant GetMigrationAssistant(ctx, clusterId).Execute()

Get the Migration Assistant for a cluster.

Includes the lifecycle status and, while RUNNING, the URL and sign-in
credentials for the assistant; treat the response as sensitive.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster whose assistant to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.GetMigrationAssistant(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.GetMigrationAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMigrationAssistant`: MigrationAssistant
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.GetMigrationAssistant`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster whose assistant to fetch. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MigrationAssistant**](MigrationAssistant.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMigrationAssistantSourceCACert

> MigrationAssistantSourceCACert GetMigrationAssistantSourceCACert(ctx, clusterId).Execute()

Get the source-database CA certificate for a cluster's Migration Assistant.

Returns the custom CA certificate bundle the assistant trusts when
connecting to its source database. Returns an empty bundle when the
assistant uses the system trust roots.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster whose source-database CA certificate to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.GetMigrationAssistantSourceCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.GetMigrationAssistantSourceCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMigrationAssistantSourceCACert`: MigrationAssistantSourceCACert
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.GetMigrationAssistantSourceCACert`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster whose source-database CA certificate to fetch. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MigrationAssistantSourceCACert**](MigrationAssistantSourceCACert.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateMigrationAssistantSourceCACert

> MigrationAssistantSourceCACert UpdateMigrationAssistantSourceCACert(ctx, clusterId).UpdateMigrationAssistantSourceCACertBody(updateMigrationAssistantSourceCACertBody).Execute()

Update the source-database CA certificate for a cluster's Migration Assistant.

Replaces the custom CA certificate bundle the assistant trusts when
connecting to its source database. Returns the updated certificate bundle.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
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
    clusterId := "clusterId_example" // string | The ID of the cluster whose source-database CA certificate to update.
    updateMigrationAssistantSourceCACertBody := *openapiclient.NewUpdateMigrationAssistantSourceCACertBody("X509PemCert_example") // UpdateMigrationAssistantSourceCACertBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MigrationAssistantApi.UpdateMigrationAssistantSourceCACert(context.Background(), clusterId).UpdateMigrationAssistantSourceCACertBody(updateMigrationAssistantSourceCACertBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MigrationAssistantApi.UpdateMigrationAssistantSourceCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMigrationAssistantSourceCACert`: MigrationAssistantSourceCACert
    fmt.Fprintf(os.Stdout, "Response from `MigrationAssistantApi.UpdateMigrationAssistantSourceCACert`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster whose source-database CA certificate to update. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateMigrationAssistantSourceCACertBody** | [**UpdateMigrationAssistantSourceCACertBody**](UpdateMigrationAssistantSourceCACertBody.md) |  | 

### Return type

[**MigrationAssistantSourceCACert**](MigrationAssistantSourceCACert.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

