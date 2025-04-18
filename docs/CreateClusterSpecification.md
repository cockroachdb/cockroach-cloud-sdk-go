# CreateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterCreateSpecification**](DedicatedClusterCreateSpecification.md) |  | [optional] 
**DeleteProtection** | Pointer to [**DeleteProtectionStateType**](DeleteProtectionStateType.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. | [optional] 
**ParentId** | Pointer to **string** | Preview: The parent ID is a folder ID. An empty string or \&quot;root\&quot; will create a cluster at the root level. | [optional] 
**Plan** | Pointer to [**PlanType**](PlanType.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterCreateSpecification**](ServerlessClusterCreateSpecification.md) |  | [optional] 

## Methods

### NewCreateClusterSpecification

`func NewCreateClusterSpecification() *CreateClusterSpecification`

NewCreateClusterSpecification instantiates a new CreateClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDedicated

`func (o *CreateClusterSpecification) GetDedicated() DedicatedClusterCreateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### SetDedicated

`func (o *CreateClusterSpecification) SetDedicated(v DedicatedClusterCreateSpecification)`

SetDedicated sets Dedicated field to given value.

### GetDeleteProtection

`func (o *CreateClusterSpecification) GetDeleteProtection() DeleteProtectionStateType`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### SetDeleteProtection

`func (o *CreateClusterSpecification) SetDeleteProtection(v DeleteProtectionStateType)`

SetDeleteProtection sets DeleteProtection field to given value.

### GetLabels

`func (o *CreateClusterSpecification) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *CreateClusterSpecification) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetParentId

`func (o *CreateClusterSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *CreateClusterSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPlan

`func (o *CreateClusterSpecification) GetPlan() PlanType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### SetPlan

`func (o *CreateClusterSpecification) SetPlan(v PlanType)`

SetPlan sets Plan field to given value.

### GetServerless

`func (o *CreateClusterSpecification) GetServerless() ServerlessClusterCreateSpecification`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### SetServerless

`func (o *CreateClusterSpecification) SetServerless(v ServerlessClusterCreateSpecification)`

SetServerless sets Serverless field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


