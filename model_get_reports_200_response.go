/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetReports200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReports200Response{}

// GetReports200Response struct for GetReports200Response
type GetReports200Response struct {
	Reports []GetReports200ResponseReportsInner `json:"reports,omitempty"`
}

// NewGetReports200Response instantiates a new GetReports200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReports200Response() *GetReports200Response {
	this := GetReports200Response{}
	return &this
}

// NewGetReports200ResponseWithDefaults instantiates a new GetReports200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReports200ResponseWithDefaults() *GetReports200Response {
	this := GetReports200Response{}
	return &this
}

// GetReports returns the Reports field value if set, zero value otherwise.
func (o *GetReports200Response) GetReports() []GetReports200ResponseReportsInner {
	if o == nil || IsNil(o.Reports) {
		var ret []GetReports200ResponseReportsInner
		return ret
	}
	return o.Reports
}

// GetReportsOk returns a tuple with the Reports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReports200Response) GetReportsOk() ([]GetReports200ResponseReportsInner, bool) {
	if o == nil || IsNil(o.Reports) {
		return nil, false
	}
	return o.Reports, true
}

// HasReports returns a boolean if a field has been set.
func (o *GetReports200Response) HasReports() bool {
	if o != nil && !IsNil(o.Reports) {
		return true
	}

	return false
}

// SetReports gets a reference to the given []GetReports200ResponseReportsInner and assigns it to the Reports field.
func (o *GetReports200Response) SetReports(v []GetReports200ResponseReportsInner) {
	o.Reports = v
}

func (o GetReports200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReports200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reports) {
		toSerialize["reports"] = o.Reports
	}
	return toSerialize, nil
}

type NullableGetReports200Response struct {
	value *GetReports200Response
	isSet bool
}

func (v NullableGetReports200Response) Get() *GetReports200Response {
	return v.value
}

func (v *NullableGetReports200Response) Set(val *GetReports200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReports200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReports200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReports200Response(val *GetReports200Response) *NullableGetReports200Response {
	return &NullableGetReports200Response{value: val, isSet: true}
}

func (v NullableGetReports200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReports200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
