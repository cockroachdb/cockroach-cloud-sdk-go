# PatchGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operations** | Pointer to [**[]ScimOperations**](ScimOperations.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPatchGroupRequest

`func NewPatchGroupRequest() *PatchGroupRequest`

NewPatchGroupRequest instantiates a new PatchGroupRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetOperations

`func (o *PatchGroupRequest) GetOperations() []ScimOperations`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### SetOperations

`func (o *PatchGroupRequest) SetOperations(v []ScimOperations)`

SetOperations sets Operations field to given value.

### GetSchemas

`func (o *PatchGroupRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *PatchGroupRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


