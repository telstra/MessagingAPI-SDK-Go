# Multimedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | the type of multimedia content file you&#39;re sending (image, audio or video) followed by the file type. Use the format \&quot;multimedia type/file type\&quot;, e.g. \&quot;image/PNG\&quot; or \&quot;audio/MP3\&quot;. Supported file types:JPEG, BMP, GIF87a, GIF89a, PNG, MP3, WAV, MPEG, MPG, MP4, 3GP and US-ASCII. | 
**FileName** | **string** | The name of the multimedia file. | 
**Payload** | **string** | The base64 encoded content. You can use [this online tool](https://elmah.io/tools/base64-image-encoder/) to encode an image, or [Base64 Guru](https://base64.guru/) to encode a video or audio file.  | 

## Methods

### NewMultimedia

`func NewMultimedia(type_ string, fileName string, payload string, ) *Multimedia`

NewMultimedia instantiates a new Multimedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultimediaWithDefaults

`func NewMultimediaWithDefaults() *Multimedia`

NewMultimediaWithDefaults instantiates a new Multimedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Multimedia) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Multimedia) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Multimedia) SetType(v string)`

SetType sets Type field to given value.


### GetFileName

`func (o *Multimedia) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Multimedia) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Multimedia) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetPayload

`func (o *Multimedia) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Multimedia) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Multimedia) SetPayload(v string)`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


