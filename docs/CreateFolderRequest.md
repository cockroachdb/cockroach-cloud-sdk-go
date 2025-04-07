# CreateFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. | [optional] 
**Name** | **string** |  | 
**ParentId** | Pointer to **string** | The parent ID is a folder ID. An empty string or \&quot;root\&quot; will create a folder at the root level. | [optional] 

## Methods

### NewCreateFolderRequest

`func NewCreateFolderRequest(name string, ) *CreateFolderRequest`

NewCreateFolderRequest instantiates a new CreateFolderRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewCreateFolderRequestWithDefaults

`func NewCreateFolderRequestWithDefaults() *CreateFolderRequest`

NewCreateFolderRequestWithDefaults instantiates a new CreateFolderRequest object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetLabels

`func (o *CreateFolderRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *CreateFolderRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetName

`func (o *CreateFolderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *CreateFolderRequest) SetName(v string)`

SetName sets Name field to given value.

### GetParentId

`func (o *CreateFolderRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *CreateFolderRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


