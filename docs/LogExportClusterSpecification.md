# LogExportClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPrincipal** | Pointer to **string** | auth_principal is used in different contexts based on integration. CloudWatch: AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch GCP Cloud Logging: GCP Project ID that the cluster service account has permissions to write to for cloud logging. Azure Log Analytics (AZURE_LOG_ANALYTICS): CustomerID or WorkspaceID. This field is empty for AZURE_LOG_ANALYTICS_V2 and OTLP_HTTP. OTLP_HTTP authenticates using the headers supplied when the integration is enabled. | [optional] 
**AwsExternalId** | Pointer to **string** | aws_external_id, if set, is included when assuming the IAM role. Supported for Advanced clusters on AWS only. | [optional] 
**AzureClientId** | Pointer to **string** | Azure client ID for the app registration used by the Logs Ingestion API. | [optional] 
**AzureClientSecret** | Pointer to **string** | Azure client secret for the app registration used by the Logs Ingestion API. | [optional] 
**AzureDceEndpoint** | Pointer to **string** | Logs ingestion endpoint of the Azure Data Collection Endpoint (DCE). For example, https://{dce-name}.{region}.ingest.monitor.azure.com. | [optional] 
**AzureDcrImmutableId** | Pointer to **string** | Immutable ID of the Azure Data Collection Rule (DCR), for example dcr-... | [optional] 
**AzureDcrResourceId** | Pointer to **string** | Full ARM resource ID of the Azure Data Collection Rule (DCR). Cockroach Cloud reads and updates this DCR to add streams for each configured log group. For example, /subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.Insights/dataCollectionRules/{dcr-name}. | [optional] 
**AzureSharedKey** | Pointer to **string** | The primary or the secondary connected sources client authentication key. This is used to export logs to Azure Log Analytics via the legacy HTTP Data Collector API. Deprecated: use azure_client_secret instead. | [optional] 
**AzureTenantId** | Pointer to **string** | Azure tenant ID for the app registration used by the Logs Ingestion API. | [optional] 
**AzureWorkspaceResourceId** | Pointer to **string** | Full ARM resource ID of the Log Analytics workspace. Cockroach Cloud creates or updates the required custom tables for each configured log group. For example, /subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.OperationalInsights/workspaces/{workspace-name}. | [optional] 
**Groups** | Pointer to [**[]LogExportGroup**](LogExportGroup.md) | groups is a collection of log group configurations to customize which CRDB channels get aggregated into different groups at the target sink. Unconfigured channels will be sent to the default locations using the top-level log export settings. | [optional] 
**LogName** | Pointer to **string** | log_name is an identifier for the logs in the customer&#39;s log sink. For AZURE_LOG_ANALYTICS_V2, it must start with a letter and contain only letters, digits, and underscores. | [optional] 
**OmittedChannels** | Pointer to **[]string** | omitted_channels is a list of channels that the user does not want to export logs for. | [optional] 
**OtlpEndpoint** | Pointer to **string** | otlp_endpoint is the OTLP/HTTP URL for the OTLP_HTTP sink type. This field accepts either a base endpoint or a full /v1/logs endpoint. | [optional] 
**OtlpHeaderNames** | Pointer to **[]string** | otlp_header_names lists the configured OTLP auth header names; values are never returned. | [optional] 
**Redact** | Pointer to **bool** | redact controls whether logs are redacted before forwarding to customer sinks. By default they are not redacted. | [optional] 
**Region** | Pointer to **string** | region controls whether all logs are sent to a specific region in the customer sink. By default, logs will remain their region of origin depending on the cluster node&#39;s region. | [optional] 
**Type** | Pointer to [**LogExportType**](LogExportType.md) |  | [optional] 

## Methods

### NewLogExportClusterSpecification

`func NewLogExportClusterSpecification() *LogExportClusterSpecification`

NewLogExportClusterSpecification instantiates a new LogExportClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAuthPrincipal

`func (o *LogExportClusterSpecification) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### SetAuthPrincipal

`func (o *LogExportClusterSpecification) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.

### GetAwsExternalId

`func (o *LogExportClusterSpecification) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### SetAwsExternalId

`func (o *LogExportClusterSpecification) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### GetAzureClientId

`func (o *LogExportClusterSpecification) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### SetAzureClientId

`func (o *LogExportClusterSpecification) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### GetAzureClientSecret

`func (o *LogExportClusterSpecification) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### SetAzureClientSecret

`func (o *LogExportClusterSpecification) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### GetAzureDceEndpoint

