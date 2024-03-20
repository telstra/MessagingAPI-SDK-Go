# MessagesReportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | Set the time period you want to generate a report for by typing the start date (inclusive) here.  Note that we only retain data for three months, so please ensure your startDate is not more than three months old.  Use ISO format(yyyy-mm-dd), e.g. 2019-08-24.  | 
**EndDate** | **string** | Type the end date (inclusive) of your reporting period here.  Your endDate must be a date in the past, and less than three months from your startDate. Use ISO format(yyyy-mm-dd).  | 
**ReportCallbackUrl** | Pointer to **string** | Tell us the callbackUrl you want us to notify when your report is ready for download.  Sample callback response:  &lt;pre&gt;&lt;code class&#x3D;\&quot;language-sh\&quot;&gt;{   \&quot;reportId\&quot;:\&quot;1520b774-46b0-4415-a05f-7bcb1c032c59\&quot;,   \&quot;reportStatus\&quot;:\&quot;completed\&quot;,   \&quot;timestamp\&quot;:\&quot;2022-11-10T05:06:42.823Z\&quot; }&lt;/code&gt;&lt;/pre&gt;  | [optional] 
**Filter** | Pointer to **string** | Filter your messages report by:   * tag - use one of the tags assigned to your message(s)   * number - either the Virtual Number used to send the message, or the Recipient Number the message was sent to.  | [optional] 

## Methods

### NewMessagesReportRequest

`func NewMessagesReportRequest(startDate string, endDate string, ) *MessagesReportRequest`

NewMessagesReportRequest instantiates a new MessagesReportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesReportRequestWithDefaults

`func NewMessagesReportRequestWithDefaults() *MessagesReportRequest`

NewMessagesReportRequestWithDefaults instantiates a new MessagesReportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *MessagesReportRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *MessagesReportRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *MessagesReportRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *MessagesReportRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *MessagesReportRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *MessagesReportRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.


### GetReportCallbackUrl

`func (o *MessagesReportRequest) GetReportCallbackUrl() string`

GetReportCallbackUrl returns the ReportCallbackUrl field if non-nil, zero value otherwise.

### GetReportCallbackUrlOk

`func (o *MessagesReportRequest) GetReportCallbackUrlOk() (*string, bool)`

GetReportCallbackUrlOk returns a tuple with the ReportCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCallbackUrl

`func (o *MessagesReportRequest) SetReportCallbackUrl(v string)`

SetReportCallbackUrl sets ReportCallbackUrl field to given value.

### HasReportCallbackUrl

`func (o *MessagesReportRequest) HasReportCallbackUrl() bool`

HasReportCallbackUrl returns a boolean if a field has been set.

### GetFilter

`func (o *MessagesReportRequest) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MessagesReportRequest) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MessagesReportRequest) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MessagesReportRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


