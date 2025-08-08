# ListRestoresResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**Restores** | [**[]Restore**](Restore.md) |  | 

## Methods

### NewListRestoresResponse

`func NewListRestoresResponse(restores []Restore, ) *ListRestoresResponse`

NewListRestoresResponse instantiates a new ListRestoresResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListRestoresResponseWithDefaults

`func NewListRestoresResponseWithDefaults() *ListRestoresResponse`

NewListRestoresResponseWithDefaults instantiates a new ListRestoresResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListRestoresResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListRestoresResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetRestores

`func (o *ListRestoresResponse) GetRestores() []Restore`

GetRestores returns the Restores field if non-nil, zero value otherwise.

### SetRestores

`func (o *ListRestoresResponse) SetRestores(v []Restore)`

SetRestores sets Restores field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


