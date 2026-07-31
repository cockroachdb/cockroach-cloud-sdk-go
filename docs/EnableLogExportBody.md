# EnableLogExportBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPrincipal** | Pointer to **string** | auth_principal is used in different contexts based on integration. CloudWatch: AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch GCP Cloud Logging: GCP Project ID that the cluster service account has permissions to write to for cloud logging. Azure Log Analytics (AZURE_LOG_ANALYTICS): CustomerID or WorkspaceID. Required for AWS_CLOUDWATCH, GCP_CLOUD_LOGGING, and AZURE_LOG_ANALYTICS. Not used for AZURE_LOG_ANALYTICS_V2 or OTLP_HTTP. OTLP_HTTP authenticates via otlp_headers (see otlp_endpoint / otlp_headers). | [optional] 
**AwsExternalId** | Pointer to **string** | aws_external_id to include when assuming the IAM role specified by role_arn. Optional. A specific value may be required by the role&#39;s trust policy. Only supported for Advanced clusters on AWS. If provided for a Standard cluster, the request is rejected. | [optional] 
**AzureClientId** | Pointer to **string** | Azure client ID for the app registration used by the Logs Ingestion API. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureClientSecret** | Pointer to **string** | Azure client secret for the app registration used by the Logs Ingestion API. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureDceEndpoint** | Pointer to **string** | Logs ingestion endpoint of the Azure Data Collection Endpoint (DCE). For example, https://{dce-name}.{region}.ingest.monitor.azure.com. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureDcrImmutableId** | Pointer to **string** | Immutable ID of the Azure Data Collection Rule (DCR), for example dcr-... Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureDcrResourceId** | Pointer to **string** | Full ARM resource ID of the Azure Data Collection Rule (DCR). Cockroach Cloud reads and updates this DCR to add streams for each configured log group. For example, /subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.Insights/dataCollectionRules/{dcr-name}. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureSharedKey** | Pointer to **string** | The primary or the secondary connected sources client authentication key. This is used to export logs to Azure Log Analytics via the legacy HTTP Data Collector API. Deprecated: use azure_client_secret instead. | [optional] 
**AzureTenantId** | Pointer to **string** | Azure tenant ID for the app registration used by the Logs Ingestion API. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**AzureWorkspaceResourceId** | Pointer to **string** | Full ARM resource ID of the Log Analytics workspace. Cockroach Cloud creates or updates the required custom tables for each configured log group. For example, /subscriptions/{subscription-id}/resourceGroups/{resource-group}/providers/Microsoft.OperationalInsights/workspaces/{workspace-name}. Required when type is AZURE_LOG_ANALYTICS_V2. | [optional] 
**Groups** | Pointer to [**[]LogExportGroup**](LogExportGroup.md) | groups is a collection of log group configurations that allows the customer to define collections of CRDB log channels that are aggregated separately at the target sink. | [optional] 
**LogName** | **string** | log_name is an identifier for the logs in the customer&#39;s log sink. For AZURE_LOG_ANALYTICS_V2, it must start with a letter and contain only letters, digits, and underscores. | 
**OmittedChannels** | Pointer to **[]string** | omitted_channels is a list of channels that the user does not want to export logs for. | [optional] 
**OtlpEndpoint** | Pointer to **string** | otlp_endpoint is the OTLP/HTTP URL for the OTLP_HTTP sink type. Customers may provide either a base endpoint or a full /v1/logs endpoint. Required when type is OTLP_HTTP. | [optional] 
**OtlpHeaders** | Pointer to **map[string]string** | otlp_headers are auth headers (name-&gt;value) for the OTLP_HTTP sink, e.g. {\&quot;authorization\&quot;: \&quot;Bearer ...\&quot;}. Write-only: values are stored securely and never returned. For existing OTLP_HTTP configurations, omitting this field or sending an empty map preserves the stored headers; sending a non-empty map replaces them. | [optional] 
**Redact** | Pointer to **bool** | redact allows the customer to set a default redaction policy for logs before they are exported to the target sink. If a group config omits a redact flag and this one is set to &#x60;true&#x60;, then that group will receive redacted logs. | [optional] 
**Region** | Pointer to **string** | region allows the customer to override the destination region for all logs for a cluster. | [optional] 
**Type** | [**LogExportType**](LogExportType.md) |  | 

## Methods

### NewEnableLogExportBody

`func NewEnableLogExportBody(logName string, type_ LogExportType, ) *EnableLogExportBody`

NewEnableLogExportBody instantiates a new EnableLogExportBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEnableLogExportBodyWithDefaults

`func NewEnableLogExportBodyWithDefaults() *EnableLogExportBody`

NewEnableLogExportBodyWithDefaults instantiates a new EnableLogExportBody object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAuthPrincipal

`func (o *EnableLogExportBody) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### SetAuthPrincipal

`func (o *EnableLogExportBody) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.

### GetAwsExternalId

`func (o *EnableLogExportBody) GetAwsExternalId() string`

