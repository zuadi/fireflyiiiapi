package ubs

import (
	"encoding/csv"
	configModels "fireflyiiiapi/config/models"
	"fireflyiiiapi/ubs/interfaces"
	"fireflyiiiapi/ubs/models"
	"io"
	"os"
	"regexp"
	"strings"
)

type CSVReader struct {
	config     configModels.UBS
	Account    []interfaces.Accounts
	CreditCard []interfaces.Accounts
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

		reader := csv.NewReader(f)
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
					r.Account = append(r.Account, account)
				} else if re.ReplaceAllString(row[0], "") == r.config.CSV.FirstItemForCreditCard {
					account = &models.CreditCard{}
					r.CreditCard = append(r.CreditCard, account)
				}
			}

			account.AddData(i, row)

			i++
		}
	}

	return nil
}
