# GetResourceTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimResourceType**](ScimResourceType.md) |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**StartIndex** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetResourceTypesResponse

`func NewGetResourceTypesResponse() *GetResourceTypesResponse`

NewGetResourceTypesResponse instantiates a new GetResourceTypesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetResources

`func (o *GetResourceTypesResponse) GetResources() []ScimResourceType`

GetResources returns the Resources field if non-nil, zero value otherwise.

### SetResources

`func (o *GetResourceTypesResponse) SetResources(v []ScimResourceType)`

SetResources sets Resources field to given value.

### GetItemsPerPage

`func (o *GetResourceTypesResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### SetItemsPerPage

`func (o *GetResourceTypesResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### GetSchemas

`func (o *GetResourceTypesResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *GetResourceTypesResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetStartIndex

`func (o *GetResourceTypesResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### SetStartIndex

`func (o *GetResourceTypesResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### GetTotalResults

`func (o *GetResourceTypesResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### SetTotalResults

`func (o *GetResourceTypesResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


