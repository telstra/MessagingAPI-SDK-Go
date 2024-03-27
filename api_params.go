/*
Messaging API v3.4.3

Request params for the Telstra Messaging API
*/
package TelstraMessaging

import "context"

type RequestParams struct {
	ContentType       string
	Accept            string
	AcceptCharset     string
	ContentLanguage   string
	TelstraApiVersion string
	GrantType         string
	Scope             string
}

func SetRequestParams() *RequestParams {
	return &RequestParams{
		ContentType:       "application/json",
		Accept:            "application/json",
		AcceptCharset:     "utf-8",
		ContentLanguage:   "en-au",
		TelstraApiVersion: "3.1.0",
		GrantType:         "client_credentials",
		Scope:             "free-trial-numbers:read free-trial-numbers:write messages:read messages:write virtual-numbers:read virtual-numbers:write reports:read reports:write",
	}
}

func GetAuthorization(apiClient *APIClient, clientId string, clientSecret string) (string, error) {
	authApi := apiClient.AuthenticationAPI.AuthToken(context.Background())
	authApi.clientId = &clientId
	authApi.clientSecret = &clientSecret

	oauthResult, _, err := authApi.AuthToken()
	if err != nil {
		return "", err
	}

	accessToken := *oauthResult.AccessToken
	return "Bearer " + accessToken, nil
}
