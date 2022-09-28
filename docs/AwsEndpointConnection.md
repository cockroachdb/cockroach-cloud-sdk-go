# AwsEndpointConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegionName** | **string** | RegionName is the cloud provider region name (i.e. us-east-1). | 
**CloudProvider** | [**ApiCloudProvider**](ApiCloudProvider.md) |  | 
**Status** | [**AWSEndpointConnectionStatus**](AWSEndpointConnectionStatus.md) |  | 
**EndpointId** | **string** | EndpointID is the client side of the PrivateLink connection. | 
**ServiceId** | **string** | ServiceID is the server side of the PrivateLink connection. This is the same as AWSPrivateLinkEndpoint.service_id. | 

## Methods

### NewAwsEndpointConnection

`func NewAwsEndpointConnection(regionName string, cloudProvider ApiCloudProvider, status AWSEndpointConnectionStatus, endpointId string, serviceId string, ) *AwsEndpointConnection`

NewAwsEndpointConnection instantiates a new AwsEndpointConnection object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAwsEndpointConnectionWithDefaults

`func NewAwsEndpointConnectionWithDefaults() *AwsEndpointConnection`

NewAwsEndpointConnectionWithDefaults instantiates a new AwsEndpointConnection object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRegionName

`func (o *AwsEndpointConnection) GetRegionName() string`

GetRegionName returns the RegionName field if non-nil, zero value otherwise.

### SetRegionName

`func (o *AwsEndpointConnection) SetRegionName(v string)`

SetRegionName sets RegionName field to given value.

### GetCloudProvider

`func (o *AwsEndpointConnection) GetCloudProvider() ApiCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *AwsEndpointConnection) SetCloudProvider(v ApiCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.

### GetStatus

`func (o *AwsEndpointConnection) GetStatus() AWSEndpointConnectionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *AwsEndpointConnection) SetStatus(v AWSEndpointConnectionStatus)`

SetStatus sets Status field to given value.

### GetEndpointId

`func (o *AwsEndpointConnection) GetEndpointId() string`

GetEndpointId returns the EndpointId field if non-nil, zero value otherwise.

### SetEndpointId

`func (o *AwsEndpointConnection) SetEndpointId(v string)`

SetEndpointId sets EndpointId field to given value.

### GetServiceId

`func (o *AwsEndpointConnection) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### SetServiceId

`func (o *AwsEndpointConnection) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


