# CreateClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Provider** | [**CloudProvider**](CloudProvider.md) |  | [default to UNSPECIFIED]
**Spec** | [**CreateClusterSpecification**](CreateClusterSpecification.md) |  | 

## Methods

### NewCreateClusterRequest

`func NewCreateClusterRequest(name string, provider CloudProvider, spec CreateClusterSpecification, ) *CreateClusterRequest`

NewCreateClusterRequest instantiates a new CreateClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterRequestWithDefaults

`func NewCreateClusterRequestWithDefaults() *CreateClusterRequest`

NewCreateClusterRequestWithDefaults instantiates a new CreateClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateClusterRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *CreateClusterRequest) GetProvider() CloudProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateClusterRequest) GetProviderOk() (*CloudProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateClusterRequest) SetProvider(v CloudProvider)`

SetProvider sets Provider field to given value.


### GetSpec

`func (o *CreateClusterRequest) GetSpec() CreateClusterSpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateClusterRequest) GetSpecOk() (*CreateClusterSpecification, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateClusterRequest) SetSpec(v CreateClusterSpecification)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


