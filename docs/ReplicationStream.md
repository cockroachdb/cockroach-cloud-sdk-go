# ReplicationStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationAt** | Pointer to **time.Time** | activation_at is the crdb system time at which failover is finalized. This may differ from the time for which failover was requested. This field will be present when a replication stream is in the COMPLETED state. | [optional] 
**CreatedAt** | **time.Time** | created_at is the timestamp at which the replication stream was created. | 
**FailoverAt** | Pointer to **time.Time** | failover_at is the time for which failover is requested. If the user sets the status to &#39;FAILING_OVER&#39; but omits failover_at, the failover time will default to the latest consistent replicated time. Otherwise, the user can pick a time up to one hour in the future to schedule a failover, or a time in the past to restore the cluster to a recent state. This field will be present if the user has requested failover at a future time. | [optional] 
**Id** | **string** | id is the UUID of the replication stream. | 
**ReplicatedTime** | Pointer to **time.Time** | replicated_time is the timestamp indicating the point up to which data has been replicated. The window between replicated_time and the actual time is known as replication lag. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**ReplicationLagSeconds** | Pointer to **int32** | replication_lag_seconds is the replication lag (current time minus replicated time) in seconds. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**RetainedTime** | Pointer to **time.Time** | retained_time is the timestamp indicating the lower bound that the replication stream can failover to. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**SourceClusterId** | **string** | source_cluster_id is the ID of the cluster that is being replicated. | 
**Status** | [**ReplicationStreamStatusType**](ReplicationStreamStatusType.md) |  | 
**TargetClusterId** | **string** | target_cluster_id is the ID of the cluster that data is being replicated to. | 

## Methods

### NewReplicationStream

`func NewReplicationStream(createdAt time.Time, id string, sourceClusterId string, status ReplicationStreamStatusType, targetClusterId string, ) *ReplicationStream`

NewReplicationStream instantiates a new ReplicationStream object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewReplicationStreamWithDefaults

`func NewReplicationStreamWithDefaults() *ReplicationStream`

NewReplicationStreamWithDefaults instantiates a new ReplicationStream object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetActivationAt

`func (o *ReplicationStream) GetActivationAt() time.Time`

GetActivationAt returns the ActivationAt field if non-nil, zero value otherwise.

### SetActivationAt

`func (o *ReplicationStream) SetActivationAt(v time.Time)`

SetActivationAt sets ActivationAt field to given value.

### GetCreatedAt

`func (o *ReplicationStream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *ReplicationStream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetFailoverAt

`func (o *ReplicationStream) GetFailoverAt() time.Time`

GetFailoverAt returns the FailoverAt field if non-nil, zero value otherwise.

### SetFailoverAt

`func (o *ReplicationStream) SetFailoverAt(v time.Time)`

SetFailoverAt sets FailoverAt field to given value.

### GetId

`func (o *ReplicationStream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *ReplicationStream) SetId(v string)`

SetId sets Id field to given value.

### GetReplicatedTime

`func (o *ReplicationStream) GetReplicatedTime() time.Time`

GetReplicatedTime returns the ReplicatedTime field if non-nil, zero value otherwise.

### SetReplicatedTime

`func (o *ReplicationStream) SetReplicatedTime(v time.Time)`

SetReplicatedTime sets ReplicatedTime field to given value.

### GetReplicationLagSeconds

`func (o *ReplicationStream) GetReplicationLagSeconds() int32`

GetReplicationLagSeconds returns the ReplicationLagSeconds field if non-nil, zero value otherwise.

### SetReplicationLagSeconds

`func (o *ReplicationStream) SetReplicationLagSeconds(v int32)`

SetReplicationLagSeconds sets ReplicationLagSeconds field to given value.

### GetRetainedTime

`func (o *ReplicationStream) GetRetainedTime() time.Time`

GetRetainedTime returns the RetainedTime field if non-nil, zero value otherwise.

### SetRetainedTime

`func (o *ReplicationStream) SetRetainedTime(v time.Time)`

SetRetainedTime sets RetainedTime field to given value.

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


