# BlackoutWindow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | The cluster ID that this blackout window applies to. | 
**EndTime** | **time.Time** | The end time of the blackout window (UTC). Can be up to 14 days after the start time. Must not be more than three months after the current time. | 
**Id** | **string** | The unique ID of the blackout window. | 
**StartTime** | **time.Time** | The start time of the blackout window (UTC). Must be scheduled at least 7 days in advance. | 

## Methods

### NewBlackoutWindow

`func NewBlackoutWindow(clusterId string, endTime time.Time, id string, startTime time.Time, ) *BlackoutWindow`

NewBlackoutWindow instantiates a new BlackoutWindow object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewBlackoutWindowWithDefaults

`func NewBlackoutWindowWithDefaults() *BlackoutWindow`

NewBlackoutWindowWithDefaults instantiates a new BlackoutWindow object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *BlackoutWindow) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *BlackoutWindow) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetEndTime

`func (o *BlackoutWindow) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### SetEndTime

`func (o *BlackoutWindow) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### GetId

`func (o *BlackoutWindow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *BlackoutWindow) SetId(v string)`

SetId sets Id field to given value.

### GetStartTime

`func (o *BlackoutWindow) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### SetStartTime

`func (o *BlackoutWindow) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


