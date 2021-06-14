package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hendrikdelarey/cash-collection/pkg/investec"
)

func getInvestecApiClient() *investec.InvestecApi {
	creds := investec.AuthCredentials{
		ClientID:     "DummyId",
		ClientSecret: "DummySecret",
	}
	return investec.NewOpenApiClient(creds)
}

func GetRecentInvestecTransactions(w http.ResponseWriter, r *http.Request) {
	var transactions []investec.Transaction

	client := getInvestecApiClient()

	accounts, err := client.GetAccounts()
	if err != nil {
		http.Error(w, "Error accessing Investec account", http.StatusBadRequest)
	}

	for _, acc := range accounts.Accounts {
		t, err := client.GetAllTransactionsFromPastDays(acc.AccountID, 30)
		if err != nil {
			http.Error(w, "Error accessing Investec account", http.StatusBadRequest)
		}

		transactions = append(transactions, t.Transactions...)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
