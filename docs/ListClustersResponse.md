# ListClustersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | [**[]Cluster**](Cluster.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListClustersResponse

`func NewListClustersResponse(clusters []Cluster, ) *ListClustersResponse`

NewListClustersResponse instantiates a new ListClustersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListClustersResponseWithDefaults

`func NewListClustersResponseWithDefaults() *ListClustersResponse`

NewListClustersResponseWithDefaults instantiates a new ListClustersResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusters

`func (o *ListClustersResponse) GetClusters() []Cluster`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### SetClusters

`func (o *ListClustersResponse) SetClusters(v []Cluster)`

SetClusters sets Clusters field to given value.

### GetPagination

`func (o *ListClustersResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListClustersResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


