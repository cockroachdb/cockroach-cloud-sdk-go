# ListDatabasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databases** | [**[]Database**](Database.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListDatabasesResponse

`func NewListDatabasesResponse(databases []Database, ) *ListDatabasesResponse`

NewListDatabasesResponse instantiates a new ListDatabasesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListDatabasesResponseWithDefaults

`func NewListDatabasesResponseWithDefaults() *ListDatabasesResponse`

NewListDatabasesResponseWithDefaults instantiates a new ListDatabasesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDatabases

`func (o *ListDatabasesResponse) GetDatabases() []Database`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### SetDatabases

`func (o *ListDatabasesResponse) SetDatabases(v []Database)`

SetDatabases sets Databases field to given value.

### GetPagination

`func (o *ListDatabasesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListDatabasesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


