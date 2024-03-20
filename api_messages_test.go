/*
Messaging API v3.4.3

Testing MessagesAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_TelstraMessaging_MessagesAPIService(t *testing.T) {

	t.Skip("skip test")

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	accept := "application/json"
	acceptCharset := "utf-8"
	contentType := "application/json"
	contentLanguage := "en-au"
	telstraApiVersion := "3.1.0"
	messageId := "cd8cbda0-4c33-11ee-966e-7bcedb926a5d"

	grantType := "client_credentials"
	scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"

	oauthResult, _, _ := apiClient.AuthenticationAPI.AuthToken(context.Background()).
		ClientId(clientId).
		ClientSecret(clientSecret).
		GrantType(grantType).
		Scope(scope).
		Execute()

	accessToken := *oauthResult.AccessToken
	authorization := "Bearer " + accessToken

	t.Run("Test MessagesAPIService DeleteMessageById", func(t *testing.T) {

		httpRes, err := apiClient.MessagesAPI.DeleteMessageById(context.Background(), messageId).
			Accept(accept).
			AcceptCharset(acceptCharset).
			Authorization(authorization).
			ContentLanguage(contentLanguage).
			ContentType(contentType).
			TelstraApiVersion(telstraApiVersion).
			Execute()

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		assert.Equal(t, 204, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)

	})

	t.Run("Test MessagesAPIService GetMessageById", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test MessagesAPIService GetMessages", func(t *testing.T) {

		resp, httpRes, err := apiClient.MessagesAPI.GetMessages(context.Background()).
			Accept(accept).
			AcceptCharset(acceptCharset).
			Authorization(authorization).
			ContentLanguage(contentLanguage).
			ContentType(contentType).
			TelstraApiVersion(telstraApiVersion).
			Reverse(true).
			StartTime(time.Date(2024, 2, 23, 5, 10, 6, 0, time.UTC)).
			EndTime(time.Date(2024, 2, 29, 5, 10, 6, 0, time.UTC)).
			Execute()

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test MessagesAPIService SendMessages", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		sendMessagesSlice := []string{"0400000001", "0400000002"}
		sendMessagesRequestTo := SendMessagesRequestTo{
			ArrayOfString: &sendMessagesSlice,
			String:        new(string),
		}

		sendMessagesFrom := "0428180739"
		sendMessagesRequest := NewSendMessagesRequest(sendMessagesRequestTo, sendMessagesFrom)
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
		assert.Equal(t, 201, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test MessagesAPIService UpdateMessageById", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		updateMessageByIdRequest := NewUpdateMessageByIdRequest("0400000001", "0428180739")

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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test MessagesAPIService UpdateMessageTags", func(t *testing.T) {

		tags := []string{"marketing", "SMS"}
		updateMessageTagsRequest := NewUpdateMessageTagsRequest(tags)

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
		assert.Equal(t, 204, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)

	})

}
