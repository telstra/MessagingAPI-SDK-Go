/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the AssignNumberRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignNumberRequest{}

// AssignNumberRequest struct for AssignNumberRequest
type AssignNumberRequest struct {
	// Tell us the URL that replies to the Virtual Number should be sent to.  Sample callback response:  <pre><code class=\"language-sh\">{   \"to\":\"0476543210\",    \"from\":\"0401234567\",    \"timestamp\":\"2022-11-10T05:06:42.823Z\",    \"messageId\":\"75f263c0-60b5-11ed-8456-71ae4c63550d\",    \"messageContent\":\"Hi, example message\",    \"multimedia\": {      \"fileName:\"image.jpeg\"      \"type:\"image/jpeg\"      \"payload\":\"base64 payload\"   } }</code></pre>
	ReplyCallbackUrl *string `json:"replyCallbackUrl,omitempty"`
	// Create your own tags and use them to fetch, sort and report on your Virtual Numbers through our other endpoints. You can assign up to 10 tags per number.
	Tags []string `json:"tags,omitempty"`
}

// NewAssignNumberRequest instantiates a new AssignNumberRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignNumberRequest() *AssignNumberRequest {
	this := AssignNumberRequest{}
	return &this
}

// NewAssignNumberRequestWithDefaults instantiates a new AssignNumberRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignNumberRequestWithDefaults() *AssignNumberRequest {
	this := AssignNumberRequest{}
	return &this
}

// GetReplyCallbackUrl returns the ReplyCallbackUrl field value if set, zero value otherwise.
func (o *AssignNumberRequest) GetReplyCallbackUrl() string {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		var ret string
		return ret
	}
	return *o.ReplyCallbackUrl
}

// GetReplyCallbackUrlOk returns a tuple with the ReplyCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignNumberRequest) GetReplyCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		return nil, false
	}
	return o.ReplyCallbackUrl, true
}

// HasReplyCallbackUrl returns a boolean if a field has been set.
func (o *AssignNumberRequest) HasReplyCallbackUrl() bool {
	if o != nil && !IsNil(o.ReplyCallbackUrl) {
		return true
	}

	return false
}

// SetReplyCallbackUrl gets a reference to the given string and assigns it to the ReplyCallbackUrl field.
func (o *AssignNumberRequest) SetReplyCallbackUrl(v string) {
	o.ReplyCallbackUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AssignNumberRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignNumberRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AssignNumberRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AssignNumberRequest) SetTags(v []string) {
	o.Tags = v
}

func (o AssignNumberRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignNumberRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReplyCallbackUrl) {
		toSerialize["replyCallbackUrl"] = o.ReplyCallbackUrl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableAssignNumberRequest struct {
	value *AssignNumberRequest
	isSet bool
}

func (v NullableAssignNumberRequest) Get() *AssignNumberRequest {
	return v.value
}

func (v *NullableAssignNumberRequest) Set(val *AssignNumberRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignNumberRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignNumberRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignNumberRequest(val *AssignNumberRequest) *NullableAssignNumberRequest {
	return &NullableAssignNumberRequest{value: val, isSet: true}
}

func (v NullableAssignNumberRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignNumberRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
