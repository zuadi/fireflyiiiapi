package test

import (
	"encoding/json"
	"fireflyiiiapi/config"
	"fireflyiiiapi/ubs"
	"os"
	"testing"
)

func TestReadUBSCSV(t *testing.T) {

	writeFiles := 0 // 0=no 1=first 2 =second 12=both

	c, err := config.ReadConfig("../dist/config.yml")
	if err != nil {
		t.Error(err)
	}

	reader := ubs.NewCSVReader(c.UBS)
	err = reader.Read("../dist/transactions.csv") //"../dist/transactionsSparen.csv", "../dist/transactionsCreditCard.csv")
	if err != nil {
		t.Error(err)
	}

	if writeFiles == 1 || writeFiles == 12 {
		f, err := os.OpenFile("../dist/account.json", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		data, err := json.Marshal(reader.Accounts)
		if err != nil {
			t.Error(err)
		}

		if _, err := f.Write(data); err != nil {
			t.Error(err)
		}
	}

	if writeFiles == 2 || writeFiles == 12 {
		f, err := os.OpenFile("../dist/credit.json", os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			t.Error(err)
		}
		defer f.Close()

		data, err := json.Marshal(reader.CreditCards)
		if err != nil {
			t.Error(err)
		}

		if _, err := f.Write(data); err != nil {
			t.Error(err)
		}
	}
}
