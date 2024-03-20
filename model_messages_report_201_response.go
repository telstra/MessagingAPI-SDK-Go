/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the MessagesReport201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagesReport201Response{}

// MessagesReport201Response struct for MessagesReport201Response
type MessagesReport201Response struct {
	// When your report is ready for download, you can use this UUID to fetch it with the GET /reports/{reportId} endpoint.
	ReportId *string `json:"reportId,omitempty"`
	// If you provided a reportCallbackUrl in your request, it will be returned here.
	ReportCallbackUrl *string `json:"reportCallbackUrl,omitempty"`
	// The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.
	ReportStatus *string `json:"reportStatus,omitempty"`
}

// NewMessagesReport201Response instantiates a new MessagesReport201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagesReport201Response() *MessagesReport201Response {
	this := MessagesReport201Response{}
	return &this
}

// NewMessagesReport201ResponseWithDefaults instantiates a new MessagesReport201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagesReport201ResponseWithDefaults() *MessagesReport201Response {
	this := MessagesReport201Response{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *MessagesReport201Response) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesReport201Response) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *MessagesReport201Response) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *MessagesReport201Response) SetReportId(v string) {
	o.ReportId = &v
}

// GetReportCallbackUrl returns the ReportCallbackUrl field value if set, zero value otherwise.
func (o *MessagesReport201Response) GetReportCallbackUrl() string {
	if o == nil || IsNil(o.ReportCallbackUrl) {
		var ret string
		return ret
	}
	return *o.ReportCallbackUrl
}

// GetReportCallbackUrlOk returns a tuple with the ReportCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesReport201Response) GetReportCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportCallbackUrl) {
		return nil, false
	}
	return o.ReportCallbackUrl, true
}

// HasReportCallbackUrl returns a boolean if a field has been set.
func (o *MessagesReport201Response) HasReportCallbackUrl() bool {
	if o != nil && !IsNil(o.ReportCallbackUrl) {
		return true
	}

	return false
}

// SetReportCallbackUrl gets a reference to the given string and assigns it to the ReportCallbackUrl field.
func (o *MessagesReport201Response) SetReportCallbackUrl(v string) {
	o.ReportCallbackUrl = &v
}

// GetReportStatus returns the ReportStatus field value if set, zero value otherwise.
func (o *MessagesReport201Response) GetReportStatus() string {
	if o == nil || IsNil(o.ReportStatus) {
		var ret string
		return ret
	}
	return *o.ReportStatus
}

// GetReportStatusOk returns a tuple with the ReportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagesReport201Response) GetReportStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReportStatus) {
		return nil, false
	}
	return o.ReportStatus, true
}

// HasReportStatus returns a boolean if a field has been set.
func (o *MessagesReport201Response) HasReportStatus() bool {
	if o != nil && !IsNil(o.ReportStatus) {
		return true
	}

	return false
}

// SetReportStatus gets a reference to the given string and assigns it to the ReportStatus field.
func (o *MessagesReport201Response) SetReportStatus(v string) {
	o.ReportStatus = &v
}

func (o MessagesReport201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagesReport201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.ReportCallbackUrl) {
		toSerialize["reportCallbackUrl"] = o.ReportCallbackUrl
	}
	if !IsNil(o.ReportStatus) {
		toSerialize["reportStatus"] = o.ReportStatus
	}
	return toSerialize, nil
}

type NullableMessagesReport201Response struct {
	value *MessagesReport201Response
	isSet bool
}

func (v NullableMessagesReport201Response) Get() *MessagesReport201Response {
	return v.value
}

func (v *NullableMessagesReport201Response) Set(val *MessagesReport201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagesReport201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagesReport201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagesReport201Response(val *MessagesReport201Response) *NullableMessagesReport201Response {
	return &NullableMessagesReport201Response{value: val, isSet: true}
}

func (v NullableMessagesReport201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagesReport201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