`func (o *LogExportClusterSpecification) GetAzureDceEndpoint() string`

GetAzureDceEndpoint returns the AzureDceEndpoint field if non-nil, zero value otherwise.

### SetAzureDceEndpoint

`func (o *LogExportClusterSpecification) SetAzureDceEndpoint(v string)`

SetAzureDceEndpoint sets AzureDceEndpoint field to given value.

### GetAzureDcrImmutableId

`func (o *LogExportClusterSpecification) GetAzureDcrImmutableId() string`

GetAzureDcrImmutableId returns the AzureDcrImmutableId field if non-nil, zero value otherwise.

### SetAzureDcrImmutableId

`func (o *LogExportClusterSpecification) SetAzureDcrImmutableId(v string)`

SetAzureDcrImmutableId sets AzureDcrImmutableId field to given value.

### GetAzureDcrResourceId

`func (o *LogExportClusterSpecification) GetAzureDcrResourceId() string`

GetAzureDcrResourceId returns the AzureDcrResourceId field if non-nil, zero value otherwise.

### SetAzureDcrResourceId

`func (o *LogExportClusterSpecification) SetAzureDcrResourceId(v string)`

SetAzureDcrResourceId sets AzureDcrResourceId field to given value.

### GetAzureSharedKey

`func (o *LogExportClusterSpecification) GetAzureSharedKey() string`

GetAzureSharedKey returns the AzureSharedKey field if non-nil, zero value otherwise.

### SetAzureSharedKey

`func (o *LogExportClusterSpecification) SetAzureSharedKey(v string)`

SetAzureSharedKey sets AzureSharedKey field to given value.

### GetAzureTenantId

`func (o *LogExportClusterSpecification) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### SetAzureTenantId

`func (o *LogExportClusterSpecification) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### GetAzureWorkspaceResourceId

`func (o *LogExportClusterSpecification) GetAzureWorkspaceResourceId() string`

GetAzureWorkspaceResourceId returns the AzureWorkspaceResourceId field if non-nil, zero value otherwise.

### SetAzureWorkspaceResourceId

`func (o *LogExportClusterSpecification) SetAzureWorkspaceResourceId(v string)`

SetAzureWorkspaceResourceId sets AzureWorkspaceResourceId field to given value.

### GetGroups

`func (o *LogExportClusterSpecification) GetGroups() []LogExportGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### SetGroups

`func (o *LogExportClusterSpecification) SetGroups(v []LogExportGroup)`

SetGroups sets Groups field to given value.

### GetLogName

`func (o *LogExportClusterSpecification) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### SetLogName

`func (o *LogExportClusterSpecification) SetLogName(v string)`

SetLogName sets LogName field to given value.

### GetOmittedChannels

`func (o *LogExportClusterSpecification) GetOmittedChannels() []string`

GetOmittedChannels returns the OmittedChannels field if non-nil, zero value otherwise.

### SetOmittedChannels

`func (o *LogExportClusterSpecification) SetOmittedChannels(v []string)`

SetOmittedChannels sets OmittedChannels field to given value.

### GetOtlpEndpoint

`func (o *LogExportClusterSpecification) GetOtlpEndpoint() string`

GetOtlpEndpoint returns the OtlpEndpoint field if non-nil, zero value otherwise.

### SetOtlpEndpoint

`func (o *LogExportClusterSpecification) SetOtlpEndpoint(v string)`

SetOtlpEndpoint sets OtlpEndpoint field to given value.

### GetOtlpHeaderNames

`func (o *LogExportClusterSpecification) GetOtlpHeaderNames() []string`

GetOtlpHeaderNames returns the OtlpHeaderNames field if non-nil, zero value otherwise.

### SetOtlpHeaderNames

`func (o *LogExportClusterSpecification) SetOtlpHeaderNames(v []string)`

SetOtlpHeaderNames sets OtlpHeaderNames field to given value.

### GetRedact

`func (o *LogExportClusterSpecification) GetRedact() bool`

GetRedact returns the Redact field if non-nil, zero value otherwise.

### SetRedact

`func (o *LogExportClusterSpecification) SetRedact(v bool)`

SetRedact sets Redact field to given value.

### GetRegion

`func (o *LogExportClusterSpecification) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *LogExportClusterSpecification) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetType

`func (o *LogExportClusterSpecification) GetType() LogExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *LogExportClusterSpecification) SetType(v LogExportType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


