# VirtualClusterWorkspaces

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVirtualClusterWorkspace**](VirtualClusterWorkspacesApi.md#GetVirtualClusterWorkspace) | **Get** /api/v1/virtual-cluster-workspaces/{workspace_id} | Show details for a specific workspace
[**ListVirtualClusterWorkspaces**](VirtualClusterWorkspacesApi.md#ListVirtualClusterWorkspaces) | **Get** /api/v1/virtual-cluster-workspaces | List workspaces in the organization
[**UpdateVirtualClusterWorkspace**](VirtualClusterWorkspacesApi.md#UpdateVirtualClusterWorkspace) | **Patch** /api/v1/virtual-cluster-workspaces/{workspace_id} | Update the configuration of a specific workspace



## GetVirtualClusterWorkspace

> VirtualClusterWorkspace GetVirtualClusterWorkspace(ctx, workspaceId).Execute()

Show details for a specific workspace

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- CLUSTER_CREATOR
- AUDITOR


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
    workspaceId := "workspaceId_example" // string | The unique identifier of the workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.VirtualClusterWorkspacesApi.GetVirtualClusterWorkspace(context.Background(), workspaceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterWorkspacesApi.GetVirtualClusterWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualClusterWorkspace`: VirtualClusterWorkspace
    fmt.Fprintf(os.Stdout, "Response from `VirtualClusterWorkspacesApi.GetVirtualClusterWorkspace`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualClusterWorkspace**](VirtualClusterWorkspace.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVirtualClusterWorkspaces

> ListVirtualClusterWorkspacesResponse ListVirtualClusterWorkspaces(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List workspaces in the organization

Sort order: workspace name

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- CLUSTER_CREATOR
- AUDITOR


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
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.VirtualClusterWorkspacesApi.ListVirtualClusterWorkspaces(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterWorkspacesApi.ListVirtualClusterWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualClusterWorkspaces`: ListVirtualClusterWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `VirtualClusterWorkspacesApi.ListVirtualClusterWorkspaces`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListVirtualClusterWorkspacesOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListVirtualClusterWorkspacesResponse**](ListVirtualClusterWorkspacesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVirtualClusterWorkspace

> VirtualClusterWorkspace UpdateVirtualClusterWorkspace(ctx, workspaceId).UpdateVirtualClusterWorkspaceSpecification(updateVirtualClusterWorkspaceSpecification).Execute()

Update the configuration of a specific workspace

Can be used by the following roles assigned at the organization, folder or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER

Changes to a workspace's configuration apply only to virtual clusters created after the change, not to virtual clusters that already exist. To modify existing virtual clusters, see [CockroachCloud_UpdateVirtualCluster](#CockroachCloud_UpdateVirtualCluster).

The role must be assigned on the host cluster backing the workspace, not on the workspace itself.

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
    workspaceId := "workspaceId_example" // string | The unique identifier of the workspace.
    updateVirtualClusterWorkspaceSpecification := *openapiclient.NewUpdateVirtualClusterWorkspaceSpecification() // UpdateVirtualClusterWorkspaceSpecification | The changes to apply to the workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.VirtualClusterWorkspacesApi.UpdateVirtualClusterWorkspace(context.Background(), workspaceId).UpdateVirtualClusterWorkspaceSpecification(updateVirtualClusterWorkspaceSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualClusterWorkspacesApi.UpdateVirtualClusterWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVirtualClusterWorkspace`: VirtualClusterWorkspace
    fmt.Fprintf(os.Stdout, "Response from `VirtualClusterWorkspacesApi.UpdateVirtualClusterWorkspace`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVirtualClusterWorkspaceSpecification** | [**UpdateVirtualClusterWorkspaceSpecification**](UpdateVirtualClusterWorkspaceSpecification.md) | The changes to apply to the workspace. | 

### Return type

[**VirtualClusterWorkspace**](VirtualClusterWorkspace.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

