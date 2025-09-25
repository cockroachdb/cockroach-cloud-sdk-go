# GcpCustomerCloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountEmail** | **string** | The customer-owned GCP service account email. This is the principal that will be impersonated by CockroachDB Cloud to perform actions in the customer-owned GCP project. The service account email contains the project ID, so CockroachDB Cloud doesn&#39;t require it to be passed separately. | 

## Methods

### NewGcpCustomerCloudAccount

`func NewGcpCustomerCloudAccount(serviceAccountEmail string, ) *GcpCustomerCloudAccount`

NewGcpCustomerCloudAccount instantiates a new GcpCustomerCloudAccount object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewGcpCustomerCloudAccountWithDefaults

`func NewGcpCustomerCloudAccountWithDefaults() *GcpCustomerCloudAccount`

NewGcpCustomerCloudAccountWithDefaults instantiates a new GcpCustomerCloudAccount object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetServiceAccountEmail

`func (o *GcpCustomerCloudAccount) GetServiceAccountEmail() string`

GetServiceAccountEmail returns the ServiceAccountEmail field if non-nil, zero value otherwise.

### SetServiceAccountEmail

`func (o *GcpCustomerCloudAccount) SetServiceAccountEmail(v string)`

SetServiceAccountEmail sets ServiceAccountEmail field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


