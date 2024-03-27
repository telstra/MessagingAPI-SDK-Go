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

	messageId := "00f30ac0-ebb5-11ee-b9c0-5f8692f4ff9d"

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"
	authorization, _ := GetAuthorization(apiClient, clientId, clientSecret)

	t.Run("Test MessagesAPIService DeleteMessageById", func(t *testing.T) {
		//t.Skip("skip test") // remove to run test
		messagesApi := apiClient.MessagesAPI.DeleteMessageById(context.Background(), messageId)
		messagesApi.authorization = &authorization

		httpRes, err := messagesApi.DeleteMessageById()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		assert.Equal(t, 204, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)

	})

	t.Run("Test MessagesAPIService GetMessageById", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		messagesApi := apiClient.MessagesAPI.GetMessageById(context.Background(), messageId)
		messagesApi.authorization = &authorization

		resp, httpRes, err := messagesApi.GetMessageById()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test MessagesAPIService GetMessages", func(t *testing.T) {

		messagesApi := apiClient.MessagesAPI.GetMessages(context.Background())
		messagesApi.authorization = &authorization
		reverse := true
		messagesApi.reverse = &reverse
		startTime := time.Now().AddDate(0, 0, -5).UTC()
		messagesApi.startTime = &startTime
		endTime := time.Now().AddDate(0, 0, -1).UTC()
		messagesApi.endTime = &endTime

		resp, httpRes, err := messagesApi.GetMessages()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

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

		sendMessagesRequest.SetScheduleSend(time.Now().AddDate(0, 0, 5).UTC())

		sendMessagesRequest.SetTags(tags)

		messagesApi := apiClient.MessagesAPI.SendMessages(context.Background())
		messagesApi.authorization = &authorization

		resp, httpRes, err := messagesApi.SendMessages(*sendMessagesRequest)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

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

		messagesApi := apiClient.MessagesAPI.UpdateMessageById(context.Background(), messageId)
		messagesApi.authorization = &authorization

		resp, httpRes, err := messagesApi.UpdateMessageById(*updateMessageByIdRequest)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

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

		messagesApi := apiClient.MessagesAPI.UpdateMessageTags(context.Background(), messageId)
		messagesApi.authorization = &authorization

		httpRes, err := messagesApi.UpdateMessageTags(*updateMessageTagsRequest)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		assert.Equal(t, 204, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)

	})

}
