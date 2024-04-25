# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | the creation time of the api key. | 
**Id** | **string** | the ID of the api key. | 
**Name** | **string** |  | 
**ServiceAccountId** | **string** | the ID of the service account the api key references. | 

## Methods

### NewApiKey

`func NewApiKey(createdAt time.Time, id string, name string, serviceAccountId string, ) *ApiKey`

NewApiKey instantiates a new ApiKey object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *ApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *ApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.

### GetServiceAccountId

`func (o *ApiKey) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### SetServiceAccountId

`func (o *ApiKey) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


