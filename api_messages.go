/*
Messaging API v3.1.0

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network. With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code. Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.

API version: 3.1.0
*/

package TelstraMessaging

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// MessagesAPIService MessagesAPI service
type MessagesAPIService service

type ApiDeleteMessageByIdRequest struct {
	ctx               context.Context
	ApiService        *MessagesAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	messageId         string
	telstraApiVersion *string
}

func (r ApiDeleteMessageByIdRequest) ContentLanguage(contentLanguage string) ApiDeleteMessageByIdRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiDeleteMessageByIdRequest) Authorization(authorization string) ApiDeleteMessageByIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteMessageByIdRequest) Accept(accept string) ApiDeleteMessageByIdRequest {
	r.accept = &accept
	return r
}

func (r ApiDeleteMessageByIdRequest) AcceptCharset(acceptCharset string) ApiDeleteMessageByIdRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiDeleteMessageByIdRequest) ContentType(contentType string) ApiDeleteMessageByIdRequest {
	r.contentType = &contentType
	return r
}

func (r ApiDeleteMessageByIdRequest) TelstraApiVersion(telstraApiVersion string) ApiDeleteMessageByIdRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiDeleteMessageByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMessageByIdExecute(r)
}

/*
DeleteMessageById delete a message

Use this endpoint to delete a message that's been scheduled for sending, but hasn't yet sent.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.
	@return ApiDeleteMessageByIdRequest
*/
func (a *MessagesAPIService) DeleteMessageById(ctx context.Context, messageId string) ApiDeleteMessageByIdRequest {
	return ApiDeleteMessageByIdRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (a *MessagesAPIService) DeleteMessageByIdExecute(r ApiDeleteMessageByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.DeleteMessageById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentLanguage == nil {
		return nil, reportError("contentLanguage is required and must be specified")
	}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.accept == nil {
		return nil, reportError("accept is required and must be specified")
	}
	if r.acceptCharset == nil {
		return nil, reportError("acceptCharset is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}
	if strlen(r.messageId) < 36 {
		return nil, reportError("messageId must have at least 36 elements")
	}
	if strlen(r.messageId) > 36 {
		return nil, reportError("messageId must have less than 36 elements")
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetMessageByIdRequest struct {
	ctx               context.Context
	ApiService        *MessagesAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	messageId         string
	telstraApiVersion *string
}

func (r ApiGetMessageByIdRequest) ContentLanguage(contentLanguage string) ApiGetMessageByIdRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetMessageByIdRequest) Authorization(authorization string) ApiGetMessageByIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetMessageByIdRequest) Accept(accept string) ApiGetMessageByIdRequest {
	r.accept = &accept
	return r
}

func (r ApiGetMessageByIdRequest) AcceptCharset(acceptCharset string) ApiGetMessageByIdRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetMessageByIdRequest) ContentType(contentType string) ApiGetMessageByIdRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetMessageByIdRequest) TelstraApiVersion(telstraApiVersion string) ApiGetMessageByIdRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiGetMessageByIdRequest) Execute() (*MessageGet, *http.Response, error) {
	return r.ApiService.GetMessageByIdExecute(r)
}

