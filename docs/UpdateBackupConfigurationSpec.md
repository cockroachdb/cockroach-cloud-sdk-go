# UpdateBackupConfigurationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether backups are enabled. | [optional] 
**FrequencyMinutes** | Pointer to **int32** | How frequently in minutes that backups are taken, which will determine the [RPO](https://www.cockroachlabs.com/docs/stable/disaster-recovery-overview#resilience-strategy) of the cluster.  Valid values are [5, 10, 15, 30, 60, 240, 1440]. | [optional] 
**RetentionDays** | Pointer to **int32** | The number of days to retain backups for. Can only be set once, further changes require opening a support ticket. Valid values are [2, 7, 30, 90, 365]. | [optional] 

## Methods

### NewUpdateBackupConfigurationSpec

`func NewUpdateBackupConfigurationSpec() *UpdateBackupConfigurationSpec`

NewUpdateBackupConfigurationSpec instantiates a new UpdateBackupConfigurationSpec object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetEnabled

`func (o *UpdateBackupConfigurationSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### SetEnabled

`func (o *UpdateBackupConfigurationSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### GetFrequencyMinutes

`func (o *UpdateBackupConfigurationSpec) GetFrequencyMinutes() int32`

GetFrequencyMinutes returns the FrequencyMinutes field if non-nil, zero value otherwise.

### SetFrequencyMinutes

`func (o *UpdateBackupConfigurationSpec) SetFrequencyMinutes(v int32)`

SetFrequencyMinutes sets FrequencyMinutes field to given value.

### GetRetentionDays

`func (o *UpdateBackupConfigurationSpec) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### SetRetentionDays

`func (o *UpdateBackupConfigurationSpec) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


