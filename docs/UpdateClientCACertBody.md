# UpdateClientCACertBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateMode** | Pointer to [**ClientCACertUpdateModeType**](ClientCACertUpdateModeType.md) |  | [optional] 
**X509PemCert** | Pointer to **string** | x509_pem_cert is the PEM-encoded X.509 CA certificate to apply. | [optional] 

## Methods

### NewUpdateClientCACertBody

`func NewUpdateClientCACertBody() *UpdateClientCACertBody`

NewUpdateClientCACertBody instantiates a new UpdateClientCACertBody object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetUpdateMode

`func (o *UpdateClientCACertBody) GetUpdateMode() ClientCACertUpdateModeType`

GetUpdateMode returns the UpdateMode field if non-nil, zero value otherwise.

### SetUpdateMode

`func (o *UpdateClientCACertBody) SetUpdateMode(v ClientCACertUpdateModeType)`

SetUpdateMode sets UpdateMode field to given value.

### GetX509PemCert

`func (o *UpdateClientCACertBody) GetX509PemCert() string`

GetX509PemCert returns the X509PemCert field if non-nil, zero value otherwise.

### SetX509PemCert

`func (o *UpdateClientCACertBody) SetX509PemCert(v string)`

SetX509PemCert sets X509PemCert field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


