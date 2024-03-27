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

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"
	authorization, _ := GetAuthorization(apiClient, clientId, clientSecret)

	t.Run("Test HealthCheckAPIService HealthCheck", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		healthCheckApi := apiClient.HealthCheckAPI.HealthCheck(context.Background(), authorization)
		resp, httpRes, err := healthCheckApi.HealthCheck()

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
