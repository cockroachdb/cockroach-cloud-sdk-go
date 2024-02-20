# AddPrivateEndpointTrustedOwnerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalOwnerId** | **string** | external_owner_id is the identifier of the owner within the cloud provider for private endpoint connections. A wildcard character (\&quot;*\&quot;) can be used to denote all owners. | 
**Type** | [**PrivateEndpointTrustedOwnerTypeType**](PrivateEndpointTrustedOwnerTypeType.md) |  | 

## Methods

### NewAddPrivateEndpointTrustedOwnerRequest

`func NewAddPrivateEndpointTrustedOwnerRequest(externalOwnerId string, type_ PrivateEndpointTrustedOwnerTypeType, ) *AddPrivateEndpointTrustedOwnerRequest`

NewAddPrivateEndpointTrustedOwnerRequest instantiates a new AddPrivateEndpointTrustedOwnerRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewAddPrivateEndpointTrustedOwnerRequestWithDefaults

`func NewAddPrivateEndpointTrustedOwnerRequestWithDefaults() *AddPrivateEndpointTrustedOwnerRequest`

NewAddPrivateEndpointTrustedOwnerRequestWithDefaults instantiates a new AddPrivateEndpointTrustedOwnerRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetExternalOwnerId

`func (o *AddPrivateEndpointTrustedOwnerRequest) GetExternalOwnerId() string`

GetExternalOwnerId returns the ExternalOwnerId field if non-nil, zero value otherwise.

### SetExternalOwnerId

`func (o *AddPrivateEndpointTrustedOwnerRequest) SetExternalOwnerId(v string)`

SetExternalOwnerId sets ExternalOwnerId field to given value.

### GetType

`func (o *AddPrivateEndpointTrustedOwnerRequest) GetType() PrivateEndpointTrustedOwnerTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *AddPrivateEndpointTrustedOwnerRequest) SetType(v PrivateEndpointTrustedOwnerTypeType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


