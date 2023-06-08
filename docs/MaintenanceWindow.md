# MaintenanceWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OffsetDuration** | **string** | OffsetDuration is the duration from the start of a week (Monday 00:00 UTC) that this maintenance window will start after.  Must be less than 1 week. | 
**WindowDuration** | **string** | WindowDuration is the duration of the maintenance window.  Must be at least 6 hours and less than 1 week. | 

## Methods

### NewMaintenanceWindow

`func NewMaintenanceWindow(offsetDuration string, windowDuration string, ) *MaintenanceWindow`

NewMaintenanceWindow instantiates a new MaintenanceWindow object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewMaintenanceWindowWithDefaults

`func NewMaintenanceWindowWithDefaults() *MaintenanceWindow`

NewMaintenanceWindowWithDefaults instantiates a new MaintenanceWindow object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetOffsetDuration

`func (o *MaintenanceWindow) GetOffsetDuration() string`

GetOffsetDuration returns the OffsetDuration field if non-nil, zero value otherwise.

### SetOffsetDuration

`func (o *MaintenanceWindow) SetOffsetDuration(v string)`

SetOffsetDuration sets OffsetDuration field to given value.

### GetWindowDuration

`func (o *MaintenanceWindow) GetWindowDuration() string`

GetWindowDuration returns the WindowDuration field if non-nil, zero value otherwise.

### SetWindowDuration

`func (o *MaintenanceWindow) SetWindowDuration(v string)`

SetWindowDuration sets WindowDuration field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


