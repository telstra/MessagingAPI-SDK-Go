# \ReportsAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReport**](ReportsAPI.md#GetReport) | **Get** /reports/{reportId} | fetch a specific report
[**GetReports**](ReportsAPI.md#GetReports) | **Get** /reports | fetch all reports
[**MessagesReport**](ReportsAPI.md#MessagesReport) | **Post** /reports/messages | submit a request for a messages report



## GetReport

> GetReport200Response GetReport(ctx, reportId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

fetch a specific report



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
    contentLanguage := "en-au" // string | 
    authorization := "Bearer <access_token>" // string | 
    accept := "application/json" // string | 
    acceptCharset := "utf-8" // string | 
    contentType := "application/json; charset=utf-8" // string | 
    reportId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Use the reportId returned in the POST /reports/{reportType} response. 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetReport(context.Background(), reportId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReport`: GetReport200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** | Use the reportId returned in the POST /reports/{reportType} response.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **telstraApiVersion** | **string** |  | 

### Return type

[**GetReport200Response**](GetReport200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReports

> GetReports200Response GetReports(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

fetch all reports



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
    contentLanguage := "en-au" // string | 
    authorization := "Bearer <access_token>" // string | 
    accept := "application/json" // string | 
    acceptCharset := "utf-8" // string | 
    contentType := "application/json; charset=utf-8" // string | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetReports(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReports`: GetReports200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**GetReports200Response**](GetReports200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MessagesReport

> MessagesReport201Response MessagesReport(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).MessagesReportRequest(messagesReportRequest).TelstraApiVersion(telstraApiVersion).Execute()

submit a request for a messages report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    contentLanguage := "en-au" // string | 
    authorization := "Bearer <access_token>" // string | 
    accept := "application/json" // string | 
    acceptCharset := "utf-8" // string | 
    contentType := "application/json; charset=utf-8" // string | 
    messagesReportRequest := *openapiclient.NewMessagesReportRequest(time.Now(), time.Now()) // MessagesReportRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.MessagesReport(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).MessagesReportRequest(messagesReportRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.MessagesReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MessagesReport`: MessagesReport201Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.MessagesReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMessagesReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **messagesReportRequest** | [**MessagesReportRequest**](MessagesReportRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**MessagesReport201Response**](MessagesReport201Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

