# ListPrivateEndpointTrustedOwnersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrustedOwners** | Pointer to [**[]PrivateEndpointTrustedOwner**](PrivateEndpointTrustedOwner.md) | trusted_owners describes the private endpoint trusted owner entries for the requested cluster. | [optional] 

## Methods

### NewListPrivateEndpointTrustedOwnersResponse

`func NewListPrivateEndpointTrustedOwnersResponse() *ListPrivateEndpointTrustedOwnersResponse`

NewListPrivateEndpointTrustedOwnersResponse instantiates a new ListPrivateEndpointTrustedOwnersResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetTrustedOwners

`func (o *ListPrivateEndpointTrustedOwnersResponse) GetTrustedOwners() []PrivateEndpointTrustedOwner`

GetTrustedOwners returns the TrustedOwners field if non-nil, zero value otherwise.

### SetTrustedOwners

`func (o *ListPrivateEndpointTrustedOwnersResponse) SetTrustedOwners(v []PrivateEndpointTrustedOwner)`

SetTrustedOwners sets TrustedOwners field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


