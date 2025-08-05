# Restore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupId** | **string** | The ID of the backup used for this restore operation. | 
**CompletionPercent** | **float32** | The percentage of the restore operation that has been completed. The value ranges from 0 to 1. | 
**CreatedAt** | **time.Time** | The time at which the restore operation was initiated. | 
**Id** | **string** | The unique identifier associated with the restore. | 
**Status** | [**RestoreStatusType**](RestoreStatusType.md) |  | 
**Type** | [**RestoreTypeType**](RestoreTypeType.md) |  | 

## Methods

### NewRestore

`func NewRestore(backupId string, completionPercent float32, createdAt time.Time, id string, status RestoreStatusType, type_ RestoreTypeType, ) *Restore`

NewRestore instantiates a new Restore object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRestoreWithDefaults

`func NewRestoreWithDefaults() *Restore`

NewRestoreWithDefaults instantiates a new Restore object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBackupId

`func (o *Restore) GetBackupId() string`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### SetBackupId

`func (o *Restore) SetBackupId(v string)`

SetBackupId sets BackupId field to given value.

### GetCompletionPercent

`func (o *Restore) GetCompletionPercent() float32`

GetCompletionPercent returns the CompletionPercent field if non-nil, zero value otherwise.

### SetCompletionPercent

`func (o *Restore) SetCompletionPercent(v float32)`

SetCompletionPercent sets CompletionPercent field to given value.

### GetCreatedAt

`func (o *Restore) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *Restore) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetId

`func (o *Restore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Restore) SetId(v string)`

SetId sets Id field to given value.

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


