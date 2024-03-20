/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetRecipientOptouts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecipientOptouts200Response{}

// GetRecipientOptouts200Response struct for GetRecipientOptouts200Response
type GetRecipientOptouts200Response struct {
	// The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.
	RecipientOptouts []RecipientOptout             `json:"recipientOptouts,omitempty"`
	Paging           *GetMessages200ResponsePaging `json:"paging,omitempty"`
}

// NewGetRecipientOptouts200Response instantiates a new GetRecipientOptouts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecipientOptouts200Response() *GetRecipientOptouts200Response {
	this := GetRecipientOptouts200Response{}
	return &this
}

// NewGetRecipientOptouts200ResponseWithDefaults instantiates a new GetRecipientOptouts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecipientOptouts200ResponseWithDefaults() *GetRecipientOptouts200Response {
	this := GetRecipientOptouts200Response{}
	return &this
}

// GetRecipientOptouts returns the RecipientOptouts field value if set, zero value otherwise.
func (o *GetRecipientOptouts200Response) GetRecipientOptouts() []RecipientOptout {
	if o == nil || IsNil(o.RecipientOptouts) {
		var ret []RecipientOptout
		return ret
	}
	return o.RecipientOptouts
}

// GetRecipientOptoutsOk returns a tuple with the RecipientOptouts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecipientOptouts200Response) GetRecipientOptoutsOk() ([]RecipientOptout, bool) {
	if o == nil || IsNil(o.RecipientOptouts) {
		return nil, false
	}
	return o.RecipientOptouts, true
}

// HasRecipientOptouts returns a boolean if a field has been set.
func (o *GetRecipientOptouts200Response) HasRecipientOptouts() bool {
	if o != nil && !IsNil(o.RecipientOptouts) {
		return true
	}

	return false
}

// SetRecipientOptouts gets a reference to the given []RecipientOptout and assigns it to the RecipientOptouts field.
func (o *GetRecipientOptouts200Response) SetRecipientOptouts(v []RecipientOptout) {
	o.RecipientOptouts = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *GetRecipientOptouts200Response) GetPaging() GetMessages200ResponsePaging {
	if o == nil || IsNil(o.Paging) {
		var ret GetMessages200ResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecipientOptouts200Response) GetPagingOk() (*GetMessages200ResponsePaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *GetRecipientOptouts200Response) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Object and assigns it to the Paging field.
func (o *GetRecipientOptouts200Response) SetPaging(v GetMessages200ResponsePaging) {
	o.Paging = &v
}

func (o GetRecipientOptouts200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecipientOptouts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecipientOptouts) {
		toSerialize["recipientOptouts"] = o.RecipientOptouts
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableGetRecipientOptouts200Response struct {
	value *GetRecipientOptouts200Response
	isSet bool
}

func (v NullableGetRecipientOptouts200Response) Get() *GetRecipientOptouts200Response {
	return v.value
}

func (v *NullableGetRecipientOptouts200Response) Set(val *GetRecipientOptouts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecipientOptouts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecipientOptouts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecipientOptouts200Response(val *GetRecipientOptouts200Response) *NullableGetRecipientOptouts200Response {
	return &NullableGetRecipientOptouts200Response{value: val, isSet: true}
}

func (v NullableGetRecipientOptouts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecipientOptouts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
