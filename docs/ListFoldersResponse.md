# ListFoldersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folders** | [**[]FolderResource**](FolderResource.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListFoldersResponse

`func NewListFoldersResponse(folders []FolderResource, ) *ListFoldersResponse`

NewListFoldersResponse instantiates a new ListFoldersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListFoldersResponseWithDefaults

`func NewListFoldersResponseWithDefaults() *ListFoldersResponse`

NewListFoldersResponseWithDefaults instantiates a new ListFoldersResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetFolders

`func (o *ListFoldersResponse) GetFolders() []FolderResource`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### SetFolders

`func (o *ListFoldersResponse) SetFolders(v []FolderResource)`

SetFolders sets Folders field to given value.

### GetPagination

`func (o *ListFoldersResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListFoldersResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


