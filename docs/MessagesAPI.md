# \MessagesAPI

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessageById**](MessagesAPI.md#DeleteMessageById) | **Delete** /messages/{messageId} | delete a message
[**GetMessageById**](MessagesAPI.md#GetMessageById) | **Get** /messages/{messageId} | fetch a specific message
[**GetMessages**](MessagesAPI.md#GetMessages) | **Get** /messages | fetch all sent/received messages
[**SendMessages**](MessagesAPI.md#SendMessages) | **Post** /messages | send messages
[**UpdateMessageById**](MessagesAPI.md#UpdateMessageById) | **Put** /messages/{messageId} | update a message
[**UpdateMessageTags**](MessagesAPI.md#UpdateMessageTags) | **Patch** /messages/{messageId} | update message tags



## DeleteMessageById

> DeleteMessageById(ctx, messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

delete a message



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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints. 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessagesAPI.DeleteMessageById(context.Background(), messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.DeleteMessageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageByIdRequest struct via the builder pattern


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


## GetMessageById

> MessageGet GetMessageById(ctx, messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()

fetch a specific message



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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | When you sent the original message, this is the UUID that was returned in the response. Use this ID to fetch, update or delete a message with the appropriate endpoints.   Incoming messages are also assigned a messageId. Use this ID with GET/ messages/{messageId} to fetch the message later. 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.GetMessageById(context.Background(), messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageById`: MessageGet
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | When you sent the original message, this is the UUID that was returned in the response. Use this ID to fetch, update or delete a message with the appropriate endpoints.   Incoming messages are also assigned a messageId. Use this ID with GET/ messages/{messageId} to fetch the message later.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **telstraApiVersion** | **string** |  | 

### Return type

[**MessageGet**](MessageGet.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessages

> GetMessages200Response GetMessages(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Direction(direction).Status(status).Filter(filter).Execute()

fetch all sent/received messages



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
    direction := "direction_example" // string | Filter your messages by direction: * **outgoing** – messages sent from your account. * **incoming** – messages sent to your account.  (optional)
    status := "status_example" // string | Filter your messages by one of the statuses below:  * **queued** – messages queued for sending or still in transit. * **sent** – messages that have been sent from the server. * **delivered** – messages successful delivered to the recipient's device. Note that we will only be able to return this status if you set deliveryNotification to **true** (paid feature).  * **expired** – message that couldn't be sent within the **retryTimeout** timeframe.  (optional)
    filter := "filter_example" // string | Filter your messages by:  * tag - use one of the tags assigned to your message(s) * number - either the Virtual Number used to send the message, or the Recipient Number the message was sent to.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.GetMessages(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).TelstraApiVersion(telstraApiVersion).Limit(limit).Offset(offset).Direction(direction).Status(status).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.GetMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessages`: GetMessages200Response
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.GetMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesRequest struct via the builder pattern


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
 **direction** | **string** | Filter your messages by direction: * **outgoing** – messages sent from your account. * **incoming** – messages sent to your account.  | 
 **status** | **string** | Filter your messages by one of the statuses below:  * **queued** – messages queued for sending or still in transit. * **sent** – messages that have been sent from the server. * **delivered** – messages successful delivered to the recipient&#39;s device. Note that we will only be able to return this status if you set deliveryNotification to **true** (paid feature).  * **expired** – message that couldn&#39;t be sent within the **retryTimeout** timeframe.  | 
 **filter** | **string** | Filter your messages by:  * tag - use one of the tags assigned to your message(s) * number - either the Virtual Number used to send the message, or the Recipient Number the message was sent to.  | 

### Return type

[**GetMessages200Response**](GetMessages200Response.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessages

> MessageSent SendMessages(ctx).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).SendMessagesRequest(sendMessagesRequest).TelstraApiVersion(telstraApiVersion).Execute()

send messages



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
    sendMessagesRequest := *openapiclient.NewSendMessagesRequest(openapiclient.sendMessages_request_to{ArrayOfString: new([]string)}, "From_example") // SendMessagesRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.SendMessages(context.Background()).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).SendMessagesRequest(sendMessagesRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.SendMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMessages`: MessageSent
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.SendMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 
 **sendMessagesRequest** | [**SendMessagesRequest**](SendMessagesRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**MessageSent**](MessageSent.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageById

> MessageUpdate UpdateMessageById(ctx, messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateMessageByIdRequest(updateMessageByIdRequest).TelstraApiVersion(telstraApiVersion).Execute()

update a message



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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints. 
    updateMessageByIdRequest := *openapiclient.NewUpdateMessageByIdRequest("To_example", "From_example") // UpdateMessageByIdRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.UpdateMessageById(context.Background(), messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateMessageByIdRequest(updateMessageByIdRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessageById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMessageById`: MessageUpdate
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.UpdateMessageById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **updateMessageByIdRequest** | [**UpdateMessageByIdRequest**](UpdateMessageByIdRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

[**MessageUpdate**](MessageUpdate.md)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMessageTags

> UpdateMessageTags(ctx, messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateMessageTagsRequest(updateMessageTagsRequest).TelstraApiVersion(telstraApiVersion).Execute()

update message tags



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
    messageId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints. 
    updateMessageTagsRequest := *openapiclient.NewUpdateMessageTagsRequest([]string{"Tags_example"}) // UpdateMessageTagsRequest | 
    telstraApiVersion := "3.1.0" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MessagesAPI.UpdateMessageTags(context.Background(), messageId).ContentLanguage(contentLanguage).Authorization(authorization).Accept(accept).AcceptCharset(acceptCharset).ContentType(contentType).UpdateMessageTagsRequest(updateMessageTagsRequest).TelstraApiVersion(telstraApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.UpdateMessageTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | When you sent the original message, this is the UUID that was returned in the call response. Use this ID to fetch, update or delete a message with the appropriate endpoints.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMessageTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentLanguage** | **string** |  | 
 **authorization** | **string** |  | 
 **accept** | **string** |  | 
 **acceptCharset** | **string** |  | 
 **contentType** | **string** |  | 

 **updateMessageTagsRequest** | [**UpdateMessageTagsRequest**](UpdateMessageTagsRequest.md) |  | 
 **telstraApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

