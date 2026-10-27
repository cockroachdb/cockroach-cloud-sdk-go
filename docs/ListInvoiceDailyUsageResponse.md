# ListInvoiceDailyUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**UsageItems** | [**[]InvoiceDailyUsageItem**](InvoiceDailyUsageItem.md) | usage_items contains the daily usage rows for this page. By default, usage rows are sorted by usage date and then by cluster, with adjustment rows on the final pages. In descending order, adjustment rows appear first, followed by usage rows from newest to oldest. | 

## Methods

### NewListInvoiceDailyUsageResponse

`func NewListInvoiceDailyUsageResponse(usageItems []InvoiceDailyUsageItem, ) *ListInvoiceDailyUsageResponse`

NewListInvoiceDailyUsageResponse instantiates a new ListInvoiceDailyUsageResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListInvoiceDailyUsageResponseWithDefaults

`func NewListInvoiceDailyUsageResponseWithDefaults() *ListInvoiceDailyUsageResponse`

NewListInvoiceDailyUsageResponseWithDefaults instantiates a new ListInvoiceDailyUsageResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ListInvoiceDailyUsageResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListInvoiceDailyUsageResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetUsageItems

`func (o *ListInvoiceDailyUsageResponse) GetUsageItems() []InvoiceDailyUsageItem`

GetUsageItems returns the UsageItems field if non-nil, zero value otherwise.

### SetUsageItems

`func (o *ListInvoiceDailyUsageResponse) SetUsageItems(v []InvoiceDailyUsageItem)`

SetUsageItems sets UsageItems field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


