# ClusterMajorVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedUpgrades** | **[]string** |  | 
**ReleaseType** | [**ReleaseTypeType**](ReleaseTypeType.md) |  | 
**SupportEnd** | **time.Time** |  | 
**SupportStatus** | [**ClusterMajorVersionSupportStatusType**](ClusterMajorVersionSupportStatusType.md) |  | 
**Version** | **string** |  | 

## Methods

### NewClusterMajorVersion

`func NewClusterMajorVersion(allowedUpgrades []string, releaseType ReleaseTypeType, supportEnd time.Time, supportStatus ClusterMajorVersionSupportStatusType, version string, ) *ClusterMajorVersion`

NewClusterMajorVersion instantiates a new ClusterMajorVersion object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewClusterMajorVersionWithDefaults

`func NewClusterMajorVersionWithDefaults() *ClusterMajorVersion`

NewClusterMajorVersionWithDefaults instantiates a new ClusterMajorVersion object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAllowedUpgrades

`func (o *ClusterMajorVersion) GetAllowedUpgrades() []string`

GetAllowedUpgrades returns the AllowedUpgrades field if non-nil, zero value otherwise.

### SetAllowedUpgrades

`func (o *ClusterMajorVersion) SetAllowedUpgrades(v []string)`

SetAllowedUpgrades sets AllowedUpgrades field to given value.

### GetReleaseType

`func (o *ClusterMajorVersion) GetReleaseType() ReleaseTypeType`

GetReleaseType returns the ReleaseType field if non-nil, zero value otherwise.

### SetReleaseType

`func (o *ClusterMajorVersion) SetReleaseType(v ReleaseTypeType)`

SetReleaseType sets ReleaseType field to given value.

### GetSupportEnd

`func (o *ClusterMajorVersion) GetSupportEnd() time.Time`

GetSupportEnd returns the SupportEnd field if non-nil, zero value otherwise.

### SetSupportEnd

`func (o *ClusterMajorVersion) SetSupportEnd(v time.Time)`

SetSupportEnd sets SupportEnd field to given value.

### GetSupportStatus

`func (o *ClusterMajorVersion) GetSupportStatus() ClusterMajorVersionSupportStatusType`

GetSupportStatus returns the SupportStatus field if non-nil, zero value otherwise.

### SetSupportStatus

`func (o *ClusterMajorVersion) SetSupportStatus(v ClusterMajorVersionSupportStatusType)`

SetSupportStatus sets SupportStatus field to given value.

### GetVersion

`func (o *ClusterMajorVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### SetVersion

`func (o *ClusterMajorVersion) SetVersion(v string)`

SetVersion sets Version field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


