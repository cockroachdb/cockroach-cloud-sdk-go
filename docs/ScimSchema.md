# ScimSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**[]ScimSchemaAttribute**](ScimSchemaAttribute.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ScimMetadata**](ScimMetadata.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewScimSchema

`func NewScimSchema() *ScimSchema`

NewScimSchema instantiates a new ScimSchema object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAttributes

`func (o *ScimSchema) GetAttributes() []ScimSchemaAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### SetAttributes

`func (o *ScimSchema) SetAttributes(v []ScimSchemaAttribute)`

SetAttributes sets Attributes field to given value.

### GetDescription

`func (o *ScimSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *ScimSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetId

`func (o *ScimSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ScimSchema) SetId(v string)`

SetId sets Id field to given value.

### GetMeta

`func (o *ScimSchema) GetMeta() ScimMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### SetMeta

`func (o *ScimSchema) SetMeta(v ScimMetadata)`

SetMeta sets Meta field to given value.

### GetName

`func (o *ScimSchema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *ScimSchema) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


