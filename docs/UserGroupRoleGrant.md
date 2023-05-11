# UserGroupRoleGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupRoles** | Pointer to [**[]BuiltInFromGroups**](BuiltInFromGroups.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewUserGroupRoleGrant

`func NewUserGroupRoleGrant() *UserGroupRoleGrant`

NewUserGroupRoleGrant instantiates a new UserGroupRoleGrant object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetGroupRoles

`func (o *UserGroupRoleGrant) GetGroupRoles() []BuiltInFromGroups`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### SetGroupRoles

`func (o *UserGroupRoleGrant) SetGroupRoles(v []BuiltInFromGroups)`

SetGroupRoles sets GroupRoles field to given value.

### GetUserId

`func (o *UserGroupRoleGrant) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### SetUserId

`func (o *UserGroupRoleGrant) SetUserId(v string)`

SetUserId sets UserId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


