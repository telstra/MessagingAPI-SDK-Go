/*Messaging API v3.4.3

Testing ReportsAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func String(s string) *string {
	return &s
}

func Test_TelstraMessaging_ReportsAPIService(t *testing.T) {

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

	t.Run("Test ReportsAPIService GetReport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test
		reportId := "2be7b580-4c34-11ee-a651-ad71114ff6eb"

		resp, httpRes, err := apiClient.ReportsAPI.GetReport(context.Background(), reportId).
			Accept(accept).
			AcceptCharset(acceptCharset).
			Authorization(authorization).
			ContentLanguage(contentLanguage).
			ContentType(contentType).
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

	t.Run("Test ReportsAPIService GetReports", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ReportsAPI.GetReports(context.Background()).
			Accept(accept).
			AcceptCharset(acceptCharset).
			Authorization(authorization).
			ContentLanguage(contentLanguage).
			ContentType(contentType).
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

	t.Run("Test ReportsAPIService MessagesReport", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		messageReportRequest := NewMessagesReportRequest("2023-09-01", "2023-09-05")

		resp, httpRes, err := apiClient.ReportsAPI.MessagesReport(context.Background()).
			Accept(accept).
			AcceptCharset(acceptCharset).
			Authorization(authorization).
			ContentLanguage(contentLanguage).
			ContentType(contentType).
			TelstraApiVersion(telstraApiVersion).
			MessagesReportRequest(*messageReportRequest).
			Execute()

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

}
