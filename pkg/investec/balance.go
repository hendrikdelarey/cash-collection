package investec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AccountBalance struct {
	AccountID        string  `json: "accountId"`
	CurrentBalance   float32 `json: "currentBalance"`
	AvailableBalance string  `json: "availableBalance"`
	Currency         string  `json: "currency"`
}

func (c *InvestecApi) GetAccountBalance(accountId string) (AccountBalance, error) {
	balance := &AccountBalance{}

	err := c.refreshTokenIfRequired()
	if err != nil {
		return *balance, err
	}

	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/balance", Base+Accounts, accountId), nil)
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.auth.TokenType, c.auth.AccessToken))
	req.Header.Add("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return *balance, err
	}

	defer res.Body.Close()

	apiResp := &ApiResponse{Data: balance}

	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return *balance, err
	}

	if balance, ok := apiResp.Data.(*AccountBalance); ok {
		return *balance, nil
	}

	return *balance, fmt.Errorf("invalid response received from API")
}
