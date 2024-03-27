/*
Messaging API v3.4.3

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


*/

package TelstraMessaging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// ReportsAPIService ReportsAPI service
type ReportsAPIService service

type ApiGetReportRequest struct {
	ctx               context.Context
	ApiService        *ReportsAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	reportId          string
	telstraApiVersion *string
}

func (r ApiGetReportRequest) ContentLanguage(contentLanguage string) ApiGetReportRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetReportRequest) Authorization(authorization string) ApiGetReportRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetReportRequest) Accept(accept string) ApiGetReportRequest {
	r.accept = &accept
	return r
}

func (r ApiGetReportRequest) AcceptCharset(acceptCharset string) ApiGetReportRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetReportRequest) ContentType(contentType string) ApiGetReportRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetReportRequest) TelstraApiVersion(telstraApiVersion string) ApiGetReportRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiGetReportRequest) Execute() (*GetReport200Response, *http.Response, error) {
	r.contentType = new(string)
	*r.contentType = SetRequestParams().ContentType
	r.accept = new(string)
	*r.accept = SetRequestParams().Accept
	r.acceptCharset = new(string)
	*r.acceptCharset = SetRequestParams().AcceptCharset
	r.contentLanguage = new(string)
	*r.contentLanguage = SetRequestParams().ContentLanguage
	r.telstraApiVersion = new(string)
	*r.telstraApiVersion = SetRequestParams().TelstraApiVersion

	return r.ApiService.GetReportExecute(r)
}

/*
GetReport fetch a specific report

Fetch a download link for a report generated with POST /reports/{reportId} using the **reportId** returned in the response. Once ready, your report will be available for download for one week.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param reportId Use the reportId returned in the POST /reports/{reportType} response.
	@return ApiGetReportRequest
*/
func (a *ReportsAPIService) GetReport(ctx context.Context, reportId string, authorization string) ApiGetReportRequest {
	return ApiGetReportRequest{
		ApiService:    a,
		ctx:           ctx,
		reportId:      reportId,
		authorization: &authorization,
	}
}

func (r ApiGetReportRequest) GetReport() (*GetReport200Response, *http.Response, error) {
	return r.Execute()
}

