# MetricExport

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCloudWatchMetricExport**](MetricExportApi.md#DeleteCloudWatchMetricExport) | **Delete** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Delete the CloudWatch Metric Export configuration for a cluster
[**DeleteDatadogMetricExport**](MetricExportApi.md#DeleteDatadogMetricExport) | **Delete** /api/v1/clusters/{cluster_id}/metricexport/datadog | Delete the Datadog Metric Export configuration for a cluster
[**EnableCloudWatchMetricExport**](MetricExportApi.md#EnableCloudWatchMetricExport) | **Post** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Create or update the CloudWatch Metric Export configuration for a cluster
[**EnableDatadogMetricExport**](MetricExportApi.md#EnableDatadogMetricExport) | **Post** /api/v1/clusters/{cluster_id}/metricexport/datadog | Create or update the Datadog Metric Export configuration for a cluster
[**GetCloudWatchMetricExportInfo**](MetricExportApi.md#GetCloudWatchMetricExportInfo) | **Get** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Get the CloudWatch Metric Export configuration for a cluster
[**GetDatadogMetricExportInfo**](MetricExportApi.md#GetDatadogMetricExportInfo) | **Get** /api/v1/clusters/{cluster_id}/metricexport/datadog | Get the Datadog Metric Export configuration for a cluster



## DeleteCloudWatchMetricExport

> DeleteMetricExportResponse DeleteCloudWatchMetricExport(ctx, clusterId).Execute()

Delete the CloudWatch Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.DeleteCloudWatchMetricExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.DeleteCloudWatchMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCloudWatchMetricExport`: DeleteMetricExportResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.DeleteCloudWatchMetricExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudWatchMetricExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMetricExportResponse**](DeleteMetricExportResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatadogMetricExport

> DeleteMetricExportResponse DeleteDatadogMetricExport(ctx, clusterId).Execute()

Delete the Datadog Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.DeleteDatadogMetricExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.DeleteDatadogMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatadogMetricExport`: DeleteMetricExportResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.DeleteDatadogMetricExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatadogMetricExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMetricExportResponse**](DeleteMetricExportResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCloudWatchMetricExport

> CloudWatchMetricExportInfo EnableCloudWatchMetricExport(ctx, clusterId).EnableCloudWatchMetricExportRequest(enableCloudWatchMetricExportRequest).Execute()

Create or update the CloudWatch Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 
    enableCloudWatchMetricExportRequest := *openapiclient.NewEnableCloudWatchMetricExportRequest("RoleArn_example") // EnableCloudWatchMetricExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.EnableCloudWatchMetricExport(context.Background(), clusterId).EnableCloudWatchMetricExportRequest(enableCloudWatchMetricExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.EnableCloudWatchMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCloudWatchMetricExport`: CloudWatchMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.EnableCloudWatchMetricExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCloudWatchMetricExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableCloudWatchMetricExportRequest** | [**EnableCloudWatchMetricExportRequest**](EnableCloudWatchMetricExportRequest.md) |  | 

### Return type

[**CloudWatchMetricExportInfo**](CloudWatchMetricExportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableDatadogMetricExport

> DatadogMetricExportInfo EnableDatadogMetricExport(ctx, clusterId).EnableDatadogMetricExportRequest(enableDatadogMetricExportRequest).Execute()

Create or update the Datadog Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 
    enableDatadogMetricExportRequest := *openapiclient.NewEnableDatadogMetricExportRequest("ApiKey_example", openapiclient.DatadogSite.Type("US1")) // EnableDatadogMetricExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.EnableDatadogMetricExport(context.Background(), clusterId).EnableDatadogMetricExportRequest(enableDatadogMetricExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.EnableDatadogMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableDatadogMetricExport`: DatadogMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.EnableDatadogMetricExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableDatadogMetricExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableDatadogMetricExportRequest** | [**EnableDatadogMetricExportRequest**](EnableDatadogMetricExportRequest.md) |  | 

### Return type

[**DatadogMetricExportInfo**](DatadogMetricExportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCloudWatchMetricExportInfo

> CloudWatchMetricExportInfo GetCloudWatchMetricExportInfo(ctx, clusterId).Execute()

Get the CloudWatch Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.GetCloudWatchMetricExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.GetCloudWatchMetricExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudWatchMetricExportInfo`: CloudWatchMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.GetCloudWatchMetricExportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudWatchMetricExportInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudWatchMetricExportInfo**](CloudWatchMetricExportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatadogMetricExportInfo

> DatadogMetricExportInfo GetDatadogMetricExportInfo(ctx, clusterId).Execute()

Get the Datadog Metric Export configuration for a cluster

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
    clusterId := "clusterId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.MetricExportApi.GetDatadogMetricExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricExportApi.GetDatadogMetricExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatadogMetricExportInfo`: DatadogMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `MetricExportApi.GetDatadogMetricExportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatadogMetricExportInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DatadogMetricExportInfo**](DatadogMetricExportInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

