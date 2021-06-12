package investec

import (
	"net/http"
	"time"
)

const (
	Base     = "https://openapi.investec.com/"
	Auth     = "identity/v2/oauth2/token"
	Accounts = "za/pb/v1/accounts"
)

type ApiResponse struct {
	Data  interface{} `json: "data"`
	Links interface{} `json "links"`
	Meta  interface{} `json "meta"`
}

type InvestecApi struct {
	auth            *AuthResponse
	client          *http.Client
	credentials     AuthCredentials
	tokenExpireTime time.Time
}

func NewOpenApiClient(credentials AuthCredentials) *InvestecApi {
	client := InvestecApi{
		credentials:     credentials,
		client:          &http.Client{},
		tokenExpireTime: time.Now(),
	}
	return &client
}
