# FolderResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is the resource&#39;s name. | 
**OrganizationId** | **string** | organization_id is the id of the organization this resource belongs to. | 
**ParentId** | **string** | parent_id is the id of the resource&#39;s parent folder. \&quot;root\&quot; represents a root level resource. | 
**Path** | [**[]PathSegment**](PathSegment.md) | path contains the ids and names of ancestors that make up the resource&#39;s lineage. | 
**ResourceId** | **string** | resource_id is the resource&#39;s id. | 
**ResourceType** | [**FolderResourceTypeType**](FolderResourceTypeType.md) |  | 

## Methods

### NewFolderResource

`func NewFolderResource(name string, organizationId string, parentId string, path []PathSegment, resourceId string, resourceType FolderResourceTypeType, ) *FolderResource`

NewFolderResource instantiates a new FolderResource object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewFolderResourceWithDefaults

`func NewFolderResourceWithDefaults() *FolderResource`

NewFolderResourceWithDefaults instantiates a new FolderResource object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *FolderResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *FolderResource) SetName(v string)`

SetName sets Name field to given value.

### GetOrganizationId

`func (o *FolderResource) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### SetOrganizationId

`func (o *FolderResource) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### GetParentId

`func (o *FolderResource) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### SetParentId

`func (o *FolderResource) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### GetPath

`func (o *FolderResource) GetPath() []PathSegment`

GetPath returns the Path field if non-nil, zero value otherwise.

### SetPath

`func (o *FolderResource) SetPath(v []PathSegment)`

SetPath sets Path field to given value.

### GetResourceId

`func (o *FolderResource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### SetResourceId

`func (o *FolderResource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### GetResourceType

`func (o *FolderResource) GetResourceType() FolderResourceTypeType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### SetResourceType

`func (o *FolderResource) SetResourceType(v FolderResourceTypeType)`

SetResourceType sets ResourceType field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


