# EgressPrivateEndpoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainNames** | **[]string** | domain_names are the domain names associated with the egress private endpoint. | 
**EndpointAddress** | **string** |  | 
**EndpointConnectionId** | **string** | endpoint_connection_id is the cloud-specific id for egress private endpoints. This connection ID is visible in CRL and CSP consoles as \&quot;VPC endpoint ID\&quot; and is used to uniquely identify the endpoint in external configurations. | 
**Id** | **string** | A generated ID that uniquely identifies the egress_private_endpoint for use with the CockroachDB Cloud API. This ID is generic and not specific to the cloud provider. | 
**Region** | **string** | region represents the cloud region that the endpoint was created in. | 
**State** | [**EgressPrivateEndpointStateType**](EgressPrivateEndpointStateType.md) |  | 
**TargetServiceIdentifier** | **string** | target_service_identifier represents the identifier of the target service. User-visible in CRL and CSP consoles as the \&quot;thing\&quot; the user is telling us to connect to, i.e. Service Name for AWS, Service Attachment for GCP, or ARN for AWS MSK. | 
**TargetServiceType** | [**EgressPrivateEndpointTargetServiceTypeType**](EgressPrivateEndpointTargetServiceTypeType.md) |  | 

## Methods

### NewEgressPrivateEndpoint

`func NewEgressPrivateEndpoint(domainNames []string, endpointAddress string, endpointConnectionId string, id string, region string, state EgressPrivateEndpointStateType, targetServiceIdentifier string, targetServiceType EgressPrivateEndpointTargetServiceTypeType, ) *EgressPrivateEndpoint`

NewEgressPrivateEndpoint instantiates a new EgressPrivateEndpoint object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEgressPrivateEndpointWithDefaults

`func NewEgressPrivateEndpointWithDefaults() *EgressPrivateEndpoint`

NewEgressPrivateEndpointWithDefaults instantiates a new EgressPrivateEndpoint object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDomainNames

`func (o *EgressPrivateEndpoint) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### SetDomainNames

`func (o *EgressPrivateEndpoint) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### GetEndpointAddress

`func (o *EgressPrivateEndpoint) GetEndpointAddress() string`

GetEndpointAddress returns the EndpointAddress field if non-nil, zero value otherwise.

### SetEndpointAddress

`func (o *EgressPrivateEndpoint) SetEndpointAddress(v string)`

SetEndpointAddress sets EndpointAddress field to given value.

### GetEndpointConnectionId

`func (o *EgressPrivateEndpoint) GetEndpointConnectionId() string`

GetEndpointConnectionId returns the EndpointConnectionId field if non-nil, zero value otherwise.

### SetEndpointConnectionId

`func (o *EgressPrivateEndpoint) SetEndpointConnectionId(v string)`

SetEndpointConnectionId sets EndpointConnectionId field to given value.

### GetId

`func (o *EgressPrivateEndpoint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *EgressPrivateEndpoint) SetId(v string)`

SetId sets Id field to given value.

### GetRegion

`func (o *EgressPrivateEndpoint) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *EgressPrivateEndpoint) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetState

`func (o *EgressPrivateEndpoint) GetState() EgressPrivateEndpointStateType`

GetState returns the State field if non-nil, zero value otherwise.

### SetState

`func (o *EgressPrivateEndpoint) SetState(v EgressPrivateEndpointStateType)`

SetState sets State field to given value.

### GetTargetServiceIdentifier

`func (o *EgressPrivateEndpoint) GetTargetServiceIdentifier() string`

GetTargetServiceIdentifier returns the TargetServiceIdentifier field if non-nil, zero value otherwise.

### SetTargetServiceIdentifier

`func (o *EgressPrivateEndpoint) SetTargetServiceIdentifier(v string)`

SetTargetServiceIdentifier sets TargetServiceIdentifier field to given value.

### GetTargetServiceType

`func (o *EgressPrivateEndpoint) GetTargetServiceType() EgressPrivateEndpointTargetServiceTypeType`

GetTargetServiceType returns the TargetServiceType field if non-nil, zero value otherwise.

### SetTargetServiceType

`func (o *EgressPrivateEndpoint) SetTargetServiceType(v EgressPrivateEndpointTargetServiceTypeType)`

SetTargetServiceType sets TargetServiceType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


