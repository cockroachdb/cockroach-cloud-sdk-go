# ServerlessClusterCreateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | **[]string** | Region values should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. | 
**SpendLimit** | **int32** |  | 

## Methods

### NewServerlessClusterCreateSpecification

`func NewServerlessClusterCreateSpecification(regions []string, spendLimit int32, ) *ServerlessClusterCreateSpecification`

NewServerlessClusterCreateSpecification instantiates a new ServerlessClusterCreateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServerlessClusterCreateSpecificationWithDefaults

`func NewServerlessClusterCreateSpecificationWithDefaults() *ServerlessClusterCreateSpecification`

NewServerlessClusterCreateSpecificationWithDefaults instantiates a new ServerlessClusterCreateSpecification object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


