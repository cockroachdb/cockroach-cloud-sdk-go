# SharedClusterUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryRegion** | Pointer to **string** | Specify which region should be made the primary region. This is only applicable to multi-region shared clusters. This field is required if the regions field contains more than one region. | [optional] 
**Regions** | Pointer to **[]string** | Region values should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. If this field is provided, the cluster&#39;s regions will be changed to match this list. Regions cannot currently be removed. | [optional] 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewSharedClusterUpdateSpecification

`func NewSharedClusterUpdateSpecification() *SharedClusterUpdateSpecification`

NewSharedClusterUpdateSpecification instantiates a new SharedClusterUpdateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetPrimaryRegion

`func (o *SharedClusterUpdateSpecification) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### SetPrimaryRegion

`func (o *SharedClusterUpdateSpecification) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### GetRegions

`func (o *SharedClusterUpdateSpecification) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *SharedClusterUpdateSpecification) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetUsageLimits

`func (o *SharedClusterUpdateSpecification) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *SharedClusterUpdateSpecification) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


