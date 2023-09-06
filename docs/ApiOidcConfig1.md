# ApiOidcConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Claim** | Pointer to **string** |  | [optional] 
**IdentityMap** | Pointer to [**[]ApiOidcIdentityMapEntry**](ApiOidcIdentityMapEntry.md) |  | [optional] 
**Issuer** | **string** |  | 
**Jwks** | **string** |  | 

## Methods

### NewApiOidcConfig1

`func NewApiOidcConfig1(audience string, issuer string, jwks string, ) *ApiOidcConfig1`

NewApiOidcConfig1 instantiates a new ApiOidcConfig1 object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApiOidcConfig1WithDefaults

`func NewApiOidcConfig1WithDefaults() *ApiOidcConfig1`

NewApiOidcConfig1WithDefaults instantiates a new ApiOidcConfig1 object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAudience

`func (o *ApiOidcConfig1) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *ApiOidcConfig1) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *ApiOidcConfig1) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *ApiOidcConfig1) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetIdentityMap

`func (o *ApiOidcConfig1) GetIdentityMap() []ApiOidcIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *ApiOidcConfig1) SetIdentityMap(v []ApiOidcIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuer

`func (o *ApiOidcConfig1) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### SetIssuer

`func (o *ApiOidcConfig1) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### GetJwks

`func (o *ApiOidcConfig1) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *ApiOidcConfig1) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


