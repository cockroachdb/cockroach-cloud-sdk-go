# Region

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InternalDns** | **string** | internal_dns is the internal DNS name of the cluster within the cloud provider&#39;s network. It is used to connect to the cluster with PrivateLink or VPC peering. | 
**Name** | **string** |  | 
**NodeCount** | **int32** | node_count will be 0 for Serverless clusters. | 
**Primary** | Pointer to **bool** | primary is true only for the primary region in a Multi Region Serverless cluster. | [optional] 
**SqlDns** | **string** | sql_dns is the DNS name of SQL interface of the cluster. It is used to connect to the cluster with IP allowlisting. | 
**UiDns** | **string** | ui_dns is the DNS name used when connecting to the DB Console for the cluster. | 

## Methods

### NewRegion

`func NewRegion(internalDns string, name string, nodeCount int32, sqlDns string, uiDns string, ) *Region`

NewRegion instantiates a new Region object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRegionWithDefaults

`func NewRegionWithDefaults() *Region`

NewRegionWithDefaults instantiates a new Region object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetInternalDns

`func (o *Region) GetInternalDns() string`

GetInternalDns returns the InternalDns field if non-nil, zero value otherwise.

### SetInternalDns

`func (o *Region) SetInternalDns(v string)`

SetInternalDns sets InternalDns field to given value.

### GetName

`func (o *Region) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Region) SetName(v string)`

SetName sets Name field to given value.

### GetNodeCount

`func (o *Region) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### SetNodeCount

`func (o *Region) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### GetPrimary

`func (o *Region) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### SetPrimary

`func (o *Region) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


