# CurrencyAmount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float64** | amount is the quantity of currency. | [optional] 
**Currency** | Pointer to [**Currency**](Currency.md) |  | [optional] 

## Methods

### NewCurrencyAmount

`func NewCurrencyAmount() *CurrencyAmount`

NewCurrencyAmount instantiates a new CurrencyAmount object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAmount

`func (o *CurrencyAmount) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### SetAmount

`func (o *CurrencyAmount) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### GetCurrency

`func (o *CurrencyAmount) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### SetCurrency

`func (o *CurrencyAmount) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


