/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
	"time"
)

// checks if the VirtualNumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualNumber{}

// VirtualNumber struct for VirtualNumber
type VirtualNumber struct {
	// The Virtual Number assigned to your account.
	VirtualNumber *string `json:"virtualNumber,omitempty"`
	// The URL that replies to the Virtual Number will be posted to.
	ReplyCallbackUrl *string `json:"replyCallbackUrl,omitempty"`
	// Any customisable tags assigned to the Virtual Number.
	Tags []string `json:"tags,omitempty"`
	// The last time the Virtual Number was used to send a message.
	LastUse *time.Time `json:"lastUse,omitempty"`
}

// NewVirtualNumber instantiates a new VirtualNumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNumber() *VirtualNumber {
	this := VirtualNumber{}
	return &this
}

// NewVirtualNumberWithDefaults instantiates a new VirtualNumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNumberWithDefaults() *VirtualNumber {
	this := VirtualNumber{}
	return &this
}

// GetVirtualNumber returns the VirtualNumber field value if set, zero value otherwise.
func (o *VirtualNumber) GetVirtualNumber() string {
	if o == nil || IsNil(o.VirtualNumber) {
		var ret string
		return ret
	}
	return *o.VirtualNumber
}

// GetVirtualNumberOk returns a tuple with the VirtualNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNumber) GetVirtualNumberOk() (*string, bool) {
	if o == nil || IsNil(o.VirtualNumber) {
		return nil, false
	}
	return o.VirtualNumber, true
}

// HasVirtualNumber returns a boolean if a field has been set.
func (o *VirtualNumber) HasVirtualNumber() bool {
	if o != nil && !IsNil(o.VirtualNumber) {
		return true
	}

	return false
}

// SetVirtualNumber gets a reference to the given string and assigns it to the VirtualNumber field.
func (o *VirtualNumber) SetVirtualNumber(v string) {
	o.VirtualNumber = &v
}

// GetReplyCallbackUrl returns the ReplyCallbackUrl field value if set, zero value otherwise.
func (o *VirtualNumber) GetReplyCallbackUrl() string {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		var ret string
		return ret
	}
	return *o.ReplyCallbackUrl
}

// GetReplyCallbackUrlOk returns a tuple with the ReplyCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNumber) GetReplyCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReplyCallbackUrl) {
		return nil, false
	}
	return o.ReplyCallbackUrl, true
}

// HasReplyCallbackUrl returns a boolean if a field has been set.
func (o *VirtualNumber) HasReplyCallbackUrl() bool {
	if o != nil && !IsNil(o.ReplyCallbackUrl) {
		return true
	}

	return false
}

// SetReplyCallbackUrl gets a reference to the given string and assigns it to the ReplyCallbackUrl field.
func (o *VirtualNumber) SetReplyCallbackUrl(v string) {
	o.ReplyCallbackUrl = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VirtualNumber) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNumber) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VirtualNumber) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VirtualNumber) SetTags(v []string) {
	o.Tags = v
}

// GetLastUse returns the LastUse field value if set, zero value otherwise.
func (o *VirtualNumber) GetLastUse() time.Time {
	if o == nil || IsNil(o.LastUse) {
		var ret time.Time
		return ret
	}
	return *o.LastUse
}

// GetLastUseOk returns a tuple with the LastUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNumber) GetLastUseOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUse) {
		return nil, false
	}
	return o.LastUse, true
}

// HasLastUse returns a boolean if a field has been set.
func (o *VirtualNumber) HasLastUse() bool {
	if o != nil && !IsNil(o.LastUse) {
		return true
	}

	return false
}

// SetLastUse gets a reference to the given time.Time and assigns it to the LastUse field.
func (o *VirtualNumber) SetLastUse(v time.Time) {
	o.LastUse = &v
}

func (o VirtualNumber) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualNumber) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VirtualNumber) {
		toSerialize["virtualNumber"] = o.VirtualNumber
	}
	if !IsNil(o.ReplyCallbackUrl) {
		toSerialize["replyCallbackUrl"] = o.ReplyCallbackUrl
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.LastUse) {
		toSerialize["lastUse"] = o.LastUse
	}
	return toSerialize, nil
}

type NullableVirtualNumber struct {
	value *VirtualNumber
	isSet bool
}

func (v NullableVirtualNumber) Get() *VirtualNumber {
	return v.value
}

func (v *NullableVirtualNumber) Set(val *VirtualNumber) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNumber) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNumber(val *VirtualNumber) *NullableVirtualNumber {
	return &NullableVirtualNumber{value: val, isSet: true}
}

func (v NullableVirtualNumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
