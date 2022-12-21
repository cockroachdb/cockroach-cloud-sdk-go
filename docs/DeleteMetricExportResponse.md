# DeleteMetricExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**MetricExportStatus**](MetricExportStatus.md) |  | [optional] 
**Type** | Pointer to [**MetricExportIntegrationType**](MetricExportIntegrationType.md) |  | [optional] 

## Methods

### NewDeleteMetricExportResponse

`func NewDeleteMetricExportResponse() *DeleteMetricExportResponse`

NewDeleteMetricExportResponse instantiates a new DeleteMetricExportResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetClusterId

`func (o *DeleteMetricExportResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *DeleteMetricExportResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetStatus

`func (o *DeleteMetricExportResponse) GetStatus() MetricExportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *DeleteMetricExportResponse) SetStatus(v MetricExportStatus)`

SetStatus sets Status field to given value.

### GetType

`func (o *DeleteMetricExportResponse) GetType() MetricExportIntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *DeleteMetricExportResponse) SetType(v MetricExportIntegrationType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


