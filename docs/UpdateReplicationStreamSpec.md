# UpdateReplicationStreamSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverAt** | Pointer to **time.Time** | failover_at is the crdb system time at which failover occurs. If the user sets the status to &#39;FAILING_OVER&#39; but omits failover_at, the failover time will default to the latest consistent replicated time. Otherwise, the user can pick a time up to one hour in the future to schedule a failover, or a time in the past to restore the cluster to a recent state. If the time is in the past, the API will make a best-effort attempt to validate that the time is not earlier than the retained time. In this case, if the retained time is updated in between validation and failover execution and the failover time becomes invalid, the stream will failover to the retained time. | [optional] 
**Status** | Pointer to [**ReplicationStreamStatusType**](ReplicationStreamStatusType.md) |  | [optional] 

## Methods

### NewUpdateReplicationStreamSpec

`func NewUpdateReplicationStreamSpec() *UpdateReplicationStreamSpec`

NewUpdateReplicationStreamSpec instantiates a new UpdateReplicationStreamSpec object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetFailoverAt

`func (o *UpdateReplicationStreamSpec) GetFailoverAt() time.Time`

GetFailoverAt returns the FailoverAt field if non-nil, zero value otherwise.

### SetFailoverAt

`func (o *UpdateReplicationStreamSpec) SetFailoverAt(v time.Time)`

SetFailoverAt sets FailoverAt field to given value.

### GetStatus

`func (o *UpdateReplicationStreamSpec) GetStatus() ReplicationStreamStatusType`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *UpdateReplicationStreamSpec) SetStatus(v ReplicationStreamStatusType)`

SetStatus sets Status field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


