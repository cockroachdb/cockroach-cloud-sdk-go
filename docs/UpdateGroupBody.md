# UpdateGroupBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]ScimResource**](ScimResource.md) |  | [optional] 
**Schemas** | **[]string** |  | 

## Methods

### NewUpdateGroupBody

`func NewUpdateGroupBody(displayName string, schemas []string, ) *UpdateGroupBody`

NewUpdateGroupBody instantiates a new UpdateGroupBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUpdateGroupBodyWithDefaults

`func NewUpdateGroupBodyWithDefaults() *UpdateGroupBody`

NewUpdateGroupBodyWithDefaults instantiates a new UpdateGroupBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplayName

`func (o *UpdateGroupBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *UpdateGroupBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetExternalId

`func (o *UpdateGroupBody) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *UpdateGroupBody) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetMembers

`func (o *UpdateGroupBody) GetMembers() []ScimResource`

GetMembers returns the Members field if non-nil, zero value otherwise.

### SetMembers

`func (o *UpdateGroupBody) SetMembers(v []ScimResource)`

SetMembers sets Members field to given value.

### GetSchemas

`func (o *UpdateGroupBody) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *UpdateGroupBody) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


