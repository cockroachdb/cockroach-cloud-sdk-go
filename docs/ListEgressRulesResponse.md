# ListEgressRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 
**Rules** | Pointer to [**[]EgressRule**](EgressRule.md) | rules are the egress rules associated with the given CockroachDB cluster. | [optional] 

## Methods

### NewListEgressRulesResponse

`func NewListEgressRulesResponse() *ListEgressRulesResponse`

NewListEgressRulesResponse instantiates a new ListEgressRulesResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetPagination

`func (o *ListEgressRulesResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListEgressRulesResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.

### GetRules

`func (o *ListEgressRulesResponse) GetRules() []EgressRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### SetRules

`func (o *ListEgressRulesResponse) SetRules(v []EgressRule)`

SetRules sets Rules field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


