# CreateUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Emails** | [**[]ScimEmail**](ScimEmail.md) |  | 
**ExternalId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to [**ScimName**](ScimName.md) |  | [optional] 
**Schemas** | **[]string** |  | 
**UserName** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserRequest

`func NewCreateUserRequest(emails []ScimEmail, schemas []string, ) *CreateUserRequest`

NewCreateUserRequest instantiates a new CreateUserRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateUserRequestWithDefaults

`func NewCreateUserRequestWithDefaults() *CreateUserRequest`

NewCreateUserRequestWithDefaults instantiates a new CreateUserRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetActive

`func (o *CreateUserRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### SetActive

`func (o *CreateUserRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### GetDisplayName

`func (o *CreateUserRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### SetDisplayName

`func (o *CreateUserRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### GetEmails

`func (o *CreateUserRequest) GetEmails() []ScimEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### SetEmails

`func (o *CreateUserRequest) SetEmails(v []ScimEmail)`

SetEmails sets Emails field to given value.

### GetExternalId

`func (o *CreateUserRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *CreateUserRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetName

`func (o *CreateUserRequest) GetName() ScimName`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateUserRequest) SetName(v ScimName)`

SetName sets Name field to given value.

### GetSchemas

`func (o *CreateUserRequest) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *CreateUserRequest) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetUserName

`func (o *CreateUserRequest) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### SetUserName

`func (o *CreateUserRequest) SetUserName(v string)`

SetUserName sets UserName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


