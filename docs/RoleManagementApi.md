# RoleManagement

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToRole**](RoleManagementApi.md#AddUserToRole) | **Post** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Add a role to a user, service account, or group
[**GetAllRolesForUser**](RoleManagementApi.md#GetAllRolesForUser) | **Get** /api/v1/roles/{user_id} | Get all role grants for a user, service account, or group
[**GetPersonUsersByEmail**](RoleManagementApi.md#GetPersonUsersByEmail) | **Get** /api/v1/users/persons-by-email | Search person users by email address
[**ListRoleGrants**](RoleManagementApi.md#ListRoleGrants) | **Get** /api/v1/roles | List all RoleGrants
[**RemoveUserFromRole**](RoleManagementApi.md#RemoveUserFromRole) | **Delete** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Remove a role from a user, service account, or group
[**SetRolesForUser**](RoleManagementApi.md#SetRolesForUser) | **Put** /api/v1/roles/{user_id} | Replace the roles for a user, service account, or group with exactly those provided



## AddUserToRole

> GetAllRolesForUserResponse AddUserToRole(ctx, userId, resourceType, resourceId, roleName).Execute()

Add a role to a user, service account, or group

Add a single role to a user, service account, or group by providing its ID in the user_id path parameter.

Roles that will be added as a result of this call must follow the CC rules for role assignment:
https://www.cockroachlabs.com/docs/cockroachcloud/authorization#organization-user-roles

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
    userId := "userId_example" // string | user_id is the ID of the user, service account, or group to add the role to.
    resourceType := "resourceType_example" // string | resource_type is the type of resource the role applies to.
    resourceId := "resourceId_example" // string | resource_id is the ID of the resource the role applies to. Pass an empty string for organization-scoped roles.
    roleName := "roleName_example" // string | role_name is the role to grant.

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
**userId** | **string** | user_id is the ID of the user, service account, or group to add the role to. | 
**resourceType** | **string** | resource_type is the type of resource the role applies to. | 
**resourceId** | **string** | resource_id is the ID of the resource the role applies to. Pass an empty string for organization-scoped roles. | 
**roleName** | **string** | role_name is the role to grant. | 

### Other Parameters


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

Get all role grants for a user, service account, or group

Get all role grants for a user, service account, or group by providing its ID in the user_id path parameter.

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN
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
    userId := "userId_example" // string | user_id is the ID of the user, service account, or group whose roles are being retrieved.

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
**userId** | **string** | user_id is the ID of the user, service account, or group whose roles are being retrieved. | 

### Other Parameters


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

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN
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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


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

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN
- FOLDER_ADMIN


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

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListRoleGrantsOptions struct.

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

Remove a role from a user, service account, or group

Remove a single role from a user, service account, or group by providing its ID in the user_id path parameter. A principal's last remaining role cannot be removed; because groups do not hold the implicit Org Member role that users and service accounts have, a group must always retain at least one other role.

Roles that will be removed as a result of this call must follow the CC rules for role assignment:
https://www.cockroachlabs.com/docs/cockroachcloud/authorization#organization-user-roles

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
    userId := "userId_example" // string | user_id is the ID of the user, service account, or group to remove the role from.
    resourceType := "resourceType_example" // string | resource_type is the type of resource the role applies to.
    resourceId := "resourceId_example" // string | resource_id is the ID of the resource the role applies to. Pass an empty string for organization-scoped roles.
    roleName := "roleName_example" // string | role_name is the role to revoke.

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
**userId** | **string** | user_id is the ID of the user, service account, or group to remove the role from. | 
**resourceType** | **string** | resource_type is the type of resource the role applies to. | 
**resourceId** | **string** | resource_id is the ID of the resource the role applies to. Pass an empty string for organization-scoped roles. | 
**roleName** | **string** | role_name is the role to revoke. | 

### Other Parameters


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

> GetAllRolesForUserResponse SetRolesForUser(ctx, userId).SetRolesForUserBody(setRolesForUserBody).Execute()

Replace the roles for a user, service account, or group with exactly those provided

Replace the entire role set for a user, service account, or group by providing its ID in the user_id path parameter.

Roles that will be removed or added as a result of this call must follow the CC rules for role assignment:
https://www.cockroachlabs.com/docs/cockroachcloud/authorization#organization-user-roles

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
    userId := "userId_example" // string | user_id is the ID of the user, service account, or group.
    setRolesForUserBody := *openapiclient.NewSetRolesForUserBody([]openapiclient.BuiltInRole{*openapiclient.NewBuiltInRole(openapiclient.OrganizationUserRole.Type("BILLING_COORDINATOR"), *openapiclient.NewResource(openapiclient.ResourceType.Type("ORGANIZATION")))}) // SetRolesForUserBody | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.RoleManagementApi.SetRolesForUser(context.Background(), userId).SetRolesForUserBody(setRolesForUserBody).Execute()
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
**userId** | **string** | user_id is the ID of the user, service account, or group. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setRolesForUserBody** | [**SetRolesForUserBody**](SetRolesForUserBody.md) |  | 

### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

