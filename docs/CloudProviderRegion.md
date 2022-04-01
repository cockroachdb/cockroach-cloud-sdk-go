# CloudProviderRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Location** | **string** |  | 
**Provider** | [**CloudProvider**](CloudProvider.md) |  | [default to UNSPECIFIED]
**Serverless** | **bool** |  | 
**Distance** | **float32** |  | 

## Methods

### NewCloudProviderRegion

`func NewCloudProviderRegion(name string, location string, provider CloudProvider, serverless bool, distance float32, ) *CloudProviderRegion`

NewCloudProviderRegion instantiates a new CloudProviderRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudProviderRegionWithDefaults

`func NewCloudProviderRegionWithDefaults() *CloudProviderRegion`

NewCloudProviderRegionWithDefaults instantiates a new CloudProviderRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CloudProviderRegion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudProviderRegion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudProviderRegion) SetName(v string)`

SetName sets Name field to given value.


### GetLocation

`func (o *CloudProviderRegion) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CloudProviderRegion) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CloudProviderRegion) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetProvider

`func (o *CloudProviderRegion) GetProvider() CloudProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudProviderRegion) GetProviderOk() (*CloudProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudProviderRegion) SetProvider(v CloudProvider)`

SetProvider sets Provider field to given value.


### GetServerless

`func (o *CloudProviderRegion) GetServerless() bool`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *CloudProviderRegion) GetServerlessOk() (*bool, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *CloudProviderRegion) SetServerless(v bool)`

SetServerless sets Serverless field to given value.


### GetDistance

`func (o *CloudProviderRegion) GetDistance() float32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *CloudProviderRegion) GetDistanceOk() (*float32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *CloudProviderRegion) SetDistance(v float32)`

SetDistance sets Distance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


