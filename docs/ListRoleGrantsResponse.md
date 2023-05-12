# ListRoleGrantsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grants** | Pointer to [**[]UserRoleGrants**](UserRoleGrants.md) |  | [optional] 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**UserGroupGrants** | Pointer to [**[]UserGroupRoleGrant**](UserGroupRoleGrant.md) |  | [optional] 

## Methods

### NewListRoleGrantsResponse

`func NewListRoleGrantsResponse() *ListRoleGrantsResponse`

NewListRoleGrantsResponse instantiates a new ListRoleGrantsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetGrants

`func (o *ListRoleGrantsResponse) GetGrants() []UserRoleGrants`

GetGrants returns the Grants field if non-nil, zero value otherwise.

### SetGrants

`func (o *ListRoleGrantsResponse) SetGrants(v []UserRoleGrants)`

SetGrants sets Grants field to given value.

### GetPagination

`func (o *ListRoleGrantsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListRoleGrantsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetUserGroupGrants

`func (o *ListRoleGrantsResponse) GetUserGroupGrants() []UserGroupRoleGrant`

GetUserGroupGrants returns the UserGroupGrants field if non-nil, zero value otherwise.

### SetUserGroupGrants

`func (o *ListRoleGrantsResponse) SetUserGroupGrants(v []UserGroupRoleGrant)`

SetUserGroupGrants sets UserGroupGrants field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


