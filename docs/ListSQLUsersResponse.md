# ListSQLUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]SQLUser**](SQLUser.md) |  | 
**Pagination** | Pointer to [**PaginationResponse**](PaginationResponse.md) |  | [optional] 

## Methods

### NewListSQLUsersResponse

`func NewListSQLUsersResponse(users []SQLUser, ) *ListSQLUsersResponse`

NewListSQLUsersResponse instantiates a new ListSQLUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSQLUsersResponseWithDefaults

`func NewListSQLUsersResponseWithDefaults() *ListSQLUsersResponse`

NewListSQLUsersResponseWithDefaults instantiates a new ListSQLUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ListSQLUsersResponse) GetUsers() []SQLUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListSQLUsersResponse) GetUsersOk() (*[]SQLUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListSQLUsersResponse) SetUsers(v []SQLUser)`

SetUsers sets Users field to given value.


### GetPagination

`func (o *ListSQLUsersResponse) GetPagination() PaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSQLUsersResponse) GetPaginationOk() (*PaginationResponse, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSQLUsersResponse) SetPagination(v PaginationResponse)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSQLUsersResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


