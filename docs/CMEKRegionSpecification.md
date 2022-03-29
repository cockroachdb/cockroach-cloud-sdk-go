# CMEKRegionSpecification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**KeySpec** | Pointer to [**CMEKKeySpecification**](CMEKKeySpecification.md) |  | [optional] 

## Methods

### NewCMEKRegionSpecification

`func NewCMEKRegionSpecification() *CMEKRegionSpecification`

NewCMEKRegionSpecification instantiates a new CMEKRegionSpecification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCMEKRegionSpecificationWithDefaults

`func NewCMEKRegionSpecificationWithDefaults() *CMEKRegionSpecification`

NewCMEKRegionSpecificationWithDefaults instantiates a new CMEKRegionSpecification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *CMEKRegionSpecification) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CMEKRegionSpecification) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CMEKRegionSpecification) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CMEKRegionSpecification) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetKeySpec

`func (o *CMEKRegionSpecification) GetKeySpec() CMEKKeySpecification`

GetKeySpec returns the KeySpec field if non-nil, zero value otherwise.

### GetKeySpecOk

`func (o *CMEKRegionSpecification) GetKeySpecOk() (*CMEKKeySpecification, bool)`

GetKeySpecOk returns a tuple with the KeySpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySpec

`func (o *CMEKRegionSpecification) SetKeySpec(v CMEKKeySpecification)`

SetKeySpec sets KeySpec field to given value.

### HasKeySpec

`func (o *CMEKRegionSpecification) HasKeySpec() bool`

HasKeySpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


