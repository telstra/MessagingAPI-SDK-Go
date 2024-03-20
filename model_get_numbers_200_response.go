/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetNumbers200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNumbers200Response{}

// GetNumbers200Response struct for GetNumbers200Response
type GetNumbers200Response struct {
	// The paginated results of your request. To fetch the next set of results, send another request and provide the succeeding offset value.
	VirtualNumbers []VirtualNumber               `json:"virtualNumbers,omitempty"`
	Paging         *GetMessages200ResponsePaging `json:"paging,omitempty"`
}

// NewGetNumbers200Response instantiates a new GetNumbers200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNumbers200Response() *GetNumbers200Response {
	this := GetNumbers200Response{}
	return &this
}

// NewGetNumbers200ResponseWithDefaults instantiates a new GetNumbers200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNumbers200ResponseWithDefaults() *GetNumbers200Response {
	this := GetNumbers200Response{}
	return &this
}

// GetVirtualNumbers returns the VirtualNumbers field value if set, zero value otherwise.
func (o *GetNumbers200Response) GetVirtualNumbers() []VirtualNumber {
	if o == nil || IsNil(o.VirtualNumbers) {
		var ret []VirtualNumber
		return ret
	}
	return o.VirtualNumbers
}

// GetVirtualNumbersOk returns a tuple with the VirtualNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumbers200Response) GetVirtualNumbersOk() ([]VirtualNumber, bool) {
	if o == nil || IsNil(o.VirtualNumbers) {
		return nil, false
	}
	return o.VirtualNumbers, true
}

// HasVirtualNumbers returns a boolean if a field has been set.
func (o *GetNumbers200Response) HasVirtualNumbers() bool {
	if o != nil && !IsNil(o.VirtualNumbers) {
		return true
	}

	return false
}

// SetVirtualNumbers gets a reference to the given []VirtualNumber and assigns it to the VirtualNumbers field.
func (o *GetNumbers200Response) SetVirtualNumbers(v []VirtualNumber) {
	o.VirtualNumbers = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *GetNumbers200Response) GetPaging() GetMessages200ResponsePaging {
	if o == nil || IsNil(o.Paging) {
		var ret GetMessages200ResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumbers200Response) GetPagingOk() (*GetMessages200ResponsePaging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *GetNumbers200Response) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Object and assigns it to the Paging field.
func (o *GetNumbers200Response) SetPaging(v GetMessages200ResponsePaging) {
	o.Paging = &v
}

func (o GetNumbers200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNumbers200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VirtualNumbers) {
		toSerialize["virtualNumbers"] = o.VirtualNumbers
	}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	return toSerialize, nil
}

type NullableGetNumbers200Response struct {
	value *GetNumbers200Response
	isSet bool
}

func (v NullableGetNumbers200Response) Get() *GetNumbers200Response {
	return v.value
}

func (v *NullableGetNumbers200Response) Set(val *GetNumbers200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNumbers200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNumbers200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNumbers200Response(val *GetNumbers200Response) *NullableGetNumbers200Response {
	return &NullableGetNumbers200Response{value: val, isSet: true}
}

func (v NullableGetNumbers200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNumbers200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
