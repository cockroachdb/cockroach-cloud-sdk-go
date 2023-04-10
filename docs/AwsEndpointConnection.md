# AwsEndpointConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**EndpointId** | **string** | endpoint_id is the client side of the PrivateLink connection. | 
**RegionName** | **string** | region_name is the cloud provider region name (i.e. us-east-1). | 
**ServiceId** | **string** | service_id is the server side of the PrivateLink connection. This is the same as AWSPrivateLinkEndpoint.service_id. | 
**Status** | [**AWSEndpointConnectionStatusType**](AWSEndpointConnectionStatusType.md) |  | 

## Methods

### NewAwsEndpointConnection

`func NewAwsEndpointConnection(cloudProvider CloudProviderType, endpointId string, regionName string, serviceId string, status AWSEndpointConnectionStatusType, ) *AwsEndpointConnection`

NewAwsEndpointConnection instantiates a new AwsEndpointConnection object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAwsEndpointConnectionWithDefaults

`func NewAwsEndpointConnectionWithDefaults() *AwsEndpointConnection`

NewAwsEndpointConnectionWithDefaults instantiates a new AwsEndpointConnection object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCloudProvider

`func (o *AwsEndpointConnection) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *AwsEndpointConnection) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### GetEndpointId

`func (o *AwsEndpointConnection) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### SetEndpointId

`func (o *AwsEndpointConnection) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### GetRegionName

`func (o *AwsEndpointConnection) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### SetRegionName

`func (o *AwsEndpointConnection) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### GetServiceId

`func (o *AwsEndpointConnection) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### SetServiceId

`func (o *AwsEndpointConnection) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### GetStatus

`func (o *AwsEndpointConnection) GetStatus() AWSEndpointConnectionStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *AwsEndpointConnection) SetStatus(v AWSEndpointConnectionStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


