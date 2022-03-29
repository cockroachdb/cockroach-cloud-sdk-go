# DedicatedMachineTypeSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineType** | Pointer to **string** | MachineType is the machine type identifier within the given cloud provider, ex. m5.xlarge, n2-standard-4. | [optional] 
**NumVirtualCpus** | Pointer to **int32** | NumVirtualCPUs may be used to automatically select a machine type according to the desired number of vCPUs. | [optional] 

## Methods

### NewDedicatedMachineTypeSpecification

`func NewDedicatedMachineTypeSpecification() *DedicatedMachineTypeSpecification`

NewDedicatedMachineTypeSpecification instantiates a new DedicatedMachineTypeSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDedicatedMachineTypeSpecificationWithDefaults

`func NewDedicatedMachineTypeSpecificationWithDefaults() *DedicatedMachineTypeSpecification`

NewDedicatedMachineTypeSpecificationWithDefaults instantiates a new DedicatedMachineTypeSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineType

`func (o *DedicatedMachineTypeSpecification) GetMachineType() string`

GetMachineType returns the MachineType field if non-nil, zero value otherwise.

### GetMachineTypeOk

`func (o *DedicatedMachineTypeSpecification) GetMachineTypeOk() (*string, bool)`

GetMachineTypeOk returns a tuple with the MachineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineType

`func (o *DedicatedMachineTypeSpecification) SetMachineType(v string)`

SetMachineType sets MachineType field to given value.

### HasMachineType

`func (o *DedicatedMachineTypeSpecification) HasMachineType() bool`

HasMachineType returns a boolean if a field has been set.

### GetNumVirtualCpus

`func (o *DedicatedMachineTypeSpecification) GetNumVirtualCpus() int32`

GetNumVirtualCpus returns the NumVirtualCpus field if non-nil, zero value otherwise.

### GetNumVirtualCpusOk

`func (o *DedicatedMachineTypeSpecification) GetNumVirtualCpusOk() (*int32, bool)`

GetNumVirtualCpusOk returns a tuple with the NumVirtualCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumVirtualCpus

`func (o *DedicatedMachineTypeSpecification) SetNumVirtualCpus(v int32)`

SetNumVirtualCpus sets NumVirtualCpus field to given value.

### HasNumVirtualCpus

`func (o *DedicatedMachineTypeSpecification) HasNumVirtualCpus() bool`

HasNumVirtualCpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


