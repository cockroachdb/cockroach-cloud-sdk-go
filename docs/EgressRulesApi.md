# EgressRules

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEgressRule**](EgressRulesApi.md#AddEgressRule) | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules | Add an egress rule
[**DeleteEgressRule**](EgressRulesApi.md#DeleteEgressRule) | **Delete** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Delete an existing egress rule
[**EditEgressRule**](EgressRulesApi.md#EditEgressRule) | **Patch** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Edit an existing egress rule
[**GetEgressRule**](EgressRulesApi.md#GetEgressRule) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Get an existing egress rule
[**ListEgressRules**](EgressRulesApi.md#ListEgressRules) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules | List all egress rules associates with a cluster
[**SetEgressTrafficPolicy**](EgressRulesApi.md#SetEgressTrafficPolicy) | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules/egress-traffic-policy | Outbound traffic management



## AddEgressRule

> AddEgressRuleResponse AddEgressRule(ctx, clusterId).AddEgressRuleRequest(addEgressRuleRequest).Execute()

Add an egress rule

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id identifies the cluster to which this egress rule applies.
    addEgressRuleRequest := *openapiclient.NewAddEgressRuleRequest("Description_example", "Destination_example", "Name_example", "Type_example") // AddEgressRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.AddEgressRule(context.Background(), clusterId).AddEgressRuleRequest(addEgressRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.AddEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEgressRule`: AddEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.AddEgressRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the cluster to which this egress rule applies. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEgressRule struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addEgressRuleRequest** | [**AddEgressRuleRequest**](AddEgressRuleRequest.md) |  | 

### Return type

[**AddEgressRuleResponse**](AddEgressRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteEgressRule

> DeleteEgressRuleResponse DeleteEgressRule(ctx, clusterId, ruleId).IdempotencyKey(idempotencyKey).Execute()

Delete an existing egress rule

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id uniquely identifies the cluster owning the egress rule.
    ruleId := "ruleId_example" // string | rule_id is the UUID of an existing egress rule. This field is required.
    idempotencyKey := "idempotencyKey_example" // string | idempotency_key uniquely identifies this request. If not set, it will be set by the server. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.DeleteEgressRule(context.Background(), clusterId, ruleId).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.DeleteEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEgressRule`: DeleteEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.DeleteEgressRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id uniquely identifies the cluster owning the egress rule. | 
**ruleId** | **string** | rule_id is the UUID of an existing egress rule. This field is required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEgressRule struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **idempotencyKey** | **string** | idempotency_key uniquely identifies this request. If not set, it will be set by the server. | 

### Return type

[**DeleteEgressRuleResponse**](DeleteEgressRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## EditEgressRule

> EditEgressRuleResponse EditEgressRule(ctx, clusterId, ruleId).EditEgressRuleRequest(editEgressRuleRequest).Execute()

Edit an existing egress rule

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id uniquely identifies the cluster owning the egress rule.
    ruleId := "ruleId_example" // string | rule_id is the UUID of an existing egress rule. This field is required.
    editEgressRuleRequest := *openapiclient.NewEditEgressRuleRequest() // EditEgressRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.EditEgressRule(context.Background(), clusterId, ruleId).EditEgressRuleRequest(editEgressRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.EditEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEgressRule`: EditEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.EditEgressRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id uniquely identifies the cluster owning the egress rule. | 
**ruleId** | **string** | rule_id is the UUID of an existing egress rule. This field is required. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditEgressRule struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **editEgressRuleRequest** | [**EditEgressRuleRequest**](EditEgressRuleRequest.md) |  | 

### Return type

[**EditEgressRuleResponse**](EditEgressRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetEgressRule

> GetEgressRuleResponse GetEgressRule(ctx, clusterId, ruleId).Execute()

Get an existing egress rule

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id uniquely identifies the cluster owning the egress rule.
    ruleId := "ruleId_example" // string | rule_id is the UUID of an existing egress rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.GetEgressRule(context.Background(), clusterId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.GetEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEgressRule`: GetEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.GetEgressRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id uniquely identifies the cluster owning the egress rule. | 
**ruleId** | **string** | rule_id is the UUID of an existing egress rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEgressRule struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetEgressRuleResponse**](GetEgressRuleResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEgressRules

> ListEgressRulesResponse ListEgressRules(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all egress rules associates with a cluster

Sort order: Name

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id identifies the CockroachDB cluster owning the set of returned egress rules.
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.ListEgressRules(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.ListEgressRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEgressRules`: ListEgressRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.ListEgressRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the CockroachDB cluster owning the set of returned egress rules. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEgressRules struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListEgressRulesResponse**](ListEgressRulesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetEgressTrafficPolicy

> map[string]interface{} SetEgressTrafficPolicy(ctx, clusterId).SetEgressTrafficPolicyRequest(setEgressTrafficPolicyRequest).Execute()

Outbound traffic management

Can be used by the following roles assigned at the organization or cluster scope:
- CLUSTER_ADMIN
- CLUSTER_OPERATOR_WRITER


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
    clusterId := "clusterId_example" // string | cluster_id identifies the cluster whose egress policy will be updated.
    setEgressTrafficPolicyRequest := *openapiclient.NewSetEgressTrafficPolicyRequest(false) // SetEgressTrafficPolicyRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.EgressRulesApi.SetEgressTrafficPolicy(context.Background(), clusterId).SetEgressTrafficPolicyRequest(setEgressTrafficPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EgressRulesApi.SetEgressTrafficPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEgressTrafficPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `EgressRulesApi.SetEgressTrafficPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id identifies the cluster whose egress policy will be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetEgressTrafficPolicy struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setEgressTrafficPolicyRequest** | [**SetEgressTrafficPolicyRequest**](SetEgressTrafficPolicyRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