GetAwsExternalId returns the AwsExternalId field if non-nil, zero value otherwise.

### SetAwsExternalId

`func (o *EnableLogExportBody) SetAwsExternalId(v string)`

SetAwsExternalId sets AwsExternalId field to given value.

### GetAzureClientId

`func (o *EnableLogExportBody) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### SetAzureClientId

`func (o *EnableLogExportBody) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### GetAzureClientSecret

`func (o *EnableLogExportBody) GetAzureClientSecret() string`

GetAzureClientSecret returns the AzureClientSecret field if non-nil, zero value otherwise.

### SetAzureClientSecret

`func (o *EnableLogExportBody) SetAzureClientSecret(v string)`

SetAzureClientSecret sets AzureClientSecret field to given value.

### GetAzureDceEndpoint

`func (o *EnableLogExportBody) GetAzureDceEndpoint() string`

GetAzureDceEndpoint returns the AzureDceEndpoint field if non-nil, zero value otherwise.

### SetAzureDceEndpoint

`func (o *EnableLogExportBody) SetAzureDceEndpoint(v string)`

SetAzureDceEndpoint sets AzureDceEndpoint field to given value.

### GetAzureDcrImmutableId

`func (o *EnableLogExportBody) GetAzureDcrImmutableId() string`

GetAzureDcrImmutableId returns the AzureDcrImmutableId field if non-nil, zero value otherwise.

### SetAzureDcrImmutableId

`func (o *EnableLogExportBody) SetAzureDcrImmutableId(v string)`

SetAzureDcrImmutableId sets AzureDcrImmutableId field to given value.

### GetAzureDcrResourceId

`func (o *EnableLogExportBody) GetAzureDcrResourceId() string`

GetAzureDcrResourceId returns the AzureDcrResourceId field if non-nil, zero value otherwise.

### SetAzureDcrResourceId

`func (o *EnableLogExportBody) SetAzureDcrResourceId(v string)`

SetAzureDcrResourceId sets AzureDcrResourceId field to given value.

### GetAzureSharedKey

`func (o *EnableLogExportBody) GetAzureSharedKey() string`

GetAzureSharedKey returns the AzureSharedKey field if non-nil, zero value otherwise.

### SetAzureSharedKey

`func (o *EnableLogExportBody) SetAzureSharedKey(v string)`

SetAzureSharedKey sets AzureSharedKey field to given value.

### GetAzureTenantId

`func (o *EnableLogExportBody) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### SetAzureTenantId

`func (o *EnableLogExportBody) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### GetAzureWorkspaceResourceId

`func (o *EnableLogExportBody) GetAzureWorkspaceResourceId() string`

GetAzureWorkspaceResourceId returns the AzureWorkspaceResourceId field if non-nil, zero value otherwise.

### SetAzureWorkspaceResourceId

`func (o *EnableLogExportBody) SetAzureWorkspaceResourceId(v string)`

SetAzureWorkspaceResourceId sets AzureWorkspaceResourceId field to given value.

### GetGroups

`func (o *EnableLogExportBody) GetGroups() []LogExportGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### SetGroups

`func (o *EnableLogExportBody) SetGroups(v []LogExportGroup)`

SetGroups sets Groups field to given value.

### GetLogName

`func (o *EnableLogExportBody) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### SetLogName

`func (o *EnableLogExportBody) SetLogName(v string)`

SetLogName sets LogName field to given value.

### GetOmittedChannels

`func (o *EnableLogExportBody) GetOmittedChannels() []string`

GetOmittedChannels returns the OmittedChannels field if non-nil, zero value otherwise.

### SetOmittedChannels

`func (o *EnableLogExportBody) SetOmittedChannels(v []string)`

SetOmittedChannels sets OmittedChannels field to given value.

### GetOtlpEndpoint

`func (o *EnableLogExportBody) GetOtlpEndpoint() string`

GetOtlpEndpoint returns the OtlpEndpoint field if non-nil, zero value otherwise.

### SetOtlpEndpoint

`func (o *EnableLogExportBody) SetOtlpEndpoint(v string)`

SetOtlpEndpoint sets OtlpEndpoint field to given value.

### GetOtlpHeaders

`func (o *EnableLogExportBody) GetOtlpHeaders() map[string]string`

GetOtlpHeaders returns the OtlpHeaders field if non-nil, zero value otherwise.

### SetOtlpHeaders

`func (o *EnableLogExportBody) SetOtlpHeaders(v map[string]string)`

SetOtlpHeaders sets OtlpHeaders field to given value.

### GetRedact

`func (o *EnableLogExportBody) GetRedact() bool`

GetRedact returns the Redact field if non-nil, zero value otherwise.

### SetRedact

`func (o *EnableLogExportBody) SetRedact(v bool)`

SetRedact sets Redact field to given value.

### GetRegion

`func (o *EnableLogExportBody) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *EnableLogExportBody) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetType

`func (o *EnableLogExportBody) GetType() LogExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *EnableLogExportBody) SetType(v LogExportType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


