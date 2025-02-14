package models

import (
	"fmt"
	"strings"
)

type AccountArray struct {
	Data []AccountRead `json:"data" validate:"required"`
	Meta Account       `json:"meta,omitempty"`
}

// give back id looking first up in iban, accountnumber, bic and name
func (a *AccountArray) GetAccountId(input string) (string, error) {
	account := strings.ReplaceAll(input, " ", "")
	for _, acc := range a.Data {
		if account == acc.Attributes.IBAN {
			return acc.Id, nil
		} else if account == acc.Attributes.AccountNumber {
			return acc.Id, nil
		} else if account == acc.Attributes.BIC {
			return acc.Id, nil
		} else if input == acc.Attributes.Name {
			return acc.Id, nil
		}
	}
	return "", fmt.Errorf("account '%s' not found", input)
}