// Execute executes the request
//
//	@return GetReport200Response
func (a *ReportsAPIService) GetReportExecute(r ApiGetReportRequest) (*GetReport200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetReport200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/{reportId}"
	localVarPath = strings.Replace(localVarPath, "{"+"reportId"+"}", url.PathEscape(parameterValueToString(r.reportId, "reportId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentLanguage == nil {
		return localVarReturnValue, nil, reportError("contentLanguage is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}
	if r.acceptCharset == nil {
		return localVarReturnValue, nil, reportError("acceptCharset is required and must be specified")
	}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if strlen(r.reportId) < 36 {
		return localVarReturnValue, nil, reportError("reportId must have at least 36 elements")
	}
	if strlen(r.reportId) > 36 {
		return localVarReturnValue, nil, reportError("reportId must have less than 36 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.telstraApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Telstra-api-version", r.telstraApiVersion, "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Language", r.contentLanguage, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Charset", r.acceptCharset, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetReportsRequest struct {
	ctx               context.Context
	ApiService        *ReportsAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	telstraApiVersion *string
}

func (r ApiGetReportsRequest) ContentLanguage(contentLanguage string) ApiGetReportsRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetReportsRequest) Authorization(authorization string) ApiGetReportsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetReportsRequest) Accept(accept string) ApiGetReportsRequest {
	r.accept = &accept
	return r
}

func (r ApiGetReportsRequest) AcceptCharset(acceptCharset string) ApiGetReportsRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetReportsRequest) ContentType(contentType string) ApiGetReportsRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetReportsRequest) TelstraApiVersion(telstraApiVersion string) ApiGetReportsRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiGetReportsRequest) Execute() (*GetReports200Response, *http.Response, error) {
	r.contentType = new(string)
	*r.contentType = SetRequestParams().ContentType
	r.accept = new(string)
	*r.accept = SetRequestParams().Accept
	r.acceptCharset = new(string)
	*r.acceptCharset = SetRequestParams().AcceptCharset
	r.contentLanguage = new(string)
	*r.contentLanguage = SetRequestParams().ContentLanguage
	r.telstraApiVersion = new(string)
	*r.telstraApiVersion = SetRequestParams().TelstraApiVersion
	return r.ApiService.GetReportsExecute(r)
}

/*
GetReports fetch all reports

Fetch details of all reports recently generated for your account. Use it to check the status of a report, plus fetch the report ID, status, report type and expiry date.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetReportsRequest
*/
func (a *ReportsAPIService) GetReports(ctx context.Context, authorization string) ApiGetReportsRequest {
	return ApiGetReportsRequest{
		ApiService:    a,
		ctx:           ctx,
		authorization: &authorization,
	}
}

func (r ApiGetReportsRequest) GetReports() (*GetReports200Response, *http.Response, error) {
	return r.Execute()
}

// Execute executes the request
//
//	@return GetReports200Response
func (a *ReportsAPIService) GetReportsExecute(r ApiGetReportsRequest) (*GetReports200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetReports200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.GetReports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentLanguage == nil {
		return localVarReturnValue, nil, reportError("contentLanguage is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}
	if r.acceptCharset == nil {
		return localVarReturnValue, nil, reportError("acceptCharset is required and must be specified")
	}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.telstraApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Telstra-api-version", r.telstraApiVersion, "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Language", r.contentLanguage, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Charset", r.acceptCharset, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiMessagesReportRequest struct {
	ctx                   context.Context
	ApiService            *ReportsAPIService
	contentLanguage       *string
	authorization         *string
	accept                *string
	acceptCharset         *string
	contentType           *string
	messagesReportRequest *MessagesReportRequest
	telstraApiVersion     *string
}

func (r ApiMessagesReportRequest) ContentLanguage(contentLanguage string) ApiMessagesReportRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiMessagesReportRequest) Authorization(authorization string) ApiMessagesReportRequest {
	r.authorization = &authorization
	return r
}

func (r ApiMessagesReportRequest) Accept(accept string) ApiMessagesReportRequest {
	r.accept = &accept
	return r
}

func (r ApiMessagesReportRequest) AcceptCharset(acceptCharset string) ApiMessagesReportRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiMessagesReportRequest) ContentType(contentType string) ApiMessagesReportRequest {
	r.contentType = &contentType
	return r
}

func (r ApiMessagesReportRequest) MessagesReportRequest(messagesReportRequest MessagesReportRequest) ApiMessagesReportRequest {
	r.messagesReportRequest = &messagesReportRequest
	return r
}

func (r ApiMessagesReportRequest) TelstraApiVersion(telstraApiVersion string) ApiMessagesReportRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiMessagesReportRequest) Execute() (*MessagesReport201Response, *http.Response, error) {
	r.contentType = new(string)
	*r.contentType = SetRequestParams().ContentType
	r.accept = new(string)
	*r.accept = SetRequestParams().Accept
	r.acceptCharset = new(string)
	*r.acceptCharset = SetRequestParams().AcceptCharset
	r.contentLanguage = new(string)
	*r.contentLanguage = SetRequestParams().ContentLanguage
	r.telstraApiVersion = new(string)
	*r.telstraApiVersion = SetRequestParams().TelstraApiVersion
	return r.ApiService.MessagesReportExecute(r)
}

/*
MessagesReport submit a request for a messages report

Request a CSV report of messages (both incoming and outgoing) that have been sent to/from your account within the last three months. You can request details for a specific timeframe, and filter your messages by tags, recipient number or Virtual Number.

A 201 Created means your report has been queued for generation. Make a note of the reportId returned in the response. You'll need this to check the status of your report and fetch your download link with GET reports/{reportId}. If you supplied a reportCallbackUrl in the request we'll also notify it when your report is ready for download.

Once your report is generated, it will be available for download for one week. The expiry date is returned in the response.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMessagesReportRequest
*/
func (a *ReportsAPIService) MessagesReport(ctx context.Context, authorization string) ApiMessagesReportRequest {
	return ApiMessagesReportRequest{
		ApiService:    a,
		ctx:           ctx,
		authorization: &authorization,
	}
}

func (r ApiMessagesReportRequest) MessagesReport(messagesReportRequest MessagesReportRequest) (*MessagesReport201Response, *http.Response, error) {
	r.messagesReportRequest = &messagesReportRequest
	return r.Execute()
}

// Execute executes the request
//
//	@return MessagesReport201Response
func (a *ReportsAPIService) MessagesReportExecute(r ApiMessagesReportRequest) (*MessagesReport201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MessagesReport201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsAPIService.MessagesReport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentLanguage == nil {
		return localVarReturnValue, nil, reportError("contentLanguage is required and must be specified")
	}
	if r.authorization == nil {
		return localVarReturnValue, nil, reportError("authorization is required and must be specified")
	}
	if r.accept == nil {
		return localVarReturnValue, nil, reportError("accept is required and must be specified")
	}
	if r.acceptCharset == nil {
		return localVarReturnValue, nil, reportError("acceptCharset is required and must be specified")
	}
	if r.contentType == nil {
		return localVarReturnValue, nil, reportError("contentType is required and must be specified")
	}
	if r.messagesReportRequest == nil {
		return localVarReturnValue, nil, reportError("messagesReportRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.telstraApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Telstra-api-version", r.telstraApiVersion, "")
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Language", r.contentLanguage, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept-Charset", r.acceptCharset, "")
	parameterAddToHeaderOrQuery(localVarHeaderParams, "Content-Type", r.contentType, "")
	// body params
	localVarPostBody = r.messagesReportRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
