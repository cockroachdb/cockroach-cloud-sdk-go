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

> ListInvoicesResponse ListInvoices(ctx).Status(status).Execute()

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
    status := "status_example" // string | Filters the response to only include invoices with the specified status. This will be sent as a query parameter on the GET request. If not specified, both Finalized and Draft invoices will be included. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.BillingApi.ListInvoices(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListInvoices`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.

### Other Parameters

Optional parameters can be passed through a pointer to the ListInvoicesOptions struct.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filters the response to only include invoices with the specified status. This will be sent as a query parameter on the GET request. If not specified, both Finalized and Draft invoices will be included. | 

### Return type

[**ListInvoicesResponse**](ListInvoicesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

