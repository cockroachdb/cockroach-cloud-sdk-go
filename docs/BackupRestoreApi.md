# BackupRestore

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRestore**](BackupRestoreApi.md#CreateRestore) | **Post** /api/v1/clusters/{destination_cluster_id}/restores | Create a restore
[**GetBackupConfiguration**](BackupRestoreApi.md#GetBackupConfiguration) | **Get** /api/v1/clusters/{cluster_id}/backups-config | Get the backup configuration for a cluster
[**GetRestore**](BackupRestoreApi.md#GetRestore) | **Get** /api/v1/clusters/{cluster_id}/restores/{restore_id} | View a restore job
[**ListBackups**](BackupRestoreApi.md#ListBackups) | **Get** /api/v1/clusters/{cluster_id}/backups | List cluster backups
[**ListRestores**](BackupRestoreApi.md#ListRestores) | **Get** /api/v1/clusters/{cluster_id}/restores | List restore operations
[**UpdateBackupConfiguration**](BackupRestoreApi.md#UpdateBackupConfiguration) | **Patch** /api/v1/clusters/{cluster_id}/backups-config | Update the backup configuration for a cluster



## CreateRestore

> Restore CreateRestore(ctx, destinationClusterId).CockroachCloudCreateRestoreRequest(cockroachCloudCreateRestoreRequest).Execute()

Create a restore

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
    destinationClusterId := "destinationClusterId_example" // string | The ID of the cluster where the backup will be restored.
    cockroachCloudCreateRestoreRequest := *openapiclient.NewCockroachCloudCreateRestoreRequest(openapiclient.RestoreType.Type("CLUSTER")) // CockroachCloudCreateRestoreRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.CreateRestore(context.Background(), destinationClusterId).CockroachCloudCreateRestoreRequest(cockroachCloudCreateRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.CreateRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRestore`: Restore
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.CreateRestore`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**destinationClusterId** | **string** | The ID of the cluster where the backup will be restored. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cockroachCloudCreateRestoreRequest** | [**CockroachCloudCreateRestoreRequest**](CockroachCloudCreateRestoreRequest.md) |  | 

### Return type

[**Restore**](Restore.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBackupConfiguration

> BackupConfiguration GetBackupConfiguration(ctx, clusterId).Execute()

Get the backup configuration for a cluster

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
    clusterId := "clusterId_example" // string | The UUID of the cluster that this backup configuration belongs to.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.GetBackupConfiguration(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.GetBackupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBackupConfiguration`: BackupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.GetBackupConfiguration`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The UUID of the cluster that this backup configuration belongs to. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetRestore

> Restore GetRestore(ctx, clusterId, restoreId).Execute()

View a restore job

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
    clusterId := "clusterId_example" // string | The ID of the cluster where the restore ran or is currently running.
    restoreId := "restoreId_example" // string | The ID of the restore.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.GetRestore(context.Background(), clusterId, restoreId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.GetRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRestore`: Restore
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.GetRestore`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster where the restore ran or is currently running. | 
**restoreId** | **string** | The ID of the restore. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Restore**](Restore.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBackups

> ListBackupsResponse ListBackups(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).StartTime(startTime).EndTime(endTime).Execute()

List cluster backups

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
    clusterId := "clusterId_example" // string | The cluster associated with the backups being retrieved.
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)
    startTime := time.Now() // time.Time | The beginning of the time range (inclusive) used to search for backups. (optional)
    endTime := time.Now() // time.Time | The end of the time range (exclusive) used to search for backups. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.ListBackups(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.ListBackups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBackups`: ListBackupsResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.ListBackups`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The cluster associated with the backups being retrieved. | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListBackupsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 
 **startTime** | **time.Time** | The beginning of the time range (inclusive) used to search for backups. | 
 **endTime** | **time.Time** | The end of the time range (exclusive) used to search for backups. | 

### Return type

[**ListBackupsResponse**](ListBackupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListRestores

> ListRestoresResponse ListRestores(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).StartTime(startTime).EndTime(endTime).Execute()

List restore operations

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
    clusterId := "clusterId_example" // string | The ID of the cluster where the restores ran or are currently running.
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)
    startTime := time.Now() // time.Time | The beginning of the time range (inclusive) used to search for restores. (optional)
    endTime := time.Now() // time.Time | The end of the time range (exclusive) used to search for restores. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.ListRestores(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).StartTime(startTime).EndTime(endTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.ListRestores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRestores`: ListRestoresResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.ListRestores`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The ID of the cluster where the restores ran or are currently running. | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListRestoresOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 
 **startTime** | **time.Time** | The beginning of the time range (inclusive) used to search for restores. | 
 **endTime** | **time.Time** | The end of the time range (exclusive) used to search for restores. | 

### Return type

[**ListRestoresResponse**](ListRestoresResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBackupConfiguration

> BackupConfiguration UpdateBackupConfiguration(ctx, clusterId).UpdateBackupConfigurationSpec(updateBackupConfigurationSpec).Execute()

Update the backup configuration for a cluster

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
    clusterId := "clusterId_example" // string | The UUID of the cluster that this backup configuration belongs to.
    updateBackupConfigurationSpec := *openapiclient.NewUpdateBackupConfigurationSpec() // UpdateBackupConfigurationSpec | spec contains the information that is being updated for the given BackupConfiguration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BackupRestoreApi.UpdateBackupConfiguration(context.Background(), clusterId).UpdateBackupConfigurationSpec(updateBackupConfigurationSpec).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupRestoreApi.UpdateBackupConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBackupConfiguration`: BackupConfiguration
    fmt.Fprintf(os.Stdout, "Response from `BackupRestoreApi.UpdateBackupConfiguration`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The UUID of the cluster that this backup configuration belongs to. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBackupConfigurationSpec** | [**UpdateBackupConfigurationSpec**](UpdateBackupConfigurationSpec.md) | spec contains the information that is being updated for the given BackupConfiguration. | 

### Return type

[**BackupConfiguration**](BackupConfiguration.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

