# KeysetPaginationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to **string** |  | [optional] 
**Last** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to SORTORDER_ASC]

## Methods

### NewKeysetPaginationResponse

`func NewKeysetPaginationResponse() *KeysetPaginationResponse`

NewKeysetPaginationResponse instantiates a new KeysetPaginationResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetNext

`func (o *KeysetPaginationResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### SetNext

`func (o *KeysetPaginationResponse) SetNext(v string)`

SetNext sets Next field to given value.

### GetLast

`func (o *KeysetPaginationResponse) GetLast() string`

GetLast returns the Last field if non-nil, zero value otherwise.

### SetLast

`func (o *KeysetPaginationResponse) SetLast(v string)`

SetLast sets Last field to given value.

### GetLimit

`func (o *KeysetPaginationResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### SetLimit

`func (o *KeysetPaginationResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### GetTime

`func (o *KeysetPaginationResponse) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### SetTime

`func (o *KeysetPaginationResponse) SetTime(v time.Time)`

SetTime sets Time field to given value.

### GetOrder

`func (o *KeysetPaginationResponse) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### SetOrder

`func (o *KeysetPaginationResponse) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


