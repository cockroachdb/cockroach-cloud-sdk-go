# CreatePhysicalReplicationStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryClusterId** | **string** | primary_cluster_id is the ID of the cluster that is being replicated. | 
**StandbyClusterId** | **string** | standby_cluster_id is the ID of the cluster that data is being replicated to. | 

## Methods

### NewCreatePhysicalReplicationStreamRequest

`func NewCreatePhysicalReplicationStreamRequest(primaryClusterId string, standbyClusterId string, ) *CreatePhysicalReplicationStreamRequest`

NewCreatePhysicalReplicationStreamRequest instantiates a new CreatePhysicalReplicationStreamRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreatePhysicalReplicationStreamRequestWithDefaults

`func NewCreatePhysicalReplicationStreamRequestWithDefaults() *CreatePhysicalReplicationStreamRequest`

NewCreatePhysicalReplicationStreamRequestWithDefaults instantiates a new CreatePhysicalReplicationStreamRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPrimaryClusterId

`func (o *CreatePhysicalReplicationStreamRequest) GetPrimaryClusterId() string`

GetPrimaryClusterId returns the PrimaryClusterId field if non-nil, zero value otherwise.

### SetPrimaryClusterId

`func (o *CreatePhysicalReplicationStreamRequest) SetPrimaryClusterId(v string)`

SetPrimaryClusterId sets PrimaryClusterId field to given value.

### GetStandbyClusterId

`func (o *CreatePhysicalReplicationStreamRequest) GetStandbyClusterId() string`

GetStandbyClusterId returns the StandbyClusterId field if non-nil, zero value otherwise.

### SetStandbyClusterId

`func (o *CreatePhysicalReplicationStreamRequest) SetStandbyClusterId(v string)`

SetStandbyClusterId sets StandbyClusterId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


