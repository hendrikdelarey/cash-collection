package investec

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
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

	c.tokenExpireTime = time.Now().Add(time.Duration(ret.ExpiresIn) * time.Second)

	return ret, nil
}

func (c *InvestecApi) refreshTokenIfRequired() error {
	if time.Now().After(c.tokenExpireTime) {
		_, err := c.Authenticate()
		if err != nil {
			return err
		}
	}
	return nil
}
