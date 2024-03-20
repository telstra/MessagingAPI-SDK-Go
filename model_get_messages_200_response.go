/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetMessages200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessages200Response{}

// GetMessages200Response struct for GetMessages200Response
type GetMessages200Response struct {
	// The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.
	Messages []MessageGet                  `json:"messages,omitempty"`
	Paging   *GetMessages200ResponsePaging `json:"paging,omitempty"`
}

// NewGetMessages200Response instantiates a new GetMessages200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessages200Response() *GetMessages200Response {
	this := GetMessages200Response{}
	return &this
}

// NewGetMessages200ResponseWithDefaults instantiates a new GetMessages200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessages200ResponseWithDefaults() *GetMessages200Response {
	this := GetMessages200Response{}
	return &this
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *GetMessages200Response) GetMessages() []MessageGet {
	if o == nil || IsNil(o.Messages) {
		var ret []MessageGet
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200Response) GetMessagesOk() ([]MessageGet, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *GetMessages200Response) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MessageGet and assigns it to the Messages field.
func (o *GetMessages200Response) SetMessages(v []MessageGet) {
	o.Messages = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *GetMessages200Response) GetPaging() GetMessages200ResponsePaging {
	if o == nil || IsNil(o.Paging) {
		var ret GetMessages200ResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200Response) GetPagingOk() (*GetMessages200ResponsePaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *GetMessages200Response) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given GetMessages200ResponsePaging and assigns it to the Paging field.
func (o *GetMessages200Response) SetPaging(v GetMessages200ResponsePaging) {
	o.Paging = &v
}

func (o GetMessages200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMessages200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableGetMessages200Response struct {
	value *GetMessages200Response
	isSet bool
}

func (v NullableGetMessages200Response) Get() *GetMessages200Response {
	return v.value
}

func (v *NullableGetMessages200Response) Set(val *GetMessages200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessages200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessages200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessages200Response(val *GetMessages200Response) *NullableGetMessages200Response {
	return &NullableGetMessages200Response{value: val, isSet: true}
}

func (v NullableGetMessages200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMessages200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
