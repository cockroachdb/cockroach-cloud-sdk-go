# CreateSQLUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**SQLUser**](SQLUser.md) |  | 
**Password** | **string** |  | 

## Methods

### NewCreateSQLUserRequest

`func NewCreateSQLUserRequest(user SQLUser, password string, ) *CreateSQLUserRequest`

NewCreateSQLUserRequest instantiates a new CreateSQLUserRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateSQLUserRequestWithDefaults

`func NewCreateSQLUserRequestWithDefaults() *CreateSQLUserRequest`

NewCreateSQLUserRequestWithDefaults instantiates a new CreateSQLUserRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetUser

`func (o *CreateSQLUserRequest) GetUser() SQLUser`

GetUser returns the User field if non-nil, zero value otherwise.

### SetUser

`func (o *CreateSQLUserRequest) SetUser(v SQLUser)`

SetUser sets User field to given value.

### GetPassword

`func (o *CreateSQLUserRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### SetPassword

`func (o *CreateSQLUserRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


