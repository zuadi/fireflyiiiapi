package ubs

import (
	"bufio"
	"encoding/csv"
	configModels "fireflyiiiapi/config/models"
	"fireflyiiiapi/ubs/interfaces"
	"fireflyiiiapi/ubs/models"
	"io"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type CSVReader struct {
	config      configModels.UBS
	Accounts    []interfaces.Accounts `json:"accounts,omitempty"`
	CreditCards []interfaces.Accounts `json:"creditCards,omitempty"`
}

func NewCSVReader(config configModels.UBS) *CSVReader {
	return &CSVReader{
		config: config,
	}
}

func (r *CSVReader) Read(files ...string) error {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return err
		}

		// Convert ISO-8859-1 (Latin-1) to UTF-8
		readerconv := charmap.ISO8859_1.NewDecoder().Reader(f)
		reader := csv.NewReader(bufio.NewReader(readerconv))

		reader.Comma = rune(r.config.CSV.Seperator[0])
		reader.FieldsPerRecord = r.config.CSV.FieldLimit

		var i int
		var account interfaces.Accounts

		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				if strings.Contains(err.Error(), "bare \" in non-quoted-field") {
					continue
				}
				return err
			}

			if i == 0 {
				re := regexp.MustCompile(`[^\x20-\x7E]+`) // Matches non-printable ASCII
				if re.ReplaceAllString(row[0], "") == r.config.CSV.FirstItemForBankaccount {
					account = &models.Account{}
					r.Accounts = append(r.Accounts, account)
				} else if re.ReplaceAllString(row[0], "") == r.config.CSV.FirstItemForCreditCard {
					account = &models.CreditCard{}
					r.CreditCards = append(r.CreditCards, account)
				}
			}

			account.AddData(i, row)

			i++
		}
	}

	return nil
}

func (r *CSVReader) GetAllAccountTransaction() []interfaces.Accounts {
	return r.Accounts
}

func (r *CSVReader) GetAllCreditCardransaction() []interfaces.Accounts {
	return r.CreditCards
}
