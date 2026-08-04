# LogExportType

## Enum
> LogExportType identifies the destination used for exported logs. Cloud-native destinations use their provider's logging service, while OTLP_HTTP sends logs to a customer-configured OTLP/HTTP endpoint.   - AZURE_LOG_ANALYTICS: AZURE_LOG_ANALYTICS is the legacy Azure export path via the HTTP Data Collector API (retiring 2026-09-14). Deprecated: use AZURE_LOG_ANALYTICS_V2 instead.  - AZURE_LOG_ANALYTICS_V2: AZURE_LOG_ANALYTICS_V2 exports to Azure Monitor via the Logs Ingestion API and DCR-based ingestion, replacing the deprecated AZURE_LOG_ANALYTICS path that uses the HTTP Data Collector API (retiring 2026-09-14). V2 authenticates with a tenant/client/secret + DCR endpoint and rule, whereas the legacy type uses a workspace ID + shared key.

* `AWS_CLOUDWATCH` (value: `"AWS_CLOUDWATCH"`)

* `GCP_CLOUD_LOGGING` (value: `"GCP_CLOUD_LOGGING"`)

* `AZURE_LOG_ANALYTICS` (value: `"AZURE_LOG_ANALYTICS"`)

* `AZURE_LOG_ANALYTICS_V2` (value: `"AZURE_LOG_ANALYTICS_V2"`)

* `OTLP_HTTP` (value: `"OTLP_HTTP"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


