package test

import (
	"fireflyiiiapi/api"
	"fmt"
	"testing"
)

const url = "https://firefly.tecamino.com"
const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiZjQ1NTJlNTkwMDAyNDE3OTljZmYzMDcxMjRkN2VhYjJjNWJhMjliMDRiNDJjYjdmZDBmNTI4ZjkzYzU3ZjY2MmUzMzIwN2MwN2RmN2NjMTkiLCJpYXQiOjE3MzkyMDc1MjYuOTA1NjM1LCJuYmYiOjE3MzkyMDc1MjYuOTA1NjUsImV4cCI6MTc3MDc0MzUyNi41MjEwODMsInN1YiI6IjEiLCJzY29wZXMiOltdfQ.dAZfSfh0r4DEf3ueU51fuUUBaIVeWy7kct9B9bwagMiaG14Z9O-bJtKsBDOaXskMGcSQNB76k-R6IcASUKkqDuYDWTJgoCyVS284Jn4okRpAroozGwK9TeMDAMfzflBU7Ab1VBRUGmhX7GwYTTsmkEOzWF7mmGuMyy4wMLWqwX74a7hdrPXNiwwM2cewTNra6Yk1Zf80yzxZq_CA7v9ZBXFSi5HlVCfoQbIDHZY8BvSiPuERx_rEIkizlT10bRact9U1yAhsWhJ5l09zCuzUJBB2CmtigIovhfimYSNajF1D8UM-nH7Rm7lhz2R5TOoll8PsC6qE_yqerZ3W6KDkwuv65hQIduenOce8ZQoIEhdImtiXttrg9IMnlodVqVDcQ8aU6gt2U6Na5G36eHBsB2oHttXdcjf7J-T_hX3KdnA54bJhtiH3RDGgY-MDd_P6avDJEhPE8H13NpGnsdsPR2gUEGciLuwDbeaHZO0gm5VWODoB0sE2TBdWjdFL9RlzJ0FrQfwdqTxvGZZEaAo-pjC--DxSg8Rd2PWdCA8bqkB9Dx6yojmvL3EUPtSQhHctoE4DlQ8xLMe9w3plT8lJh-ZK4hLD44Bo9k4jmlnUc0ZANC63j5PTPToKUpfWjhHhcqEbWNpgap878fpiSFoVHL_HU9VGZnu1P_VNkmD6XmA"

func TestWrongToken(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token[1:])
	if _, err := fireflyAPI.GetSystemInfo(); err != nil {
		t.Log(err)
	}

}

func TestWrongUrl(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url[5:], token)
	if _, err := fireflyAPI.GetSystemInfo(); err != nil {
		t.Log(err)
	}
}

func TestGetSystemInfo(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	data, err := fireflyAPI.GetSystemInfo()
	if err != nil {
		t.Error(err)
	}

	t.Log(data)
}

func TestGetAboutUser(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	data, err := fireflyAPI.GetUser()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(data)
}

func TestGetAllTransactions(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	data, err := fireflyAPI.AllTransaction(api.Income)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data.Meta)
	for i, item := range data.Data {
		t.Log(i)
		t.Log(item)
		t.Log(item.Type)
	}
}

func TestGetAccounts(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	_, err := fireflyAPI.GetAccounts()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetAccountId(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	accounts, err := fireflyAPI.GetAccounts()
	if err != nil {
		t.Error(err)
		return
	}

	for _, o := range accounts.Data {
		fmt.Println(1111, o)
	}
	t.Log("test account not found")
	if _, err := accounts.GetAccountId("test"); err != nil {
		t.Log(err.Error())
	}
	t.Log("test account found")
	if id, err := accounts.GetAccountId("CH14 0023 5235 6639 7040W"); err != nil {
		t.Error(err)
	} else {
		t.Log("id:", id)
	}
}

func TestAddTransaction(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	accounts, err := fireflyAPI.GetAccounts()
	if err != nil {
		t.Error(err)
		return
	}

	incomeId, err := accounts.GetAccountId("CH14 0023 5235 6639 7040 W")
	if err != nil {
		t.Error(err)
		return
	}

	savingId, err := accounts.GetAccountId("CH29 0023 5235 6639 70M1 U")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("add deposit")
	err = fireflyAPI.AddTransactionDeposit(incomeId, "123.25", "Income", "golang test")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("add withrawal")
	err = fireflyAPI.AddTransactionWithdrawals(incomeId, "25.45", "Food", "golang test 2")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("add transfer")
	err = fireflyAPI.AddTransactionTransfer(incomeId, savingId, "25.45", "Umbuchung", "golang test 3")
	if err != nil {
		t.Error(err)
		return
	}

	err = fireflyAPI.SendTransactions()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSaveToken(t *testing.T) {
	if err := api.SaveTokenToFile("../dist/secret.bin", token); err != nil {
		t.Error(err)
	}

	savedToken, err := api.ReadTokenToFile("../dist/secret.bin")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token == savedToken)
}
