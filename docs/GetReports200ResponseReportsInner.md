# GetReports200ResponseReportsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** | Use this UUID to fetch your report with the GET /reports/{reportId} endpoint.  | [optional] 
**ReportStatus** | Pointer to **string** | The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.  | [optional] 
**ReportType** | Pointer to **string** | The type of report generated. | [optional] 
**ReportExpiry** | Pointer to **string** | The expiry date of your report. After this date, you will be unable to download your report.  | [optional] 

## Methods

### NewGetReports200ResponseReportsInner

`func NewGetReports200ResponseReportsInner() *GetReports200ResponseReportsInner`

NewGetReports200ResponseReportsInner instantiates a new GetReports200ResponseReportsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReports200ResponseReportsInnerWithDefaults

`func NewGetReports200ResponseReportsInnerWithDefaults() *GetReports200ResponseReportsInner`

NewGetReports200ResponseReportsInnerWithDefaults instantiates a new GetReports200ResponseReportsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *GetReports200ResponseReportsInner) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *GetReports200ResponseReportsInner) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *GetReports200ResponseReportsInner) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *GetReports200ResponseReportsInner) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportStatus

`func (o *GetReports200ResponseReportsInner) GetReportStatus() string`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *GetReports200ResponseReportsInner) GetReportStatusOk() (*string, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *GetReports200ResponseReportsInner) SetReportStatus(v string)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *GetReports200ResponseReportsInner) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.

### GetReportType

`func (o *GetReports200ResponseReportsInner) GetReportType() string`

GetReportType returns the ReportType field if non-nil, zero value otherwise.

### GetReportTypeOk

`func (o *GetReports200ResponseReportsInner) GetReportTypeOk() (*string, bool)`

GetReportTypeOk returns a tuple with the ReportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportType

`func (o *GetReports200ResponseReportsInner) SetReportType(v string)`

SetReportType sets ReportType field to given value.

### HasReportType

`func (o *GetReports200ResponseReportsInner) HasReportType() bool`

HasReportType returns a boolean if a field has been set.

### GetReportExpiry

`func (o *GetReports200ResponseReportsInner) GetReportExpiry() string`

GetReportExpiry returns the ReportExpiry field if non-nil, zero value otherwise.

### GetReportExpiryOk

`func (o *GetReports200ResponseReportsInner) GetReportExpiryOk() (*string, bool)`

GetReportExpiryOk returns a tuple with the ReportExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportExpiry

`func (o *GetReports200ResponseReportsInner) SetReportExpiry(v string)`

SetReportExpiry sets ReportExpiry field to given value.

### HasReportExpiry

`func (o *GetReports200ResponseReportsInner) HasReportExpiry() bool`

HasReportExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


