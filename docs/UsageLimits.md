# UsageLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestUnitLimit** | **int64** | request_unit_limit is the maximum number of request units that the cluster can consume during the month. If this limit is exceeded, then the cluster is disabled until the limit is increased, or until the beginning of the next month when more free request units are granted. It is an error for this to be zero. | 
**StorageMibLimit** | **int64** | storage_mib_limit is the maximum number of Mebibytes of storage that the cluster can have at any time during the month. If this limit is exceeded, then the cluster is throttled; only one SQL connection is allowed at a time, with the expectation that it is used to delete data to reduce storage usage. It is an error for this to be zero. | 

## Methods

### NewUsageLimits

`func NewUsageLimits(requestUnitLimit int64, storageMibLimit int64, ) *UsageLimits`

NewUsageLimits instantiates a new UsageLimits object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUsageLimitsWithDefaults

`func NewUsageLimitsWithDefaults() *UsageLimits`

NewUsageLimitsWithDefaults instantiates a new UsageLimits object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRequestUnitLimit

`func (o *UsageLimits) GetRequestUnitLimit() int64`

GetRequestUnitLimit returns the RequestUnitLimit field if non-nil, zero value otherwise.

### SetRequestUnitLimit

`func (o *UsageLimits) SetRequestUnitLimit(v int64)`

SetRequestUnitLimit sets RequestUnitLimit field to given value.

### GetStorageMibLimit

`func (o *UsageLimits) GetStorageMibLimit() int64`

GetStorageMibLimit returns the StorageMibLimit field if non-nil, zero value otherwise.

### SetStorageMibLimit

`func (o *UsageLimits) SetStorageMibLimit(v int64)`

SetStorageMibLimit sets StorageMibLimit field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


