/*
Messaging API v3.4.3

Testing FreeTrialNumbersAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TelstraMessaging_FreeTrialNumbersAPIService(t *testing.T) {

	t.Skip("skip test")

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	accept := "application/json"
	acceptCharset := "utf-8"
	contentType := "application/json"
	contentLanguage := "en-au"
	telstraApiVersion := "3.1.0"
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

	t.Run("Test FreeTrialNumbersAPIService CreateTrialNumbers", func(t *testing.T) {

		trialNumbers := []string{"0400000001", "0400000002"}
		createTrialNumbersRequestFreeTrialNumbers := ArrayOfStringAsCreateTrialNumbersRequestFreeTrialNumbers(&trialNumbers)
		createTrialNumbersRequest := NewCreateTrialNumbersRequest(createTrialNumbersRequestFreeTrialNumbers)

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
		assert.Equal(t, 201, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

	t.Run("Test FreeTrialNumbersAPIService GetTrialNumbers", func(t *testing.T) {

		//t.Skip("skip test")
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
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.NotEmpty(t, httpRes, "httpRes should not be empty")
		assert.NotNil(t, resp, "resp should not be nil")

	})

}
