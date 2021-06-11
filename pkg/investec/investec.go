package investec

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	Base = "https://openapi.investec.com/"
	Auth = "identity/v2/oauth2/token"
)

const (
	GrantType = "client_credentials"
	Scope     = "accounts"
)

type AuthCredentials struct {
	ClientID     string
	ClientSecret string
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

type InvestecAccount struct {
}

type Transaction struct {
}

type InvestecApi struct {
	client      *http.Client
	credentials AuthCredentials
	auth        *AuthResponse
}

func NewOpenApiClient(credentials AuthCredentials) *InvestecApi {
	client := InvestecApi{
		credentials: credentials,
		client:      &http.Client{},
	}
	return &client
}

func (c *InvestecApi) Authenticate() (AuthResponse, error) {

	ret := AuthResponse{}

	data := url.Values{}
	data.Add("grant_type", GrantType)
	data.Add("scope", Scope)
	data.Add("client_id", c.credentials.ClientID)
	data.Add("client_secret", c.credentials.ClientSecret)

	req, _ := http.NewRequest("POST", Base+Auth, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return ret, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&ret)
	if err != nil {
		return ret, err
	}
	c.auth = &ret

	return ret, nil
}

func (api *InvestecApi) GetAccounts() []InvestecAccount {
	accounts := make([]InvestecAccount, 0)
	return accounts
}

func (api *InvestecApi) GetAccountTransactions(from time.Time, to time.Time) []Transaction {
	accounts := make([]Transaction, 0)

	return accounts
}
