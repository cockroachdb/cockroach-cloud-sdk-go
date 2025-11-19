# Restore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupEndTime** | **time.Time** | The timestamp at which the backup data was captured. | 
**BackupId** | **string** | The ID of the backup used for this restore job. | 
**ClientErrorCode** | Pointer to **int32** | Error code from the restore job, only populated if it has failed. | [optional] 
**ClientErrorMessage** | Pointer to **string** | Error message from the restore job, only populated if it has failed. | [optional] 
**CompletedAt** | Pointer to **time.Time** | The timestamp at which the restore job completed. | [optional] 
**CompletionPercent** | **float32** | The percentage of the restore job that has been completed. The value ranges from 0 to 1. | 
**CrdbJobId** | Pointer to **string** | The CockroachDB internal job ID for the restore job. | [optional] 
**CreatedAt** | **time.Time** | The time at which the restore job was initiated. | 
**DestinationClusterName** | **string** | The name of the cluster to which the restore is being applied. | 
**Id** | **string** | The unique identifier associated with the restore. | 
**Objects** | Pointer to [**[]RestoreItem**](RestoreItem.md) | The list of database objects (databases, tables) that were restored. | [optional] 
**RestoreOpts** | Pointer to [**RestoreOpts**](RestoreOpts.md) |  | [optional] 
**SourceClusterName** | **string** | The name of the cluster from which the backup was taken. | 
**Status** | [**RestoreStatusType**](RestoreStatusType.md) |  | 
**Type** | [**RestoreTypeType**](RestoreTypeType.md) |  | 

## Methods

### NewRestore

`func NewRestore(backupEndTime time.Time, backupId string, completionPercent float32, createdAt time.Time, destinationClusterName string, id string, sourceClusterName string, status RestoreStatusType, type_ RestoreTypeType, ) *Restore`

NewRestore instantiates a new Restore object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRestoreWithDefaults

`func NewRestoreWithDefaults() *Restore`

NewRestoreWithDefaults instantiates a new Restore object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBackupEndTime

`func (o *Restore) GetBackupEndTime() time.Time`

GetBackupEndTime returns the BackupEndTime field if non-nil, zero value otherwise.

### SetBackupEndTime

`func (o *Restore) SetBackupEndTime(v time.Time)`

SetBackupEndTime sets BackupEndTime field to given value.

### GetBackupId

`func (o *Restore) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### SetBackupId

`func (o *Restore) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### GetClientErrorCode

`func (o *Restore) GetClientErrorCode() int32`

GetClientErrorCode returns the ClientErrorCode field if non-nil, zero value otherwise.

### SetClientErrorCode

`func (o *Restore) SetClientErrorCode(v int32)`

SetClientErrorCode sets ClientErrorCode field to given value.

### GetClientErrorMessage

`func (o *Restore) GetClientErrorMessage() string`

GetClientErrorMessage returns the ClientErrorMessage field if non-nil, zero value otherwise.

### SetClientErrorMessage

`func (o *Restore) SetClientErrorMessage(v string)`

SetClientErrorMessage sets ClientErrorMessage field to given value.

### GetCompletedAt

`func (o *Restore) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### SetCompletedAt

`func (o *Restore) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### GetCompletionPercent

`func (o *Restore) GetCompletionPercent() float32`

GetCompletionPercent returns the CompletionPercent field if non-nil, zero value otherwise.

### SetCompletionPercent

`func (o *Restore) SetCompletionPercent(v float32)`

SetCompletionPercent sets CompletionPercent field to given value.

### GetCrdbJobId

`func (o *Restore) GetCrdbJobId() string`

GetCrdbJobId returns the CrdbJobId field if non-nil, zero value otherwise.

### SetCrdbJobId

`func (o *Restore) SetCrdbJobId(v string)`

SetCrdbJobId sets CrdbJobId field to given value.

### GetCreatedAt

`func (o *Restore) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *Restore) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetDestinationClusterName

`func (o *Restore) GetDestinationClusterName() string`

GetDestinationClusterName returns the DestinationClusterName field if non-nil, zero value otherwise.

### SetDestinationClusterName

`func (o *Restore) SetDestinationClusterName(v string)`

SetDestinationClusterName sets DestinationClusterName field to given value.

### GetId

`func (o *Restore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Restore) SetId(v string)`

SetId sets Id field to given value.

### GetObjects

`func (o *Restore) GetObjects() []RestoreItem`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### SetObjects

`func (o *Restore) SetObjects(v []RestoreItem)`

SetObjects sets Objects field to given value.

### GetRestoreOpts

`func (o *Restore) GetRestoreOpts() RestoreOpts`

GetRestoreOpts returns the RestoreOpts field if non-nil, zero value otherwise.

### SetRestoreOpts

`func (o *Restore) SetRestoreOpts(v RestoreOpts)`

SetRestoreOpts sets RestoreOpts field to given value.

### GetSourceClusterName

`func (o *Restore) GetSourceClusterName() string`

GetSourceClusterName returns the SourceClusterName field if non-nil, zero value otherwise.

### SetSourceClusterName

`func (o *Restore) SetSourceClusterName(v string)`

SetSourceClusterName sets SourceClusterName field to given value.

### GetStatus

`func (o *Restore) GetStatus() RestoreStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *Restore) SetStatus(v RestoreStatusType)`

SetStatus sets Status field to given value.

### GetType

`func (o *Restore) GetType() RestoreTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *Restore) SetType(v RestoreTypeType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


