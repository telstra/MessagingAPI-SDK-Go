/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetReport200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReport200Response{}

// GetReport200Response struct for GetReport200Response
type GetReport200Response struct {
	// The UUID assigned to your report.
	ReportId *string `json:"reportId,omitempty"`
	// The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.
	ReportStatus *string `json:"reportStatus,omitempty"`
	// Use this link to download your CSV file.
	ReportUrl *string `json:"reportUrl,omitempty"`
}

// NewGetReport200Response instantiates a new GetReport200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReport200Response() *GetReport200Response {
	this := GetReport200Response{}
	return &this
}

// NewGetReport200ResponseWithDefaults instantiates a new GetReport200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReport200ResponseWithDefaults() *GetReport200Response {
	this := GetReport200Response{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *GetReport200Response) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReport200Response) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *GetReport200Response) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *GetReport200Response) SetReportId(v string) {
	o.ReportId = &v
}

// GetReportStatus returns the ReportStatus field value if set, zero value otherwise.
func (o *GetReport200Response) GetReportStatus() string {
	if o == nil || IsNil(o.ReportStatus) {
		var ret string
		return ret
	}
	return *o.ReportStatus
}

// GetReportStatusOk returns a tuple with the ReportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReport200Response) GetReportStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReportStatus) {
		return nil, false
	}
	return o.ReportStatus, true
}

// HasReportStatus returns a boolean if a field has been set.
func (o *GetReport200Response) HasReportStatus() bool {
	if o != nil && !IsNil(o.ReportStatus) {
		return true
	}

	return false
}

// SetReportStatus gets a reference to the given string and assigns it to the ReportStatus field.
func (o *GetReport200Response) SetReportStatus(v string) {
	o.ReportStatus = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *GetReport200Response) GetReportUrl() string {
	if o == nil || IsNil(o.ReportUrl) {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReport200Response) GetReportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportUrl) {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *GetReport200Response) HasReportUrl() bool {
	if o != nil && !IsNil(o.ReportUrl) {
		return true
	}

	return false
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *GetReport200Response) SetReportUrl(v string) {
	o.ReportUrl = &v
}

func (o GetReport200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReport200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.ReportStatus) {
		toSerialize["reportStatus"] = o.ReportStatus
	}
	if !IsNil(o.ReportUrl) {
		toSerialize["reportUrl"] = o.ReportUrl
	}
	return toSerialize, nil
}

type NullableGetReport200Response struct {
	value *GetReport200Response
	isSet bool
}

func (v NullableGetReport200Response) Get() *GetReport200Response {
	return v.value
}

func (v *NullableGetReport200Response) Set(val *GetReport200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReport200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReport200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReport200Response(val *GetReport200Response) *NullableGetReport200Response {
	return &NullableGetReport200Response{value: val, isSet: true}
}

func (v NullableGetReport200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReport200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
