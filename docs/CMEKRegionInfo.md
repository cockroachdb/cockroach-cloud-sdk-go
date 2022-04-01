# CMEKRegionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**KeyInfos** | Pointer to [**[]CMEKKeyInfo**](CMEKKeyInfo.md) |  | [optional] 

## Methods

### NewCMEKRegionInfo

`func NewCMEKRegionInfo() *CMEKRegionInfo`

NewCMEKRegionInfo instantiates a new CMEKRegionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKRegionInfoWithDefaults

`func NewCMEKRegionInfoWithDefaults() *CMEKRegionInfo`

NewCMEKRegionInfoWithDefaults instantiates a new CMEKRegionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CMEKRegionInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CMEKRegionInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CMEKRegionInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CMEKRegionInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetKeyInfos

`func (o *CMEKRegionInfo) GetKeyInfos() []CMEKKeyInfo`

GetKeyInfos returns the KeyInfos field if non-nil, zero value otherwise.

### GetKeyInfosOk

`func (o *CMEKRegionInfo) GetKeyInfosOk() (*[]CMEKKeyInfo, bool)`

GetKeyInfosOk returns a tuple with the KeyInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyInfos

`func (o *CMEKRegionInfo) SetKeyInfos(v []CMEKKeyInfo)`

SetKeyInfos sets KeyInfos field to given value.

### HasKeyInfos

`func (o *CMEKRegionInfo) HasKeyInfos() bool`

HasKeyInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


