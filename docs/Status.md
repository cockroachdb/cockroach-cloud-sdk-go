# Status

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]Any**](Any.md) |  | [optional] 

## Methods

### NewStatus

`func NewStatus() *Status`

NewStatus instantiates a new Status object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetCode

`func (o *Status) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### SetCode

`func (o *Status) SetCode(v int32)`

SetCode sets Code field to given value.

### GetMessage

`func (o *Status) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### SetMessage

`func (o *Status) SetMessage(v string)`

SetMessage sets Message field to given value.

### GetDetails

`func (o *Status) GetDetails() []Any`

GetDetails returns the Details field if non-nil, zero value otherwise.

### SetDetails

`func (o *Status) SetDetails(v []Any)`

SetDetails sets Details field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


