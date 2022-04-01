# DedicatedHardwareUpdateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineSpec** | Pointer to [**DedicatedMachineTypeSpecification**](DedicatedMachineTypeSpecification.md) |  | [optional] 
**StorageGib** | Pointer to **int32** | StorageGiB is the number of storage GiB per node in the cluster. | [optional] 
**DiskIops** | Pointer to **int32** | DiskIOPs is the number of disk I/O operations per second that are permitted on each node in the cluster. Zero indicates the cloud provider-specific default. Only available for AWS clusters. | [optional] 

## Methods

### NewDedicatedHardwareUpdateSpecification

`func NewDedicatedHardwareUpdateSpecification() *DedicatedHardwareUpdateSpecification`

NewDedicatedHardwareUpdateSpecification instantiates a new DedicatedHardwareUpdateSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedHardwareUpdateSpecificationWithDefaults

`func NewDedicatedHardwareUpdateSpecificationWithDefaults() *DedicatedHardwareUpdateSpecification`

NewDedicatedHardwareUpdateSpecificationWithDefaults instantiates a new DedicatedHardwareUpdateSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineSpec

`func (o *DedicatedHardwareUpdateSpecification) GetMachineSpec() DedicatedMachineTypeSpecification`

GetMachineSpec returns the MachineSpec field if non-nil, zero value otherwise.

### GetMachineSpecOk

`func (o *DedicatedHardwareUpdateSpecification) GetMachineSpecOk() (*DedicatedMachineTypeSpecification, bool)`

GetMachineSpecOk returns a tuple with the MachineSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineSpec

`func (o *DedicatedHardwareUpdateSpecification) SetMachineSpec(v DedicatedMachineTypeSpecification)`

SetMachineSpec sets MachineSpec field to given value.

### HasMachineSpec

`func (o *DedicatedHardwareUpdateSpecification) HasMachineSpec() bool`

HasMachineSpec returns a boolean if a field has been set.

### GetStorageGib

`func (o *DedicatedHardwareUpdateSpecification) GetStorageGib() int32`

GetStorageGib returns the StorageGib field if non-nil, zero value otherwise.

### GetStorageGibOk

`func (o *DedicatedHardwareUpdateSpecification) GetStorageGibOk() (*int32, bool)`

GetStorageGibOk returns a tuple with the StorageGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGib

`func (o *DedicatedHardwareUpdateSpecification) SetStorageGib(v int32)`

SetStorageGib sets StorageGib field to given value.

### HasStorageGib

`func (o *DedicatedHardwareUpdateSpecification) HasStorageGib() bool`

HasStorageGib returns a boolean if a field has been set.

### GetDiskIops

`func (o *DedicatedHardwareUpdateSpecification) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *DedicatedHardwareUpdateSpecification) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *DedicatedHardwareUpdateSpecification) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *DedicatedHardwareUpdateSpecification) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


