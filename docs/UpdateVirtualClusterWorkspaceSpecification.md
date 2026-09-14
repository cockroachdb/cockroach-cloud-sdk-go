# UpdateVirtualClusterWorkspaceSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVirtualCpuBurstLimit** | Pointer to **int32** | Omitted leaves the current default. Applies only to virtual clusters created after this change. | [optional] 
**Name** | Pointer to **string** | Omitted leaves the current name. | [optional] 

## Methods

### NewUpdateVirtualClusterWorkspaceSpecification

`func NewUpdateVirtualClusterWorkspaceSpecification() *UpdateVirtualClusterWorkspaceSpecification`

NewUpdateVirtualClusterWorkspaceSpecification instantiates a new UpdateVirtualClusterWorkspaceSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDefaultVirtualCpuBurstLimit

`func (o *UpdateVirtualClusterWorkspaceSpecification) GetDefaultVirtualCpuBurstLimit() int32`

GetDefaultVirtualCpuBurstLimit returns the DefaultVirtualCpuBurstLimit field if non-nil, zero value otherwise.

### SetDefaultVirtualCpuBurstLimit

`func (o *UpdateVirtualClusterWorkspaceSpecification) SetDefaultVirtualCpuBurstLimit(v int32)`

SetDefaultVirtualCpuBurstLimit sets DefaultVirtualCpuBurstLimit field to given value.

### GetName

`func (o *UpdateVirtualClusterWorkspaceSpecification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *UpdateVirtualClusterWorkspaceSpecification) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


