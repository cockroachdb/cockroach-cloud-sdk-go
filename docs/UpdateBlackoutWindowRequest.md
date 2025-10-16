# UpdateBlackoutWindowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **time.Time** | Optional new UTC end time for the blackout window. Can be up to 14 days after the start time. Must not be more than three months after the current time. | [optional] 
**StartTime** | Pointer to **time.Time** | Optional new UTC start time for the blackout window. Must be scheduled at least 7 days in advance. | [optional] 

## Methods

### NewUpdateBlackoutWindowRequest

`func NewUpdateBlackoutWindowRequest() *UpdateBlackoutWindowRequest`

NewUpdateBlackoutWindowRequest instantiates a new UpdateBlackoutWindowRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetEndTime

`func (o *UpdateBlackoutWindowRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### SetEndTime

`func (o *UpdateBlackoutWindowRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### GetStartTime

`func (o *UpdateBlackoutWindowRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### SetStartTime

`func (o *UpdateBlackoutWindowRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


