# PrometheusMetricExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | Pointer to [**MetricExportStatusType**](MetricExportStatusType.md) |  | [optional] 
**Targets** | Pointer to **map[string]string** | targets is a map of ports exposing metrics to regions. | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewPrometheusMetricExportInfo

`func NewPrometheusMetricExportInfo(clusterId string, ) *PrometheusMetricExportInfo`

NewPrometheusMetricExportInfo instantiates a new PrometheusMetricExportInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPrometheusMetricExportInfoWithDefaults

`func NewPrometheusMetricExportInfoWithDefaults() *PrometheusMetricExportInfo`

NewPrometheusMetricExportInfoWithDefaults instantiates a new PrometheusMetricExportInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *PrometheusMetricExportInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *PrometheusMetricExportInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetStatus

`func (o *PrometheusMetricExportInfo) GetStatus() MetricExportStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *PrometheusMetricExportInfo) SetStatus(v MetricExportStatusType)`

SetStatus sets Status field to given value.

### GetTargets

`func (o *PrometheusMetricExportInfo) GetTargets() map[string]string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### SetTargets

`func (o *PrometheusMetricExportInfo) SetTargets(v map[string]string)`

SetTargets sets Targets field to given value.

### GetUserMessage

`func (o *PrometheusMetricExportInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *PrometheusMetricExportInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


