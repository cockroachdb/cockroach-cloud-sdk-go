# CockroachCloudUpdateClusterDisruptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionalDisruptorSpecifications** | Pointer to [**[]RegionalDisruptorSpecification**](RegionalDisruptorSpecification.md) | regional_disruptor_specifications specify how regions are to be disrupted. Any Cluster region that is not specified here will not be disrupted. A cluster region that was previously disrupted but is not listed here will be removed from disruption. To stop all disruptions, set this to an empty list or omit it from the request. | [optional] 

## Methods

### NewCockroachCloudUpdateClusterDisruptionRequest

`func NewCockroachCloudUpdateClusterDisruptionRequest() *CockroachCloudUpdateClusterDisruptionRequest`

NewCockroachCloudUpdateClusterDisruptionRequest instantiates a new CockroachCloudUpdateClusterDisruptionRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetRegionalDisruptorSpecifications

`func (o *CockroachCloudUpdateClusterDisruptionRequest) GetRegionalDisruptorSpecifications() []RegionalDisruptorSpecification`

GetRegionalDisruptorSpecifications returns the RegionalDisruptorSpecifications field if non-nil, zero value otherwise.

### SetRegionalDisruptorSpecifications

`func (o *CockroachCloudUpdateClusterDisruptionRequest) SetRegionalDisruptorSpecifications(v []RegionalDisruptorSpecification)`

SetRegionalDisruptorSpecifications sets RegionalDisruptorSpecifications field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


