# CMEKKeySpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CMEKKeyType**](CMEKKeyType.md) |  | [optional] [default to UNKNOWN_KEY_TYPE]
**Uri** | Pointer to **string** |  | [optional] 
**AuthPrincipal** | Pointer to **string** |  | [optional] 

## Methods

### NewCMEKKeySpecification

`func NewCMEKKeySpecification() *CMEKKeySpecification`

NewCMEKKeySpecification instantiates a new CMEKKeySpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKKeySpecificationWithDefaults

`func NewCMEKKeySpecificationWithDefaults() *CMEKKeySpecification`

NewCMEKKeySpecificationWithDefaults instantiates a new CMEKKeySpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CMEKKeySpecification) GetType() CMEKKeyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CMEKKeySpecification) GetTypeOk() (*CMEKKeyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CMEKKeySpecification) SetType(v CMEKKeyType)`

SetType sets Type field to given value.

### HasType

`func (o *CMEKKeySpecification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUri

`func (o *CMEKKeySpecification) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CMEKKeySpecification) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CMEKKeySpecification) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CMEKKeySpecification) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetAuthPrincipal

`func (o *CMEKKeySpecification) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### GetAuthPrincipalOk

`func (o *CMEKKeySpecification) GetAuthPrincipalOk() (*string, bool)`

GetAuthPrincipalOk returns a tuple with the AuthPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthPrincipal

`func (o *CMEKKeySpecification) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.

### HasAuthPrincipal

`func (o *CMEKKeySpecification) HasAuthPrincipal() bool`

HasAuthPrincipal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


