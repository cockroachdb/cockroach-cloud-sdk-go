# EnableDatadogMetricExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | api_key is a Datadog API key. | 
**Site** | [**ApiDatadogSite**](ApiDatadogSite.md) |  | 

## Methods

### NewEnableDatadogMetricExportRequest

`func NewEnableDatadogMetricExportRequest(apiKey string, site ApiDatadogSite, ) *EnableDatadogMetricExportRequest`

NewEnableDatadogMetricExportRequest instantiates a new EnableDatadogMetricExportRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEnableDatadogMetricExportRequestWithDefaults

`func NewEnableDatadogMetricExportRequestWithDefaults() *EnableDatadogMetricExportRequest`

NewEnableDatadogMetricExportRequestWithDefaults instantiates a new EnableDatadogMetricExportRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiKey

`func (o *EnableDatadogMetricExportRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### SetApiKey

`func (o *EnableDatadogMetricExportRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### GetSite

`func (o *EnableDatadogMetricExportRequest) GetSite() ApiDatadogSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### SetSite

`func (o *EnableDatadogMetricExportRequest) SetSite(v ApiDatadogSite)`

SetSite sets Site field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


