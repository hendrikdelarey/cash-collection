package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/hendrikdelarey/cash-collection/pkg/investec"
)

func findPaymentsMadeWithReferences(client *investec.InvestecApi, references map[string]string) {
	accounts, err := client.GetAccounts()
	if err != nil {
		fmt.Print(err.Error())
	}

	for _, acc := range accounts.Accounts {
		transaction, err := client.GetAllAccountTransactions(acc.AccountID, time.Now().AddDate(0, -1, 0), time.Now())
		if err != nil {
			fmt.Print(err.Error())
		}

		for _, trans := range transaction.Transactions {
			if trans.Type == "CREDIT" {
				// since EFT payments can cause extra strings to be added to a description we need to check all the words in the description

				s := strings.ToUpper(trans.Description)
				replacer := strings.NewReplacer(",", "", ".", "", ";", "")
				s = replacer.Replace(s)

				words := strings.Fields(s)
				for _, word := range words {
					// see if a reference is found
					if val, ok := references[word]; ok {
						// a reference is found
						fmt.Print(val)

						// break out of the inner loop as there should only be one reference per transaction
						break
					}
				}
			}
		}
		fmt.Print(len(transaction.Transactions))
	}

}
