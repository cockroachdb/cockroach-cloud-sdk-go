# UpdateUserBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Emails** | Pointer to [**[]ScimEmail**](ScimEmail.md) |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**ScimName**](ScimName.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserBody

`func NewUpdateUserBody() *UpdateUserBody`

NewUpdateUserBody instantiates a new UpdateUserBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetActive

`func (o *UpdateUserBody) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### SetActive

`func (o *UpdateUserBody) SetActive(v bool)`

SetActive sets Active field to given value.

### GetDisplayName

`func (o *UpdateUserBody) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *UpdateUserBody) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetEmails

`func (o *UpdateUserBody) GetEmails() []ScimEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### SetEmails

`func (o *UpdateUserBody) SetEmails(v []ScimEmail)`

SetEmails sets Emails field to given value.

### GetExternalId

`func (o *UpdateUserBody) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *UpdateUserBody) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetName

`func (o *UpdateUserBody) GetName() ScimName`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *UpdateUserBody) SetName(v ScimName)`

SetName sets Name field to given value.

### GetSchemas

`func (o *UpdateUserBody) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *UpdateUserBody) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetUserName

`func (o *UpdateUserBody) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### SetUserName

`func (o *UpdateUserBody) SetUserName(v string)`

SetUserName sets UserName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


