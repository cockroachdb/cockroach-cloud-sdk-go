# PrivateEndpointService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionName** | **string** | RegionName is the cloud provider region name (i.e. us-east-1). | 
**CloudProvider** | [**ApiCloudProvider**](ApiCloudProvider.md) |  | 
**Status** | [**PrivateEndpointServiceStatus**](PrivateEndpointServiceStatus.md) |  | 
**Aws** | [**AWSPrivateLinkServiceDetail**](AWSPrivateLinkServiceDetail.md) |  | 

## Methods

### NewPrivateEndpointService

`func NewPrivateEndpointService(regionName string, cloudProvider ApiCloudProvider, status PrivateEndpointServiceStatus, aws AWSPrivateLinkServiceDetail, ) *PrivateEndpointService`

NewPrivateEndpointService instantiates a new PrivateEndpointService object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPrivateEndpointServiceWithDefaults

`func NewPrivateEndpointServiceWithDefaults() *PrivateEndpointService`

NewPrivateEndpointServiceWithDefaults instantiates a new PrivateEndpointService object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRegionName

`func (o *PrivateEndpointService) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### SetRegionName

`func (o *PrivateEndpointService) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### GetCloudProvider

`func (o *PrivateEndpointService) GetCloudProvider() ApiCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *PrivateEndpointService) SetCloudProvider(v ApiCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.

### GetStatus

`func (o *PrivateEndpointService) GetStatus() PrivateEndpointServiceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *PrivateEndpointService) SetStatus(v PrivateEndpointServiceStatus)`

SetStatus sets Status field to given value.

### GetAws

`func (o *PrivateEndpointService) GetAws() AWSPrivateLinkServiceDetail`

GetAws returns the Aws field if non-nil, zero value otherwise.

### SetAws

`func (o *PrivateEndpointService) SetAws(v AWSPrivateLinkServiceDetail)`

SetAws sets Aws field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


