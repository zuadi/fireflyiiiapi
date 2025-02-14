package models

type Transaction struct {
	EndDate        string `json:"endDate,omitempty"`
	EndTime        string `json:"endTime,omitempty"`
	BookingDate    string `json:"bokingDate,omitempty"`
	ValueDate      string `json:"valueDate,omitempty"`
	Currency       string `json:"curency,omitempty"`
	Credit         string `json:"credit,omitempty"`
	Debit          string `json:"debit,omitempty"`
	SingleAmount   string `json:"singleAmount,omitempty"`
	Balance        string `json:"balance,omitempty"`
	TransactionsNr string `json:"transactionNr,omitempty"`
	Description1   string `json:"description1,omitempty"`
	Description2   string `json:"description2,omitempty"`
	Description3   string `json:"description3,omitempty"`
	Footnotes      string `json:"footenotes,omitempty"`
}

type CreditCardTransaction struct {
	AccountNumber   string `json:"accountNumber,omitempty"`
	CardNumber      string `json:"cardNumber,omitempty"`
	CardHolder      string `json:"cardHolder,omitempty"`
	PurcaseDate     string `json:"purcaseDate,omitempty"`
	BookingText     string `json:"bookingText,omitempty"`
	Branch          string `json:"branch,omitempty"`
	Amount          string `json:"amount,omitempty"`
	OriginalCurency string `json:"originalCurrency,omitempty"`
	ExchangeRate    string `json:"exchangeRate,omitempty"`
	Currency        string `json:"currency,omitempty"`
	Credit          string `json:"credit,omitempty"`
	Debit           string `json:"debit,omitempty"`
	Booking         string `json:"booking,omitempty"`
}
