/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the HealthCheck200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheck200Response{}

// HealthCheck200Response struct for HealthCheck200Response
type HealthCheck200Response struct {
	Status *string `json:"status,omitempty"`
}

// NewHealthCheck200Response instantiates a new HealthCheck200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheck200Response() *HealthCheck200Response {
	this := HealthCheck200Response{}
	return &this
}

// NewHealthCheck200ResponseWithDefaults instantiates a new HealthCheck200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheck200ResponseWithDefaults() *HealthCheck200Response {
	this := HealthCheck200Response{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheck200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheck200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheck200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *HealthCheck200Response) SetStatus(v string) {
	o.Status = &v
}

func (o HealthCheck200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheck200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableHealthCheck200Response struct {
	value *HealthCheck200Response
	isSet bool
}

func (v NullableHealthCheck200Response) Get() *HealthCheck200Response {
	return v.value
}

func (v *NullableHealthCheck200Response) Set(val *HealthCheck200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheck200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheck200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheck200Response(val *HealthCheck200Response) *NullableHealthCheck200Response {
	return &NullableHealthCheck200Response{value: val, isSet: true}
}

func (v NullableHealthCheck200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheck200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
