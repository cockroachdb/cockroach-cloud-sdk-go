# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedHardwareConfig**](DedicatedHardwareConfig.md) |  | [optional] 
**Shared** | Pointer to [**SharedClusterConfig**](SharedClusterConfig.md) |  | [optional] 

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

### GetShared

`func (o *ClusterConfig) GetShared() SharedClusterConfig`

GetShared returns the Shared field if non-nil, zero value otherwise.

### SetShared

`func (o *ClusterConfig) SetShared(v SharedClusterConfig)`

SetShared sets Shared field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


