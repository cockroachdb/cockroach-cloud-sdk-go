# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SqlDns** | **string** |  | 
**UiDns** | **string** |  | 
**NodeCount** | **int32** | NodeCount will be 0 for serverless clusters. | 

## Methods

### NewRegion

`func NewRegion(name string, sqlDns string, uiDns string, nodeCount int32, ) *Region`

NewRegion instantiates a new Region object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### GetSqlDns

`func (o *Region) GetSqlDns() string`

GetSqlDns returns the SqlDns field if non-nil, zero value otherwise.

### SetSqlDns

`func (o *Region) SetSqlDns(v string)`

SetSqlDns sets SqlDns field to given value.

### GetUiDns

`func (o *Region) GetUiDns() string`

GetUiDns returns the UiDns field if non-nil, zero value otherwise.

### SetUiDns

`func (o *Region) SetUiDns(v string)`

SetUiDns sets UiDns field to given value.

### GetNodeCount

`func (o *Region) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### SetNodeCount

`func (o *Region) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


