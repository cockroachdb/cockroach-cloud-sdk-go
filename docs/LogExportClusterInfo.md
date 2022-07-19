# LogExportClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**LogExportStatus**](LogExportStatus.md) |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**LogExportClusterSpecification**](LogExportClusterSpecification.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLogExportClusterInfo

`func NewLogExportClusterInfo() *LogExportClusterInfo`

NewLogExportClusterInfo instantiates a new LogExportClusterInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetClusterId

`func (o *LogExportClusterInfo) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *LogExportClusterInfo) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetStatus

`func (o *LogExportClusterInfo) GetStatus() LogExportStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *LogExportClusterInfo) SetStatus(v LogExportStatus)`

SetStatus sets Status field to given value.

### GetUserMessage

`func (o *LogExportClusterInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *LogExportClusterInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### GetSpec

`func (o *LogExportClusterInfo) GetSpec() LogExportClusterSpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### SetSpec

`func (o *LogExportClusterInfo) SetSpec(v LogExportClusterSpecification)`

SetSpec sets Spec field to given value.

### GetCreatedAt

`func (o *LogExportClusterInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *LogExportClusterInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetUpdatedAt

`func (o *LogExportClusterInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *LogExportClusterInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


