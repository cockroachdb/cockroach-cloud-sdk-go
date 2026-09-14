# CreateMigrationAssistantBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The cluster region to deploy into. Optional; empty selects a default region (the cluster&#39;s first region ordered by locality, which is not necessarily its original primary region). Must be one of the cluster&#39;s existing regions. The chosen region is reported as the assistant&#39;s deployed_region. | [optional] 

## Methods

### NewCreateMigrationAssistantBody

`func NewCreateMigrationAssistantBody() *CreateMigrationAssistantBody`

NewCreateMigrationAssistantBody instantiates a new CreateMigrationAssistantBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetRegion

`func (o *CreateMigrationAssistantBody) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *CreateMigrationAssistantBody) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


