# CMEKKeySpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CMEKKeyType**](CMEKKeyType.md) |  | [optional] [default to CMEKKEYTYPE_UNKNOWN_KEY_TYPE]
**Uri** | Pointer to **string** |  | [optional] 
**AuthPrincipal** | Pointer to **string** |  | [optional] 

## Methods

### NewCMEKKeySpecification

`func NewCMEKKeySpecification() *CMEKKeySpecification`

NewCMEKKeySpecification instantiates a new CMEKKeySpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetType

`func (o *CMEKKeySpecification) GetType() CMEKKeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *CMEKKeySpecification) SetType(v CMEKKeyType)`

SetType sets Type field to given value.

### GetUri

`func (o *CMEKKeySpecification) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### SetUri

`func (o *CMEKKeySpecification) SetUri(v string)`

SetUri sets Uri field to given value.

### GetAuthPrincipal

`func (o *CMEKKeySpecification) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### SetAuthPrincipal

`func (o *CMEKKeySpecification) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


