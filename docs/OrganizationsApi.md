# Organizations

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrganizationInfo**](OrganizationsApi.md#GetOrganizationInfo) | **Get** /api/v1/organization | Get information about the caller&#39;s organization



## GetOrganizationInfo

> Organization GetOrganizationInfo(ctx).Execute()

Get information about the caller's organization

Can be used by the following roles assigned at the organization scope:
- ORG_ADMIN
- ORG_MEMBER


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
    resp, r, err := api_client.OrganizationsApi.GetOrganizationInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsApi.GetOrganizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInfo`: Organization
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsApi.GetOrganizationInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationInfo struct via the builder pattern


### Return type

[**Organization**](Organization.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

