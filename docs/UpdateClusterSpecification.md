# UpdateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterUpdateSpecification**](DedicatedClusterUpdateSpecification.md) |  | [optional] 
**ParentId** | Pointer to **string** | Limited Access: The parent ID is a folder ID. An empty string or \&quot;root\&quot; represents the root level. | [optional] 
**Shared** | Pointer to [**SharedClusterUpdateSpecification**](SharedClusterUpdateSpecification.md) |  | [optional] 
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

### GetParentId

`func (o *UpdateClusterSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *UpdateClusterSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetShared

`func (o *UpdateClusterSpecification) GetShared() SharedClusterUpdateSpecification`

GetShared returns the Shared field if non-nil, zero value otherwise.

### SetShared

`func (o *UpdateClusterSpecification) SetShared(v SharedClusterUpdateSpecification)`

SetShared sets Shared field to given value.

### GetUpgradeStatus

`func (o *UpdateClusterSpecification) GetUpgradeStatus() ClusterUpgradeStatusType`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### SetUpgradeStatus

`func (o *UpdateClusterSpecification) SetUpgradeStatus(v ClusterUpgradeStatusType)`

SetUpgradeStatus sets UpgradeStatus field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


