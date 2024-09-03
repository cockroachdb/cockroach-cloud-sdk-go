# JWTIssuer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Claim** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IdentityMap** | Pointer to [**[]JWTIssuerIdentityMapEntry**](JWTIssuerIdentityMapEntry.md) |  | [optional] 
**IssuerUrl** | **string** |  | 
**Jwks** | Pointer to **string** |  | [optional] 

## Methods

### NewJWTIssuer

`func NewJWTIssuer(audience string, id string, issuerUrl string, ) *JWTIssuer`

NewJWTIssuer instantiates a new JWTIssuer object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewJWTIssuerWithDefaults

`func NewJWTIssuerWithDefaults() *JWTIssuer`

NewJWTIssuerWithDefaults instantiates a new JWTIssuer object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAudience

`func (o *JWTIssuer) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *JWTIssuer) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *JWTIssuer) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *JWTIssuer) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetId

`func (o *JWTIssuer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *JWTIssuer) SetId(v string)`

SetId sets Id field to given value.

### GetIdentityMap

`func (o *JWTIssuer) GetIdentityMap() []JWTIssuerIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *JWTIssuer) SetIdentityMap(v []JWTIssuerIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuerUrl

`func (o *JWTIssuer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### SetIssuerUrl

`func (o *JWTIssuer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### GetJwks

`func (o *JWTIssuer) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *JWTIssuer) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


