# GetMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | Pointer to [**[]MessageGet**](MessageGet.md) | The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.  | [optional] 
**Paging** | Pointer to [**GetMessages200ResponsePaging**](GetMessages200ResponsePaging.md) |  | [optional] 

## Methods

### NewGetMessages200Response

`func NewGetMessages200Response() *GetMessages200Response`

NewGetMessages200Response instantiates a new GetMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMessages200ResponseWithDefaults

`func NewGetMessages200ResponseWithDefaults() *GetMessages200Response`

NewGetMessages200ResponseWithDefaults instantiates a new GetMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetMessages200Response) GetMessages() []MessageGet`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetMessages200Response) GetMessagesOk() (*[]MessageGet, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetMessages200Response) SetMessages(v []MessageGet)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetPaging

`func (o *GetMessages200Response) GetPaging() GetMessages200ResponsePaging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetMessages200Response) GetPagingOk() (*GetMessages200ResponsePaging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetMessages200Response) SetPaging(v GetMessages200ResponsePaging)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetMessages200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


