# Folders

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FoldersApi.md#CreateFolder) | **Post** /api/v1/folders | Create a folder
[**DeleteFolder**](FoldersApi.md#DeleteFolder) | **Delete** /api/v1/folders/{folder_id} | Delete a folder
[**GetFolder**](FoldersApi.md#GetFolder) | **Get** /api/v1/folders/{folder_id} | Get folder info for a folder
[**ListFolderContents**](FoldersApi.md#ListFolderContents) | **Get** /api/v1/folders/{folder_id}/contents | List contents of a folder
[**UpdateFolder**](FoldersApi.md#UpdateFolder) | **Patch** /api/v1/folders/{folder_id} | Update a folder



## CreateFolder

> FolderResource CreateFolder(ctx).CreateFolderRequest(createFolderRequest).Execute()

Create a folder

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



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolder struct via the builder pattern


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

Other parameters are passed through a pointer to a apiDeleteFolder struct via the builder pattern


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

Other parameters are passed through a pointer to a apiGetFolder struct via the builder pattern


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

Set `folder_id` to 'root' to list root level contents.

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

Other parameters are passed through a pointer to a apiListFolderContents struct via the builder pattern


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


## UpdateFolder

> FolderResource UpdateFolder(ctx, folderId).UpdateFolderSpecification(updateFolderSpecification).Execute()

Update a folder

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

Other parameters are passed through a pointer to a apiUpdateFolder struct via the builder pattern


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

