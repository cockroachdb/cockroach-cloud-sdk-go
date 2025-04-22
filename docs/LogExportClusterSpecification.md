# LogExportClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPrincipal** | Pointer to **string** | auth_principal is either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP Project ID that the cluster service account has permissions to write to for cloud logging. | [optional] 
**AwsExternalId** | Pointer to **string** | aws_external_id, if set, is included when assuming the IAM role. Supported for Advanced clusters on AWS only. | [optional] 
**AzureSharedKey** | Pointer to **string** | The primary or the secondary connected sources client authentication key. This is used to export logs to Azure Log Analytics. | [optional] 
**Groups** | Pointer to [**[]LogExportGroup**](LogExportGroup.md) | groups is a collection of log group configurations to customize which CRDB channels get aggregated into different groups at the target sink. Unconfigured channels will be sent to the default locations via the settings above. | [optional] 
**LogName** | Pointer to **string** | log_name is an identifier for the logs in the customer&#39;s log sink. | [optional] 
**OmittedChannels** | Pointer to **[]string** | omitted_channels is a list of channels that the user does not want to export logs for. | [optional] 
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

### GetAzureSharedKey

`func (o *LogExportClusterSpecification) GetAzureSharedKey() string`

GetAzureSharedKey returns the AzureSharedKey field if non-nil, zero value otherwise.

### SetAzureSharedKey

`func (o *LogExportClusterSpecification) SetAzureSharedKey(v string)`

SetAzureSharedKey sets AzureSharedKey field to given value.

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


