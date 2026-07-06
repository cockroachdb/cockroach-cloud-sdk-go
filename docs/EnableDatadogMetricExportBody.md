# EnableDatadogMetricExportBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | api_key is a Datadog API key. | 
**Site** | [**DatadogSiteType**](DatadogSiteType.md) |  | 

## Methods

### NewEnableDatadogMetricExportBody

`func NewEnableDatadogMetricExportBody(apiKey string, site DatadogSiteType, ) *EnableDatadogMetricExportBody`

NewEnableDatadogMetricExportBody instantiates a new EnableDatadogMetricExportBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEnableDatadogMetricExportBodyWithDefaults

`func NewEnableDatadogMetricExportBodyWithDefaults() *EnableDatadogMetricExportBody`

NewEnableDatadogMetricExportBodyWithDefaults instantiates a new EnableDatadogMetricExportBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetApiKey

`func (o *EnableDatadogMetricExportBody) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### SetApiKey

`func (o *EnableDatadogMetricExportBody) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### GetSite

`func (o *EnableDatadogMetricExportBody) GetSite() DatadogSiteType`

GetSite returns the Site field if non-nil, zero value otherwise.

### SetSite

`func (o *EnableDatadogMetricExportBody) SetSite(v DatadogSiteType)`

SetSite sets Site field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


