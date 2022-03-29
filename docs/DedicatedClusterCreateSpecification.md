# DedicatedClusterCreateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionNodes** | **map[string]int32** | Region keys should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. Values represent the node count. | 
**Hardware** | [**DedicatedHardwareCreateSpecification**](DedicatedHardwareCreateSpecification.md) |  | 
**CockroachVersion** | Pointer to **string** | The CockroachDB version for the cluster. The current version is used if omitted. | [optional] 

## Methods

### NewDedicatedClusterCreateSpecification

`func NewDedicatedClusterCreateSpecification(regionNodes map[string]int32, hardware DedicatedHardwareCreateSpecification, ) *DedicatedClusterCreateSpecification`

NewDedicatedClusterCreateSpecification instantiates a new DedicatedClusterCreateSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedClusterCreateSpecificationWithDefaults

`func NewDedicatedClusterCreateSpecificationWithDefaults() *DedicatedClusterCreateSpecification`

NewDedicatedClusterCreateSpecificationWithDefaults instantiates a new DedicatedClusterCreateSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegionNodes

`func (o *DedicatedClusterCreateSpecification) GetRegionNodes() map[string]int32`

GetRegionNodes returns the RegionNodes field if non-nil, zero value otherwise.

### GetRegionNodesOk

`func (o *DedicatedClusterCreateSpecification) GetRegionNodesOk() (*map[string]int32, bool)`

GetRegionNodesOk returns a tuple with the RegionNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionNodes

`func (o *DedicatedClusterCreateSpecification) SetRegionNodes(v map[string]int32)`

SetRegionNodes sets RegionNodes field to given value.


### GetHardware

`func (o *DedicatedClusterCreateSpecification) GetHardware() DedicatedHardwareCreateSpecification`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *DedicatedClusterCreateSpecification) GetHardwareOk() (*DedicatedHardwareCreateSpecification, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *DedicatedClusterCreateSpecification) SetHardware(v DedicatedHardwareCreateSpecification)`

SetHardware sets Hardware field to given value.


### GetCockroachVersion

`func (o *DedicatedClusterCreateSpecification) GetCockroachVersion() string`

GetCockroachVersion returns the CockroachVersion field if non-nil, zero value otherwise.

### GetCockroachVersionOk

`func (o *DedicatedClusterCreateSpecification) GetCockroachVersionOk() (*string, bool)`

GetCockroachVersionOk returns a tuple with the CockroachVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCockroachVersion

`func (o *DedicatedClusterCreateSpecification) SetCockroachVersion(v string)`

SetCockroachVersion sets CockroachVersion field to given value.

### HasCockroachVersion

`func (o *DedicatedClusterCreateSpecification) HasCockroachVersion() bool`

HasCockroachVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


