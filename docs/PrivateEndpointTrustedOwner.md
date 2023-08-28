# PrivateEndpointTrustedOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterId** | **string** | cluster_id identifies the cluster to which this trusted owner entry applies. | 
**CreatedAt** | **time.Time** | created_at is the time at which the entry was created. | 
**ExternalOwnerId** | **string** | external_owner_id is the identifier of the owner within the cloud provider for private endpoint connections. A wildcard character (\&quot;*\&quot;) can be used to denote all owners. | 
**Id** | **string** | id is a UUID that uniquely identifies this trusted owner entry. | 
**Type** | [**PrivateEndpointTrustedOwnerTypeType**](PrivateEndpointTrustedOwnerTypeType.md) |  | 

## Methods

### NewPrivateEndpointTrustedOwner

`func NewPrivateEndpointTrustedOwner(clusterId string, createdAt time.Time, externalOwnerId string, id string, type_ PrivateEndpointTrustedOwnerTypeType, ) *PrivateEndpointTrustedOwner`

NewPrivateEndpointTrustedOwner instantiates a new PrivateEndpointTrustedOwner object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPrivateEndpointTrustedOwnerWithDefaults

`func NewPrivateEndpointTrustedOwnerWithDefaults() *PrivateEndpointTrustedOwner`

NewPrivateEndpointTrustedOwnerWithDefaults instantiates a new PrivateEndpointTrustedOwner object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetClusterId

`func (o *PrivateEndpointTrustedOwner) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *PrivateEndpointTrustedOwner) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetCreatedAt

`func (o *PrivateEndpointTrustedOwner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *PrivateEndpointTrustedOwner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetExternalOwnerId

`func (o *PrivateEndpointTrustedOwner) GetExternalOwnerId() string`

GetExternalOwnerId returns the ExternalOwnerId field if non-nil, zero value otherwise.

### SetExternalOwnerId

`func (o *PrivateEndpointTrustedOwner) SetExternalOwnerId(v string)`

SetExternalOwnerId sets ExternalOwnerId field to given value.

### GetId

`func (o *PrivateEndpointTrustedOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *PrivateEndpointTrustedOwner) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *PrivateEndpointTrustedOwner) GetType() PrivateEndpointTrustedOwnerTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *PrivateEndpointTrustedOwner) SetType(v PrivateEndpointTrustedOwnerTypeType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


