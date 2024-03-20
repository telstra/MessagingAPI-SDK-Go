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

	accept := "application/json"
	acceptCharset := "utf-8"
	contentType := "application/json"
	contentLanguage := "en-au"
	telstraApiVersion := "3.1.0"
	virtualNumber := "0428180739"

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
		assert.Equal(t, 201, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService DeleteNumber", func(t *testing.T) {

		httpRes, err := apiClient.VirtualNumbersAPI.DeleteNumber(context.Background(), virtualNumber).
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

	t.Run("Test VirtualNumbersAPIService GetNumbers", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService GetRecipientOptouts", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test VirtualNumbersAPIService GetVirtualNumber", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

}
