# ListSQLUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**Users** | [**[]SQLUser**](SQLUser.md) |  | 

## Methods

### NewListSQLUsersResponse

`func NewListSQLUsersResponse(users []SQLUser, ) *ListSQLUsersResponse`

NewListSQLUsersResponse instantiates a new ListSQLUsersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListSQLUsersResponseWithDefaults

`func NewListSQLUsersResponseWithDefaults() *ListSQLUsersResponse`

NewListSQLUsersResponseWithDefaults instantiates a new ListSQLUsersResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListSQLUsersResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListSQLUsersResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetUsers

`func (o *ListSQLUsersResponse) GetUsers() []SQLUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### SetUsers

`func (o *ListSQLUsersResponse) SetUsers(v []SQLUser)`

SetUsers sets Users field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


