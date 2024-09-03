# PrivateEndpointServices

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPrivateEndpointConnection**](PrivateEndpointServicesApi.md#AddPrivateEndpointConnection) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-connections | Add a connection to a cluster&#39;s private endpoint service.
[**AddPrivateEndpointTrustedOwner**](PrivateEndpointServicesApi.md#AddPrivateEndpointTrustedOwner) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-trusted-owners | Add a private endpoint trusted owner to a cluster
[**CreatePrivateEndpointServices**](PrivateEndpointServicesApi.md#CreatePrivateEndpointServices) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Create all PrivateEndpointServices for a cluster
[**DeletePrivateEndpointConnection**](PrivateEndpointServicesApi.md#DeletePrivateEndpointConnection) | **Delete** /api/v1/clusters/{cluster_id}/networking/private-endpoint-connections/{endpoint_id} | Delete a connection from a cluster&#39;s private endpoint service.
[**GetPrivateEndpointTrustedOwner**](PrivateEndpointServicesApi.md#GetPrivateEndpointTrustedOwner) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-trusted-owners/{owner_id} | Get a private endpoint trusted owner entry for a cluster
[**ListAwsEndpointConnections**](PrivateEndpointServicesApi.md#ListAwsEndpointConnections) | **Get** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections | List all AwsEndpointConnections for a cluster
[**ListPrivateEndpointConnections**](PrivateEndpointServicesApi.md#ListPrivateEndpointConnections) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-connections | List all connections to a cluster&#39;s private endpoint service.
[**ListPrivateEndpointServices**](PrivateEndpointServicesApi.md#ListPrivateEndpointServices) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | List all PrivateEndpointServices for a cluster
[**ListPrivateEndpointTrustedOwners**](PrivateEndpointServicesApi.md#ListPrivateEndpointTrustedOwners) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-trusted-owners | List all private endpoint trusted owners for a cluster
[**RemovePrivateEndpointTrustedOwner**](PrivateEndpointServicesApi.md#RemovePrivateEndpointTrustedOwner) | **Delete** /api/v1/clusters/{cluster_id}/networking/private-endpoint-trusted-owners/{owner_id} | Remove a private endpoint trusted owner from a cluster
[**SetAwsEndpointConnectionState**](PrivateEndpointServicesApi.md#SetAwsEndpointConnectionState) | **Patch** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections/{endpoint_id} | Set the AWS Endpoint Connection state



## AddPrivateEndpointConnection

> PrivateEndpointConnection AddPrivateEndpointConnection(ctx, clusterId).AddPrivateEndpointConnectionRequest(addPrivateEndpointConnectionRequest).Execute()

Add a connection to a cluster's private endpoint service.

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the id of the cluster to which the private endpoint connection will be added.
    addPrivateEndpointConnectionRequest := *openapiclient.NewAddPrivateEndpointConnectionRequest("EndpointId_example") // AddPrivateEndpointConnectionRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.AddPrivateEndpointConnection(context.Background(), clusterId).AddPrivateEndpointConnectionRequest(addPrivateEndpointConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.AddPrivateEndpointConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrivateEndpointConnection`: PrivateEndpointConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.AddPrivateEndpointConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the id of the cluster to which the private endpoint connection will be added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateEndpointConnection struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPrivateEndpointConnectionRequest** | [**AddPrivateEndpointConnectionRequest**](AddPrivateEndpointConnectionRequest.md) |  | 

### Return type

[**PrivateEndpointConnection**](PrivateEndpointConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## AddPrivateEndpointTrustedOwner

> AddPrivateEndpointTrustedOwnerResponse AddPrivateEndpointTrustedOwner(ctx, clusterId).AddPrivateEndpointTrustedOwnerRequest(addPrivateEndpointTrustedOwnerRequest).Execute()

Add a private endpoint trusted owner to a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.
    addPrivateEndpointTrustedOwnerRequest := *openapiclient.NewAddPrivateEndpointTrustedOwnerRequest("ExternalOwnerId_example", openapiclient.PrivateEndpointTrustedOwnerType.Type("AWS_ACCOUNT_ID")) // AddPrivateEndpointTrustedOwnerRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.AddPrivateEndpointTrustedOwner(context.Background(), clusterId).AddPrivateEndpointTrustedOwnerRequest(addPrivateEndpointTrustedOwnerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.AddPrivateEndpointTrustedOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPrivateEndpointTrustedOwner`: AddPrivateEndpointTrustedOwnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.AddPrivateEndpointTrustedOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPrivateEndpointTrustedOwner struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addPrivateEndpointTrustedOwnerRequest** | [**AddPrivateEndpointTrustedOwnerRequest**](AddPrivateEndpointTrustedOwnerRequest.md) |  | 

### Return type

[**AddPrivateEndpointTrustedOwnerResponse**](AddPrivateEndpointTrustedOwnerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePrivateEndpointServices

> PrivateEndpointServices CreatePrivateEndpointServices(ctx, clusterId).Execute()

Create all PrivateEndpointServices for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.CreatePrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.CreatePrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.CreatePrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrivateEndpointServices struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateEndpointServices**](PrivateEndpointServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeletePrivateEndpointConnection

> map[string]interface{} DeletePrivateEndpointConnection(ctx, clusterId, endpointId).Execute()

Delete a connection from a cluster's private endpoint service.

Remove a private endpoint from a service's trusted endpoints list. Caller should make sure to URL encode the endpoint_id before calling this method.

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the id of the cluster from which the private endpoint connection will be removed.
    endpointId := "endpointId_example" // string | endpoint_id is the id of the private endpoint associated with a cluster's private endpoint service.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.DeletePrivateEndpointConnection(context.Background(), clusterId, endpointId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.DeletePrivateEndpointConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePrivateEndpointConnection`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.DeletePrivateEndpointConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the id of the cluster from which the private endpoint connection will be removed. | 
**endpointId** | **string** | endpoint_id is the id of the private endpoint associated with a cluster&#39;s private endpoint service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrivateEndpointConnection struct via the builder pattern


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


## GetPrivateEndpointTrustedOwner

> GetPrivateEndpointTrustedOwnerResponse GetPrivateEndpointTrustedOwner(ctx, clusterId, ownerId).Execute()

Get a private endpoint trusted owner entry for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.
    ownerId := "ownerId_example" // string | owner_id corresponds to the UUID of the private endpoint trusted owner entry.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.GetPrivateEndpointTrustedOwner(context.Background(), clusterId, ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.GetPrivateEndpointTrustedOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateEndpointTrustedOwner`: GetPrivateEndpointTrustedOwnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.GetPrivateEndpointTrustedOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 
**ownerId** | **string** | owner_id corresponds to the UUID of the private endpoint trusted owner entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateEndpointTrustedOwner struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPrivateEndpointTrustedOwnerResponse**](GetPrivateEndpointTrustedOwnerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListAwsEndpointConnections

> AwsEndpointConnections ListAwsEndpointConnections(ctx, clusterId).Execute()

List all AwsEndpointConnections for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListAwsEndpointConnections(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListAwsEndpointConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsEndpointConnections`: AwsEndpointConnections
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListAwsEndpointConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsEndpointConnections struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsEndpointConnections**](AwsEndpointConnections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPrivateEndpointConnections

> PrivateEndpointConnections ListPrivateEndpointConnections(ctx, clusterId).Execute()

List all connections to a cluster's private endpoint service.

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListPrivateEndpointConnections(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListPrivateEndpointConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointConnections`: PrivateEndpointConnections
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListPrivateEndpointConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointConnections struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateEndpointConnections**](PrivateEndpointConnections.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPrivateEndpointServices

> PrivateEndpointServices ListPrivateEndpointServices(ctx, clusterId).Execute()

List all PrivateEndpointServices for a cluster

The internal_dns property from the regions field in the ListClusters response can be used to connect to PrivateEndpointServices.

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListPrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListPrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListPrivateEndpointServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointServices struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrivateEndpointServices**](PrivateEndpointServices.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPrivateEndpointTrustedOwners

> ListPrivateEndpointTrustedOwnersResponse ListPrivateEndpointTrustedOwners(ctx, clusterId).Execute()

List all private endpoint trusted owners for a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.ListPrivateEndpointTrustedOwners(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.ListPrivateEndpointTrustedOwners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointTrustedOwners`: ListPrivateEndpointTrustedOwnersResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.ListPrivateEndpointTrustedOwners`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPrivateEndpointTrustedOwners struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListPrivateEndpointTrustedOwnersResponse**](ListPrivateEndpointTrustedOwnersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## RemovePrivateEndpointTrustedOwner

> RemovePrivateEndpointTrustedOwnerResponse RemovePrivateEndpointTrustedOwner(ctx, clusterId, ownerId).Execute()

Remove a private endpoint trusted owner from a cluster

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.
    ownerId := "ownerId_example" // string | owner_id corresponds to the UUID of the private endpoint trusted owner entry.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.RemovePrivateEndpointTrustedOwner(context.Background(), clusterId, ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.RemovePrivateEndpointTrustedOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovePrivateEndpointTrustedOwner`: RemovePrivateEndpointTrustedOwnerResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.RemovePrivateEndpointTrustedOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 
**ownerId** | **string** | owner_id corresponds to the UUID of the private endpoint trusted owner entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePrivateEndpointTrustedOwner struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemovePrivateEndpointTrustedOwnerResponse**](RemovePrivateEndpointTrustedOwnerResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## SetAwsEndpointConnectionState

> AwsEndpointConnection SetAwsEndpointConnectionState(ctx, clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()

Set the AWS Endpoint Connection state

The "status" in the response does not reflect the latest post-update
status, but rather the status before the state is transitioned.

Can be used by the following roles assigned at the organization, folder or cluster scope:
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
    clusterId := "clusterId_example" // string | cluster_id is the ID for the cluster.
    endpointId := "endpointId_example" // string | endpoint_id is the ID for the VPC endpoint on the customer's side.
    setAwsEndpointConnectionStateRequest := *openapiclient.NewSetAwsEndpointConnectionStateRequest(openapiclient.SetAWSEndpointConnectionStatus.Type("AVAILABLE")) // SetAwsEndpointConnectionStateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.PrivateEndpointServicesApi.SetAwsEndpointConnectionState(context.Background(), clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateEndpointServicesApi.SetAwsEndpointConnectionState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAwsEndpointConnectionState`: AwsEndpointConnection
    fmt.Fprintf(os.Stdout, "Response from `PrivateEndpointServicesApi.SetAwsEndpointConnectionState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | cluster_id is the ID for the cluster. | 
**endpointId** | **string** | endpoint_id is the ID for the VPC endpoint on the customer&#39;s side. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAwsEndpointConnectionState struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setAwsEndpointConnectionStateRequest** | [**SetAwsEndpointConnectionStateRequest**](SetAwsEndpointConnectionStateRequest.md) |  | 

### Return type

[**AwsEndpointConnection**](AwsEndpointConnection.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

