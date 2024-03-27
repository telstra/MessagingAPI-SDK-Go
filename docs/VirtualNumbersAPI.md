# \VirtualNumbersAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignNumber**](VirtualNumbersAPI.md#AssignNumber) | **Post** /virtual-numbers | assign a virtual number
[**DeleteNumber**](VirtualNumbersAPI.md#DeleteNumber) | **Delete** /virtual-numbers/{virtual-number} | delete a virtual number
[**GetNumbers**](VirtualNumbersAPI.md#GetNumbers) | **Get** /virtual-numbers | fetch all virtual numbers
[**GetRecipientOptouts**](VirtualNumbersAPI.md#GetRecipientOptouts) | **Get** /virtual-numbers/{virtual-number}/optouts | Get recipient optouts list
[**GetVirtualNumber**](VirtualNumbersAPI.md#GetVirtualNumber) | **Get** /virtual-numbers/{virtual-number} | fetch a virtual number
[**UpdateNumber**](VirtualNumbersAPI.md#UpdateNumber) | **Put** /virtual-numbers/{virtual-number} | update a virtual number



## AssignNumber

> VirtualNumber AssignNumber(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).AssignNumberRequest(assignNumberRequest).TelstraApiVersion(telstraApiVersion).Execute()

assign a virtual number



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
    assignNumberRequest := *openapiclient.NewAssignNumberRequest() // AssignNumberRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualNumbersAPI.AssignNumber(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).AssignNumberRequest(assignNumberRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.AssignNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignNumber`: VirtualNumber
    fmt.Fprintf(os.Stdout, "Response from `VirtualNumbersAPI.AssignNumber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **assignNumberRequest** | [**AssignNumberRequest**](AssignNumberRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**VirtualNumber**](VirtualNumber.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNumber

> DeleteNumber(ctx, virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

delete a virtual number



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
    virtualNumber := "0412345678" // string | Write the Virtual Number here, using national format (e.g. 0412345678). 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VirtualNumbersAPI.DeleteNumber(context.Background(), virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.DeleteNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualNumber** | **string** | Write the Virtual Number here, using national format (e.g. 0412345678).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **telstraApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNumbers

> GetNumbers200Response GetNumbers(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Filter(filter).Execute()

fetch all virtual numbers



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
    limit := int32(56) // int32 | Tell us how many results you want us to return, up to a maximum of 50.  (optional) (default to 10)
    offset := int32(56) // int32 | Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.  (optional) (default to 0)
    filter := "filter_example" // string | Filter your Virtual Numbers by tag or by number. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualNumbersAPI.GetNumbers(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.GetNumbers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNumbers`: GetNumbers200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualNumbersAPI.GetNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **telstraApiVersion** | **string** |  | 
 **limit** | **int32** | Tell us how many results you want us to return, up to a maximum of 50.  | [default to 10]
 **offset** | **int32** | Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.  | [default to 0]
 **filter** | **string** | Filter your Virtual Numbers by tag or by number. | 

### Return type

[**GetNumbers200Response**](GetNumbers200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecipientOptouts

> GetRecipientOptouts200Response GetRecipientOptouts(ctx, virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Execute()

Get recipient optouts list



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
    virtualNumber := "0412345678" // string | Write the Virtual Number here, using national format (e.g. 0412345678). 
    telstraApiVersion := "3.1.0" // string |  (optional)
    limit := int32(56) // int32 | Tell us how many results you want us to return, up to a maximum of 50.  (optional) (default to 10)
    offset := int32(56) // int32 | Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualNumbersAPI.GetRecipientOptouts(context.Background(), virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.GetRecipientOptouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecipientOptouts`: GetRecipientOptouts200Response
    fmt.Fprintf(os.Stdout, "Response from `VirtualNumbersAPI.GetRecipientOptouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualNumber** | **string** | Write the Virtual Number here, using national format (e.g. 0412345678).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecipientOptoutsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **telstraApiVersion** | **string** |  | 
 **limit** | **int32** | Tell us how many results you want us to return, up to a maximum of 50.  | [default to 10]
 **offset** | **int32** | Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.  | [default to 0]

### Return type

[**GetRecipientOptouts200Response**](GetRecipientOptouts200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualNumber

> VirtualNumber GetVirtualNumber(ctx, virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

fetch a virtual number



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
    virtualNumber := "0412345678" // string | Write the Virtual Number here, using national format (e.g. 0412345678). 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualNumbersAPI.GetVirtualNumber(context.Background(), virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.GetVirtualNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualNumber`: VirtualNumber
    fmt.Fprintf(os.Stdout, "Response from `VirtualNumbersAPI.GetVirtualNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualNumber** | **string** | Write the Virtual Number here, using national format (e.g. 0412345678).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **telstraApiVersion** | **string** |  | 

### Return type

[**VirtualNumber**](VirtualNumber.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNumber

> VirtualNumber UpdateNumber(ctx, virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateNumberRequest(updateNumberRequest).TelstraApiVersion(telstraApiVersion).Execute()

update a virtual number



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
    virtualNumber := "0412345678" // string | Write the Virtual Number here, using national format (e.g. 0412345678). 
    updateNumberRequest := *openapiclient.NewUpdateNumberRequest() // UpdateNumberRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualNumbersAPI.UpdateNumber(context.Background(), virtualNumber).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateNumberRequest(updateNumberRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualNumbersAPI.UpdateNumber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNumber`: VirtualNumber
    fmt.Fprintf(os.Stdout, "Response from `VirtualNumbersAPI.UpdateNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**virtualNumber** | **string** | Write the Virtual Number here, using national format (e.g. 0412345678).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **updateNumberRequest** | [**UpdateNumberRequest**](UpdateNumberRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**VirtualNumber**](VirtualNumber.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

