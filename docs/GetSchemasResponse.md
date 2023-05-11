# GetSchemasResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimSchema**](ScimSchema.md) |  | [optional] 
**ItemsPerPage** | Pointer to **int32** |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**StartIndex** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetSchemasResponse

`func NewGetSchemasResponse() *GetSchemasResponse`

NewGetSchemasResponse instantiates a new GetSchemasResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetResources

`func (o *GetSchemasResponse) GetResources() []ScimSchema`

GetResources returns the Resources field if non-nil, zero value otherwise.

### SetResources

`func (o *GetSchemasResponse) SetResources(v []ScimSchema)`

SetResources sets Resources field to given value.

### GetItemsPerPage

`func (o *GetSchemasResponse) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### SetItemsPerPage

`func (o *GetSchemasResponse) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.

### GetSchemas

`func (o *GetSchemasResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *GetSchemasResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetStartIndex

`func (o *GetSchemasResponse) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### SetStartIndex

`func (o *GetSchemasResponse) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### GetTotalResults

`func (o *GetSchemasResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### SetTotalResults

`func (o *GetSchemasResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


