package investec

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Transaction struct {
	AccountID       string  `json: "accountId"`
	Type            string  `json: "type"`
	TransactionType string  `json: "transactionType"`
	Status          string  `json: "status"`
	Description     string  `json: "description"`
	CardNumber      string  `json: "cardNumber"`
	PostedOrder     int     `json: "postedOrder"`
	PostingDate     string  `json: "postingDate"`
	ValueDate       string  `json: "postingDate"`
	ActionDate      string  `json: "postingDate"`
	TransactionDate string  `json: "postingDate"`
	Amount          float32 `json: "amount"`
	RunningBalance  float32 `json: "runningBalance"`
}

type AccountTransactions struct {
	Transactions []Transaction `json "transactions"`
}

func (c *InvestecApi) GetAccountTransactions(accountId string, from time.Time, to time.Time, transactionType string) (AccountTransactions, error) {
	transactions := &AccountTransactions{}

	err := c.refreshTokenIfRequired()
	if err != nil {
		return *transactions, err
	}

	endpoint := fmt.Sprintf("%s/%s/transactions?from=%s&to=%s", Base+Accounts, accountId, from.Format(time.RFC3339), to.Format(time.RFC3339))
	if transactionType != "" {
		endpoint += "&" + string(transactionType)
	}

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Authorization", fmt.Sprintf("%s %s", c.auth.TokenType, c.auth.AccessToken))
	req.Header.Add("Accept", "application/json")

	res, err := c.client.Do(req)
	if err != nil {
		return *transactions, err
	}

	defer res.Body.Close()

	apiResp := &ApiResponse{Data: transactions}

	err = json.NewDecoder(res.Body).Decode(&apiResp)
	if err != nil {
		return *transactions, err
	}

	if transactions, ok := apiResp.Data.(*AccountTransactions); ok {
		return *transactions, nil
	}

	return *transactions, fmt.Errorf("invalid response received from API")
}

func (c *InvestecApi) GetAllAccountTransactions(accountId string, from time.Time, to time.Time) (AccountTransactions, error) {
	return c.GetAccountTransactions(accountId, from, to, "")
}

func (c *InvestecApi) GetAllTransactionsFromPastDays(accountID string, days int) (AccountTransactions, error) {
	return c.GetAccountTransactions(accountID, time.Now().AddDate(0, 0, -days), time.Now(), "")
}

func (c *InvestecApi) GetAllTransactionsForPastDuration(accountID string, duration time.Duration) (AccountTransactions, error) {
	return c.GetAccountTransactions(accountID, time.Now().Add(-duration), time.Now(), "")
}

func (c *InvestecApi) GetAllTransactionsFromPastDaysForTransactionType(accountID string, days int, transactiontype string) (AccountTransactions, error) {
	return c.GetAccountTransactions(accountID, time.Now().AddDate(0, 0, -days), time.Now(), transactiontype)
}

func (c *InvestecApi) GetAllTransactionsForPastDurationForTransactionType(accountID string, duration time.Duration, transactiontype string) (AccountTransactions, error) {
	return c.GetAccountTransactions(accountID, time.Now().Add(-duration), time.Now(), transactiontype)
}
