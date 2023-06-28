# SQLUsers

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSQLUser**](SQLUsersApi.md#CreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user
[**DeleteSQLUser**](SQLUsersApi.md#DeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user
[**ListSQLUsers**](SQLUsersApi.md#ListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster
[**UpdateSQLUserPassword**](SQLUsersApi.md#UpdateSQLUserPassword) | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Update a SQL user&#39;s password



## CreateSQLUser

> SQLUser CreateSQLUser(ctx, clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()

Create a new SQL user
    
Can be used by the following roles assigned at the organization or cluster scope:
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
    clusterId := "clusterId_example" // string | 
    createSQLUserRequest := *openapiclient.NewCreateSQLUserRequest("Name_example", "Password_example") // CreateSQLUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SQLUsersApi.CreateSQLUser(context.Background(), clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLUsersApi.CreateSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `SQLUsersApi.CreateSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSQLUserRequest** | [**CreateSQLUserRequest**](CreateSQLUserRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteSQLUser

> SQLUser DeleteSQLUser(ctx, clusterId, name).Execute()

Delete a SQL user
    
Can be used by the following roles assigned at the organization or cluster scope:
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
    clusterId := "clusterId_example" // string | 
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SQLUsersApi.DeleteSQLUser(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLUsersApi.DeleteSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `SQLUsersApi.DeleteSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSQLUsers

> ListSQLUsersResponse ListSQLUsers(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List SQL users for a cluster
    
Can be used by the following roles assigned at the organization or cluster scope:
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
    resp, r, err := api_client.SQLUsersApi.ListSQLUsers(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLUsersApi.ListSQLUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSQLUsers`: ListSQLUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `SQLUsersApi.ListSQLUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSQLUsers struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListSQLUsersResponse**](ListSQLUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSQLUserPassword

> SQLUser UpdateSQLUserPassword(ctx, clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()

Update a SQL user's password
    
Can be used by the following roles assigned at the organization or cluster scope:
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
    clusterId := "clusterId_example" // string | 
    name := "name_example" // string | 
    updateSQLUserPasswordRequest := *openapiclient.NewUpdateSQLUserPasswordRequest("Password_example") // UpdateSQLUserPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SQLUsersApi.UpdateSQLUserPassword(context.Background(), clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SQLUsersApi.UpdateSQLUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSQLUserPassword`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `SQLUsersApi.UpdateSQLUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSQLUserPassword struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSQLUserPasswordRequest** | [**UpdateSQLUserPasswordRequest**](UpdateSQLUserPasswordRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

