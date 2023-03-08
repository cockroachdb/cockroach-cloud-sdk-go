# ListMajorClusterVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**Versions** | [**[]ClusterMajorVersion**](ClusterMajorVersion.md) |  | 

## Methods

### NewListMajorClusterVersionsResponse

`func NewListMajorClusterVersionsResponse(versions []ClusterMajorVersion, ) *ListMajorClusterVersionsResponse`

NewListMajorClusterVersionsResponse instantiates a new ListMajorClusterVersionsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListMajorClusterVersionsResponseWithDefaults

`func NewListMajorClusterVersionsResponseWithDefaults() *ListMajorClusterVersionsResponse`

NewListMajorClusterVersionsResponseWithDefaults instantiates a new ListMajorClusterVersionsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListMajorClusterVersionsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListMajorClusterVersionsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetVersions

`func (o *ListMajorClusterVersionsResponse) GetVersions() []ClusterMajorVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### SetVersions

`func (o *ListMajorClusterVersionsResponse) SetVersions(v []ClusterMajorVersion)`

SetVersions sets Versions field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


