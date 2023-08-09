# ConnectionStringParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | **string** |  | 
**Host** | **string** |  | 
**Port** | **string** |  | 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectionStringParameters

`func NewConnectionStringParameters(database string, host string, port string, ) *ConnectionStringParameters`

NewConnectionStringParameters instantiates a new ConnectionStringParameters object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewConnectionStringParametersWithDefaults

`func NewConnectionStringParametersWithDefaults() *ConnectionStringParameters`

NewConnectionStringParametersWithDefaults instantiates a new ConnectionStringParameters object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDatabase

`func (o *ConnectionStringParameters) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### SetDatabase

`func (o *ConnectionStringParameters) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### GetHost

`func (o *ConnectionStringParameters) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### SetHost

`func (o *ConnectionStringParameters) SetHost(v string)`

SetHost sets Host field to given value.

### GetPort

`func (o *ConnectionStringParameters) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### SetPort

`func (o *ConnectionStringParameters) SetPort(v string)`

SetPort sets Port field to given value.

### GetUsername

`func (o *ConnectionStringParameters) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### SetUsername

`func (o *ConnectionStringParameters) SetUsername(v string)`

SetUsername sets Username field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


