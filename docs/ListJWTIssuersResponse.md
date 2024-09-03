# ListJWTIssuersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtIssuers** | Pointer to [**[]JWTIssuer**](JWTIssuer.md) |  | [optional] 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListJWTIssuersResponse

`func NewListJWTIssuersResponse() *ListJWTIssuersResponse`

NewListJWTIssuersResponse instantiates a new ListJWTIssuersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetJwtIssuers

`func (o *ListJWTIssuersResponse) GetJwtIssuers() []JWTIssuer`

GetJwtIssuers returns the JwtIssuers field if non-nil, zero value otherwise.

### SetJwtIssuers

`func (o *ListJWTIssuersResponse) SetJwtIssuers(v []JWTIssuer)`

SetJwtIssuers sets JwtIssuers field to given value.

### GetPagination

`func (o *ListJWTIssuersResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListJWTIssuersResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


