# CreateGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Members** | Pointer to [**[]ScimResource**](ScimResource.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCreateGroupRequest

`func NewCreateGroupRequest(displayName string, ) *CreateGroupRequest`

NewCreateGroupRequest instantiates a new CreateGroupRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateGroupRequestWithDefaults

`func NewCreateGroupRequestWithDefaults() *CreateGroupRequest`

NewCreateGroupRequestWithDefaults instantiates a new CreateGroupRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplayName

`func (o *CreateGroupRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *CreateGroupRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetExternalId

`func (o *CreateGroupRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *CreateGroupRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetMembers

`func (o *CreateGroupRequest) GetMembers() []ScimResource`

GetMembers returns the Members field if non-nil, zero value otherwise.

### SetMembers

`func (o *CreateGroupRequest) SetMembers(v []ScimResource)`

SetMembers sets Members field to given value.

### GetSchemas

`func (o *CreateGroupRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *CreateGroupRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


