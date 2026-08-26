# ServiceAccountCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | [**ServiceAccountCredentialAlgorithmType**](ServiceAccountCredentialAlgorithmType.md) |  | 
**ClientId** | **string** | The public OAuth client identifier; also the iss and sub a signed assertion must carry. | 
**CreatedAt** | **time.Time** | When the credential was created. | 
**CredentialType** | [**ServiceAccountCredentialTypeType**](ServiceAccountCredentialTypeType.md) |  | 
**Description** | Pointer to **string** | An optional longer description. | [optional] 
**DisabledAt** | Pointer to **time.Time** | When the credential was disabled; unset if active. | [optional] 
**ExpiresAt** | **time.Time** | When the credential expires. Always set; a credential cannot be created without an expiry. | 
**Id** | **string** | The unique ID of the credential. | 
**Name** | **string** | A human-readable name for the credential. | 
**ServiceAccountId** | **string** | The ID of the service account the credential belongs to. | 

## Methods

### NewServiceAccountCredential

`func NewServiceAccountCredential(algorithm ServiceAccountCredentialAlgorithmType, clientId string, createdAt time.Time, credentialType ServiceAccountCredentialTypeType, expiresAt time.Time, id string, name string, serviceAccountId string, ) *ServiceAccountCredential`

NewServiceAccountCredential instantiates a new ServiceAccountCredential object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServiceAccountCredentialWithDefaults

`func NewServiceAccountCredentialWithDefaults() *ServiceAccountCredential`

NewServiceAccountCredentialWithDefaults instantiates a new ServiceAccountCredential object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAlgorithm

`func (o *ServiceAccountCredential) GetAlgorithm() ServiceAccountCredentialAlgorithmType`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### SetAlgorithm

`func (o *ServiceAccountCredential) SetAlgorithm(v ServiceAccountCredentialAlgorithmType)`

SetAlgorithm sets Algorithm field to given value.

### GetClientId

`func (o *ServiceAccountCredential) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### SetClientId

`func (o *ServiceAccountCredential) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### GetCreatedAt

`func (o *ServiceAccountCredential) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *ServiceAccountCredential) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetCredentialType

`func (o *ServiceAccountCredential) GetCredentialType() ServiceAccountCredentialTypeType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### SetCredentialType

`func (o *ServiceAccountCredential) SetCredentialType(v ServiceAccountCredentialTypeType)`

SetCredentialType sets CredentialType field to given value.

### GetDescription

`func (o *ServiceAccountCredential) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *ServiceAccountCredential) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetDisabledAt

`func (o *ServiceAccountCredential) GetDisabledAt() time.Time`

GetDisabledAt returns the DisabledAt field if non-nil, zero value otherwise.

### SetDisabledAt

`func (o *ServiceAccountCredential) SetDisabledAt(v time.Time)`

SetDisabledAt sets DisabledAt field to given value.

### GetExpiresAt

`func (o *ServiceAccountCredential) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### SetExpiresAt

`func (o *ServiceAccountCredential) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### GetId

`func (o *ServiceAccountCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ServiceAccountCredential) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *ServiceAccountCredential) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *ServiceAccountCredential) SetName(v string)`

SetName sets Name field to given value.

### GetServiceAccountId

`func (o *ServiceAccountCredential) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### SetServiceAccountId

`func (o *ServiceAccountCredential) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


