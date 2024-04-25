# ListApiKeysResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeys** | [**[]ApiKey**](ApiKey.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListApiKeysResponse

`func NewListApiKeysResponse(apiKeys []ApiKey, ) *ListApiKeysResponse`

NewListApiKeysResponse instantiates a new ListApiKeysResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListApiKeysResponseWithDefaults

`func NewListApiKeysResponseWithDefaults() *ListApiKeysResponse`

NewListApiKeysResponseWithDefaults instantiates a new ListApiKeysResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiKeys

`func (o *ListApiKeysResponse) GetApiKeys() []ApiKey`

GetApiKeys returns the ApiKeys field if non-nil, zero value otherwise.

### SetApiKeys

`func (o *ListApiKeysResponse) SetApiKeys(v []ApiKey)`

SetApiKeys sets ApiKeys field to given value.

### GetPagination

`func (o *ListApiKeysResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListApiKeysResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


