/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the Errors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Errors{}

// Errors struct for Errors
type Errors struct {
	Errors []Error `json:"errors,omitempty"`
}

// NewErrors instantiates a new Errors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrors() *Errors {
	this := Errors{}
	return &this
}

// NewErrorsWithDefaults instantiates a new Errors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorsWithDefaults() *Errors {
	this := Errors{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Errors) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Errors) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Errors) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *Errors) SetErrors(v []Error) {
	o.Errors = v
}

func (o Errors) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Errors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableErrors struct {
	value *Errors
	isSet bool
}

func (v NullableErrors) Get() *Errors {
	return v.value
}

func (v *NullableErrors) Set(val *Errors) {
	v.value = val
	v.isSet = true
}

func (v NullableErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrors(val *Errors) *NullableErrors {
	return &NullableErrors{value: val, isSet: true}
}

func (v NullableErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
