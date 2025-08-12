# GetGroupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimGroup**](ScimGroup.md) |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Schemas** | **[]string** |  | 
**StartIndex** | Pointer to **int32** |  | [optional] 
**TotalResults** | **int32** |  | 

## Methods

### NewGetGroupsResponse

`func NewGetGroupsResponse(schemas []string, totalResults int32, ) *GetGroupsResponse`

NewGetGroupsResponse instantiates a new GetGroupsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewGetGroupsResponseWithDefaults

`func NewGetGroupsResponseWithDefaults() *GetGroupsResponse`

NewGetGroupsResponseWithDefaults instantiates a new GetGroupsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetResources

`func (o *GetGroupsResponse) GetResources() []ScimGroup`

GetResources returns the Resources field if non-nil, zero value otherwise.

### SetResources

`func (o *GetGroupsResponse) SetResources(v []ScimGroup)`

SetResources sets Resources field to given value.

### GetItemsPerPage

`func (o *GetGroupsResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### SetItemsPerPage

`func (o *GetGroupsResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### GetSchemas

`func (o *GetGroupsResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *GetGroupsResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetStartIndex

`func (o *GetGroupsResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### SetStartIndex

`func (o *GetGroupsResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### GetTotalResults

`func (o *GetGroupsResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### SetTotalResults

`func (o *GetGroupsResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


