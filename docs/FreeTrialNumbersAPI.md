# \FreeTrialNumbersAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrialNumbers**](FreeTrialNumbersAPI.md#CreateTrialNumbers) | **Post** /free-trial-numbers | create free trial number list
[**GetTrialNumbers**](FreeTrialNumbersAPI.md#GetTrialNumbers) | **Get** /free-trial-numbers | get all free trial numbers



## CreateTrialNumbers

> FreeTrialNumbers CreateTrialNumbers(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).CreateTrialNumbersRequest(createTrialNumbersRequest).TelstraApiVersion(telstraApiVersion).Execute()

create free trial number list



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
    createTrialNumbersRequest := *openapiclient.NewCreateTrialNumbersRequest(openapiclient.createTrialNumbers_request_freeTrialNumbers{ArrayOfString: new([]string)}) // CreateTrialNumbersRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FreeTrialNumbersAPI.CreateTrialNumbers(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).CreateTrialNumbersRequest(createTrialNumbersRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeTrialNumbersAPI.CreateTrialNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrialNumbers`: FreeTrialNumbers
    fmt.Fprintf(os.Stdout, "Response from `FreeTrialNumbersAPI.CreateTrialNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTrialNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **createTrialNumbersRequest** | [**CreateTrialNumbersRequest**](CreateTrialNumbersRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**FreeTrialNumbers**](FreeTrialNumbers.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrialNumbers

> FreeTrialNumbers GetTrialNumbers(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

get all free trial numbers



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
    resp, r, err := apiClient.FreeTrialNumbersAPI.GetTrialNumbers(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FreeTrialNumbersAPI.GetTrialNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrialNumbers`: FreeTrialNumbers
    fmt.Fprintf(os.Stdout, "Response from `FreeTrialNumbersAPI.GetTrialNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**FreeTrialNumbers**](FreeTrialNumbers.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

