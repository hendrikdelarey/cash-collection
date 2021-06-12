package investec

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InvestecAccount struct {
	AccountID     string `json: "accountId"`
	AccountNumber string `json: "accountNumber"`
	AccountName   string `json: "accountName"`
	ReferenceName string `json: "referenceName"`
	ProductName   string `json: "productName"`
}

type InvestecAccounts struct {
	Accounts []InvestecAccount `json "accounts"`
}

func (c *InvestecApi) GetAccounts() (InvestecAccounts, error) {
	accounts := &InvestecAccounts{}

	err := c.refreshTokenIfRequired()
	if err != nil {
		return *accounts, err
	}

	req, _ := http.NewRequest("GET", Base+Accounts, nil)
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.auth.TokenType, c.auth.AccessToken))
	req.Header.Add("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return *accounts, err
	}

	defer res.Body.Close()

	apiResp := &ApiResponse{Data: accounts}

	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return *accounts, err
	}

	if accounts, ok := apiResp.Data.(*InvestecAccounts); ok {
		return *accounts, nil
	}
	return *accounts, fmt.Errorf("invalid response received from API")
}
