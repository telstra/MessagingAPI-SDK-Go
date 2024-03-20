# \LogsAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogs**](LogsAPI.md#GetLogs) | **Get** /logs | get logs



## GetLogs

> GetLogs200Response GetLogs(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).StartDate(startDate).EndDate(endDate).LogsFilter(logsFilter).Execute()

get logs



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
    telstraApiVersion := "3.1.0" // string |  (optional)
    startDate := time.Now() // time.Time | Fetch logs for calls made from a specific time and date. Default value is events logged within the last week. (optional)
    endDate := time.Now() // time.Time | Fetch logs for calls made until a specific time and date. Default value is current time and date (optional)
    logsFilter := "logsFilter_example" // string | Filter your logs by:  * date - rather than searching within a time range, you can choose to return logs for a specific date instead. Use string <date-time> format. * resource – return logs for a specific resource, e.g. POST /messages.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.GetLogs(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).StartDate(startDate).EndDate(endDate).LogsFilter(logsFilter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: GetLogs200Response
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **telstraApiVersion** | **string** |  | 
 **startDate** | **time.Time** | Fetch logs for calls made from a specific time and date. Default value is events logged within the last week. | 
 **endDate** | **time.Time** | Fetch logs for calls made until a specific time and date. Default value is current time and date | 
 **logsFilter** | **string** | Filter your logs by:  * date - rather than searching within a time range, you can choose to return logs for a specific date instead. Use string &lt;date-time&gt; format. * resource – return logs for a specific resource, e.g. POST /messages.  | 

### Return type

[**GetLogs200Response**](GetLogs200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

