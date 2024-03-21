/*
Messaging API v3.4.3

Testing HealthCheckAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TelstraMessaging_HealthCheckAPIService(t *testing.T) {

	t.Skip("skip test")

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)
	telstraApiVersion := "3.1.0"
	grantType := "client_credentials"
	scope := "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write"

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"

	oauthResult, _, err := apiClient.AuthenticationAPI.AuthToken(context.Background()).
		ClientId(clientId).
		ClientSecret(clientSecret).
		GrantType(grantType).
		Scope(scope).
		Execute()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	accessToken := *oauthResult.AccessToken
	authorization := "Bearer " + accessToken

	t.Run("Test HealthCheckAPIService HealthCheck", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.HealthCheckAPI.HealthCheck(context.Background()).
			Authorization(authorization).
			TelstraApiVersion(telstraApiVersion).
			Execute()

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

}
