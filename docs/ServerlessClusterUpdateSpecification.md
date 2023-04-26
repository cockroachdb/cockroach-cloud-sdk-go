# ServerlessClusterUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendLimit** | Pointer to **Nullable[int32]** | spend_limit is the maximum monthly charge for a cluster, in US cents. We recommend using usage_limits instead, since spend_limit will be deprecated in the future. | [optional] 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewServerlessClusterUpdateSpecification

`func NewServerlessClusterUpdateSpecification() *ServerlessClusterUpdateSpecification`

NewServerlessClusterUpdateSpecification instantiates a new ServerlessClusterUpdateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetSpendLimit

`func (o *ServerlessClusterUpdateSpecification) GetSpendLimit() int32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### SetSpendLimit

`func (o *ServerlessClusterUpdateSpecification) SetSpendLimit(v int32)`

SetSpendLimit sets SpendLimit field to given value.

### GetUsageLimits

`func (o *ServerlessClusterUpdateSpecification) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *ServerlessClusterUpdateSpecification) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


