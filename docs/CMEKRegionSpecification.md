# CMEKRegionSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeySpec** | Pointer to [**CMEKKeySpecification**](CMEKKeySpecification.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewCMEKRegionSpecification

`func NewCMEKRegionSpecification() *CMEKRegionSpecification`

NewCMEKRegionSpecification instantiates a new CMEKRegionSpecification object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetKeySpec

`func (o *CMEKRegionSpecification) GetKeySpec() CMEKKeySpecification`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### SetKeySpec

`func (o *CMEKRegionSpecification) SetKeySpec(v CMEKKeySpecification)`

SetKeySpec sets KeySpec field to given value.

### GetRegion

`func (o *CMEKRegionSpecification) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### SetRegion

`func (o *CMEKRegionSpecification) SetRegion(v string)`

SetRegion sets Region field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


