# PhysicalReplicationStreamList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**PhysicalReplicationStreams** | [**[]PhysicalReplicationStream**](PhysicalReplicationStream.md) | physical_replication_streams is a list of PhysicalReplicationStreams. | 

## Methods

### NewPhysicalReplicationStreamList

`func NewPhysicalReplicationStreamList(physicalReplicationStreams []PhysicalReplicationStream, ) *PhysicalReplicationStreamList`

NewPhysicalReplicationStreamList instantiates a new PhysicalReplicationStreamList object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewPhysicalReplicationStreamListWithDefaults

`func NewPhysicalReplicationStreamListWithDefaults() *PhysicalReplicationStreamList`

NewPhysicalReplicationStreamListWithDefaults instantiates a new PhysicalReplicationStreamList object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetPagination

`func (o *PhysicalReplicationStreamList) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *PhysicalReplicationStreamList) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetPhysicalReplicationStreams

`func (o *PhysicalReplicationStreamList) GetPhysicalReplicationStreams() []PhysicalReplicationStream`

GetPhysicalReplicationStreams returns the PhysicalReplicationStreams field if non-nil, zero value otherwise.

### SetPhysicalReplicationStreams

`func (o *PhysicalReplicationStreamList) SetPhysicalReplicationStreams(v []PhysicalReplicationStream)`

SetPhysicalReplicationStreams sets PhysicalReplicationStreams field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


