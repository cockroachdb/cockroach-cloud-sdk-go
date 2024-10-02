# JWTIssuerIdentityMapEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CcIdentity** | **string** | Specifies how to map the fetched token identity to an identity in CockroachDB Cloud. In case of a regular expression for token_identity, this must contain a \\1 placeholder for the matched content. Note that you will need to escape the backslash in the string as in the example usage (\\\\\\\\1). | 
**TokenIdentity** | **string** | Specifies how to fetch external identity from the token claim. A regular expression must start with a forward slash. The regular expression must be in RE2 compatible syntax. For further details, please see https://github.com/google/re2/wiki/Syntax. | 

## Methods

### NewJWTIssuerIdentityMapEntry

`func NewJWTIssuerIdentityMapEntry(ccIdentity string, tokenIdentity string, ) *JWTIssuerIdentityMapEntry`

NewJWTIssuerIdentityMapEntry instantiates a new JWTIssuerIdentityMapEntry object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewJWTIssuerIdentityMapEntryWithDefaults

`func NewJWTIssuerIdentityMapEntryWithDefaults() *JWTIssuerIdentityMapEntry`

NewJWTIssuerIdentityMapEntryWithDefaults instantiates a new JWTIssuerIdentityMapEntry object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCcIdentity

`func (o *JWTIssuerIdentityMapEntry) GetCcIdentity() string`

GetCcIdentity returns the CcIdentity field if non-nil, zero value otherwise.

### SetCcIdentity

`func (o *JWTIssuerIdentityMapEntry) SetCcIdentity(v string)`

SetCcIdentity sets CcIdentity field to given value.

### GetTokenIdentity

`func (o *JWTIssuerIdentityMapEntry) GetTokenIdentity() string`

GetTokenIdentity returns the TokenIdentity field if non-nil, zero value otherwise.

### SetTokenIdentity

`func (o *JWTIssuerIdentityMapEntry) SetTokenIdentity(v string)`

SetTokenIdentity sets TokenIdentity field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


