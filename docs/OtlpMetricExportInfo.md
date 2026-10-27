# OtlpMetricExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloudwatch** | Pointer to [**OtlpMetricExportCloudWatchInfo**](OtlpMetricExportCloudWatchInfo.md) |  | [optional] 
**ClusterId** | **string** | cluster_id is the ID of the cluster this metric export is configured for. | 
**Status** | [**MetricExportStatusType**](MetricExportStatusType.md) |  | 
**UserMessage** | **string** | user_message is a human-readable description of the current export status. | 

## Methods

### NewOtlpMetricExportInfo

`func NewOtlpMetricExportInfo(clusterId string, status MetricExportStatusType, userMessage string, ) *OtlpMetricExportInfo`

NewOtlpMetricExportInfo instantiates a new OtlpMetricExportInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOtlpMetricExportInfoWithDefaults

`func NewOtlpMetricExportInfoWithDefaults() *OtlpMetricExportInfo`

NewOtlpMetricExportInfoWithDefaults instantiates a new OtlpMetricExportInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCloudwatch

`func (o *OtlpMetricExportInfo) GetCloudwatch() OtlpMetricExportCloudWatchInfo`

GetCloudwatch returns the Cloudwatch field if non-nil, zero value otherwise.

### SetCloudwatch

`func (o *OtlpMetricExportInfo) SetCloudwatch(v OtlpMetricExportCloudWatchInfo)`

SetCloudwatch sets Cloudwatch field to given value.

### GetClusterId

`func (o *OtlpMetricExportInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *OtlpMetricExportInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetStatus

`func (o *OtlpMetricExportInfo) GetStatus() MetricExportStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *OtlpMetricExportInfo) SetStatus(v MetricExportStatusType)`

SetStatus sets Status field to given value.

### GetUserMessage

`func (o *OtlpMetricExportInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *OtlpMetricExportInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


