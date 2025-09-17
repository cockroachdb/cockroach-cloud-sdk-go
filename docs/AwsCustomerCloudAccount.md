# AwsCustomerCloudAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arn** | **string** | The ARN string of the role that CockroachDB Cloud will assume to perform actions in the customer-owned AWS account. The ARN contains the account ID so CockroachDB Cloud doesn&#39;t require it to be passed separately. | 

## Methods

### NewAwsCustomerCloudAccount

`func NewAwsCustomerCloudAccount(arn string, ) *AwsCustomerCloudAccount`

NewAwsCustomerCloudAccount instantiates a new AwsCustomerCloudAccount object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAwsCustomerCloudAccountWithDefaults

`func NewAwsCustomerCloudAccountWithDefaults() *AwsCustomerCloudAccount`

NewAwsCustomerCloudAccountWithDefaults instantiates a new AwsCustomerCloudAccount object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetArn

`func (o *AwsCustomerCloudAccount) GetArn() string`

GetArn returns the Arn field if non-nil, zero value otherwise.

### SetArn

`func (o *AwsCustomerCloudAccount) SetArn(v string)`

SetArn sets Arn field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


