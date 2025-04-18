# UpdateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CockroachVersion** | Pointer to **string** | The desired CockroachDB major version for the cluster.  It can be used to orchestrate version changes.  Setting the version to a later version will initiate an upgrade to that version.  After an upgrade is initiated but before it&#39;s finalized, setting the version back to the previous version will initiate a rollback. | [optional] 
**Dedicated** | Pointer to [**DedicatedClusterUpdateSpecification**](DedicatedClusterUpdateSpecification.md) |  | [optional] 
**DeleteProtection** | Pointer to [**DeleteProtectionStateType**](DeleteProtectionStateType.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. If the labels field is included in the request: Any existing labels on the cluster that are not included will be removed, and any new labels specified will be added. If the labels field is omitted from the request entirely, all existing labels will remain unchanged. | [optional] 
**ParentId** | Pointer to **string** | Preview: The parent ID is a folder ID. An empty string or \&quot;root\&quot; represents the root level. | [optional] 
**Plan** | Pointer to [**PlanType**](PlanType.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterUpdateSpecification**](ServerlessClusterUpdateSpecification.md) |  | [optional] 
**UpgradeStatus** | Pointer to [**ClusterUpgradeStatusType**](ClusterUpgradeStatusType.md) |  | [optional] 

## Methods

### NewUpdateClusterSpecification

`func NewUpdateClusterSpecification() *UpdateClusterSpecification`

NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetCockroachVersion

`func (o *UpdateClusterSpecification) GetCockroachVersion() string`

GetCockroachVersion returns the CockroachVersion field if non-nil, zero value otherwise.

### SetCockroachVersion

`func (o *UpdateClusterSpecification) SetCockroachVersion(v string)`

SetCockroachVersion sets CockroachVersion field to given value.

### GetDedicated

`func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### SetDedicated

`func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification)`

SetDedicated sets Dedicated field to given value.

### GetDeleteProtection

`func (o *UpdateClusterSpecification) GetDeleteProtection() DeleteProtectionStateType`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### SetDeleteProtection

`func (o *UpdateClusterSpecification) SetDeleteProtection(v DeleteProtectionStateType)`

SetDeleteProtection sets DeleteProtection field to given value.

### GetLabels

`func (o *UpdateClusterSpecification) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *UpdateClusterSpecification) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetParentId

`func (o *UpdateClusterSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *UpdateClusterSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPlan

`func (o *UpdateClusterSpecification) GetPlan() PlanType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### SetPlan

`func (o *UpdateClusterSpecification) SetPlan(v PlanType)`

SetPlan sets Plan field to given value.

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


