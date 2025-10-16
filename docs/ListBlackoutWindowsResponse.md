# ListBlackoutWindowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlackoutWindows** | Pointer to [**[]BlackoutWindow**](BlackoutWindow.md) |  | [optional] 
**Pagination** | Pointer to [**KeysetPaginationResponse**](KeysetPaginationResponse.md) |  | [optional] 

## Methods

### NewListBlackoutWindowsResponse

`func NewListBlackoutWindowsResponse() *ListBlackoutWindowsResponse`

NewListBlackoutWindowsResponse instantiates a new ListBlackoutWindowsResponse object.
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed.

### GetBlackoutWindows

`func (o *ListBlackoutWindowsResponse) GetBlackoutWindows() []BlackoutWindow`

GetBlackoutWindows returns the BlackoutWindows field if non-nil, zero value otherwise.

### SetBlackoutWindows

`func (o *ListBlackoutWindowsResponse) SetBlackoutWindows(v []BlackoutWindow)`

SetBlackoutWindows sets BlackoutWindows field to given value.

### GetPagination

`func (o *ListBlackoutWindowsResponse) GetPagination() KeysetPaginationResponse`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### SetPagination

`func (o *ListBlackoutWindowsResponse) SetPagination(v KeysetPaginationResponse)`

SetPagination sets Pagination field to given value.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


