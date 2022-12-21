# EnableMetricExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cloudwatch** | Pointer to [**CloudWatchConfig**](CloudWatchConfig.md) |  | [optional] 
**Datadog** | Pointer to [**DatadogConfig**](DatadogConfig.md) |  | [optional] 

## Methods

### NewEnableMetricExportRequest

`func NewEnableMetricExportRequest() *EnableMetricExportRequest`

NewEnableMetricExportRequest instantiates a new EnableMetricExportRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetCloudwatch

`func (o *EnableMetricExportRequest) GetCloudwatch() CloudWatchConfig`

GetCloudwatch returns the Cloudwatch field if non-nil, zero value otherwise.

### SetCloudwatch

`func (o *EnableMetricExportRequest) SetCloudwatch(v CloudWatchConfig)`

SetCloudwatch sets Cloudwatch field to given value.

### GetDatadog

`func (o *EnableMetricExportRequest) GetDatadog() DatadogConfig`

GetDatadog returns the Datadog field if non-nil, zero value otherwise.

### SetDatadog

`func (o *EnableMetricExportRequest) SetDatadog(v DatadogConfig)`

SetDatadog sets Datadog field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


