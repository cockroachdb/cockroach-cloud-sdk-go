# Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**RegionName** | **string** |  | 
**Status** | [**NodeStatusType**](NodeStatusType.md) |  | 

## Methods

### NewNode

`func NewNode(name string, regionName string, status NodeStatusType, ) *Node`

NewNode instantiates a new Node object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewNodeWithDefaults

`func NewNodeWithDefaults() *Node`

NewNodeWithDefaults instantiates a new Node object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Node) SetName(v string)`

SetName sets Name field to given value.

### GetRegionName

`func (o *Node) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### SetRegionName

`func (o *Node) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### GetStatus

`func (o *Node) GetStatus() NodeStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *Node) SetStatus(v NodeStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


