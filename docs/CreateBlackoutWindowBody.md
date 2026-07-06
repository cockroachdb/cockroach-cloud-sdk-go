# CreateBlackoutWindowBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** | The end time of the blackout window (UTC). Can be up to 14 days after the start time. Must not be more than three months after the current time. | 
**StartTime** | **time.Time** | The start time of the blackout window (UTC). Must be scheduled at least 7 days in advance. | 

## Methods

### NewCreateBlackoutWindowBody

`func NewCreateBlackoutWindowBody(endTime time.Time, startTime time.Time, ) *CreateBlackoutWindowBody`

NewCreateBlackoutWindowBody instantiates a new CreateBlackoutWindowBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateBlackoutWindowBodyWithDefaults

`func NewCreateBlackoutWindowBodyWithDefaults() *CreateBlackoutWindowBody`

NewCreateBlackoutWindowBodyWithDefaults instantiates a new CreateBlackoutWindowBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEndTime

`func (o *CreateBlackoutWindowBody) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### SetEndTime

`func (o *CreateBlackoutWindowBody) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### GetStartTime

`func (o *CreateBlackoutWindowBody) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### SetStartTime

`func (o *CreateBlackoutWindowBody) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


