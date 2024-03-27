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

	clientId := "YOUR CLIENT ID"
	clientSecret := "YOUR CLIENT SECRET"
	authorization, _ := GetAuthorization(apiClient, clientId, clientSecret)

	t.Run("Test FreeTrialNumbersAPIService CreateTrialNumbers", func(t *testing.T) {

		trialNumbers := []string{"0400000001", "0400000002"}
		createTrialNumbersRequestFreeTrialNumbers := ArrayOfStringAsCreateTrialNumbersRequestFreeTrialNumbers(&trialNumbers)
		createTrialNumbersRequest := NewCreateTrialNumbersRequest(createTrialNumbersRequestFreeTrialNumbers)

		trialNumbersApi := apiClient.FreeTrialNumbersAPI.CreateTrialNumbers(context.Background())
		trialNumbersApi.authorization = &authorization

		resp, httpRes, err := trialNumbersApi.CreateTrialNumbers(*createTrialNumbersRequest)

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

	t.Run("Test FreeTrialNumbersAPIService GetTrialNumbers", func(t *testing.T) {

		//t.Skip("skip test")
		trialNumbersApi := apiClient.FreeTrialNumbersAPI.GetTrialNumbers(context.Background())
		trialNumbersApi.authorization = &authorization

		resp, httpRes, err := trialNumbersApi.GetTrialNumbers()

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
