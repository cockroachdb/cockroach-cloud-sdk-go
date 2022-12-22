# KeysetPaginationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOfTime** | Pointer to **time.Time** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 
**Page** | Pointer to **string** |  | [optional] 
**SortOrder** | Pointer to [**SortOrder**](SortOrder.md) |  | [optional] 

## Methods

### NewKeysetPaginationRequest

`func NewKeysetPaginationRequest() *KeysetPaginationRequest`

NewKeysetPaginationRequest instantiates a new KeysetPaginationRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAsOfTime

`func (o *KeysetPaginationRequest) GetAsOfTime() time.Time`

GetAsOfTime returns the AsOfTime field if non-nil, zero value otherwise.

### SetAsOfTime

`func (o *KeysetPaginationRequest) SetAsOfTime(v time.Time)`

SetAsOfTime sets AsOfTime field to given value.

### GetLimit

`func (o *KeysetPaginationRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### SetLimit

`func (o *KeysetPaginationRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### GetPage

`func (o *KeysetPaginationRequest) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### SetPage

`func (o *KeysetPaginationRequest) SetPage(v string)`

SetPage sets Page field to given value.

### GetSortOrder

`func (o *KeysetPaginationRequest) GetSortOrder() SortOrder`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### SetSortOrder

`func (o *KeysetPaginationRequest) SetSortOrder(v SortOrder)`

SetSortOrder sets SortOrder field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


