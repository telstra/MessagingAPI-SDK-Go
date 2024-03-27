# GetRecipientOptouts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientOptouts** | Pointer to [**[]RecipientOptout**](RecipientOptout.md) | The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.  | [optional] 
**Paging** | Pointer to [**Object**](Object.md) |  | [optional] 

## Methods

### NewGetRecipientOptouts200Response

`func NewGetRecipientOptouts200Response() *GetRecipientOptouts200Response`

NewGetRecipientOptouts200Response instantiates a new GetRecipientOptouts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecipientOptouts200ResponseWithDefaults

`func NewGetRecipientOptouts200ResponseWithDefaults() *GetRecipientOptouts200Response`

NewGetRecipientOptouts200ResponseWithDefaults instantiates a new GetRecipientOptouts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientOptouts

`func (o *GetRecipientOptouts200Response) GetRecipientOptouts() []RecipientOptout`

GetRecipientOptouts returns the RecipientOptouts field if non-nil, zero value otherwise.

### GetRecipientOptoutsOk

`func (o *GetRecipientOptouts200Response) GetRecipientOptoutsOk() (*[]RecipientOptout, bool)`

GetRecipientOptoutsOk returns a tuple with the RecipientOptouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientOptouts

`func (o *GetRecipientOptouts200Response) SetRecipientOptouts(v []RecipientOptout)`

SetRecipientOptouts sets RecipientOptouts field to given value.

### HasRecipientOptouts

`func (o *GetRecipientOptouts200Response) HasRecipientOptouts() bool`

HasRecipientOptouts returns a boolean if a field has been set.

### GetPaging

`func (o *GetRecipientOptouts200Response) GetPaging() Object`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *GetRecipientOptouts200Response) GetPagingOk() (*Object, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *GetRecipientOptouts200Response) SetPaging(v Object)`

SetPaging sets Paging field to given value.

### HasPaging

`func (o *GetRecipientOptouts200Response) HasPaging() bool`

HasPaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


