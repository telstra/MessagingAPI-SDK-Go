# GetNumbers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualNumbers** | Pointer to [**[]VirtualNumber**](VirtualNumber.md) | The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.  | [optional] 
**Paging** | Pointer to [**Object**](Object.md) |  | [optional] 

## Methods

### NewGetNumbers200Response

`func NewGetNumbers200Response() *GetNumbers200Response`

NewGetNumbers200Response instantiates a new GetNumbers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNumbers200ResponseWithDefaults

`func NewGetNumbers200ResponseWithDefaults() *GetNumbers200Response`

NewGetNumbers200ResponseWithDefaults instantiates a new GetNumbers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualNumbers

`func (o *GetNumbers200Response) GetVirtualNumbers() []VirtualNumber`

GetVirtualNumbers returns the VirtualNumbers field if non-nil, zero value otherwise.

### GetVirtualNumbersOk

`func (o *GetNumbers200Response) GetVirtualNumbersOk() (*[]VirtualNumber, bool)`

GetVirtualNumbersOk returns a tuple with the VirtualNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualNumbers

`func (o *GetNumbers200Response) SetVirtualNumbers(v []VirtualNumber)`

SetVirtualNumbers sets VirtualNumbers field to given value.

### HasVirtualNumbers

`func (o *GetNumbers200Response) HasVirtualNumbers() bool`

HasVirtualNumbers returns a boolean if a field has been set.

### GetPaging

`func (o *GetNumbers200Response) GetPaging() Object`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetNumbers200Response) GetPagingOk() (*Object, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetNumbers200Response) SetPaging(v Object)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetNumbers200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


