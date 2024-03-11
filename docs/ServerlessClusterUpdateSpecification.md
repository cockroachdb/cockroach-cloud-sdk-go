# ServerlessClusterUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryRegion** | Pointer to **string** | Specify which region should be made the primary region. This is only applicable to multi-region Serverless clusters. This field is required if the regions field contains more than one region. | [optional] 
**Regions** | Pointer to **[]string** | Region values should match the cloud provider&#39;s zone code. For example, for Oregon, set region_name to \&quot;us-west2\&quot; for GCP and \&quot;us-west-2\&quot; for AWS. If this field is provided, the cluster&#39;s regions will be changed to match this list. Regions cannot currently be removed. | [optional] 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewServerlessClusterUpdateSpecification

`func NewServerlessClusterUpdateSpecification() *ServerlessClusterUpdateSpecification`

NewServerlessClusterUpdateSpecification instantiates a new ServerlessClusterUpdateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetPrimaryRegion

`func (o *ServerlessClusterUpdateSpecification) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### SetPrimaryRegion

`func (o *ServerlessClusterUpdateSpecification) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### GetRegions

`func (o *ServerlessClusterUpdateSpecification) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *ServerlessClusterUpdateSpecification) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetUsageLimits

`func (o *ServerlessClusterUpdateSpecification) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *ServerlessClusterUpdateSpecification) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


