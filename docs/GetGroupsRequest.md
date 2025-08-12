# GetGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** | The maximum number of resources to return. If omitted, defaults to 20. If set to 0, the response will contain no resources but will include metadata such as &#x60;totalResults&#x60;, complying with [RFC 7644, Section 3.4.2.4: Pagination](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.4). | [optional] 
**ExcludedAttributes** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**StartIndex** | Pointer to **int32** | The 1-based index of the first resource to return in the response. If omitted or less than 1, defaults to 1. This behavior complies with [RFC 7644, Section 3.4.2.4: Pagination](https://datatracker.ietf.org/doc/html/rfc7644#section-3.4.2.4). | [optional] 

## Methods

### NewGetGroupsRequest

`func NewGetGroupsRequest() *GetGroupsRequest`

NewGetGroupsRequest instantiates a new GetGroupsRequest object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetAttributes

`func (o *GetGroupsRequest) GetAttributes() string`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### SetAttributes

`func (o *GetGroupsRequest) SetAttributes(v string)`

SetAttributes sets Attributes field to given value.

### GetCount

`func (o *GetGroupsRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### SetCount

`func (o *GetGroupsRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### GetExcludedAttributes

`func (o *GetGroupsRequest) GetExcludedAttributes() string`

GetExcludedAttributes returns the ExcludedAttributes field if non-nil, zero value otherwise.

### SetExcludedAttributes

`func (o *GetGroupsRequest) SetExcludedAttributes(v string)`

SetExcludedAttributes sets ExcludedAttributes field to given value.

### GetFilter

`func (o *GetGroupsRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### SetFilter

`func (o *GetGroupsRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### GetStartIndex

`func (o *GetGroupsRequest) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### SetStartIndex

`func (o *GetGroupsRequest) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


