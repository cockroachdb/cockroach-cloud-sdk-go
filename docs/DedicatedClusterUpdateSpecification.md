# DedicatedClusterUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionNodes** | Pointer to **map[string]int32** | Region keys should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. Values represent the node count. | [optional] 
**Hardware** | Pointer to [**DedicatedHardwareUpdateSpecification**](DedicatedHardwareUpdateSpecification.md) |  | [optional] 

## Methods

### NewDedicatedClusterUpdateSpecification

`func NewDedicatedClusterUpdateSpecification() *DedicatedClusterUpdateSpecification`

NewDedicatedClusterUpdateSpecification instantiates a new DedicatedClusterUpdateSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedClusterUpdateSpecificationWithDefaults

`func NewDedicatedClusterUpdateSpecificationWithDefaults() *DedicatedClusterUpdateSpecification`

NewDedicatedClusterUpdateSpecificationWithDefaults instantiates a new DedicatedClusterUpdateSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionNodes

`func (o *DedicatedClusterUpdateSpecification) GetRegionNodes() map[string]int32`

GetRegionNodes returns the RegionNodes field if non-nil, zero value otherwise.

### GetRegionNodesOk

`func (o *DedicatedClusterUpdateSpecification) GetRegionNodesOk() (*map[string]int32, bool)`

GetRegionNodesOk returns a tuple with the RegionNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNodes

`func (o *DedicatedClusterUpdateSpecification) SetRegionNodes(v map[string]int32)`

SetRegionNodes sets RegionNodes field to given value.

### HasRegionNodes

`func (o *DedicatedClusterUpdateSpecification) HasRegionNodes() bool`

HasRegionNodes returns a boolean if a field has been set.

### GetHardware

`func (o *DedicatedClusterUpdateSpecification) GetHardware() DedicatedHardwareUpdateSpecification`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DedicatedClusterUpdateSpecification) GetHardwareOk() (*DedicatedHardwareUpdateSpecification, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DedicatedClusterUpdateSpecification) SetHardware(v DedicatedHardwareUpdateSpecification)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *DedicatedClusterUpdateSpecification) HasHardware() bool`

HasHardware returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


