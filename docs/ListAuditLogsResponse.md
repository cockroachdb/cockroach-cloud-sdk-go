# ListAuditLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]AuditLogEntry**](AuditLogEntry.md) | entries is the contiguous list of audit log entries matching the pagination request, sorted in the order requested. | [optional] 
**NextStartingFrom** | Pointer to **time.Time** | next_starting_from is the timestamp the caller should use to continue paginating in the same direction. If the timestamp is unset, it means that there are no more entries in the direction of the request, and there never will be. | [optional] 

## Methods

### NewListAuditLogsResponse

`func NewListAuditLogsResponse() *ListAuditLogsResponse`

NewListAuditLogsResponse instantiates a new ListAuditLogsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetEntries

`func (o *ListAuditLogsResponse) GetEntries() []AuditLogEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### SetEntries

`func (o *ListAuditLogsResponse) SetEntries(v []AuditLogEntry)`

SetEntries sets Entries field to given value.

### GetNextStartingFrom

`func (o *ListAuditLogsResponse) GetNextStartingFrom() time.Time`

GetNextStartingFrom returns the NextStartingFrom field if non-nil, zero value otherwise.

### SetNextStartingFrom

`func (o *ListAuditLogsResponse) SetNextStartingFrom(v time.Time)`

SetNextStartingFrom sets NextStartingFrom field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


