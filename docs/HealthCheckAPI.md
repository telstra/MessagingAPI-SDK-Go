# \HealthCheckAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthCheck**](HealthCheckAPI.md#HealthCheck) | **Get** /health-check | health check



## HealthCheck

> HealthCheck200Response HealthCheck(ctx).Authorization(authorization).TelstraApiVersion(telstraApiVersion).Execute()

health check



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
    authorization := "Bearer <access_token>" // string | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthCheckAPI.HealthCheck(context.Background()).Authorization(authorization).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthCheckAPI.HealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HealthCheck`: HealthCheck200Response
    fmt.Fprintf(os.Stdout, "Response from `HealthCheckAPI.HealthCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**HealthCheck200Response**](HealthCheck200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

