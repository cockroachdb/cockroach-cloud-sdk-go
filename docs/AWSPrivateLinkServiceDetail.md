# AWSPrivateLinkServiceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneIds** | **[]string** | availability_zone_ids are the identifiers for the availability zones that the service is available in. | 
**ServiceId** | **string** | service_id is the server side of the PrivateLink connection. This is the same as AWSPrivateLinkEndpoint.service_id. | 
**ServiceName** | **string** | service_name is the AWS service name customers use to create endpoints on their end. | 

## Methods

### NewAWSPrivateLinkServiceDetail

`func NewAWSPrivateLinkServiceDetail(availabilityZoneIds []string, serviceId string, serviceName string, ) *AWSPrivateLinkServiceDetail`

NewAWSPrivateLinkServiceDetail instantiates a new AWSPrivateLinkServiceDetail object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAWSPrivateLinkServiceDetailWithDefaults

`func NewAWSPrivateLinkServiceDetailWithDefaults() *AWSPrivateLinkServiceDetail`

NewAWSPrivateLinkServiceDetailWithDefaults instantiates a new AWSPrivateLinkServiceDetail object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAvailabilityZoneIds

`func (o *AWSPrivateLinkServiceDetail) GetAvailabilityZoneIds() []string`

GetAvailabilityZoneIds returns the AvailabilityZoneIds field if non-nil, zero value otherwise.

### SetAvailabilityZoneIds

`func (o *AWSPrivateLinkServiceDetail) SetAvailabilityZoneIds(v []string)`

SetAvailabilityZoneIds sets AvailabilityZoneIds field to given value.

### GetServiceId

`func (o *AWSPrivateLinkServiceDetail) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### SetServiceId

`func (o *AWSPrivateLinkServiceDetail) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### GetServiceName

`func (o *AWSPrivateLinkServiceDetail) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### SetServiceName

`func (o *AWSPrivateLinkServiceDetail) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


