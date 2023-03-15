# ClientCACertInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**ClientCACertStatus**](ClientCACertStatus.md) |  | [optional] 
**X509PemCert** | Pointer to **string** |  | [optional] 

## Methods

### NewClientCACertInfo

`func NewClientCACertInfo() *ClientCACertInfo`

NewClientCACertInfo instantiates a new ClientCACertInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetStatus

`func (o *ClientCACertInfo) GetStatus() ClientCACertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *ClientCACertInfo) SetStatus(v ClientCACertStatus)`

SetStatus sets Status field to given value.

### GetX509PemCert

`func (o *ClientCACertInfo) GetX509PemCert() string`

GetX509PemCert returns the X509PemCert field if non-nil, zero value otherwise.

### SetX509PemCert

`func (o *ClientCACertInfo) SetX509PemCert(v string)`

SetX509PemCert sets X509PemCert field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


