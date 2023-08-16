# ListApiOidcConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiOidcConfigs** | Pointer to [**[]ApiOidcConfig**](ApiOidcConfig.md) |  | [optional] 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListApiOidcConfigResponse

`func NewListApiOidcConfigResponse() *ListApiOidcConfigResponse`

NewListApiOidcConfigResponse instantiates a new ListApiOidcConfigResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetApiOidcConfigs

`func (o *ListApiOidcConfigResponse) GetApiOidcConfigs() []ApiOidcConfig`

GetApiOidcConfigs returns the ApiOidcConfigs field if non-nil, zero value otherwise.

### SetApiOidcConfigs

`func (o *ListApiOidcConfigResponse) SetApiOidcConfigs(v []ApiOidcConfig)`

SetApiOidcConfigs sets ApiOidcConfigs field to given value.

### GetPagination

`func (o *ListApiOidcConfigResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListApiOidcConfigResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


