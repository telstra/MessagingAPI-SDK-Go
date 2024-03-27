# MultimediaGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of multimedia content (image, audio or video) followed by the type (e.g. jpeg, png, text). | 
**FileName** | **string** | The name of the multimedia file. | 

## Methods

### NewMultimediaGet

`func NewMultimediaGet(type_ string, fileName string, ) *MultimediaGet`

NewMultimediaGet instantiates a new MultimediaGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultimediaGetWithDefaults

`func NewMultimediaGetWithDefaults() *MultimediaGet`

NewMultimediaGetWithDefaults instantiates a new MultimediaGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MultimediaGet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MultimediaGet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MultimediaGet) SetType(v string)`

SetType sets Type field to given value.


### GetFileName

`func (o *MultimediaGet) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *MultimediaGet) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *MultimediaGet) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


