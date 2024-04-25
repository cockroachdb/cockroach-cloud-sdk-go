# CreateApiKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | [**ApiKey**](ApiKey.md) |  | 
**Secret** | **string** | The full api key. This is the value that would be passed in the Authorization header. It is not stored by the backend and is therefore not recoverable if lost. | 

## Methods

### NewCreateApiKeyResponse

`func NewCreateApiKeyResponse(apiKey ApiKey, secret string, ) *CreateApiKeyResponse`

NewCreateApiKeyResponse instantiates a new CreateApiKeyResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateApiKeyResponseWithDefaults

`func NewCreateApiKeyResponseWithDefaults() *CreateApiKeyResponse`

NewCreateApiKeyResponseWithDefaults instantiates a new CreateApiKeyResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiKey

`func (o *CreateApiKeyResponse) GetApiKey() ApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### SetApiKey

`func (o *CreateApiKeyResponse) SetApiKey(v ApiKey)`

SetApiKey sets ApiKey field to given value.

### GetSecret

`func (o *CreateApiKeyResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### SetSecret

`func (o *CreateApiKeyResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


