# BackupRestore

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBackupConfiguration**](BackupRestoreApi.md#GetBackupConfiguration) | **Get** /api/v1/clusters/{cluster_id}/backups-config | Get the backup configuration for a cluster
[**UpdateBackupConfiguration**](BackupRestoreApi.md#UpdateBackupConfiguration) | **Patch** /api/v1/clusters/{cluster_id}/backups-config | Update the backup configuration for a cluster



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

