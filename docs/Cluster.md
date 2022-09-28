# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CockroachVersion** | **string** |  | 
**Plan** | [**Plan**](Plan.md) |  | 
**CloudProvider** | [**ApiCloudProvider**](ApiCloudProvider.md) |  | 
**AccountId** | Pointer to **string** |  | [optional] 
**State** | [**ClusterStateType**](ClusterStateType.md) |  | 
**CreatorId** | **string** |  | 
**OperationStatus** | [**ClusterStatusType**](ClusterStatusType.md) |  | 
**Config** | [**ClusterConfig**](ClusterConfig.md) |  | 
**Regions** | [**[]Region**](Region.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCluster

`func NewCluster(id string, name string, cockroachVersion string, plan Plan, cloudProvider ApiCloudProvider, state ClusterStateType, creatorId string, operationStatus ClusterStatusType, config ClusterConfig, regions []Region, ) *Cluster`

NewCluster instantiates a new Cluster object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### GetCockroachVersion

`func (o *Cluster) GetCockroachVersion() string`

GetCockroachVersion returns the CockroachVersion field if non-nil, zero value otherwise.

### SetCockroachVersion

`func (o *Cluster) SetCockroachVersion(v string)`

SetCockroachVersion sets CockroachVersion field to given value.

### GetPlan

`func (o *Cluster) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### SetPlan

`func (o *Cluster) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### GetCloudProvider

`func (o *Cluster) GetCloudProvider() ApiCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v ApiCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.

### GetAccountId

`func (o *Cluster) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### SetAccountId

`func (o *Cluster) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### GetState

`func (o *Cluster) GetState() ClusterStateType`

GetState returns the State field if non-nil, zero value otherwise.

### SetState

`func (o *Cluster) SetState(v ClusterStateType)`

SetState sets State field to given value.

### GetCreatorId

`func (o *Cluster) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### SetCreatorId

`func (o *Cluster) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### GetOperationStatus

`func (o *Cluster) GetOperationStatus() ClusterStatusType`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### SetOperationStatus

`func (o *Cluster) SetOperationStatus(v ClusterStatusType)`

SetOperationStatus sets OperationStatus field to given value.

### GetConfig

`func (o *Cluster) GetConfig() ClusterConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### SetConfig

`func (o *Cluster) SetConfig(v ClusterConfig)`

SetConfig sets Config field to given value.

### GetRegions

`func (o *Cluster) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *Cluster) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### GetDeletedAt

`func (o *Cluster) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### SetDeletedAt

`func (o *Cluster) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


