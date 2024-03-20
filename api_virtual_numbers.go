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

// VirtualNumbersAPIService VirtualNumbersAPI service
type VirtualNumbersAPIService service

type ApiAssignNumberRequest struct {
	ctx                 context.Context
	ApiService          *VirtualNumbersAPIService
	contentLanguage     *string
	authorization       *string
	accept              *string
	acceptCharset       *string
	contentType         *string
	assignNumberRequest *AssignNumberRequest
	telstraApiVersion   *string
}

func (r ApiAssignNumberRequest) ContentLanguage(contentLanguage string) ApiAssignNumberRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiAssignNumberRequest) Authorization(authorization string) ApiAssignNumberRequest {
	r.authorization = &authorization
	return r
}

func (r ApiAssignNumberRequest) Accept(accept string) ApiAssignNumberRequest {
	r.accept = &accept
	return r
}

func (r ApiAssignNumberRequest) AcceptCharset(acceptCharset string) ApiAssignNumberRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiAssignNumberRequest) ContentType(contentType string) ApiAssignNumberRequest {
	r.contentType = &contentType
	return r
}

func (r ApiAssignNumberRequest) AssignNumberRequest(assignNumberRequest AssignNumberRequest) ApiAssignNumberRequest {
	r.assignNumberRequest = &assignNumberRequest
	return r
}

func (r ApiAssignNumberRequest) TelstraApiVersion(telstraApiVersion string) ApiAssignNumberRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiAssignNumberRequest) Execute() (*VirtualNumber, *http.Response, error) {
	return r.ApiService.AssignNumberExecute(r)
}

/*
AssignNumber assign a virtual number

When a recipient receives your message, you can choose whether they'll see a Virtual Number or senderName (paid plans only) in the **from** field. If you want to use a Virtual Number, use this endpoint to assign one. Free Trial users can assign one Virtual Number, and those on a paid plan can assign up to 100.

Virtual Numbers that have not sent a message in 30 days (Free Trial) or sent/received a message in 18 months (paid plans) will be automatically unassigned from your account. You can check the **lastUse** date of your Virtual Number at any time using GET /virtual-numbers/{virtual-number}.

Note that Virtual Numbers used in v2 of the Messaging API cannot be used to send messages in v3. Please assign a new Virtual Number instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAssignNumberRequest
*/
func (a *VirtualNumbersAPIService) AssignNumber(ctx context.Context) ApiAssignNumberRequest {
	return ApiAssignNumberRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return VirtualNumber
func (a *VirtualNumbersAPIService) AssignNumberExecute(r ApiAssignNumberRequest) (*VirtualNumber, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VirtualNumber
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.AssignNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers"

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
	if r.assignNumberRequest == nil {
		return localVarReturnValue, nil, reportError("assignNumberRequest is required and must be specified")
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
	localVarPostBody = r.assignNumberRequest
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

type ApiDeleteNumberRequest struct {
	ctx               context.Context
	ApiService        *VirtualNumbersAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	virtualNumber     string
	telstraApiVersion *string
}

func (r ApiDeleteNumberRequest) ContentLanguage(contentLanguage string) ApiDeleteNumberRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiDeleteNumberRequest) Authorization(authorization string) ApiDeleteNumberRequest {
	r.authorization = &authorization
	return r
}

func (r ApiDeleteNumberRequest) Accept(accept string) ApiDeleteNumberRequest {
	r.accept = &accept
	return r
}

func (r ApiDeleteNumberRequest) AcceptCharset(acceptCharset string) ApiDeleteNumberRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiDeleteNumberRequest) ContentType(contentType string) ApiDeleteNumberRequest {
	r.contentType = &contentType
	return r
}

func (r ApiDeleteNumberRequest) TelstraApiVersion(telstraApiVersion string) ApiDeleteNumberRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiDeleteNumberRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNumberExecute(r)
}

/*
DeleteNumber delete a virtual number

Use **virtual-number** to remove a Virtual Number that's been assigned to your account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualNumber Write the Virtual Number here, using national format (e.g. 0412345678).
	@return ApiDeleteNumberRequest
*/
func (a *VirtualNumbersAPIService) DeleteNumber(ctx context.Context, virtualNumber string) ApiDeleteNumberRequest {
	return ApiDeleteNumberRequest{
		ApiService:    a,
		ctx:           ctx,
		virtualNumber: virtualNumber,
	}
}

// Execute executes the request
func (a *VirtualNumbersAPIService) DeleteNumberExecute(r ApiDeleteNumberRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.DeleteNumber")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers/{virtual-number}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtual-number"+"}", url.PathEscape(parameterValueToString(r.virtualNumber, "virtualNumber")), -1)

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
	if strlen(r.virtualNumber) < 10 {
		return nil, reportError("virtualNumber must have at least 10 elements")
	}
	if strlen(r.virtualNumber) > 10 {
		return nil, reportError("virtualNumber must have less than 10 elements")
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

type ApiGetNumbersRequest struct {
	ctx               context.Context
	ApiService        *VirtualNumbersAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	telstraApiVersion *string
	limit             *int32
	offset            *int32
	filter            *string
}

func (r ApiGetNumbersRequest) ContentLanguage(contentLanguage string) ApiGetNumbersRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetNumbersRequest) Authorization(authorization string) ApiGetNumbersRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetNumbersRequest) Accept(accept string) ApiGetNumbersRequest {
	r.accept = &accept
	return r
}

