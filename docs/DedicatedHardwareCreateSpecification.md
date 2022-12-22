# DedicatedHardwareCreateSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskIops** | Pointer to **int32** | disk_iops is the number of disk I/O operations per second that are permitted on each node in the cluster. Zero indicates the cloud provider-specific default. Only available for AWS clusters. | [optional] 
**MachineSpec** | [**DedicatedMachineTypeSpecification**](DedicatedMachineTypeSpecification.md) |  | 
**StorageGib** | **int32** | storage_gib is the number of storage GiB per node in the cluster. Zero indicates default to the lowest storage GiB available given machine specs. | 

## Methods

### NewDedicatedHardwareCreateSpecification

`func NewDedicatedHardwareCreateSpecification(machineSpec DedicatedMachineTypeSpecification, storageGib int32, ) *DedicatedHardwareCreateSpecification`

NewDedicatedHardwareCreateSpecification instantiates a new DedicatedHardwareCreateSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDedicatedHardwareCreateSpecificationWithDefaults

`func NewDedicatedHardwareCreateSpecificationWithDefaults() *DedicatedHardwareCreateSpecification`

NewDedicatedHardwareCreateSpecificationWithDefaults instantiates a new DedicatedHardwareCreateSpecification object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDiskIops

`func (o *DedicatedHardwareCreateSpecification) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### SetDiskIops

`func (o *DedicatedHardwareCreateSpecification) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### GetMachineSpec

`func (o *DedicatedHardwareCreateSpecification) GetMachineSpec() DedicatedMachineTypeSpecification`

GetMachineSpec returns the MachineSpec field if non-nil, zero value otherwise.

### SetMachineSpec

`func (o *DedicatedHardwareCreateSpecification) SetMachineSpec(v DedicatedMachineTypeSpecification)`

SetMachineSpec sets MachineSpec field to given value.

### GetStorageGib

`func (o *DedicatedHardwareCreateSpecification) GetStorageGib() int32`

GetStorageGib returns the StorageGib field if non-nil, zero value otherwise.

### SetStorageGib

`func (o *DedicatedHardwareCreateSpecification) SetStorageGib(v int32)`

SetStorageGib sets StorageGib field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


