/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the UpdateNumberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNumberRequest{}

// UpdateNumberRequest struct for UpdateNumberRequest
type UpdateNumberRequest struct {
	// Tell us the URL that replies to the Virtual Number should be sent to.  Note that if you don't include this field, any pre-existing replyCallbackUrl will be wiped.  Sample callback response:  <pre><code class=\"language-sh\">{   \"to\":\"0476543210\",    \"from\":\"0401234567\",    \"timestamp\":\"2022-11-10T05:06:42.823Z\",   \"messageId\":\"75f263c0-60b5-11ed-8456-71ae4c63550d\",    \"messageContent\":\"Hi, example message\",    \"multimedia\": {      \"fileName:\"image.jpeg\"      \"type:\"image/jpeg\"      \"payload\":\"base64 payload\"    } }</code></pre>
	ReplyCallbackUrl *string `json:"replyCallbackUrl,omitempty"`
	// Create your own tags and use them to fetch, sort and report on your Virtual Numbers through our other endpoints. You can assign up to 10 tags per number.   Note that if you don't include this field, any pre-existing tags will be wiped.
	Tags []string `json:"tags,omitempty"`
}

// NewUpdateNumberRequest instantiates a new UpdateNumberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNumberRequest() *UpdateNumberRequest {
	this := UpdateNumberRequest{}
	return &this
}

// NewUpdateNumberRequestWithDefaults instantiates a new UpdateNumberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNumberRequestWithDefaults() *UpdateNumberRequest {
	this := UpdateNumberRequest{}
	return &this
}

// GetReplyCallbackUrl returns the ReplyCallbackUrl field value if set, zero value otherwise.
func (o *UpdateNumberRequest) GetReplyCallbackUrl() string {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		var ret string
		return ret
	}
	return *o.ReplyCallbackUrl
}

// GetReplyCallbackUrlOk returns a tuple with the ReplyCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNumberRequest) GetReplyCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		return nil, false
	}
	return o.ReplyCallbackUrl, true
}

// HasReplyCallbackUrl returns a boolean if a field has been set.
func (o *UpdateNumberRequest) HasReplyCallbackUrl() bool {
	if o != nil && !IsNil(o.ReplyCallbackUrl) {
		return true
	}

	return false
}

// SetReplyCallbackUrl gets a reference to the given string and assigns it to the ReplyCallbackUrl field.
func (o *UpdateNumberRequest) SetReplyCallbackUrl(v string) {
	o.ReplyCallbackUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateNumberRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNumberRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateNumberRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateNumberRequest) SetTags(v []string) {
	o.Tags = v
}

func (o UpdateNumberRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNumberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReplyCallbackUrl) {
		toSerialize["replyCallbackUrl"] = o.ReplyCallbackUrl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableUpdateNumberRequest struct {
	value *UpdateNumberRequest
	isSet bool
}

func (v NullableUpdateNumberRequest) Get() *UpdateNumberRequest {
	return v.value
}

func (v *NullableUpdateNumberRequest) Set(val *UpdateNumberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNumberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNumberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNumberRequest(val *UpdateNumberRequest) *NullableUpdateNumberRequest {
	return &NullableUpdateNumberRequest{value: val, isSet: true}
}

func (v NullableUpdateNumberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNumberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
