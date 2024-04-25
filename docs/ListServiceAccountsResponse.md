# ListServiceAccountsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**ServiceAccounts** | [**[]ServiceAccount**](ServiceAccount.md) |  | 

## Methods

### NewListServiceAccountsResponse

`func NewListServiceAccountsResponse(serviceAccounts []ServiceAccount, ) *ListServiceAccountsResponse`

NewListServiceAccountsResponse instantiates a new ListServiceAccountsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListServiceAccountsResponseWithDefaults

`func NewListServiceAccountsResponseWithDefaults() *ListServiceAccountsResponse`

NewListServiceAccountsResponseWithDefaults instantiates a new ListServiceAccountsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListServiceAccountsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListServiceAccountsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetServiceAccounts

`func (o *ListServiceAccountsResponse) GetServiceAccounts() []ServiceAccount`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### SetServiceAccounts

`func (o *ListServiceAccountsResponse) SetServiceAccounts(v []ServiceAccount)`

SetServiceAccounts sets ServiceAccounts field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


