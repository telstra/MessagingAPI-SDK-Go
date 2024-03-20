/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the AuthToken400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthToken400Response{}

// AuthToken400Response OAuth Error Responses
type AuthToken400Response struct {
	// Unique code of the error
	Error string `json:"error"`
	// Description of the error
	ErrorDescription *string `json:"error_description,omitempty"`
}

// NewAuthToken400Response instantiates a new AuthToken400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthToken400Response(error_ string) *AuthToken400Response {
	this := AuthToken400Response{}
	this.Error = error_
	return &this
}

// NewAuthToken400ResponseWithDefaults instantiates a new AuthToken400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthToken400ResponseWithDefaults() *AuthToken400Response {
	this := AuthToken400Response{}
	return &this
}

// GetError returns the Error field value
func (o *AuthToken400Response) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *AuthToken400Response) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *AuthToken400Response) SetError(v string) {
	o.Error = v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *AuthToken400Response) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthToken400Response) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *AuthToken400Response) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *AuthToken400Response) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

func (o AuthToken400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthToken400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	if !IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	return toSerialize, nil
}

type NullableAuthToken400Response struct {
	value *AuthToken400Response
	isSet bool
}

func (v NullableAuthToken400Response) Get() *AuthToken400Response {
	return v.value
}

func (v *NullableAuthToken400Response) Set(val *AuthToken400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthToken400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthToken400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthToken400Response(val *AuthToken400Response) *NullableAuthToken400Response {
	return &NullableAuthToken400Response{value: val, isSet: true}
}

func (v NullableAuthToken400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthToken400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
