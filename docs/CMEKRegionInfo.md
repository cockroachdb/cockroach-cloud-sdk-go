# CMEKRegionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] 
**KeyInfos** | Pointer to [**[]CMEKKeyInfo**](CMEKKeyInfo.md) |  | [optional] 

## Methods

### NewCMEKRegionInfo

`func NewCMEKRegionInfo() *CMEKRegionInfo`

NewCMEKRegionInfo instantiates a new CMEKRegionInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetRegion

`func (o *CMEKRegionInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *CMEKRegionInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetStatus

`func (o *CMEKRegionInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *CMEKRegionInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### GetKeyInfos

`func (o *CMEKRegionInfo) GetKeyInfos() []CMEKKeyInfo`

GetKeyInfos returns the KeyInfos field if non-nil, zero value otherwise.

### SetKeyInfos

`func (o *CMEKRegionInfo) SetKeyInfos(v []CMEKKeyInfo)`

SetKeyInfos sets KeyInfos field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


