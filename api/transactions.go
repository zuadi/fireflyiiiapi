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

func (h *APIHandler) AddTransactionDeposit(accountId, amount, category, description string) error {
	if len(h.Deposits.Transactions) == 0 {
		h.Deposits.Transactions = make([]models.TransactionSplitStore, 0)
	}

	h.Deposits.Transactions = append(h.Deposits.Transactions, models.TransactionSplitStore{
		Type:          "deposit",
		Amount:        amount,
		Description:   description,
		DestinationId: accountId,
		Date:          time.Now().Format("2006-01-02T15:04:05-07:00"),
		CategoryName:  category,
	})
	return nil
}

func (h *APIHandler) AddTransactionWithdrawals(accountId, amount, category, description string) error {
	if len(h.Withdrawals.Transactions) == 0 {
		h.Withdrawals.Transactions = make([]models.TransactionSplitStore, 0)
	}

	h.Withdrawals.Transactions = append(h.Withdrawals.Transactions, models.TransactionSplitStore{
		Type:         "withdrawal",
		Amount:       amount,
		Description:  description,
		SourceId:     accountId,
		Date:         time.Now().Format("2006-01-02T15:04:05-07:00"),
		CategoryName: category,
	})
	return nil
}

func (h *APIHandler) AddTransactionTransfer(srcAccountId, dstAccountId, amount, category, description string) error {
	if len(h.Transfers.Transactions) == 0 {
		h.Transfers.Transactions = make([]models.TransactionSplitStore, 0)
	}

	h.Transfers.Transactions = append(h.Transfers.Transactions, models.TransactionSplitStore{
		Type:          "transfer",
		Amount:        amount,
		Description:   description,
		SourceId:      srcAccountId,
		DestinationId: dstAccountId,
		Date:          time.Now().Format("2006-01-02T15:04:05-07:00"),
		CategoryName:  category,
	})
	return nil
}

func (h *APIHandler) AddTransactionReconsiliation(srcAccountId, dstAccountId, amount, category, description string) error {
	if len(h.Reconsiliations.Transactions) == 0 {
		h.Reconsiliations.Transactions = make([]models.TransactionSplitStore, 0)
	}

	h.Reconsiliations.Transactions = append(h.Reconsiliations.Transactions, models.TransactionSplitStore{
		Type:          "reconsiliation",
		Amount:        amount,
		Description:   description,
		SourceId:      srcAccountId,
		DestinationId: dstAccountId,
		Date:          time.Now().Format("2006-01-02T15:04:05-07:00"),
		CategoryName:  category,
	})
	return nil
}

func (h *APIHandler) AddTransactionOpeningBalance(srcAccountId, dstAccountId, amount, category, description string) error {
	if len(h.OpeningBalances.Transactions) == 0 {
		h.OpeningBalances.Transactions = make([]models.TransactionSplitStore, 0)
	}

	h.OpeningBalances.Transactions = append(h.OpeningBalances.Transactions, models.TransactionSplitStore{
		Type:          "opening_balance",
		Amount:        amount,
		Description:   description,
		SourceId:      srcAccountId,
		DestinationId: dstAccountId,
		Date:          time.Now().Format("2006-01-02T15:04:05-07:00"),
		CategoryName:  category,
	})
	return nil
}

func (h *APIHandler) SendTransactions() error {
	transactions := models.TransactionStore{}
	for i := 0; i < 4; i++ {
		switch i {
		case 0:
			if len(h.Deposits.Transactions) > 0 {
				transactions = h.Deposits
			} else {
				continue
			}
		case 1:
			if len(h.Withdrawals.Transactions) > 0 {
				transactions = h.Withdrawals
			} else {
				continue
			}
		case 2:
			if len(h.Transfers.Transactions) > 0 {
				transactions = h.Transfers
			} else {
				continue
			}
		case 3:
			if len(h.Reconsiliations.Transactions) > 0 {
				transactions = h.Reconsiliations
			} else {
				continue
			}
		case 4:
			if len(h.OpeningBalances.Transactions) > 0 {
				transactions = h.OpeningBalances
			} else {
				continue
			}
		}

		jsonData, err := json.Marshal(transactions)
		if err != nil {
			return err
		}

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
	}
	return nil
}
