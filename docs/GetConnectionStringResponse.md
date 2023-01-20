# GetConnectionStringResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | Pointer to **string** | connection_string contains the full connection string, with parameters formatted inline. | [optional] 
**Params** | Pointer to **map[string]string** | params contains a list of individual key parameters, for generating nonstandard connection strings. | [optional] 

## Methods

### NewGetConnectionStringResponse

`func NewGetConnectionStringResponse() *GetConnectionStringResponse`

NewGetConnectionStringResponse instantiates a new GetConnectionStringResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetConnectionString

`func (o *GetConnectionStringResponse) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### SetConnectionString

`func (o *GetConnectionStringResponse) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### GetParams

`func (o *GetConnectionStringResponse) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### SetParams

`func (o *GetConnectionStringResponse) SetParams(v map[string]string)`

SetParams sets Params field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


