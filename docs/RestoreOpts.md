# RestoreOpts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntoDb** | Pointer to **string** | Specifies the target database to restore the table into during a table restore job. If not set, the table will be restored into the database it belonged to in the source backup. | [optional] 
**NewDbName** | Pointer to **string** | Specifies the name of the database to create during a database restore job. If not set, the name defaults to the original database name from the source cluster. | [optional] 
**SchemaOnly** | Pointer to **bool** | If set, only the schema will be restored and no user data will be included. | [optional] 
**SkipLocalitiesCheck** | Pointer to **bool** | Allows the restore job to continue in the event that there are mismatched localities between the backup and target cluster. Useful when restoring multi-region tables to a cluster missing some localities. | [optional] 
**SkipMissingForeignKeys** | Pointer to **bool** | Allows a table to be restored even if it has foreign key constraints referencing rows that no longer exist in the target cluster. | [optional] 
**SkipMissingSequences** | Pointer to **bool** |  | [optional] 
**SkipMissingViews** | Pointer to **bool** | Allows the job to skip restoring views that cannot be restored because their dependencies are not included in the current restore job. | [optional] 

## Methods

### NewRestoreOpts

`func NewRestoreOpts() *RestoreOpts`

NewRestoreOpts instantiates a new RestoreOpts object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetIntoDb

`func (o *RestoreOpts) GetIntoDb() string`

GetIntoDb returns the IntoDb field if non-nil, zero value otherwise.

### SetIntoDb

`func (o *RestoreOpts) SetIntoDb(v string)`

SetIntoDb sets IntoDb field to given value.

### GetNewDbName

`func (o *RestoreOpts) GetNewDbName() string`

GetNewDbName returns the NewDbName field if non-nil, zero value otherwise.

### SetNewDbName

`func (o *RestoreOpts) SetNewDbName(v string)`

SetNewDbName sets NewDbName field to given value.

### GetSchemaOnly

`func (o *RestoreOpts) GetSchemaOnly() bool`

GetSchemaOnly returns the SchemaOnly field if non-nil, zero value otherwise.

### SetSchemaOnly

`func (o *RestoreOpts) SetSchemaOnly(v bool)`

SetSchemaOnly sets SchemaOnly field to given value.

### GetSkipLocalitiesCheck

`func (o *RestoreOpts) GetSkipLocalitiesCheck() bool`

GetSkipLocalitiesCheck returns the SkipLocalitiesCheck field if non-nil, zero value otherwise.

### SetSkipLocalitiesCheck

`func (o *RestoreOpts) SetSkipLocalitiesCheck(v bool)`

SetSkipLocalitiesCheck sets SkipLocalitiesCheck field to given value.

### GetSkipMissingForeignKeys

`func (o *RestoreOpts) GetSkipMissingForeignKeys() bool`

GetSkipMissingForeignKeys returns the SkipMissingForeignKeys field if non-nil, zero value otherwise.

### SetSkipMissingForeignKeys

`func (o *RestoreOpts) SetSkipMissingForeignKeys(v bool)`

SetSkipMissingForeignKeys sets SkipMissingForeignKeys field to given value.

### GetSkipMissingSequences

`func (o *RestoreOpts) GetSkipMissingSequences() bool`

GetSkipMissingSequences returns the SkipMissingSequences field if non-nil, zero value otherwise.

### SetSkipMissingSequences

`func (o *RestoreOpts) SetSkipMissingSequences(v bool)`

SetSkipMissingSequences sets SkipMissingSequences field to given value.

### GetSkipMissingViews

`func (o *RestoreOpts) GetSkipMissingViews() bool`

GetSkipMissingViews returns the SkipMissingViews field if non-nil, zero value otherwise.

### SetSkipMissingViews

`func (o *RestoreOpts) SetSkipMissingViews(v bool)`

SetSkipMissingViews sets SkipMissingViews field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


