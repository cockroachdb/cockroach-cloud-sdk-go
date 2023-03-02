# LogExportGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channels** | **[]string** | channels is a list of CRDB log channels to include in this group. | 
**LogName** | **string** |  | 
**MinLevel** | Pointer to [**LogLevel**](LogLevel.md) |  | [optional] 
**Redact** | Pointer to **bool** | redact is a boolean that governs whether this log group should aggregate redacted logs. Redaction settings will inherit from the cluster log export defaults if unset. | [optional] 

## Methods

### NewLogExportGroup

`func NewLogExportGroup(channels []string, logName string, ) *LogExportGroup`

NewLogExportGroup instantiates a new LogExportGroup object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLogExportGroupWithDefaults

`func NewLogExportGroupWithDefaults() *LogExportGroup`

NewLogExportGroupWithDefaults instantiates a new LogExportGroup object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetChannels

`func (o *LogExportGroup) GetChannels() []string`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### SetChannels

`func (o *LogExportGroup) SetChannels(v []string)`

SetChannels sets Channels field to given value.

### GetLogName

`func (o *LogExportGroup) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### SetLogName

`func (o *LogExportGroup) SetLogName(v string)`

SetLogName sets LogName field to given value.

### GetMinLevel

`func (o *LogExportGroup) GetMinLevel() LogLevel`

GetMinLevel returns the MinLevel field if non-nil, zero value otherwise.

### SetMinLevel

`func (o *LogExportGroup) SetMinLevel(v LogLevel)`

SetMinLevel sets MinLevel field to given value.

### GetRedact

`func (o *LogExportGroup) GetRedact() bool`

GetRedact returns the Redact field if non-nil, zero value otherwise.

### SetRedact

`func (o *LogExportGroup) SetRedact(v bool)`

SetRedact sets Redact field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


