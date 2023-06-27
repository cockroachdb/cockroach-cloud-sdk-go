# AuditLogs

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditLogs**](AuditLogsApi.md#ListAuditLogs) | **Get** /api/v1/auditlogevents | List audit logs



## ListAuditLogs

> ListAuditLogsResponse ListAuditLogs(ctx).StartingFrom(startingFrom).SortOrder(sortOrder).Limit(limit).Execute()

List audit logs



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
    startingFrom := time.Now() // time.Time | starting_from is the (exclusive) timestamp from which log entries will be returned in the response based on their created_at time, respecting the sort order specified in pagination. If unset, the default will be the current time if results are returned in descending order and the beginning of time if results are in ascending order. (optional)
    sortOrder := "sortOrder_example" // string | sort_order is the direction of pagination, with starting_from as the start point. If unset, the default is ascending order.   - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)
    limit := int32(56) // int32 | limit is the number of entries requested in the response. Note that the response may still contain slightly more results, since the response will always contain every entry at a particular timestamp. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.AuditLogsApi.ListAuditLogs(context.Background()).StartingFrom(startingFrom).SortOrder(sortOrder).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsApi.ListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditLogs`: ListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `AuditLogsApi.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogs struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startingFrom** | **time.Time** | starting_from is the (exclusive) timestamp from which log entries will be returned in the response based on their created_at time, respecting the sort order specified in pagination. If unset, the default will be the current time if results are returned in descending order and the beginning of time if results are in ascending order. | 
 **sortOrder** | **string** | sort_order is the direction of pagination, with starting_from as the start point. If unset, the default is ascending order.   - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 
 **limit** | **int32** | limit is the number of entries requested in the response. Note that the response may still contain slightly more results, since the response will always contain every entry at a particular timestamp. | 

### Return type

[**ListAuditLogsResponse**](ListAuditLogsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

