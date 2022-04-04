# ListClusterNodesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**[]Node**](Node.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListClusterNodesResponse

`func NewListClusterNodesResponse(nodes []Node, ) *ListClusterNodesResponse`

NewListClusterNodesResponse instantiates a new ListClusterNodesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListClusterNodesResponseWithDefaults

`func NewListClusterNodesResponseWithDefaults() *ListClusterNodesResponse`

NewListClusterNodesResponseWithDefaults instantiates a new ListClusterNodesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetNodes

`func (o *ListClusterNodesResponse) GetNodes() []Node`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### SetNodes

`func (o *ListClusterNodesResponse) SetNodes(v []Node)`

SetNodes sets Nodes field to given value.

### GetPagination

`func (o *ListClusterNodesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListClusterNodesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


