# AddEgressRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description is text that serves to document the rules purpose. | 
**Destination** | **string** | destination is the endpoint (or subnetwork if CIDR) to which traffic is allowed. | 
**IdempotencyKey** | Pointer to **string** | idempotency_key uniquely identifies this request. If not set, it will be set by the server. | [optional] 
**Name** | **string** | name is the name of the egress rule. | 
**Paths** | Pointer to **[]string** | Deprecated: This field is ignored and will be removed in the next version. paths are the allowed URL paths. If empty, all paths are allowed. Only valid if Type&#x3D;\&quot;FQDN\&quot;. | [optional] 
**Ports** | Pointer to **[]int32** | ports are the allowed ports for TCP protocol. If Empty, all ports are allowed. | [optional] 
**Type** | **string** | type classifies the Destination field. Valid types include: \&quot;FQDN\&quot;, \&quot;CIDR\&quot;. | 

## Methods

### NewAddEgressRuleRequest

`func NewAddEgressRuleRequest(description string, destination string, name string, type_ string, ) *AddEgressRuleRequest`

NewAddEgressRuleRequest instantiates a new AddEgressRuleRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAddEgressRuleRequestWithDefaults

`func NewAddEgressRuleRequestWithDefaults() *AddEgressRuleRequest`

NewAddEgressRuleRequestWithDefaults instantiates a new AddEgressRuleRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *AddEgressRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *AddEgressRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetDestination

`func (o *AddEgressRuleRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### SetDestination

`func (o *AddEgressRuleRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### GetIdempotencyKey

`func (o *AddEgressRuleRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### SetIdempotencyKey

`func (o *AddEgressRuleRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### GetName

`func (o *AddEgressRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *AddEgressRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### GetPaths

`func (o *AddEgressRuleRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### SetPaths

`func (o *AddEgressRuleRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### GetPorts

`func (o *AddEgressRuleRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### SetPorts

`func (o *AddEgressRuleRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### GetType

`func (o *AddEgressRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *AddEgressRuleRequest) SetType(v string)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


