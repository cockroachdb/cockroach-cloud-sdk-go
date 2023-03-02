# Organization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewOrganization

`func NewOrganization(createdAt time.Time, id string, label string, name string, ) *Organization`

NewOrganization instantiates a new Organization object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewOrganizationWithDefaults

`func NewOrganizationWithDefaults() *Organization`

NewOrganizationWithDefaults instantiates a new Organization object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetCreatedAt

`func (o *Organization) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *Organization) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetId

`func (o *Organization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Organization) SetId(v string)`

SetId sets Id field to given value.

### GetLabel

`func (o *Organization) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### SetLabel

`func (o *Organization) SetLabel(v string)`

SetLabel sets Label field to given value.

### GetName

`func (o *Organization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Organization) SetName(v string)`

SetName sets Name field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


