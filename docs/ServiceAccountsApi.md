# ServiceAccounts

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceAccount**](ServiceAccountsApi.md#CreateServiceAccount) | **Post** /api/v1/service-accounts | Create a service account
[**DeleteServiceAccount**](ServiceAccountsApi.md#DeleteServiceAccount) | **Delete** /api/v1/service-accounts/{id} | Delete a service account
[**GetServiceAccount**](ServiceAccountsApi.md#GetServiceAccount) | **Get** /api/v1/service-accounts/{id} | Get a service account by ID
[**ListServiceAccounts**](ServiceAccountsApi.md#ListServiceAccounts) | **Get** /api/v1/service-accounts | List service accounts for an organization
[**UpdateServiceAccount**](ServiceAccountsApi.md#UpdateServiceAccount) | **Patch** /api/v1/service-accounts/{id} | Update a service account



## CreateServiceAccount

> ServiceAccount CreateServiceAccount(ctx).CreateServiceAccountRequest(createServiceAccountRequest).Execute()

Create a service account

Can be used by the following roles assigned at the organization scope:
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
    createServiceAccountRequest := *openapiclient.NewCreateServiceAccountRequest("Description_example", "Name_example", []openapiclient.BuiltInRole{*openapiclient.NewBuiltInRole(openapiclient.OrganizationUserRole.Type("BILLING_COORDINATOR"), *openapiclient.NewResource(openapiclient.ResourceType.Type("ORGANIZATION")))}) // CreateServiceAccountRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.CreateServiceAccount(context.Background()).CreateServiceAccountRequest(createServiceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.CreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createServiceAccountRequest** | [**CreateServiceAccountRequest**](CreateServiceAccountRequest.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteServiceAccount

> ServiceAccount DeleteServiceAccount(ctx, id).Execute()

Delete a service account

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    id := "id_example" // string | the ID of the service account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.DeleteServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.DeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.DeleteServiceAccount`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the ID of the service account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetServiceAccount

> ServiceAccount GetServiceAccount(ctx, id).Execute()

Get a service account by ID

Can be used by the following roles assigned at the organization scope:
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
    id := "id_example" // string | The ID of the service account.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.GetServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.GetServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.GetServiceAccount`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the service account. | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListServiceAccounts

> ListServiceAccountsResponse ListServiceAccounts(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List service accounts for an organization

Sort order: Service account name

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- CLUSTER_ADMIN


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
    resp, r, err := api_client.ServiceAccountsApi.ListServiceAccounts(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.ListServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAccounts`: ListServiceAccountsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListServiceAccountsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListServiceAccountsResponse**](ListServiceAccountsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateServiceAccount

> ServiceAccount UpdateServiceAccount(ctx, id).UpdateServiceAccountSpecification(updateServiceAccountSpecification).Execute()

Update a service account

To manage roles associated with a service account after creation, pass the service_account_id instead of a user_id to any [Role Management endpoint](#tag--Role-Management).

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    id := "id_example" // string | 
    updateServiceAccountSpecification := *openapiclient.NewUpdateServiceAccountSpecification() // UpdateServiceAccountSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.ServiceAccountsApi.UpdateServiceAccount(context.Background(), id).UpdateServiceAccountSpecification(updateServiceAccountSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccountsApi.UpdateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccountsApi.UpdateServiceAccount`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateServiceAccountSpecification** | [**UpdateServiceAccountSpecification**](UpdateServiceAccountSpecification.md) |  | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

