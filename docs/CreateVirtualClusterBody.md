# CreateVirtualClusterBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to [**DeleteProtectionStateType**](DeleteProtectionStateType.md) |  | [optional] 
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. | [optional] 
**Name** | **string** | The name of the cluster. | 
**ParentId** | Pointer to **string** | The parent ID is a folder ID. An empty string or \&quot;root\&quot; will create a cluster at the root level. | [optional] 
**PrimaryRegion** | Pointer to **string** | Specify which region should be made the primary region. This field is required if the cluster spans more than one region, and must be one of the regions the cluster spans. | [optional] 
**Regions** | Pointer to **[]string** | regions is the subset of the workspace&#39;s regions to create this cluster in. Omit to use every region in the workspace. Gives an error if a region is specified that is not in the workspace. Values are the cloud provider&#39;s region codes, for example \&quot;us-east-1\&quot;. | [optional] 
**VirtualCpuBurstLimit** | Pointer to **int32** | Sets a custom value for the maximum vCPUs this cluster may use at once. Omit to inherit the workspace&#39;s default. Setting a custom vCPU burst limit requires edit permission on the workspace. | [optional] 

## Methods

### NewCreateVirtualClusterBody

`func NewCreateVirtualClusterBody(name string, ) *CreateVirtualClusterBody`

NewCreateVirtualClusterBody instantiates a new CreateVirtualClusterBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateVirtualClusterBodyWithDefaults

`func NewCreateVirtualClusterBodyWithDefaults() *CreateVirtualClusterBody`

NewCreateVirtualClusterBodyWithDefaults instantiates a new CreateVirtualClusterBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDeleteProtection

`func (o *CreateVirtualClusterBody) GetDeleteProtection() DeleteProtectionStateType`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### SetDeleteProtection

`func (o *CreateVirtualClusterBody) SetDeleteProtection(v DeleteProtectionStateType)`

SetDeleteProtection sets DeleteProtection field to given value.

### GetLabels

`func (o *CreateVirtualClusterBody) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *CreateVirtualClusterBody) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetName

`func (o *CreateVirtualClusterBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateVirtualClusterBody) SetName(v string)`

SetName sets Name field to given value.

### GetParentId

`func (o *CreateVirtualClusterBody) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *CreateVirtualClusterBody) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPrimaryRegion

`func (o *CreateVirtualClusterBody) GetPrimaryRegion() string`

GetPrimaryRegion returns the PrimaryRegion field if non-nil, zero value otherwise.

### SetPrimaryRegion

`func (o *CreateVirtualClusterBody) SetPrimaryRegion(v string)`

SetPrimaryRegion sets PrimaryRegion field to given value.

### GetRegions

`func (o *CreateVirtualClusterBody) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *CreateVirtualClusterBody) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetVirtualCpuBurstLimit

`func (o *CreateVirtualClusterBody) GetVirtualCpuBurstLimit() int32`

GetVirtualCpuBurstLimit returns the VirtualCpuBurstLimit field if non-nil, zero value otherwise.

### SetVirtualCpuBurstLimit

`func (o *CreateVirtualClusterBody) SetVirtualCpuBurstLimit(v int32)`

SetVirtualCpuBurstLimit sets VirtualCpuBurstLimit field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


