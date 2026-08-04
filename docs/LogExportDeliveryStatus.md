# LogExportDeliveryStatus

## Enum
> LogExportDeliveryStatus reports whether logs are reaching the customer's destination. This is independent of LogExportStatus, which reports whether the integration is configured and running. Delivery health is currently reported only for OTLP_HTTP sinks.   - DELIVERY_HEALTHY: The destination is receiving log records normally.  - DELIVERY_UNHEALTHY: The destination is rejecting or failing to receive a substantial fraction of log records. The cause is intentionally not distinguished — it may be invalid credentials, a wrong endpoint, or a transient outage.

* `HEALTHY` (value: `"DELIVERY_HEALTHY"`)

* `UNHEALTHY` (value: `"DELIVERY_UNHEALTHY"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


