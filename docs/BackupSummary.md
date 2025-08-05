# BackupSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOfTime** | **time.Time** | The point in time (in UTC) the backup restores to. | 
**Id** | **string** | The unique identifier associated with the backup. | 

## Methods

### NewBackupSummary

`func NewBackupSummary(asOfTime time.Time, id string, ) *BackupSummary`

NewBackupSummary instantiates a new BackupSummary object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewBackupSummaryWithDefaults

`func NewBackupSummaryWithDefaults() *BackupSummary`

NewBackupSummaryWithDefaults instantiates a new BackupSummary object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAsOfTime

`func (o *BackupSummary) GetAsOfTime() time.Time`

GetAsOfTime returns the AsOfTime field if non-nil, zero value otherwise.

### SetAsOfTime

`func (o *BackupSummary) SetAsOfTime(v time.Time)`

SetAsOfTime sets AsOfTime field to given value.

### GetId

`func (o *BackupSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *BackupSummary) SetId(v string)`

SetId sets Id field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


