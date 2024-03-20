/*
Messaging API v3.4.3

Testing AuthenticationAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TelstraMessaging_AuthenticationAPIService(t *testing.T) {

	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	t.Run("Test AuthenticationAPIService AuthToken", func(t *testing.T) {
		t.Skip("skip test") // remove to run test
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
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "Response should not be nil")
		assert.NotEmpty(t, oauthResult.AccessToken, "Access token should not be empty")

	})

}
