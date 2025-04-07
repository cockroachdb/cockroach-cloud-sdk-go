# UpdateFolderSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **map[string]string** | labels are key-value pairs used to organize and categorize resources. If the labels field is included in the request: Any existing labels on the folder that are not included will be removed, and any new labels specified will be added. If the labels field is omitted from the request entirely, all existing labels will remain unchanged. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateFolderSpecification

`func NewUpdateFolderSpecification() *UpdateFolderSpecification`

NewUpdateFolderSpecification instantiates a new UpdateFolderSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetLabels

`func (o *UpdateFolderSpecification) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### SetLabels

`func (o *UpdateFolderSpecification) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### GetName

`func (o *UpdateFolderSpecification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *UpdateFolderSpecification) SetName(v string)`

SetName sets Name field to given value.

### GetParentId

`func (o *UpdateFolderSpecification) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *UpdateFolderSpecification) SetParentId(v string)`

SetParentId sets ParentId field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


