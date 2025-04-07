# ClusterDisruptionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionalDisruptorSpecifications** | [**[]RegionalDisruptorSpecification**](RegionalDisruptorSpecification.md) | regional_disruptor_specifications represents a region that is disrupted. If there are no entries, it means the cluster is not disrupted. | 

## Methods

### NewClusterDisruptionInfo

`func NewClusterDisruptionInfo(regionalDisruptorSpecifications []RegionalDisruptorSpecification, ) *ClusterDisruptionInfo`

NewClusterDisruptionInfo instantiates a new ClusterDisruptionInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewClusterDisruptionInfoWithDefaults

`func NewClusterDisruptionInfoWithDefaults() *ClusterDisruptionInfo`

NewClusterDisruptionInfoWithDefaults instantiates a new ClusterDisruptionInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRegionalDisruptorSpecifications

`func (o *ClusterDisruptionInfo) GetRegionalDisruptorSpecifications() []RegionalDisruptorSpecification`

GetRegionalDisruptorSpecifications returns the RegionalDisruptorSpecifications field if non-nil, zero value otherwise.

### SetRegionalDisruptorSpecifications

`func (o *ClusterDisruptionInfo) SetRegionalDisruptorSpecifications(v []RegionalDisruptorSpecification)`

SetRegionalDisruptorSpecifications sets RegionalDisruptorSpecifications field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


