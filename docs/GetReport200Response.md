# GetReport200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** | The UUID assigned to your report.  | [optional] 
**ReportStatus** | Pointer to **string** | The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.  | [optional] 
**ReportUrl** | Pointer to **string** | Use this link to download your CSV file. | [optional] 

## Methods

### NewGetReport200Response

`func NewGetReport200Response() *GetReport200Response`

NewGetReport200Response instantiates a new GetReport200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReport200ResponseWithDefaults

`func NewGetReport200ResponseWithDefaults() *GetReport200Response`

NewGetReport200ResponseWithDefaults instantiates a new GetReport200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *GetReport200Response) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *GetReport200Response) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *GetReport200Response) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *GetReport200Response) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportStatus

`func (o *GetReport200Response) GetReportStatus() string`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *GetReport200Response) GetReportStatusOk() (*string, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *GetReport200Response) SetReportStatus(v string)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *GetReport200Response) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.

### GetReportUrl

`func (o *GetReport200Response) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *GetReport200Response) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *GetReport200Response) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *GetReport200Response) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


