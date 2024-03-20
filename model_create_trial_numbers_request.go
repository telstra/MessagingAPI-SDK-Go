/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the CreateTrialNumbersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTrialNumbersRequest{}

// CreateTrialNumbersRequest struct for CreateTrialNumbersRequest
type CreateTrialNumbersRequest struct {
	FreeTrialNumbers CreateTrialNumbersRequestFreeTrialNumbers `json:"freeTrialNumbers"`
}

// NewCreateTrialNumbersRequest instantiates a new CreateTrialNumbersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTrialNumbersRequest(freeTrialNumbers CreateTrialNumbersRequestFreeTrialNumbers) *CreateTrialNumbersRequest {
	this := CreateTrialNumbersRequest{}
	this.FreeTrialNumbers = freeTrialNumbers
	return &this
}

// NewCreateTrialNumbersRequestWithDefaults instantiates a new CreateTrialNumbersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTrialNumbersRequestWithDefaults() *CreateTrialNumbersRequest {
	this := CreateTrialNumbersRequest{}
	return &this
}

// GetFreeTrialNumbers returns the FreeTrialNumbers field value
func (o *CreateTrialNumbersRequest) GetFreeTrialNumbers() CreateTrialNumbersRequestFreeTrialNumbers {
	if o == nil {
		var ret CreateTrialNumbersRequestFreeTrialNumbers
		return ret
	}

	return o.FreeTrialNumbers
}

// GetFreeTrialNumbersOk returns a tuple with the FreeTrialNumbers field value
// and a boolean to check if the value has been set.
func (o *CreateTrialNumbersRequest) GetFreeTrialNumbersOk() (*CreateTrialNumbersRequestFreeTrialNumbers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeTrialNumbers, true
}

// SetFreeTrialNumbers sets field value
func (o *CreateTrialNumbersRequest) SetFreeTrialNumbers(v CreateTrialNumbersRequestFreeTrialNumbers) {
	o.FreeTrialNumbers = v
}

func (o CreateTrialNumbersRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTrialNumbersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["freeTrialNumbers"] = o.FreeTrialNumbers
	return toSerialize, nil
}

type NullableCreateTrialNumbersRequest struct {
	value *CreateTrialNumbersRequest
	isSet bool
}

func (v NullableCreateTrialNumbersRequest) Get() *CreateTrialNumbersRequest {
	return v.value
}

func (v *NullableCreateTrialNumbersRequest) Set(val *CreateTrialNumbersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTrialNumbersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTrialNumbersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTrialNumbersRequest(val *CreateTrialNumbersRequest) *NullableCreateTrialNumbersRequest {
	return &NullableCreateTrialNumbersRequest{value: val, isSet: true}
}

func (v NullableCreateTrialNumbersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTrialNumbersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
