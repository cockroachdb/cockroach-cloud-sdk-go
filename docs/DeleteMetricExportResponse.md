# DeleteMetricExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** |  | 
**Status** | Pointer to [**MetricExportStatusType**](MetricExportStatusType.md) |  | [optional] 

## Methods

### NewDeleteMetricExportResponse

`func NewDeleteMetricExportResponse(clusterId string, ) *DeleteMetricExportResponse`

NewDeleteMetricExportResponse instantiates a new DeleteMetricExportResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDeleteMetricExportResponseWithDefaults

`func NewDeleteMetricExportResponseWithDefaults() *DeleteMetricExportResponse`

NewDeleteMetricExportResponseWithDefaults instantiates a new DeleteMetricExportResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *DeleteMetricExportResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *DeleteMetricExportResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetStatus

`func (o *DeleteMetricExportResponse) GetStatus() MetricExportStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *DeleteMetricExportResponse) SetStatus(v MetricExportStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


