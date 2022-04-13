# ListAvailableRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | [**[]CloudProviderRegion**](CloudProviderRegion.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListAvailableRegionsResponse

`func NewListAvailableRegionsResponse(regions []CloudProviderRegion, ) *ListAvailableRegionsResponse`

NewListAvailableRegionsResponse instantiates a new ListAvailableRegionsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListAvailableRegionsResponseWithDefaults

`func NewListAvailableRegionsResponseWithDefaults() *ListAvailableRegionsResponse`

NewListAvailableRegionsResponseWithDefaults instantiates a new ListAvailableRegionsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRegions

`func (o *ListAvailableRegionsResponse) GetRegions() []CloudProviderRegion`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *ListAvailableRegionsResponse) SetRegions(v []CloudProviderRegion)`

SetRegions sets Regions field to given value.

### GetPagination

`func (o *ListAvailableRegionsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListAvailableRegionsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


