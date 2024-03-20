/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
	"time"
)

// checks if the RecipientOptout type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecipientOptout{}

// RecipientOptout struct for RecipientOptout
type RecipientOptout struct {
	// The mobile number that sent the optout request.
	OptoutNumber *string `json:"optoutNumber,omitempty"`
	// The date and time we received the optout request.
	CreateTimestamp *time.Time `json:"createTimestamp,omitempty"`
}

// NewRecipientOptout instantiates a new RecipientOptout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecipientOptout() *RecipientOptout {
	this := RecipientOptout{}
	return &this
}

// NewRecipientOptoutWithDefaults instantiates a new RecipientOptout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecipientOptoutWithDefaults() *RecipientOptout {
	this := RecipientOptout{}
	return &this
}

// GetOptoutNumber returns the OptoutNumber field value if set, zero value otherwise.
func (o *RecipientOptout) GetOptoutNumber() string {
	if o == nil || IsNil(o.OptoutNumber) {
		var ret string
		return ret
	}
	return *o.OptoutNumber
}

// GetOptoutNumberOk returns a tuple with the OptoutNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientOptout) GetOptoutNumberOk() (*string, bool) {
	if o == nil || IsNil(o.OptoutNumber) {
		return nil, false
	}
	return o.OptoutNumber, true
}

// HasOptoutNumber returns a boolean if a field has been set.
func (o *RecipientOptout) HasOptoutNumber() bool {
	if o != nil && !IsNil(o.OptoutNumber) {
		return true
	}

	return false
}

// SetOptoutNumber gets a reference to the given string and assigns it to the OptoutNumber field.
func (o *RecipientOptout) SetOptoutNumber(v string) {
	o.OptoutNumber = &v
}

// GetCreateTimestamp returns the CreateTimestamp field value if set, zero value otherwise.
func (o *RecipientOptout) GetCreateTimestamp() time.Time {
	if o == nil || IsNil(o.CreateTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.CreateTimestamp
}

// GetCreateTimestampOk returns a tuple with the CreateTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecipientOptout) GetCreateTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTimestamp) {
		return nil, false
	}
	return o.CreateTimestamp, true
}

// HasCreateTimestamp returns a boolean if a field has been set.
func (o *RecipientOptout) HasCreateTimestamp() bool {
	if o != nil && !IsNil(o.CreateTimestamp) {
		return true
	}

	return false
}

// SetCreateTimestamp gets a reference to the given time.Time and assigns it to the CreateTimestamp field.
func (o *RecipientOptout) SetCreateTimestamp(v time.Time) {
	o.CreateTimestamp = &v
}

func (o RecipientOptout) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecipientOptout) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptoutNumber) {
		toSerialize["optoutNumber"] = o.OptoutNumber
	}
	if !IsNil(o.CreateTimestamp) {
		toSerialize["createTimestamp"] = o.CreateTimestamp
	}
	return toSerialize, nil
}

type NullableRecipientOptout struct {
	value *RecipientOptout
	isSet bool
}

func (v NullableRecipientOptout) Get() *RecipientOptout {
	return v.value
}

func (v *NullableRecipientOptout) Set(val *RecipientOptout) {
	v.value = val
	v.isSet = true
}

func (v NullableRecipientOptout) IsSet() bool {
	return v.isSet
}

func (v *NullableRecipientOptout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecipientOptout(val *RecipientOptout) *NullableRecipientOptout {
	return &NullableRecipientOptout{value: val, isSet: true}
}

func (v NullableRecipientOptout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecipientOptout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