/*
GetMessageById fetch a specific message

Use the **messageId** to fetch a message that's been sent from/to your account within the last 30 days.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId When you sent the original message, this is the UUID that was returned in the response. Use this ID to fetch, update or delete a message with the appropriate endpoints.   Incoming messages are also assigned a messageId. Use this ID with GET/ messages/{messageId} to fetch the message later.
	@return ApiGetMessageByIdRequest
*/
func (a *MessagesAPIService) GetMessageById(ctx context.Context, messageId string) ApiGetMessageByIdRequest {
	return ApiGetMessageByIdRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
//
//	@return MessageGet
func (a *MessagesAPIService) GetMessageByIdExecute(r ApiGetMessageByIdRequest) (*MessageGet, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MessageGet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.GetMessageById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

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
	if strlen(r.messageId) < 36 {
		return localVarReturnValue, nil, reportError("messageId must have at least 36 elements")
	}
	if strlen(r.messageId) > 36 {
		return localVarReturnValue, nil, reportError("messageId must have less than 36 elements")
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
		if localVarHTTPResponse.StatusCode == 406 {
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
		if localVarHTTPResponse.StatusCode == 415 {
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
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiGetMessagesRequest struct {
	ctx               context.Context
	ApiService        *MessagesAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	telstraApiVersion *string
	limit *int32
	offset *int32
	direction *string
	status *string
	filter *string
	reverse *bool
	startTime *time.Time
	endTime *time.Time
}

func (r ApiGetMessagesRequest) ContentLanguage(contentLanguage string) ApiGetMessagesRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetMessagesRequest) Authorization(authorization string) ApiGetMessagesRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetMessagesRequest) Accept(accept string) ApiGetMessagesRequest {
	r.accept = &accept
	return r
}

func (r ApiGetMessagesRequest) AcceptCharset(acceptCharset string) ApiGetMessagesRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetMessagesRequest) ContentType(contentType string) ApiGetMessagesRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetMessagesRequest) TelstraApiVersion(telstraApiVersion string) ApiGetMessagesRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

// Tell us how many results you want us to return, up to a maximum of 50.
func (r ApiGetMessagesRequest) Limit(limit int32) ApiGetMessagesRequest {
	r.limit = &limit
	return r
}

// Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.
func (r ApiGetMessagesRequest) Offset(offset int32) ApiGetMessagesRequest {
	r.offset = &offset
	return r
}

// Filter your messages by direction: * **outgoing** – messages sent from your account. * **incoming** – messages sent to your account.
func (r ApiGetMessagesRequest) Direction(direction string) ApiGetMessagesRequest {
	r.direction = &direction
	return r
}

// Filter your messages by one of the statuses below:  * **queued** – messages queued for sending or still in transit. * **sent** – messages that have been sent from the server. * **delivered** – messages successful delivered to the recipient&#39;s device. Note that we will only be able to return this status if you set deliveryNotification to **true** (paid feature).  * **expired** – message that couldn&#39;t be sent within the **retryTimeout** timeframe.
func (r ApiGetMessagesRequest) Status(status string) ApiGetMessagesRequest {
	r.status = &status
	return r
}

// Filter your messages by:  * tag - use one of the tags assigned to your message(s) * number - either the Virtual Number used to send the message, or the Recipient Number the message was sent to.
func (r ApiGetMessagesRequest) Filter(filter string) ApiGetMessagesRequest {
	r.filter = &filter
	return r
}

// If set to **true** the results will be returned in reverse order. Recent messages will be returned first. By default, the results will be returned in the order they were sent/received. 
func (r ApiGetMessagesRequest) Reverse(reverse bool) ApiGetMessagesRequest {
	r.reverse = &reverse
	return r
}

// By Default messages from last 30 days will be fetched.Use this parameter to fetch the messages from the time you want.  Set the time in London Greenwich Mean Time (adjusting for any time difference) and use ISO format, e.g. \&quot;2024-01-24T15:39:00Z\&quot;.  You can set a time from last 30 days. If you specify a timestamp outside of this limit, the API will return a FIELD_INVALID error. If startTime alone set, messages from this time till current time will be fetched. 
func (r ApiGetMessagesRequest) StartTime(startTime time.Time) ApiGetMessagesRequest {
	r.startTime = &startTime
	return r
}

// By Default messages from last 30 days will be fetched. Use this parameter to fetch the messages till the time you want.  Set the time in London Greenwich Mean Time (adjusting for any time difference) and use ISO format, e.g. \&quot;2024-01-24T16:39:00Z\&quot;.  You can set a time from last 30 days and endTime should be after startTime (if set). If you specify a timestamp outside of this limit, the API will return a FIELD_INVALID error. 
func (r ApiGetMessagesRequest) EndTime(endTime time.Time) ApiGetMessagesRequest {
	r.endTime = &endTime
	return r
}

func (r ApiGetMessagesRequest) Execute() (*GetMessages200Response, *http.Response, error) {
	return r.ApiService.GetMessagesExecute(r)
}

/*
GetMessages fetch all sent/received messages

Fetch messages that have been sent from/to your account in the last 30 days.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetMessagesRequest
*/
func (a *MessagesAPIService) GetMessages(ctx context.Context) ApiGetMessagesRequest {
	return ApiGetMessagesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetMessages200Response
func (a *MessagesAPIService) GetMessagesExecute(r ApiGetMessagesRequest) (*GetMessages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetMessages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.GetMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages"

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

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	} else {
		var defaultValue int32 = 10
		r.limit = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.direction != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.reverse != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "reverse", r.reverse, "")
	} else {
		var defaultValue bool = false
		r.reverse = &defaultValue
	}
	if r.startTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startTime", r.startTime, "")
	}
	if r.endTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endTime", r.endTime, "")
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

