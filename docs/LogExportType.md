# LogExportType

## Enum
> LogExportType encodes the cloud selection that we're exporting to along with the cloud logging platform.  Currently, each cloud has a single logging platform.   - AZURE_LOG_ANALYTICS: AZURE_LOG_ANALYTICS is the legacy Azure export path via Fluent Bit's HTTP Data Collector API (retiring 2026-09-14). Deprecated: use AZURE_LOG_ANALYTICS_V2 instead.  - AZURE_LOG_ANALYTICS_V2: AZURE_LOG_ANALYTICS_V2 exports to Azure Monitor via the Logs Ingestion API and DCR-based ingestion (OTel pipeline), replacing the deprecated AZURE_LOG_ANALYTICS path that uses Fluent Bit's HTTP Data Collector API (retiring 2026-09-14). V2 authenticates with a tenant/client/secret + DCR endpoint and rule, whereas the legacy type uses a workspace ID + shared key.

* `AWS_CLOUDWATCH` (value: `"AWS_CLOUDWATCH"`)

* `GCP_CLOUD_LOGGING` (value: `"GCP_CLOUD_LOGGING"`)

* `AZURE_LOG_ANALYTICS` (value: `"AZURE_LOG_ANALYTICS"`)

* `AZURE_LOG_ANALYTICS_V2` (value: `"AZURE_LOG_ANALYTICS_V2"`)


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


