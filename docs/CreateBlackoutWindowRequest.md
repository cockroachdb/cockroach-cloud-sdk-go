# CreateBlackoutWindowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **time.Time** | The end time of the blackout window (UTC). Can be up to 14 days after the start time. Must not be more than three months after the current time. | 
**StartTime** | **time.Time** | The start time of the blackout window (UTC). Must be scheduled at least 7 days in advance. | 

## Methods

### NewCreateBlackoutWindowRequest

`func NewCreateBlackoutWindowRequest(endTime time.Time, startTime time.Time, ) *CreateBlackoutWindowRequest`

NewCreateBlackoutWindowRequest instantiates a new CreateBlackoutWindowRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateBlackoutWindowRequestWithDefaults

`func NewCreateBlackoutWindowRequestWithDefaults() *CreateBlackoutWindowRequest`

NewCreateBlackoutWindowRequestWithDefaults instantiates a new CreateBlackoutWindowRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEndTime

`func (o *CreateBlackoutWindowRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### SetEndTime

`func (o *CreateBlackoutWindowRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### GetStartTime

`func (o *CreateBlackoutWindowRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### SetStartTime

`func (o *CreateBlackoutWindowRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


