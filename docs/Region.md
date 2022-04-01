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

NewRegion instantiates a new Region object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Region) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.


### GetSqlDns

`func (o *Region) GetSqlDns() string`

GetSqlDns returns the SqlDns field if non-nil, zero value otherwise.

### GetSqlDnsOk

`func (o *Region) GetSqlDnsOk() (*string, bool)`

GetSqlDnsOk returns a tuple with the SqlDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqlDns

`func (o *Region) SetSqlDns(v string)`

SetSqlDns sets SqlDns field to given value.


### GetUiDns

`func (o *Region) GetUiDns() string`

GetUiDns returns the UiDns field if non-nil, zero value otherwise.

### GetUiDnsOk

`func (o *Region) GetUiDnsOk() (*string, bool)`

GetUiDnsOk returns a tuple with the UiDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiDns

`func (o *Region) SetUiDns(v string)`

SetUiDns sets UiDns field to given value.


### GetNodeCount

`func (o *Region) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *Region) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *Region) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


