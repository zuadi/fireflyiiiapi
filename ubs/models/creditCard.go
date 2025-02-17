package models

type CreditCard struct {
	Transactions []CreditCardTransaction
}

func (c *CreditCard) AddData(index int, data []string) {
	switch index {
	case 0:
		return
	default:
		c.Transactions = append(c.Transactions, CreditCardTransaction{
			AccountNumber:   data[0],
			CardNUmber:      data[1],
			CardHolder:      data[2],
			BuyDate:         data[3],
			BookingText:     data[4],
			Branch:          data[5],
			Amount:          data[6],
			OriginalCurency: data[7],
			ExchangeRate:    data[8],
			Currency:        data[9],
			Credit:          data[10],
			Debit:           data[11],
			Booking:         data[12],
		})
	}
}
