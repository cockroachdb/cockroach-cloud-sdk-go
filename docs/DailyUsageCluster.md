# DailyUsageCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [**CloudProviderType**](CloudProviderType.md) |  | 
**Edition** | Pointer to [**EditionType**](EditionType.md) |  | [optional] 
**Id** | **string** | id is the unique ID of the cluster. | 
**Name** | **string** | name is the cluster&#39;s name. Retained so rows for deleted clusters stay human-readable. | 
**PlanType** | Pointer to [**PlanType**](PlanType.md) |  | [optional] 

## Methods

### NewDailyUsageCluster

`func NewDailyUsageCluster(cloudProvider CloudProviderType, id string, name string, ) *DailyUsageCluster`

NewDailyUsageCluster instantiates a new DailyUsageCluster object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDailyUsageClusterWithDefaults

`func NewDailyUsageClusterWithDefaults() *DailyUsageCluster`

NewDailyUsageClusterWithDefaults instantiates a new DailyUsageCluster object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCloudProvider

`func (o *DailyUsageCluster) GetCloudProvider() CloudProviderType`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### SetCloudProvider

`func (o *DailyUsageCluster) SetCloudProvider(v CloudProviderType)`

SetCloudProvider sets CloudProvider field to given value.

### GetEdition

`func (o *DailyUsageCluster) GetEdition() EditionType`

GetEdition returns the Edition field if non-nil, zero value otherwise.

### SetEdition

`func (o *DailyUsageCluster) SetEdition(v EditionType)`

SetEdition sets Edition field to given value.

### GetId

`func (o *DailyUsageCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *DailyUsageCluster) SetId(v string)`

SetId sets Id field to given value.

### GetName

`func (o *DailyUsageCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *DailyUsageCluster) SetName(v string)`

SetName sets Name field to given value.

### GetPlanType

`func (o *DailyUsageCluster) GetPlanType() PlanType`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### SetPlanType

`func (o *DailyUsageCluster) SetPlanType(v PlanType)`

SetPlanType sets PlanType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


