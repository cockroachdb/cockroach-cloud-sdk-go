# MigrationAssistant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketGrantee** | Pointer to **string** | The cloud-native identity the customer must grant bucket access to; populated when the status is RUNNING. Format depends on the cluster&#39;s cloud (AWS IAM role ARN, GCP service account email, or Azure object ID). | [optional] 
**ClusterId** | **string** | The ID of the cluster this assistant is deployed for. | 
**CreatedAt** | **time.Time** | Timestamp when the assistant was created. | 
**DeployedRegion** | **string** | The cluster region the assistant is deployed in. | 
**Health** | [**MigrationAssistantHealthType**](MigrationAssistantHealthType.md) |  | 
**Id** | **string** | The unique ID of the migration assistant. | 
**LastError** | Pointer to **string** | The failure detail; populated when the status is FAILED. | [optional] 
**Password** | Pointer to **string** | The password used to sign in to the assistant; populated when the status is RUNNING. | [optional] 
**Status** | [**MigrationAssistantStatusType**](MigrationAssistantStatusType.md) |  | 
**UpdatedAt** | **time.Time** | Timestamp when the assistant was last updated. | 
**Url** | Pointer to **string** | The URL used to reach the assistant; populated when the status is RUNNING. | [optional] 
**Username** | Pointer to **string** | The username used to sign in to the assistant; populated when the status is RUNNING. | [optional] 

## Methods

### NewMigrationAssistant

`func NewMigrationAssistant(clusterId string, createdAt time.Time, deployedRegion string, health MigrationAssistantHealthType, id string, status MigrationAssistantStatusType, updatedAt time.Time, ) *MigrationAssistant`

NewMigrationAssistant instantiates a new MigrationAssistant object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMigrationAssistantWithDefaults

`func NewMigrationAssistantWithDefaults() *MigrationAssistant`

NewMigrationAssistantWithDefaults instantiates a new MigrationAssistant object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBucketGrantee

`func (o *MigrationAssistant) GetBucketGrantee() string`

GetBucketGrantee returns the BucketGrantee field if non-nil, zero value otherwise.

### SetBucketGrantee

`func (o *MigrationAssistant) SetBucketGrantee(v string)`

SetBucketGrantee sets BucketGrantee field to given value.

### GetClusterId

`func (o *MigrationAssistant) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *MigrationAssistant) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetCreatedAt

`func (o *MigrationAssistant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *MigrationAssistant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetDeployedRegion

`func (o *MigrationAssistant) GetDeployedRegion() string`

GetDeployedRegion returns the DeployedRegion field if non-nil, zero value otherwise.

### SetDeployedRegion

`func (o *MigrationAssistant) SetDeployedRegion(v string)`

SetDeployedRegion sets DeployedRegion field to given value.

### GetHealth

`func (o *MigrationAssistant) GetHealth() MigrationAssistantHealthType`

GetHealth returns the Health field if non-nil, zero value otherwise.

### SetHealth

`func (o *MigrationAssistant) SetHealth(v MigrationAssistantHealthType)`

SetHealth sets Health field to given value.

### GetId

`func (o *MigrationAssistant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *MigrationAssistant) SetId(v string)`

SetId sets Id field to given value.

### GetLastError

`func (o *MigrationAssistant) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### SetLastError

`func (o *MigrationAssistant) SetLastError(v string)`

SetLastError sets LastError field to given value.

### GetPassword

`func (o *MigrationAssistant) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### SetPassword

`func (o *MigrationAssistant) SetPassword(v string)`

SetPassword sets Password field to given value.

### GetStatus

`func (o *MigrationAssistant) GetStatus() MigrationAssistantStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *MigrationAssistant) SetStatus(v MigrationAssistantStatusType)`

SetStatus sets Status field to given value.

### GetUpdatedAt

`func (o *MigrationAssistant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *MigrationAssistant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### GetUrl

`func (o *MigrationAssistant) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### SetUrl

`func (o *MigrationAssistant) SetUrl(v string)`

SetUrl sets Url field to given value.

### GetUsername

`func (o *MigrationAssistant) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### SetUsername

`func (o *MigrationAssistant) SetUsername(v string)`

SetUsername sets Username field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


