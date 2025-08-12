# ReplicationStreamStatusType

## Enum
> status describes the desired status of the replication stream.   - STARTING: signifies that a stream is starting  - REPLICATING: during an update request, status 'REPLICATING' is allowed to transition to 'FAILING_OVER' or 'CANCELED'. Other status transitions are not supported.  - FAILING_OVER: used to trigger a failover, or to signify that a failover is occurring.  - COMPLETED: signifies that a failover is complete.  - CANCELED: signifies that a stream is canceled.

* `STARTING` (value: `"STARTING"`)

* `REPLICATING` (value: `"REPLICATING"`)

* `FAILING_OVER` (value: `"FAILING_OVER"`)

* `COMPLETED` (value: `"COMPLETED"`)

* `CANCELED` (value: `"CANCELED"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


