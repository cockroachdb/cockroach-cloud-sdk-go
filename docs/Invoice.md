# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjustments** | Pointer to [**[]InvoiceAdjustment**](InvoiceAdjustment.md) | adjustments is a list of credits or costs that adjust the value of the invoice (e.g. a Serverless Free Credit or Premium Support adjustment). Unlike line items, adjustments are not tied to a particular cluster. | [optional] 
**Balances** | [**[]CurrencyAmount**](CurrencyAmount.md) | balances are the amounts of currency left at the time of the invoice. | 
**FolderResources** | Pointer to [**[]FolderResource**](FolderResource.md) | Preview: FolderResources is a list of resources in the organization and their folder paths. | [optional] 
**InvoiceId** | **string** | invoice_id is the unique ID representing the invoice. | 
**InvoiceItems** | [**[]InvoiceItem**](InvoiceItem.md) | invoice_items are sorted by the cluster name. | 
**PeriodEnd** | **time.Time** | period_end is the end of the billing period (exclusive). | 
**PeriodStart** | **time.Time** | period_start is the start of the billing period (inclusive). | 
**Status** | Pointer to [**InvoiceStatusType**](InvoiceStatusType.md) |  | [optional] 
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

### GetAdjustments

`func (o *Invoice) GetAdjustments() []InvoiceAdjustment`

GetAdjustments returns the Adjustments field if non-nil, zero value otherwise.

### SetAdjustments

`func (o *Invoice) SetAdjustments(v []InvoiceAdjustment)`

SetAdjustments sets Adjustments field to given value.

### GetBalances

`func (o *Invoice) GetBalances() []CurrencyAmount`

GetBalances returns the Balances field if non-nil, zero value otherwise.

### SetBalances

`func (o *Invoice) SetBalances(v []CurrencyAmount)`

SetBalances sets Balances field to given value.

### GetFolderResources

`func (o *Invoice) GetFolderResources() []FolderResource`

GetFolderResources returns the FolderResources field if non-nil, zero value otherwise.

### SetFolderResources

`func (o *Invoice) SetFolderResources(v []FolderResource)`

SetFolderResources sets FolderResources field to given value.

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

### GetStatus

`func (o *Invoice) GetStatus() InvoiceStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *Invoice) SetStatus(v InvoiceStatusType)`

SetStatus sets Status field to given value.

### GetTotals

`func (o *Invoice) GetTotals() []CurrencyAmount`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### SetTotals

`func (o *Invoice) SetTotals(v []CurrencyAmount)`

SetTotals sets Totals field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