func (r ApiGetNumbersRequest) AcceptCharset(acceptCharset string) ApiGetNumbersRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetNumbersRequest) ContentType(contentType string) ApiGetNumbersRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetNumbersRequest) TelstraApiVersion(telstraApiVersion string) ApiGetNumbersRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

// Tell us how many results you want us to return, up to a maximum of 50.
func (r ApiGetNumbersRequest) Limit(limit int32) ApiGetNumbersRequest {
	r.limit = &limit
	return r
}

// Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.
func (r ApiGetNumbersRequest) Offset(offset int32) ApiGetNumbersRequest {
	r.offset = &offset
	return r
}

// Filter your Virtual Numbers by tag or by number.
func (r ApiGetNumbersRequest) Filter(filter string) ApiGetNumbersRequest {
	r.filter = &filter
	return r
}

func (r ApiGetNumbersRequest) Execute() (*GetNumbers200Response, *http.Response, error) {
	return r.ApiService.GetNumbersExecute(r)
}

/*
GetNumbers fetch all virtual numbers

Use this endpoint to fetch all Virtual Numbers currently assigned to your account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetNumbersRequest
*/
func (a *VirtualNumbersAPIService) GetNumbers(ctx context.Context) ApiGetNumbersRequest {
	return ApiGetNumbersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return GetNumbers200Response
func (a *VirtualNumbersAPIService) GetNumbersExecute(r ApiGetNumbersRequest) (*GetNumbers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetNumbers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.GetNumbers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers"

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
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
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

type ApiGetRecipientOptoutsRequest struct {
	ctx               context.Context
	ApiService        *VirtualNumbersAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	virtualNumber     string
	telstraApiVersion *string
	limit             *int32
	offset            *int32
}

func (r ApiGetRecipientOptoutsRequest) ContentLanguage(contentLanguage string) ApiGetRecipientOptoutsRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetRecipientOptoutsRequest) Authorization(authorization string) ApiGetRecipientOptoutsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetRecipientOptoutsRequest) Accept(accept string) ApiGetRecipientOptoutsRequest {
	r.accept = &accept
	return r
}

func (r ApiGetRecipientOptoutsRequest) AcceptCharset(acceptCharset string) ApiGetRecipientOptoutsRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetRecipientOptoutsRequest) ContentType(contentType string) ApiGetRecipientOptoutsRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetRecipientOptoutsRequest) TelstraApiVersion(telstraApiVersion string) ApiGetRecipientOptoutsRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

// Tell us how many results you want us to return, up to a maximum of 50.
func (r ApiGetRecipientOptoutsRequest) Limit(limit int32) ApiGetRecipientOptoutsRequest {
	r.limit = &limit
	return r
}

// Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.
func (r ApiGetRecipientOptoutsRequest) Offset(offset int32) ApiGetRecipientOptoutsRequest {
	r.offset = &offset
	return r
}

func (r ApiGetRecipientOptoutsRequest) Execute() (*GetRecipientOptouts200Response, *http.Response, error) {
	return r.ApiService.GetRecipientOptoutsExecute(r)
}

/*
GetRecipientOptouts Get recipient optouts list

Use this endpoint to fetch any mobile number(s) that have opted out of receiving messages from a Virtual Number assigned to your account.

Recipients can opt out at any time by sending a message with industry standard keywords such as STOP, STOPALL, UNSUBSCRIBE, QUIT, END and CANCEL.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualNumber Write the Virtual Number here, using national format (e.g. 0412345678).
	@return ApiGetRecipientOptoutsRequest
*/
func (a *VirtualNumbersAPIService) GetRecipientOptouts(ctx context.Context, virtualNumber string) ApiGetRecipientOptoutsRequest {
	return ApiGetRecipientOptoutsRequest{
		ApiService:    a,
		ctx:           ctx,
		virtualNumber: virtualNumber,
	}
}

// Execute executes the request
//
//	@return GetRecipientOptouts200Response
func (a *VirtualNumbersAPIService) GetRecipientOptoutsExecute(r ApiGetRecipientOptoutsRequest) (*GetRecipientOptouts200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetRecipientOptouts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.GetRecipientOptouts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers/{virtual-number}/optouts"
	localVarPath = strings.Replace(localVarPath, "{"+"virtual-number"+"}", url.PathEscape(parameterValueToString(r.virtualNumber, "virtualNumber")), -1)

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
	if strlen(r.virtualNumber) < 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have at least 10 elements")
	}
	if strlen(r.virtualNumber) > 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have less than 10 elements")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiGetVirtualNumberRequest struct {
	ctx               context.Context
	ApiService        *VirtualNumbersAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	virtualNumber     string
	telstraApiVersion *string
}

