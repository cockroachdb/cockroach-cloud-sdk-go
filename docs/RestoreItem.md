# RestoreItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** |  | 
**Schema** | Pointer to **string** |  | [optional] 
**Table** | Pointer to **string** |  | [optional] 

## Methods

### NewRestoreItem

`func NewRestoreItem(database string, ) *RestoreItem`

NewRestoreItem instantiates a new RestoreItem object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewRestoreItemWithDefaults

`func NewRestoreItemWithDefaults() *RestoreItem`

NewRestoreItemWithDefaults instantiates a new RestoreItem object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDatabase

`func (o *RestoreItem) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### SetDatabase

`func (o *RestoreItem) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### GetSchema

`func (o *RestoreItem) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### SetSchema

`func (o *RestoreItem) SetSchema(v string)`

SetSchema sets Schema field to given value.

### GetTable

`func (o *RestoreItem) GetTable() string`

GetTable returns the Table field if non-nil, zero value otherwise.

### SetTable

`func (o *RestoreItem) SetTable(v string)`

SetTable sets Table field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


