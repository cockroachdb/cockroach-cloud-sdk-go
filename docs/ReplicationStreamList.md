# ReplicationStreamList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**ReplicationStreams** | [**[]ReplicationStream**](ReplicationStream.md) | replication_streams is a list of ReplicationStreams. | 

## Methods

### NewReplicationStreamList

`func NewReplicationStreamList(replicationStreams []ReplicationStream, ) *ReplicationStreamList`

NewReplicationStreamList instantiates a new ReplicationStreamList object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewReplicationStreamListWithDefaults

`func NewReplicationStreamListWithDefaults() *ReplicationStreamList`

NewReplicationStreamListWithDefaults instantiates a new ReplicationStreamList object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *ReplicationStreamList) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ReplicationStreamList) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetReplicationStreams

`func (o *ReplicationStreamList) GetReplicationStreams() []ReplicationStream`

GetReplicationStreams returns the ReplicationStreams field if non-nil, zero value otherwise.

### SetReplicationStreams

`func (o *ReplicationStreamList) SetReplicationStreams(v []ReplicationStream)`

SetReplicationStreams sets ReplicationStreams field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


