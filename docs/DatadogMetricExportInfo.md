# DatadogMetricExportInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | api_key is the last 4 digits of a Datadog API key. | [optional] 
**ClusterId** | **string** |  | 
**Site** | [**DatadogSiteType**](DatadogSiteType.md) |  | 
**Status** | Pointer to [**MetricExportStatusType**](MetricExportStatusType.md) |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewDatadogMetricExportInfo

`func NewDatadogMetricExportInfo(clusterId string, site DatadogSiteType, ) *DatadogMetricExportInfo`

NewDatadogMetricExportInfo instantiates a new DatadogMetricExportInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDatadogMetricExportInfoWithDefaults

`func NewDatadogMetricExportInfoWithDefaults() *DatadogMetricExportInfo`

NewDatadogMetricExportInfoWithDefaults instantiates a new DatadogMetricExportInfo object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiKey

`func (o *DatadogMetricExportInfo) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### SetApiKey

`func (o *DatadogMetricExportInfo) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### GetClusterId

`func (o *DatadogMetricExportInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *DatadogMetricExportInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetSite

`func (o *DatadogMetricExportInfo) GetSite() DatadogSiteType`

GetSite returns the Site field if non-nil, zero value otherwise.

### SetSite

`func (o *DatadogMetricExportInfo) SetSite(v DatadogSiteType)`

SetSite sets Site field to given value.

### GetStatus

`func (o *DatadogMetricExportInfo) GetStatus() MetricExportStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *DatadogMetricExportInfo) SetStatus(v MetricExportStatusType)`

SetStatus sets Status field to given value.

### GetUserMessage

`func (o *DatadogMetricExportInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *DatadogMetricExportInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


