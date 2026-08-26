# CreateServiceAccountCredentialBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | [**ServiceAccountCredentialAlgorithmType**](ServiceAccountCredentialAlgorithmType.md) |  | 
**Description** | Pointer to **string** | An optional longer description. | [optional] 
**ExpiresAt** | **time.Time** | Required. When the credential expires; must be in the future. There is no maximum lifetime. | 
**Name** | **string** | A human-readable name for the credential. | 
**PublicKey** | **string** | The PEM-encoded public key used to verify signed JWT assertions. | 

## Methods

### NewCreateServiceAccountCredentialBody

`func NewCreateServiceAccountCredentialBody(algorithm ServiceAccountCredentialAlgorithmType, expiresAt time.Time, name string, publicKey string, ) *CreateServiceAccountCredentialBody`

NewCreateServiceAccountCredentialBody instantiates a new CreateServiceAccountCredentialBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateServiceAccountCredentialBodyWithDefaults

`func NewCreateServiceAccountCredentialBodyWithDefaults() *CreateServiceAccountCredentialBody`

NewCreateServiceAccountCredentialBodyWithDefaults instantiates a new CreateServiceAccountCredentialBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAlgorithm

`func (o *CreateServiceAccountCredentialBody) GetAlgorithm() ServiceAccountCredentialAlgorithmType`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### SetAlgorithm

`func (o *CreateServiceAccountCredentialBody) SetAlgorithm(v ServiceAccountCredentialAlgorithmType)`

SetAlgorithm sets Algorithm field to given value.

### GetDescription

`func (o *CreateServiceAccountCredentialBody) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *CreateServiceAccountCredentialBody) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetExpiresAt

`func (o *CreateServiceAccountCredentialBody) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### SetExpiresAt

`func (o *CreateServiceAccountCredentialBody) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetName

`func (o *CreateServiceAccountCredentialBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateServiceAccountCredentialBody) SetName(v string)`

SetName sets Name field to given value.

### GetPublicKey

`func (o *CreateServiceAccountCredentialBody) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### SetPublicKey

`func (o *CreateServiceAccountCredentialBody) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


