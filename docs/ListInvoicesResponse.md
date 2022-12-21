# ListInvoicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | [**[]Invoice**](Invoice.md) | invoices are sorted by period_start time. | 

## Methods

### NewListInvoicesResponse

`func NewListInvoicesResponse(invoices []Invoice, ) *ListInvoicesResponse`

NewListInvoicesResponse instantiates a new ListInvoicesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListInvoicesResponseWithDefaults

`func NewListInvoicesResponseWithDefaults() *ListInvoicesResponse`

NewListInvoicesResponseWithDefaults instantiates a new ListInvoicesResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetInvoices

`func (o *ListInvoicesResponse) GetInvoices() []Invoice`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### SetInvoices

`func (o *ListInvoicesResponse) SetInvoices(v []Invoice)`

SetInvoices sets Invoices field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


