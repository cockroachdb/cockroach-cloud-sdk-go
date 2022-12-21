# InvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**Cluster**](Cluster.md) |  | 
**LineItems** | [**[]LineItem**](LineItem.md) | line_items contain all the relevant line items from the Metronome invoice. | 
**Totals** | [**[]CurrencyAmount**](CurrencyAmount.md) | totals is a list of the total amounts of line items per currency. | 

## Methods

### NewInvoiceItem

`func NewInvoiceItem(cluster Cluster, lineItems []LineItem, totals []CurrencyAmount, ) *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCluster

`func (o *InvoiceItem) GetCluster() Cluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### SetCluster

`func (o *InvoiceItem) SetCluster(v Cluster)`

SetCluster sets Cluster field to given value.

### GetLineItems

`func (o *InvoiceItem) GetLineItems() []LineItem`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### SetLineItems

`func (o *InvoiceItem) SetLineItems(v []LineItem)`

SetLineItems sets LineItems field to given value.

### GetTotals

`func (o *InvoiceItem) GetTotals() []CurrencyAmount`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### SetTotals

`func (o *InvoiceItem) SetTotals(v []CurrencyAmount)`

SetTotals sets Totals field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


