# BackupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates whether backups are enabled. | 
**FrequencyMinutes** | **int32** | How frequently in minutes that backups are taken, which will determine the [RPO](https://www.cockroachlabs.com/docs/stable/disaster-recovery-overview#resilience-strategy) of the cluster. | 
**RetentionDays** | **int32** | The number of days backups are retained for. | 

## Methods

### NewBackupConfiguration

`func NewBackupConfiguration(enabled bool, frequencyMinutes int32, retentionDays int32, ) *BackupConfiguration`

NewBackupConfiguration instantiates a new BackupConfiguration object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewBackupConfigurationWithDefaults

`func NewBackupConfigurationWithDefaults() *BackupConfiguration`

NewBackupConfigurationWithDefaults instantiates a new BackupConfiguration object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEnabled

`func (o *BackupConfiguration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### SetEnabled

`func (o *BackupConfiguration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### GetFrequencyMinutes

`func (o *BackupConfiguration) GetFrequencyMinutes() int32`

GetFrequencyMinutes returns the FrequencyMinutes field if non-nil, zero value otherwise.

### SetFrequencyMinutes

`func (o *BackupConfiguration) SetFrequencyMinutes(v int32)`

SetFrequencyMinutes sets FrequencyMinutes field to given value.

### GetRetentionDays

`func (o *BackupConfiguration) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### SetRetentionDays

`func (o *BackupConfiguration) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


