# LineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description contains the details of the line item (i.e t3 micro). | 
**Quantity** | **float64** | quantity is the number of the specific line items used. | 
**QuantityUnit** | [**QuantityUnitType**](QuantityUnitType.md) |  | 
**Total** | [**CurrencyAmount**](CurrencyAmount.md) |  | 
**UnitCost** | **float64** | unit_cost is the cost per unit of line item. | 

## Methods

### NewLineItem

`func NewLineItem(description string, quantity float64, quantityUnit QuantityUnitType, total CurrencyAmount, unitCost float64, ) *LineItem`

NewLineItem instantiates a new LineItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewLineItemWithDefaults

`func NewLineItemWithDefaults() *LineItem`

NewLineItemWithDefaults instantiates a new LineItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *LineItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *LineItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetQuantity

`func (o *LineItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### SetQuantity

`func (o *LineItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### GetQuantityUnit

`func (o *LineItem) GetQuantityUnit() QuantityUnitType`

GetQuantityUnit returns the QuantityUnit field if non-nil, zero value otherwise.

### SetQuantityUnit

`func (o *LineItem) SetQuantityUnit(v QuantityUnitType)`

SetQuantityUnit sets QuantityUnit field to given value.

### GetTotal

`func (o *LineItem) GetTotal() CurrencyAmount`

GetTotal returns the Total field if non-nil, zero value otherwise.

### SetTotal

`func (o *LineItem) SetTotal(v CurrencyAmount)`

SetTotal sets Total field to given value.

### GetUnitCost

`func (o *LineItem) GetUnitCost() float64`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### SetUnitCost

`func (o *LineItem) SetUnitCost(v float64)`

SetUnitCost sets UnitCost field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


