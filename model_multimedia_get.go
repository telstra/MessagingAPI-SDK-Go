/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the MultimediaGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultimediaGet{}

// MultimediaGet struct for MultimediaGet
type MultimediaGet struct {
	// The type of multimedia content (image, audio or video) followed by the type (e.g. jpeg, png, text).
	Type string `json:"type"`
	// The name of the multimedia file.
	FileName string `json:"fileName"`
}

// NewMultimediaGet instantiates a new MultimediaGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultimediaGet(type_ string, fileName string) *MultimediaGet {
	this := MultimediaGet{}
	this.Type = type_
	this.FileName = fileName
	return &this
}

// NewMultimediaGetWithDefaults instantiates a new MultimediaGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultimediaGetWithDefaults() *MultimediaGet {
	this := MultimediaGet{}
	return &this
}

// GetType returns the Type field value
func (o *MultimediaGet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MultimediaGet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MultimediaGet) SetType(v string) {
	o.Type = v
}

// GetFileName returns the FileName field value
func (o *MultimediaGet) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *MultimediaGet) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *MultimediaGet) SetFileName(v string) {
	o.FileName = v
}

func (o MultimediaGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultimediaGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableMultimediaGet struct {
	value *MultimediaGet
	isSet bool
}

func (v NullableMultimediaGet) Get() *MultimediaGet {
	return v.value
}

func (v *NullableMultimediaGet) Set(val *MultimediaGet) {
	v.value = val
	v.isSet = true
}

func (v NullableMultimediaGet) IsSet() bool {
	return v.isSet
}

func (v *NullableMultimediaGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultimediaGet(val *MultimediaGet) *NullableMultimediaGet {
	return &NullableMultimediaGet{value: val, isSet: true}
}

func (v NullableMultimediaGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultimediaGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
