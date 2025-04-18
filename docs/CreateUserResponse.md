# CreateUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]ScimEmail**](ScimEmail.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]ScimResource**](ScimResource.md) |  | [optional] 
**Id** | **string** |  | 
**Meta** | Pointer to [**ScimMetadata**](ScimMetadata.md) |  | [optional] 
**Name** | Pointer to [**ScimName**](ScimName.md) |  | [optional] 
**Schemas** | **[]string** |  | 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserResponse

`func NewCreateUserResponse(id string, schemas []string, ) *CreateUserResponse`

NewCreateUserResponse instantiates a new CreateUserResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateUserResponseWithDefaults

`func NewCreateUserResponseWithDefaults() *CreateUserResponse`

NewCreateUserResponseWithDefaults instantiates a new CreateUserResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetActive

`func (o *CreateUserResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### SetActive

`func (o *CreateUserResponse) SetActive(v bool)`

SetActive sets Active field to given value.

### GetDisplayName

`func (o *CreateUserResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *CreateUserResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetEmails

`func (o *CreateUserResponse) GetEmails() []ScimEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### SetEmails

`func (o *CreateUserResponse) SetEmails(v []ScimEmail)`

SetEmails sets Emails field to given value.

### GetExternalId

`func (o *CreateUserResponse) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *CreateUserResponse) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetGroups

`func (o *CreateUserResponse) GetGroups() []ScimResource`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### SetGroups

`func (o *CreateUserResponse) SetGroups(v []ScimResource)`

SetGroups sets Groups field to given value.

### GetId

`func (o *CreateUserResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *CreateUserResponse) SetId(v string)`

SetId sets Id field to given value.

### GetMeta

`func (o *CreateUserResponse) GetMeta() ScimMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### SetMeta

`func (o *CreateUserResponse) SetMeta(v ScimMetadata)`

SetMeta sets Meta field to given value.

### GetName

`func (o *CreateUserResponse) GetName() ScimName`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateUserResponse) SetName(v ScimName)`

SetName sets Name field to given value.

### GetSchemas

`func (o *CreateUserResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *CreateUserResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetUserName

`func (o *CreateUserResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### SetUserName

`func (o *CreateUserResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


