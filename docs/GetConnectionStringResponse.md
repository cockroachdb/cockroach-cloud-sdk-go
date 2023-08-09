# GetConnectionStringResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionString** | **string** | connection_string contains the full connection string with parameters formatted inline. | 
**Params** | [**ConnectionStringParameters**](ConnectionStringParameters.md) |  | 

## Methods

### NewGetConnectionStringResponse

`func NewGetConnectionStringResponse(connectionString string, params ConnectionStringParameters, ) *GetConnectionStringResponse`

NewGetConnectionStringResponse instantiates a new GetConnectionStringResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewGetConnectionStringResponseWithDefaults

`func NewGetConnectionStringResponseWithDefaults() *GetConnectionStringResponse`

NewGetConnectionStringResponseWithDefaults instantiates a new GetConnectionStringResponse object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetConnectionString

`func (o *GetConnectionStringResponse) GetConnectionString() string`

GetConnectionString returns the ConnectionString field if non-nil, zero value otherwise.

### SetConnectionString

`func (o *GetConnectionStringResponse) SetConnectionString(v string)`

SetConnectionString sets ConnectionString field to given value.

### GetParams

`func (o *GetConnectionStringResponse) GetParams() ConnectionStringParameters`

GetParams returns the Params field if non-nil, zero value otherwise.

### SetParams

`func (o *GetConnectionStringResponse) SetParams(v ConnectionStringParameters)`

SetParams sets Params field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


