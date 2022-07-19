# CMEKKeyInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**CMEKStatus**](CMEKStatus.md) |  | [optional] 
**UserMessage** | Pointer to **string** |  | [optional] 
**Spec** | Pointer to [**CMEKKeySpecification**](CMEKKeySpecification.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCMEKKeyInfo

`func NewCMEKKeyInfo() *CMEKKeyInfo`

NewCMEKKeyInfo instantiates a new CMEKKeyInfo object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetStatus

`func (o *CMEKKeyInfo) GetStatus() CMEKStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### SetStatus

`func (o *CMEKKeyInfo) SetStatus(v CMEKStatus)`

SetStatus sets Status field to given value.

### GetUserMessage

`func (o *CMEKKeyInfo) GetUserMessage() string`

GetUserMessage returns the UserMessage field if non-nil, zero value otherwise.

### SetUserMessage

`func (o *CMEKKeyInfo) SetUserMessage(v string)`

SetUserMessage sets UserMessage field to given value.

### GetSpec

`func (o *CMEKKeyInfo) GetSpec() CMEKKeySpecification`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### SetSpec

`func (o *CMEKKeyInfo) SetSpec(v CMEKKeySpecification)`

SetSpec sets Spec field to given value.

### GetCreatedAt

`func (o *CMEKKeyInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### SetCreatedAt

`func (o *CMEKKeyInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### GetUpdatedAt

`func (o *CMEKKeyInfo) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### SetUpdatedAt

`func (o *CMEKKeyInfo) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


