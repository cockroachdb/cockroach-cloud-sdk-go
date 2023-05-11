# GetServiceProviderConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationSchemes** | Pointer to [**[]ScimAuthenticationScheme**](ScimAuthenticationScheme.md) |  | [optional] 
**Bulk** | Pointer to [**ScimBulkSupport**](ScimBulkSupport.md) |  | [optional] 
**ChangePassword** | Pointer to [**ScimChangePasswordSupport**](ScimChangePasswordSupport.md) |  | [optional] 
**Etag** | Pointer to [**ScimEtagSupport**](ScimEtagSupport.md) |  | [optional] 
**Filter** | Pointer to [**ScimFilterSupport**](ScimFilterSupport.md) |  | [optional] 
**Meta** | Pointer to [**ScimMetadata**](ScimMetadata.md) |  | [optional] 
**Schemas** | Pointer to **[]string** |  | [optional] 
**Sort** | Pointer to [**ScimSortSupport**](ScimSortSupport.md) |  | [optional] 

## Methods

### NewGetServiceProviderConfigResponse

`func NewGetServiceProviderConfigResponse() *GetServiceProviderConfigResponse`

NewGetServiceProviderConfigResponse instantiates a new GetServiceProviderConfigResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAuthenticationSchemes

`func (o *GetServiceProviderConfigResponse) GetAuthenticationSchemes() []ScimAuthenticationScheme`

GetAuthenticationSchemes returns the AuthenticationSchemes field if non-nil, zero value otherwise.

### SetAuthenticationSchemes

`func (o *GetServiceProviderConfigResponse) SetAuthenticationSchemes(v []ScimAuthenticationScheme)`

SetAuthenticationSchemes sets AuthenticationSchemes field to given value.

### GetBulk

`func (o *GetServiceProviderConfigResponse) GetBulk() ScimBulkSupport`

GetBulk returns the Bulk field if non-nil, zero value otherwise.

### SetBulk

`func (o *GetServiceProviderConfigResponse) SetBulk(v ScimBulkSupport)`

SetBulk sets Bulk field to given value.

### GetChangePassword

`func (o *GetServiceProviderConfigResponse) GetChangePassword() ScimChangePasswordSupport`

GetChangePassword returns the ChangePassword field if non-nil, zero value otherwise.

### SetChangePassword

`func (o *GetServiceProviderConfigResponse) SetChangePassword(v ScimChangePasswordSupport)`

SetChangePassword sets ChangePassword field to given value.

### GetEtag

`func (o *GetServiceProviderConfigResponse) GetEtag() ScimEtagSupport`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### SetEtag

`func (o *GetServiceProviderConfigResponse) SetEtag(v ScimEtagSupport)`

SetEtag sets Etag field to given value.

### GetFilter

`func (o *GetServiceProviderConfigResponse) GetFilter() ScimFilterSupport`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### SetFilter

`func (o *GetServiceProviderConfigResponse) SetFilter(v ScimFilterSupport)`

SetFilter sets Filter field to given value.

### GetMeta

`func (o *GetServiceProviderConfigResponse) GetMeta() ScimMetadata`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### SetMeta

`func (o *GetServiceProviderConfigResponse) SetMeta(v ScimMetadata)`

SetMeta sets Meta field to given value.

### GetSchemas

`func (o *GetServiceProviderConfigResponse) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### SetSchemas

`func (o *GetServiceProviderConfigResponse) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.

### GetSort

`func (o *GetServiceProviderConfigResponse) GetSort() ScimSortSupport`

GetSort returns the Sort field if non-nil, zero value otherwise.

### SetSort

`func (o *GetServiceProviderConfigResponse) SetSort(v ScimSortSupport)`

SetSort sets Sort field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