func (r ApiGetVirtualNumberRequest) ContentLanguage(contentLanguage string) ApiGetVirtualNumberRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetVirtualNumberRequest) Authorization(authorization string) ApiGetVirtualNumberRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetVirtualNumberRequest) Accept(accept string) ApiGetVirtualNumberRequest {
	r.accept = &accept
	return r
}

func (r ApiGetVirtualNumberRequest) AcceptCharset(acceptCharset string) ApiGetVirtualNumberRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetVirtualNumberRequest) ContentType(contentType string) ApiGetVirtualNumberRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetVirtualNumberRequest) TelstraApiVersion(telstraApiVersion string) ApiGetVirtualNumberRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiGetVirtualNumberRequest) Execute() (*VirtualNumber, *http.Response, error) {
	return r.ApiService.GetVirtualNumberExecute(r)
}

/*
GetVirtualNumber fetch a virtual number

Fetch the tags, replyCallbackUrl and lastUse date for a Virtual Number.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualNumber Write the Virtual Number here, using national format (e.g. 0412345678).
	@return ApiGetVirtualNumberRequest
*/
func (a *VirtualNumbersAPIService) GetVirtualNumber(ctx context.Context, virtualNumber string) ApiGetVirtualNumberRequest {
	return ApiGetVirtualNumberRequest{
		ApiService:    a,
		ctx:           ctx,
		virtualNumber: virtualNumber,
	}
}

// Execute executes the request
//
//	@return VirtualNumber
func (a *VirtualNumbersAPIService) GetVirtualNumberExecute(r ApiGetVirtualNumberRequest) (*VirtualNumber, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VirtualNumber
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.GetVirtualNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers/{virtual-number}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtual-number"+"}", url.PathEscape(parameterValueToString(r.virtualNumber, "virtualNumber")), -1)

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
	if strlen(r.virtualNumber) < 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have at least 10 elements")
	}
	if strlen(r.virtualNumber) > 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have less than 10 elements")
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

type ApiUpdateNumberRequest struct {
	ctx                 context.Context
	ApiService          *VirtualNumbersAPIService
	contentLanguage     *string
	authorization       *string
	accept              *string
	acceptCharset       *string
	contentType         *string
	virtualNumber       string
	updateNumberRequest *UpdateNumberRequest
	telstraApiVersion   *string
}

func (r ApiUpdateNumberRequest) ContentLanguage(contentLanguage string) ApiUpdateNumberRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiUpdateNumberRequest) Authorization(authorization string) ApiUpdateNumberRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateNumberRequest) Accept(accept string) ApiUpdateNumberRequest {
	r.accept = &accept
	return r
}

func (r ApiUpdateNumberRequest) AcceptCharset(acceptCharset string) ApiUpdateNumberRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiUpdateNumberRequest) ContentType(contentType string) ApiUpdateNumberRequest {
	r.contentType = &contentType
	return r
}

func (r ApiUpdateNumberRequest) UpdateNumberRequest(updateNumberRequest UpdateNumberRequest) ApiUpdateNumberRequest {
	r.updateNumberRequest = &updateNumberRequest
	return r
}

func (r ApiUpdateNumberRequest) TelstraApiVersion(telstraApiVersion string) ApiUpdateNumberRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiUpdateNumberRequest) Execute() (*VirtualNumber, *http.Response, error) {
	return r.ApiService.UpdateNumberExecute(r)
}

/*
UpdateNumber update a virtual number

Use **virtual-number** to update the tags and/or replyCallbackUrl of a Virtual Number.

This request body will override the original POST/ virtual-numbers call.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param virtualNumber Write the Virtual Number here, using national format (e.g. 0412345678).
	@return ApiUpdateNumberRequest
*/
func (a *VirtualNumbersAPIService) UpdateNumber(ctx context.Context, virtualNumber string) ApiUpdateNumberRequest {
	return ApiUpdateNumberRequest{
		ApiService:    a,
		ctx:           ctx,
		virtualNumber: virtualNumber,
	}
}

// Execute executes the request
//
//	@return VirtualNumber
func (a *VirtualNumbersAPIService) UpdateNumberExecute(r ApiUpdateNumberRequest) (*VirtualNumber, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *VirtualNumber
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VirtualNumbersAPIService.UpdateNumber")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/virtual-numbers/{virtual-number}"
	localVarPath = strings.Replace(localVarPath, "{"+"virtual-number"+"}", url.PathEscape(parameterValueToString(r.virtualNumber, "virtualNumber")), -1)

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
	if strlen(r.virtualNumber) < 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have at least 10 elements")
	}
	if strlen(r.virtualNumber) > 10 {
		return localVarReturnValue, nil, reportError("virtualNumber must have less than 10 elements")
	}
	if r.updateNumberRequest == nil {
		return localVarReturnValue, nil, reportError("updateNumberRequest is required and must be specified")
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
	localVarPostBody = r.updateNumberRequest
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