type ApiSendMessagesRequest struct {
	ctx                 context.Context
	ApiService          *MessagesAPIService
	contentLanguage     *string
	authorization       *string
	accept              *string
	acceptCharset       *string
	contentType         *string
	sendMessagesRequest *SendMessagesRequest
	telstraApiVersion   *string
}

func (r ApiSendMessagesRequest) ContentLanguage(contentLanguage string) ApiSendMessagesRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiSendMessagesRequest) Authorization(authorization string) ApiSendMessagesRequest {
	r.authorization = &authorization
	return r
}

func (r ApiSendMessagesRequest) Accept(accept string) ApiSendMessagesRequest {
	r.accept = &accept
	return r
}

func (r ApiSendMessagesRequest) AcceptCharset(acceptCharset string) ApiSendMessagesRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiSendMessagesRequest) ContentType(contentType string) ApiSendMessagesRequest {
	r.contentType = &contentType
	return r
}

func (r ApiSendMessagesRequest) SendMessagesRequest(sendMessagesRequest SendMessagesRequest) ApiSendMessagesRequest {
	r.sendMessagesRequest = &sendMessagesRequest
	return r
}

func (r ApiSendMessagesRequest) TelstraApiVersion(telstraApiVersion string) ApiSendMessagesRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiSendMessagesRequest) Execute() (*MessageSent, *http.Response, error) {
	return r.ApiService.SendMessagesExecute(r)
}

/*
SendMessages send messages

Send an SMS/MMS to a mobile number, or to multiple mobile numbers.

Free Trial users can message to up to 10 unique recipient numbers for free. The first five recipients will be automatically added to your Free Trial Numbers list. Need more? Just use the POST /free-trial-numbers call to add another five.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiSendMessagesRequest
*/
func (a *MessagesAPIService) SendMessages(ctx context.Context) ApiSendMessagesRequest {
	return ApiSendMessagesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return MessageSent
func (a *MessagesAPIService) SendMessagesExecute(r ApiSendMessagesRequest) (*MessageSent, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MessageSent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.SendMessages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages"

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
	if r.sendMessagesRequest == nil {
		return localVarReturnValue, nil, reportError("sendMessagesRequest is required and must be specified")
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
	localVarPostBody = r.sendMessagesRequest
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
		if localVarHTTPResponse.StatusCode == 413 {
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

type ApiUpdateMessageByIdRequest struct {
	ctx                      context.Context
	ApiService               *MessagesAPIService
	contentLanguage          *string
	authorization            *string
	accept                   *string
	acceptCharset            *string
	contentType              *string
	messageId                string
	updateMessageByIdRequest *UpdateMessageByIdRequest
	telstraApiVersion        *string
}

func (r ApiUpdateMessageByIdRequest) ContentLanguage(contentLanguage string) ApiUpdateMessageByIdRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiUpdateMessageByIdRequest) Authorization(authorization string) ApiUpdateMessageByIdRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateMessageByIdRequest) Accept(accept string) ApiUpdateMessageByIdRequest {
	r.accept = &accept
	return r
}

func (r ApiUpdateMessageByIdRequest) AcceptCharset(acceptCharset string) ApiUpdateMessageByIdRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiUpdateMessageByIdRequest) ContentType(contentType string) ApiUpdateMessageByIdRequest {
	r.contentType = &contentType
	return r
}

func (r ApiUpdateMessageByIdRequest) UpdateMessageByIdRequest(updateMessageByIdRequest UpdateMessageByIdRequest) ApiUpdateMessageByIdRequest {
	r.updateMessageByIdRequest = &updateMessageByIdRequest
	return r
}

func (r ApiUpdateMessageByIdRequest) TelstraApiVersion(telstraApiVersion string) ApiUpdateMessageByIdRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiUpdateMessageByIdRequest) Execute() (*MessageUpdate, *http.Response, error) {
	return r.ApiService.UpdateMessageByIdExecute(r)
}

