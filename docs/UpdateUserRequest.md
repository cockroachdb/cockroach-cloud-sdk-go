# UpdateUserRequest

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

### NewUpdateUserRequest

`func NewUpdateUserRequest() *UpdateUserRequest`

NewUpdateUserRequest instantiates a new UpdateUserRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetActive

`func (o *UpdateUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### SetActive

`func (o *UpdateUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### GetDisplayName

`func (o *UpdateUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *UpdateUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetEmails

`func (o *UpdateUserRequest) GetEmails() []ScimEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### SetEmails

`func (o *UpdateUserRequest) SetEmails(v []ScimEmail)`

SetEmails sets Emails field to given value.

### GetExternalId

`func (o *UpdateUserRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *UpdateUserRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetName

`func (o *UpdateUserRequest) GetName() ScimName`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *UpdateUserRequest) SetName(v ScimName)`

SetName sets Name field to given value.

### GetSchemas

`func (o *UpdateUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *UpdateUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetUserName

`func (o *UpdateUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### SetUserName

`func (o *UpdateUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


