# ApiOidcConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Claim** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**IdentityMap** | Pointer to [**[]ApiOidcIdentityMapEntry**](ApiOidcIdentityMapEntry.md) |  | [optional] 
**Issuer** | **string** |  | 
**Jwks** | **string** |  | 

## Methods

### NewApiOidcConfig

`func NewApiOidcConfig(audience string, id string, issuer string, jwks string, ) *ApiOidcConfig`

NewApiOidcConfig instantiates a new ApiOidcConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApiOidcConfigWithDefaults

`func NewApiOidcConfigWithDefaults() *ApiOidcConfig`

NewApiOidcConfigWithDefaults instantiates a new ApiOidcConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAudience

`func (o *ApiOidcConfig) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *ApiOidcConfig) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *ApiOidcConfig) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *ApiOidcConfig) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetId

`func (o *ApiOidcConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ApiOidcConfig) SetId(v string)`

SetId sets Id field to given value.

### GetIdentityMap

`func (o *ApiOidcConfig) GetIdentityMap() []ApiOidcIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *ApiOidcConfig) SetIdentityMap(v []ApiOidcIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuer

`func (o *ApiOidcConfig) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### SetIssuer

`func (o *ApiOidcConfig) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### GetJwks

`func (o *ApiOidcConfig) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *ApiOidcConfig) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


