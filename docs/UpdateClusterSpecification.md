# UpdateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterUpdateSpecification**](DedicatedClusterUpdateSpecification.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterUpdateSpecification**](ServerlessClusterUpdateSpecification.md) |  | [optional] 
**UpgradeStatus** | Pointer to [**ClusterUpgradeStatusType**](ClusterUpgradeStatusType.md) |  | [optional] 

## Methods

### NewUpdateClusterSpecification

`func NewUpdateClusterSpecification() *UpdateClusterSpecification`

NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDedicated

`func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### SetDedicated

`func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification)`

SetDedicated sets Dedicated field to given value.

### GetServerless

`func (o *UpdateClusterSpecification) GetServerless() ServerlessClusterUpdateSpecification`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### SetServerless

`func (o *UpdateClusterSpecification) SetServerless(v ServerlessClusterUpdateSpecification)`

SetServerless sets Serverless field to given value.

### GetUpgradeStatus

`func (o *UpdateClusterSpecification) GetUpgradeStatus() ClusterUpgradeStatusType`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### SetUpgradeStatus

`func (o *UpdateClusterSpecification) SetUpgradeStatus(v ClusterUpgradeStatusType)`

SetUpgradeStatus sets UpgradeStatus field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


