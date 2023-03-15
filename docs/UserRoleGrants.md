# UserRoleGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Roles** | Pointer to [**[]BuiltInRole**](BuiltInRole.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserRoleGrants

`func NewUserRoleGrants() *UserRoleGrants`

NewUserRoleGrants instantiates a new UserRoleGrants object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

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


