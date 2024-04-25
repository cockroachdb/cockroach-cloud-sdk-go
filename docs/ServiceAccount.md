# ServiceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | created_at represents the creation time of the service account. | 
**CreatorName** | **string** | creator_name is the name of user that created the service account. | 
**Description** | **string** | description of the service account. | 
**GroupRoles** | [**[]BuiltInFromGroups**](BuiltInFromGroups.md) | roles that the service account has inherited via a group. | 
**Id** | **string** | The ID of the service account. | 
**Name** | **string** | name of the service account. | 
**Roles** | [**[]BuiltInRole**](BuiltInRole.md) | roles that are assigned to the service account. | 

## Methods

### NewServiceAccount

`func NewServiceAccount(createdAt time.Time, creatorName string, description string, groupRoles []BuiltInFromGroups, id string, name string, roles []BuiltInRole, ) *ServiceAccount`

NewServiceAccount instantiates a new ServiceAccount object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServiceAccountWithDefaults

`func NewServiceAccountWithDefaults() *ServiceAccount`

NewServiceAccountWithDefaults instantiates a new ServiceAccount object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *ServiceAccount) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *ServiceAccount) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetCreatorName

`func (o *ServiceAccount) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### SetCreatorName

`func (o *ServiceAccount) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### GetDescription

`func (o *ServiceAccount) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *ServiceAccount) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetGroupRoles

`func (o *ServiceAccount) GetGroupRoles() []BuiltInFromGroups`

GetGroupRoles returns the GroupRoles field if non-nil, zero value otherwise.

### SetGroupRoles

`func (o *ServiceAccount) SetGroupRoles(v []BuiltInFromGroups)`

SetGroupRoles sets GroupRoles field to given value.

### GetId

`func (o *ServiceAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ServiceAccount) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *ServiceAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *ServiceAccount) SetName(v string)`

SetName sets Name field to given value.

### GetRoles

`func (o *ServiceAccount) GetRoles() []BuiltInRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### SetRoles

`func (o *ServiceAccount) SetRoles(v []BuiltInRole)`

SetRoles sets Roles field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


