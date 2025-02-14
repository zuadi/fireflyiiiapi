package test

import (
	"fireflyiiiapi/config"
	"testing"
)

func TestReadConfig(t *testing.T) {
	_, err := config.ReadConfig("../dist/config.yml")
	if err != nil {
		t.Log(err)
	}
}
