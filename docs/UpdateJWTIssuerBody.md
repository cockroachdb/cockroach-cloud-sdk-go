# UpdateJWTIssuerBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Claim** | Pointer to **string** |  | [optional] 
**IdentityMap** | Pointer to [**[]JWTIssuerIdentityMapEntry**](JWTIssuerIdentityMapEntry.md) |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**Jwks** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateJWTIssuerBody

`func NewUpdateJWTIssuerBody() *UpdateJWTIssuerBody`

NewUpdateJWTIssuerBody instantiates a new UpdateJWTIssuerBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAudience

`func (o *UpdateJWTIssuerBody) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *UpdateJWTIssuerBody) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *UpdateJWTIssuerBody) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *UpdateJWTIssuerBody) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetIdentityMap

`func (o *UpdateJWTIssuerBody) GetIdentityMap() []JWTIssuerIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *UpdateJWTIssuerBody) SetIdentityMap(v []JWTIssuerIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuerUrl

`func (o *UpdateJWTIssuerBody) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### SetIssuerUrl

`func (o *UpdateJWTIssuerBody) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### GetJwks

`func (o *UpdateJWTIssuerBody) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *UpdateJWTIssuerBody) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


