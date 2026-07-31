# UserRoleGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | [**[]BuiltInRole**](BuiltInRole.md) | roles is the complete set of roles for the user, service account, or group. | 
**UserId** | **string** | user_id is the ID of the user, service account, or group. | 

## Methods

### NewUserRoleGrants

`func NewUserRoleGrants(roles []BuiltInRole, userId string, ) *UserRoleGrants`

NewUserRoleGrants instantiates a new UserRoleGrants object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUserRoleGrantsWithDefaults

`func NewUserRoleGrantsWithDefaults() *UserRoleGrants`

NewUserRoleGrantsWithDefaults instantiates a new UserRoleGrants object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRoles

`func (o *UserRoleGrants) GetRoles() []BuiltInRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### SetRoles

`func (o *UserRoleGrants) SetRoles(v []BuiltInRole)`

SetRoles sets Roles field to given value.

### GetUserId

`func (o *UserRoleGrants) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### SetUserId

`func (o *UserRoleGrants) SetUserId(v string)`

SetUserId sets UserId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


