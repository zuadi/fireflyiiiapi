package models

type Account struct {
	AccountNumber        string
	IBAN                 string
	From                 string
	To                   string
	StartBalance         string
	EndBalance           string
	Currency             string
	NumberOfTransactions string
	Transactions         []Transaction
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
	default:
		a.Transactions = append(a.Transactions, Transaction{
			EndDate:        data[0],
			EndTime:        data[1],
			BookDate:       data[2],
			ValudaDate:     data[3],
			Valutadatum:    data[4],
			Currency:       data[5],
			Credit:         data[6],
			Depit:          data[7],
			SingleAmount:   data[8],
			Balance:        data[9],
			TransactionsNr: data[10],
			Description1:   data[11],
			Description2:   data[12],
			Description3:   data[13],
			Fotenote:       data[14],
		})
	}
}
