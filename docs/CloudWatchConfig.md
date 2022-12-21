# CloudWatchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogGroupName** | Pointer to **string** | log_group_name is the customized log group name. | [optional] 
**RoleArn** | Pointer to **string** | role_arn is the IAM role used to upload metric segments to the target AWS account. | [optional] 
**TargetRegion** | Pointer to **string** | target_region specifies the specific AWS region that the metrics will be exported to. | [optional] 

## Methods

### NewCloudWatchConfig

`func NewCloudWatchConfig() *CloudWatchConfig`

NewCloudWatchConfig instantiates a new CloudWatchConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetLogGroupName

`func (o *CloudWatchConfig) GetLogGroupName() string`

GetLogGroupName returns the LogGroupName field if non-nil, zero value otherwise.

### SetLogGroupName

`func (o *CloudWatchConfig) SetLogGroupName(v string)`

SetLogGroupName sets LogGroupName field to given value.

### GetRoleArn

`func (o *CloudWatchConfig) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### SetRoleArn

`func (o *CloudWatchConfig) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### GetTargetRegion

`func (o *CloudWatchConfig) GetTargetRegion() string`

GetTargetRegion returns the TargetRegion field if non-nil, zero value otherwise.

### SetTargetRegion

`func (o *CloudWatchConfig) SetTargetRegion(v string)`

SetTargetRegion sets TargetRegion field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


