# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balances** | [**[]CurrencyAmount**](CurrencyAmount.md) | balances are the amounts of currency left at the time of the invoice. | 
**InvoiceId** | **string** | invoice_id is the unique ID representing the invoice. | 
**InvoiceItems** | [**[]InvoiceItem**](InvoiceItem.md) | invoice_items are sorted by the cluster name. | 
**PeriodEnd** | **time.Time** | period_end is the end of the billing period (exclusive). | 
**PeriodStart** | **time.Time** | period_start is the start of the billing period (inclusive). | 
**Totals** | [**[]CurrencyAmount**](CurrencyAmount.md) | totals is a list of the total amounts per currency. | 

## Methods

### NewInvoice

`func NewInvoice(balances []CurrencyAmount, invoiceId string, invoiceItems []InvoiceItem, periodEnd time.Time, periodStart time.Time, totals []CurrencyAmount, ) *Invoice`

NewInvoice instantiates a new Invoice object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBalances

`func (o *Invoice) GetBalances() []CurrencyAmount`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### SetBalances

`func (o *Invoice) SetBalances(v []CurrencyAmount)`

SetBalances sets Balances field to given value.

### GetInvoiceId

`func (o *Invoice) GetInvoiceId() string`

GetInvoiceId returns the InvoiceId field if non-nil, zero value otherwise.

### SetInvoiceId

`func (o *Invoice) SetInvoiceId(v string)`

SetInvoiceId sets InvoiceId field to given value.

### GetInvoiceItems

`func (o *Invoice) GetInvoiceItems() []InvoiceItem`

GetInvoiceItems returns the InvoiceItems field if non-nil, zero value otherwise.

### SetInvoiceItems

`func (o *Invoice) SetInvoiceItems(v []InvoiceItem)`

SetInvoiceItems sets InvoiceItems field to given value.

### GetPeriodEnd

`func (o *Invoice) GetPeriodEnd() time.Time`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### SetPeriodEnd

`func (o *Invoice) SetPeriodEnd(v time.Time)`

SetPeriodEnd sets PeriodEnd field to given value.

### GetPeriodStart

`func (o *Invoice) GetPeriodStart() time.Time`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### SetPeriodStart

`func (o *Invoice) SetPeriodStart(v time.Time)`

SetPeriodStart sets PeriodStart field to given value.

### GetTotals

`func (o *Invoice) GetTotals() []CurrencyAmount`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### SetTotals

`func (o *Invoice) SetTotals(v []CurrencyAmount)`

SetTotals sets Totals field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


