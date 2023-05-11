# ScimEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Display** | Pointer to **string** |  | [optional] 
**Primary** | **bool** |  | 
**Type** | Pointer to **string** |  | [optional] 
**Value** | **string** |  | 

## Methods

### NewScimEmail

`func NewScimEmail(primary bool, value string, ) *ScimEmail`

NewScimEmail instantiates a new ScimEmail object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewScimEmailWithDefaults

`func NewScimEmailWithDefaults() *ScimEmail`

NewScimEmailWithDefaults instantiates a new ScimEmail object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDisplay

`func (o *ScimEmail) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### SetDisplay

`func (o *ScimEmail) SetDisplay(v string)`

SetDisplay sets Display field to given value.

### GetPrimary

`func (o *ScimEmail) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### SetPrimary

`func (o *ScimEmail) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### GetType

`func (o *ScimEmail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### SetType

`func (o *ScimEmail) SetType(v string)`

SetType sets Type field to given value.

### GetValue

`func (o *ScimEmail) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### SetValue

`func (o *ScimEmail) SetValue(v string)`

SetValue sets Value field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


