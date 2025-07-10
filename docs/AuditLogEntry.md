# AuditLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**AuditLogAction**](AuditLogAction.md) |  | [optional] 
**ClusterId** | Pointer to **string** | ClusterId is the ID of the cluster to which this log entry applies, if it applies to a single cluster. | [optional] 
**ClusterName** | Pointer to **string** | ClusterName is the name of the cluster to which this log entry applies, if it applies to a single cluster. | [optional] 
**CreatedAt** | Pointer to **time.Time** | CreatedAt is the time that this log entry was recorded. | [optional] 
**Error** | Pointer to **string** | Error is the error that applies to this entry if it represents a failure. | [optional] 
**Id** | Pointer to **string** | Id uniquely identifies this entry. | [optional] 
**Metadata** | Pointer to [**AuditLogMetadata**](AuditLogMetadata.md) |  | [optional] 
**Payload** | Pointer to **string** | Payload is a representation of the essential details relating to this log entry. | [optional] 
**ServiceAccountName** | Pointer to **string** | ServiceAccountName is the name of the service account that triggered this log entry. If it was not a service account, it will be empty. | [optional] 
**SessionId** | Pointer to **string** | SessionId is an ID that can be used to correlate this log entry with others that are emitted as part of the same user session, typically for users interacting through the UI. It should be treated as an opaque string with no guaranteed structure. | [optional] 
**Source** | Pointer to [**AuditLogSource**](AuditLogSource.md) |  | [optional] 
**TraceId** | Pointer to **string** | TraceId is an ID that can be used to correlate this log entry with others that are emitted as part of the same process. It should be treated as an opaque string with no guaranteed structure. | [optional] 
**UserEmail** | Pointer to **string** | UserEmail is the email address of the user that triggered this log entry. If it was not a human user, it will be empty. | [optional] 

## Methods

### NewAuditLogEntry

`func NewAuditLogEntry() *AuditLogEntry`

NewAuditLogEntry instantiates a new AuditLogEntry object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAction

`func (o *AuditLogEntry) GetAction() AuditLogAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### SetAction

`func (o *AuditLogEntry) SetAction(v AuditLogAction)`

SetAction sets Action field to given value.

### GetClusterId

`func (o *AuditLogEntry) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### SetClusterId

`func (o *AuditLogEntry) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### GetClusterName

`func (o *AuditLogEntry) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### SetClusterName

`func (o *AuditLogEntry) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### GetCreatedAt

`func (o *AuditLogEntry) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *AuditLogEntry) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetError

`func (o *AuditLogEntry) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### SetError

`func (o *AuditLogEntry) SetError(v string)`

SetError sets Error field to given value.

### GetId

`func (o *AuditLogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *AuditLogEntry) SetId(v string)`

SetId sets Id field to given value.

### GetMetadata

`func (o *AuditLogEntry) GetMetadata() AuditLogMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### SetMetadata

`func (o *AuditLogEntry) SetMetadata(v AuditLogMetadata)`

SetMetadata sets Metadata field to given value.

### GetPayload

`func (o *AuditLogEntry) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### SetPayload

`func (o *AuditLogEntry) SetPayload(v string)`

SetPayload sets Payload field to given value.

### GetServiceAccountName

`func (o *AuditLogEntry) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### SetServiceAccountName

`func (o *AuditLogEntry) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### GetSessionId

`func (o *AuditLogEntry) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### SetSessionId

`func (o *AuditLogEntry) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### GetSource

`func (o *AuditLogEntry) GetSource() AuditLogSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### SetSource

`func (o *AuditLogEntry) SetSource(v AuditLogSource)`

SetSource sets Source field to given value.

### GetTraceId

`func (o *AuditLogEntry) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### SetTraceId

`func (o *AuditLogEntry) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### GetUserEmail

`func (o *AuditLogEntry) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### SetUserEmail

`func (o *AuditLogEntry) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


