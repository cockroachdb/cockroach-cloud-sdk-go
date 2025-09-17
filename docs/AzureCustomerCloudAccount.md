# AzureCustomerCloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubscriptionId** | **string** | The ID of the subscription in the customer-owned Azure tenant in which clusters will be created. | 
**TenantId** | **string** | The ID of the customer-owned tenant in which the subscription exists. | 

## Methods

### NewAzureCustomerCloudAccount

`func NewAzureCustomerCloudAccount(subscriptionId string, tenantId string, ) *AzureCustomerCloudAccount`

NewAzureCustomerCloudAccount instantiates a new AzureCustomerCloudAccount object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAzureCustomerCloudAccountWithDefaults

`func NewAzureCustomerCloudAccountWithDefaults() *AzureCustomerCloudAccount`

NewAzureCustomerCloudAccountWithDefaults instantiates a new AzureCustomerCloudAccount object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetSubscriptionId

`func (o *AzureCustomerCloudAccount) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### SetSubscriptionId

`func (o *AzureCustomerCloudAccount) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### GetTenantId

`func (o *AzureCustomerCloudAccount) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### SetTenantId

`func (o *AzureCustomerCloudAccount) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


