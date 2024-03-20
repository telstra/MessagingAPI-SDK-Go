# GetMessages200ResponsePaging

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPage** | Pointer to **string** | The next page of results. | [optional] 
**PreviousPage** | Pointer to **string** | The previous page of results. | [optional] 
**LastPage** | Pointer to **string** | The last page of results. | [optional] 
**TotalCount** | Pointer to **float32** | The total number of pages. | [optional] 

## Methods

### NewGetMessages200ResponsePaging

`func NewGetMessages200ResponsePaging() *GetMessages200ResponsePaging`

NewGetMessages200ResponsePaging instantiates a new GetMessages200ResponsePaging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessages200ResponsePagingWithDefaults

`func NewGetMessages200ResponsePagingWithDefaults() *GetMessages200ResponsePaging`

NewGetMessages200ResponsePagingWithDefaults instantiates a new GetMessages200ResponsePaging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPage

`func (o *GetMessages200ResponsePaging) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *GetMessages200ResponsePaging) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *GetMessages200ResponsePaging) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *GetMessages200ResponsePaging) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetPreviousPage

`func (o *GetMessages200ResponsePaging) GetPreviousPage() string`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *GetMessages200ResponsePaging) GetPreviousPageOk() (*string, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *GetMessages200ResponsePaging) SetPreviousPage(v string)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *GetMessages200ResponsePaging) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### GetLastPage

`func (o *GetMessages200ResponsePaging) GetLastPage() string`

GetLastPage returns the LastPage field if non-nil, zero value otherwise.

### GetLastPageOk

`func (o *GetMessages200ResponsePaging) GetLastPageOk() (*string, bool)`

GetLastPageOk returns a tuple with the LastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPage

`func (o *GetMessages200ResponsePaging) SetLastPage(v string)`

SetLastPage sets LastPage field to given value.

### HasLastPage

`func (o *GetMessages200ResponsePaging) HasLastPage() bool`

HasLastPage returns a boolean if a field has been set.

### GetTotalCount

`func (o *GetMessages200ResponsePaging) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetMessages200ResponsePaging) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetMessages200ResponsePaging) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GetMessages200ResponsePaging) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


