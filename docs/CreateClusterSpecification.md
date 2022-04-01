# CreateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterCreateSpecification**](DedicatedClusterCreateSpecification.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterSpecification**](ServerlessClusterSpecification.md) |  | [optional] 

## Methods

### NewCreateClusterSpecification

`func NewCreateClusterSpecification() *CreateClusterSpecification`

NewCreateClusterSpecification instantiates a new CreateClusterSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterSpecificationWithDefaults

`func NewCreateClusterSpecificationWithDefaults() *CreateClusterSpecification`

NewCreateClusterSpecificationWithDefaults instantiates a new CreateClusterSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *CreateClusterSpecification) GetDedicated() DedicatedClusterCreateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *CreateClusterSpecification) GetDedicatedOk() (*DedicatedClusterCreateSpecification, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *CreateClusterSpecification) SetDedicated(v DedicatedClusterCreateSpecification)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *CreateClusterSpecification) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetServerless

`func (o *CreateClusterSpecification) GetServerless() ServerlessClusterSpecification`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *CreateClusterSpecification) GetServerlessOk() (*ServerlessClusterSpecification, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *CreateClusterSpecification) SetServerless(v ServerlessClusterSpecification)`

SetServerless sets Serverless field to given value.

### HasServerless

`func (o *CreateClusterSpecification) HasServerless() bool`

HasServerless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


