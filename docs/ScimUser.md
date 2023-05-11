# ScimUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]ScimEmail**](ScimEmail.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Groups** | Pointer to [**[]ScimResource**](ScimResource.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Meta** | Pointer to [**ScimMetadata**](ScimMetadata.md) |  | [optional] 
**Name** | Pointer to [**ScimName**](ScimName.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewScimUser

`func NewScimUser() *ScimUser`

NewScimUser instantiates a new ScimUser object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetActive

`func (o *ScimUser) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### SetActive

`func (o *ScimUser) SetActive(v bool)`

SetActive sets Active field to given value.

### GetDisplayName

`func (o *ScimUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *ScimUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetEmails

`func (o *ScimUser) GetEmails() []ScimEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### SetEmails

`func (o *ScimUser) SetEmails(v []ScimEmail)`

SetEmails sets Emails field to given value.

### GetExternalId

`func (o *ScimUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *ScimUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetGroups

`func (o *ScimUser) GetGroups() []ScimResource`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### SetGroups

`func (o *ScimUser) SetGroups(v []ScimResource)`

SetGroups sets Groups field to given value.

### GetId

`func (o *ScimUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ScimUser) SetId(v string)`

SetId sets Id field to given value.

### GetMeta

`func (o *ScimUser) GetMeta() ScimMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### SetMeta

`func (o *ScimUser) SetMeta(v ScimMetadata)`

SetMeta sets Meta field to given value.

### GetName

`func (o *ScimUser) GetName() ScimName`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *ScimUser) SetName(v ScimName)`

SetName sets Name field to given value.

### GetSchemas

`func (o *ScimUser) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *ScimUser) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetUserName

`func (o *ScimUser) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### SetUserName

`func (o *ScimUser) SetUserName(v string)`

SetUserName sets UserName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


