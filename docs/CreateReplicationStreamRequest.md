# CreateReplicationStreamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceClusterId** | **string** | source_cluster_id is the ID of the cluster that is being replicated. | 
**TargetClusterId** | **string** | target_cluster_id is the ID of the cluster that data is being replicated to. | 

## Methods

### NewCreateReplicationStreamRequest

`func NewCreateReplicationStreamRequest(sourceClusterId string, targetClusterId string, ) *CreateReplicationStreamRequest`

NewCreateReplicationStreamRequest instantiates a new CreateReplicationStreamRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateReplicationStreamRequestWithDefaults

`func NewCreateReplicationStreamRequestWithDefaults() *CreateReplicationStreamRequest`

NewCreateReplicationStreamRequestWithDefaults instantiates a new CreateReplicationStreamRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetSourceClusterId

`func (o *CreateReplicationStreamRequest) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### SetSourceClusterId

`func (o *CreateReplicationStreamRequest) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### GetTargetClusterId

`func (o *CreateReplicationStreamRequest) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### SetTargetClusterId

`func (o *CreateReplicationStreamRequest) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


