# \CockroachCloudApi

All URIs are relative to *https://cockroachlabs.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistEntry**](CockroachCloudApi.md#AddAllowlistEntry) | **Post** /api/v1/clusters/{cluster_id}/networking/allowlist | Add a new CIDR address to the IP allowlist
[**AddAllowlistEntry2**](CockroachCloudApi.md#AddAllowlistEntry2) | **Put** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Add a new CIDR address to the IP allowlist
[**AddEgressRule**](CockroachCloudApi.md#AddEgressRule) | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules | Add an egress rule
[**AddUserToRole**](CockroachCloudApi.md#AddUserToRole) | **Post** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Adds the user to the given role
[**CreateCluster**](CockroachCloudApi.md#CreateCluster) | **Post** /api/v1/clusters | Create and initialize a new cluster
[**CreateDatabase**](CockroachCloudApi.md#CreateDatabase) | **Post** /api/v1/clusters/{cluster_id}/databases | Create a new database
[**CreateGroup**](CockroachCloudApi.md#CreateGroup) | **Post** /api/scim/v2/Groups | Creates a group
[**CreatePrivateEndpointServices**](CockroachCloudApi.md#CreatePrivateEndpointServices) | **Post** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Creates all PrivateEndpointServices for a given cluster
[**CreateSQLUser**](CockroachCloudApi.md#CreateSQLUser) | **Post** /api/v1/clusters/{cluster_id}/sql-users | Create a new SQL user
[**CreateUser**](CockroachCloudApi.md#CreateUser) | **Post** /api/scim/v2/Users | Creates a user
[**DeleteAllowlistEntry**](CockroachCloudApi.md#DeleteAllowlistEntry) | **Delete** /api/v1/clusters/{cluster_id}/networking/allowlist/{cidr_ip}/{cidr_mask} | Delete an IP allowlist entry
[**DeleteClientCACert**](CockroachCloudApi.md#DeleteClientCACert) | **Delete** /api/v1/clusters/{cluster_id}/client-ca-cert | Delete Client CA Cert for a cluster
[**DeleteCloudWatchMetricExport**](CockroachCloudApi.md#DeleteCloudWatchMetricExport) | **Delete** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Delete the CloudWatch Metric Export configuration for a cluster
[**DeleteCluster**](CockroachCloudApi.md#DeleteCluster) | **Delete** /api/v1/clusters/{cluster_id} | Delete a cluster and all of its data
[**DeleteDatabase**](CockroachCloudApi.md#DeleteDatabase) | **Delete** /api/v1/clusters/{cluster_id}/databases/{name} | Delete a database
[**DeleteDatadogMetricExport**](CockroachCloudApi.md#DeleteDatadogMetricExport) | **Delete** /api/v1/clusters/{cluster_id}/metricexport/datadog | Delete the Datadog Metric Export configuration for a cluster
[**DeleteEgressRule**](CockroachCloudApi.md#DeleteEgressRule) | **Delete** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Delete an existing egress rule
[**DeleteGroup**](CockroachCloudApi.md#DeleteGroup) | **Delete** /api/scim/v2/Groups/{id} | Deletes a group based on id
[**DeleteLogExport**](CockroachCloudApi.md#DeleteLogExport) | **Delete** /api/v1/clusters/{cluster_id}/logexport | Delete the Log Export configuration for a cluster
[**DeleteSQLUser**](CockroachCloudApi.md#DeleteSQLUser) | **Delete** /api/v1/clusters/{cluster_id}/sql-users/{name} | Delete a SQL user
[**DeleteUser**](CockroachCloudApi.md#DeleteUser) | **Delete** /api/scim/v2/Users/{id} | Deletes a user based on id
[**EditDatabase**](CockroachCloudApi.md#EditDatabase) | **Patch** /api/v1/clusters/{cluster_id}/databases/{name} | Update a database
[**EditDatabase2**](CockroachCloudApi.md#EditDatabase2) | **Patch** /api/v1/clusters/{cluster_id}/databases | Update a database
[**EditEgressRule**](CockroachCloudApi.md#EditEgressRule) | **Patch** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Edit an existing egress rule
[**EnableCMEKSpec**](CockroachCloudApi.md#EnableCMEKSpec) | **Post** /api/v1/clusters/{cluster_id}/cmek | Enable CMEK for a cluster
[**EnableCloudWatchMetricExport**](CockroachCloudApi.md#EnableCloudWatchMetricExport) | **Post** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Create or update the CloudWatch Metric Export configuration for a cluster
[**EnableDatadogMetricExport**](CockroachCloudApi.md#EnableDatadogMetricExport) | **Post** /api/v1/clusters/{cluster_id}/metricexport/datadog | Create or update the Datadog Metric Export configuration for a cluster
[**EnableLogExport**](CockroachCloudApi.md#EnableLogExport) | **Post** /api/v1/clusters/{cluster_id}/logexport | Create or update the Log Export configuration for a cluster
[**GetAllRolesForUser**](CockroachCloudApi.md#GetAllRolesForUser) | **Get** /api/v1/roles/{user_id} | Gets All Role Grants for the specified user
[**GetCMEKClusterInfo**](CockroachCloudApi.md#GetCMEKClusterInfo) | **Get** /api/v1/clusters/{cluster_id}/cmek | Get CMEK-related information for a cluster
[**GetClientCACert**](CockroachCloudApi.md#GetClientCACert) | **Get** /api/v1/clusters/{cluster_id}/client-ca-cert | Get Client CA Cert information for a cluster
[**GetCloudWatchMetricExportInfo**](CockroachCloudApi.md#GetCloudWatchMetricExportInfo) | **Get** /api/v1/clusters/{cluster_id}/metricexport/cloudwatch | Get the CloudWatch Metric Export configuration for a cluster
[**GetCluster**](CockroachCloudApi.md#GetCluster) | **Get** /api/v1/clusters/{cluster_id} | Get extended information about a cluster
[**GetConnectionString**](CockroachCloudApi.md#GetConnectionString) | **Get** /api/v1/clusters/{cluster_id}/connection-string | Get a formatted generic connection string for a cluster
[**GetDatadogMetricExportInfo**](CockroachCloudApi.md#GetDatadogMetricExportInfo) | **Get** /api/v1/clusters/{cluster_id}/metricexport/datadog | Get the Datadog Metric Export configuration for a cluster
[**GetEgressRule**](CockroachCloudApi.md#GetEgressRule) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules/{rule_id} | Get an existing egress rule
[**GetGroup**](CockroachCloudApi.md#GetGroup) | **Get** /api/scim/v2/Groups/{id} | Gets a group based on id
[**GetGroup2**](CockroachCloudApi.md#GetGroup2) | **Put** /api/scim/v2/Groups/{id}/.search | Gets a group based on id
[**GetGroups**](CockroachCloudApi.md#GetGroups) | **Get** /api/scim/v2/Groups | Gets groups based on query parameters
[**GetGroups2**](CockroachCloudApi.md#GetGroups2) | **Put** /api/scim/v2/Groups/.search | Gets groups based on query parameters
[**GetInvoice**](CockroachCloudApi.md#GetInvoice) | **Get** /api/v1/invoices/{invoice_id} | Gets a specific invoice for an organization
[**GetLogExportInfo**](CockroachCloudApi.md#GetLogExportInfo) | **Get** /api/v1/clusters/{cluster_id}/logexport | Get the Log Export configuration for a cluster
[**GetOrganizationInfo**](CockroachCloudApi.md#GetOrganizationInfo) | **Get** /api/v1/organization | Get information about the caller&#39;s organization
[**GetPersonUsersByEmail**](CockroachCloudApi.md#GetPersonUsersByEmail) | **Get** /api/v1/users/persons-by-email | Search person users by email address
[**GetResourceType**](CockroachCloudApi.md#GetResourceType) | **Get** /api/scim/v2/ResourceTypes/{resourceId} | 
[**GetResourceTypes**](CockroachCloudApi.md#GetResourceTypes) | **Get** /api/scim/v2/ResourceTypes | 
[**GetSchema**](CockroachCloudApi.md#GetSchema) | **Get** /api/scim/v2/Schemas/{schemaId} | 
[**GetSchemas**](CockroachCloudApi.md#GetSchemas) | **Get** /api/scim/v2/Schemas | 
[**GetServiceProviderConfig**](CockroachCloudApi.md#GetServiceProviderConfig) | **Get** /api/scim/v2/ServiceProviderConfig | Returns our SCIM configuration
[**GetUser**](CockroachCloudApi.md#GetUser) | **Get** /api/scim/v2/Users/{id} | Gets a user based on id
[**GetUser2**](CockroachCloudApi.md#GetUser2) | **Put** /api/scim/v2/Users/{id}/.search | Gets a user based on id
[**GetUsers**](CockroachCloudApi.md#GetUsers) | **Get** /api/scim/v2/Users | Gets Users based on query parameters
[**GetUsers2**](CockroachCloudApi.md#GetUsers2) | **Put** /api/scim/v2/Users/.search | Gets Users based on query parameters
[**ListAllowlistEntries**](CockroachCloudApi.md#ListAllowlistEntries) | **Get** /api/v1/clusters/{cluster_id}/networking/allowlist | Get the IP allowlist and propagation status for a cluster
[**ListAuditLogs**](CockroachCloudApi.md#ListAuditLogs) | **Get** /api/v1/auditlogevents | Limited Access: List audit logs.
[**ListAvailableRegions**](CockroachCloudApi.md#ListAvailableRegions) | **Get** /api/v1/clusters/available-regions | List the regions available for new clusters and nodes
[**ListAwsEndpointConnections**](CockroachCloudApi.md#ListAwsEndpointConnections) | **Get** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections | Lists all AwsEndpointConnections for a given cluster
[**ListClusterNodes**](CockroachCloudApi.md#ListClusterNodes) | **Get** /api/v1/clusters/{cluster_id}/nodes | List nodes for a cluster
[**ListClusters**](CockroachCloudApi.md#ListClusters) | **Get** /api/v1/clusters | List clusters owned by an organization
[**ListDatabases**](CockroachCloudApi.md#ListDatabases) | **Get** /api/v1/clusters/{cluster_id}/databases | List databases for a cluster
[**ListEgressRules**](CockroachCloudApi.md#ListEgressRules) | **Get** /api/v1/clusters/{cluster_id}/networking/egress-rules | List all egress rules associates with a cluster
[**ListInvoices**](CockroachCloudApi.md#ListInvoices) | **Get** /api/v1/invoices | List invoices for a given organization
[**ListMajorClusterVersions**](CockroachCloudApi.md#ListMajorClusterVersions) | **Get** /api/v1/cluster-versions | List available major cluster versions
[**ListPrivateEndpointServices**](CockroachCloudApi.md#ListPrivateEndpointServices) | **Get** /api/v1/clusters/{cluster_id}/networking/private-endpoint-services | Lists all PrivateEndpointServices for a given cluster
[**ListRoleGrants**](CockroachCloudApi.md#ListRoleGrants) | **Get** /api/v1/roles | Lists all RoleGrants
[**ListSQLUsers**](CockroachCloudApi.md#ListSQLUsers) | **Get** /api/v1/clusters/{cluster_id}/sql-users | List SQL users for a cluster
[**PatchGroup**](CockroachCloudApi.md#PatchGroup) | **Patch** /api/scim/v2/Groups/{id} | Updates a group by specifying individual values to update
[**PatchUser**](CockroachCloudApi.md#PatchUser) | **Patch** /api/scim/v2/Users/{id} | Updates a user by specifying individual values to update
[**RemoveUserFromRole**](CockroachCloudApi.md#RemoveUserFromRole) | **Delete** /api/v1/roles/{user_id}/{resource_type}/{resource_id}/{role_name} | Removes the user from the given role
[**SetAwsEndpointConnectionState**](CockroachCloudApi.md#SetAwsEndpointConnectionState) | **Patch** /api/v1/clusters/{cluster_id}/networking/aws-endpoint-connections/{endpoint_id} | Sets the AWS Endpoint Connection state based on what is passed in the body
[**SetClientCACert**](CockroachCloudApi.md#SetClientCACert) | **Post** /api/v1/clusters/{cluster_id}/client-ca-cert | Set Client CA Cert for a cluster
[**SetEgressTrafficPolicy**](CockroachCloudApi.md#SetEgressTrafficPolicy) | **Post** /api/v1/clusters/{cluster_id}/networking/egress-rules/egress-traffic-policy | Outbound traffic management
[**SetRolesForUser**](CockroachCloudApi.md#SetRolesForUser) | **Put** /api/v1/roles/{user_id} | Makes the users roles exactly those provided
[**UpdateAllowlistEntry**](CockroachCloudApi.md#UpdateAllowlistEntry) | **Patch** /api/v1/clusters/{cluster_id}/networking/allowlist/{entry.cidr_ip}/{entry.cidr_mask} | Update an IP allowlist entry
[**UpdateCMEKSpec**](CockroachCloudApi.md#UpdateCMEKSpec) | **Put** /api/v1/clusters/{cluster_id}/cmek | Enable or update the CMEK spec for a cluster
[**UpdateCMEKStatus**](CockroachCloudApi.md#UpdateCMEKStatus) | **Patch** /api/v1/clusters/{cluster_id}/cmek | Update the CMEK-related status for a cluster
[**UpdateClientCACert**](CockroachCloudApi.md#UpdateClientCACert) | **Patch** /api/v1/clusters/{cluster_id}/client-ca-cert | Update Client CA Cert for a cluster
[**UpdateCluster**](CockroachCloudApi.md#UpdateCluster) | **Patch** /api/v1/clusters/{cluster_id} | Scale or edit a cluster
[**UpdateGroup**](CockroachCloudApi.md#UpdateGroup) | **Put** /api/scim/v2/Groups/{id} | Updates a group by supplying all values of the user object
[**UpdateSQLUserPassword**](CockroachCloudApi.md#UpdateSQLUserPassword) | **Put** /api/v1/clusters/{cluster_id}/sql-users/{name}/password | Update a SQL user&#39;s password
[**UpdateUser**](CockroachCloudApi.md#UpdateUser) | **Put** /api/scim/v2/Users/{id} | Updates a user by supplying all values of the user object



## AddAllowlistEntry

> AllowlistEntry AddAllowlistEntry(ctx, clusterId).AllowlistEntry(allowlistEntry).Execute()

Add a new CIDR address to the IP allowlist

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
    allowlistEntry := *openapiclient.NewAllowlistEntry("CidrIp_example", int32(123), false, false) // AllowlistEntry | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddAllowlistEntry(context.Background(), clusterId).AllowlistEntry(allowlistEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowlistEntry** | [**AllowlistEntry**](AllowlistEntry.md) |  | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddAllowlistEntry2

> AllowlistEntry AddAllowlistEntry2(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()

Add a new CIDR address to the IP allowlist

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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddAllowlistEntry2(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddAllowlistEntry2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistEntry2`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddAllowlistEntry2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAllowlistEntry2 struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry1** | [**AllowlistEntry1**](AllowlistEntry1.md) | AllowlistEntry | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddEgressRule

> AddEgressRuleResponse AddEgressRule(ctx, clusterId).AddEgressRuleRequest(addEgressRuleRequest).Execute()

Add an egress rule

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
    resp, r, err := api_client.CockroachCloudApi.AddEgressRule(context.Background(), clusterId).AddEgressRuleRequest(addEgressRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddEgressRule`: AddEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddEgressRule`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddUserToRole

> GetAllRolesForUserResponse AddUserToRole(ctx, userId, resourceType, resourceId, roleName).Execute()

Adds the user to the given role

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
    userId := "userId_example" // string | 
    resourceType := "resourceType_example" // string | 
    resourceId := "resourceId_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.AddUserToRole(context.Background(), userId, resourceType, resourceId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.AddUserToRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToRole`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.AddUserToRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToRole struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCluster

> Cluster CreateCluster(ctx).CreateClusterRequest(createClusterRequest).Execute()

Create and initialize a new cluster

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
    createClusterRequest := *openapiclient.NewCreateClusterRequest("Name_example", openapiclient.CloudProvider.Type("GCP"), *openapiclient.NewCreateClusterSpecification()) // CreateClusterRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateCluster(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateCluster`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCluster struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createClusterRequest** | [**CreateClusterRequest**](CreateClusterRequest.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDatabase

> ApiDatabase CreateDatabase(ctx, clusterId).CreateDatabaseRequest(createDatabaseRequest).Execute()

Create a new database

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
    createDatabaseRequest := *openapiclient.NewCreateDatabaseRequest("Name_example") // CreateDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateDatabase(context.Background(), clusterId).CreateDatabaseRequest(createDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDatabaseRequest** | [**CreateDatabaseRequest**](CreateDatabaseRequest.md) |  | 

### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> ScimGroup CreateGroup(ctx).CreateGroupRequest(createGroupRequest).Execute()

Creates a group

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
    resp, r, err := api_client.CockroachCloudApi.CreateGroup(context.Background()).CreateGroupRequest(createGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroup struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePrivateEndpointServices

> PrivateEndpointServices CreatePrivateEndpointServices(ctx, clusterId).Execute()

Creates all PrivateEndpointServices for a given cluster

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
    resp, r, err := api_client.CockroachCloudApi.CreatePrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreatePrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreatePrivateEndpointServices`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSQLUser

> SQLUser CreateSQLUser(ctx, clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()

Create a new SQL user

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
    createSQLUserRequest := *openapiclient.NewCreateSQLUserRequest("Name_example", "Password_example") // CreateSQLUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.CreateSQLUser(context.Background(), clusterId).CreateSQLUserRequest(createSQLUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSQLUserRequest** | [**CreateSQLUserRequest**](CreateSQLUserRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUser

> ScimUser CreateUser(ctx).CreateUserRequest(createUserRequest).Execute()

Creates a user

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
    resp, r, err := api_client.CockroachCloudApi.CreateUser(context.Background()).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.CreateUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUser struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistEntry

> AllowlistEntry DeleteAllowlistEntry(ctx, clusterId, cidrIp, cidrMask).Execute()

Delete an IP allowlist entry

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
    cidrIp := "cidrIp_example" // string | 
    cidrMask := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteAllowlistEntry(context.Background(), clusterId, cidrIp, cidrMask).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**cidrIp** | **string** |  | 
**cidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClientCACert

> ClientCACertInfo DeleteClientCACert(ctx, clusterId).Execute()

Delete Client CA Cert for a cluster

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
    resp, r, err := api_client.CockroachCloudApi.DeleteClientCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    resp, r, err := api_client.CockroachCloudApi.DeleteCloudWatchMetricExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteCloudWatchMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCloudWatchMetricExport`: DeleteMetricExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteCloudWatchMetricExport`: %v\n", resp)
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


## DeleteCluster

> Cluster DeleteCluster(ctx, clusterId).Execute()

Delete a cluster and all of its data

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
    resp, r, err := api_client.CockroachCloudApi.DeleteCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCluster struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDatabase

> ApiDatabase DeleteDatabase(ctx, clusterId, name).Execute()

Delete a database

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteDatabase(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiDatabase**](ApiDatabase.md)

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
    resp, r, err := api_client.CockroachCloudApi.DeleteDatadogMetricExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteDatadogMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDatadogMetricExport`: DeleteMetricExportResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteDatadogMetricExport`: %v\n", resp)
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


## DeleteEgressRule

> DeleteEgressRuleResponse DeleteEgressRule(ctx, clusterId, ruleId).IdempotencyKey(idempotencyKey).Execute()

Delete an existing egress rule

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
    resp, r, err := api_client.CockroachCloudApi.DeleteEgressRule(context.Background(), clusterId, ruleId).IdempotencyKey(idempotencyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEgressRule`: DeleteEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteEgressRule`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroup

> map[string]interface{} DeleteGroup(ctx, id).Execute()

Deletes a group based on id

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
    resp, r, err := api_client.CockroachCloudApi.DeleteGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteGroup`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroup struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLogExport

> LogExportClusterInfo DeleteLogExport(ctx, clusterId).Execute()

Delete the Log Export configuration for a cluster

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
    resp, r, err := api_client.CockroachCloudApi.DeleteLogExport(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSQLUser

> SQLUser DeleteSQLUser(ctx, clusterId, name).Execute()

Delete a SQL user

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.DeleteSQLUser(context.Background(), clusterId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteSQLUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSQLUser`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteSQLUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSQLUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> map[string]interface{} DeleteUser(ctx, id).Execute()

Deletes a user based on id

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
    resp, r, err := api_client.CockroachCloudApi.DeleteUser(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUser`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUser struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDatabase

> ApiDatabase EditDatabase(ctx, clusterId, name).UpdateDatabaseRequest1(updateDatabaseRequest1).Execute()

Update a database

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
    name := "name_example" // string | 
    updateDatabaseRequest1 := *openapiclient.NewUpdateDatabaseRequest1("NewName_example") // UpdateDatabaseRequest1 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EditDatabase(context.Background(), clusterId, name).UpdateDatabaseRequest1(updateDatabaseRequest1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EditDatabase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EditDatabase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDatabase struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDatabaseRequest1** | [**UpdateDatabaseRequest1**](UpdateDatabaseRequest1.md) |  | 

### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditDatabase2

> ApiDatabase EditDatabase2(ctx, clusterId).UpdateDatabaseRequest(updateDatabaseRequest).Execute()

Update a database

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
    updateDatabaseRequest := *openapiclient.NewUpdateDatabaseRequest("Name_example", "NewName_example") // UpdateDatabaseRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EditDatabase2(context.Background(), clusterId).UpdateDatabaseRequest(updateDatabaseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EditDatabase2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditDatabase2`: ApiDatabase
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EditDatabase2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditDatabase2 struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDatabaseRequest** | [**UpdateDatabaseRequest**](UpdateDatabaseRequest.md) |  | 

### Return type

[**ApiDatabase**](ApiDatabase.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditEgressRule

> EditEgressRuleResponse EditEgressRule(ctx, clusterId, ruleId).EditEgressRuleRequest(editEgressRuleRequest).Execute()

Edit an existing egress rule

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
    resp, r, err := api_client.CockroachCloudApi.EditEgressRule(context.Background(), clusterId, ruleId).EditEgressRuleRequest(editEgressRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EditEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditEgressRule`: EditEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EditEgressRule`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableCMEKSpec

> CMEKClusterInfo EnableCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable CMEK for a cluster

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
    cMEKClusterSpecification := *openapiclient.NewCMEKClusterSpecification([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // CMEKClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EnableCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableCMEKSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableCMEKSpec struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cMEKClusterSpecification** | [**CMEKClusterSpecification**](CMEKClusterSpecification.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
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
    resp, r, err := api_client.CockroachCloudApi.EnableCloudWatchMetricExport(context.Background(), clusterId).EnableCloudWatchMetricExportRequest(enableCloudWatchMetricExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableCloudWatchMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableCloudWatchMetricExport`: CloudWatchMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableCloudWatchMetricExport`: %v\n", resp)
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
    resp, r, err := api_client.CockroachCloudApi.EnableDatadogMetricExport(context.Background(), clusterId).EnableDatadogMetricExportRequest(enableDatadogMetricExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableDatadogMetricExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableDatadogMetricExport`: DatadogMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableDatadogMetricExport`: %v\n", resp)
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


## EnableLogExport

> LogExportClusterInfo EnableLogExport(ctx, clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()

Create or update the Log Export configuration for a cluster

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
    enableLogExportRequest := *openapiclient.NewEnableLogExportRequest("AuthPrincipal_example", "LogName_example", openapiclient.LogExportType("AWS_CLOUDWATCH")) // EnableLogExportRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.EnableLogExport(context.Background(), clusterId).EnableLogExportRequest(enableLogExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.EnableLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableLogExport`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.EnableLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableLogExport struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enableLogExportRequest** | [**EnableLogExportRequest**](EnableLogExportRequest.md) |  | 

### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllRolesForUser

> GetAllRolesForUserResponse GetAllRolesForUser(ctx, userId).Execute()

Gets All Role Grants for the specified user

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetAllRolesForUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetAllRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRolesForUser`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetAllRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRolesForUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCMEKClusterInfo

> CMEKClusterInfo GetCMEKClusterInfo(ctx, clusterId).Execute()

Get CMEK-related information for a cluster

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
    resp, r, err := api_client.CockroachCloudApi.GetCMEKClusterInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetCMEKClusterInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCMEKClusterInfo`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetCMEKClusterInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCMEKClusterInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClientCACert

> ClientCACertInfo GetClientCACert(ctx, clusterId).Execute()

Get Client CA Cert information for a cluster

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
    resp, r, err := api_client.CockroachCloudApi.GetClientCACert(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
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
    resp, r, err := api_client.CockroachCloudApi.GetCloudWatchMetricExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetCloudWatchMetricExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudWatchMetricExportInfo`: CloudWatchMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetCloudWatchMetricExportInfo`: %v\n", resp)
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


## GetCluster

> Cluster GetCluster(ctx, clusterId).Execute()

Get extended information about a cluster

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
    resp, r, err := api_client.CockroachCloudApi.GetCluster(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCluster struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectionString

> GetConnectionStringResponse GetConnectionString(ctx, clusterId).Database(database).SqlUser(sqlUser).Os(os).Execute()

Get a formatted generic connection string for a cluster

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
    database := "database_example" // string |  (optional) (default to "defaultdb")
    sqlUser := "sqlUser_example" // string |  (optional)
    os := "os_example" // string | os indicates the target operating system, used with formatting the default SSL certificate path. Required only for dedicated clusters. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetConnectionString(context.Background(), clusterId).Database(database).SqlUser(sqlUser).Os(os).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetConnectionString``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionString`: GetConnectionStringResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetConnectionString`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionString struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **database** | **string** |  | [default to &quot;defaultdb&quot;]
 **sqlUser** | **string** |  | 
 **os** | **string** | os indicates the target operating system, used with formatting the default SSL certificate path. Required only for dedicated clusters. | 

### Return type

[**GetConnectionStringResponse**](GetConnectionStringResponse.md)

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
    resp, r, err := api_client.CockroachCloudApi.GetDatadogMetricExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetDatadogMetricExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatadogMetricExportInfo`: DatadogMetricExportInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetDatadogMetricExportInfo`: %v\n", resp)
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


## GetEgressRule

> GetEgressRuleResponse GetEgressRule(ctx, clusterId, ruleId).Execute()

Get an existing egress rule

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
    resp, r, err := api_client.CockroachCloudApi.GetEgressRule(context.Background(), clusterId, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetEgressRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEgressRule`: GetEgressRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetEgressRule`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> ScimGroup GetGroup(ctx, id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Gets a group based on id

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
    resp, r, err := api_client.CockroachCloudApi.GetGroup(context.Background(), id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroup struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup2

> ScimGroup GetGroup2(ctx, id).GetGroupRequest(getGroupRequest).Execute()

Gets a group based on id

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
    resp, r, err := api_client.CockroachCloudApi.GetGroup2(context.Background(), id).GetGroupRequest(getGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetGroup2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroup2`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetGroup2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroup2 struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups

> GetGroupsResponse GetGroups(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Gets groups based on query parameters

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
    resp, r, err := api_client.CockroachCloudApi.GetGroups(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups`: GetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroups struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributes** | **string** |  | 
 **excludedAttributes** | **string** |  | 

### Return type

[**GetGroupsResponse**](GetGroupsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroups2

> GetGroupsResponse GetGroups2(ctx).GetGroupsRequest(getGroupsRequest).Execute()

Gets groups based on query parameters

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
    resp, r, err := api_client.CockroachCloudApi.GetGroups2(context.Background()).GetGroupsRequest(getGroupsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetGroups2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroups2`: GetGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetGroups2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroups2 struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> Invoice GetInvoice(ctx, invoiceId).Execute()

Gets a specific invoice for an organization

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
    resp, r, err := api_client.CockroachCloudApi.GetInvoice(context.Background(), invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInvoice`: Invoice
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetInvoice`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogExportInfo

> LogExportClusterInfo GetLogExportInfo(ctx, clusterId).Execute()

Get the Log Export configuration for a cluster

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
    resp, r, err := api_client.CockroachCloudApi.GetLogExportInfo(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetLogExportInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogExportInfo`: LogExportClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetLogExportInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogExportInfo struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LogExportClusterInfo**](LogExportClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizationInfo

> Organization GetOrganizationInfo(ctx).Execute()

Get information about the caller's organization

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
    resp, r, err := api_client.CockroachCloudApi.GetOrganizationInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetOrganizationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizationInfo`: Organization
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetOrganizationInfo`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonUsersByEmail

> GetPersonUsersByEmailResponse GetPersonUsersByEmail(ctx).Email(email).Execute()

Search person users by email address

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
    email := "email_example" // string | an email address is required.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.GetPersonUsersByEmail(context.Background()).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetPersonUsersByEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonUsersByEmail`: GetPersonUsersByEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetPersonUsersByEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonUsersByEmail struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | an email address is required. | 

### Return type

[**GetPersonUsersByEmailResponse**](GetPersonUsersByEmailResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceType

> ScimResourceType GetResourceType(ctx, resourceId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.GetResourceType(context.Background(), resourceId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetResourceType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceType`: ScimResourceType
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetResourceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceType struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceTypes

> GetResourceTypesResponse GetResourceTypes(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.GetResourceTypes(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetResourceTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceTypes`: GetResourceTypesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetResourceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceTypes struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchema

> ScimSchema GetSchema(ctx, schemaId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.GetSchema(context.Background(), schemaId).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchema`: ScimSchema
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchema struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemas

> GetSchemasResponse GetSchemas(ctx).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()



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
    resp, r, err := api_client.CockroachCloudApi.GetSchemas(context.Background()).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemas`: GetSchemasResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetSchemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemas struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceProviderConfig

> GetServiceProviderConfigResponse GetServiceProviderConfig(ctx).Execute()

Returns our SCIM configuration

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
    resp, r, err := api_client.CockroachCloudApi.GetServiceProviderConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetServiceProviderConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceProviderConfig`: GetServiceProviderConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetServiceProviderConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceProviderConfig struct via the builder pattern


### Return type

[**GetServiceProviderConfigResponse**](GetServiceProviderConfigResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> ScimUser GetUser(ctx, id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Gets a user based on id

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
    resp, r, err := api_client.CockroachCloudApi.GetUser(context.Background(), id).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUser struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser2

> ScimUser GetUser2(ctx, id).GetUserRequest(getUserRequest).Execute()

Gets a user based on id

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
    resp, r, err := api_client.CockroachCloudApi.GetUser2(context.Background(), id).GetUserRequest(getUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetUser2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser2`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetUser2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUser2 struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> GetUsersResponse GetUsers(ctx).Filter(filter).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()

Gets Users based on query parameters

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
    resp, r, err := api_client.CockroachCloudApi.GetUsers(context.Background()).Filter(filter).Attributes(attributes).ExcludedAttributes(excludedAttributes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: GetUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsers struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers2

> GetUsersResponse GetUsers2(ctx).GetUsersRequest(getUsersRequest).Execute()

Gets Users based on query parameters

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
    resp, r, err := api_client.CockroachCloudApi.GetUsers2(context.Background()).GetUsersRequest(getUsersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.GetUsers2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers2`: GetUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.GetUsers2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsers2 struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllowlistEntries

> ListAllowlistEntriesResponse ListAllowlistEntries(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

Get the IP allowlist and propagation status for a cluster



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
    clusterId := "clusterId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAllowlistEntries(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAllowlistEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowlistEntries`: ListAllowlistEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAllowlistEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowlistEntries struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListAllowlistEntriesResponse**](ListAllowlistEntriesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAuditLogs

> ListAuditLogsResponse ListAuditLogs(ctx).StartingFrom(startingFrom).SortOrder(sortOrder).Limit(limit).Execute()

Limited Access: List audit logs.

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
    resp, r, err := api_client.CockroachCloudApi.ListAuditLogs(context.Background()).StartingFrom(startingFrom).SortOrder(sortOrder).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuditLogs`: ListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAuditLogs`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableRegions

> ListAvailableRegionsResponse ListAvailableRegions(ctx).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List the regions available for new clusters and nodes



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
    provider := "provider_example" // string | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider.  - AZURE: Limited Access: The Azure cloud provider. (optional)
    serverless := true // bool | Optional filter to only show regions available for serverless clusters. (optional) (default to false)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListAvailableRegions(context.Background()).Provider(provider).Serverless(serverless).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAvailableRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableRegions`: ListAvailableRegionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAvailableRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableRegions struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Optional CloudProvider for filtering.   - GCP: The Google Cloud Platform cloud provider.  - AWS: The Amazon Web Services cloud provider.  - AZURE: Limited Access: The Azure cloud provider. | 
 **serverless** | **bool** | Optional filter to only show regions available for serverless clusters. | [default to false]
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListAvailableRegionsResponse**](ListAvailableRegionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAwsEndpointConnections

> AwsEndpointConnections ListAwsEndpointConnections(ctx, clusterId).Execute()

Lists all AwsEndpointConnections for a given cluster

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
    resp, r, err := api_client.CockroachCloudApi.ListAwsEndpointConnections(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListAwsEndpointConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsEndpointConnections`: AwsEndpointConnections
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListAwsEndpointConnections`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusterNodes

> ListClusterNodesResponse ListClusterNodes(ctx, clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List nodes for a cluster



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
    clusterId := "clusterId_example" // string | 
    regionName := "regionName_example" // string | Optional filter to limit response to a single region. (optional)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusterNodes(context.Background(), clusterId).RegionName(regionName).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListClusterNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusterNodes`: ListClusterNodesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListClusterNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClusterNodes struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **regionName** | **string** | Optional filter to limit response to a single region. | 
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListClusterNodesResponse**](ListClusterNodesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> ListClustersResponse ListClusters(ctx).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List clusters owned by an organization



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
    showInactive := true // bool | If `true`, show clusters that have been deleted or failed to initialize. (optional) (default to false)
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListClusters(context.Background()).ShowInactive(showInactive).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListClusters`: ListClustersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListClusters struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showInactive** | **bool** | If &#x60;true&#x60;, show clusters that have been deleted or failed to initialize. | [default to false]
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListClustersResponse**](ListClustersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDatabases

> ApiListDatabasesResponse ListDatabases(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List databases for a cluster



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
    clusterId := "clusterId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListDatabases(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListDatabases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDatabases`: ApiListDatabasesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListDatabases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDatabases struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ApiListDatabasesResponse**](ApiListDatabasesResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEgressRules

> ListEgressRulesResponse ListEgressRules(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List all egress rules associates with a cluster



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
    resp, r, err := api_client.CockroachCloudApi.ListEgressRules(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListEgressRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEgressRules`: ListEgressRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListEgressRules`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> ListInvoicesResponse ListInvoices(ctx).Execute()

List invoices for a given organization



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
    resp, r, err := api_client.CockroachCloudApi.ListInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvoices`: ListInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListInvoices`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMajorClusterVersions

> ListMajorClusterVersionsResponse ListMajorClusterVersions(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List available major cluster versions



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
    resp, r, err := api_client.CockroachCloudApi.ListMajorClusterVersions(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListMajorClusterVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMajorClusterVersions`: ListMajorClusterVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListMajorClusterVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMajorClusterVersions struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListMajorClusterVersionsResponse**](ListMajorClusterVersionsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPrivateEndpointServices

> PrivateEndpointServices ListPrivateEndpointServices(ctx, clusterId).Execute()

Lists all PrivateEndpointServices for a given cluster



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
    resp, r, err := api_client.CockroachCloudApi.ListPrivateEndpointServices(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListPrivateEndpointServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrivateEndpointServices`: PrivateEndpointServices
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListPrivateEndpointServices`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRoleGrants

> ListRoleGrantsResponse ListRoleGrants(ctx).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

Lists all RoleGrants

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
    resp, r, err := api_client.CockroachCloudApi.ListRoleGrants(context.Background()).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListRoleGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRoleGrants`: ListRoleGrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListRoleGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRoleGrants struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListRoleGrantsResponse**](ListRoleGrantsResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSQLUsers

> ListSQLUsersResponse ListSQLUsers(ctx, clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()

List SQL users for a cluster



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
    clusterId := "clusterId_example" // string | 
    paginationPage := "paginationPage_example" // string |  (optional)
    paginationLimit := int32(56) // int32 |  (optional)
    paginationAsOfTime := time.Now() // time.Time |  (optional)
    paginationSortOrder := "paginationSortOrder_example" // string |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.ListSQLUsers(context.Background(), clusterId).PaginationPage(paginationPage).PaginationLimit(paginationLimit).PaginationAsOfTime(paginationAsOfTime).PaginationSortOrder(paginationSortOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.ListSQLUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSQLUsers`: ListSQLUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.ListSQLUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSQLUsers struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationPage** | **string** |  | 
 **paginationLimit** | **int32** |  | 
 **paginationAsOfTime** | **time.Time** |  | 
 **paginationSortOrder** | **string** |  - ASC: Sort in ascending order. This is the default unless otherwise specified.  - DESC: Sort in descending order. | 

### Return type

[**ListSQLUsersResponse**](ListSQLUsersResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchGroup

> ScimGroup PatchGroup(ctx, id).PatchGroupRequest(patchGroupRequest).Execute()

Updates a group by specifying individual values to update

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
    patchGroupRequest := *openapiclient.NewPatchGroupRequest() // PatchGroupRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.PatchGroup(context.Background(), id).PatchGroupRequest(patchGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.PatchGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.PatchGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchGroup struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchGroupRequest** | [**PatchGroupRequest**](PatchGroupRequest.md) |  | 

### Return type

[**ScimGroup**](ScimGroup.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchUser

> ScimUser PatchUser(ctx, id).PatchUserRequest(patchUserRequest).Execute()

Updates a user by specifying individual values to update

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
    patchUserRequest := *openapiclient.NewPatchUserRequest() // PatchUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.PatchUser(context.Background(), id).PatchUserRequest(patchUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.PatchUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.PatchUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchUserRequest** | [**PatchUserRequest**](PatchUserRequest.md) |  | 

### Return type

[**ScimUser**](ScimUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromRole

> GetAllRolesForUserResponse RemoveUserFromRole(ctx, userId, resourceType, resourceId, roleName).Execute()

Removes the user from the given role

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
    userId := "userId_example" // string | 
    resourceType := "resourceType_example" // string | 
    resourceId := "resourceId_example" // string | 
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.RemoveUserFromRole(context.Background(), userId, resourceType, resourceId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.RemoveUserFromRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveUserFromRole`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.RemoveUserFromRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 
**resourceType** | **string** |  | 
**resourceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromRole struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAwsEndpointConnectionState

> AwsEndpointConnection SetAwsEndpointConnectionState(ctx, clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()

Sets the AWS Endpoint Connection state based on what is passed in the body



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
    resp, r, err := api_client.CockroachCloudApi.SetAwsEndpointConnectionState(context.Background(), clusterId, endpointId).SetAwsEndpointConnectionStateRequest(setAwsEndpointConnectionStateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.SetAwsEndpointConnectionState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAwsEndpointConnectionState`: AwsEndpointConnection
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.SetAwsEndpointConnectionState`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetClientCACert

> ClientCACertInfo SetClientCACert(ctx, clusterId).SetClientCACertRequest(setClientCACertRequest).Execute()

Set Client CA Cert for a cluster

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
    setClientCACertRequest := *openapiclient.NewSetClientCACertRequest("X509PemCert_example") // SetClientCACertRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.SetClientCACert(context.Background(), clusterId).SetClientCACertRequest(setClientCACertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.SetClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.SetClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setClientCACertRequest** | [**SetClientCACertRequest**](SetClientCACertRequest.md) |  | 

### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetEgressTrafficPolicy

> map[string]interface{} SetEgressTrafficPolicy(ctx, clusterId).SetEgressTrafficPolicyRequest(setEgressTrafficPolicyRequest).Execute()

Outbound traffic management

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
    resp, r, err := api_client.CockroachCloudApi.SetEgressTrafficPolicy(context.Background(), clusterId).SetEgressTrafficPolicyRequest(setEgressTrafficPolicyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.SetEgressTrafficPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetEgressTrafficPolicy`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.SetEgressTrafficPolicy`: %v\n", resp)
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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetRolesForUser

> GetAllRolesForUserResponse SetRolesForUser(ctx, userId).CockroachCloudSetRolesForUserRequest(cockroachCloudSetRolesForUserRequest).Execute()

Makes the users roles exactly those provided

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
    userId := "userId_example" // string | 
    cockroachCloudSetRolesForUserRequest := *openapiclient.NewCockroachCloudSetRolesForUserRequest([]openapiclient.BuiltInRole{*openapiclient.NewBuiltInRole(openapiclient.OrganizationUserRole.Type("DEVELOPER"), *openapiclient.NewResource(openapiclient.ResourceType.Type("ORGANIZATION")))}) // CockroachCloudSetRolesForUserRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.SetRolesForUser(context.Background(), userId).CockroachCloudSetRolesForUserRequest(cockroachCloudSetRolesForUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.SetRolesForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetRolesForUser`: GetAllRolesForUserResponse
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.SetRolesForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRolesForUser struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cockroachCloudSetRolesForUserRequest** | [**CockroachCloudSetRolesForUserRequest**](CockroachCloudSetRolesForUserRequest.md) |  | 

### Return type

[**GetAllRolesForUserResponse**](GetAllRolesForUserResponse.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAllowlistEntry

> AllowlistEntry UpdateAllowlistEntry(ctx, clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()

Update an IP allowlist entry

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
    entryCidrIp := "entryCidrIp_example" // string | 
    entryCidrMask := int32(56) // int32 | 
    allowlistEntry1 := *openapiclient.NewAllowlistEntry1(false, false) // AllowlistEntry1 | AllowlistEntry

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateAllowlistEntry(context.Background(), clusterId, entryCidrIp, entryCidrMask).AllowlistEntry1(allowlistEntry1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateAllowlistEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAllowlistEntry`: AllowlistEntry
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateAllowlistEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**entryCidrIp** | **string** |  | 
**entryCidrMask** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAllowlistEntry struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **allowlistEntry1** | [**AllowlistEntry1**](AllowlistEntry1.md) | AllowlistEntry | 

### Return type

[**AllowlistEntry**](AllowlistEntry.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCMEKSpec

> CMEKClusterInfo UpdateCMEKSpec(ctx, clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()

Enable or update the CMEK spec for a cluster

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
    cMEKClusterSpecification := *openapiclient.NewCMEKClusterSpecification([]openapiclient.CMEKRegionSpecification{*openapiclient.NewCMEKRegionSpecification()}) // CMEKClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateCMEKSpec(context.Background(), clusterId).CMEKClusterSpecification(cMEKClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCMEKSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKSpec`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCMEKSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCMEKSpec struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cMEKClusterSpecification** | [**CMEKClusterSpecification**](CMEKClusterSpecification.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCMEKStatus

> CMEKClusterInfo UpdateCMEKStatus(ctx, clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()

Update the CMEK-related status for a cluster

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
    updateCMEKStatusRequest := *openapiclient.NewUpdateCMEKStatusRequest(openapiclient.CMEKCustomerAction("REVOKE")) // UpdateCMEKStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateCMEKStatus(context.Background(), clusterId).UpdateCMEKStatusRequest(updateCMEKStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCMEKStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCMEKStatus`: CMEKClusterInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCMEKStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCMEKStatus struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCMEKStatusRequest** | [**UpdateCMEKStatusRequest**](UpdateCMEKStatusRequest.md) |  | 

### Return type

[**CMEKClusterInfo**](CMEKClusterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateClientCACert

> ClientCACertInfo UpdateClientCACert(ctx, clusterId).UpdateClientCACertRequest(updateClientCACertRequest).Execute()

Update Client CA Cert for a cluster

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
    updateClientCACertRequest := *openapiclient.NewUpdateClientCACertRequest() // UpdateClientCACertRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateClientCACert(context.Background(), clusterId).UpdateClientCACertRequest(updateClientCACertRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateClientCACert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateClientCACert`: ClientCACertInfo
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateClientCACert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateClientCACert struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClientCACertRequest** | [**UpdateClientCACertRequest**](UpdateClientCACertRequest.md) |  | 

### Return type

[**ClientCACertInfo**](ClientCACertInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCluster

> Cluster UpdateCluster(ctx, clusterId).UpdateClusterSpecification(updateClusterSpecification).Execute()

Scale or edit a cluster

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
    updateClusterSpecification := *openapiclient.NewUpdateClusterSpecification() // UpdateClusterSpecification | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateCluster(context.Background(), clusterId).UpdateClusterSpecification(updateClusterSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateCluster``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCluster`: Cluster
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCluster struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateClusterSpecification** | [**UpdateClusterSpecification**](UpdateClusterSpecification.md) |  | 

### Return type

[**Cluster**](Cluster.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroup

> ScimGroup UpdateGroup(ctx, id).UpdateGroupRequest(updateGroupRequest).Execute()

Updates a group by supplying all values of the user object

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
    resp, r, err := api_client.CockroachCloudApi.UpdateGroup(context.Background(), id).UpdateGroupRequest(updateGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroup`: ScimGroup
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroup struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSQLUserPassword

> SQLUser UpdateSQLUserPassword(ctx, clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()

Update a SQL user's password

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
    name := "name_example" // string | 
    updateSQLUserPasswordRequest := *openapiclient.NewUpdateSQLUserPasswordRequest("Password_example") // UpdateSQLUserPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewClient(configuration)
    resp, r, err := api_client.CockroachCloudApi.UpdateSQLUserPassword(context.Background(), clusterId, name).UpdateSQLUserPasswordRequest(updateSQLUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateSQLUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSQLUserPassword`: SQLUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateSQLUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSQLUserPassword struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateSQLUserPasswordRequest** | [**UpdateSQLUserPasswordRequest**](UpdateSQLUserPasswordRequest.md) |  | 

### Return type

[**SQLUser**](SQLUser.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> ScimUser UpdateUser(ctx, id).UpdateUserRequest(updateUserRequest).Execute()

Updates a user by supplying all values of the user object

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
    resp, r, err := api_client.CockroachCloudApi.UpdateUser(context.Background(), id).UpdateUserRequest(updateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CockroachCloudApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUser`: ScimUser
    fmt.Fprintf(os.Stdout, "Response from `CockroachCloudApi.UpdateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUser struct via the builder pattern


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
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

