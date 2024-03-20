/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the FreeTrialNumbers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FreeTrialNumbers{}

// FreeTrialNumbers struct for FreeTrialNumbers
type FreeTrialNumbers struct {
	// The recipient number(s) registered to your Free Trial Numbers List. These are the mobile numbers that you can message during the Free Trial.
	FreeTrialNumbers []string `json:"freeTrialNumbers"`
}

// NewFreeTrialNumbers instantiates a new FreeTrialNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFreeTrialNumbers(freeTrialNumbers []string) *FreeTrialNumbers {
	this := FreeTrialNumbers{}
	this.FreeTrialNumbers = freeTrialNumbers
	return &this
}

// NewFreeTrialNumbersWithDefaults instantiates a new FreeTrialNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFreeTrialNumbersWithDefaults() *FreeTrialNumbers {
	this := FreeTrialNumbers{}
	return &this
}

// GetFreeTrialNumbers returns the FreeTrialNumbers field value
func (o *FreeTrialNumbers) GetFreeTrialNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FreeTrialNumbers
}

// GetFreeTrialNumbersOk returns a tuple with the FreeTrialNumbers field value
// and a boolean to check if the value has been set.
func (o *FreeTrialNumbers) GetFreeTrialNumbersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FreeTrialNumbers, true
}

// SetFreeTrialNumbers sets field value
func (o *FreeTrialNumbers) SetFreeTrialNumbers(v []string) {
	o.FreeTrialNumbers = v
}

func (o FreeTrialNumbers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FreeTrialNumbers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["freeTrialNumbers"] = o.FreeTrialNumbers
	return toSerialize, nil
}

type NullableFreeTrialNumbers struct {
	value *FreeTrialNumbers
	isSet bool
}

func (v NullableFreeTrialNumbers) Get() *FreeTrialNumbers {
	return v.value
}

func (v *NullableFreeTrialNumbers) Set(val *FreeTrialNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableFreeTrialNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableFreeTrialNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFreeTrialNumbers(val *FreeTrialNumbers) *NullableFreeTrialNumbers {
	return &NullableFreeTrialNumbers{value: val, isSet: true}
}

func (v NullableFreeTrialNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFreeTrialNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
