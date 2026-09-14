# ListVirtualClusterWorkspacesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**Workspaces** | [**[]VirtualClusterWorkspace**](VirtualClusterWorkspace.md) | The workspaces in the organization. | 

## Methods

### NewListVirtualClusterWorkspacesResponse

`func NewListVirtualClusterWorkspacesResponse(workspaces []VirtualClusterWorkspace, ) *ListVirtualClusterWorkspacesResponse`

NewListVirtualClusterWorkspacesResponse instantiates a new ListVirtualClusterWorkspacesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListVirtualClusterWorkspacesResponseWithDefaults

`func NewListVirtualClusterWorkspacesResponseWithDefaults() *ListVirtualClusterWorkspacesResponse`

NewListVirtualClusterWorkspacesResponseWithDefaults instantiates a new ListVirtualClusterWorkspacesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListVirtualClusterWorkspacesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListVirtualClusterWorkspacesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetWorkspaces

`func (o *ListVirtualClusterWorkspacesResponse) GetWorkspaces() []VirtualClusterWorkspace`

GetWorkspaces returns the Workspaces field if non-nil, zero value otherwise.

### SetWorkspaces

`func (o *ListVirtualClusterWorkspacesResponse) SetWorkspaces(v []VirtualClusterWorkspace)`

SetWorkspaces sets Workspaces field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


