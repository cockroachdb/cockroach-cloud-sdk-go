# CMEKStatus

## Enum
> `CMEKStatus describes the current status of CMEK for an entire CRDB cluster or a CMEK key within a region.   - DISABLED: DISABLED corresponds to the state of a cluster or region-level key when CMEK has finished being disabled. By default, CMEK will be disabled for new clusters.  - DISABLING: DISABLING corresponds to the state of a cluster or region-level key when CMEK is in the process of being disabled.  - DISABLE_FAILED: DISABLE_FAILED corresponds to the state of a cluster or region-level key when CMEK has failed to be disabled.  - ENABLED: ENABLED corresponds to the state of a cluster or region-level key when CMEK is enabled.  - ENABLING: ENABLING corresponds to the state of a cluster or region-level key when CMEK is in the process of being enabled.  - ENABLE_FAILED: ENABLE_FAILED corresponds to the state of a cluster or region-level key when CMEK has failed to be enabled.  - ROTATING: ROTATING corresponds to the state of a cluster or region when the a new spec is in the process of being enabled while an existing spec is being disabled.  - ROTATE_FAILED: ROTATE_FAILED corresponds to the state of a cluster or region if there was a failure to update from one CMEK spec to another.  - REVOKED: REVOKED corresponds to the state of a cluster or region-level key when the customer has revoked CockroachLab's permissions for their key.  - REVOKING: REVOKING corresponds to the state of a cluster or region-level key when CMEK is in the process of being revoked.  - REVOKE_FAILED: REVOKE_FAILED corresponds to the state of a cluster or region-level key when CMEK has failed to be revoked.`

* `DISABLED` (value: `"DISABLED"`)

* `DISABLING` (value: `"DISABLING"`)

* `DISABLE_FAILED` (value: `"DISABLE_FAILED"`)

* `ENABLED` (value: `"ENABLED"`)

* `ENABLING` (value: `"ENABLING"`)

* `ENABLE_FAILED` (value: `"ENABLE_FAILED"`)

* `ROTATING` (value: `"ROTATING"`)

* `ROTATE_FAILED` (value: `"ROTATE_FAILED"`)

* `REVOKED` (value: `"REVOKED"`)

* `REVOKING` (value: `"REVOKING"`)

* `REVOKE_FAILED` (value: `"REVOKE_FAILED"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


