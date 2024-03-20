/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the Multimedia type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Multimedia{}

// Multimedia struct for Multimedia
type Multimedia struct {
	// the type of multimedia content file you're sending (image, audio or video) followed by the file type. Use the format \"multimedia type/file type\", e.g. \"image/PNG\" or \"audio/MP3\". Supported file types:JPEG, BMP, GIF87a, GIF89a, PNG, MP3, WAV, MPEG, MPG, MP4, 3GP and US-ASCII.
	Type string `json:"type"`
	// The name of the multimedia file.
	FileName string `json:"fileName"`
	// The base64 encoded content. You can use [this online tool](https://elmah.io/tools/base64-image-encoder/) to encode an image, or [Base64 Guru](https://base64.guru/) to encode a video or audio file.
	Payload string `json:"payload"`
}

// NewMultimedia instantiates a new Multimedia object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultimedia(type_ string, fileName string, payload string) *Multimedia {
	this := Multimedia{}
	this.Type = type_
	this.FileName = fileName
	this.Payload = payload
	return &this
}

// NewMultimediaWithDefaults instantiates a new Multimedia object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultimediaWithDefaults() *Multimedia {
	this := Multimedia{}
	return &this
}

// GetType returns the Type field value
func (o *Multimedia) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Multimedia) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Multimedia) SetType(v string) {
	o.Type = v
}

// GetFileName returns the FileName field value
func (o *Multimedia) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *Multimedia) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *Multimedia) SetFileName(v string) {
	o.FileName = v
}

// GetPayload returns the Payload field value
func (o *Multimedia) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Multimedia) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *Multimedia) SetPayload(v string) {
	o.Payload = v
}

func (o Multimedia) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Multimedia) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["fileName"] = o.FileName
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullableMultimedia struct {
	value *Multimedia
	isSet bool
}

func (v NullableMultimedia) Get() *Multimedia {
	return v.value
}

func (v *NullableMultimedia) Set(val *Multimedia) {
	v.value = val
	v.isSet = true
}

func (v NullableMultimedia) IsSet() bool {
	return v.isSet
}

func (v *NullableMultimedia) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultimedia(val *Multimedia) *NullableMultimedia {
	return &NullableMultimedia{value: val, isSet: true}
}

func (v NullableMultimedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultimedia) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
