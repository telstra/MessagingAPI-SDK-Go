/*Messaging API v3.4.3

Testing ReportsAPIService

*/

package TelstraMessaging

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func String(s string) *string {
	return &s
}

func Test_TelstraMessaging_ReportsAPIService(t *testing.T) {
	t.Skip("skip test")
	configuration := NewConfiguration()
	apiClient := NewAPIClient(configuration)

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"
	authorization, _ := GetAuthorization(apiClient, clientId, clientSecret)

	t.Run("Test ReportsAPIService GetReport", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test
		reportId := "c22d2050-eb37-11ee-ad5e-3d9263246ccb"

		reportsApi := apiClient.ReportsAPI.GetReport(context.Background(), reportId)
		reportsApi.authorization = &authorization

		resp, httpRes, err := reportsApi.GetReport()

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

		//t.Skip("skip test") // remove to run test
		reportsApi := apiClient.ReportsAPI.GetReports(context.Background())
		reportsApi.authorization = &authorization

		resp, httpRes, err := reportsApi.GetReports()

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

		//t.Skip("skip test") // remove to run test

		now := time.Now()
		threeDaysAgo := now.AddDate(0, 0, -3)
		oneDayAgo := now.AddDate(0, 0, -1)
		threeDaysAgoFormatted := threeDaysAgo.Format("2006-01-02")
		oneDayAgoFormatted := oneDayAgo.Format("2006-01-02")
		messageReportRequest := NewMessagesReportRequest(threeDaysAgoFormatted, oneDayAgoFormatted)

		reportsApi := apiClient.ReportsAPI.MessagesReport(context.Background())
		reportsApi.authorization = &authorization

		resp, httpRes, err := reportsApi.MessagesReport(*messageReportRequest)

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
