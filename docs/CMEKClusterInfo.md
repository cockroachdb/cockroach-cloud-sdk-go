# CMEKClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] [default to CMEKSTATUS_UNKNOWN_STATUS]
**RegionInfos** | Pointer to [**[]CMEKRegionInfo**](CMEKRegionInfo.md) |  | [optional] 

## Methods

### NewCMEKClusterInfo

`func NewCMEKClusterInfo() *CMEKClusterInfo`

NewCMEKClusterInfo instantiates a new CMEKClusterInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetStatus

`func (o *CMEKClusterInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *CMEKClusterInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### GetRegionInfos

`func (o *CMEKClusterInfo) GetRegionInfos() []CMEKRegionInfo`

GetRegionInfos returns the RegionInfos field if non-nil, zero value otherwise.

### SetRegionInfos

`func (o *CMEKClusterInfo) SetRegionInfos(v []CMEKRegionInfo)`

SetRegionInfos sets RegionInfos field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


