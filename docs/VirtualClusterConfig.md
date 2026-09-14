# VirtualClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingId** | **string** | routing_id is used to identify the cluster in a connection string. | 
**VirtualCpuBurstLimit** | **int32** | virtual_cpu_burst_limit is a positive integer describing the most vCPUs the cluster can use at once. | 
**WorkspaceId** | **string** | workspace_id is the workspace the cluster is currently a member of. | 

## Methods

### NewVirtualClusterConfig

`func NewVirtualClusterConfig(routingId string, virtualCpuBurstLimit int32, workspaceId string, ) *VirtualClusterConfig`

NewVirtualClusterConfig instantiates a new VirtualClusterConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewVirtualClusterConfigWithDefaults

`func NewVirtualClusterConfigWithDefaults() *VirtualClusterConfig`

NewVirtualClusterConfigWithDefaults instantiates a new VirtualClusterConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRoutingId

`func (o *VirtualClusterConfig) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### SetRoutingId

`func (o *VirtualClusterConfig) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### GetVirtualCpuBurstLimit

`func (o *VirtualClusterConfig) GetVirtualCpuBurstLimit() int32`

GetVirtualCpuBurstLimit returns the VirtualCpuBurstLimit field if non-nil, zero value otherwise.

### SetVirtualCpuBurstLimit

`func (o *VirtualClusterConfig) SetVirtualCpuBurstLimit(v int32)`

SetVirtualCpuBurstLimit sets VirtualCpuBurstLimit field to given value.

### GetWorkspaceId

`func (o *VirtualClusterConfig) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### SetWorkspaceId

`func (o *VirtualClusterConfig) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


