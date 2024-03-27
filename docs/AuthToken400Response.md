# AuthToken400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Unique code of the error | 
**ErrorDescription** | Pointer to **string** | Description of the error | [optional] 

## Methods

### NewAuthToken400Response

`func NewAuthToken400Response(error_ string, ) *AuthToken400Response`

NewAuthToken400Response instantiates a new AuthToken400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthToken400ResponseWithDefaults

`func NewAuthToken400ResponseWithDefaults() *AuthToken400Response`

NewAuthToken400ResponseWithDefaults instantiates a new AuthToken400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AuthToken400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthToken400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthToken400Response) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *AuthToken400Response) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *AuthToken400Response) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *AuthToken400Response) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *AuthToken400Response) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


