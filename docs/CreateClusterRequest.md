# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provider** | [**ApiCloudProvider**](ApiCloudProvider.md) |  | [default to APICLOUDPROVIDER_UNSPECIFIED]
**Spec** | [**CreateClusterSpecification**](CreateClusterSpecification.md) |  | 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest(name string, provider ApiCloudProvider, spec CreateClusterSpecification, ) *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.

### GetProvider

`func (o *CreateClusterRequest) GetProvider() ApiCloudProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### SetProvider

`func (o *CreateClusterRequest) SetProvider(v ApiCloudProvider)`

SetProvider sets Provider field to given value.

### GetSpec

`func (o *CreateClusterRequest) GetSpec() CreateClusterSpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### SetSpec

`func (o *CreateClusterRequest) SetSpec(v CreateClusterSpecification)`

SetSpec sets Spec field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


