# ServerlessClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoutingId** | **string** | routing_id is used to identify the cluster in a connection string. | 
**UpgradeType** | [**UpgradeTypeType**](UpgradeTypeType.md) |  | 
**UsageLimits** | Pointer to [**UsageLimits**](UsageLimits.md) |  | [optional] 

## Methods

### NewServerlessClusterConfig

`func NewServerlessClusterConfig(routingId string, upgradeType UpgradeTypeType, ) *ServerlessClusterConfig`

NewServerlessClusterConfig instantiates a new ServerlessClusterConfig object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewServerlessClusterConfigWithDefaults

`func NewServerlessClusterConfigWithDefaults() *ServerlessClusterConfig`

NewServerlessClusterConfigWithDefaults instantiates a new ServerlessClusterConfig object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRoutingId

`func (o *ServerlessClusterConfig) GetRoutingId() string`

GetRoutingId returns the RoutingId field if non-nil, zero value otherwise.

### SetRoutingId

`func (o *ServerlessClusterConfig) SetRoutingId(v string)`

SetRoutingId sets RoutingId field to given value.

### GetUpgradeType

`func (o *ServerlessClusterConfig) GetUpgradeType() UpgradeTypeType`

GetUpgradeType returns the UpgradeType field if non-nil, zero value otherwise.

### SetUpgradeType

`func (o *ServerlessClusterConfig) SetUpgradeType(v UpgradeTypeType)`

SetUpgradeType sets UpgradeType field to given value.

### GetUsageLimits

`func (o *ServerlessClusterConfig) GetUsageLimits() UsageLimits`

GetUsageLimits returns the UsageLimits field if non-nil, zero value otherwise.

### SetUsageLimits

`func (o *ServerlessClusterConfig) SetUsageLimits(v UsageLimits)`

SetUsageLimits sets UsageLimits field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


