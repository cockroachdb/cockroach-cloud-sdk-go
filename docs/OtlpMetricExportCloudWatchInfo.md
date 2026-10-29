# OtlpMetricExportCloudWatchInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | external_id, if set, is included when assuming the IAM role. | [optional] 
**RoleArn** | **string** | role_arn is the IAM role assumed to send metrics via SigV4. | 
**TargetRegion** | Pointer to **string** | target_region specifies the AWS region that metrics are exported to. If empty, metrics export to each deployment region. | [optional] 

## Methods

### NewOtlpMetricExportCloudWatchInfo

`func NewOtlpMetricExportCloudWatchInfo(roleArn string, ) *OtlpMetricExportCloudWatchInfo`

NewOtlpMetricExportCloudWatchInfo instantiates a new OtlpMetricExportCloudWatchInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOtlpMetricExportCloudWatchInfoWithDefaults

`func NewOtlpMetricExportCloudWatchInfoWithDefaults() *OtlpMetricExportCloudWatchInfo`

NewOtlpMetricExportCloudWatchInfoWithDefaults instantiates a new OtlpMetricExportCloudWatchInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExternalId

`func (o *OtlpMetricExportCloudWatchInfo) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### SetExternalId

`func (o *OtlpMetricExportCloudWatchInfo) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### GetRoleArn

`func (o *OtlpMetricExportCloudWatchInfo) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### SetRoleArn

`func (o *OtlpMetricExportCloudWatchInfo) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### GetTargetRegion

`func (o *OtlpMetricExportCloudWatchInfo) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### SetTargetRegion

`func (o *OtlpMetricExportCloudWatchInfo) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


