# AddJWTIssuerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Claim** | Pointer to **string** |  | [optional] 
**IdentityMap** | Pointer to [**[]JWTIssuerIdentityMapEntry**](JWTIssuerIdentityMapEntry.md) | A list of mappings to map the external token identity into CockroachDB Cloud. | [optional] 
**IssuerUrl** | **string** |  | 
**Jwks** | Pointer to **string** |  | [optional] 

## Methods

### NewAddJWTIssuerRequest

`func NewAddJWTIssuerRequest(audience string, issuerUrl string, ) *AddJWTIssuerRequest`

NewAddJWTIssuerRequest instantiates a new AddJWTIssuerRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAddJWTIssuerRequestWithDefaults

`func NewAddJWTIssuerRequestWithDefaults() *AddJWTIssuerRequest`

NewAddJWTIssuerRequestWithDefaults instantiates a new AddJWTIssuerRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAudience

`func (o *AddJWTIssuerRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *AddJWTIssuerRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *AddJWTIssuerRequest) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *AddJWTIssuerRequest) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetIdentityMap

`func (o *AddJWTIssuerRequest) GetIdentityMap() []JWTIssuerIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *AddJWTIssuerRequest) SetIdentityMap(v []JWTIssuerIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuerUrl

`func (o *AddJWTIssuerRequest) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### SetIssuerUrl

`func (o *AddJWTIssuerRequest) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### GetJwks

`func (o *AddJWTIssuerRequest) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *AddJWTIssuerRequest) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


