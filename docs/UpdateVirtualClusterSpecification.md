# UpdateVirtualClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to [**DeleteProtectionStateType**](DeleteProtectionStateType.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. If the labels field is included in the request, any existing labels on the cluster that are not included are removed, and any new labels specified are added. If the labels field is omitted from the request entirely, all existing labels remain unchanged. | [optional] 
**ParentId** | Pointer to **string** | The parent ID is a folder ID. An empty string or \&quot;root\&quot; moves the cluster to the root level. Omit to leave the cluster where it is. | [optional] 
**PrimaryRegion** | Pointer to **string** | Specify which region should be made the primary region. It must be one the cluster will span, either from the regions above or, when those are omitted, from the regions the cluster already has. When omitted, the current primary region is kept if it is still listed. | [optional] 
**Regions** | Pointer to **[]string** | The regions the cluster should span, as cloud provider region codes, for example \&quot;us-east-1\&quot;. Get the workspace to see which are available. Omit to leave the current regions alone. | [optional] 
**VirtualCpuBurstLimit** | Pointer to **int32** | Sets a custom value for the maximum vCPUs this cluster may use at once. Setting a custom vCPU burst limit requires edit permission on the host cluster backing the workspace. Omit to leave the current limit alone. | [optional] 

## Methods

### NewUpdateVirtualClusterSpecification

`func NewUpdateVirtualClusterSpecification() *UpdateVirtualClusterSpecification`

NewUpdateVirtualClusterSpecification instantiates a new UpdateVirtualClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDeleteProtection

`func (o *UpdateVirtualClusterSpecification) GetDeleteProtection() DeleteProtectionStateType`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### SetDeleteProtection

`func (o *UpdateVirtualClusterSpecification) SetDeleteProtection(v DeleteProtectionStateType)`

SetDeleteProtection sets DeleteProtection field to given value.

### GetLabels

`func (o *UpdateVirtualClusterSpecification) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *UpdateVirtualClusterSpecification) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetParentId

`func (o *UpdateVirtualClusterSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *UpdateVirtualClusterSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPrimaryRegion

`func (o *UpdateVirtualClusterSpecification) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### SetPrimaryRegion

`func (o *UpdateVirtualClusterSpecification) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### GetRegions

`func (o *UpdateVirtualClusterSpecification) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *UpdateVirtualClusterSpecification) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetVirtualCpuBurstLimit

`func (o *UpdateVirtualClusterSpecification) GetVirtualCpuBurstLimit() int32`

GetVirtualCpuBurstLimit returns the VirtualCpuBurstLimit field if non-nil, zero value otherwise.

### SetVirtualCpuBurstLimit

`func (o *UpdateVirtualClusterSpecification) SetVirtualCpuBurstLimit(v int32)`

SetVirtualCpuBurstLimit sets VirtualCpuBurstLimit field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


