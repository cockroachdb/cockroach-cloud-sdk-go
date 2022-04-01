# CMEKClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] [default to UNKNOWN_STATUS]
**RegionInfos** | Pointer to [**[]CMEKRegionInfo**](CMEKRegionInfo.md) |  | [optional] 

## Methods

### NewCMEKClusterInfo

`func NewCMEKClusterInfo() *CMEKClusterInfo`

NewCMEKClusterInfo instantiates a new CMEKClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKClusterInfoWithDefaults

`func NewCMEKClusterInfoWithDefaults() *CMEKClusterInfo`

NewCMEKClusterInfoWithDefaults instantiates a new CMEKClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CMEKClusterInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CMEKClusterInfo) GetStatusOk() (*CMEKStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CMEKClusterInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CMEKClusterInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegionInfos

`func (o *CMEKClusterInfo) GetRegionInfos() []CMEKRegionInfo`

GetRegionInfos returns the RegionInfos field if non-nil, zero value otherwise.

### GetRegionInfosOk

`func (o *CMEKClusterInfo) GetRegionInfosOk() (*[]CMEKRegionInfo, bool)`

GetRegionInfosOk returns a tuple with the RegionInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInfos

`func (o *CMEKClusterInfo) SetRegionInfos(v []CMEKRegionInfo)`

SetRegionInfos sets RegionInfos field to given value.

### HasRegionInfos

`func (o *CMEKClusterInfo) HasRegionInfos() bool`

HasRegionInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


