/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the MessagesReportRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesReportRequest{}

// MessagesReportRequest struct for MessagesReportRequest
type MessagesReportRequest struct {
	// Set the time period you want to generate a report for by typing the start date (inclusive) here.  Note that we only retain data for three months, so please ensure your startDate is not more than three months old.  Use ISO format(yyyy-mm-dd), e.g. 2019-08-24.
	StartDate string `json:"startDate"`
	// Type the end date (inclusive) of your reporting period here.  Your endDate must be a date in the past, and less than three months from your startDate. Use ISO format(yyyy-mm-dd).
	EndDate string `json:"endDate"`
	// Tell us the callbackUrl you want us to notify when your report is ready for download.  Sample callback response:  <pre><code class=\"language-sh\">{   \"reportId\":\"1520b774-46b0-4415-a05f-7bcb1c032c59\",   \"reportStatus\":\"completed\",   \"timestamp\":\"2022-11-10T05:06:42.823Z\" }</code></pre>
	ReportCallbackUrl *string `json:"reportCallbackUrl,omitempty"`
	// Filter your messages report by:   * tag - use one of the tags assigned to your message(s)   * number - either the Virtual Number used to send the message, or the Recipient Number the message was sent to.
	Filter *string `json:"filter,omitempty"`
}

// NewMessagesReportRequest instantiates a new MessagesReportRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesReportRequest(startDate string, endDate string) *MessagesReportRequest {
	this := MessagesReportRequest{}
	this.StartDate = startDate
	this.EndDate = endDate
	return &this
}

// NewMessagesReportRequestWithDefaults instantiates a new MessagesReportRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesReportRequestWithDefaults() *MessagesReportRequest {
	this := MessagesReportRequest{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *MessagesReportRequest) GetStartDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *MessagesReportRequest) GetStartDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *MessagesReportRequest) SetStartDate(v string) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value
func (o *MessagesReportRequest) GetEndDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *MessagesReportRequest) GetEndDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *MessagesReportRequest) SetEndDate(v string) {
	o.EndDate = v
}

// GetReportCallbackUrl returns the ReportCallbackUrl field value if set, zero value otherwise.
func (o *MessagesReportRequest) GetReportCallbackUrl() string {
	if o == nil || IsNil(o.ReportCallbackUrl) {
		var ret string
		return ret
	}
	return *o.ReportCallbackUrl
}

// GetReportCallbackUrlOk returns a tuple with the ReportCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesReportRequest) GetReportCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportCallbackUrl) {
		return nil, false
	}
	return o.ReportCallbackUrl, true
}

// HasReportCallbackUrl returns a boolean if a field has been set.
func (o *MessagesReportRequest) HasReportCallbackUrl() bool {
	if o != nil && !IsNil(o.ReportCallbackUrl) {
		return true
	}

	return false
}

// SetReportCallbackUrl gets a reference to the given string and assigns it to the ReportCallbackUrl field.
func (o *MessagesReportRequest) SetReportCallbackUrl(v string) {
	o.ReportCallbackUrl = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MessagesReportRequest) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesReportRequest) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MessagesReportRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *MessagesReportRequest) SetFilter(v string) {
	o.Filter = &v
}

func (o MessagesReportRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesReportRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startDate"] = o.StartDate
	toSerialize["endDate"] = o.EndDate
	if !IsNil(o.ReportCallbackUrl) {
		toSerialize["reportCallbackUrl"] = o.ReportCallbackUrl
	}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	return toSerialize, nil
}

type NullableMessagesReportRequest struct {
	value *MessagesReportRequest
	isSet bool
}

func (v NullableMessagesReportRequest) Get() *MessagesReportRequest {
	return v.value
}

func (v *NullableMessagesReportRequest) Set(val *MessagesReportRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesReportRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesReportRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesReportRequest(val *MessagesReportRequest) *NullableMessagesReportRequest {
	return &NullableMessagesReportRequest{value: val, isSet: true}
}

func (v NullableMessagesReportRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesReportRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
