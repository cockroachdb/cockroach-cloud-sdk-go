# CMEKKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] [default to UNKNOWN_STATUS]
**Spec** | Pointer to [**CMEKKeySpecification**](CMEKKeySpecification.md) |  | [optional] 

## Methods

### NewCMEKKeyInfo

`func NewCMEKKeyInfo() *CMEKKeyInfo`

NewCMEKKeyInfo instantiates a new CMEKKeyInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKKeyInfoWithDefaults

`func NewCMEKKeyInfoWithDefaults() *CMEKKeyInfo`

NewCMEKKeyInfoWithDefaults instantiates a new CMEKKeyInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CMEKKeyInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CMEKKeyInfo) GetStatusOk() (*CMEKStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CMEKKeyInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CMEKKeyInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSpec

`func (o *CMEKKeyInfo) GetSpec() CMEKKeySpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CMEKKeyInfo) GetSpecOk() (*CMEKKeySpecification, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CMEKKeyInfo) SetSpec(v CMEKKeySpecification)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *CMEKKeyInfo) HasSpec() bool`

HasSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


