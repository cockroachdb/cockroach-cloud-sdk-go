# Databases

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDatabase**](DatabasesApi.md#CreateDatabase) | **Post** /api/v1/clusters/{cluster_id}/databases | Create a new database
[**DeleteDatabase**](DatabasesApi.md#DeleteDatabase) | **Delete** /api/v1/clusters/{cluster_id}/databases/{name} | Delete a database
[**EditDatabase**](DatabasesApi.md#EditDatabase) | **Patch** /api/v1/clusters/{cluster_id}/databases/{name} | Update a database
[**EditDatabase2**](DatabasesApi.md#EditDatabase2) | **Patch** /api/v1/clusters/{cluster_id}/databases | Update a database
[**ListDatabases**](DatabasesApi.md#ListDatabases) | **Get** /api/v1/clusters/{cluster_id}/databases | List databases for a cluster



## CreateDatabase

> Database CreateDatabase(ctx, clusterId).CreateDatabaseBody(createDatabaseBody).Execute()

Create a new database

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
    createDatabaseBody := *openapiclient.NewCreateDatabaseBody("Name_example") // CreateDatabaseBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DatabasesApi.CreateDatabase(context.Background(), clusterId).CreateDatabaseBody(createDatabaseBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.CreateDatabase`: %v\n", resp)
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

 **createDatabaseBody** | [**CreateDatabaseBody**](CreateDatabaseBody.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteDatabase

> Database DeleteDatabase(ctx, clusterId, name).Execute()

Delete a database

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DatabasesApi.DeleteDatabase(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Database**](Database.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## EditDatabase

> Database EditDatabase(ctx, clusterId, name).EditDatabaseBody(editDatabaseBody).Execute()

Update a database

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
    name := "name_example" // string | 
    editDatabaseBody := *openapiclient.NewEditDatabaseBody("Name_example", "NewName_example") // EditDatabaseBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DatabasesApi.EditDatabase(context.Background(), clusterId, name).EditDatabaseBody(editDatabaseBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.EditDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.EditDatabase`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editDatabaseBody** | [**EditDatabaseBody**](EditDatabaseBody.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## EditDatabase2

> Database EditDatabase2(ctx, clusterId).EditDatabaseBody(editDatabaseBody).Execute()

Update a database

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
    editDatabaseBody := *openapiclient.NewEditDatabaseBody("Name_example", "NewName_example") // EditDatabaseBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.DatabasesApi.EditDatabase2(context.Background(), clusterId).EditDatabaseBody(editDatabaseBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.EditDatabase2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase2`: Database
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.EditDatabase2`: %v\n", resp)
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

 **editDatabaseBody** | [**EditDatabaseBody**](EditDatabaseBody.md) |  | 

### Return type

[**Database**](Database.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDatabases

> ListDatabasesResponse ListDatabases(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List databases for a cluster

Sort order: Database name ascending

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    resp, r, err := api_client.DatabasesApi.ListDatabases(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DatabasesApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: ListDatabasesResponse
    fmt.Fprintf(os.Stdout, "Response from `DatabasesApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the ListDatabasesOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListDatabasesResponse**](ListDatabasesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

