# UpdateCMEKStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**CMEKCustomerAction**](CMEKCustomerAction.md) |  | [default to UNKNOWN_ACTION]

## Methods

### NewUpdateCMEKStatusRequest

`func NewUpdateCMEKStatusRequest(action CMEKCustomerAction, ) *UpdateCMEKStatusRequest`

NewUpdateCMEKStatusRequest instantiates a new UpdateCMEKStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCMEKStatusRequestWithDefaults

`func NewUpdateCMEKStatusRequestWithDefaults() *UpdateCMEKStatusRequest`

NewUpdateCMEKStatusRequestWithDefaults instantiates a new UpdateCMEKStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateCMEKStatusRequest) GetAction() CMEKCustomerAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateCMEKStatusRequest) GetActionOk() (*CMEKCustomerAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateCMEKStatusRequest) SetAction(v CMEKCustomerAction)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


