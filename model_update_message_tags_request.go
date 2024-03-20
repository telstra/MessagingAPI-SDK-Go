/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the UpdateMessageTagsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessageTagsRequest{}

// UpdateMessageTagsRequest struct for UpdateMessageTagsRequest
type UpdateMessageTagsRequest struct {
	// Write the updated list of tag(s) here. You can assign up to 10 tags per message.  Note that if you provide an empty array, any pre-existing tags will be wiped.
	Tags []string `json:"tags"`
}

// NewUpdateMessageTagsRequest instantiates a new UpdateMessageTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessageTagsRequest(tags []string) *UpdateMessageTagsRequest {
	this := UpdateMessageTagsRequest{}
	this.Tags = tags
	return &this
}

// NewUpdateMessageTagsRequestWithDefaults instantiates a new UpdateMessageTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessageTagsRequestWithDefaults() *UpdateMessageTagsRequest {
	this := UpdateMessageTagsRequest{}
	return &this
}

// GetTags returns the Tags field value
func (o *UpdateMessageTagsRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UpdateMessageTagsRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *UpdateMessageTagsRequest) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateMessageTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMessageTagsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableUpdateMessageTagsRequest struct {
	value *UpdateMessageTagsRequest
	isSet bool
}

func (v NullableUpdateMessageTagsRequest) Get() *UpdateMessageTagsRequest {
	return v.value
}

func (v *NullableUpdateMessageTagsRequest) Set(val *UpdateMessageTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessageTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessageTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessageTagsRequest(val *UpdateMessageTagsRequest) *NullableUpdateMessageTagsRequest {
	return &NullableUpdateMessageTagsRequest{value: val, isSet: true}
}

func (v NullableUpdateMessageTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMessageTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
