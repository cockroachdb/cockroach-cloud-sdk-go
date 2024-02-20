# Database

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**TableCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewDatabase

`func NewDatabase(name string, ) *Database`

NewDatabase instantiates a new Database object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewDatabaseWithDefaults

`func NewDatabaseWithDefaults() *Database`

NewDatabaseWithDefaults instantiates a new Database object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetName

`func (o *Database) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### SetName

`func (o *Database) SetName(v string)`

SetName sets Name field to given value.

### GetTableCount

`func (o *Database) GetTableCount() int64`

GetTableCount returns the TableCount field if non-nil, zero value otherwise.

### SetTableCount

`func (o *Database) SetTableCount(v int64)`

SetTableCount sets TableCount field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


