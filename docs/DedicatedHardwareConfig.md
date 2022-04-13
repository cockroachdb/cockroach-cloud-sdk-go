# DedicatedHardwareConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineType** | **string** | MachineType is the machine type identifier within the given cloud provider, ex. m5.xlarge, n2-standard-4. | 
**NumVirtualCpus** | **int32** | NumVirtualCPUs is the number of virtual CPUs per node in the cluster. | 
**StorageGib** | **int32** | StorageGiB is the number of storage GiB per node in the cluster. | 
**MemoryGib** | **float32** | MemoryGiB is the memory GiB per node in the cluster. | 
**DiskIops** | **int32** | DiskIOPs is the number of disk I/O operations per second that are permitted on each node in the cluster. Zero indicates the cloud provider-specific default. | 

## Methods

### NewDedicatedHardwareConfig

`func NewDedicatedHardwareConfig(machineType string, numVirtualCpus int32, storageGib int32, memoryGib float32, diskIops int32, ) *DedicatedHardwareConfig`

NewDedicatedHardwareConfig instantiates a new DedicatedHardwareConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDedicatedHardwareConfigWithDefaults

`func NewDedicatedHardwareConfigWithDefaults() *DedicatedHardwareConfig`

NewDedicatedHardwareConfigWithDefaults instantiates a new DedicatedHardwareConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetMachineType

`func (o *DedicatedHardwareConfig) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### SetMachineType

`func (o *DedicatedHardwareConfig) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### GetNumVirtualCpus

`func (o *DedicatedHardwareConfig) GetNumVirtualCpus() int32`

GetNumVirtualCpus returns the NumVirtualCpus field if non-nil, zero value otherwise.

### SetNumVirtualCpus

`func (o *DedicatedHardwareConfig) SetNumVirtualCpus(v int32)`

SetNumVirtualCpus sets NumVirtualCpus field to given value.

### GetStorageGib

`func (o *DedicatedHardwareConfig) GetStorageGib() int32`

GetStorageGib returns the StorageGib field if non-nil, zero value otherwise.

### SetStorageGib

`func (o *DedicatedHardwareConfig) SetStorageGib(v int32)`

SetStorageGib sets StorageGib field to given value.

### GetMemoryGib

`func (o *DedicatedHardwareConfig) GetMemoryGib() float32`

GetMemoryGib returns the MemoryGib field if non-nil, zero value otherwise.

### SetMemoryGib

`func (o *DedicatedHardwareConfig) SetMemoryGib(v float32)`

SetMemoryGib sets MemoryGib field to given value.

### GetDiskIops

`func (o *DedicatedHardwareConfig) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### SetDiskIops

`func (o *DedicatedHardwareConfig) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