/*
UpdateMessageById update a message

Need to update a message that's scheduled for sending? You can change any of the below parameters, as long as the message hasn't been sent yet. This request body will override the original POST/ messages call.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.
	@return ApiUpdateMessageByIdRequest
*/
func (a *MessagesAPIService) UpdateMessageById(ctx context.Context, messageId string) ApiUpdateMessageByIdRequest {
	return ApiUpdateMessageByIdRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
//
//	@return MessageUpdate
func (a *MessagesAPIService) UpdateMessageByIdExecute(r ApiUpdateMessageByIdRequest) (*MessageUpdate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *MessageUpdate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.UpdateMessageById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

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
	if strlen(r.messageId) < 36 {
		return localVarReturnValue, nil, reportError("messageId must have at least 36 elements")
	}
	if strlen(r.messageId) > 36 {
		return localVarReturnValue, nil, reportError("messageId must have less than 36 elements")
	}
	if r.updateMessageByIdRequest == nil {
		return localVarReturnValue, nil, reportError("updateMessageByIdRequest is required and must be specified")
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
	localVarPostBody = r.updateMessageByIdRequest
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
		if localVarHTTPResponse.StatusCode == 413 {
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

type ApiUpdateMessageTagsRequest struct {
	ctx                      context.Context
	ApiService               *MessagesAPIService
	contentLanguage          *string
	authorization            *string
	accept                   *string
	acceptCharset            *string
	contentType              *string
	messageId                string
	updateMessageTagsRequest *UpdateMessageTagsRequest
	telstraApiVersion        *string
}

func (r ApiUpdateMessageTagsRequest) ContentLanguage(contentLanguage string) ApiUpdateMessageTagsRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiUpdateMessageTagsRequest) Authorization(authorization string) ApiUpdateMessageTagsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateMessageTagsRequest) Accept(accept string) ApiUpdateMessageTagsRequest {
	r.accept = &accept
	return r
}

func (r ApiUpdateMessageTagsRequest) AcceptCharset(acceptCharset string) ApiUpdateMessageTagsRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiUpdateMessageTagsRequest) ContentType(contentType string) ApiUpdateMessageTagsRequest {
	r.contentType = &contentType
	return r
}

func (r ApiUpdateMessageTagsRequest) UpdateMessageTagsRequest(updateMessageTagsRequest UpdateMessageTagsRequest) ApiUpdateMessageTagsRequest {
	r.updateMessageTagsRequest = &updateMessageTagsRequest
	return r
}

func (r ApiUpdateMessageTagsRequest) TelstraApiVersion(telstraApiVersion string) ApiUpdateMessageTagsRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiUpdateMessageTagsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateMessageTagsExecute(r)
}

/*
UpdateMessageTags update message tags

Use the **messageId** to update the tag(s) assigned to a message. You can use this endpoint any time, even after your message has been delivered.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.
	@return ApiUpdateMessageTagsRequest
*/
func (a *MessagesAPIService) UpdateMessageTags(ctx context.Context, messageId string) ApiUpdateMessageTagsRequest {
	return ApiUpdateMessageTagsRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
func (a *MessagesAPIService) UpdateMessageTagsExecute(r ApiUpdateMessageTagsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.UpdateMessageTags")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.contentLanguage == nil {
		return nil, reportError("contentLanguage is required and must be specified")
	}
	if r.authorization == nil {
		return nil, reportError("authorization is required and must be specified")
	}
	if r.accept == nil {
		return nil, reportError("accept is required and must be specified")
	}
	if r.acceptCharset == nil {
		return nil, reportError("acceptCharset is required and must be specified")
	}
	if r.contentType == nil {
		return nil, reportError("contentType is required and must be specified")
	}
	if strlen(r.messageId) < 36 {
		return nil, reportError("messageId must have at least 36 elements")
	}
	if strlen(r.messageId) > 36 {
		return nil, reportError("messageId must have less than 36 elements")
	}
	if r.updateMessageTagsRequest == nil {
		return nil, reportError("updateMessageTagsRequest is required and must be specified")
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
	localVarPostBody = r.updateMessageTagsRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 405 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Errors
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		var v Errors
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
