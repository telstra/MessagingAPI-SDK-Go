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

// AuthenticationAPIService AuthenticationAPI service
type AuthenticationAPIService service

type ApiAuthTokenRequest struct {
	ctx          context.Context
	ApiService   *AuthenticationAPIService
	clientId     *string
	clientSecret *string
	grantType    *string
	scope        *string
}

// Copy and paste your &#x60;client_id&#x60; here. Find it in your [dashboard](https://dev.telstra.com).
func (r ApiAuthTokenRequest) ClientId(clientId string) ApiAuthTokenRequest {
	r.clientId = &clientId
	return r
}

// Copy and paste your &#x60;client_secret&#x60; here. Find it in your [dashboard](https://dev.telstra.com).
func (r ApiAuthTokenRequest) ClientSecret(clientSecret string) ApiAuthTokenRequest {
	r.clientSecret = &clientSecret
	return r
}

// Ensure the &#x60;grant_type&#x60; is **client_credentials**.
func (r ApiAuthTokenRequest) GrantType(grantType string) ApiAuthTokenRequest {
	r.grantType = &grantType
	return r
}

// Ensure the &#x60;scope&#x60; is **verification:read**.
func (r ApiAuthTokenRequest) Scope(scope string) ApiAuthTokenRequest {
	r.scope = &scope
	return r
}

func (r ApiAuthTokenRequest) Execute() (*OAuth, *http.Response, error) {
	r.grantType = new(string)
	*r.grantType = setRequestParams().GrantType
	r.scope = new(string)
	*r.scope = setRequestParams().Scope

	return r.ApiService.AuthTokenExecute(r)
}

/*
AuthToken Generate an access token

An OAuth 2.0 access token is required to access the API features. To create a token, use the unique `client_id` and `client_secret` you received when you registered for the API. You can find these credentials [here](https://dev.telstra.com). Note that your access token will expire in 1 hour.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAuthTokenRequest
*/
func (a *AuthenticationAPIService) AuthToken(ctx context.Context) ApiAuthTokenRequest {
	return ApiAuthTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

func (r ApiAuthTokenRequest) AuthToken() (*OAuth, *http.Response, error) {
	return r.Execute()
}

// Execute executes the request
//
//	@return OAuth
func (a *AuthenticationAPIService) AuthTokenExecute(r ApiAuthTokenRequest) (*OAuth, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *OAuth
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationAPIService.AuthToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if strlen(*r.clientId) < 0 {
		return localVarReturnValue, nil, reportError("clientId must have at least 0 elements")
	}
	if strlen(*r.clientId) > 99999 {
		return localVarReturnValue, nil, reportError("clientId must have less than 99999 elements")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}
	if strlen(*r.clientSecret) < 0 {
		return localVarReturnValue, nil, reportError("clientSecret must have at least 0 elements")
	}
	if strlen(*r.clientSecret) > 99999 {
		return localVarReturnValue, nil, reportError("clientSecret must have less than 99999 elements")
	}
	if r.grantType == nil {
		return localVarReturnValue, nil, reportError("grantType is required and must be specified")
	}
	if strlen(*r.grantType) < 0 {
		return localVarReturnValue, nil, reportError("grantType must have at least 0 elements")
	}
	if strlen(*r.grantType) > 99999 {
		return localVarReturnValue, nil, reportError("grantType must have less than 99999 elements")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "client_secret", r.clientSecret, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "")
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "scope", r.scope, "")
	}
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
			var v AuthToken400Response
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
			var v AuthToken400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v AuthToken400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v AuthToken400Response
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
