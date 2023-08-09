# EnableLogExportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthPrincipal** | **string** | auth_principal is either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP Project ID that the cluster service account has permissions to write to for cloud logging. | 
**Groups** | Pointer to [**[]LogExportGroup**](LogExportGroup.md) | groups is a collection of log group configurations that allows the customer to define collections of CRDB log channels that are aggregated separately at the target sink. | [optional] 
**LogName** | **string** | log_name is an identifier for the logs in the customer&#39;s log sink. | 
**OmittedChannels** | Pointer to **[]string** | omitted_channels is a list of channels that the user does not want to export logs for. | [optional] 
**Redact** | Pointer to **bool** | redact allows the customer to set a default redaction policy for logs before they are exported to the target sink. If a group config omits a redact flag and this one is set to &#x60;true&#x60;, then that group will receive redacted logs. | [optional] 
**Region** | Pointer to **string** | region allows the customer to override the destination region for all logs for a cluster. | [optional] 
**Type** | [**LogExportType**](LogExportType.md) |  | 

## Methods

### NewEnableLogExportRequest

`func NewEnableLogExportRequest(authPrincipal string, logName string, type_ LogExportType, ) *EnableLogExportRequest`

NewEnableLogExportRequest instantiates a new EnableLogExportRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewEnableLogExportRequestWithDefaults

`func NewEnableLogExportRequestWithDefaults() *EnableLogExportRequest`

NewEnableLogExportRequestWithDefaults instantiates a new EnableLogExportRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAuthPrincipal

`func (o *EnableLogExportRequest) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### SetAuthPrincipal

`func (o *EnableLogExportRequest) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.

### GetGroups

`func (o *EnableLogExportRequest) GetGroups() []LogExportGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### SetGroups

`func (o *EnableLogExportRequest) SetGroups(v []LogExportGroup)`

SetGroups sets Groups field to given value.

### GetLogName

`func (o *EnableLogExportRequest) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### SetLogName

`func (o *EnableLogExportRequest) SetLogName(v string)`

SetLogName sets LogName field to given value.

### GetOmittedChannels

`func (o *EnableLogExportRequest) GetOmittedChannels() []string`

GetOmittedChannels returns the OmittedChannels field if non-nil, zero value otherwise.

### SetOmittedChannels

`func (o *EnableLogExportRequest) SetOmittedChannels(v []string)`

SetOmittedChannels sets OmittedChannels field to given value.

### GetRedact

`func (o *EnableLogExportRequest) GetRedact() bool`

GetRedact returns the Redact field if non-nil, zero value otherwise.

### SetRedact

`func (o *EnableLogExportRequest) SetRedact(v bool)`

SetRedact sets Redact field to given value.

### GetRegion

`func (o *EnableLogExportRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *EnableLogExportRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### GetType

`func (o *EnableLogExportRequest) GetType() LogExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *EnableLogExportRequest) SetType(v LogExportType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


