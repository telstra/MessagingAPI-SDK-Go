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
)

// FreeTrialNumbersAPIService FreeTrialNumbersAPI service
type FreeTrialNumbersAPIService service

type ApiCreateTrialNumbersRequest struct {
	ctx                       context.Context
	ApiService                *FreeTrialNumbersAPIService
	contentLanguage           *string
	authorization             *string
	accept                    *string
	acceptCharset             *string
	contentType               *string
	createTrialNumbersRequest *CreateTrialNumbersRequest
	telstraApiVersion         *string
}

func (r ApiCreateTrialNumbersRequest) ContentLanguage(contentLanguage string) ApiCreateTrialNumbersRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiCreateTrialNumbersRequest) Authorization(authorization string) ApiCreateTrialNumbersRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateTrialNumbersRequest) Accept(accept string) ApiCreateTrialNumbersRequest {
	r.accept = &accept
	return r
}

func (r ApiCreateTrialNumbersRequest) AcceptCharset(acceptCharset string) ApiCreateTrialNumbersRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiCreateTrialNumbersRequest) ContentType(contentType string) ApiCreateTrialNumbersRequest {
	r.contentType = &contentType
	return r
}

func (r ApiCreateTrialNumbersRequest) CreateTrialNumbersRequest(createTrialNumbersRequest CreateTrialNumbersRequest) ApiCreateTrialNumbersRequest {
	r.createTrialNumbersRequest = &createTrialNumbersRequest
	return r
}

func (r ApiCreateTrialNumbersRequest) TelstraApiVersion(telstraApiVersion string) ApiCreateTrialNumbersRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiCreateTrialNumbersRequest) Execute() (*FreeTrialNumbers, *http.Response, error) {
	return r.ApiService.CreateTrialNumbersExecute(r)
}

/*
CreateTrialNumbers create free trial number list

Your Free Trial Numbers are the 10 recipient mobile numbers that you can message during the Free Trial. The first five numbers you send an SMS/MMS to will automatically be added to your Free Trial Numbers list. After that, you can use this endpoint to register another five. Alternatively, you can use this endpoint to register all 10 numbers.

Use this endpoint to register a Free Trial Number to your account. To test out all the features that the trial has to offer, we recommend registering your own mobile number to your Free Trial Numbers list.

Note that you can only message mobile numbers that have been added to your Free Trial list and once registered, a Free Trial Number cannot be removed or replaced.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTrialNumbersRequest
*/
func (a *FreeTrialNumbersAPIService) CreateTrialNumbers(ctx context.Context) ApiCreateTrialNumbersRequest {
	return ApiCreateTrialNumbersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FreeTrialNumbers
func (a *FreeTrialNumbersAPIService) CreateTrialNumbersExecute(r ApiCreateTrialNumbersRequest) (*FreeTrialNumbers, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FreeTrialNumbers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeTrialNumbersAPIService.CreateTrialNumbers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free-trial-numbers"

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
	if r.createTrialNumbersRequest == nil {
		return localVarReturnValue, nil, reportError("createTrialNumbersRequest is required and must be specified")
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
	localVarPostBody = r.createTrialNumbersRequest
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

type ApiGetTrialNumbersRequest struct {
	ctx               context.Context
	ApiService        *FreeTrialNumbersAPIService
	contentLanguage   *string
	authorization     *string
	accept            *string
	acceptCharset     *string
	contentType       *string
	telstraApiVersion *string
}

func (r ApiGetTrialNumbersRequest) ContentLanguage(contentLanguage string) ApiGetTrialNumbersRequest {
	r.contentLanguage = &contentLanguage
	return r
}

func (r ApiGetTrialNumbersRequest) Authorization(authorization string) ApiGetTrialNumbersRequest {
	r.authorization = &authorization
	return r
}

func (r ApiGetTrialNumbersRequest) Accept(accept string) ApiGetTrialNumbersRequest {
	r.accept = &accept
	return r
}

func (r ApiGetTrialNumbersRequest) AcceptCharset(acceptCharset string) ApiGetTrialNumbersRequest {
	r.acceptCharset = &acceptCharset
	return r
}

func (r ApiGetTrialNumbersRequest) ContentType(contentType string) ApiGetTrialNumbersRequest {
	r.contentType = &contentType
	return r
}

func (r ApiGetTrialNumbersRequest) TelstraApiVersion(telstraApiVersion string) ApiGetTrialNumbersRequest {
	r.telstraApiVersion = &telstraApiVersion
	return r
}

func (r ApiGetTrialNumbersRequest) Execute() (*FreeTrialNumbers, *http.Response, error) {
	return r.ApiService.GetTrialNumbersExecute(r)
}

/*
GetTrialNumbers get all free trial numbers

Use this endpoint to fetch the Free Trial Number(s) currently assigned to your account. These are the mobile numbers that you can message during the Free Trial.

If you're using a paid plan, there's no limit to the number of recipients that you can message, so you don't need to register Free Trial Numbers.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetTrialNumbersRequest
*/
func (a *FreeTrialNumbersAPIService) GetTrialNumbers(ctx context.Context) ApiGetTrialNumbersRequest {
	return ApiGetTrialNumbersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FreeTrialNumbers
func (a *FreeTrialNumbersAPIService) GetTrialNumbersExecute(r ApiGetTrialNumbersRequest) (*FreeTrialNumbers, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *FreeTrialNumbers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FreeTrialNumbersAPIService.GetTrialNumbers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/free-trial-numbers"

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
