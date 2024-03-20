# MessagesReport201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReportId** | Pointer to **string** | When your report is ready for download, you can use this UUID to fetch it with the GET /reports/{reportId} endpoint.  | [optional] 
**ReportCallbackUrl** | Pointer to **string** | If you provided a reportCallbackUrl in your request, it will be returned here. | [optional] 
**ReportStatus** | Pointer to **string** | The status of the report. It will be either:        * **queued** – the report is in the queue for generation.        * **completed** – the report is ready for download.        * **failed** – the report failed to generate. Please try again.  | [optional] 

## Methods

### NewMessagesReport201Response

`func NewMessagesReport201Response() *MessagesReport201Response`

NewMessagesReport201Response instantiates a new MessagesReport201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagesReport201ResponseWithDefaults

`func NewMessagesReport201ResponseWithDefaults() *MessagesReport201Response`

NewMessagesReport201ResponseWithDefaults instantiates a new MessagesReport201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReportId

`func (o *MessagesReport201Response) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *MessagesReport201Response) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *MessagesReport201Response) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *MessagesReport201Response) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportCallbackUrl

`func (o *MessagesReport201Response) GetReportCallbackUrl() string`

GetReportCallbackUrl returns the ReportCallbackUrl field if non-nil, zero value otherwise.

### GetReportCallbackUrlOk

`func (o *MessagesReport201Response) GetReportCallbackUrlOk() (*string, bool)`

GetReportCallbackUrlOk returns a tuple with the ReportCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCallbackUrl

`func (o *MessagesReport201Response) SetReportCallbackUrl(v string)`

SetReportCallbackUrl sets ReportCallbackUrl field to given value.

### HasReportCallbackUrl

`func (o *MessagesReport201Response) HasReportCallbackUrl() bool`

HasReportCallbackUrl returns a boolean if a field has been set.

### GetReportStatus

`func (o *MessagesReport201Response) GetReportStatus() string`

GetReportStatus returns the ReportStatus field if non-nil, zero value otherwise.

### GetReportStatusOk

`func (o *MessagesReport201Response) GetReportStatusOk() (*string, bool)`

GetReportStatusOk returns a tuple with the ReportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportStatus

`func (o *MessagesReport201Response) SetReportStatus(v string)`

SetReportStatus sets ReportStatus field to given value.

### HasReportStatus

`func (o *MessagesReport201Response) HasReportStatus() bool`

HasReportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


