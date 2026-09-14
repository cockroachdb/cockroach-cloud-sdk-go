# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedHardwareConfig**](DedicatedHardwareConfig.md) |  | [optional] 
**Host** | Pointer to [**HostClusterConfig**](HostClusterConfig.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterConfig**](ServerlessClusterConfig.md) |  | [optional] 
**Virtual** | Pointer to [**VirtualClusterConfig**](VirtualClusterConfig.md) |  | [optional] 

## Methods

### NewClusterConfig

`func NewClusterConfig() *ClusterConfig`

NewClusterConfig instantiates a new ClusterConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDedicated

`func (o *ClusterConfig) GetDedicated() DedicatedHardwareConfig`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### SetDedicated

`func (o *ClusterConfig) SetDedicated(v DedicatedHardwareConfig)`

SetDedicated sets Dedicated field to given value.

### GetHost

`func (o *ClusterConfig) GetHost() HostClusterConfig`

GetHost returns the Host field if non-nil, zero value otherwise.

### SetHost

`func (o *ClusterConfig) SetHost(v HostClusterConfig)`

SetHost sets Host field to given value.

### GetServerless

`func (o *ClusterConfig) GetServerless() ServerlessClusterConfig`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### SetServerless

`func (o *ClusterConfig) SetServerless(v ServerlessClusterConfig)`

SetServerless sets Serverless field to given value.

### GetVirtual

`func (o *ClusterConfig) GetVirtual() VirtualClusterConfig`

GetVirtual returns the Virtual field if non-nil, zero value otherwise.

### SetVirtual

`func (o *ClusterConfig) SetVirtual(v VirtualClusterConfig)`

SetVirtual sets Virtual field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


