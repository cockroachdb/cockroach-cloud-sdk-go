# DedicatedClusterUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CmekRegionSpecs** | Pointer to [**[]CMEKRegionSpecification**](CMEKRegionSpecification.md) | This field should contain the CMEK specs for newly added regions. If a CMEK spec is provided for an existing region, the request is invalid and will fail. | [optional] 
**Hardware** | Pointer to [**DedicatedHardwareUpdateSpecification**](DedicatedHardwareUpdateSpecification.md) |  | [optional] 
**RegionNodes** | Pointer to **map[string]int32** | Region keys should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. Values represent the node count. | [optional] 

## Methods

### NewDedicatedClusterUpdateSpecification

`func NewDedicatedClusterUpdateSpecification() *DedicatedClusterUpdateSpecification`

NewDedicatedClusterUpdateSpecification instantiates a new DedicatedClusterUpdateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetCmekRegionSpecs

`func (o *DedicatedClusterUpdateSpecification) GetCmekRegionSpecs() []CMEKRegionSpecification`

GetCmekRegionSpecs returns the CmekRegionSpecs field if non-nil, zero value otherwise.

### SetCmekRegionSpecs

`func (o *DedicatedClusterUpdateSpecification) SetCmekRegionSpecs(v []CMEKRegionSpecification)`

SetCmekRegionSpecs sets CmekRegionSpecs field to given value.

### GetHardware

`func (o *DedicatedClusterUpdateSpecification) GetHardware() DedicatedHardwareUpdateSpecification`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### SetHardware

`func (o *DedicatedClusterUpdateSpecification) SetHardware(v DedicatedHardwareUpdateSpecification)`

SetHardware sets Hardware field to given value.

### GetRegionNodes

`func (o *DedicatedClusterUpdateSpecification) GetRegionNodes() map[string]int32`

GetRegionNodes returns the RegionNodes field if non-nil, zero value otherwise.

### SetRegionNodes

`func (o *DedicatedClusterUpdateSpecification) SetRegionNodes(v map[string]int32)`

SetRegionNodes sets RegionNodes field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


