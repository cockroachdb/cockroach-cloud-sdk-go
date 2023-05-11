# ClusterUpgradeStatusType

## Enum
> ` - FINALIZED: The cluster is running the latest available CockroachDB version, and all upgrades have been finalized.  - MAJOR_UPGRADE_RUNNING: An major version upgrade is currently in progress.  - UPGRADE_AVAILABLE: An upgrade is available. If preview builds are enabled for the parent organization, this could indicate that a preview upgrade is available.  - PENDING_FINALIZATION: An upgrade is complete, but pending finalization. Upgrades are automatically finalized after 72 hours. For more information, see https://www.cockroachlabs.com/docs/stable/upgrade-cockroach-version.html  - ROLLBACK_RUNNING: A rollback operation is currently in progress.`

* `FINALIZED` (value: `"FINALIZED"`)

* `MAJOR_UPGRADE_RUNNING` (value: `"MAJOR_UPGRADE_RUNNING"`)

* `UPGRADE_AVAILABLE` (value: `"UPGRADE_AVAILABLE"`)

* `PENDING_FINALIZATION` (value: `"PENDING_FINALIZATION"`)

* `ROLLBACK_RUNNING` (value: `"ROLLBACK_RUNNING"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


