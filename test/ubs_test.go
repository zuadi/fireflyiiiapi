package test

import (
	"fireflyiiiapi/config"
	"fireflyiiiapi/ubs"
	"testing"
)

func TestReadUBSCSV(t *testing.T) {
	c, err := config.ReadConfig("../dist/config.yml")
	if err != nil {
		t.Error(err)
	}

	reader := ubs.NewCSVReader(c.UBS)
	err = reader.Read("../dist/transactionsSparen.csv", "../dist/transactionsCreditCard.csv")
	if err != nil {
		t.Error(err)
	}
}
