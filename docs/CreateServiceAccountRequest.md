# CreateServiceAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description of the service account. | 
**Name** | **string** | name of the service account. | 
**Roles** | [**[]BuiltInRole**](BuiltInRole.md) | roles that are assigned to the service account. | 

## Methods

### NewCreateServiceAccountRequest

`func NewCreateServiceAccountRequest(description string, name string, roles []BuiltInRole, ) *CreateServiceAccountRequest`

NewCreateServiceAccountRequest instantiates a new CreateServiceAccountRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateServiceAccountRequestWithDefaults

`func NewCreateServiceAccountRequestWithDefaults() *CreateServiceAccountRequest`

NewCreateServiceAccountRequestWithDefaults instantiates a new CreateServiceAccountRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDescription

`func (o *CreateServiceAccountRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### SetDescription

`func (o *CreateServiceAccountRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### GetName

`func (o *CreateServiceAccountRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateServiceAccountRequest) SetName(v string)`

SetName sets Name field to given value.

### GetRoles

`func (o *CreateServiceAccountRequest) GetRoles() []BuiltInRole`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### SetRoles

`func (o *CreateServiceAccountRequest) SetRoles(v []BuiltInRole)`

SetRoles sets Roles field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


