# MetricExportIntegrationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloudwatch** | Pointer to [**CloudWatchConfig**](CloudWatchConfig.md) |  | [optional] 
**Datadog** | Pointer to [**DatadogConfig**](DatadogConfig.md) |  | [optional] 
**Status** | Pointer to [**MetricExportStatus**](MetricExportStatus.md) |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricExportIntegrationInfo

`func NewMetricExportIntegrationInfo() *MetricExportIntegrationInfo`

NewMetricExportIntegrationInfo instantiates a new MetricExportIntegrationInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetCloudwatch

`func (o *MetricExportIntegrationInfo) GetCloudwatch() CloudWatchConfig`

GetCloudwatch returns the Cloudwatch field if non-nil, zero value otherwise.

### SetCloudwatch

`func (o *MetricExportIntegrationInfo) SetCloudwatch(v CloudWatchConfig)`

SetCloudwatch sets Cloudwatch field to given value.

### GetDatadog

`func (o *MetricExportIntegrationInfo) GetDatadog() DatadogConfig`

GetDatadog returns the Datadog field if non-nil, zero value otherwise.

### SetDatadog

`func (o *MetricExportIntegrationInfo) SetDatadog(v DatadogConfig)`

SetDatadog sets Datadog field to given value.

### GetStatus

`func (o *MetricExportIntegrationInfo) GetStatus() MetricExportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *MetricExportIntegrationInfo) SetStatus(v MetricExportStatus)`

SetStatus sets Status field to given value.

### GetUserMessage

`func (o *MetricExportIntegrationInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *MetricExportIntegrationInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


