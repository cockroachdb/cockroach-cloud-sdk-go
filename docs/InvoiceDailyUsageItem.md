# InvoiceDailyUsageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChargeCategory** | [**ChargeCategoryType**](ChargeCategoryType.md) |  | 
**Cluster** | Pointer to [**DailyUsageCluster**](DailyUsageCluster.md) |  | [optional] 
**LineItem** | [**DailyUsageLineItem**](DailyUsageLineItem.md) |  | 
**UsageDate** | Pointer to **time.Time** | usage_date is the day the usage was incurred. Null for adjustment rows. | [optional] 

## Methods

### NewInvoiceDailyUsageItem

`func NewInvoiceDailyUsageItem(chargeCategory ChargeCategoryType, lineItem DailyUsageLineItem, ) *InvoiceDailyUsageItem`

NewInvoiceDailyUsageItem instantiates a new InvoiceDailyUsageItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewInvoiceDailyUsageItemWithDefaults

`func NewInvoiceDailyUsageItemWithDefaults() *InvoiceDailyUsageItem`

NewInvoiceDailyUsageItemWithDefaults instantiates a new InvoiceDailyUsageItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetChargeCategory

`func (o *InvoiceDailyUsageItem) GetChargeCategory() ChargeCategoryType`

GetChargeCategory returns the ChargeCategory field if non-nil, zero value otherwise.

### SetChargeCategory

`func (o *InvoiceDailyUsageItem) SetChargeCategory(v ChargeCategoryType)`

SetChargeCategory sets ChargeCategory field to given value.

### GetCluster

`func (o *InvoiceDailyUsageItem) GetCluster() DailyUsageCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### SetCluster

`func (o *InvoiceDailyUsageItem) SetCluster(v DailyUsageCluster)`

SetCluster sets Cluster field to given value.

### GetLineItem

`func (o *InvoiceDailyUsageItem) GetLineItem() DailyUsageLineItem`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### SetLineItem

`func (o *InvoiceDailyUsageItem) SetLineItem(v DailyUsageLineItem)`

SetLineItem sets LineItem field to given value.

### GetUsageDate

`func (o *InvoiceDailyUsageItem) GetUsageDate() time.Time`

GetUsageDate returns the UsageDate field if non-nil, zero value otherwise.

### SetUsageDate

`func (o *InvoiceDailyUsageItem) SetUsageDate(v time.Time)`

SetUsageDate sets UsageDate field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


