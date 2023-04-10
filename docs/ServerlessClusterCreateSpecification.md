# ServerlessClusterCreateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryRegion** | Pointer to **string** | Preview: Specify which region should be made the primary region. This is only applicable to multi-region Serverless clusters. This field is required if you create the cluster in more than one region. | [optional] 
**Regions** | **[]string** | Region values should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. | 
**SpendLimit** | Pointer to **int32** | spend_limit is the maximum monthly charge for a cluster, in US cents. We recommend using usage_limits instead, since spend_limit will be deprecated in the future. | [optional] 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewServerlessClusterCreateSpecification

`func NewServerlessClusterCreateSpecification(regions []string, ) *ServerlessClusterCreateSpecification`

NewServerlessClusterCreateSpecification instantiates a new ServerlessClusterCreateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServerlessClusterCreateSpecificationWithDefaults

`func NewServerlessClusterCreateSpecificationWithDefaults() *ServerlessClusterCreateSpecification`

NewServerlessClusterCreateSpecificationWithDefaults instantiates a new ServerlessClusterCreateSpecification object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPrimaryRegion

`func (o *ServerlessClusterCreateSpecification) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### SetPrimaryRegion

`func (o *ServerlessClusterCreateSpecification) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### GetRegions

`func (o *ServerlessClusterCreateSpecification) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *ServerlessClusterCreateSpecification) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetSpendLimit

`func (o *ServerlessClusterCreateSpecification) GetSpendLimit() int32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### SetSpendLimit

`func (o *ServerlessClusterCreateSpecification) SetSpendLimit(v int32)`

SetSpendLimit sets SpendLimit field to given value.

### GetUsageLimits

`func (o *ServerlessClusterCreateSpecification) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *ServerlessClusterCreateSpecification) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


