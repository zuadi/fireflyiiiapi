package test

import (
	"fireflyiiiapi/api"
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
	err := fireflyAPI.GetAccounts()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestAddTransaction(t *testing.T) {
	fireflyAPI := api.NewAPIHandler(url, token)
	err := fireflyAPI.AddTransaction(api.Deposit, "123.25", "golang test")
	if err != nil {
		t.Error(err)
		return
	}
}
