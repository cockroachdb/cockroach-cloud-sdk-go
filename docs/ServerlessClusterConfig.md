# ServerlessClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendLimit** | **int32** | Spend limit in US cents. | 
**RoutingId** | **string** | Used to build a connection string. | 

## Methods

### NewServerlessClusterConfig

`func NewServerlessClusterConfig(spendLimit int32, routingId string, ) *ServerlessClusterConfig`

NewServerlessClusterConfig instantiates a new ServerlessClusterConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServerlessClusterConfigWithDefaults

`func NewServerlessClusterConfigWithDefaults() *ServerlessClusterConfig`

NewServerlessClusterConfigWithDefaults instantiates a new ServerlessClusterConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetSpendLimit

`func (o *ServerlessClusterConfig) GetSpendLimit() int32`

GetSpendLimit returns the SpendLimit field if non-nil, zero value otherwise.

### SetSpendLimit

`func (o *ServerlessClusterConfig) SetSpendLimit(v int32)`

SetSpendLimit sets SpendLimit field to given value.

### GetRoutingId

`func (o *ServerlessClusterConfig) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### SetRoutingId

`func (o *ServerlessClusterConfig) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


