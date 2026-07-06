# SetEgressTrafficPolicyBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAll** | **bool** | allow_all, if true results in unrestricted egress traffic. If false, egress traffic is set to default-deny and is managed via the Egress Rule Management API. | 
**IdempotencyKey** | Pointer to **string** | idempotency_key uniquely identifies this request. If not set, it will be set by the server. | [optional] 

## Methods

### NewSetEgressTrafficPolicyBody

`func NewSetEgressTrafficPolicyBody(allowAll bool, ) *SetEgressTrafficPolicyBody`

NewSetEgressTrafficPolicyBody instantiates a new SetEgressTrafficPolicyBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewSetEgressTrafficPolicyBodyWithDefaults

`func NewSetEgressTrafficPolicyBodyWithDefaults() *SetEgressTrafficPolicyBody`

NewSetEgressTrafficPolicyBodyWithDefaults instantiates a new SetEgressTrafficPolicyBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAllowAll

`func (o *SetEgressTrafficPolicyBody) GetAllowAll() bool`

GetAllowAll returns the AllowAll field if non-nil, zero value otherwise.

### SetAllowAll

`func (o *SetEgressTrafficPolicyBody) SetAllowAll(v bool)`

SetAllowAll sets AllowAll field to given value.

### GetIdempotencyKey

`func (o *SetEgressTrafficPolicyBody) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### SetIdempotencyKey

`func (o *SetEgressTrafficPolicyBody) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


