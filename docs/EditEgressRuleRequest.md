# EditEgressRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description is text that serves to document the rules purpose. | [optional] 
**Destination** | Pointer to **string** | destination is a CIDR range or fully-qualified domain name to which outgoing traffic should be allowed. This field is required. | [optional] 
**IdempotencyKey** | Pointer to **string** | idempotency_key uniquely identifies this request. If not set, it will be set by the server. | [optional] 
**Paths** | Pointer to **[]string** | paths are the allowed URL paths. If empty, all paths are allowed. Only valid if Type&#x3D;\&quot;FQDN\&quot;. | [optional] 
**Ports** | Pointer to **[]int32** | ports are the allowed ports for TCP protocol. If empty, all ports are allowed. | [optional] 
**Type** | Pointer to **string** | type is the destination type of this rule. Example values are FQDN or CIDR. This field is required. | [optional] 

## Methods

### NewEditEgressRuleRequest

`func NewEditEgressRuleRequest() *EditEgressRuleRequest`

NewEditEgressRuleRequest instantiates a new EditEgressRuleRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetDescription

`func (o *EditEgressRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *EditEgressRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetDestination

`func (o *EditEgressRuleRequest) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### SetDestination

`func (o *EditEgressRuleRequest) SetDestination(v string)`

SetDestination sets Destination field to given value.

### GetIdempotencyKey

`func (o *EditEgressRuleRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### SetIdempotencyKey

`func (o *EditEgressRuleRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.

### GetPaths

`func (o *EditEgressRuleRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### SetPaths

`func (o *EditEgressRuleRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### GetPorts

`func (o *EditEgressRuleRequest) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### SetPorts

`func (o *EditEgressRuleRequest) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### GetType

`func (o *EditEgressRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *EditEgressRuleRequest) SetType(v string)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


