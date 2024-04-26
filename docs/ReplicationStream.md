# ReplicationStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | created_at is the timestamp at which the replication stream was created. | 
**Cutover** | **bool** | cutover describes whether or not cutover has been triggered. | 
**CutoverAt** | **time.Time** | cutover_at is the timestamp at which cutover occurs. If the user sets &#x60;cutover&#x60; to &#x60;true&#x60; but omits cutover_at, the cutover time will default to the latest consistent replicated time. Otherwise, the user can pick a time in the future to schedule a cutover in the future, or a time in the past to restore the cluster to a recent state. | 
**Id** | **string** | id is the UUID of the replication stream. | 
**SourceClusterId** | **string** | source_cluster_id is the ID of the cluster that is being replicated. | 
**Status** | [**ReplicationStreamStatusType**](ReplicationStreamStatusType.md) |  | 
**TargetClusterId** | **string** | target_cluster_id is the ID of the cluster that data is being replicated to. | 

## Methods

### NewReplicationStream

`func NewReplicationStream(createdAt time.Time, cutover bool, cutoverAt time.Time, id string, sourceClusterId string, status ReplicationStreamStatusType, targetClusterId string, ) *ReplicationStream`

NewReplicationStream instantiates a new ReplicationStream object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewReplicationStreamWithDefaults

`func NewReplicationStreamWithDefaults() *ReplicationStream`

NewReplicationStreamWithDefaults instantiates a new ReplicationStream object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *ReplicationStream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *ReplicationStream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetCutover

`func (o *ReplicationStream) GetCutover() bool`

GetCutover returns the Cutover field if non-nil, zero value otherwise.

### SetCutover

`func (o *ReplicationStream) SetCutover(v bool)`

SetCutover sets Cutover field to given value.

### GetCutoverAt

`func (o *ReplicationStream) GetCutoverAt() time.Time`

GetCutoverAt returns the CutoverAt field if non-nil, zero value otherwise.

### SetCutoverAt

`func (o *ReplicationStream) SetCutoverAt(v time.Time)`

SetCutoverAt sets CutoverAt field to given value.

### GetId

`func (o *ReplicationStream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ReplicationStream) SetId(v string)`

SetId sets Id field to given value.

### GetSourceClusterId

`func (o *ReplicationStream) GetSourceClusterId() string`

GetSourceClusterId returns the SourceClusterId field if non-nil, zero value otherwise.

### SetSourceClusterId

`func (o *ReplicationStream) SetSourceClusterId(v string)`

SetSourceClusterId sets SourceClusterId field to given value.

### GetStatus

`func (o *ReplicationStream) GetStatus() ReplicationStreamStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *ReplicationStream) SetStatus(v ReplicationStreamStatusType)`

SetStatus sets Status field to given value.

### GetTargetClusterId

`func (o *ReplicationStream) GetTargetClusterId() string`

GetTargetClusterId returns the TargetClusterId field if non-nil, zero value otherwise.

### SetTargetClusterId

`func (o *ReplicationStream) SetTargetClusterId(v string)`

SetTargetClusterId sets TargetClusterId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


