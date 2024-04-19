# EgressRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | cluster_id identifies the cluster to which this egress rule applies. | 
**CreatedAt** | Pointer to **time.Time** | created_at is the time at which the time at which the egress rule was created. | [optional] 
**CrlManaged** | **bool** | crl_managed indicates this egress rule is managed by CockroachDB Cloud services. This field is set by the server. | 
**Description** | **string** | description is a longer that serves to document the rules purpose. | 
**Destination** | **string** | destination is the endpoint (or subnetwork if CIDR) to which traffic is allowed. | 
**Id** | **string** | id uniquely identifies this egress rule. | 
**Name** | **string** | name is the name of the egress rule. | 
**Paths** | Pointer to **[]string** | Deprecated: This field is ignored and will be removed in the next version. paths are the allowed URL paths. Only valid if Type&#x3D;\&quot;FQDN\&quot;. | [optional] 
**Ports** | Pointer to **[]int32** | ports are the allowed ports for TCP protocol. If Empty, all ports are allowed. | [optional] 
**State** | **string** | state indicates the state of the egress rule. | 
**Type** | **string** | type classifies the destination field. Valid types include: \&quot;FQDN\&quot;, \&quot;CIDR\&quot;. | 

## Methods

### NewEgressRule

`func NewEgressRule(clusterId string, crlManaged bool, description string, destination string, id string, name string, state string, type_ string, ) *EgressRule`

NewEgressRule instantiates a new EgressRule object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEgressRuleWithDefaults

`func NewEgressRuleWithDefaults() *EgressRule`

NewEgressRuleWithDefaults instantiates a new EgressRule object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *EgressRule) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *EgressRule) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetCreatedAt

`func (o *EgressRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *EgressRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetCrlManaged

`func (o *EgressRule) GetCrlManaged() bool`

GetCrlManaged returns the CrlManaged field if non-nil, zero value otherwise.

### SetCrlManaged

`func (o *EgressRule) SetCrlManaged(v bool)`

SetCrlManaged sets CrlManaged field to given value.

### GetDescription

`func (o *EgressRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *EgressRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetDestination

`func (o *EgressRule) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### SetDestination

`func (o *EgressRule) SetDestination(v string)`

SetDestination sets Destination field to given value.

### GetId

`func (o *EgressRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *EgressRule) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *EgressRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *EgressRule) SetName(v string)`

SetName sets Name field to given value.

### GetPaths

`func (o *EgressRule) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### SetPaths

`func (o *EgressRule) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### GetPorts

`func (o *EgressRule) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### SetPorts

`func (o *EgressRule) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### GetState

`func (o *EgressRule) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### SetState

`func (o *EgressRule) SetState(v string)`

SetState sets State field to given value.

### GetType

`func (o *EgressRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *EgressRule) SetType(v string)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


