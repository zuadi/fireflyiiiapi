package api

import (
	"bytes"
	"encoding/json"
	"fireflyiiiapi/api/models"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Type string

const (
	All            Type = "all"
	Withdrawal     Type = "withdrawal"
	Withdrawals    Type = "withdrawals"
	Expense        Type = "expense"
	Deposit        Type = "deposit"
	Deposits       Type = "deposits"
	Income         Type = "income"
	Transfer       Type = "transfer"
	Transfers      Type = "transfers"
	OpeningBalance Type = "opening_balance"
	Reconciliation Type = "reconciliation"
	Special        Type = "special"
	Specials       Type = "specials"
	Default        Type = "default"
)

func (h *APIHandler) AllTransaction(typ Type) (*models.TransactionArray, error) {
	baseURL := fmt.Sprintf("%s/api/v1/transactions", h.Url)

	params := url.Values{}
	//params.Add("limit", "10")         // Number of items per page
	//params.Add("page", "1")           // Page number
	//params.Add("start", "2018-09-17") // Start date (YYYY-MM-DD)
	//params.Add("end", "2018-09-17")   // End date (YYYY-MM-DD)
	params.Add("type", string(typ)) // Filter by transaction type

	req, err := http.NewRequest("GET", fmt.Sprintf("%s?%s", baseURL, params.Encode()), nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.AccesToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	body, err := SendRequest(req)
	if err != nil {
		return nil, err
	}
	transactions := models.TransactionArray{}
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %s", err.Error())
	}

	return &transactions, nil
}

func (h *APIHandler) AddTransaction(typ Type, amount, description string) error {

	transactions := models.TransactionStore{}
	transactions.Transactions = make([]models.TransactionSplitStore, 0)
	transactions.Transactions = append(transactions.Transactions, models.TransactionSplitStore{
		Type:          string(typ),
		Amount:        amount,
		Description:   description,
		Date:          time.Now().Format("2006-01-02T15:04:05-07:00"),
		SourceId:      "1",
		DestinationId: "2",
		CategoryName:  "Food",
	})

	jsonData, err := json.Marshal(transactions)
	if err != nil {
		return err
	}
	fmt.Println(123, string(jsonData))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/v1/transactions", h.Url), bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %s", err.Error())
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", h.AccesToken))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	_, err = SendRequest(req)
	if err != nil {
		return err
	}

	return nil
}
