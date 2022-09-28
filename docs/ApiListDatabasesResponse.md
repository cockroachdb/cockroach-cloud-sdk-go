# ApiListDatabasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Databases** | [**[]ApiDatabase**](ApiDatabase.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewApiListDatabasesResponse

`func NewApiListDatabasesResponse(databases []ApiDatabase, ) *ApiListDatabasesResponse`

NewApiListDatabasesResponse instantiates a new ApiListDatabasesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApiListDatabasesResponseWithDefaults

`func NewApiListDatabasesResponseWithDefaults() *ApiListDatabasesResponse`

NewApiListDatabasesResponseWithDefaults instantiates a new ApiListDatabasesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDatabases

`func (o *ApiListDatabasesResponse) GetDatabases() []ApiDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### SetDatabases

`func (o *ApiListDatabasesResponse) SetDatabases(v []ApiDatabase)`

SetDatabases sets Databases field to given value.

### GetPagination

`func (o *ApiListDatabasesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ApiListDatabasesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


