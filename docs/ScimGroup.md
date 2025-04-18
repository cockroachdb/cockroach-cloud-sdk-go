# ScimGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Members** | Pointer to [**[]ScimResource**](ScimResource.md) |  | [optional] 
**Meta** | Pointer to [**ScimMetadata**](ScimMetadata.md) |  | [optional] 
**Schemas** | **[]string** |  | 

## Methods

### NewScimGroup

`func NewScimGroup(displayName string, id string, schemas []string, ) *ScimGroup`

NewScimGroup instantiates a new ScimGroup object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewScimGroupWithDefaults

`func NewScimGroupWithDefaults() *ScimGroup`

NewScimGroupWithDefaults instantiates a new ScimGroup object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplayName

`func (o *ScimGroup) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *ScimGroup) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetExternalId

`func (o *ScimGroup) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *ScimGroup) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetId

`func (o *ScimGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ScimGroup) SetId(v string)`

SetId sets Id field to given value.

### GetMembers

`func (o *ScimGroup) GetMembers() []ScimResource`

GetMembers returns the Members field if non-nil, zero value otherwise.

### SetMembers

`func (o *ScimGroup) SetMembers(v []ScimResource)`

SetMembers sets Members field to given value.

### GetMeta

`func (o *ScimGroup) GetMeta() ScimMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### SetMeta

`func (o *ScimGroup) SetMeta(v ScimMetadata)`

SetMeta sets Meta field to given value.

### GetSchemas

`func (o *ScimGroup) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *ScimGroup) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


