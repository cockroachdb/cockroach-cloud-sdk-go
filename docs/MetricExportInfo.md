# MetricExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Integration** | Pointer to [**[]MetricExportIntegrationInfo**](MetricExportIntegrationInfo.md) |  | [optional] 

## Methods

### NewMetricExportInfo

`func NewMetricExportInfo() *MetricExportInfo`

NewMetricExportInfo instantiates a new MetricExportInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetClusterId

`func (o *MetricExportInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *MetricExportInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetIntegration

`func (o *MetricExportInfo) GetIntegration() []MetricExportIntegrationInfo`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### SetIntegration

`func (o *MetricExportInfo) SetIntegration(v []MetricExportIntegrationInfo)`

SetIntegration sets Integration field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


