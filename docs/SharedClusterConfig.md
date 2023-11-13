# SharedClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingId** | **string** | routing_id is used to identify the cluster in a connection string. | 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewSharedClusterConfig

`func NewSharedClusterConfig(routingId string, ) *SharedClusterConfig`

NewSharedClusterConfig instantiates a new SharedClusterConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSharedClusterConfigWithDefaults

`func NewSharedClusterConfigWithDefaults() *SharedClusterConfig`

NewSharedClusterConfigWithDefaults instantiates a new SharedClusterConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRoutingId

`func (o *SharedClusterConfig) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### SetRoutingId

`func (o *SharedClusterConfig) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### GetUsageLimits

`func (o *SharedClusterConfig) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *SharedClusterConfig) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


