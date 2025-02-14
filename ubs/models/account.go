package models

type Account struct {
	AccountNumber        string        `json:"accountNumber,omitempty"`
	IBAN                 string        `json:"iban,omitempty"`
	From                 string        `json:"from,omitempty"`
	To                   string        `json:"to,omitempty"`
	StartBalance         string        `json:"startBalance,omitempty"`
	EndBalance           string        `json:"endBalance,omitempty"`
	Currency             string        `json:"curency,omitempty"`
	NumberOfTransactions string        `json:"numberOfTrnsactions,omitempty"`
	Transactions         []Transaction `json:"transactions,omitempty"`
}

func (a *Account) AddData(index int, data []string) {
	switch index {
	case 0:
		a.AccountNumber = data[1]
	case 1:
		a.IBAN = data[1]
	case 2:
		a.From = data[1]
	case 3:
		a.To = data[1]
	case 4:
		a.StartBalance = data[1]
	case 5:
		a.EndBalance = data[1]
	case 6:
		a.Currency = data[1]
	case 7:
		a.NumberOfTransactions = data[1]
	case 8: //skip its empty
		return
	default:
		a.Transactions = append(a.Transactions, Transaction{
			EndDate:        data[0],
			EndTime:        data[1],
			BookingDate:    data[2],
			ValueDate:      data[3],
			Currency:       data[4],
			Credit:         data[5],
			Debit:          data[6],
			SingleAmount:   data[7],
			Balance:        data[8],
			TransactionsNr: data[9],
			Description1:   data[10],
			Description2:   data[11],
			Description3:   data[12],
			Footnotes:      data[13],
		})
	}
}
