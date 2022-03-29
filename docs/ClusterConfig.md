# ClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedHardwareConfig**](DedicatedHardwareConfig.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterConfig**](ServerlessClusterConfig.md) |  | [optional] 

## Methods

### NewClusterConfig

`func NewClusterConfig() *ClusterConfig`

NewClusterConfig instantiates a new ClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterConfigWithDefaults

`func NewClusterConfigWithDefaults() *ClusterConfig`

NewClusterConfigWithDefaults instantiates a new ClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *ClusterConfig) GetDedicated() DedicatedHardwareConfig`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *ClusterConfig) GetDedicatedOk() (*DedicatedHardwareConfig, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *ClusterConfig) SetDedicated(v DedicatedHardwareConfig)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *ClusterConfig) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetServerless

`func (o *ClusterConfig) GetServerless() ServerlessClusterConfig`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *ClusterConfig) GetServerlessOk() (*ServerlessClusterConfig, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *ClusterConfig) SetServerless(v ServerlessClusterConfig)`

SetServerless sets Serverless field to given value.

### HasServerless

`func (o *ClusterConfig) HasServerless() bool`

HasServerless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


