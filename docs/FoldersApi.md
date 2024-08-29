# Folders

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersApi.md#CreateFolder) | **Post** /api/v1/folders | Create a folder
[**DeleteFolder**](FoldersApi.md#DeleteFolder) | **Delete** /api/v1/folders/{folder_id} | Delete a folder
[**GetFolder**](FoldersApi.md#GetFolder) | **Get** /api/v1/folders/{folder_id} | Get folder info for a folder
[**ListFolderContents**](FoldersApi.md#ListFolderContents) | **Get** /api/v1/folders/{folder_id}/contents | List contents of a folder
[**ListFolders**](FoldersApi.md#ListFolders) | **Get** /api/v1/folders | List folders owned by an organization
[**UpdateFolder**](FoldersApi.md#UpdateFolder) | **Patch** /api/v1/folders/{folder_id} | Update a folder



## CreateFolder

> FolderResource CreateFolder(ctx).CreateFolderRequest(createFolderRequest).Execute()

Create a folder

Can be used by the following roles assigned at the organization or folder scope:
- FOLDER_ADMIN


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
    createFolderRequest := *openapiclient.NewCreateFolderRequest("Name_example") // CreateFolderRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.CreateFolder(context.Background()).CreateFolderRequest(createFolderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.CreateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFolder`: FolderResource
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.CreateFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFolderRequest** | [**CreateFolderRequest**](CreateFolderRequest.md) |  | 

### Return type

[**FolderResource**](FolderResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteFolder

> map[string]interface{} DeleteFolder(ctx, folderId).Execute()

Delete a folder

Can be used by the following roles assigned at the organization or folder scope:
- FOLDER_ADMIN


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
    folderId := "folderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.DeleteFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.DeleteFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFolder`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.DeleteFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

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


## GetFolder

> FolderResource GetFolder(ctx, folderId).Execute()

Get folder info for a folder

Can be used by the following roles assigned at the organization or folder scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- CLUSTER_CREATOR
- FOLDER_ADMIN
- FOLDER_MOVER


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
    folderId := "folderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.GetFolder(context.Background(), folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.GetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolder`: FolderResource
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.GetFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FolderResource**](FolderResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFolderContents

> FolderResourceList ListFolderContents(ctx, folderId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List contents of a folder

Set `folder_id` to 'root' to list root level contents.  Sort order: Folders sorted by name, followed by Clusters sorted by name.

Can be used by the following roles assigned at the organization, folder or cluster scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- FOLDER_ADMIN
- FOLDER_MOVER


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
    folderId := "folderId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.ListFolderContents(context.Background(), folderId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.ListFolderContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFolderContents`: FolderResourceList
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.ListFolderContents`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListFolderContentsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**FolderResourceList**](FolderResourceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFolders

> ListFoldersResponse ListFolders(ctx).Path(path).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List folders owned by an organization

Sort order: Folder name

Can be used by the following roles assigned at the organization or folder scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER
- CLUSTER_DEVELOPER
- CLUSTER_CREATOR
- FOLDER_ADMIN
- FOLDER_MOVER


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
    path := "path_example" // string | Optional filter to limit the response to include only results that match the given absolute path to that folder. Preceding and ending \"/\" are optional. For example /folder1/folder2, /folder1/folder2/, folder1/folder2, and folder1/folder2/ are all equivalent. If no matching folder is found, an empty list is returned. Because folder paths are passed via the query parameters, they must be URL-encoded. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.ListFolders(context.Background()).Path(path).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.ListFolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFolders`: ListFoldersResponse
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.ListFolders`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListFoldersOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Optional filter to limit the response to include only results that match the given absolute path to that folder. Preceding and ending \&quot;/\&quot; are optional. For example /folder1/folder2, /folder1/folder2/, folder1/folder2, and folder1/folder2/ are all equivalent. If no matching folder is found, an empty list is returned. Because folder paths are passed via the query parameters, they must be URL-encoded. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListFoldersResponse**](ListFoldersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateFolder

> FolderResource UpdateFolder(ctx, folderId).UpdateFolderSpecification(updateFolderSpecification).Execute()

Update a folder

Can be used by the following roles assigned at the organization or folder scope:
- FOLDER_ADMIN
- FOLDER_MOVER


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
    folderId := "folderId_example" // string | 
    updateFolderSpecification := *openapiclient.NewUpdateFolderSpecification() // UpdateFolderSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.FoldersApi.UpdateFolder(context.Background(), folderId).UpdateFolderSpecification(updateFolderSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FoldersApi.UpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFolder`: FolderResource
    fmt.Fprintf(os.Stdout, "Response from `FoldersApi.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFolderSpecification** | [**UpdateFolderSpecification**](UpdateFolderSpecification.md) |  | 

### Return type

[**FolderResource**](FolderResource.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

