package models

type Transaction struct {
	EndDate        string
	EndTime        string
	BookDate       string
	ValudaDate     string
	Valutadatum    string
	Currency       string
	Credit         string
	Depit          string
	SingleAmount   string
	Balance        string
	TransactionsNr string
	Description1   string
	Description2   string
	Description3   string
	Fotenote       string
}

type CreditCardTransaction struct {
	AccountNumber   string
	CardNUmber      string
	CardHolder      string
	BuyDate         string
	BookingText     string
	Branch          string
	Amount          string
	OriginalCurency string
	ExchangeRate    string
	Currency        string
	Credit          string
	Debit           string
	Booking         string
}
