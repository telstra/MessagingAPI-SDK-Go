/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"encoding/json"
)

// checks if the GetReports200ResponseReportsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetReports200ResponseReportsInner{}

// GetReports200ResponseReportsInner struct for GetReports200ResponseReportsInner
type GetReports200ResponseReportsInner struct {
	// Use this UUID to fetch your report with the GET /reports/{reportId} endpoint.
	ReportId *string `json:"reportId,omitempty"`
	// The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.
	ReportStatus *string `json:"reportStatus,omitempty"`
	// The type of report generated.
	ReportType *string `json:"reportType,omitempty"`
	// The expiry date of your report. After this date, you will be unable to download your report.
	ReportExpiry *string `json:"reportExpiry,omitempty"`
}

// NewGetReports200ResponseReportsInner instantiates a new GetReports200ResponseReportsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetReports200ResponseReportsInner() *GetReports200ResponseReportsInner {
	this := GetReports200ResponseReportsInner{}
	return &this
}

// NewGetReports200ResponseReportsInnerWithDefaults instantiates a new GetReports200ResponseReportsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetReports200ResponseReportsInnerWithDefaults() *GetReports200ResponseReportsInner {
	this := GetReports200ResponseReportsInner{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *GetReports200ResponseReportsInner) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReports200ResponseReportsInner) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *GetReports200ResponseReportsInner) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *GetReports200ResponseReportsInner) SetReportId(v string) {
	o.ReportId = &v
}

// GetReportStatus returns the ReportStatus field value if set, zero value otherwise.
func (o *GetReports200ResponseReportsInner) GetReportStatus() string {
	if o == nil || IsNil(o.ReportStatus) {
		var ret string
		return ret
	}
	return *o.ReportStatus
}

// GetReportStatusOk returns a tuple with the ReportStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReports200ResponseReportsInner) GetReportStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReportStatus) {
		return nil, false
	}
	return o.ReportStatus, true
}

// HasReportStatus returns a boolean if a field has been set.
func (o *GetReports200ResponseReportsInner) HasReportStatus() bool {
	if o != nil && !IsNil(o.ReportStatus) {
		return true
	}

	return false
}

// SetReportStatus gets a reference to the given string and assigns it to the ReportStatus field.
func (o *GetReports200ResponseReportsInner) SetReportStatus(v string) {
	o.ReportStatus = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *GetReports200ResponseReportsInner) GetReportType() string {
	if o == nil || IsNil(o.ReportType) {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReports200ResponseReportsInner) GetReportTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReportType) {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *GetReports200ResponseReportsInner) HasReportType() bool {
	if o != nil && !IsNil(o.ReportType) {
		return true
	}

	return false
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *GetReports200ResponseReportsInner) SetReportType(v string) {
	o.ReportType = &v
}

// GetReportExpiry returns the ReportExpiry field value if set, zero value otherwise.
func (o *GetReports200ResponseReportsInner) GetReportExpiry() string {
	if o == nil || IsNil(o.ReportExpiry) {
		var ret string
		return ret
	}
	return *o.ReportExpiry
}

// GetReportExpiryOk returns a tuple with the ReportExpiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetReports200ResponseReportsInner) GetReportExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.ReportExpiry) {
		return nil, false
	}
	return o.ReportExpiry, true
}

// HasReportExpiry returns a boolean if a field has been set.
func (o *GetReports200ResponseReportsInner) HasReportExpiry() bool {
	if o != nil && !IsNil(o.ReportExpiry) {
		return true
	}

	return false
}

// SetReportExpiry gets a reference to the given string and assigns it to the ReportExpiry field.
func (o *GetReports200ResponseReportsInner) SetReportExpiry(v string) {
	o.ReportExpiry = &v
}

func (o GetReports200ResponseReportsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetReports200ResponseReportsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.ReportStatus) {
		toSerialize["reportStatus"] = o.ReportStatus
	}
	if !IsNil(o.ReportType) {
		toSerialize["reportType"] = o.ReportType
	}
	if !IsNil(o.ReportExpiry) {
		toSerialize["reportExpiry"] = o.ReportExpiry
	}
	return toSerialize, nil
}

type NullableGetReports200ResponseReportsInner struct {
	value *GetReports200ResponseReportsInner
	isSet bool
}

func (v NullableGetReports200ResponseReportsInner) Get() *GetReports200ResponseReportsInner {
	return v.value
}

func (v *NullableGetReports200ResponseReportsInner) Set(val *GetReports200ResponseReportsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetReports200ResponseReportsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetReports200ResponseReportsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetReports200ResponseReportsInner(val *GetReports200ResponseReportsInner) *NullableGetReports200ResponseReportsInner {
	return &NullableGetReports200ResponseReportsInner{value: val, isSet: true}
}

func (v NullableGetReports200ResponseReportsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetReports200ResponseReportsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
