# PrivateEndpointService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneIds** | **[]string** | availability_zone_ids are the unique identifiers for the availability zones in which this service is available. Note these identifiers are unique even across typical cloud provider boundaries, for example AWS accounts or organizations. In AWS, availability zone ids for us-east-1 are use1-az1, use1-az2, use1-az3. | 
**Aws** | Pointer to [**AWSPrivateLinkServiceDetail**](AWSPrivateLinkServiceDetail.md) |  | [optional] 
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**EndpointServiceId** | **string** | endpoint_service_id uniquely identifies this private endpoint service. This is the cloud provider generated id for the service. | 
**Name** | **string** | name is the name of the private endpoints service. | 
**RegionName** | **string** | region_name is the cloud provider region name (e.g. us-east-1). | 
**Status** | [**PrivateEndpointServiceStatusType**](PrivateEndpointServiceStatusType.md) |  | 

## Methods

### NewPrivateEndpointService

`func NewPrivateEndpointService(availabilityZoneIds []string, cloudProvider CloudProviderType, endpointServiceId string, name string, regionName string, status PrivateEndpointServiceStatusType, ) *PrivateEndpointService`

NewPrivateEndpointService instantiates a new PrivateEndpointService object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPrivateEndpointServiceWithDefaults

`func NewPrivateEndpointServiceWithDefaults() *PrivateEndpointService`

NewPrivateEndpointServiceWithDefaults instantiates a new PrivateEndpointService object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAvailabilityZoneIds

`func (o *PrivateEndpointService) GetAvailabilityZoneIds() []string`

GetAvailabilityZoneIds returns the AvailabilityZoneIds field if non-nil, zero value otherwise.

### SetAvailabilityZoneIds

`func (o *PrivateEndpointService) SetAvailabilityZoneIds(v []string)`

SetAvailabilityZoneIds sets AvailabilityZoneIds field to given value.

### GetAws

`func (o *PrivateEndpointService) GetAws() AWSPrivateLinkServiceDetail`

GetAws returns the Aws field if non-nil, zero value otherwise.

### SetAws

`func (o *PrivateEndpointService) SetAws(v AWSPrivateLinkServiceDetail)`

SetAws sets Aws field to given value.

### GetCloudProvider

`func (o *PrivateEndpointService) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *PrivateEndpointService) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### GetEndpointServiceId

`func (o *PrivateEndpointService) GetEndpointServiceId() string`

GetEndpointServiceId returns the EndpointServiceId field if non-nil, zero value otherwise.

### SetEndpointServiceId

`func (o *PrivateEndpointService) SetEndpointServiceId(v string)`

SetEndpointServiceId sets EndpointServiceId field to given value.

### GetName

`func (o *PrivateEndpointService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *PrivateEndpointService) SetName(v string)`

SetName sets Name field to given value.

### GetRegionName

`func (o *PrivateEndpointService) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### SetRegionName

`func (o *PrivateEndpointService) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### GetStatus

`func (o *PrivateEndpointService) GetStatus() PrivateEndpointServiceStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *PrivateEndpointService) SetStatus(v PrivateEndpointServiceStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


