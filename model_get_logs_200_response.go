/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetLogs200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLogs200Response{}

// GetLogs200Response struct for GetLogs200Response
type GetLogs200Response struct {
	Logs []Log `json:"logs,omitempty"`
}

// NewGetLogs200Response instantiates a new GetLogs200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLogs200Response() *GetLogs200Response {
	this := GetLogs200Response{}
	return &this
}

// NewGetLogs200ResponseWithDefaults instantiates a new GetLogs200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLogs200ResponseWithDefaults() *GetLogs200Response {
	this := GetLogs200Response{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *GetLogs200Response) GetLogs() []Log {
	if o == nil || IsNil(o.Logs) {
		var ret []Log
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLogs200Response) GetLogsOk() ([]Log, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *GetLogs200Response) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []Log and assigns it to the Logs field.
func (o *GetLogs200Response) SetLogs(v []Log) {
	o.Logs = v
}

func (o GetLogs200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLogs200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableGetLogs200Response struct {
	value *GetLogs200Response
	isSet bool
}

func (v NullableGetLogs200Response) Get() *GetLogs200Response {
	return v.value
}

func (v *NullableGetLogs200Response) Set(val *GetLogs200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLogs200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLogs200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLogs200Response(val *GetLogs200Response) *NullableGetLogs200Response {
	return &NullableGetLogs200Response{value: val, isSet: true}
}

func (v NullableGetLogs200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLogs200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
