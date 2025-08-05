# ListBackupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | [**[]BackupSummary**](BackupSummary.md) |  | 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListBackupsResponse

`func NewListBackupsResponse(backups []BackupSummary, ) *ListBackupsResponse`

NewListBackupsResponse instantiates a new ListBackupsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewListBackupsResponseWithDefaults

`func NewListBackupsResponseWithDefaults() *ListBackupsResponse`

NewListBackupsResponseWithDefaults instantiates a new ListBackupsResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetBackups

`func (o *ListBackupsResponse) GetBackups() []BackupSummary`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### SetBackups

`func (o *ListBackupsResponse) SetBackups(v []BackupSummary)`

SetBackups sets Backups field to given value.

### GetPagination

`func (o *ListBackupsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListBackupsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


