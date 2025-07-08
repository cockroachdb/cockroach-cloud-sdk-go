# PhysicalReplicationStream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatedAt** | Pointer to **time.Time** | activated_at is the crdb system time at which failover is finalized. This may differ from the time for which failover was requested. This field will be present when a replication stream is in the COMPLETED state. | [optional] 
**CanceledAt** | Pointer to **time.Time** | canceled_at is the timestamp at which the replication stream was canceled. | [optional] 
**CreatedAt** | **time.Time** | created_at is the timestamp at which the replication stream was created. | 
**FailoverAt** | Pointer to **time.Time** | failover_at is the time for which failover is requested. If the user sets the status to &#39;FAILING_OVER&#39; but omits failover_at, the failover time will default to the latest consistent replicated time. Otherwise, the user can pick a time up to one hour in the future to schedule a failover, or a time in the past to restore the cluster to a recent state. This field will be present if the user has requested failover at a future time. | [optional] 
**Id** | **string** | id is the UUID of the replication stream. | 
**PrimaryClusterId** | **string** | primary_cluster_id is the ID of the cluster that is being replicated. | 
**ReplicatedTime** | Pointer to **time.Time** | replicated_time is the timestamp indicating the point up to which data has been replicated. The window between replicated_time and the actual time is known as replication lag. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**ReplicationLagSeconds** | Pointer to **int32** | replication_lag_seconds is the replication lag (current time minus replicated time) in seconds. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**RetainedTime** | Pointer to **time.Time** | retained_time is the timestamp indicating the lower bound that the replication stream can failover to. This field will be present when a replication stream is in the REPLICATING state. | [optional] 
**StandbyClusterId** | **string** | standby_cluster_id is the ID of the standby cluster that data is being replicated to. | 
**Status** | [**ReplicationStreamStatusType**](ReplicationStreamStatusType.md) |  | 

## Methods

### NewPhysicalReplicationStream

`func NewPhysicalReplicationStream(createdAt time.Time, id string, primaryClusterId string, standbyClusterId string, status ReplicationStreamStatusType, ) *PhysicalReplicationStream`

NewPhysicalReplicationStream instantiates a new PhysicalReplicationStream object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPhysicalReplicationStreamWithDefaults

`func NewPhysicalReplicationStreamWithDefaults() *PhysicalReplicationStream`

NewPhysicalReplicationStreamWithDefaults instantiates a new PhysicalReplicationStream object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetActivatedAt

`func (o *PhysicalReplicationStream) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### SetActivatedAt

`func (o *PhysicalReplicationStream) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### GetCanceledAt

`func (o *PhysicalReplicationStream) GetCanceledAt() time.Time`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### SetCanceledAt

`func (o *PhysicalReplicationStream) SetCanceledAt(v time.Time)`

SetCanceledAt sets CanceledAt field to given value.

### GetCreatedAt

`func (o *PhysicalReplicationStream) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *PhysicalReplicationStream) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetFailoverAt

`func (o *PhysicalReplicationStream) GetFailoverAt() time.Time`

GetFailoverAt returns the FailoverAt field if non-nil, zero value otherwise.

### SetFailoverAt

`func (o *PhysicalReplicationStream) SetFailoverAt(v time.Time)`

SetFailoverAt sets FailoverAt field to given value.

### GetId

`func (o *PhysicalReplicationStream) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *PhysicalReplicationStream) SetId(v string)`

SetId sets Id field to given value.

### GetPrimaryClusterId

`func (o *PhysicalReplicationStream) GetPrimaryClusterId() string`

GetPrimaryClusterId returns the PrimaryClusterId field if non-nil, zero value otherwise.

### SetPrimaryClusterId

`func (o *PhysicalReplicationStream) SetPrimaryClusterId(v string)`

SetPrimaryClusterId sets PrimaryClusterId field to given value.

### GetReplicatedTime

`func (o *PhysicalReplicationStream) GetReplicatedTime() time.Time`

GetReplicatedTime returns the ReplicatedTime field if non-nil, zero value otherwise.

### SetReplicatedTime

`func (o *PhysicalReplicationStream) SetReplicatedTime(v time.Time)`

SetReplicatedTime sets ReplicatedTime field to given value.

### GetReplicationLagSeconds

`func (o *PhysicalReplicationStream) GetReplicationLagSeconds() int32`

GetReplicationLagSeconds returns the ReplicationLagSeconds field if non-nil, zero value otherwise.

### SetReplicationLagSeconds

`func (o *PhysicalReplicationStream) SetReplicationLagSeconds(v int32)`

SetReplicationLagSeconds sets ReplicationLagSeconds field to given value.

### GetRetainedTime

`func (o *PhysicalReplicationStream) GetRetainedTime() time.Time`

GetRetainedTime returns the RetainedTime field if non-nil, zero value otherwise.

### SetRetainedTime

`func (o *PhysicalReplicationStream) SetRetainedTime(v time.Time)`

SetRetainedTime sets RetainedTime field to given value.

### GetStandbyClusterId

`func (o *PhysicalReplicationStream) GetStandbyClusterId() string`

GetStandbyClusterId returns the StandbyClusterId field if non-nil, zero value otherwise.

### SetStandbyClusterId

`func (o *PhysicalReplicationStream) SetStandbyClusterId(v string)`

SetStandbyClusterId sets StandbyClusterId field to given value.

### GetStatus

`func (o *PhysicalReplicationStream) GetStatus() ReplicationStreamStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *PhysicalReplicationStream) SetStatus(v ReplicationStreamStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


