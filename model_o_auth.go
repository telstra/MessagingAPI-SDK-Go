/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the OAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuth{}

// OAuth struct for OAuth
type OAuth struct {
	// This is your OAuth 2.0 Authentication Token. It will be valid for one hour.
	AccessToken *string `json:"access_token,omitempty"`
	// This is the Authentication Token type.
	TokenType *string `json:"token_type,omitempty"`
	// This is when your token will expire.
	ExpiresIn *string `json:"expires_in,omitempty"`
}

// NewOAuth instantiates a new OAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuth() *OAuth {
	this := OAuth{}
	return &this
}

// NewOAuthWithDefaults instantiates a new OAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthWithDefaults() *OAuth {
	this := OAuth{}
	return &this
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *OAuth) GetAccessToken() string {
	if o == nil || IsNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth) GetAccessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *OAuth) HasAccessToken() bool {
	if o != nil && !IsNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *OAuth) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTokenType returns the TokenType field value if set, zero value otherwise.
func (o *OAuth) GetTokenType() string {
	if o == nil || IsNil(o.TokenType) {
		var ret string
		return ret
	}
	return *o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth) GetTokenTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TokenType) {
		return nil, false
	}
	return o.TokenType, true
}

// HasTokenType returns a boolean if a field has been set.
func (o *OAuth) HasTokenType() bool {
	if o != nil && !IsNil(o.TokenType) {
		return true
	}

	return false
}

// SetTokenType gets a reference to the given string and assigns it to the TokenType field.
func (o *OAuth) SetTokenType(v string) {
	o.TokenType = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *OAuth) GetExpiresIn() string {
	if o == nil || IsNil(o.ExpiresIn) {
		var ret string
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuth) GetExpiresInOk() (*string, bool) {
	if o == nil || IsNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *OAuth) HasExpiresIn() bool {
	if o != nil && !IsNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given string and assigns it to the ExpiresIn field.
func (o *OAuth) SetExpiresIn(v string) {
	o.ExpiresIn = &v
}

func (o OAuth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !IsNil(o.TokenType) {
		toSerialize["token_type"] = o.TokenType
	}
	if !IsNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	return toSerialize, nil
}

type NullableOAuth struct {
	value *OAuth
	isSet bool
}

func (v NullableOAuth) Get() *OAuth {
	return v.value
}

func (v *NullableOAuth) Set(val *OAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuth(val *OAuth) *NullableOAuth {
	return &NullableOAuth{value: val, isSet: true}
}

func (v NullableOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
