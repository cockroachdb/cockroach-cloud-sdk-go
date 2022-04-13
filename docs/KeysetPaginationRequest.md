# KeysetPaginationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartKey** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to [**PageDirection**](PageDirection.md) |  | [optional] [default to PAGEDIRECTION_NEXT]
**Limit** | Pointer to **int32** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**Order** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] [default to SORTORDER_ASC]

## Methods

### NewKeysetPaginationRequest

`func NewKeysetPaginationRequest() *KeysetPaginationRequest`

NewKeysetPaginationRequest instantiates a new KeysetPaginationRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetStartKey

`func (o *KeysetPaginationRequest) GetStartKey() string`

GetStartKey returns the StartKey field if non-nil, zero value otherwise.

### SetStartKey

`func (o *KeysetPaginationRequest) SetStartKey(v string)`

SetStartKey sets StartKey field to given value.

### GetDirection

`func (o *KeysetPaginationRequest) GetDirection() PageDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### SetDirection

`func (o *KeysetPaginationRequest) SetDirection(v PageDirection)`

SetDirection sets Direction field to given value.

### GetLimit

`func (o *KeysetPaginationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### SetLimit

`func (o *KeysetPaginationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### GetTime

`func (o *KeysetPaginationRequest) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### SetTime

`func (o *KeysetPaginationRequest) SetTime(v time.Time)`

SetTime sets Time field to given value.

### GetOrder

`func (o *KeysetPaginationRequest) GetOrder() SortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### SetOrder

`func (o *KeysetPaginationRequest) SetOrder(v SortOrder)`

SetOrder sets Order field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


