# ListEgressPrivateEndpointsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressPrivateEndpoints** | [**[]EgressPrivateEndpoint**](EgressPrivateEndpoint.md) | egress_private_endpoints are the egress private endpoints associated with the given cluster. | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListEgressPrivateEndpointsResponse

`func NewListEgressPrivateEndpointsResponse(egressPrivateEndpoints []EgressPrivateEndpoint, ) *ListEgressPrivateEndpointsResponse`

NewListEgressPrivateEndpointsResponse instantiates a new ListEgressPrivateEndpointsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListEgressPrivateEndpointsResponseWithDefaults

`func NewListEgressPrivateEndpointsResponseWithDefaults() *ListEgressPrivateEndpointsResponse`

NewListEgressPrivateEndpointsResponseWithDefaults instantiates a new ListEgressPrivateEndpointsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetEgressPrivateEndpoints

`func (o *ListEgressPrivateEndpointsResponse) GetEgressPrivateEndpoints() []EgressPrivateEndpoint`

GetEgressPrivateEndpoints returns the EgressPrivateEndpoints field if non-nil, zero value otherwise.

### SetEgressPrivateEndpoints

`func (o *ListEgressPrivateEndpointsResponse) SetEgressPrivateEndpoints(v []EgressPrivateEndpoint)`

SetEgressPrivateEndpoints sets EgressPrivateEndpoints field to given value.

### GetPagination

`func (o *ListEgressPrivateEndpointsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListEgressPrivateEndpointsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


