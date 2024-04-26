# UpdateReplicationStreamSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cutover** | **bool** | cutover describes whether or not to trigger cutover. | 
**CutoverAt** | Pointer to **time.Time** | cutover_at is the timestamp at which cutover occurs. If the user sets &#x60;cutover&#x60; to &#x60;true&#x60; but omits cutover_at, the cutover time will default to the latest consistent replicated time. Otherwise, the user can pick a time in the future to schedule a cutover in the future, or a time in the past to restore the cluster to a recent state. | [optional] 

## Methods

### NewUpdateReplicationStreamSpec

`func NewUpdateReplicationStreamSpec(cutover bool, ) *UpdateReplicationStreamSpec`

NewUpdateReplicationStreamSpec instantiates a new UpdateReplicationStreamSpec object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewUpdateReplicationStreamSpecWithDefaults

`func NewUpdateReplicationStreamSpecWithDefaults() *UpdateReplicationStreamSpec`

NewUpdateReplicationStreamSpecWithDefaults instantiates a new UpdateReplicationStreamSpec object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCutover

`func (o *UpdateReplicationStreamSpec) GetCutover() bool`

GetCutover returns the Cutover field if non-nil, zero value otherwise.

### SetCutover

`func (o *UpdateReplicationStreamSpec) SetCutover(v bool)`

SetCutover sets Cutover field to given value.

### GetCutoverAt

`func (o *UpdateReplicationStreamSpec) GetCutoverAt() time.Time`

GetCutoverAt returns the CutoverAt field if non-nil, zero value otherwise.

### SetCutoverAt

`func (o *UpdateReplicationStreamSpec) SetCutoverAt(v time.Time)`

SetCutoverAt sets CutoverAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


