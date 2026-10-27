# OtlpMetricExportCloudWatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | external_id to include when assuming the IAM role specified by role_arn. Optional. A specific value may be required by the role&#39;s trust policy. | [optional] 
**RoleArn** | **string** | role_arn is the IAM role assumed to send metrics via SigV4. It requires the cloudwatch:PutMetricData permission. | 
**TargetRegion** | Pointer to **string** | target_region specifies the AWS region that metrics are exported to. It must support CloudWatch OTLP ingestion and PromQL. If empty, metrics export to each deployment region. | [optional] 

## Methods

### NewOtlpMetricExportCloudWatch

`func NewOtlpMetricExportCloudWatch(roleArn string, ) *OtlpMetricExportCloudWatch`

NewOtlpMetricExportCloudWatch instantiates a new OtlpMetricExportCloudWatch object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOtlpMetricExportCloudWatchWithDefaults

`func NewOtlpMetricExportCloudWatchWithDefaults() *OtlpMetricExportCloudWatch`

NewOtlpMetricExportCloudWatchWithDefaults instantiates a new OtlpMetricExportCloudWatch object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExternalId

`func (o *OtlpMetricExportCloudWatch) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *OtlpMetricExportCloudWatch) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetRoleArn

`func (o *OtlpMetricExportCloudWatch) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### SetRoleArn

`func (o *OtlpMetricExportCloudWatch) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### GetTargetRegion

`func (o *OtlpMetricExportCloudWatch) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### SetTargetRegion

`func (o *OtlpMetricExportCloudWatch) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


