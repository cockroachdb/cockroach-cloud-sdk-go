# UpdateJWTIssuerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** |  | [optional] 
**Claim** | Pointer to **string** |  | [optional] 
**IdentityMap** | Pointer to [**[]JWTIssuerIdentityMapEntry**](JWTIssuerIdentityMapEntry.md) |  | [optional] 
**IssuerUrl** | Pointer to **string** |  | [optional] 
**Jwks** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateJWTIssuerRequest

`func NewUpdateJWTIssuerRequest() *UpdateJWTIssuerRequest`

NewUpdateJWTIssuerRequest instantiates a new UpdateJWTIssuerRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAudience

`func (o *UpdateJWTIssuerRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### SetAudience

`func (o *UpdateJWTIssuerRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### GetClaim

`func (o *UpdateJWTIssuerRequest) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### SetClaim

`func (o *UpdateJWTIssuerRequest) SetClaim(v string)`

SetClaim sets Claim field to given value.

### GetIdentityMap

`func (o *UpdateJWTIssuerRequest) GetIdentityMap() []JWTIssuerIdentityMapEntry`

GetIdentityMap returns the IdentityMap field if non-nil, zero value otherwise.

### SetIdentityMap

`func (o *UpdateJWTIssuerRequest) SetIdentityMap(v []JWTIssuerIdentityMapEntry)`

SetIdentityMap sets IdentityMap field to given value.

### GetIssuerUrl

`func (o *UpdateJWTIssuerRequest) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### SetIssuerUrl

`func (o *UpdateJWTIssuerRequest) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### GetJwks

`func (o *UpdateJWTIssuerRequest) GetJwks() string`

GetJwks returns the Jwks field if non-nil, zero value otherwise.

### SetJwks

`func (o *UpdateJWTIssuerRequest) SetJwks(v string)`

SetJwks sets Jwks field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


