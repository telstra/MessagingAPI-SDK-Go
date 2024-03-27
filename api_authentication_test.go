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

		authApi := apiClient.AuthenticationAPI.AuthToken(context.Background(), clientId, clientSecret)
		oauthResult, resp, err := authApi.AuthToken()

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Access Token: %s\n", *oauthResult.AccessToken)
		fmt.Printf("OAuth Result: %+v\n", oauthResult)
		assert.NoError(t, err, "Error encountered: %v", err)
		assert.NotNil(t, resp, "Response should not be nil")
		assert.NotEmpty(t, oauthResult.AccessToken, "Access token should not be empty")

	})

}
