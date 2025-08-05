# CockroachCloudCreateRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | Pointer to **string** | The ID of the backup from which data will be restored. If backup_id is not set, the restore will use the latest available backup from source_cluster_id. | [optional] 
**Objects** | Pointer to [**[]RestoreItem**](RestoreItem.md) | The list of objects to restore. Required when type is DATABASE or TABLE. | [optional] 
**RestoreOpts** | Pointer to [**RestoreOpts**](RestoreOpts.md) |  | [optional] 
**SourceClusterId** | Pointer to **string** | The ID of the cluster containing the backup to be restored. | [optional] 
**Type** | [**RestoreTypeType**](RestoreTypeType.md) |  | 

## Methods

### NewCockroachCloudCreateRestoreRequest

`func NewCockroachCloudCreateRestoreRequest(type_ RestoreTypeType, ) *CockroachCloudCreateRestoreRequest`

NewCockroachCloudCreateRestoreRequest instantiates a new CockroachCloudCreateRestoreRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCockroachCloudCreateRestoreRequestWithDefaults

`func NewCockroachCloudCreateRestoreRequestWithDefaults() *CockroachCloudCreateRestoreRequest`

NewCockroachCloudCreateRestoreRequestWithDefaults instantiates a new CockroachCloudCreateRestoreRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBackupId

`func (o *CockroachCloudCreateRestoreRequest) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### SetBackupId

`func (o *CockroachCloudCreateRestoreRequest) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### GetObjects

`func (o *CockroachCloudCreateRestoreRequest) GetObjects() []RestoreItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### SetObjects

`func (o *CockroachCloudCreateRestoreRequest) SetObjects(v []RestoreItem)`

SetObjects sets Objects field to given value.

### GetRestoreOpts

`func (o *CockroachCloudCreateRestoreRequest) GetRestoreOpts() RestoreOpts`

GetRestoreOpts returns the RestoreOpts field if non-nil, zero value otherwise.

### SetRestoreOpts

`func (o *CockroachCloudCreateRestoreRequest) SetRestoreOpts(v RestoreOpts)`

SetRestoreOpts sets RestoreOpts field to given value.

### GetSourceClusterId

`func (o *CockroachCloudCreateRestoreRequest) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### SetSourceClusterId

`func (o *CockroachCloudCreateRestoreRequest) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### GetType

`func (o *CockroachCloudCreateRestoreRequest) GetType() RestoreTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *CockroachCloudCreateRestoreRequest) SetType(v RestoreTypeType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


