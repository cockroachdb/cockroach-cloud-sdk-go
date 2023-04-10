# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | [**ResourceTypeType**](ResourceTypeType.md) |  | 

## Methods

### NewResource

`func NewResource(type_ ResourceTypeType, ) *Resource`

NewResource instantiates a new Resource object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.

### GetType

`func (o *Resource) GetType() ResourceTypeType`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *Resource) SetType(v ResourceTypeType)`

SetType sets Type field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


