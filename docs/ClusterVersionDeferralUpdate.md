# ClusterVersionDeferralUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeferralPolicy** | [**ClusterVersionDeferralPolicyType**](ClusterVersionDeferralPolicyType.md) |  | 
**DeferredUntil** | Pointer to **time.Time** | Deprecated: This field is ignored and will be removed in a future version. | [optional] 

## Methods

### NewClusterVersionDeferralUpdate

`func NewClusterVersionDeferralUpdate(deferralPolicy ClusterVersionDeferralPolicyType, ) *ClusterVersionDeferralUpdate`

NewClusterVersionDeferralUpdate instantiates a new ClusterVersionDeferralUpdate object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### NewClusterVersionDeferralUpdateWithDefaults

`func NewClusterVersionDeferralUpdateWithDefaults() *ClusterVersionDeferralUpdate`

NewClusterVersionDeferralUpdateWithDefaults instantiates a new ClusterVersionDeferralUpdate object.
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set.

### GetDeferralPolicy

`func (o *ClusterVersionDeferralUpdate) GetDeferralPolicy() ClusterVersionDeferralPolicyType`

GetDeferralPolicy returns the DeferralPolicy field if non-nil, zero value otherwise.

### SetDeferralPolicy

`func (o *ClusterVersionDeferralUpdate) SetDeferralPolicy(v ClusterVersionDeferralPolicyType)`

SetDeferralPolicy sets DeferralPolicy field to given value.

### GetDeferredUntil

`func (o *ClusterVersionDeferralUpdate) GetDeferredUntil() time.Time`

GetDeferredUntil returns the DeferredUntil field if non-nil, zero value otherwise.

### SetDeferredUntil

`func (o *ClusterVersionDeferralUpdate) SetDeferredUntil(v time.Time)`

SetDeferredUntil sets DeferredUntil field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


