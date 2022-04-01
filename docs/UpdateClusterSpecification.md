# UpdateClusterSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to [**DedicatedClusterUpdateSpecification**](DedicatedClusterUpdateSpecification.md) |  | [optional] 
**Serverless** | Pointer to [**ServerlessClusterSpecification**](ServerlessClusterSpecification.md) |  | [optional] 

## Methods

### NewUpdateClusterSpecification

`func NewUpdateClusterSpecification() *UpdateClusterSpecification`

NewUpdateClusterSpecification instantiates a new UpdateClusterSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterSpecificationWithDefaults

`func NewUpdateClusterSpecificationWithDefaults() *UpdateClusterSpecification`

NewUpdateClusterSpecificationWithDefaults instantiates a new UpdateClusterSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *UpdateClusterSpecification) GetDedicated() DedicatedClusterUpdateSpecification`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *UpdateClusterSpecification) GetDedicatedOk() (*DedicatedClusterUpdateSpecification, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *UpdateClusterSpecification) SetDedicated(v DedicatedClusterUpdateSpecification)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *UpdateClusterSpecification) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetServerless

`func (o *UpdateClusterSpecification) GetServerless() ServerlessClusterSpecification`

GetServerless returns the Serverless field if non-nil, zero value otherwise.

### GetServerlessOk

`func (o *UpdateClusterSpecification) GetServerlessOk() (*ServerlessClusterSpecification, bool)`

GetServerlessOk returns a tuple with the Serverless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerless

`func (o *UpdateClusterSpecification) SetServerless(v ServerlessClusterSpecification)`

SetServerless sets Serverless field to given value.

### HasServerless

`func (o *UpdateClusterSpecification) HasServerless() bool`

HasServerless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


