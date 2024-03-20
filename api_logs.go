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
	"time"
)

// LogsAPIService LogsAPI service
type LogsAPIService service

type ApiGetLogsRequest struct {
	ctx               context.Context
	ApiService        *LogsAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	telstraApiVersion *string
	startDate         *time.Time
	endDate           *time.Time
	logsFilter        *string
}

func (r ApiGetLogsRequest) ContentLanguage(contentLanguage string) ApiGetLogsRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetLogsRequest) Authorization(authorization string) ApiGetLogsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetLogsRequest) Accept(accept string) ApiGetLogsRequest {
	r.accept = &accept
	return r
}

func (r ApiGetLogsRequest) AcceptCharset(acceptCharset string) ApiGetLogsRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetLogsRequest) ContentType(contentType string) ApiGetLogsRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetLogsRequest) TelstraApiVersion(telstraApiVersion string) ApiGetLogsRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

// Fetch logs for calls made from a specific time and date. Default value is events logged within the last week.
func (r ApiGetLogsRequest) StartDate(startDate time.Time) ApiGetLogsRequest {
	r.startDate = &startDate
	return r
}

// Fetch logs for calls made until a specific time and date. Default value is current time and date
func (r ApiGetLogsRequest) EndDate(endDate time.Time) ApiGetLogsRequest {
	r.endDate = &endDate
	return r
}

// Filter your logs by:  * date - rather than searching within a time range, you can choose to return logs for a specific date instead. Use string &lt;date-time&gt; format. * resource – return logs for a specific resource, e.g. POST /messages.
func (r ApiGetLogsRequest) LogsFilter(logsFilter string) ApiGetLogsRequest {
	r.logsFilter = &logsFilter
	return r
}

func (r ApiGetLogsRequest) Execute() (*GetLogs200Response, *http.Response, error) {
	return r.ApiService.GetLogsExecute(r)
}

/*
GetLogs get logs

Fetch a log of events made from your account. Get details of the last 50 calls (default), display results for a specific timeframe, or filter your logs by date or by resource.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetLogsRequest
*/
func (a *LogsAPIService) GetLogs(ctx context.Context) ApiGetLogsRequest {
	return ApiGetLogsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetLogs200Response
func (a *LogsAPIService) GetLogsExecute(r ApiGetLogsRequest) (*GetLogs200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetLogs200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogsAPIService.GetLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs"

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

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.logsFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "logsFilter", r.logsFilter, "")
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
