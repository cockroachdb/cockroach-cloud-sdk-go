# VirtualClusterWorkspace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**CreatedAt** | **time.Time** | The date and time the workspace was created. | 
**DefaultVirtualCpuBurstLimit** | **int32** | The vCPU burst limit a virtual cluster gets when its create request does not ask for a specific one. Changing this value in a workspace does not affect clusters that already exist in the workspace, only newly created clusters. | 
**Id** | **string** | The unique identifier of the workspace. | 
**Name** | **string** | The name of the workspace. | 
**Regions** | **[]string** | The regions the workspace spans. A virtual cluster created into it may use any subset of these and no others. | 
**UpdatedAt** | **time.Time** | The date and time the workspace was last modified. | 

## Methods

### NewVirtualClusterWorkspace

`func NewVirtualClusterWorkspace(cloudProvider CloudProviderType, createdAt time.Time, defaultVirtualCpuBurstLimit int32, id string, name string, regions []string, updatedAt time.Time, ) *VirtualClusterWorkspace`

NewVirtualClusterWorkspace instantiates a new VirtualClusterWorkspace object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewVirtualClusterWorkspaceWithDefaults

`func NewVirtualClusterWorkspaceWithDefaults() *VirtualClusterWorkspace`

NewVirtualClusterWorkspaceWithDefaults instantiates a new VirtualClusterWorkspace object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCloudProvider

`func (o *VirtualClusterWorkspace) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *VirtualClusterWorkspace) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### GetCreatedAt

`func (o *VirtualClusterWorkspace) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *VirtualClusterWorkspace) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetDefaultVirtualCpuBurstLimit

`func (o *VirtualClusterWorkspace) GetDefaultVirtualCpuBurstLimit() int32`

GetDefaultVirtualCpuBurstLimit returns the DefaultVirtualCpuBurstLimit field if non-nil, zero value otherwise.

### SetDefaultVirtualCpuBurstLimit

`func (o *VirtualClusterWorkspace) SetDefaultVirtualCpuBurstLimit(v int32)`

SetDefaultVirtualCpuBurstLimit sets DefaultVirtualCpuBurstLimit field to given value.

### GetId

`func (o *VirtualClusterWorkspace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *VirtualClusterWorkspace) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *VirtualClusterWorkspace) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *VirtualClusterWorkspace) SetName(v string)`

SetName sets Name field to given value.

### GetRegions

`func (o *VirtualClusterWorkspace) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *VirtualClusterWorkspace) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### GetUpdatedAt

`func (o *VirtualClusterWorkspace) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *VirtualClusterWorkspace) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


