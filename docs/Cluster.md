# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**CidrRange** | **string** | cidr_range is the IPv4 range in CIDR format that will be used by the cluster. It is only set on GCP Advanced tier clusters and is otherwise empty. | 
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**CockroachVersion** | **string** |  | 
**Config** | [**ClusterConfig**](ClusterConfig.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | **string** |  | 
**DeleteProtection** | Pointer to [**DeleteProtectionStateType**](DeleteProtectionStateType.md) |  | [optional] 
**DeletedAt** | Pointer to **time.Time** |  | [optional] 
**EgressTrafficPolicy** | Pointer to [**EgressTrafficPolicyType**](EgressTrafficPolicyType.md) |  | [optional] 
**Id** | **string** |  | 
**Labels** | **map[string]string** | labels are key-value pairs used to organize and categorize resources. | 
**Name** | **string** |  | 
**NetworkVisibility** | Pointer to [**NetworkVisibilityType**](NetworkVisibilityType.md) |  | [optional] 
**OperationStatus** | [**ClusterStatusType**](ClusterStatusType.md) |  | 
**ParentId** | Pointer to **string** | Preview: The parent ID is a folder ID. A \&quot;root\&quot; valued parent ID refers to a cluster at the root level. | [optional] 
**Plan** | [**PlanType**](PlanType.md) |  | 
**Regions** | [**[]Region**](Region.md) |  | 
**SqlDns** | Pointer to **string** | sql_dns is the DNS name of SQL interface of the cluster. | [optional] 
**State** | [**ClusterStateType**](ClusterStateType.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**UpgradeStatus** | [**ClusterUpgradeStatusType**](ClusterUpgradeStatusType.md) |  | 

## Methods

### NewCluster

`func NewCluster(cidrRange string, cloudProvider CloudProviderType, cockroachVersion string, config ClusterConfig, creatorId string, id string, labels map[string]string, name string, operationStatus ClusterStatusType, plan PlanType, regions []Region, state ClusterStateType, upgradeStatus ClusterUpgradeStatusType, ) *Cluster`

NewCluster instantiates a new Cluster object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetAccountId

`func (o *Cluster) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### SetAccountId

`func (o *Cluster) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### GetCidrRange

`func (o *Cluster) GetCidrRange() string`

GetCidrRange returns the CidrRange field if non-nil, zero value otherwise.

### SetCidrRange

`func (o *Cluster) SetCidrRange(v string)`

SetCidrRange sets CidrRange field to given value.

### GetCloudProvider

`func (o *Cluster) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### GetCockroachVersion

`func (o *Cluster) GetCockroachVersion() string`

GetCockroachVersion returns the CockroachVersion field if non-nil, zero value otherwise.

### SetCockroachVersion

`func (o *Cluster) SetCockroachVersion(v string)`

SetCockroachVersion sets CockroachVersion field to given value.

### GetConfig

`func (o *Cluster) GetConfig() ClusterConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### SetConfig

`func (o *Cluster) SetConfig(v ClusterConfig)`

SetConfig sets Config field to given value.

### GetCreatedAt

`func (o *Cluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *Cluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetCreatorId

`func (o *Cluster) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### SetCreatorId

`func (o *Cluster) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### GetDeleteProtection

`func (o *Cluster) GetDeleteProtection() DeleteProtectionStateType`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### SetDeleteProtection

`func (o *Cluster) SetDeleteProtection(v DeleteProtectionStateType)`

SetDeleteProtection sets DeleteProtection field to given value.

### GetDeletedAt

`func (o *Cluster) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### SetDeletedAt

`func (o *Cluster) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### GetEgressTrafficPolicy

`func (o *Cluster) GetEgressTrafficPolicy() EgressTrafficPolicyType`

GetEgressTrafficPolicy returns the EgressTrafficPolicy field if non-nil, zero value otherwise.

### SetEgressTrafficPolicy

`func (o *Cluster) SetEgressTrafficPolicy(v EgressTrafficPolicyType)`

SetEgressTrafficPolicy sets EgressTrafficPolicy field to given value.

### GetId

`func (o *Cluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Cluster) SetId(v string)`

SetId sets Id field to given value.

### GetLabels

`func (o *Cluster) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *Cluster) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.

### GetNetworkVisibility

`func (o *Cluster) GetNetworkVisibility() NetworkVisibilityType`

GetNetworkVisibility returns the NetworkVisibility field if non-nil, zero value otherwise.

### SetNetworkVisibility

`func (o *Cluster) SetNetworkVisibility(v NetworkVisibilityType)`

SetNetworkVisibility sets NetworkVisibility field to given value.

### GetOperationStatus

`func (o *Cluster) GetOperationStatus() ClusterStatusType`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### SetOperationStatus

`func (o *Cluster) SetOperationStatus(v ClusterStatusType)`

SetOperationStatus sets OperationStatus field to given value.

### GetParentId

`func (o *Cluster) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *Cluster) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPlan

`func (o *Cluster) GetPlan() PlanType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### SetPlan

`func (o *Cluster) SetPlan(v PlanType)`

SetPlan sets Plan field to given value.

### GetRegions

`func (o *Cluster) GetRegions() []Region`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### SetRegions

`func (o *Cluster) SetRegions(v []Region)`

SetRegions sets Regions field to given value.

### GetSqlDns

`func (o *Cluster) GetSqlDns() string`

GetSqlDns returns the SqlDns field if non-nil, zero value otherwise.

### SetSqlDns

`func (o *Cluster) SetSqlDns(v string)`

SetSqlDns sets SqlDns field to given value.

### GetState

`func (o *Cluster) GetState() ClusterStateType`

GetState returns the State field if non-nil, zero value otherwise.

### SetState

`func (o *Cluster) SetState(v ClusterStateType)`

SetState sets State field to given value.

### GetUpdatedAt

`func (o *Cluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *Cluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### GetUpgradeStatus

`func (o *Cluster) GetUpgradeStatus() ClusterUpgradeStatusType`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### SetUpgradeStatus

`func (o *Cluster) SetUpgradeStatus(v ClusterUpgradeStatusType)`

SetUpgradeStatus sets UpgradeStatus field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


