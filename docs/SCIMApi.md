# SCIM

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroup**](SCIMApi.md#CreateGroup) | **Post** /api/scim/v2/Groups | Create a group
[**CreateUser**](SCIMApi.md#CreateUser) | **Post** /api/scim/v2/Users | Create a user
[**DeleteGroup**](SCIMApi.md#DeleteGroup) | **Delete** /api/scim/v2/Groups/{id} | Delete a group based on ID
[**DeleteUser**](SCIMApi.md#DeleteUser) | **Delete** /api/scim/v2/Users/{id} | Delete a user based on ID
[**GetGroup**](SCIMApi.md#GetGroup) | **Get** /api/scim/v2/Groups/{id} | Get a group by ID
[**GetGroup2**](SCIMApi.md#GetGroup2) | **Put** /api/scim/v2/Groups/{id}/.search | Search a group by ID (Deprecated)
[**GetGroups**](SCIMApi.md#GetGroups) | **Get** /api/scim/v2/Groups | List groups
[**GetGroups2**](SCIMApi.md#GetGroups2) | **Put** /api/scim/v2/Groups/.search | Search groups (Deprecated)
[**GetResourceType**](SCIMApi.md#GetResourceType) | **Get** /api/scim/v2/ResourceTypes/{resourceId} | Get a SCIM resource type by ID
[**GetResourceTypes**](SCIMApi.md#GetResourceTypes) | **Get** /api/scim/v2/ResourceTypes | List the SCIM resource types
[**GetSchema**](SCIMApi.md#GetSchema) | **Get** /api/scim/v2/Schemas/{schemaId} | Get a SCIM schema by ID
[**GetSchemas**](SCIMApi.md#GetSchemas) | **Get** /api/scim/v2/Schemas | List the SCIM schemas
[**GetServiceProviderConfig**](SCIMApi.md#GetServiceProviderConfig) | **Get** /api/scim/v2/ServiceProviderConfig | Return the SCIM Service Provider configuration
[**GetUser**](SCIMApi.md#GetUser) | **Get** /api/scim/v2/Users/{id} | Get a user by ID
[**GetUser2**](SCIMApi.md#GetUser2) | **Put** /api/scim/v2/Users/{id}/.search | Search for a user by ID (Deprecated)
[**GetUsers**](SCIMApi.md#GetUsers) | **Get** /api/scim/v2/Users | List Users
[**GetUsers2**](SCIMApi.md#GetUsers2) | **Put** /api/scim/v2/Users/.search | Search User (Deprecated)
[**SearchGroup**](SCIMApi.md#SearchGroup) | **Post** /api/scim/v2/Groups/{id}/.search | Search a group by ID
[**SearchGroups**](SCIMApi.md#SearchGroups) | **Post** /api/scim/v2/Groups/.search | Search groups
[**SearchUser**](SCIMApi.md#SearchUser) | **Post** /api/scim/v2/Users/{id}/.search | Search for a user by ID
[**SearchUsers**](SCIMApi.md#SearchUsers) | **Post** /api/scim/v2/Users/.search | Search Users
[**UpdateGroup**](SCIMApi.md#UpdateGroup) | **Put** /api/scim/v2/Groups/{id} | Update a group by supplying all values of the user object
[**UpdateUser**](SCIMApi.md#UpdateUser) | **Put** /api/scim/v2/Users/{id} | Update a user by supplying all values of the user object



## CreateGroup

> ScimGroup CreateGroup(ctx).CreateGroupRequest(createGroupRequest).Execute()

Create a group

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
    createGroupRequest := *openapiclient.NewCreateGroupRequest("DisplayName_example") // CreateGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.CreateGroup(context.Background()).CreateGroupRequest(createGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createGroupRequest** | [**CreateGroupRequest**](CreateGroupRequest.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateUser

> ScimUser CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Create a user

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
    createUserRequest := *openapiclient.NewCreateUserRequest(false, "DisplayName_example", []openapiclient.ScimEmail{*openapiclient.NewScimEmail(false, "Value_example")}, *openapiclient.NewScimName()) // CreateUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteGroup

> map[string]interface{} DeleteGroup(ctx, id).Execute()

Delete a group based on ID

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.DeleteGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.DeleteGroup`: %v\n", resp)
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


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteUser

> map[string]interface{} DeleteUser(ctx, id).Execute()

Delete a user based on ID

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.DeleteUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.DeleteUser`: %v\n", resp)
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


### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroup

> ScimGroup GetGroup(ctx, id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Get a group by ID

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
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetGroup(context.Background(), id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the GetGroupOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroup2

> ScimGroup GetGroup2(ctx, id).GetGroupRequest(getGroupRequest).Execute()

Search a group by ID (Deprecated)

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
    getGroupRequest := *openapiclient.NewGetGroupRequest() // GetGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetGroup2(context.Background(), id).GetGroupRequest(getGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetGroup2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetGroup2`: %v\n", resp)
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

 **getGroupRequest** | [**GetGroupRequest**](GetGroupRequest.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroups

> GetGroupsResponse GetGroups(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Filter(filter).Execute()

List groups

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
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetGroups(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: GetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the GetGroupsOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 
 **filter** | **string** |  | 

### Return type

[**GetGroupsResponse**](GetGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroups2

> GetGroupsResponse GetGroups2(ctx).GetGroupsRequest(getGroupsRequest).Execute()

Search groups (Deprecated)

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
    getGroupsRequest := *openapiclient.NewGetGroupsRequest() // GetGroupsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetGroups2(context.Background()).GetGroupsRequest(getGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetGroups2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups2`: GetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetGroups2`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getGroupsRequest** | [**GetGroupsRequest**](GetGroupsRequest.md) |  | 

### Return type

[**GetGroupsResponse**](GetGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetResourceType

> ScimResourceType GetResourceType(ctx, resourceId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Get a SCIM resource type by ID

This endpoint may be used by any member of the organization.

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
    resourceId := "resourceId_example" // string | 
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetResourceType(context.Background(), resourceId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceType`: ScimResourceType
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetResourceType`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the GetResourceTypeOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**ScimResourceType**](ScimResourceType.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetResourceTypes

> GetResourceTypesResponse GetResourceTypes(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

List the SCIM resource types

This endpoint may be used by any member of the organization.

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
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetResourceTypes(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetResourceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceTypes`: GetResourceTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetResourceTypes`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the GetResourceTypesOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**GetResourceTypesResponse**](GetResourceTypesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSchema

> ScimSchema GetSchema(ctx, schemaId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Get a SCIM schema by ID

This endpoint may be used by any member of the organization.

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
    schemaId := "schemaId_example" // string | 
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetSchema(context.Background(), schemaId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: ScimSchema
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the GetSchemaOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**ScimSchema**](ScimSchema.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSchemas

> GetSchemasResponse GetSchemas(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

List the SCIM schemas

This endpoint may be used by any member of the organization.

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
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetSchemas(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemas`: GetSchemasResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetSchemas`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the GetSchemasOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**GetSchemasResponse**](GetSchemasResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetServiceProviderConfig

> GetServiceProviderConfigResponse GetServiceProviderConfig(ctx).Execute()

Return the SCIM Service Provider configuration

This endpoint may be used by any member of the organization.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetServiceProviderConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetServiceProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceProviderConfig`: GetServiceProviderConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetServiceProviderConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


### Return type

[**GetServiceProviderConfigResponse**](GetServiceProviderConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetUser

> ScimUser GetUser(ctx, id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Get a user by ID

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
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetUser(context.Background(), id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetUser`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Optional parameters can be passed through a pointer to the GetUserOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetUser2

> ScimUser GetUser2(ctx, id).GetUserRequest(getUserRequest).Execute()

Search for a user by ID (Deprecated)

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
    getUserRequest := *openapiclient.NewGetUserRequest() // GetUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetUser2(context.Background(), id).GetUserRequest(getUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser2`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetUser2`: %v\n", resp)
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

 **getUserRequest** | [**GetUserRequest**](GetUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetUsers

> GetUsersResponse GetUsers(ctx).Filter(filter).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

List Users

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
    filter := "filter_example" // string |  (optional)
    attributes := "attributes_example" // string |  (optional)
    excludedAttributes := "excludedAttributes_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetUsers(context.Background()).Filter(filter).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: GetUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the GetUsersOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 
 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**GetUsersResponse**](GetUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetUsers2

> GetUsersResponse GetUsers2(ctx).GetUsersRequest(getUsersRequest).Execute()

Search User (Deprecated)

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
    getUsersRequest := *openapiclient.NewGetUsersRequest() // GetUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.GetUsers2(context.Background()).GetUsersRequest(getUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.GetUsers2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers2`: GetUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.GetUsers2`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUsersRequest** | [**GetUsersRequest**](GetUsersRequest.md) |  | 

### Return type

[**GetUsersResponse**](GetUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SearchGroup

> ScimGroup SearchGroup(ctx, id).GetGroupRequest(getGroupRequest).Execute()

Search a group by ID

Similar to GetGroup however search parameters are passed via the POST body. See https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3 for more details.

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
    getGroupRequest := *openapiclient.NewGetGroupRequest() // GetGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.SearchGroup(context.Background(), id).GetGroupRequest(getGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.SearchGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.SearchGroup`: %v\n", resp)
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

 **getGroupRequest** | [**GetGroupRequest**](GetGroupRequest.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SearchGroups

> GetGroupsResponse SearchGroups(ctx).GetGroupsRequest(getGroupsRequest).Execute()

Search groups

Similar to GetGroups however search parameters are passed via the POST body. See https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3 for more details.

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
    getGroupsRequest := *openapiclient.NewGetGroupsRequest() // GetGroupsRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.SearchGroups(context.Background()).GetGroupsRequest(getGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.SearchGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchGroups`: GetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.SearchGroups`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getGroupsRequest** | [**GetGroupsRequest**](GetGroupsRequest.md) |  | 

### Return type

[**GetGroupsResponse**](GetGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SearchUser

> ScimUser SearchUser(ctx, id).GetUserRequest(getUserRequest).Execute()

Search for a user by ID

Similar to GetUser however search parameters are passed via the POST body. See https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3 for more details.

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
    getUserRequest := *openapiclient.NewGetUserRequest() // GetUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.SearchUser(context.Background(), id).GetUserRequest(getUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.SearchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.SearchUser`: %v\n", resp)
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

 **getUserRequest** | [**GetUserRequest**](GetUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SearchUsers

> GetUsersResponse SearchUsers(ctx).GetUsersRequest(getUsersRequest).Execute()

Search Users

Similar to GetUsers however search parameters are passed via the POST body. See https://www.rfc-editor.org/rfc/rfc7644.html#section-3.4.3 for more details.

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
    getUsersRequest := *openapiclient.NewGetUsersRequest() // GetUsersRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.SearchUsers(context.Background()).GetUsersRequest(getUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.SearchUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsers`: GetUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.SearchUsers`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getUsersRequest** | [**GetUsersRequest**](GetUsersRequest.md) |  | 

### Return type

[**GetUsersResponse**](GetUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateGroup

> ScimGroup UpdateGroup(ctx, id).UpdateGroupRequest(updateGroupRequest).Execute()

Update a group by supplying all values of the user object

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
    updateGroupRequest := *openapiclient.NewUpdateGroupRequest("DisplayName_example") // UpdateGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.UpdateGroup(context.Background(), id).UpdateGroupRequest(updateGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.UpdateGroup`: %v\n", resp)
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

 **updateGroupRequest** | [**UpdateGroupRequest**](UpdateGroupRequest.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateUser

> ScimUser UpdateUser(ctx, id).UpdateUserRequest(updateUserRequest).Execute()

Update a user by supplying all values of the user object

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
    updateUserRequest := *openapiclient.NewUpdateUserRequest(false) // UpdateUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.SCIMApi.UpdateUser(context.Background(), id).UpdateUserRequest(updateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SCIMApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `SCIMApi.UpdateUser`: %v\n", resp)
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

 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

