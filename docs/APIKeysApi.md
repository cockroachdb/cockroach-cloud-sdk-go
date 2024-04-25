# APIKeys

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKey**](APIKeysApi.md#CreateApiKey) | **Post** /api/v1/api-keys | Create a new API Key
[**DeleteApiKey**](APIKeysApi.md#DeleteApiKey) | **Delete** /api/v1/api-keys/{id} | Delete an API Key
[**GetApiKey**](APIKeysApi.md#GetApiKey) | **Get** /api/v1/api-keys/{id} | Get an API Key by ID
[**ListApiKeys**](APIKeysApi.md#ListApiKeys) | **Get** /api/v1/api-keys | List API Keys
[**UpdateApiKey**](APIKeysApi.md#UpdateApiKey) | **Patch** /api/v1/api-keys/{id} | Update an API Key



## CreateApiKey

> CreateApiKeyResponse CreateApiKey(ctx).CreateApiKeyRequest(createApiKeyRequest).Execute()

Create a new API Key

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
    createApiKeyRequest := *openapiclient.NewCreateApiKeyRequest("example name", "1234abcd-1234-1234-abcd-12345678abcd") // CreateApiKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.APIKeysApi.CreateApiKey(context.Background()).CreateApiKeyRequest(createApiKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.CreateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKey`: CreateApiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.CreateApiKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKey struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApiKeyRequest** | [**CreateApiKeyRequest**](CreateApiKeyRequest.md) |  | 

### Return type

[**CreateApiKeyResponse**](CreateApiKeyResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteApiKey

> ApiKey DeleteApiKey(ctx, id).Execute()

Delete an API Key

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
    id := "id_example" // string | The ID of the api key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.APIKeysApi.DeleteApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.DeleteApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteApiKey`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.DeleteApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKey struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetApiKey

> ApiKey GetApiKey(ctx, id).Execute()

Get an API Key by ID

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
    id := "id_example" // string | the ID of the api key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.APIKeysApi.GetApiKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.GetApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKey`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the ID of the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKey struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListApiKeys

> ListApiKeysResponse ListApiKeys(ctx).ServiceAccountId(serviceAccountId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List API Keys

Sort order: created_at

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
    serviceAccountId := "serviceAccountId_example" // string | Optional filter to limit the response to include only api keys for a specific service account. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.APIKeysApi.ListApiKeys(context.Background()).ServiceAccountId(serviceAccountId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.ListApiKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApiKeys`: ListApiKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.ListApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApiKeys struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountId** | **string** | Optional filter to limit the response to include only api keys for a specific service account. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListApiKeysResponse**](ListApiKeysResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateApiKey

> ApiKey UpdateApiKey(ctx, id).UpdateApiKeySpecification(updateApiKeySpecification).Execute()

Update an API Key

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
    id := "id_example" // string | The ID of the api key.
    updateApiKeySpecification := *openapiclient.NewUpdateApiKeySpecification() // UpdateApiKeySpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.APIKeysApi.UpdateApiKey(context.Background(), id).UpdateApiKeySpecification(updateApiKeySpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APIKeysApi.UpdateApiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKey`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `APIKeysApi.UpdateApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the api key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKey struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateApiKeySpecification** | [**UpdateApiKeySpecification**](UpdateApiKeySpecification.md) |  | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

