# RoleManagement

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToRole**](RoleManagementApi.md#AddUserToRole) | **Post** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Add the user to the given role
[**GetAllRolesForUser**](RoleManagementApi.md#GetAllRolesForUser) | **Get** /api/v1/roles/{user_id} | Get all Role Grants for a user
[**GetPersonUsersByEmail**](RoleManagementApi.md#GetPersonUsersByEmail) | **Get** /api/v1/users/persons-by-email | Search person users by email address
[**ListRoleGrants**](RoleManagementApi.md#ListRoleGrants) | **Get** /api/v1/roles | List all RoleGrants
[**RemoveUserFromRole**](RoleManagementApi.md#RemoveUserFromRole) | **Delete** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Remove the user from the given role
[**SetRolesForUser**](RoleManagementApi.md#SetRolesForUser) | **Put** /api/v1/roles/{user_id} | Make a user&#39;s roles exactly those provided



## AddUserToRole

> GetAllRolesForUserResponse AddUserToRole(ctx, userId, resourceType, resourceId, roleName).Execute()

Add the user to the given role

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
    userId := "userId_example" // string | 
    resourceType := "resourceType_example" // string | 
    resourceId := "resourceId_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.AddUserToRole(context.Background(), userId, resourceType, resourceId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.AddUserToRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToRole`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.AddUserToRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToRole struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAllRolesForUser

> GetAllRolesForUserResponse GetAllRolesForUser(ctx, userId).Execute()

Get all Role Grants for a user

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.GetAllRolesForUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.GetAllRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRolesForUser`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.GetAllRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesForUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPersonUsersByEmail

> GetPersonUsersByEmailResponse GetPersonUsersByEmail(ctx).Email(email).Execute()

Search person users by email address

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
    email := "email_example" // string | an email address is required.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.GetPersonUsersByEmail(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.GetPersonUsersByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonUsersByEmail`: GetPersonUsersByEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.GetPersonUsersByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonUsersByEmail struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | an email address is required. | 

### Return type

[**GetPersonUsersByEmailResponse**](GetPersonUsersByEmailResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListRoleGrants

> ListRoleGrantsResponse ListRoleGrants(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all RoleGrants

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
    resp, r, err := api_client.RoleManagementApi.ListRoleGrants(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.ListRoleGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleGrants`: ListRoleGrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.ListRoleGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleGrants struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListRoleGrantsResponse**](ListRoleGrantsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## RemoveUserFromRole

> GetAllRolesForUserResponse RemoveUserFromRole(ctx, userId, resourceType, resourceId, roleName).Execute()

Remove the user from the given role

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
    userId := "userId_example" // string | 
    resourceType := "resourceType_example" // string | 
    resourceId := "resourceId_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.RemoveUserFromRole(context.Background(), userId, resourceType, resourceId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.RemoveUserFromRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserFromRole`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.RemoveUserFromRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromRole struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetRolesForUser

> GetAllRolesForUserResponse SetRolesForUser(ctx, userId).CockroachCloudSetRolesForUserRequest(cockroachCloudSetRolesForUserRequest).Execute()

Make a user's roles exactly those provided

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
    userId := "userId_example" // string | 
    cockroachCloudSetRolesForUserRequest := *openapiclient.NewCockroachCloudSetRolesForUserRequest([]openapiclient.BuiltInRole{*openapiclient.NewBuiltInRole(openapiclient.OrganizationUserRole.Type("DEVELOPER"), *openapiclient.NewResource(openapiclient.ResourceType.Type("ORGANIZATION")))}) // CockroachCloudSetRolesForUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.SetRolesForUser(context.Background(), userId).CockroachCloudSetRolesForUserRequest(cockroachCloudSetRolesForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleManagementApi.SetRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRolesForUser`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `RoleManagementApi.SetRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRolesForUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cockroachCloudSetRolesForUserRequest** | [**CockroachCloudSetRolesForUserRequest**](CockroachCloudSetRolesForUserRequest.md) |  | 

### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

