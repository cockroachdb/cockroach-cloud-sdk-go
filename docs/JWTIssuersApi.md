# JWTIssuers

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddJWTIssuer**](JWTIssuersApi.md#AddJWTIssuer) | **Post** /api/v1/jwt-issuers | Add a JWT Issuer
[**DeleteJWTIssuer**](JWTIssuersApi.md#DeleteJWTIssuer) | **Delete** /api/v1/jwt-issuers/{id} | Delete a JWT Issuer
[**GetJWTIssuer**](JWTIssuersApi.md#GetJWTIssuer) | **Get** /api/v1/jwt-issuers/{id} | Get a JWT Issuer
[**ListJWTIssuers**](JWTIssuersApi.md#ListJWTIssuers) | **Get** /api/v1/jwt-issuers | List all JWT Issuers
[**UpdateJWTIssuer**](JWTIssuersApi.md#UpdateJWTIssuer) | **Patch** /api/v1/jwt-issuers/{id} | Update a JWT Issuer



## AddJWTIssuer

> JWTIssuer AddJWTIssuer(ctx).AddJWTIssuerRequest(addJWTIssuerRequest).Execute()

Add a JWT Issuer

Registers a JWT Issuer with the CockroachDB Cloud to allow verifying JWTs during API authentication

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
    addJWTIssuerRequest := *openapiclient.NewAddJWTIssuerRequest("1234567890abcd", "https://jwt-issuer.example.com") // AddJWTIssuerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.JWTIssuersApi.AddJWTIssuer(context.Background()).AddJWTIssuerRequest(addJWTIssuerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWTIssuersApi.AddJWTIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddJWTIssuer`: JWTIssuer
    fmt.Fprintf(os.Stdout, "Response from `JWTIssuersApi.AddJWTIssuer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddJWTIssuer struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addJWTIssuerRequest** | [**AddJWTIssuerRequest**](AddJWTIssuerRequest.md) |  | 

### Return type

[**JWTIssuer**](JWTIssuer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteJWTIssuer

> JWTIssuer DeleteJWTIssuer(ctx, id).Execute()

Delete a JWT Issuer

Deletes the JWT Issuer configuration

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the JWT Issuer resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.JWTIssuersApi.DeleteJWTIssuer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWTIssuersApi.DeleteJWTIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteJWTIssuer`: JWTIssuer
    fmt.Fprintf(os.Stdout, "Response from `JWTIssuersApi.DeleteJWTIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the JWT Issuer resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJWTIssuer struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JWTIssuer**](JWTIssuer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetJWTIssuer

> JWTIssuer GetJWTIssuer(ctx, id).Execute()

Get a JWT Issuer

Retrieves the JWT Issuer configuration

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the JWT Issuer resource

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.JWTIssuersApi.GetJWTIssuer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWTIssuersApi.GetJWTIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJWTIssuer`: JWTIssuer
    fmt.Fprintf(os.Stdout, "Response from `JWTIssuersApi.GetJWTIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the JWT Issuer resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJWTIssuer struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JWTIssuer**](JWTIssuer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListJWTIssuers

> ListJWTIssuersResponse ListJWTIssuers(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all JWT Issuers

Lists all the JWT Issuer configurations registered for the CockroachDB Cloud organization

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN


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
    resp, r, err := api_client.JWTIssuersApi.ListJWTIssuers(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWTIssuersApi.ListJWTIssuers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListJWTIssuers`: ListJWTIssuersResponse
    fmt.Fprintf(os.Stdout, "Response from `JWTIssuersApi.ListJWTIssuers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListJWTIssuers struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListJWTIssuersResponse**](ListJWTIssuersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateJWTIssuer

> JWTIssuer UpdateJWTIssuer(ctx, id).UpdateJWTIssuerRequest(updateJWTIssuerRequest).Execute()

Update a JWT Issuer

Updates the JWT Issuer configuration

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique identifier of the JWT Issuer resource
    updateJWTIssuerRequest := *openapiclient.NewUpdateJWTIssuerRequest() // UpdateJWTIssuerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.JWTIssuersApi.UpdateJWTIssuer(context.Background(), id).UpdateJWTIssuerRequest(updateJWTIssuerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JWTIssuersApi.UpdateJWTIssuer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJWTIssuer`: JWTIssuer
    fmt.Fprintf(os.Stdout, "Response from `JWTIssuersApi.UpdateJWTIssuer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The unique identifier of the JWT Issuer resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJWTIssuer struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateJWTIssuerRequest** | [**UpdateJWTIssuerRequest**](UpdateJWTIssuerRequest.md) |  | 

### Return type

[**JWTIssuer**](JWTIssuer.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

