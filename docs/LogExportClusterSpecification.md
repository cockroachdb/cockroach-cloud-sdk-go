# LogExportClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LogExportType**](LogExportType.md) |  | [optional] 
**LogName** | Pointer to **string** | log_name is an identifier for the logs in the customer&#39;s log sink. | [optional] 
**AuthPrincipal** | Pointer to **string** | auth_principal is either the AWS Role ARN that identifies a role that the cluster account can assume to write to CloudWatch or the GCP Project ID that the cluster service account has permissions to write to for cloud logging. | [optional] 

## Methods

### NewLogExportClusterSpecification

`func NewLogExportClusterSpecification() *LogExportClusterSpecification`

NewLogExportClusterSpecification instantiates a new LogExportClusterSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetType

`func (o *LogExportClusterSpecification) GetType() LogExportType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *LogExportClusterSpecification) SetType(v LogExportType)`

SetType sets Type field to given value.

### GetLogName

`func (o *LogExportClusterSpecification) GetLogName() string`

GetLogName returns the LogName field if non-nil, zero value otherwise.

### SetLogName

`func (o *LogExportClusterSpecification) SetLogName(v string)`

SetLogName sets LogName field to given value.

### GetAuthPrincipal

`func (o *LogExportClusterSpecification) GetAuthPrincipal() string`

GetAuthPrincipal returns the AuthPrincipal field if non-nil, zero value otherwise.

### SetAuthPrincipal

`func (o *LogExportClusterSpecification) SetAuthPrincipal(v string)`

SetAuthPrincipal sets AuthPrincipal field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


