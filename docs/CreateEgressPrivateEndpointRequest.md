# CreateEgressPrivateEndpointRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | **string** | region represents the region that the endpoint will be created in. | 
**TargetServiceIdentifier** | **string** | target_service_identifier represents the identifier of the target service. User-visible in CRL and CSP consoles as the \&quot;thing\&quot; the user is telling us to connect to, i.e. Service Name for AWS, Service Attachment for GCP, or ARN for AWS MSK. | 
**TargetServiceType** | [**EgressPrivateEndpointTargetServiceTypeType**](EgressPrivateEndpointTargetServiceTypeType.md) |  | 

## Methods

### NewCreateEgressPrivateEndpointRequest

`func NewCreateEgressPrivateEndpointRequest(region string, targetServiceIdentifier string, targetServiceType EgressPrivateEndpointTargetServiceTypeType, ) *CreateEgressPrivateEndpointRequest`

NewCreateEgressPrivateEndpointRequest instantiates a new CreateEgressPrivateEndpointRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateEgressPrivateEndpointRequestWithDefaults

`func NewCreateEgressPrivateEndpointRequestWithDefaults() *CreateEgressPrivateEndpointRequest`

NewCreateEgressPrivateEndpointRequestWithDefaults instantiates a new CreateEgressPrivateEndpointRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetRegion

`func (o *CreateEgressPrivateEndpointRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *CreateEgressPrivateEndpointRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetTargetServiceIdentifier

`func (o *CreateEgressPrivateEndpointRequest) GetTargetServiceIdentifier() string`

GetTargetServiceIdentifier returns the TargetServiceIdentifier field if non-nil, zero value otherwise.

### SetTargetServiceIdentifier

`func (o *CreateEgressPrivateEndpointRequest) SetTargetServiceIdentifier(v string)`

SetTargetServiceIdentifier sets TargetServiceIdentifier field to given value.

### GetTargetServiceType

`func (o *CreateEgressPrivateEndpointRequest) GetTargetServiceType() EgressPrivateEndpointTargetServiceTypeType`

GetTargetServiceType returns the TargetServiceType field if non-nil, zero value otherwise.

### SetTargetServiceType

`func (o *CreateEgressPrivateEndpointRequest) SetTargetServiceType(v EgressPrivateEndpointTargetServiceTypeType)`

SetTargetServiceType sets TargetServiceType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


