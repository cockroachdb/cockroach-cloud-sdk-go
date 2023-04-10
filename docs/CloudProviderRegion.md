# CloudProviderRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distance** | **float32** | Distance in miles, based on client IP address. | 
**Location** | **string** |  | 
**Name** | **string** |  | 
**Provider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**Serverless** | **bool** |  | 

## Methods

### NewCloudProviderRegion

`func NewCloudProviderRegion(distance float32, location string, name string, provider CloudProviderType, serverless bool, ) *CloudProviderRegion`

NewCloudProviderRegion instantiates a new CloudProviderRegion object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCloudProviderRegionWithDefaults

`func NewCloudProviderRegionWithDefaults() *CloudProviderRegion`

NewCloudProviderRegionWithDefaults instantiates a new CloudProviderRegion object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDistance

`func (o *CloudProviderRegion) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### SetDistance

`func (o *CloudProviderRegion) SetDistance(v float32)`

SetDistance sets Distance field to given value.

### GetLocation

`func (o *CloudProviderRegion) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### SetLocation

`func (o *CloudProviderRegion) SetLocation(v string)`

SetLocation sets Location field to given value.

### GetName

`func (o *CloudProviderRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CloudProviderRegion) SetName(v string)`

SetName sets Name field to given value.

### GetProvider

`func (o *CloudProviderRegion) GetProvider() CloudProviderType`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### SetProvider

`func (o *CloudProviderRegion) SetProvider(v CloudProviderType)`

SetProvider sets Provider field to given value.

### GetServerless

`func (o *CloudProviderRegion) GetServerless() bool`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### SetServerless

`func (o *CloudProviderRegion) SetServerless(v bool)`

SetServerless sets Serverless field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


