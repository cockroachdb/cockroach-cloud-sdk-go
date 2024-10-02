# DedicatedClusterCreateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CockroachVersion** | Pointer to **string** | The CockroachDB major version for the cluster. i.e. v24.1  The latest version is used if omitted. The version passed must be one of the currently supported versions. | [optional] 
**Hardware** | [**DedicatedHardwareCreateSpecification**](DedicatedHardwareCreateSpecification.md) |  | 
**NetworkVisibility** | Pointer to [**NetworkVisibilityType**](NetworkVisibilityType.md) |  | [optional] 
**RegionNodes** | **map[string]int32** | Region keys should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. Values represent the node count. | 
**RestrictEgressTraffic** | Pointer to **bool** | Preview: RestrictEgressTraffic if set, results in an egress traffic policy of default-deny at creation time. | [optional] 

## Methods

### NewDedicatedClusterCreateSpecification

`func NewDedicatedClusterCreateSpecification(hardware DedicatedHardwareCreateSpecification, regionNodes map[string]int32, ) *DedicatedClusterCreateSpecification`

NewDedicatedClusterCreateSpecification instantiates a new DedicatedClusterCreateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDedicatedClusterCreateSpecificationWithDefaults

`func NewDedicatedClusterCreateSpecificationWithDefaults() *DedicatedClusterCreateSpecification`

NewDedicatedClusterCreateSpecificationWithDefaults instantiates a new DedicatedClusterCreateSpecification object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCockroachVersion

`func (o *DedicatedClusterCreateSpecification) GetCockroachVersion() string`

GetCockroachVersion returns the CockroachVersion field if non-nil, zero value otherwise.

### SetCockroachVersion

`func (o *DedicatedClusterCreateSpecification) SetCockroachVersion(v string)`

SetCockroachVersion sets CockroachVersion field to given value.

### GetHardware

`func (o *DedicatedClusterCreateSpecification) GetHardware() DedicatedHardwareCreateSpecification`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### SetHardware

`func (o *DedicatedClusterCreateSpecification) SetHardware(v DedicatedHardwareCreateSpecification)`

SetHardware sets Hardware field to given value.

### GetNetworkVisibility

`func (o *DedicatedClusterCreateSpecification) GetNetworkVisibility() NetworkVisibilityType`

GetNetworkVisibility returns the NetworkVisibility field if non-nil, zero value otherwise.

### SetNetworkVisibility

`func (o *DedicatedClusterCreateSpecification) SetNetworkVisibility(v NetworkVisibilityType)`

SetNetworkVisibility sets NetworkVisibility field to given value.

### GetRegionNodes

`func (o *DedicatedClusterCreateSpecification) GetRegionNodes() map[string]int32`

GetRegionNodes returns the RegionNodes field if non-nil, zero value otherwise.

### SetRegionNodes

`func (o *DedicatedClusterCreateSpecification) SetRegionNodes(v map[string]int32)`

SetRegionNodes sets RegionNodes field to given value.

### GetRestrictEgressTraffic

`func (o *DedicatedClusterCreateSpecification) GetRestrictEgressTraffic() bool`

GetRestrictEgressTraffic returns the RestrictEgressTraffic field if non-nil, zero value otherwise.

### SetRestrictEgressTraffic

`func (o *DedicatedClusterCreateSpecification) SetRestrictEgressTraffic(v bool)`

SetRestrictEgressTraffic sets RestrictEgressTraffic field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


