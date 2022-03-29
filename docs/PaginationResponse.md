# PaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **int32** |  | [optional] 
**Last** | Pointer to **int32** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**TotalResults** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to ASC]

## Methods

### NewPaginationResponse

`func NewPaginationResponse() *PaginationResponse`

NewPaginationResponse instantiates a new PaginationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseWithDefaults

`func NewPaginationResponseWithDefaults() *PaginationResponse`

NewPaginationResponseWithDefaults instantiates a new PaginationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *PaginationResponse) GetNext() int32`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginationResponse) GetNextOk() (*int32, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginationResponse) SetNext(v int32)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginationResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetLast

`func (o *PaginationResponse) GetLast() int32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *PaginationResponse) GetLastOk() (*int32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *PaginationResponse) SetLast(v int32)`

SetLast sets Last field to given value.

### HasLast

`func (o *PaginationResponse) HasLast() bool`

HasLast returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotalResults

`func (o *PaginationResponse) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *PaginationResponse) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *PaginationResponse) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.

### HasTotalResults

`func (o *PaginationResponse) HasTotalResults() bool`

HasTotalResults returns a boolean if a field has been set.

### GetTime

`func (o *PaginationResponse) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PaginationResponse) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PaginationResponse) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *PaginationResponse) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetOrder

`func (o *PaginationResponse) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PaginationResponse) GetOrderOk() (*SortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PaginationResponse) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PaginationResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


