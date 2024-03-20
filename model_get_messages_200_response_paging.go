/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetMessages200ResponsePaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMessages200ResponsePaging{}

// GetMessages200ResponsePaging struct for GetMessages200ResponsePaging
type GetMessages200ResponsePaging struct {
	// The next page of results.
	NextPage *string `json:"nextPage,omitempty"`
	// The previous page of results.
	PreviousPage *string `json:"previousPage,omitempty"`
	// The last page of results.
	LastPage *string `json:"lastPage,omitempty"`
	// The total number of pages.
	TotalCount *float32 `json:"totalCount,omitempty"`
}

// NewGetMessages200ResponsePaging instantiates a new GetMessages200ResponsePaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMessages200ResponsePaging() *GetMessages200ResponsePaging {
	this := GetMessages200ResponsePaging{}
	return &this
}

// NewGetMessages200ResponsePagingWithDefaults instantiates a new GetMessages200ResponsePaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMessages200ResponsePagingWithDefaults() *GetMessages200ResponsePaging {
	this := GetMessages200ResponsePaging{}
	return &this
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *GetMessages200ResponsePaging) GetNextPage() string {
	if o == nil || IsNil(o.NextPage) {
		var ret string
		return ret
	}
	return *o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200ResponsePaging) GetNextPageOk() (*string, bool) {
	if o == nil || IsNil(o.NextPage) {
		return nil, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *GetMessages200ResponsePaging) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given string and assigns it to the NextPage field.
func (o *GetMessages200ResponsePaging) SetNextPage(v string) {
	o.NextPage = &v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise.
func (o *GetMessages200ResponsePaging) GetPreviousPage() string {
	if o == nil || IsNil(o.PreviousPage) {
		var ret string
		return ret
	}
	return *o.PreviousPage
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200ResponsePaging) GetPreviousPageOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousPage) {
		return nil, false
	}
	return o.PreviousPage, true
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *GetMessages200ResponsePaging) HasPreviousPage() bool {
	if o != nil && !IsNil(o.PreviousPage) {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given string and assigns it to the PreviousPage field.
func (o *GetMessages200ResponsePaging) SetPreviousPage(v string) {
	o.PreviousPage = &v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *GetMessages200ResponsePaging) GetLastPage() string {
	if o == nil || IsNil(o.LastPage) {
		var ret string
		return ret
	}
	return *o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200ResponsePaging) GetLastPageOk() (*string, bool) {
	if o == nil || IsNil(o.LastPage) {
		return nil, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *GetMessages200ResponsePaging) HasLastPage() bool {
	if o != nil && !IsNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given string and assigns it to the LastPage field.
func (o *GetMessages200ResponsePaging) SetLastPage(v string) {
	o.LastPage = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *GetMessages200ResponsePaging) GetTotalCount() float32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret float32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMessages200ResponsePaging) GetTotalCountOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *GetMessages200ResponsePaging) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given float32 and assigns it to the TotalCount field.
func (o *GetMessages200ResponsePaging) SetTotalCount(v float32) {
	o.TotalCount = &v
}

func (o GetMessages200ResponsePaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMessages200ResponsePaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PreviousPage) {
		toSerialize["previousPage"] = o.PreviousPage
	}
	if !IsNil(o.LastPage) {
		toSerialize["lastPage"] = o.LastPage
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableGetMessages200ResponsePaging struct {
	value *GetMessages200ResponsePaging
	isSet bool
}

func (v NullableGetMessages200ResponsePaging) Get() *GetMessages200ResponsePaging {
	return v.value
}

func (v *NullableGetMessages200ResponsePaging) Set(val *GetMessages200ResponsePaging) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMessages200ResponsePaging) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMessages200ResponsePaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMessages200ResponsePaging(val *GetMessages200ResponsePaging) *NullableGetMessages200ResponsePaging {
	return &NullableGetMessages200ResponsePaging{value: val, isSet: true}
}

func (v NullableGetMessages200ResponsePaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMessages200ResponsePaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
