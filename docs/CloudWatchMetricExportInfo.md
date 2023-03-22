# CloudWatchMetricExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**LogGroupName** | Pointer to **string** | log_group_name is the customized log group name. | [optional] 
**RoleArn** | **string** | role_arn is the IAM role used to upload metric segments to the target AWS account. | 
**Status** | Pointer to [**MetricExportStatus**](MetricExportStatus.md) |  | [optional] 
**TargetRegion** | Pointer to **string** | target_region specifies the specific AWS region that the metrics will be exported to. | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudWatchMetricExportInfo

`func NewCloudWatchMetricExportInfo(clusterId string, roleArn string, ) *CloudWatchMetricExportInfo`

NewCloudWatchMetricExportInfo instantiates a new CloudWatchMetricExportInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudWatchMetricExportInfoWithDefaults

`func NewCloudWatchMetricExportInfoWithDefaults() *CloudWatchMetricExportInfo`

NewCloudWatchMetricExportInfoWithDefaults instantiates a new CloudWatchMetricExportInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *CloudWatchMetricExportInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *CloudWatchMetricExportInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetLogGroupName

`func (o *CloudWatchMetricExportInfo) GetLogGroupName() string`

GetLogGroupName returns the LogGroupName field if non-nil, zero value otherwise.

### SetLogGroupName

`func (o *CloudWatchMetricExportInfo) SetLogGroupName(v string)`

SetLogGroupName sets LogGroupName field to given value.

### GetRoleArn

`func (o *CloudWatchMetricExportInfo) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### SetRoleArn

`func (o *CloudWatchMetricExportInfo) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### GetStatus

`func (o *CloudWatchMetricExportInfo) GetStatus() MetricExportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *CloudWatchMetricExportInfo) SetStatus(v MetricExportStatus)`

SetStatus sets Status field to given value.

### GetTargetRegion

`func (o *CloudWatchMetricExportInfo) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### SetTargetRegion

`func (o *CloudWatchMetricExportInfo) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.

### GetUserMessage

`func (o *CloudWatchMetricExportInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *CloudWatchMetricExportInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


