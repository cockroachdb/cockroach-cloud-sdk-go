# ServerlessClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | **[]string** | Region values should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. | 
**SpendLimit** | **int32** |  | 

## Methods

### NewServerlessClusterSpecification

`func NewServerlessClusterSpecification(regions []string, spendLimit int32, ) *ServerlessClusterSpecification`

NewServerlessClusterSpecification instantiates a new ServerlessClusterSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerlessClusterSpecificationWithDefaults

`func NewServerlessClusterSpecificationWithDefaults() *ServerlessClusterSpecification`

NewServerlessClusterSpecificationWithDefaults instantiates a new ServerlessClusterSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *ServerlessClusterSpecification) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *ServerlessClusterSpecification) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *ServerlessClusterSpecification) SetRegions(v []string)`

SetRegions sets Regions field to given value.


### GetSpendLimit

`func (o *ServerlessClusterSpecification) GetSpendLimit() int32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### GetSpendLimitOk

`func (o *ServerlessClusterSpecification) GetSpendLimitOk() (*int32, bool)`

GetSpendLimitOk returns a tuple with the SpendLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendLimit

`func (o *ServerlessClusterSpecification) SetSpendLimit(v int32)`

SetSpendLimit sets SpendLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


