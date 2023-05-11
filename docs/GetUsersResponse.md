# GetUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Resources** | Pointer to [**[]ScimUser**](ScimUser.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetUsersResponse

`func NewGetUsersResponse() *GetUsersResponse`

NewGetUsersResponse instantiates a new GetUsersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetResources

`func (o *GetUsersResponse) GetResources() []ScimUser`

GetResources returns the Resources field if non-nil, zero value otherwise.

### SetResources

`func (o *GetUsersResponse) SetResources(v []ScimUser)`

SetResources sets Resources field to given value.

### GetSchemas

`func (o *GetUsersResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *GetUsersResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetTotalResults

`func (o *GetUsersResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### SetTotalResults

`func (o *GetUsersResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


