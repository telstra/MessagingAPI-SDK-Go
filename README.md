# Go API client for TelstraMessaging

Send and receive SMS & MMS programmatically, leveraging Australia's leading mobile network.
With Telstra's Messaging API, we take out the complexity to allow seamless messaging integration into your app, with just a few lines of code.
Our REST API is enterprise grade, allowing you to communicate with engaging SMS & MMS messaging in your web and mobile apps in near real-time on a global scale.


## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import TelstraMessaging "github.com/telstra/MessagingAPI-SDK-Go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), TelstraMessaging.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), TelstraMessaging.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), TelstraMessaging.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), TelstraMessaging.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://products.api.telstra.com/messaging/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationAPI* | [**AuthToken**](docs/AuthenticationAPI.md#authtoken) | **Post** /oauth/token | Generate an access token
*FreeTrialNumbersAPI* | [**CreateTrialNumbers**](docs/FreeTrialNumbersAPI.md#createtrialnumbers) | **Post** /free-trial-numbers | create free trial number list
*FreeTrialNumbersAPI* | [**GetTrialNumbers**](docs/FreeTrialNumbersAPI.md#gettrialnumbers) | **Get** /free-trial-numbers | get all free trial numbers
*HealthCheckAPI* | [**HealthCheck**](docs/HealthCheckAPI.md#healthcheck) | **Get** /health-check | health check
*LogsAPI* | [**GetLogs**](docs/LogsAPI.md#getlogs) | **Get** /logs | get logs
*MessagesAPI* | [**DeleteMessageById**](docs/MessagesAPI.md#deletemessagebyid) | **Delete** /messages/{messageId} | delete a message
*MessagesAPI* | [**GetMessageById**](docs/MessagesAPI.md#getmessagebyid) | **Get** /messages/{messageId} | fetch a specific message
*MessagesAPI* | [**GetMessages**](docs/MessagesAPI.md#getmessages) | **Get** /messages | fetch all sent/received messages
*MessagesAPI* | [**SendMessages**](docs/MessagesAPI.md#sendmessages) | **Post** /messages | send messages
*MessagesAPI* | [**UpdateMessageById**](docs/MessagesAPI.md#updatemessagebyid) | **Put** /messages/{messageId} | update a message
*MessagesAPI* | [**UpdateMessageTags**](docs/MessagesAPI.md#updatemessagetags) | **Patch** /messages/{messageId} | update message tags
*ReportsAPI* | [**GetReport**](docs/ReportsAPI.md#getreport) | **Get** /reports/{reportId} | fetch a specific report
*ReportsAPI* | [**GetReports**](docs/ReportsAPI.md#getreports) | **Get** /reports | fetch all reports
*ReportsAPI* | [**MessagesReport**](docs/ReportsAPI.md#messagesreport) | **Post** /reports/messages | submit a request for a messages report
*VirtualNumbersAPI* | [**AssignNumber**](docs/VirtualNumbersAPI.md#assignnumber) | **Post** /virtual-numbers | assign a virtual number
*VirtualNumbersAPI* | [**DeleteNumber**](docs/VirtualNumbersAPI.md#deletenumber) | **Delete** /virtual-numbers/{virtual-number} | delete a virtual number
*VirtualNumbersAPI* | [**GetNumbers**](docs/VirtualNumbersAPI.md#getnumbers) | **Get** /virtual-numbers | fetch all virtual numbers
*VirtualNumbersAPI* | [**GetRecipientOptouts**](docs/VirtualNumbersAPI.md#getrecipientoptouts) | **Get** /virtual-numbers/{virtual-number}/optouts | Get recipient optouts list
*VirtualNumbersAPI* | [**GetVirtualNumber**](docs/VirtualNumbersAPI.md#getvirtualnumber) | **Get** /virtual-numbers/{virtual-number} | fetch a virtual number
*VirtualNumbersAPI* | [**UpdateNumber**](docs/VirtualNumbersAPI.md#updatenumber) | **Put** /virtual-numbers/{virtual-number} | update a virtual number


## Documentation For Models

 - [AssignNumberRequest](docs/AssignNumberRequest.md)
 - [AuthToken400Response](docs/AuthToken400Response.md)
 - [CreateTrialNumbersRequest](docs/CreateTrialNumbersRequest.md)
 - [CreateTrialNumbersRequestFreeTrialNumbers](docs/CreateTrialNumbersRequestFreeTrialNumbers.md)
 - [Error](docs/Error.md)
 - [Errors](docs/Errors.md)
 - [FreeTrialNumbers](docs/FreeTrialNumbers.md)
 - [GetLogs200Response](docs/GetLogs200Response.md)
 - [GetMessages200Response](docs/GetMessages200Response.md)
 - [GetMessages200ResponsePaging](docs/GetMessages200ResponsePaging.md)
 - [GetNumbers200Response](docs/GetNumbers200Response.md)
 - [GetRecipientOptouts200Response](docs/GetRecipientOptouts200Response.md)
 - [GetReport200Response](docs/GetReport200Response.md)
 - [GetReports200Response](docs/GetReports200Response.md)
 - [GetReports200ResponseReportsInner](docs/GetReports200ResponseReportsInner.md)
 - [HealthCheck200Response](docs/HealthCheck200Response.md)
 - [Log](docs/Log.md)
 - [MessageGet](docs/MessageGet.md)
 - [MessageSent](docs/MessageSent.md)
 - [MessageSentMessageId](docs/MessageSentMessageId.md)
 - [MessageSentTo](docs/MessageSentTo.md)
 - [MessageUpdate](docs/MessageUpdate.md)
 - [MessagesReport201Response](docs/MessagesReport201Response.md)
 - [MessagesReportRequest](docs/MessagesReportRequest.md)
 - [Multimedia](docs/Multimedia.md)
 - [MultimediaGet](docs/MultimediaGet.md)
 - [OAuth](docs/OAuth.md)
 - [RecipientOptout](docs/RecipientOptout.md)
 - [SendMessagesRequest](docs/SendMessagesRequest.md)
 - [SendMessagesRequestTo](docs/SendMessagesRequestTo.md)
 - [UpdateMessageByIdRequest](docs/UpdateMessageByIdRequest.md)
 - [UpdateMessageTagsRequest](docs/UpdateMessageTagsRequest.md)
 - [UpdateNumberRequest](docs/UpdateNumberRequest.md)
 - [VirtualNumber](docs/VirtualNumber.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### bearer_auth

- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
  - free-trial-numbers:read: read information for free trial numbers
  - free-trial-numbers:write: write information for free trial numbers
  - messages:read: read information for messages
  - messages:write: write information for messages
  - reports:read: read information for reports
  - reports:write: write information for reports
  - virtual-numbers:read: read information for virtual-numbers
  - virtual-numbers:write: write information for virtual numbers


## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issues.

## Author

## Documentation for Authorization

Authentication schemes defined for the API:
### bearer_auth

- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
  - free-trial-numbers:read: read information for free trial numbers
  - free-trial-numbers:write: write information for free trial numbers
  - messages:read: read information for messages
  - messages:write: write information for messages
  - reports:read: read information for reports
  - reports:write: write information for reports
  - virtual-numbers:read: read information for virtual-numbers
  - virtual-numbers:write: write information for virtual numbers


```go
import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

fmt.Printf("Access Token: %s\n", *oauthResult.AccessToken)
fmt.Printf("OAuth Result: %+v\n", oauthResult)

```

## Recommendation

It's recommended to create an instance of `ApiClient` per thread in a multithreaded environment to avoid any potential issues.

## Author

## Free Trial

Telstra offers a free trial for the messaging API to help you evaluate whether
it meets your needs. There are some restrictions that apply compared to the
full API, including a maximum number of messages that can be sent and requiring the
registration of a limited number of destinations before a message can be sent to that
destination. For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#FreeTrial>.

### Registering Free Trial Numbers

> :information_source: **Only required for the free trial accounts**

Register destinations for the free trial. For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#RegisteraFreeTrialNumber>.

The function `freeTrialNumbers.create` can be used to register destinations.

It takes an object with following properties as argument:

-   `freeTrialNumbers`: A list of destinations, expected to be phone numbers of the form `04XXXXXXXX`.

It returns the list of phone numbers that have been registered.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
trialNumbers := []string{"0400000001", "0400000002"}
createTrialNumbersRequestFreeTrialNumbers := msg_sdk.ArrayOfStringAsCreateTrialNumbersRequestFreeTrialNumbers(&trialNumbers)
createTrialNumbersRequest := msg_sdk.NewCreateTrialNumbersRequest(createTrialNumbersRequestFreeTrialNumbers)
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"
	
resp, httpRes, err := apiClient.FreeTrialNumbersAPI.CreateTrialNumbers(context.Background()).
	ContentType(contentType).
	Authorization(authorization).
	Accept(accept).
	AcceptCharset(acceptCharset).
	ContentLanguage(contentLanguage).
	TelstraApiVersion(telstraApiVersion).
	CreateTrialNumbersRequest(*createTrialNumbersRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)

```

### Fetch all Free Trial Numbers

> :information_source: **Only required for the free trial**

Fetch the Free Trial Number(s) currently assigned to your account. For more information,
please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#FetchyourFreeTrialNumbers>.

The function `freeTrialNumbers.getAll` can be used to retrieve registered destinations.

It takes no arguments.

It returns the list of phone numbers that have been registered.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.FreeTrialNumbersAPI.GetTrialNumbers(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

// Print out the entire oauthResult object
fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)

```

## Virtual Number

Gives you a dedicated mobile number tied to an application which
enables you to receive replies from your customers. For more information,
please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#VirtualNumbers>.

### Assign Virtual Number

When a recipient receives your message, you can choose whether they'll see a
Virtual Number or senderName (paid plans only) in the from field.
If you want to use a Virtual Number, use this function to assign one.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#AssignaVirtualNumber>.

The function `virtualNumbers.assign` can be used to create a subscription.

It takes a object with following properties as argument:

-   `replyCallbackUrl` (optional): The URL that replies to the Virtual Number will be posted to.
-   `tags` (optional): Create your own tags and use them to fetch, sort and report on your Virtual Numbers through our other endpoints.
    You can assign up to 10 tags per number.

It returns an object with the following properties:

-   `virtualNumber`: The Virtual Number assigned to your account.
-   `lastUse`: The last time the Virtual Number was used to send a message.
-   `replyCallbackUrl`: The URL that replies to the Virtual Number will be posted to.
-   `tags`: Any customisable tags assigned to the Virtual Number.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
tags := []string{
	"reprehenderit",
	"Excepteur non labore",
	"labore in consequat culpa",
	"qui voluptate",
	"n",
	"incididunt aliqua tempor",
	"incididunt dolor Lorem",
	"adipisicing aliquip elit eiusm",
	"consequat id sunt enim",
	"co",
}
assignNumberRequest := msg_sdk.NewAssignNumberRequest()
assignNumberRequest.SetReplyCallbackUrl("http://www.example.com")
assignNumberRequest.SetTags(tags)
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.VirtualNumbersAPI.AssignNumber(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	AssignNumberRequest(*assignNumberRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)

```

### Fetch a Virtual Number

Fetch the tags, replyCallbackUrl and lastUse date for a Virtual Number.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#FetchaVirtualNumber>.

The function `virtualNumbers.get` can be used to get the details of a Virtual Number.

It takes the following arguments:

-   `virtualNumber`: The Virtual Number assigned to your account.

It returns an object with the following properties:

-   `virtualNumber`: The Virtual Number assigned to your account.
-   `lastUse`: The last time the Virtual Number was used to send a message.
-   `replyCallbackUrl`: The URL that replies to the Virtual Number will be posted to.
-   `tags`: Any customisable tags assigned to the Virtual Number.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.VirtualNumbersAPI.GetVirtualNumber(context.Background(), virtualNumber).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)

```

### Fetch all Virtual Numbers

Fetch all Virtual Numbers currently assigned to your account.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#FetchallVirtualNumbers>.

The function `virtualNumbers.getAll` can be used to get the all virtual numbers associated to your account.

It takes an object with following prperties as argument:

-   `limit` (optional): Tell us how many results you want us to return, up to a maximum of 50.
-   `offset` (optional): Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.
-   `filter` (optional): Filter your Virtual Numbers by tag or by number.

It returns an object with the following properties:

-   `virtualNumbers`: A list of Virtual Numbers assigned to your account.
-   `paging`: Paging information.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.VirtualNumbersAPI.GetNumbers(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)
```

### Update a Virtual Number

Update a virtual number attributes. For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#UpdateaVirtualNumber>.

The function `virtualNumbers.update` can be used to update a virtual number.

It takes an object with following properties as argument:

-   `virtualNumber`: The Virtual Number assigned to your account.
-   `updateData` (optional):
    -   `reply_callback_url` (optional): The URL that replies to the Virtual Number will be posted to.
    -   `tags` (optional): Create your own tags and use them to fetch, sort and report on your Virtual Numbers through our other endpoints.
        You can assign up to 10 tags per number.

It returns an object with the following properties:

-   `virtualNumber`: The Virtual Number assigned to your account.
-   `lastUse`: The last time the Virtual Number was used to send a message.
-   `replyCallbackUrl`: The URL that replies to the Virtual Number will be posted to.
-   `tags`: Any customisable tags assigned to the Virtual Number.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
updateNumberRequest := msg_sdk.NewUpdateNumberRequest()
updateNumberRequest.SetReplyCallbackUrl("http://www.example.com")
tags := []string{
	"minim qui",
	"commodo",
	"nostrud laborum minim",
	"nulla proident ut voluptat",
	"et consectetur dolor",
	"est amet cillum",
	"exercitation",
	"non occaecat cupidatat Duis",
	"adipisicing",
	"ea aliqua incididunt",
}
updateNumberRequest.SetTags(tags)		
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.VirtualNumbersAPI.UpdateNumber(context.Background(), virtualNumber).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	UpdateNumberRequest(*updateNumberRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
fmt.Printf("resp err: %+v\n", err)

```

### Delete Virtual Number

Delete the a virtual number. For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#DeleteaVirtualNumber>.

The function `virtualNumbers.get` can be used to unassign a Virtual Number.

It takes the following arguments:

-   `virtualNumber`: The Virtual Number assigned to your account.

It returns nothing.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

httpRes, err := apiClient.VirtualNumbersAPI.DeleteNumber(context.Background(), virtualNumber).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)

```


### Fetch all Recipient Optouts list

Fetch any mobile number(s) that have opted out of receiving messages
from a Virtual Number assigned to your account.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Fetchallrecipientoptoutslist>.

The function `telstra.messaging.virtual_number.get_optouts` can be used to
get the list of mobile numbers that have opted out of receiving messages
from a virtual number associated to your account.
It takes the following arguments:

- `virtual_number`: The Virtual Number assigned to your account.
- `limit`: Tell us how many results you want us to return, up to a maximum of 50.
- `offset`: Use the offset to navigate between the response results.
  An offset of 0 will display the first page of results, and so on.

Raises `telstra.messaging.exceptions.VirtualNumbersError` if anything goes wrong.

It returns an object with the following
properties:

- `recipient_optouts`: A list of recipient optouts.
- `paging`: Paging information.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.VirtualNumbersAPI.GetRecipientOptouts(context.Background(), virtualNumber).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)

```


## Message

Send and receive messages. For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Messages>.

### Send Message

Send a message to a mobile number, or to multiple mobile numbers.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#SendanSMSorMMS>.

The function `messages.send` can be used to send a message.

It takes an object with following properties as argument:

-   `to`: The destination address, expected to be a phone number of the form `+614XXXXXXXX` or `04XXXXXXXX`.
-   `from`: This will be either one of your Virtual Numbers or your senderName.
-   `messageContent` (Either one of messageContent or multimedia is required): The content of the message.
-   `multimedia` (Either one of messageContent or multimedia is required): MMS multimedia content.
-   `retryTimeout` (optional): How many minutes you asked the server to keep trying to send the message.
-   `scheduleSend` (optional): The time (in Central Standard Time) the message is scheduled to send.
-   `deliveryNotification` (optional): If set to true, you will receive a notification to the statusCallbackUrl when your
    SMS or MMS is delivered (paid feature).
-   `statusCallbackUrl` (optional): The URL the API will call when the status of the message changes.
-   `tags` (optional): Any customisable tags assigned to the message.

The type `TMultimedia`can be used to build an mms payload. It has following properties:

-   `type`: The content type of the attachment, for example <TMultimediaContentType.IMAGE_GIF>.
-   `fileName` (optional): Optional field, for example `image.png`.
-   `payload`: The payload of an mms encoded as base64.

It returns an object with the following properties:

-   `messageId`: Use this UUID with our other endpoints to fetch, update or delete the message.
-   `status`: The status will be either queued, sent, delivered or expired.
-   `to`: The recipient's mobile number(s).
-   `from`: This will be either one of your Virtual Numbers or your senderName.
-   `messageContent`: The content of the message.
-   `multimedia`: The multimedia content of the message (MMS only).
-   `retryTimeout`: How many minutes you asked the server to keep trying to send the message.
-   `scheduleSend`: The time (in Central Standard Time) a message is scheduled to send.
-   `deliveryNotification`: If set to true, you will receive a notification to the
    statusCallbackUrl when your SMS or MMS is delivered (paid feature).
-   `statusCallbackUrl`: The URL the API will call when the status of the message changes.
-   `tags`: Any customisable tags assigned to the message.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

sendMessagesSlice := []string{"0400000001", "0400000002"}
sendMessagesRequestTo := msg_sdk.SendMessagesRequestTo{
	ArrayOfString: &sendMessagesSlice,
	String:        new(string),
}

sendMessagesFrom := "0428180739"
sendMessagesRequest := msg_sdk.NewSendMessagesRequest(sendMessagesRequestTo, sendMessagesFrom)
setMessageContent := "Hello customer, this is from CBA to confirme your offer!"
sendMessagesRequest.SetMessageContent(setMessageContent)
sendMessagesRequest.SetDeliveryNotification(false)
sendMessagesRequest.SetRetryTimeout(10)
sendMessagesRequest.SetStatusCallbackUrl("http://www.example.com")

tags := []string{
	"ip",
	"deserunt exercitation",
	"id mollit magna proident ipsum",
	"consequat proident",
	"u",
	"nulla ve",
	"deserunt proident",
	"deserunt nulla id esse",
	"laboris velit",
	"pr",
}

sendMessagesRequest.SetTags(tags)

resp, httpRes, err := apiClient.MessagesAPI.SendMessages(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	SendMessagesRequest(*sendMessagesRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)

```

### Get a Message

Use the messageId to fetch a message that's been sent from/to
your account within the last 30 days.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Fetchamessage>.

The function `messages.get` can be used to retrieve the a message.

It takes the following arguments:

-   `messageId`: Unique identifier for the message.

It returns an object with the following properties:

-   `messageId`: Use this UUID with our other endpoints to fetch, update or delete the message.
-   `status`: The status will be either queued, sent, delivered or expired.
-   `createTimestamp`: The time you submitted the message to the queue for sending.
-   `sentTimestamp`: The time the message was sent from the server.
-   `receivedTimestamp`: The time the message was received by the recipient's device.
-   `to`: The recipient's mobile number(s).
-   `from`: This will be either one of your Virtual Numbers or your senderName.
-   `messageContent`: The content of the message.
-   `multimedia`: The multimedia content of the message (MMS only).
-   `direction`: Direction of the message (outgoing or incoming).
-   `retryTimeout`: How many minutes you asked the server to keep trying to send the message.
-   `scheduleSend`: The time (in Central Standard Time) the message is scheduled to send.
-   `deliveryNotification`: If set to true, you will receive a notification to the statusCallbackUrl when your SMS or MMS
    is delivered (paid feature).
-   `statusCallbackUrl`: The URL the API will call when the status of the message changes.
-   `queuePriority`: The priority assigned to the message.
-   `tags`: Any customisable tags assigned to the message.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.MessagesAPI.GetMessageById(context.Background(), messageId).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
```

### Get all Messages

Fetch messages that have been sent from/to your account in the last 30 days.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Fetchallsent/receivedmessages>.

The function `messages.getAll` can be used to fetch all messages.

It takes an object with following properties as argument:

-   `limit` (optional): Tell us how many results you want us to return, up to a maximum of 50.
-   `offset` (optional): Use the offset to navigate between the response results. An offset of 0 will display the first page of results, and so on.
-   `filter` (optional): Filter your Virtual Numbers by tag or by number.

It returns an object with the following properties:

-   `messages`: List of all messages.
-   `paging`: Paging information.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.MessagesAPI.GetMessages(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
		
```

### Update a Message

Update a message that's scheduled for sending, you can change any of
the below parameters, as long as the message hasn't been sent yet.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Updateamessage>.

The function `messages.send` can be used to send a message.

It takes an object with following properties as argument:

-   `messageId`: Use this UUID with our other endpoints to fetch, update or delete the message.
-   `to`: The destination address, expected to be a phone number of the form `+614XXXXXXXX` or `04XXXXXXXX`.
-   `from`: This will be either one of your Virtual Numbers or your senderName.
-   `messageContent` (Either one of messageContent or multimedia is required): The content of the message.
-   `multimedia` (Either one of messageContent or multimedia is required): MMS multimedia content.
-   `retryTimeout` (optional): How many minutes you asked the server to keep trying to send the message.
-   `scheduleSend` (optional): The time (in Central Standard Time) the message is scheduled to send.
-   `deliveryNotification` (optional): If set to true, you will receive a notification to the statusCallbackUrl when your
    SMS or MMS is delivered (paid feature).
-   `statusCallbackUrl` (optional): The URL the API will call when the status of the message changes.
-   `tags` (optional): Any customisable tags assigned to the message.

The type `TMultimedia`can be used to build an mms payload. It has following properties:

-   `type`: The content type of the attachment, for example <TMultimediaContentType.IMAGE_GIF>.
-   `fileName` (optional): Optional field, for example `image.png`.
-   `payload`: The payload of an mms encoded as base64.

The dataclass `telstra.messaging.message.Multimedia` can be used to build
a mms payload. It takes the following arguments:

-   `type`: The content type of the attachment, for example `image/png`.
-   `filename` (optional): Optional field, for example `image.png`.
-   `payload`: The payload of an mms encoded as base64.

It returns an object with the following properties:

-   `messageId`: Use this UUID with our other endpoints to fetch, update or delete the message.
-   `status`: The status will be either queued, sent, delivered or expired.
-   `to`: The recipient's mobile number(s).
-   `from`: This will be either one of your Virtual Numbers or your senderName.
-   `messageContent`: The content of the message.
-   `multimedia`: The multimedia content of the message (MMS only).
-   `retryTimeout`: How many minutes you asked the server to keep trying to send the message.
-   `scheduleSend`: The time (in Central Standard Time) a message is scheduled to send.
-   `deliveryNotification`: If set to true, you will receive a notification to the
    statusCallbackUrl when your SMS or MMS is delivered (paid feature).
-   `statusCallbackUrl`: The URL the API will call when the status of the message changes.
-   `tags`: Any customisable tags assigned to the message.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)	
	
configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"
updateMessageByIdRequest := msg_sdk.NewUpdateMessageByIdRequest("0400000001", "0428180739")
setMessageContent := "Ut veniam in ipsum exercitation"
updateMessageByIdRequest.SetMessageContent(setMessageContent)
updateMessageByIdRequest.SetDeliveryNotification(false)
updateMessageByIdRequest.SetRetryTimeout(10)
updateMessageByIdRequest.SetStatusCallbackUrl("http://www.example.com")

tags := []string{
	"ip",
	"deserunt exercitation",
	"id mollit magna proident ipsum",
	"consequat proident",
	"u",
	"nulla ve",
	"deserunt proident",
	"deserunt nulla id esse",
	"laboris velit",
	"pr",
}
updateMessageByIdRequest.SetTags(tags)

resp, httpRes, err := apiClient.MessagesAPI.UpdateMessageById(context.Background(), messageId).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	UpdateMessageByIdRequest(*updateMessageByIdRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
```

### Update Message Tags

Update message tags, you can update them even after your message has been delivered.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Updatemessagetags>.

The function `messages.updateTags` can be used to update message tags.

It takes the following arguments:

-   `messageId`: Unique identifier for the message.
-   `tags` (optional): Any customisable tags assigned to the message.

It returns nothing.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"
tags := []string{"marketing", "SMS"}
updateMessageTagsRequest := msg_sdk.NewUpdateMessageTagsRequest(tags)

httpRes, err := apiClient.MessagesAPI.UpdateMessageTags(context.Background(), messageId).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	UpdateMessageTagsRequest(*updateMessageTagsRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
```

### Delete a Message

Delete a scheduled message, but hasn't yet sent.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Deleteamessage>.

The function `messages.delete` can be used to delete a message.

It takes the following arguments:

-   `messageId`: Unique identifier for the message.

It returns nothing.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"

httpRes, err := apiClient.MessagesAPI.DeleteMessageById(context.Background(), messageId).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
```

## Reports

Create and fetch reports. For more information, please see here:
<https://dev.telstra.com/content/messaging-api-v3#tag/reports>.

### Request a Messages Report

Request a CSV report of messages (both incoming and outgoing)
that have been sent to/from your account within the last three months.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Submitarequestforamessagesreport>.

The function `reports.create` can be used to create a report.

It takes the following arguments:

-   `startDate`: Set the time period you want to generate a report for by typing the start date (inclusive) here.
    Note that we only retain data for three months, so please ensure your startDate is not more than three months old.
    Use ISO format(yyyy-mm-dd), e.g. "2019-08-24".
-   `endDate`: Type the end date (inclusive) of your reporting period here.
    Your endDate must be a date in the past, and less than three months from your startDate.
    Use ISO format(yyyy-mm-dd), e.g. "2019-08-24".
-   `reportCallbackUrl` (optional): The callbackUrl where notification is sent when report is ready for download.
-   `filter` (optional): Filter report messages by:
    tag - use one of the tags assigned to your message(s)
    number - either the Virtual Number used to send the message,
    or the Recipient Number the message was sent to.

It returns an object with the following properties:

-   `reportId`: Use this UUID with our other endpoints to fetch the report.
-   `reportCallbackUrl`: If you provided a reportCallbackUrl in your request, it will be returned here.
-   `reportStatus`: The status of the report. It will be either:
    -   queued – the report is in the queue for generation.
    -   completed – the report is ready for download.
    -   failed – the report failed to generate, please try again.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"
messageReportRequest := msg_sdk.NewMessagesReportRequest("2023-09-01", "2023-09-05")

resp, httpRes, err := apiClient.ReportsAPI.MessagesReport(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	MessagesReportRequest(*messageReportRequest).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
```



### Fetch a specific report

Use the report_id to fetch a download link for a report generated.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#FetchaReport>.

The function `reports.get` can be used to retrieve
the a report download link. It takes the following arguments:

-   `reportId`: Unique identifier for the report.

It returns an object with the following properties:

-   `reportId`: Use this UUID with our other endpoints to fetch the report.
-   `reportStatus`: The status of the report.
-   `reportUrl`: Download link to download the CSV file.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"
reportId := "2be7b580-4c34-11ee-a651-ad71114ff6eb"

resp, httpRes, err := apiClient.ReportsAPI.GetReport(context.Background(), reportId).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
```

### Fetch all reports

Fetch details of all reports recently generated for your account.
Use it to check the status of a report, plus fetch the report ID,
status, report type and expiry date.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#Fetchallreports>.

The function `reports.getAll` can be used to fetch all reports.

It doesn't take any arguments.

It returns a list of objects with the following properties:

-   `reportId`: Use this UUID with our other endpoints to fetch the report.
-   `reportStatus`: The status of the report.
-   `reportType`: The type of report generated.
-   `reportExpiry`: The expiry date of your report. After this date, you will be unable to download your report.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)	

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
accept := "application/json"
acceptCharset := "utf-8"
contentType := "application/json"
contentLanguage := "en-au"
telstraApiVersion := "3.1.0"


resp, httpRes, err := apiClient.ReportsAPI.GetReports(context.Background()).
	Accept(accept).
	AcceptCharset(acceptCharset).
	Authorization(authorization).
	ContentLanguage(contentLanguage).
	ContentType(contentType).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)
```

## Health Check

### Get operational status of the messaging service

Check the operational status of the messaging service.
For more information, please see here:
<https://dev.telstra.com/docs/messaging-api/apiReference/apiReferenceOverviewEndpoints?version=3.x#HealthCheck>.

The function `healthCheck.get` can be used to get the status.

It takes no arguments.

It returns an object with following properties:

-   `status`: Denotes the status of the services Up/Down.

For example:

```go

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	msg_sdk "github.com/telstra/MessagingAPI-SDK-Go"
)

configuration := msg_sdk.NewConfiguration()
apiClient := msg_sdk.NewAPIClient(configuration)

clientId := "YOUR CLIENT ID"
clientSecret := "YOUR CLIENT SECRET"
grantType := "client_credentials"
scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

oauthResult, resp, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
	ClientId(clientId).
	ClientSecret(clientSecret).
	GrantType(grantType).
	Scope(scope).
	Execute()

accessToken := *oauthResult.AccessToken
authorization := "Bearer " + accessToken
telstraApiVersion := "3.1.0"

resp, httpRes, err := apiClient.HealthCheckAPI.HealthCheck(context.Background()).
	Authorization(authorization).
	TelstraApiVersion(telstraApiVersion).
	Execute()

fmt.Printf("httpRes Result: %+v\n", httpRes)
fmt.Printf("resp Result: %+v\n", resp)


```
