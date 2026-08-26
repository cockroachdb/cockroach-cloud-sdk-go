# ListServiceAccountCredentialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | [**[]ServiceAccountCredential**](ServiceAccountCredential.md) | The service account&#39;s credentials, ordered by created_at. | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListServiceAccountCredentialsResponse

`func NewListServiceAccountCredentialsResponse(credentials []ServiceAccountCredential, ) *ListServiceAccountCredentialsResponse`

NewListServiceAccountCredentialsResponse instantiates a new ListServiceAccountCredentialsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListServiceAccountCredentialsResponseWithDefaults

`func NewListServiceAccountCredentialsResponseWithDefaults() *ListServiceAccountCredentialsResponse`

NewListServiceAccountCredentialsResponseWithDefaults instantiates a new ListServiceAccountCredentialsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCredentials

`func (o *ListServiceAccountCredentialsResponse) GetCredentials() []ServiceAccountCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### SetCredentials

`func (o *ListServiceAccountCredentialsResponse) SetCredentials(v []ServiceAccountCredential)`

SetCredentials sets Credentials field to given value.

### GetPagination

`func (o *ListServiceAccountCredentialsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListServiceAccountCredentialsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


