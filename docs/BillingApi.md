# Billing

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInvoice**](BillingApi.md#GetInvoice) | **Get** /api/v1/invoices/{invoice_id} | Get a specific invoice for an organization
[**ListInvoices**](BillingApi.md#ListInvoices) | **Get** /api/v1/invoices | List invoices for a given organization



## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Get a specific invoice for an organization

Can be used by the following roles assigned at the organization scope:
- BILLING_COORDINATOR
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
    invoiceId := "invoiceId_example" // string | invoice_id is the unique ID representing the invoice. invoice_id is used to retrieve a specific billing period's invoice.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BillingApi.GetInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **string** | invoice_id is the unique ID representing the invoice. invoice_id is used to retrieve a specific billing period&#39;s invoice. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoice struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Invoice**](Invoice.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListInvoices

> ListInvoicesResponse ListInvoices(ctx).Execute()

List invoices for a given organization

Sort order: invoice start date ascending

Can be used by the following roles assigned at the organization scope:
- BILLING_COORDINATOR
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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BillingApi.ListInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoices struct via the builder pattern


### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

