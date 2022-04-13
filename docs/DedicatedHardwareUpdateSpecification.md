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

NewDedicatedHardwareUpdateSpecification instantiates a new DedicatedHardwareUpdateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetMachineSpec

`func (o *DedicatedHardwareUpdateSpecification) GetMachineSpec() DedicatedMachineTypeSpecification`

GetMachineSpec returns the MachineSpec field if non-nil, zero value otherwise.

### SetMachineSpec

`func (o *DedicatedHardwareUpdateSpecification) SetMachineSpec(v DedicatedMachineTypeSpecification)`

SetMachineSpec sets MachineSpec field to given value.

### GetStorageGib

`func (o *DedicatedHardwareUpdateSpecification) GetStorageGib() int32`

GetStorageGib returns the StorageGib field if non-nil, zero value otherwise.

### SetStorageGib

`func (o *DedicatedHardwareUpdateSpecification) SetStorageGib(v int32)`

SetStorageGib sets StorageGib field to given value.

### GetDiskIops

`func (o *DedicatedHardwareUpdateSpecification) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### SetDiskIops

`func (o *DedicatedHardwareUpdateSpecification) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


