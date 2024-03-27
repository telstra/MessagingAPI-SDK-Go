/*
Messaging API v3.4.3

Testing VirtualNumbersAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TelstraMessaging_VirtualNumbersAPIService(t *testing.T) {

	t.Skip("skip test")

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)
	virtualNumber := "0428180739"

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"
	authorization, _ := getAuthorization(apiClient, clientId, clientSecret)

	t.Run("Test VirtualNumbersAPIService AssignNumber", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

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

		assignNumberRequest := NewAssignNumberRequest()
		assignNumberRequest.SetReplyCallbackUrl("http://www.example.com")
		assignNumberRequest.SetTags(tags)

		vnApi := apiClient.VirtualNumbersAPI.AssignNumber(context.Background())
		vnApi.authorization = &authorization

		resp, httpRes, err := vnApi.AssignNumber(*assignNumberRequest)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		fmt.Printf("resp err: %+v\n", err)
		assert.Equal(t, 201, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService DeleteNumber", func(t *testing.T) {

		vnApi := apiClient.VirtualNumbersAPI.DeleteNumber(context.Background(), virtualNumber)
		vnApi.authorization = &authorization

		httpRes, err := vnApi.DeleteNumber()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		assert.Equal(t, 204, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)

	})

	t.Run("Test VirtualNumbersAPIService GetNumbers", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		vnApi := apiClient.VirtualNumbersAPI.GetNumbers(context.Background())
		vnApi.authorization = &authorization

		resp, httpRes, err := vnApi.GetNumbers()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		fmt.Printf("resp err: %+v\n", err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService GetRecipientOptouts", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test
		vnApi := apiClient.VirtualNumbersAPI.GetRecipientOptouts(context.Background(), virtualNumber)
		vnApi.authorization = &authorization

		resp, httpRes, err := vnApi.GetRecipientOptouts()

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

	t.Run("Test VirtualNumbersAPIService GetVirtualNumber", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		vnApi := apiClient.VirtualNumbersAPI.GetVirtualNumber(context.Background(), virtualNumber)
		vnApi.authorization = &authorization

		resp, httpRes, err := vnApi.GetVirtualNumber()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		fmt.Printf("resp err: %+v\n", err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService UpdateNumber", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		updateNumberRequest := NewUpdateNumberRequest()
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

		vnApi := apiClient.VirtualNumbersAPI.UpdateNumber(context.Background(), virtualNumber)
		vnApi.authorization = &authorization

		resp, httpRes, err := vnApi.UpdateNumber(*updateNumberRequest)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("httpRes Result: %+v\n", httpRes)
		fmt.Printf("resp Result: %+v\n", resp)
		fmt.Printf("resp err: %+v\n", err)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

}
