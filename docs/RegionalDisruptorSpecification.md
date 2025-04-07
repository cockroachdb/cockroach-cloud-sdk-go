# RegionalDisruptorSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azs** | Pointer to **[]string** |  | [optional] 
**IsWholeRegion** | **bool** | is_whole_region denotes whether the whole region should be disrupted.  If a specific pod or az is passed, this value should be false. | 
**Pods** | Pointer to **[]string** | pods represents each individual CRDB pod name that should be disrupted. Pod names can be found using the [ListClusterNodes api](https://www.cockroachlabs.com/docs/api/cloud/v1#get-/api/v1/clusters/-cluster_id-/nodes). | [optional] 
**RegionCode** | **string** | region_code is the cloud provider specific region code of the region. (i.e us-east1 for gcp, us-east-1 for aws, eastus for azure). Region names for each node the cluster can be found using the [ListClusterNodes api](https://www.cockroachlabs.com/docs/api/cloud/v1#get-/api/v1/clusters/-cluster_id-/nodes). | 

## Methods

### NewRegionalDisruptorSpecification

`func NewRegionalDisruptorSpecification(isWholeRegion bool, regionCode string, ) *RegionalDisruptorSpecification`

NewRegionalDisruptorSpecification instantiates a new RegionalDisruptorSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRegionalDisruptorSpecificationWithDefaults

`func NewRegionalDisruptorSpecificationWithDefaults() *RegionalDisruptorSpecification`

NewRegionalDisruptorSpecificationWithDefaults instantiates a new RegionalDisruptorSpecification object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAzs

`func (o *RegionalDisruptorSpecification) GetAzs() []string`

GetAzs returns the Azs field if non-nil, zero value otherwise.

### SetAzs

`func (o *RegionalDisruptorSpecification) SetAzs(v []string)`

SetAzs sets Azs field to given value.

### GetIsWholeRegion

`func (o *RegionalDisruptorSpecification) GetIsWholeRegion() bool`

GetIsWholeRegion returns the IsWholeRegion field if non-nil, zero value otherwise.

### SetIsWholeRegion

`func (o *RegionalDisruptorSpecification) SetIsWholeRegion(v bool)`

SetIsWholeRegion sets IsWholeRegion field to given value.

### GetPods

`func (o *RegionalDisruptorSpecification) GetPods() []string`

GetPods returns the Pods field if non-nil, zero value otherwise.

### SetPods

`func (o *RegionalDisruptorSpecification) SetPods(v []string)`

SetPods sets Pods field to given value.

### GetRegionCode

`func (o *RegionalDisruptorSpecification) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### SetRegionCode

`func (o *RegionalDisruptorSpecification) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


