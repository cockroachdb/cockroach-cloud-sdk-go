# InvoiceAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | [**CurrencyAmount**](CurrencyAmount.md) |  | 
**Name** | **string** | name identifies the adjustment. | 

## Methods

### NewInvoiceAdjustment

`func NewInvoiceAdjustment(amount CurrencyAmount, name string, ) *InvoiceAdjustment`

NewInvoiceAdjustment instantiates a new InvoiceAdjustment object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewInvoiceAdjustmentWithDefaults

`func NewInvoiceAdjustmentWithDefaults() *InvoiceAdjustment`

NewInvoiceAdjustmentWithDefaults instantiates a new InvoiceAdjustment object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAmount

`func (o *InvoiceAdjustment) GetAmount() CurrencyAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### SetAmount

`func (o *InvoiceAdjustment) SetAmount(v CurrencyAmount)`

SetAmount sets Amount field to given value.

### GetName

`func (o *InvoiceAdjustment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *InvoiceAdjustment) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


