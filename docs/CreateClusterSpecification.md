# CreateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterCreateSpecification**](DedicatedClusterCreateSpecification.md) |  | [optional] 
**ParentId** | Pointer to **string** | Limited Access: The parent ID is a folder ID. An empty string or \&quot;root\&quot; will create a cluster at the root level. | [optional] 
**Shared** | Pointer to [**SharedClusterCreateSpecification**](SharedClusterCreateSpecification.md) |  | [optional] 

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

### GetParentId

`func (o *CreateClusterSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *CreateClusterSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetShared

`func (o *CreateClusterSpecification) GetShared() SharedClusterCreateSpecification`

GetShared returns the Shared field if non-nil, zero value otherwise.

### SetShared

`func (o *CreateClusterSpecification) SetShared(v SharedClusterCreateSpecification)`

SetShared sets Shared field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


