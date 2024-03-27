# \AuthenticationAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthToken**](AuthenticationAPI.md#AuthToken) | **Post** /oauth/token | Generate an access token



## AuthToken

> OAuth AuthToken(ctx).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Scope(scope).Execute()

Generate an access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    clientId := "clientId_example" // string | Copy and paste your `client_id` here. Find it in your [dashboard](https://dev.telstra.com).
    clientSecret := "clientSecret_example" // string | Copy and paste your `client_secret` here. Find it in your [dashboard](https://dev.telstra.com).
    grantType := "grantType_example" // string | Ensure the `grant_type` is **client_credentials**. (default to "client_credentials")
    scope := "scope_example" // string | Ensure the `scope` is **verification:read**. (optional) (default to "verification:read")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).ClientId(clientId).ClientSecret(clientSecret).GrantType(grantType).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.AuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthToken`: OAuth
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.AuthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Copy and paste your &#x60;client_id&#x60; here. Find it in your [dashboard](https://dev.telstra.com). | 
 **clientSecret** | **string** | Copy and paste your &#x60;client_secret&#x60; here. Find it in your [dashboard](https://dev.telstra.com). | 
 **grantType** | **string** | Ensure the &#x60;grant_type&#x60; is **client_credentials**. | [default to &quot;client_credentials&quot;]
 **scope** | **string** | Ensure the &#x60;scope&#x60; is **verification:read**. | [default to &quot;verification:read&quot;]

### Return type

[**OAuth**](OAuth.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

