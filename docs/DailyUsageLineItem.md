# DailyUsageLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description is the charge name, or the adjustment name on adjustment rows. | 
**Locality** | Pointer to **string** | locality is the region the charge was incurred in. Empty for adjustment rows and for charges not attributable to a specific region. | [optional] 
**Quantity** | **float64** | quantity is the number of units consumed. Zero for adjustment rows. | 
**QuantityUnit** | Pointer to [**QuantityUnitType**](QuantityUnitType.md) |  | [optional] 
**Total** | [**CurrencyAmount**](CurrencyAmount.md) |  | 
**UnitCategory** | Pointer to [**UnitCategoryType**](UnitCategoryType.md) |  | [optional] 
**UnitCost** | **float64** | unit_cost is the cost per unit. Zero for adjustment rows. | 

## Methods

### NewDailyUsageLineItem

`func NewDailyUsageLineItem(description string, quantity float64, total CurrencyAmount, unitCost float64, ) *DailyUsageLineItem`

NewDailyUsageLineItem instantiates a new DailyUsageLineItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDailyUsageLineItemWithDefaults

`func NewDailyUsageLineItemWithDefaults() *DailyUsageLineItem`

NewDailyUsageLineItemWithDefaults instantiates a new DailyUsageLineItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *DailyUsageLineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *DailyUsageLineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetLocality

`func (o *DailyUsageLineItem) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### SetLocality

`func (o *DailyUsageLineItem) SetLocality(v string)`

SetLocality sets Locality field to given value.

### GetQuantity

`func (o *DailyUsageLineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### SetQuantity

`func (o *DailyUsageLineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### GetQuantityUnit

`func (o *DailyUsageLineItem) GetQuantityUnit() QuantityUnitType`

GetQuantityUnit returns the QuantityUnit field if non-nil, zero value otherwise.

### SetQuantityUnit

`func (o *DailyUsageLineItem) SetQuantityUnit(v QuantityUnitType)`

SetQuantityUnit sets QuantityUnit field to given value.

### GetTotal

`func (o *DailyUsageLineItem) GetTotal() CurrencyAmount`

GetTotal returns the Total field if non-nil, zero value otherwise.

### SetTotal

`func (o *DailyUsageLineItem) SetTotal(v CurrencyAmount)`

SetTotal sets Total field to given value.

### GetUnitCategory

`func (o *DailyUsageLineItem) GetUnitCategory() UnitCategoryType`

GetUnitCategory returns the UnitCategory field if non-nil, zero value otherwise.

### SetUnitCategory

`func (o *DailyUsageLineItem) SetUnitCategory(v UnitCategoryType)`

SetUnitCategory sets UnitCategory field to given value.

### GetUnitCost

`func (o *DailyUsageLineItem) GetUnitCost() float64`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### SetUnitCost

`func (o *DailyUsageLineItem) SetUnitCost(v float64)`

SetUnitCost sets UnitCost field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


