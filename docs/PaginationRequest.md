# PaginationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** | Optional timestamp of the original request, for consistency. In a future release, the server will use it in an &#x60;AS OF SYSTEM TIME&#x60; clause. Currently non-functional. | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to ASC]

## Methods

### NewPaginationRequest

`func NewPaginationRequest() *PaginationRequest`

NewPaginationRequest instantiates a new PaginationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationRequestWithDefaults

`func NewPaginationRequestWithDefaults() *PaginationRequest`

NewPaginationRequestWithDefaults instantiates a new PaginationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *PaginationRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *PaginationRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *PaginationRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *PaginationRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTime

`func (o *PaginationRequest) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PaginationRequest) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PaginationRequest) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *PaginationRequest) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOrder

`func (o *PaginationRequest) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaginationRequest) GetOrderOk() (*SortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaginationRequest) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaginationRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


